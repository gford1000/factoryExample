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
	ctx []factory.FactoryContext
	a   interfaces.A
	b   interfaces.B
}

func (p *processor) setContext(ctx factory.ExecutionContext) {
	for _, c := range p.ctx {
		c.SetExecutionContext(ctx)
	}
}

func (p *processor) processRequest() string {
	return p.b.World(p.a)
}

func (p *processor) Process(ctx factory.ExecutionContext) string {
	p.setContext(ctx)
	return p.processRequest()
}
