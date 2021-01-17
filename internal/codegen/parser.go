package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"gopkg.in/yaml.v3"
)

type KabusAPIDocument struct {
	Methods     []MethodDef
	Definitions []DefinitionDef
}

type DefinitionDef struct {
	Name string
	Type TypeDef
}

func ParseKabusAPIDocument(r io.Reader) (KabusAPIDocument, error) {
	var y YAMLDoc
	if err := yaml.NewDecoder(r).Decode(&y); err != nil {
		return KabusAPIDocument{}, fmt.Errorf("failed to decode yaml: %w", err)
	}

	definitions := make([]DefinitionDef, len(y.Definitions))
	for i, d := range y.Definitions {
		if d.Name == "ErrorResponse" {
			// 特別対応
			var codeListText string
			for i, p := range d.Properties {
				if p.Name == "Message" {
					ind := strings.Index(p.Description, "|")
					codeListText = p.Description[ind:]
					d.Properties[i].Description = strings.TrimSpace(p.Description[:ind])
					d.Properties[i].Required = YAMLRequiredUnmarshaller{
						Bool: true,
					}
				}
			}
			for i, p := range d.Properties {
				if p.Name == "Code" {
					d.Properties[i].Description = p.Description + "\n" + strings.ReplaceAll(codeListText, "<br>", "")
					d.Properties[i].Required = YAMLRequiredUnmarshaller{
						Bool: true,
					}
				}
			}
		}
		t, err := YAMLObjectDefToTypeDef(d, nil)
		if err != nil {
			return KabusAPIDocument{}, fmt.Errorf("failed to parse definition: %+v: %w", d, err)
		}
		definitions[i] = DefinitionDef{
			Name: d.Name,
			Type: t,
		}
	}

	methods, err := ParseMethods(y)
	if err != nil {
		return KabusAPIDocument{}, fmt.Errorf("failed to parse methods: %w", err)
	}

	return KabusAPIDocument{
		Definitions: definitions,
		Methods:     methods,
	}, nil
}

func YAMLObjectDefToTypeDef(yd YAMLObjectDef, parentYD *YAMLObjectDef) (TypeDef, error) {
	if yd.Ref != "" {
		return yd.Ref, nil
	}

	if strings.Contains(yd.Description, "|-|-|") && yd.Name != "Message" {
		// enum
		return parseEnum(yd, parentYD.Name)
	}
	switch yd.Type {
	case "boolean":
		return BoolDef, nil
	case "string":
		return StringDef, nil
	case "integer":
		return Int32Def, nil
	case "number":
		return Float64Def, nil
	case "object":
		if yd.Name == "RegistList" {
			// fix bug of RegistList of UnregisterAllSuccess
			return ArrayDef{
				Elem: RefDef("RegistListItem"),
			}, nil
		}

		if parentYD != nil && parentYD.Type == "array" {
			yd.Name = parentYD.Name + "Item"
		}

		properties := make([]PropertyDef, len(yd.Properties))
		for i, yp := range yd.Properties {
			t, err := YAMLObjectDefToTypeDef(yp, &yd)
			if err != nil {
				return nil, fmt.Errorf("failed to YAMLObjectDefToTypeDef as object at property %d: %w", i, err)
			}
			if yp.Name == "Message" {
				// 特別対応
				yp.Description = strings.ReplaceAll(yp.Description, "<br>", "")
			}
			required := yp.Required.Bool || containsString(yd.Required.List, yp.Name)
			properties[i] = PropertyDef{
				Name:        yp.Name,
				Required:    required,
				Description: yp.Description,
				Type:        t,
			}
		}
		return StructDef{
			Name:        yd.Name,
			Description: yd.Description,
			Properties:  properties,
		}, nil
	case "array":
		t, err := YAMLObjectDefToTypeDef(*yd.Items, &yd)
		if err != nil {
			return nil, fmt.Errorf("failed to YAMLObjectDefToTypeDef as array: %w", err)
		}
		return ArrayDef{
			Elem: t,
		}, nil
	default:
		return nil, fmt.Errorf("unknown YAMLObjectDef.Type: %s", yd.Type)
	}
}

func containsString(ss []string, s string) bool {
	for _, t := range ss {
		if t == s {
			return true
		}
	}
	return false
}

func parseEnum(yd YAMLObjectDef, prefix string) (EnumDef, error) {
	scanner := bufio.NewScanner(strings.NewReader(yd.Description))
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "|-|-|") {
			break
		}
	}
	enums := make([]EnumValue, 0, 20)
	for scanner.Scan() {
		vv := strings.Split(scanner.Text(), "|")
		enums = append(enums, EnumValue{
			Name:        vv[2],
			Value:       vv[1],
			Description: scanner.Text(),
		})
	}
	baseType, err := YAMLObjectDefToTypeDef(YAMLObjectDef{
		Type: yd.Type,
	}, nil)
	if err != nil {
		return EnumDef{}, fmt.Errorf("failed to YAMLObjectDefToTypeDef as enum: %w", err)
	}
	switch yd.Name {
	// 特別対応が必要なものを個別処理
	case "PriceRangeGroup":
		es := make([]EnumValue, 0, len(enums))
		exists := make(map[string]bool, len(enums))
		for _, e := range enums {
			if exists[e.Value] {
				continue
			}
			es = append(es, EnumValue{
				Name:  e.Value,
				Value: e.Value,
			})
			exists[e.Value] = true
		}
		enums = es
	case "Code":
		exists := make(map[string]bool, len(enums))
		duplicated := make(map[string]bool, len(enums))
		for i := range enums {
			enums[i].Name = strings.ReplaceAll(enums[i].Name, "<br>", "")
			found := strings.Index(enums[i].Name, " - ")
			if found > 0 {
				enums[i].Name = enums[i].Name[:found]
			}
			if exists[enums[i].Name] {
				duplicated[enums[i].Name] = true
			}
			exists[enums[i].Name] = true
		}
		for i := range enums {
			if !duplicated[enums[i].Name] {
				continue
			}
			enums[i].Name = enums[i].Name + "_" + enums[i].Value
			enums[i].Description = strings.ReplaceAll(enums[i].Description, "<br>", "")
		}
	}
	return EnumDef{
		Prefix:   prefix,
		Name:     yd.Name,
		BaseType: baseType,
		Values:   enums,
	}, nil
}
