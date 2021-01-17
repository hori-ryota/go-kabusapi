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
			In     string
			Schema struct {
				Ref string `yaml:"$ref"`
			}
			YAMLObjectDef `yaml:",inline"`
		}
		Responses map[ /* status code */ string]struct {
			Description string
			Schema      struct {
				Ref string `yaml:"$ref"`
			}
		}
	}
	Definitions YAMLObjectDefs
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

type YAMLObjectDef struct {
	Name        string
	Type        string
	Required    YAMLRequiredUnmarshaller
	Description string
	Format      string
	Example     interface{}
	// if type is object
	Properties YAMLObjectDefs
	// if type is array
	Items *YAMLObjectDef
	Ref   RefDef `yaml:"$ref"`
}

type YAMLObjectDefs []YAMLObjectDef

func (d *YAMLObjectDefs) UnmarshalYAML(node *yaml.Node) error {
	switch node.Kind {
	case 0:
		return nil
	case yaml.MappingNode:
		list := make(YAMLObjectDefs, 0, len(node.Content)/2)
		for i := 0; i < len(node.Content); i += 2 {
			var v YAMLObjectDef
			if err := node.Content[i+1].Decode(&v); err != nil {
				return fmt.Errorf("failed to decode in YAMLObjectDefs.UnmarshalYAML with YAMLObjectDef type: yamlNode: %+v: %w", node, err)
			}
			if err := node.Content[i].Decode(&v.Name); err != nil {
				return fmt.Errorf("failed to decode in YAMLObjectDefs.UnmarshalYAML with YAMLObjectDef.Name: yamlNode: %+v: %w", node, err)
			}
			list = append(list, v)
		}
		*d = list
		return nil
	default:
		return fmt.Errorf("invalid yaml node %+v", node)
	}
}
