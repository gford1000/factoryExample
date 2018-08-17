package exemplars

import (
	"crypto/sha256"
	"fmt"
	"github.com/gford1000/factory"
	"reflect"
)

// Creates an Exemplar that can create new instances of the
// supplied srcPtr Type.  srcPtr must be a pointer to the Type.
func NewInstanceExemplar(srcPtr interface{}) factory.Exemplar {

	typ := reflect.TypeOf(srcPtr)

	if typ.Kind() != reflect.Ptr {
		panic("Should be a pointer")
	}

	typ = typ.Elem()

	return &nie{
		typ: typ,
		h:   fmt.Sprintf("%x", sha256.Sum256([]byte(typ.String()))),
	}
}

// Hidden class, implements Exemplar
type nie struct {
	typ reflect.Type
	h   string
}

func (n *nie) Hash() string {
	return n.h
}

func (n *nie) Implements(t reflect.Type) bool {
	return reflect.TypeOf(reflect.New(n.typ).Interface()).Implements(t)
}

func (n *nie) GetInstance() interface{} {
	return reflect.New(n.typ).Interface()
}
