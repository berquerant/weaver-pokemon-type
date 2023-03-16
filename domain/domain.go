package domain

import "github.com/ServiceWeaver/weaver"

//go:generate weaver generate ./...

type Type struct {
	weaver.AutoMarshal
	ID   int
	Name string
}

type Effectivity struct {
	weaver.AutoMarshal
	ID         int
	Attack     Type
	Defense    Type
	Multiplier float32
}
