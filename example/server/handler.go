package main

import (
	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/assembly"
)

type handler struct {
	p assembly.Processor
}

func (h *handler) initialise() {
	policy := &myPolicy{}
	h.p = assembly.New(policy)
}

func (h *handler) handleRequest(ctx factory.ExecutionContext) string {
	return h.p.Process(ctx)
}
