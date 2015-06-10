package terrain

import (
	"errors"
	"math"
)

type Terrain struct {
	layout     [][]uint16
	max        uint16
	max_height uint16
}

func New(scale uint16, max_height uint16) *Terrain {
	max := uint16(math.Exp2(float64(scale)) + 1)
	t := Terrain{max: max, layout: make([][]uint16, max), max_height: max_height}
	for i := range t.layout {
		t.layout[i] = make([]uint16, max)
	}

	return &t
}

func (t *Terrain) GetHeight(x uint16, y uint16) (num uint16, err error) {
	// Below-zero is unrepresentable by unsigned ints
	if x > t.max-1 || y > t.max-1 {
		return 0, errors.New("Index out of range")
	}

	return t.layout[x][y], nil
}

func (t *Terrain) SetHeight(x uint16, y uint16, height uint16) (err error) {
	if x > t.max-1 || y > t.max-1 {
		return errors.New("Index out of range")
	}

	t.layout[x][y] = height

	return nil
}
