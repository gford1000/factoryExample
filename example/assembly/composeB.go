package assembly

import (
	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/domain_logic"
	"github.com/gford1000/factory/example/interfaces"
	"github.com/gford1000/factory/example/shims"
	"github.com/gford1000/factory/exemplars"
	"reflect"
)

// Set up the domain logic implementations that can support interfaces.B
func composeB(policy factory.Policy) (reflect.Type, factory.FactoryContext, factory.ShimBuilder) {

	return composeInterface(
		reflect.TypeOf((*interfaces.B)(nil)).Elem(),
		shims.NewB,
		policy,
		exemplars.NewInstanceExemplar(&domain_logic.B1{}),
		exemplars.NewInstanceExemplar(&domain_logic.B2{}))
}
