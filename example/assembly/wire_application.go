package assembly

import (
	"github.com/gford1000/factory/example/interfaces"
)

// Wires domain logic instances so that requests can be processed
func WireSimpleProcessor(a Assembler) Processor {
	oA, err := a.GetInterface("interfaces.A")
	if err != nil {
		panic("Unable to create object of type interfaces.A")
	}

	oB, err := a.GetInterface("interfaces.B")
	if err != nil {
		panic("Unable to create object of type interfaces.B")
	}

	return &processor{
		a: oA.(interfaces.A),
		b: oB.(interfaces.B),
	}
}

// A basic implementation of Processor
type processor struct {
	a interfaces.A
	b interfaces.B
}

func (p *processor) processRequest() string {
	return p.b.World(p.a)
}

func (p *processor) Process() string {
	return p.b.World(p.a)
}
