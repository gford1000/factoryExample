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

	a := populateAssembler(policy)

	return wireProcessor(a)
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
func wireProcessor(a *assembler) Processor {
	oA, err := a.getInterface("interfaces.A")
	if err != nil {
		panic("Unable to create object of type interfaces.A")
	}

	oB, err := a.getInterface("interfaces.B")
	if err != nil {
		panic("Unable to create object of type interfaces.B")
	}

	ctx := []factory.FactoryContext{}
	for k := range a.ctx {
		ctx = append(ctx, k)
	}

	return &processor{
		ctx: ctx,
		a:   oA.(interfaces.A),
		b:   oB.(interfaces.B),
	}
}
