package assembly

import (
	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/domain_logic"
	"github.com/gford1000/factory/example/factories"
	"github.com/gford1000/factory/example/interfaces"
	"github.com/gford1000/factory/exemplars"
	"reflect"
)

// Set up the domain logic implementations that can support interfaces.A
func composeA(policy factory.Policy) (reflect.Type, factory.FactoryContext, factory.ShimBuilder) {

	typeA := reflect.TypeOf((*interfaces.A)(nil)).Elem()

	factoryRegistrationA, factoryContext, shimBuilder := factory.Init(typeA, factories.NewA)

	factoryRegistrationA.SetPolicy(policy)

	factoryRegistrationA.AddExemplar(exemplars.NewInstanceExemplar(&domain_logic.A1{}))
	factoryRegistrationA.AddExemplar(exemplars.NewInstanceExemplar(&domain_logic.A2{}))

	return typeA, factoryContext, shimBuilder
}
