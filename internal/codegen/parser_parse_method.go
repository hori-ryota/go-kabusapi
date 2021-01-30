package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	strcase "github.com/hori-ryota/go-strcase"
)

func ParseMethods(y YAMLDoc) ([]MethodDef, error) {
	pathParamRegexp := regexp.MustCompile(`\{[^}].*\}`)

	methods := make([]MethodDef, 0, len(y.Paths)*4)
	for pat, mm := range y.Paths {
		if strings.Contains(pat, "/ranking") {
			// NOTE: oneofなどの経過を見守ってから実装に加える [【要望】oneofの仕様を避けたい、用途ごとにAPIを分割して欲しい · Issue \#234 · kabucom/kabusapi](https://github.com/kabucom/kabusapi/issues/234)
			continue
		}
		for httpMethod, yp := range mm {
			// def method type
			var methodName string
			if pathParamRegexp.MatchString(pat) {
				methodName = fmt.Sprintf(
					"%s%sOf",
					strcase.ToUpperCamel(httpMethod),
					strcase.ToUpperCamel(pathParamRegexp.ReplaceAllString(pat, "")),
				)
			} else {
				methodName = fmt.Sprintf(
					"%s%s",
					strcase.ToUpperCamel(httpMethod),
					strcase.ToUpperCamel(pat),
				)
			}

			for i, p := range yp.Parameters {
				yp.Parameters[i].Schema.Name = p.Name
				yp.Parameters[i].Schema.Description = p.Description
			}

			// parse PathParams
			pathParams := make([]PathParamDef, 0, len(yp.Parameters))
			for _, p := range yp.Parameters {
				if p.In != "path" {
					continue
				}
				t, err := YAMLSchemaDefToTypeDef(p.Schema, &YAMLSchemaDef{
					Name: methodName + "Param",
				})
				if err != nil {
					return nil, fmt.Errorf("failed to parse pathParam: %+v: %w", p, err)
				}
				pathParams = append(pathParams, PathParamDef{
					Name:        p.Name,
					Type:        t,
					Description: p.Description,
				})
			}

			// parse InputType
			var inputType TypeDef
			if httpMethod == "get" {
				inputTypeName := methodName + "Query"
				queryParams := make([]PropertyDef, 0, len(yp.Parameters))
				for _, p := range yp.Parameters {
					if p.In != "query" {
						continue
					}
					t, err := YAMLSchemaDefToTypeDef(p.Schema, &YAMLSchemaDef{
						Name: inputTypeName,
					})
					if err != nil {
						return nil, fmt.Errorf("failed to parse query: %+v: %w", p, err)
					}
					queryParams = append(queryParams, PropertyDef{
						Name:        p.Name,
						Type:        t,
						Required:    p.Required,
						Description: p.Description,
					})
				}
				if len(queryParams) > 0 {
					inputType = StructDef{
						Name:        inputTypeName,
						Properties:  queryParams,
						Description: fmt.Sprintf("%s is query of %s.", inputTypeName, methodName),
					}
				}
			} else {
				if c, ok := yp.RequestBody.Content["application/json"]; ok {
					inputType = RefDef(c.Schema.Ref)
				}
			}

			// parse OutputType
			outputType := RefDef(yp.Responses["200"].Content["application/json"].Schema.Ref)

			methods = append(methods, MethodDef{
				Name:        methodName,
				HTTPMethod:  httpMethod,
				HTTPPath:    pat,
				PathParams:  pathParams,
				InputType:   inputType,
				OutputType:  outputType,
				Summary:     yp.Summary,
				Description: yp.Description,
			})
		}
	}
	sort.Slice(methods, func(i, j int) bool {
		if methods[i].HTTPPath == methods[j].HTTPPath {
			return methods[i].HTTPMethod < methods[j].HTTPMethod
		}
		return methods[i].HTTPPath < methods[j].HTTPPath
	})
	return methods, nil
}
