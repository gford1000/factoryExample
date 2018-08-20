package main

import (
	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/assembly"
)

// Handler has two responsibilities:
//
// 1. On application start, create the Processor needed to handle requests successfully.
//    This requires the runtime security policy to be loaded and the application to be assembled
//
// 2. Forward requests to the Processor once initialised
type handler struct {
	p assembly.ProcessorWrapper
}

func (h *handler) initialise() {
	policy := &myPolicy{} // This would be loaded from config system for production systems
	h.p = assembly.New(policy, assembly.WireSimpleProcessor)
}

func (h *handler) handleRequest(ctx factory.ExecutionContext) string {
	return h.p.Process(ctx)
}
