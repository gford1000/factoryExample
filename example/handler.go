package main

import (
	"fmt"
	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/assembly"
	"github.com/gford1000/factory/example/interfaces"
)

type handler struct {
	assembler assembly.Assembler
	myA       interfaces.A
	myB       interfaces.B
}

func (h *handler) initialise() {
	policy := &myPolicy{}
	h.assembler = assembly.New(policy)

	o, err := h.assembler.GetInterface("interfaces.A")

	if err != nil {
		panic(err)
	}

	h.myA = o.(interfaces.A)

	o, err = h.assembler.GetInterface("interfaces.B")

	if err != nil {
		panic(err)
	}

	h.myB = o.(interfaces.B)
}

func (h *handler) applyContext(ctx factory.ExecutionContext) {
	for _, c := range h.assembler.GetFactoryContexts() {
		c.SetExecutionContext(ctx)
	}
}

func (h *handler) handleRequest() {
	fmt.Println(h.myB.World(h.myA))
}

func (h *handler) HandleRequest(ctx factory.ExecutionContext) {
	h.applyContext(ctx)
	h.handleRequest()
}
