package terrain

import (
	"errors"
	"fmt"
	"math"
)

type Terrain struct {
	layout    []uint16
	Max       uint16
	MaxHeight uint16
}

func New(scale uint16, max_height uint16) *Terrain {
	max := uint16(math.Exp2(float64(scale)) + 1)
	t := Terrain{Max: max, layout: make([]uint16, int64(max)*int64(max)), MaxHeight: max_height}
	return &t
}

func (t *Terrain) GetHeight(x uint16, y uint16) (num uint16, err error) {
	// Below-zero is unrepresentable by unsigned ints
	if x > t.Max-1 || y > t.Max-1 {
		return 0, errors.New("Index out of range")
	}
	index := x*(t.Max-1) + y

	return t.layout[index], nil
}

func (t *Terrain) SetHeight(x uint16, y uint16, height uint16) (err error) {
	if x > t.Max-1 || y > t.Max-1 {
		return errors.New(fmt.Sprintf("%d,%d is out of range (max %d)", x, y, t.Max))
	}
	index := x*(t.Max-1) + y

	t.layout[index] = height

	return nil
}
