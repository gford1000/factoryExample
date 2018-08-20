package domain_logic

import (
	"fmt"

	"github.com/gford1000/factory/example/interfaces"
)

type B1 struct {
}

func (b *B1) World(a interfaces.A) string {
	return fmt.Sprintf("%v World", a.Hello())
}

type B2 struct {
}

func (b *B2) World(a interfaces.A) string {
	return fmt.Sprintf("%v everyone", a.Hello())
}
