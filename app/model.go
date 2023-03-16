package app

//go:generate weaver generate ./...
import (
	"errors"

	"github.com/ServiceWeaver/weaver"
)

type Type struct {
	weaver.AutoMarshal
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Effectivity struct {
	weaver.AutoMarshal
	ID         int     `json:"id"`
	Attack     Type    `json:"attack"`
	Defense    Type    `json:"defense"`
	Multiplier float32 `json:"multiplier"`
}

type PiledEffectivity []Effectivity

func (p PiledEffectivity) Multiplier() float32 {
	if len(p) == 0 {
		return -1
	}
	var v float32 = 1
	for _, e := range p {
		v *= e.Multiplier
	}
	return v
}

var (
	ErrNilPiledEffectivity = errors.New("NilPiledEffectivity")
)

func (p PiledEffectivity) Validate() error {
	if len(p) == 0 {
		return ErrNilPiledEffectivity
	}
	return nil
}

func (p PiledEffectivity) Append(e Effectivity) PiledEffectivity {
	return append(p, e)
}

type PileIndex int

var (
	ErrPileIndexTooSmall = errors.New("PileIndexTooSmall")
	ErrPileIndexTooBig   = errors.New("ErrPileIndexTooBig")
)

func (p PileIndex) Validate() error {
	switch {
	case p < 1:
		return ErrPileIndexTooSmall
	case p > 4:
		return ErrPileIndexTooBig
	default:
		return nil
	}
}

type DefenseTypeIDList []int

var (
	ErrDefenseTypeIDListTooSmall = errors.New("DefenseTypeIDListTooSmall")
	ErrDefenseTypeIDListTooBig   = errors.New("DefenseTypeIDListTooBig")
)

func (d DefenseTypeIDList) Validate() error {
	switch {
	case len(d) < 1:
		return ErrDefenseTypeIDListTooSmall
	case len(d) > 4:
		return ErrDefenseTypeIDListTooBig
	default:
		return nil
	}
}
