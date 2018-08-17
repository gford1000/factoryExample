package factories

import (
	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/interfaces"
)

func NewA() factory.Shim {
	return &shimA{}
}

// Hidden type that checks for and handles execution context changes
// for Interface instances of type A
type shimA struct {
	chkr     factory.Checker // Used to respond to execution context changes
	instance interfaces.A    // Actual object vended by Factory
}

func (s *shimA) SetInstance(o interface{}) {
	s.instance = o.(interfaces.A)
}

func (s *shimA) SetContextChecker(c factory.Checker) {
	s.chkr = c
}

// This is shimA's implementation of each method of interfaces.A
//
// All methods do the same, simple thing:
//
// 1. Invoke Check() to reset the Shim if necessary
// 2. Delegate to the underlying Interface instance to do the work

// Hello is a defined method on interfaces.A
func (s *shimA) Hello() string {
	s.chkr.Check()
	return s.instance.Hello()
}
