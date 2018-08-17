package main

func main() {

	// Wire up application
	h := handler{}
	h.initialise()

	h.HandleRequest(&myContext{h: "Context1"})
	h.HandleRequest(&myContext{h: "Context2"})
	//h.HandleRequest(&myContext{h: "Unknown"})
}
