package factory

// Represents an intermediary between the caller
// and the code to be called
type Shim interface {
	SetInstance(o interface{})
	SetContextChecker(c Checker)
}

type EmptyShimCreator func() Shim

// Shims need to be able to respond to changes in execution context,
// requesting the appropriate instance of the Interface from its Factory
// This function sets up the wiring necessary to achieve this behaviour
func PrepShim(fact *Factory, emptyShimCreator EmptyShimCreator) Shim {

	s := emptyShimCreator()

	fn := func() {
		o, err := fact.GetInterface()
		if err != nil {
			panic(err)
		}
		if o == nil {
			panic("No Interface available from the Factory")
		}
		s.SetInstance(o)
	}

	s.SetContextChecker(&contextChecker{
		ctx:     fact.GetExecutionContextHash(),
		onChg:   fn,
		factory: fact,
	})

	return s
}
