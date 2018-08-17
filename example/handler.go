package main

import (
	"fmt"
	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/assembly"
	"github.com/gford1000/factory/example/interfaces"
)

type handler struct {
	a assembly.Assembler
	o interfaces.A
}

func (h *handler) initialise() {
	policy := &myPolicy{}
	h.a = assembly.New(policy)

	o, err := h.a.GetInterface("interfaces.A")

	if err != nil {
		panic(err)
	}

	h.o = o.(interfaces.A)
}

func (h *handler) applyContext(ctx factory.ExecutionContext) {
	for _, c := range h.a.GetFactoryContexts() {
		c.SetExecutionContext(ctx)
	}
}

func (h *handler) handleRequest() {
	fmt.Println(h.o.Hello())
}

func (h *handler) HandleRequest(ctx factory.ExecutionContext) {
	h.applyContext(ctx)
	h.handleRequest()
}
