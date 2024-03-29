package api

import (
	"fmt"
	"strings"

	"github.com/cuishu/zero-api/ast"
)

type Field struct {
	Name      string
	Type      string
	Documents string
	Doc       string
	Validate  string
	isJSON    bool
}

type Type struct {
	Name      string
	Fields    []Field
	Documents string
	Doc       string
}

func (t Type) IsJSON() bool {
	for _, field := range t.Fields {
		return field.isJSON
	}
	return false
}

func goTypeToTsType(t string) string {
	if strings.Index(t, "[]") == 0 {
		t = fmt.Sprintf("Array<%s>", goTypeToTsType(strings.TrimLeft(t, "[]")))
	}
	switch t {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64":
		return "number"
	case "bool":
		return "boolean"
	case "file":
		return "any"
	default:
		return t
	}
}

func memberToField(member ast.Field) Field {
	tags := strings.Split(strings.Trim(member.Tag, "`"), " ")
	if len(tags) == 0 {
		return Field{}
	}
	var field Field
	field.Name = member.Name
	field.Documents = toDocument(member.Comment, 4)
	field.Doc = strings.Trim(strings.Trim(member.Comment, "/"), "*")
	field.Type = goTypeToTsType(member.Type)
	for _, tag := range tags {
		slice := strings.Split(tag, ":")
		switch slice[0] {
		case "json", "form":
			field.isJSON = slice[0] == "json"
			field.Name = strings.Trim(slice[1], `"`)
		case "binding":
			field.Validate = strings.Trim(slice[1], `"`)
		}
	}
	return field
}

func convertSpecType(item ast.Type) Type {
	var t Type
	t.Name = item.Name
	t.Documents = toDocument(item.Comment, 1)
	t.Doc = strings.TrimSpace(strings.Trim(strings.Trim(item.Comment, "/"), "*"))
	for _, member := range item.Fields {
		t.Fields = append(t.Fields, memberToField(member))
	}
	return t
}
