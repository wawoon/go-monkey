package object

import (
    "bytes"
    "fmt"
    "strings"
    "github.com/wawoon/go-monkey/ast"
)

// ObjectType is a string alias to identify object types at runtime.
type ObjectType string

const (
    INTEGER_OBJ      = "INTEGER"
    BOOLEAN_OBJ      = "BOOLEAN"
    NULL_OBJ         = "NULL"
    RETURN_VALUE_OBJ = "RETURN_VALUE"
    FUNCTION_OBJ     = "FUNCTION"
)

// Object represents a value at runtime.
type Object interface {
    Type() ObjectType
    Inspect() string
}

// Integer value
type Integer struct {
    Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// Boolean value
type Boolean struct {
    Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Null represents absence of value
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// ReturnValue is used to unwind the call stack when encountering a return statement.
type ReturnValue struct {
    Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// Environment holds variable bindings.
type Environment struct {
    store map[string]Object
    outer *Environment
}

func NewEnvironment() *Environment { return &Environment{store: make(map[string]Object)} }

func NewEnclosedEnvironment(outer *Environment) *Environment {
    env := NewEnvironment()
    env.outer = outer
    return env
}

func (e *Environment) Get(name string) (Object, bool) {
    obj, ok := e.store[name]
    if !ok && e.outer != nil {
        obj, ok = e.outer.Get(name)
    }
    return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
    e.store[name] = val
    return val
}

// Function object
type Function struct {
    Parameters []*ast.Identifier
    Body       *ast.BlockStatement
    Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
    var out bytes.Buffer
    params := []string{}
    for _, p := range f.Parameters {
        params = append(params, p.String())
    }
    out.WriteString("fn(")
    out.WriteString(strings.Join(params, ", "))
    out.WriteString(") {")
    out.WriteString(f.Body.String())
    out.WriteString("}")
    return out.String()
}

