package factory

import (
	"reflect"
)

type ShimBuilder func() interface{}

// Creates a new empty factory, ready for domain object registration
// and a shim creator for the Interface type that can be wired as part of dependency injection
func Init(t reflect.Type, f EmptyShimCreator) (FactoryRegistration, FactoryContext, ShimBuilder) {

	fact, err := NewFactory(t)
	if err != nil {
		panic(err) // For now
	}

	fn := func() interface{} {

		return PrepShim(fact, f)
	}

	return fact, fact, fn
}
