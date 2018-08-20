package assembly

import (
	"github.com/gford1000/factory"
)

// Processor is an interface that can generate a response to a
// user request by invoking the dominan logic
type Processor interface {
	Process() string
}

// ProcessorWrapper applies runtime context prior to calling the Processor
type ProcessorWrapper interface {
	Process(ctx factory.ExecutionContext) string
}

// A basic implementation of Processor
type processorWrapper struct {
	ctx []factory.FactoryContext
	p   Processor
}

func (p *processorWrapper) Process(ctx factory.ExecutionContext) string {
	for _, c := range p.ctx {
		c.SetExecutionContext(ctx)
	}

	return p.p.Process()
}
