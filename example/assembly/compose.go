package assembly

import (
	"errors"
	"fmt"
	"github.com/gford1000/factory"
)

type Assembler interface {
	GetFactoryContexts() []factory.FactoryContext
	GetInterface(interfaceTypeName string) (interface{}, error)
}

// Wire up the application, applying the policy to the Factories
func New(policy factory.Policy) Assembler {

	a := assembler{
		ctx: map[factory.FactoryContext]bool{},
		m:   map[string]func() interface{}{},
	}

	t, fc, f := composeA(policy)

	a.setType(t.String(), fc, func() interface{} {
		return f()
	})

	return &a
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

// Returns the FactoryContext instances that are registered
func (a *assembler) GetFactoryContexts() []factory.FactoryContext {
	keys := make([]factory.FactoryContext, 0, len(a.ctx))
	for k := range a.ctx {
		keys = append(keys, k)
	}
	return keys
}

// Returns an instance of the wired object
func (a *assembler) GetInterface(interfaceTypeName string) (interface{}, error) {
	if fn, ok := a.m[interfaceTypeName]; !ok {
		return nil, errors.New(fmt.Sprintf("%v is not available", interfaceTypeName))
	} else {
		return fn(), nil
	}
}
