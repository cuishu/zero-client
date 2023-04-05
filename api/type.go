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
	isJSON    bool
}

type Type struct {
	Name      string
	Fields    []Field
	Documents string
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
	slice := strings.Split(tags[0], ":")
	return Field{
		isJSON:    slice[0] == "json",
		Name:      strings.Trim(slice[1], `"`),
		Type:      goTypeToTsType(member.Type),
		Documents: member.Comment,
	}
}

func convertSpecType(item ast.Type) Type {
	var t Type
	t.Name = item.Name
	t.Documents = item.Comment
	for _, member := range item.Fields {
		t.Fields = append(t.Fields, memberToField(member))
	}
	return t
}
