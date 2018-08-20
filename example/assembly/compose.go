package assembly

import (
	"reflect"

	"github.com/gford1000/factory"
)

// All Instance types must have a composer function emitting these items:
// - The type of the Interface
// - The FactoryContext interface to the underlying Factory, allowing execution context to be updated
// - A function that returns a populated Shim to an instance of the Interface
type Composer func(policy factory.Policy) (reflect.Type, factory.FactoryContext, factory.ShimBuilder)

// This is the function that wires together instances of Interfaces to create a Processor
// that can deal with runtime requests.
type WireProcessor func(a Assembler) Processor

// Wire together the business logic, which is
// exposed as a Processor instance
func New(policy factory.Policy, fn WireProcessor) ProcessorWrapper {

	a := populateAssembler(policy)

	return wireProcessorWrapper(a, fn)
}

// Creates factories and shims for each domain logic interface
func populateAssembler(policy factory.Policy) *assembler {

	a := &assembler{
		ctx: map[factory.FactoryContext]bool{},
		m:   map[string]func() interface{}{},
	}

	fns := []Composer{
		composeA,
		composeB,
	}

	for _, composerFn := range fns {
		t, fc, f := composerFn(policy)

		a.setType(t.String(), fc, func() interface{} {
			return f()
		})
	}

	return a
}

// Wires domain logic instances so that requests can be processed
func wireProcessorWrapper(a *assembler, fn WireProcessor) ProcessorWrapper {
	ctx := []factory.FactoryContext{}
	for k := range a.ctx {
		ctx = append(ctx, k)
	}

	return &processorWrapper{
		ctx: ctx,
		p:   fn(a),
	}
}
