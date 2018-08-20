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

// Declare the Instances to be used by the application
func New(policy factory.Policy) Assembler {

	a := assembler{
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

	return &a
}
