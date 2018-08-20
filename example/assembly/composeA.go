package assembly

import (
	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/domain_logic"
	"github.com/gford1000/factory/example/interfaces"
	"github.com/gford1000/factory/example/shims"
	"github.com/gford1000/factory/exemplars"
	"reflect"
)

// Set up the domain logic implementations that can support interfaces.A
func composeA(policy factory.Policy) (reflect.Type, factory.FactoryContext, factory.ShimBuilder) {
	return composeInterface(
		reflect.TypeOf((*interfaces.A)(nil)).Elem(),
		shims.NewA,
		policy,
		exemplars.NewInstanceExemplar(&domain_logic.A1{}),
		exemplars.NewInstanceExemplar(&domain_logic.A2{}))
}
