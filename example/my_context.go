package main

type myContext struct {
	h string
}

func (c *myContext) Hash() string {
	return c.h
}
