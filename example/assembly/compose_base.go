package assembly

import (
	"fmt"
	"reflect"

	"github.com/gford1000/factory"
)

// Set up the domain logic implementations that can support the specified Interface type
func composeInterface(t reflect.Type, newShim factory.EmptyShimCreator, policy factory.Policy, exemplars ...factory.Exemplar) (reflect.Type, factory.FactoryContext, factory.ShimBuilder) {

	factoryRegistration, factoryContext, shimBuilder := factory.Init(t, newShim)

	factoryRegistration.SetPolicy(policy)

	for _, e := range exemplars {
		err := factoryRegistration.AddExemplar(e)
		if err != nil {
			errType := reflect.TypeOf(e.GetInstance()).Elem()
			panic(fmt.Sprintf("Cannot register %v for Interface %v.  Check the type implements the interface", errType, t))
		}
	}

	return t, factoryContext, shimBuilder
}
