package assembly

import (
	"reflect"

	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/interfaces"
)

// All Instance types must have a composer function emitting these items:
// - The type of the Interface
// - The FactoryContext interface to the underlying Factory, allowing execution context to be updated
// - A function that returns a populated Shim to an instance of the Interface
type Composer func(policy factory.Policy) (reflect.Type, factory.FactoryContext, factory.ShimBuilder)

// Wire together the business logic, which is
// exposed as a Processor instance
func New(policy factory.Policy) Processor {

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

	oA, err := a.getInterface("interfaces.A")
	if err != nil {
		panic("Unable to create object of type interfaces.A")
	}

	oB, err := a.getInterface("interfaces.B")
	if err != nil {
		panic("Unable to create object of type interfaces.B")
	}

	return &processor{
		assembler: a,
		a:         oA.(interfaces.A),
		b:         oB.(interfaces.B),
	}
}
