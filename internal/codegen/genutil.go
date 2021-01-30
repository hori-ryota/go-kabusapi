package main

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
)

// FormatGoFileFromString format Go file string.
func FormatGoFileFromString(s string) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", s, parser.ParseComments)
	if err != nil {
		return "", err
	}

	b := new(bytes.Buffer)
	if err := format.Node(b, fset, file); err != nil {
		return "", nil
	}
	return b.String(), nil
}

func toDocComment(text string) string {
	if text == "" {
		return ""
	}
	text = strings.TrimPrefix(text, `"`)
	text = strings.TrimSuffix(text, `"`)
	text = htmlTableToMarkdown(text)
	lines := strings.Split(text, "\n")
	ls := make([]string, 0, len(lines))
	for _, l := range lines {
		l = strings.TrimSpace(l)
		l = strings.TrimSuffix(l, "<br>")
		ls = append(ls, strings.Split(l, "<br>")...)
	}
	lines = ls
	for i := range lines {
		lines[i] = "// " + lines[i]
	}
	return strings.TrimSpace(strings.Join(lines, "\n")) + "\n"
}

func normalizeEnumName(name string) string {
	i := strings.Index(name, "<br>")
	if i > 0 {
		name = name[:i]
	}
	name = strings.ReplaceAll(name, "（", "_")
	name = strings.ReplaceAll(name, "）", "_")
	name = strings.ReplaceAll(name, "、", "_")
	name = strings.ReplaceAll(name, "・", "_")
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.ReplaceAll(name, "：", "_")
	name = strings.ReplaceAll(name, "。", "_")
	name = strings.TrimSuffix(name, "_")
	return name
}

func htmlTableToMarkdown(text string) string {
	lines := strings.Split(text, "\n")
	ls := make([]string, 0, len(lines))
	for i := 0; i < len(lines); i++ {
		l := lines[i]
		l = strings.TrimSpace(l)
		if l != "<table>" {
			if strings.Contains(l, "<tr>") {
				l = strings.ReplaceAll(l, "<tr><td>", "")
				l = strings.ReplaceAll(l, "</td>", "")
				l = strings.ReplaceAll(l, "</tr>", "")
				l = strings.ReplaceAll(l, "<td>", "|")
				ls = append(ls, "|"+l+"|")
				continue
			}
			ls = append(ls, l)
			continue
		}
		var tableHeader []string
		tableRows := make([][]string, 0, len(lines))
		for {
			i++
			if i == len(lines) {
				// NOTE: 閉じタグ漏れのバグ対応。直ったら消す
				// [【不具合】schema定義のenum系descriptionが複数行定義になっていない · Issue \#235 · kabucom/kabusapi](https://github.com/kabucom/kabusapi/issues/235#issuecomment-769665622)
				break
			}
			l := lines[i]
			l = strings.TrimSpace(l)
			if l == "</table>" {
				break
			}
			if l == "<thead>" || l == "</thead>" || l == "<tbody>" || l == "</tbody>" {
				continue
			}
			l = strings.ReplaceAll(l, "<br>", " ")
			if strings.Contains(l, "<th>") {
				l = strings.ReplaceAll(l, "<tr><th>", "")
				l = strings.ReplaceAll(l, "</th>", "")
				l = strings.ReplaceAll(l, "</tr>", "")
				tableHeader = strings.Split(l, "<th>")
				continue
			}
			l = strings.ReplaceAll(l, "<tr><td>", "")
			l = strings.ReplaceAll(l, "</td>", "")
			l = strings.ReplaceAll(l, "</tr>", "")
			tableRows = append(tableRows, strings.Split(l, "<td>"))
		}
		columnWidth := make([]int, len(tableHeader))
		countWidth := func(s string) int {
			width := 0
			for _, c := range strings.Split(s, "") {
				if len([]byte(c)) == 1 {
					width++
					continue
				}
				width += 2
			}
			return width
		}
		for j := range columnWidth {
			for _, l := range append(tableRows, tableHeader) {
				width := countWidth(l[j])
				if columnWidth[j] < width {
					columnWidth[j] = width
				}
			}
		}
		devider := make([]string, len(tableHeader))
		for j := range tableHeader {
			tableHeader[j] = tableHeader[j] + strings.Repeat(" ", columnWidth[j]-countWidth(tableHeader[j]))
			devider[j] = strings.Repeat("-", columnWidth[j])
		}
		ls = append(ls, "|"+strings.Join(tableHeader, "|")+"|")
		ls = append(ls, "|"+strings.Join(devider, "|")+"|")
		for _, row := range tableRows {
			for j := range row {
				row[j] = row[j] + strings.Repeat(" ", columnWidth[j]-countWidth(row[j]))
			}
			ls = append(ls, "|"+strings.Join(row, "|")+"|")
		}
	}
	return strings.Join(ls, "\n")
}
