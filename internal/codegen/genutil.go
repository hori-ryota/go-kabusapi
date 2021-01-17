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
	lines := strings.Split(text, "\n")
	ls := make([]string, 0, len(lines))
	for _, l := range lines {
		l = strings.TrimSuffix(l, "<br>")
		if strings.HasPrefix(l, "|") {
			ls = append(ls, strings.ReplaceAll(l, "<br>", " "))
			continue
		}
		ls = append(ls, strings.Split(l, "<br>")...)
	}
	lines = ls
	for i := range lines {
		lines[i] = "// " + lines[i]
	}
	return strings.Join(lines, "\n") + "\n"
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
