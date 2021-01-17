package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	strcase "github.com/hori-ryota/go-strcase"
)

func GenerateRequests(
	dstDir string,
	packageName string,
	doc KabusAPIDocument,
) error {
	out := new(bytes.Buffer)

	err := template.Must(template.New("").Funcs(map[string]interface{}{
		"PathWithParam": func(method MethodDef) string {
			if len(method.PathParams) == 0 {
				return `"` + method.HTTPPath + `"`
			}
			pat := method.HTTPPath
			paramNames := make([]string, 0, len(method.PathParams))
			for _, p := range method.PathParams {
				paramName := strcase.ToLowerCamel(p.Name)
				paramNames = append(paramNames, paramName)
				pat = strings.ReplaceAll(pat, "{"+paramName+"}", "%v")
			}
			return `fmt.Sprintf("` + pat + `", ` + strings.Join(paramNames, `, `) + `)`
		},
		"DescriptionToDoc": func(s string) string {
			return toDocComment(s)
		},
		"StructDoc": func(s StructDef) string {
			return toDocComment(s.Description)
		},
		"PropertyDoc": func(p PropertyDef) string {
			return toDocComment(p.Description)
		},
		"EnumDoc": func(v EnumValue) string {
			return toDocComment(strings.ReplaceAll(v.Description, "<br>", ""))
		},
		"IsEnum": func(t TypeDef) bool {
			_, ok := t.(EnumDef)
			return ok
		},
		"IsRef": func(t TypeDef) bool {
			_, ok := t.(RefDef)
			return ok
		},
		"IsPointer": func(p PropertyDef) bool {
			if _, ok := p.Type.(ArrayDef); ok {
				return false
			}
			return !p.Required
		},
		"AddPropertyToURLQuery": func(p PropertyDef) string {
			v := fmt.Sprintf(`req.%s`, strcase.ToUpperCamel(p.Name))
			f := fmt.Sprintf(`q.Add("%s", %%s)`, p.Name)
			if !p.Required {
				v = "*" + v
				f = fmt.Sprintf(`if req.%s != nil {
					q.Add("%s", %%s)
				}`,
					strcase.ToUpperCamel(p.Name),
					p.Name,
				)
			}

			t := p.Type

			if enum, ok := t.(EnumDef); ok {
				v = fmt.Sprintf(`%s(%s)`, enum.BaseType.ToGoType(), v)
				t = enum.BaseType
			}

			switch t.ToGoType() {
			case "string":
				// noop
			case "int32":
				v = fmt.Sprintf(`strconv.Itoa(int(%s))`, v)
			case "float64":
				v = fmt.Sprintf(`strconv.FormatFloat(%s, f, -1, 64)`, v)
			default:
				panic(fmt.Errorf("unknown property type %s", p.Type.ToGoType()))
			}
			return fmt.Sprintf(f, v)
		},
		"NormalizeEnumName": normalizeEnumName,
		"ToLower":           strings.ToLower,
		"ToLowerCamel":      strcase.ToLowerCamel,
		"ToUpperCamel":      strcase.ToUpperCamel,
	}).Parse(`
// Code generated by internal/codegen; DO NOT EDIT.
package {{.PackageName}}

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

{{- range .Methods}}

{{- range .PathParams}}
{{- if IsEnum .Type}}
{{DescriptionToDoc .Description -}}
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
{{- end}}{{/* with .Type */}}
{{- end}}{{/* if IsEnum .Type */}}
{{- end}}{{/* range .PathParams */}}

{{- if not (IsRef .InputType)}}
{{with .InputType}}

{{/* enum of input type */}}
{{- range .Properties}}
{{- if IsEnum .Type}}
{{DescriptionToDoc .Description -}}
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
{{- end}}{{/* with .Type */}}
{{- end}}{{/* if IsEnum .Type */}}
{{- end}}{{/* range .Properties */}}
{{/* end of enum of input type */}}

// {{.Name}} is definition of {{.Name}}.
type {{.Name}} struct {
	{{- range .Properties}}
	{{PropertyDoc . -}}
	{{ToUpperCamel .Name}} {{if IsPointer .}}*{{end}}{{.Type.ToGoType}} `+"`"+`json:"{{.Name}}"`+"`"+`
	{{- end}}
}
{{- end}}
{{- end}}

{{DescriptionToDoc .Summary -}}
{{DescriptionToDoc .Description -}}
func (c Client) {{.Name}}(
	ctx context.Context,
	{{- range .PathParams}}
	{{ToLowerCamel .Name}} {{.Type.ToGoType}},
	{{- end}}
	{{- if .InputType}}
	req {{.InputType.ToGoType}},
	{{- end}}
) ({{.OutputType.ToGoType}}, error) {
	pat := {{PathWithParam .}}
	res := {{.OutputType.ToGoType}}{}

	{{- if eq (ToLower .HTTPMethod) "get"}}
	q := url.Values{}
	{{- if .InputType}}
	{{- range .InputType.Properties}}
	{{AddPropertyToURLQuery .}}
	{{- end}}{{/* range .InputType.Properties */}}
	{{- end}}{{/* if .InputType */}}
	err := c.{{ToLower .HTTPMethod}}Request(ctx, pat, q, &res)
	{{- else }}
	err := c.{{ToLower .HTTPMethod}}Request(ctx, pat, {{if .InputType}}req{{else}}nil{{end}}, &res)
	{{- end}}
	return res, err
}
{{- end}}{{/* range Methods */}}
	`)).Execute(out, map[string]interface{}{
		"PackageName": packageName,
		"Methods":     doc.Methods,
	})
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// NOTE: debug時はここで一度出力しておくとフォーマットが失敗するファイルでも出力ファイルを直接見れる（そしてコンパイルエラーなどが確認できる）ので便利
	if err := ioutil.WriteFile(filepath.Join(dstDir, "requests_gen.go"), out.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	formatted, err := FormatGoFileFromString(out.String())
	if err != nil {
		return fmt.Errorf("failed to format output: %w", err)
	}
	if err := ioutil.WriteFile(filepath.Join(dstDir, "requests_gen.go"), []byte(formatted), 0644); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}
	return nil
}