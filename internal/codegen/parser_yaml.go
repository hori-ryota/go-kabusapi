package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type YAMLDoc struct {
	Paths map[ /* path */ string]map[ /* method */ string]struct {
		// Tags        []string
		Summary     string
		Description string
		Parameters  []struct {
			In          string
			Name        string
			Description string
			Schema      YAMLSchemaDef
			Required    bool
		}
		RequestBody struct {
			Required bool
			Content  map[ /* content type */ string]struct {
				Schema struct {
					Ref string `yaml:"$ref"`
				}
			}
		} `yaml:"requestBody"`
		Responses map[ /* status code */ string]struct {
			Description string
			Content     map[ /* content type */ string]struct {
				Schema struct {
					Ref string `yaml:"$ref"`
				}
			}
		}
	}
	Components struct {
		Schemas YAMLSchemaDefs
	}
}

// bool or []string
type YAMLRequiredUnmarshaller struct {
	List []string
	Bool bool
}

func (m *YAMLRequiredUnmarshaller) UnmarshalYAML(node *yaml.Node) error {
	switch node.Kind {
	case 0:
		return nil
	case yaml.ScalarNode:
		var b bool
		if err := node.Decode(&b); err != nil {
			return fmt.Errorf("failed to decode in YAMLRequiredUnmarshaller as ScalarNode as bool: yamlNode: %+v: %w", node, err)
		}
		m.Bool = b
		return nil
	case yaml.SequenceNode:
		var list []string
		if err := node.Decode(&list); err != nil {
			return fmt.Errorf("failed to decode in YAMLRequiredUnmarshaller as SequenceNode as []string: yamlNode: %+v: %w", node, err)
		}
		m.List = list
		return nil
	default:
		return fmt.Errorf("invalid yaml node %+v", node)
	}
}

type YAMLSchemaDef struct {
	Name        string
	Type        string
	Required    YAMLRequiredUnmarshaller
	Description string
	Format      string
	Example     interface{}
	// if type is object
	Properties YAMLSchemaDefs
	// if type is array
	Items *YAMLSchemaDef
	Ref   RefDef `yaml:"$ref"`
}

type YAMLSchemaDefs []YAMLSchemaDef

func (d *YAMLSchemaDefs) UnmarshalYAML(node *yaml.Node) error {
	switch node.Kind {
	case 0:
		return nil
	case yaml.MappingNode:
		list := make(YAMLSchemaDefs, 0, len(node.Content)/2)
		for i := 0; i < len(node.Content); i += 2 {
			var v YAMLSchemaDef
			if err := node.Content[i+1].Decode(&v); err != nil {
				return fmt.Errorf("failed to decode in YAMLSchemaDefs.UnmarshalYAML with YAMLSchemaDef type: yamlNode: %+v: %w", node, err)
			}
			if err := node.Content[i].Decode(&v.Name); err != nil {
				return fmt.Errorf("failed to decode in YAMLSchemaDefs.UnmarshalYAML with YAMLSchemaDef.Name: yamlNode: %+v: %w", node, err)
			}
			list = append(list, v)
		}
		*d = list
		return nil
	default:
		return fmt.Errorf("invalid yaml node %+v", node)
	}
}
