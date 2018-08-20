package assembly

import (
	"errors"
	"fmt"

	"github.com/gford1000/factory"
)

type Assembler interface {
	GetInterface(interfaceTypeName string) (interface{}, error)
}

// Hidden implementation of Assembler
type assembler struct {
	ctx map[factory.FactoryContext]bool
	m   map[string]func() interface{}
}

// Saves the composer against its type
func (a *assembler) setType(interfaceTypeName string, ctx factory.FactoryContext, creator func() interface{}) {
	a.m[interfaceTypeName] = creator
	a.ctx[ctx] = true
}

// Returns an instance of the wired object
func (a *assembler) GetInterface(interfaceTypeName string) (interface{}, error) {
	if fn, ok := a.m[interfaceTypeName]; !ok {
		return nil, errors.New(fmt.Sprintf("%v is not available", interfaceTypeName))
	} else {
		return fn(), nil
	}
}
