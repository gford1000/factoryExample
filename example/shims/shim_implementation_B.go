package shims

import (
	"github.com/gford1000/factory"
	"github.com/gford1000/factory/example/interfaces"
)

func NewB() factory.Shim {
	return &shimB{}
}

// Hidden type that checks for and handles execution context changes
// for Interface instances of type B
type shimB struct {
	chkr     factory.Checker // Used to respond to execution context changes
	instance interfaces.B    // Actual object vended by Factory
}

func (s *shimB) SetInstance(o interface{}) {
	s.instance = o.(interfaces.B)
}

func (s *shimB) SetContextChecker(c factory.Checker) {
	s.chkr = c
}

// This is shimA's implementation of each method of interfaces.A
//
// All methods do the same, simple thing:
//
// 1. Invoke Check() to reset the Shim if necessary
// 2. Delegate to the underlying Interface instance to do the work

// Hello is a defined method on interfaces.B
func (s *shimB) World(a interfaces.A) string {
	s.chkr.Check()
	return s.instance.World(a)
}
