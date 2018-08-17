package factory

import (
	"errors"
	"fmt"
	"reflect"
)

// Constructs a Factory to manage the specified Type
func NewFactory(interfaceType reflect.Type) (*Factory, error) {
	if interfaceType == nil {
		return nil, errors.New("Cannot provide nil as interfaceType")
	}
	if interfaceType.Kind() != reflect.Interface {
		return nil, errors.New("interfaceType must be of Kind reflect.Interface.  Check you are passing reflect.TypeOf((*A)(nil)).Elem() where A is the interface to vend instances for")
	}

	return &Factory{
		typ:          interfaceType,
		interfaceMap: map[string]Exemplar{},
	}, nil
}

// Base of all Factory types
type Factory struct {
	executionContext ExecutionContext    // The current execution context used to determine Exemplars
	policy           Policy              // The Policy which determines Exemplars from context
	typ              reflect.Type        // The Interface Type that the Factory is vending
	interfaceMap     map[string]Exemplar // The available Exemplars that could be used to vend instances
}

// Assigns an execution context to the Factory
func (f *Factory) SetExecutionContext(ctx ExecutionContext) error {
	f.executionContext = ctx
	return nil
}

// Retrieve the hash of the current execution context
// Empty string is returned if no execution context is defined
func (f *Factory) GetExecutionContextHash() string {
	if f.executionContext != nil {
		return f.executionContext.Hash()
	} else {
		return ""
	}
}

// Assign the specified security Policy to the Factory
func (f *Factory) SetPolicy(policy Policy) error {
	if policy == nil {
		return errors.New("Cannot assign a nil Policy")
	}

	if err := policy.Validate(); err != nil {
		return errors.New(fmt.Sprintf("Invalid Policy: %v", err))
	}

	f.policy = policy
	return nil
}

// Assigns an Exemplar.  Repeated additions of the same Exemplar against a type should be idempotent
func (f *Factory) AddExemplar(exemplar Exemplar) error {
	if exemplar == nil {
		return errors.New("Cannot provide nil as exemplar")
	}
	if !exemplar.Implements(f.typ) {
		return errors.New(fmt.Sprintf("Exemplar %v does not implement %v", exemplar.Hash(), f.typ))
	}

	// Add to map if not already present
	h := exemplar.Hash()
	if _, ok := f.interfaceMap[h]; !ok {
		f.interfaceMap[h] = exemplar
	}

	return nil
}

// Returns an instance constructed from an Exemplar, determined by applying the
// current execution context to the security Policy
func (f *Factory) GetInterface() (interface{}, error) {
	// Must have a Policy assigned
	if f.policy == nil {
		return nil, errors.New("No Policy assigned to Factory - cannot vend interfaces")
	}

	// Not having an execution context is not an error, but
	// no object is vended
	if f.executionContext == nil {
		return nil, nil
	}

	// Apply policy to find the appropriate Exemplar to use
	exemplarHash := f.policy.GetExemplarHash(f.executionContext, f.typ)

	// Policy doesn't provide an Exemplar is an error - we don't know how
	// to manage security for this interface
	if exemplarHash == "" {
		return nil, errors.New(fmt.Sprintf("No Policy defined for interface %v in the current execution context", f.typ))
	}

	var obj interface{}
	var err error

	if exemplar, ok := f.interfaceMap[exemplarHash]; !ok {
		// Not having the specified Exemplar is a registration error
		err = errors.New(fmt.Sprintf("Required Exemplar %v is not registered for %v", exemplarHash, f.typ))
	} else {
		// Build instance as concrete implementation of the requested interface
		obj = exemplar.GetInstance()
	}

	return obj, err
}
