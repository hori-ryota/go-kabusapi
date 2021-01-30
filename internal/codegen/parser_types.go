package main

import (
	"strings"

	strcase "github.com/hori-ryota/go-strcase"
)

type TypeDef interface {
	ToGoType() string
}

type StructDef struct {
	Name        string
	Description string
	Properties  []PropertyDef
}

func (s StructDef) ToGoType() string {
	return s.Name
}

type PropertyDef struct {
	Name        string
	Type        TypeDef
	Required    bool
	Description string
}

type ArrayDef struct {
	Elem TypeDef
}

func (s ArrayDef) ToGoType() string {
	return "[]" + s.Elem.ToGoType()
}

type EnumDef struct {
	Prefix   string
	Name     string
	BaseType TypeDef
	Values   []EnumValue
}

type EnumValue struct {
	Name        string
	Value       string
	Description string
}

func (s EnumDef) ToGoType() string {
	return s.Prefix + strcase.ToUpperCamel(s.Name)
}

type PrimitiveTypeDef string

func (s PrimitiveTypeDef) ToGoType() string {
	return string(s)
}

const (
	BoolDef    PrimitiveTypeDef = "bool"
	StringDef  PrimitiveTypeDef = "string"
	Int32Def   PrimitiveTypeDef = "int32"
	Float64Def PrimitiveTypeDef = "float64"
)

type RefDef string

func (s RefDef) ToGoType() string {
	return strings.TrimPrefix(string(s), "#/components/schemas/")
}

type MethodDef struct {
	Name        string
	HTTPMethod  string
	HTTPPath    string
	PathParams  []PathParamDef
	InputType   TypeDef
	OutputType  TypeDef
	Summary     string
	Description string
}

type PathParamDef struct {
	Name        string
	Type        TypeDef
	Description string
}
