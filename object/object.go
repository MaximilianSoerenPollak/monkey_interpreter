package object

import (
	"fmt"
)


type ObjectType string


const (
	INTEGER_OBJ = "INTEGER" 
	BOOLEAN_OBJ = "BOOLEAN" 
	NULL_OBJ = "NULL"
)

const (
	RETURN_VALUE_OBJ = "RETURN_VALUE" 
)

type ReturnValue struct {
	Value Object 
}

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64 
}

type Boolean struct {
	Value bool 
}
type Null struct {
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (i *Boolean) Inspect() string { return fmt.Sprintf("%t", i.Value) }
func (i *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

func (i *Null) Inspect() string { return "null" }
func (i *Null) Type() ObjectType { return NULL_OBJ }

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
