package factory

import (
	"reflect"
)

// Generalised hash retrieval interface
type Hasher interface {
	Hash() string
}

// Generalised checking interface
type Checker interface {
	Check()
}

// ExecutionContext provides information that the Policy
// can use to determine the correct Exemplar to be vended
type ExecutionContext interface {
	Hasher
}

// Exemplar provides the means for objects registered in a Factory
// to emit a suitable instance of themselves.  The object has
// control to emit itself, an empty copy or a cloned (initialised) copy
type Exemplar interface {
	Hasher
	// True if Exemplar implements the type
	Implements(t reflect.Type) bool
	// Return an instance representing the Exemplar
	GetInstance() interface{}
}

// FactoryContext allows the current execution context to be supplied to
// a Factory, which can be applied to a Policy when retrieving an interface
type FactoryContext interface {
	// Assign the context to the Factory
	SetExecutionContext(ctx ExecutionContext) error
}

// Allows Factories to indicate a change in execution context
type FactoryContextChecker interface {
	// Returns the current execution context hash, allowing
	// changes to be detected by the caller
	GetExecutionContextHash() string
}

// Called during application initialisation, this allows the security Policy
// and the set of interfaces supported by the Factory to be defined
type FactoryRegistration interface {
	// Assigns a Policy to the Factory.  An error should be thrown
	// if a Policy is already assigned
	SetPolicy(policy Policy) error
	// Assigns an Exemplar for the interface type managed by the Factory.
	// Should return an error if the Exemplar does not implement the interface
	// Repeated additions of the same Exemplar against a type should be idempotent
	AddExemplar(exemplar Exemplar) error
}

// Policy allows Factory implementations to contextually determine which
// Exemplar should be used to create a concrete implementation
// of any given interface
type Policy interface {
	// Returns error if the policy isn't valid
	Validate() error
	// Returns the hash of the Exemplar which can generate an object
	// supporting the interfaceType, given the execution context.
	// Should return "" if no Exemplars are available for this context.
	GetExemplarHash(ctx ExecutionContext, interfaceType reflect.Type) string
}
