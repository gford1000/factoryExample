package assembly

import (
	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/interfaces"
)

// Processor is an interface that can generate a response to a
// user request by invoking the dominan logic
type Processor interface {
	Process(ctx factory.ExecutionContext) string
}

// A basic implementation of Processor
type processor struct {
	assembler *assembler
	a         interfaces.A
	b         interfaces.B
}

func (p *processor) Process(ctx factory.ExecutionContext) string {
	for c := range p.assembler.ctx {
		c.SetExecutionContext(ctx)
	}
	return p.b.World(p.a)
}
