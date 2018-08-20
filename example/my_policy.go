package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"

	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/domain_logic"
	"github.com/gford1000/factory/example/interfaces"
)

type myPolicy struct {
}

func (p *myPolicy) Validate() error {
	return nil
}

func (p *myPolicy) GetExemplarHash(ctx factory.ExecutionContext, interfaceType reflect.Type) string {

	// This (bad) implementation is demonstrating that
	// we can set policies for which concrete class should
	// be returned for a given execution context

	if ctx.Hash() == "Context1" {

		if interfaceType == reflect.TypeOf((*interfaces.A)(nil)).Elem() {

			tA1 := reflect.TypeOf(domain_logic.A1{})
			return fmt.Sprintf("%x", sha256.Sum256([]byte(tA1.String())))
		}

		if interfaceType == reflect.TypeOf((*interfaces.B)(nil)).Elem() {

			tB1 := reflect.TypeOf(domain_logic.B1{})
			return fmt.Sprintf("%x", sha256.Sum256([]byte(tB1.String())))
		}
	}

	if ctx.Hash() == "Context2" {

		if interfaceType == reflect.TypeOf((*interfaces.A)(nil)).Elem() {

			tA2 := reflect.TypeOf(domain_logic.A2{})
			return fmt.Sprintf("%x", sha256.Sum256([]byte(tA2.String())))
		}

		if interfaceType == reflect.TypeOf((*interfaces.B)(nil)).Elem() {

			tB2 := reflect.TypeOf(domain_logic.B2{})
			return fmt.Sprintf("%x", sha256.Sum256([]byte(tB2.String())))
		}
	}

	return ""
}
