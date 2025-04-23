package evaluator

import (
	"github.com/Martin-Martinez4/compiler_in_go/object"
)

var builtins = map[string]*object.Builtin{
	"len":   object.GetBuiltinByName("len"),
	"first": object.GetBuiltinByName("first"),
	"last":  object.GetBuiltinByName("last"),
	"rest":  object.GetBuiltinByName("rest"),
	"push":  object.GetBuiltinByName("push"),
	"put":   object.GetBuiltinByName("put"),
}
