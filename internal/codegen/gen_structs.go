package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
)

func GenerateStructs(
	dstDir string,
	packageName string,
	doc KabusAPIDocument,
) error {
	out := new(bytes.Buffer)

	err := template.Must(template.New("").Funcs(map[string]interface{}{
		"StructDoc": func(s StructDef) string {
			return toDocComment(s.Description)
		},
		"PropertyDoc": func(p PropertyDef) string {
			return toDocComment(p.Description)
		},
		"EnumDoc": func(v EnumValue) string {
			return toDocComment(strings.ReplaceAll(v.Description, "<br>", ""))
		},
		"GetArrayElem": func(t TypeDef) TypeDef {
			return t.(ArrayDef).Elem
		},
		"IsPointer": func(p PropertyDef) bool {
			if _, ok := p.Type.(ArrayDef); ok {
				return false
			}
			return !p.Required
		},
		"IsEnum": func(t TypeDef) bool {
			_, ok := t.(EnumDef)
			return ok
		},
		"NormalizeEnumName": normalizeEnumName,
	}).Parse(`
// Code generated by internal/codegen; DO NOT EDIT.
package {{.PackageName}}

{{- range .Arrays}}

// {{.Name}} is array of {{.Type.Elem.ToGoType}}.
type {{.Name}} {{.Type.ToGoType}}
{{- end}}

{{- range .Structs}}

{{ if eq .Name "RegistList" -}}

{{- else -}}
// {{.Name}} is definition of {{.Name}}.
{{StructDoc . -}}
type {{.Name}} struct {
	{{- range .Properties}}
	{{PropertyDoc . -}}
	{{.Name}} {{if IsPointer .}}*{{end}}{{.Type.ToGoType}} `+"`"+`json:"{{.Name}}"`+"`"+`
	{{- end}}
}

{{- range .Properties}}
{{- if IsEnum .Type}}
{{PropertyDoc . -}}
{{with .Type -}}
type {{.ToGoType}} {{.BaseType.ToGoType}}
{{- $enumType := .ToGoType}}
{{- $baseType := .BaseType.ToGoType}}
const (
	{{- range .Values}}
	{{EnumDoc . -}}
	{{$enumType}}{{NormalizeEnumName .Name}} {{$enumType}} = {{if eq $baseType "string"}}"{{.Value}}"{{else}}{{.Value}}{{end}}
	{{- end}}
)

// P returns pointer of {{.ToGoType}}.
func (e {{.ToGoType}}) P() *{{.ToGoType}} {
	v := e
	return &v
}
{{- end}}{{/* with .Type */}}
{{- end}}{{/* if IsEnum .Type */}}
{{- end}}{{/* range .Properties */}}

{{- end}}{{/* if eq .Name "RegistList" */}}
{{- end}}{{/* range .Structs */}}
	`)).Execute(out, map[string]interface{}{
		"PackageName": packageName,
		"Arrays":      doc.TypedArrays(),
		"Structs":     doc.TypedStructs(),
	})
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// NOTE: debug時はここで一度出力しておくとフォーマットが失敗するファイルでも出力ファイルを直接見れる（そしてコンパイルエラーなどが確認できる）ので便利
	if err := ioutil.WriteFile(filepath.Join(dstDir, "structs_gen.go"), out.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	formatted, err := FormatGoFileFromString(out.String())
	if err != nil {
		return fmt.Errorf("failed to format output: %w", err)
	}
	if err := ioutil.WriteFile(filepath.Join(dstDir, "structs_gen.go"), []byte(formatted), 0644); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}
	return nil
}
