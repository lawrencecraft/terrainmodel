package terrain

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type Terrain struct {
	layout    []uint16
	X         uint16
	Y         uint16
	MaxHeight uint16
}

func New(x uint16, y uint16, max_height uint16) *Terrain {
	t := Terrain{X: x, Y: y, layout: make([]uint16, int64(x)*int64(y)), MaxHeight: max_height}
	return &t
}

func (t *Terrain) getIndex(x uint16, y uint16) uint16 {
	return x*(t.X-1) + y
}

func (t *Terrain) GetHeight(x uint16, y uint16) (num uint16, err error) {
	// Below-zero is unrepresentable by unsigned ints
	if x > t.X-1 || y > t.Y-1 {
		return 0, errors.New("Index out of range")
	}
	index := t.getIndex(x, y)

	return t.layout[index], nil
}

func (t *Terrain) SetHeight(x uint16, y uint16, height uint16) (err error) {
	if x > t.X-1 || y > t.Y-1 {
		return errors.New(fmt.Sprintf("%d,%d is out of range (max %d,%d)", x, y, t.X, t.Y))
	}
	index := t.getIndex(x, y)

	t.layout[index] = height

	return nil
}

func (t *Terrain) CopyTo(destination *Terrain, xOffset uint16, yOffset uint16) error {
	maxX := destination.X + xOffset
	maxY := destination.Y + yOffset
	if maxX > t.X {
		return errors.New(fmt.Sprintf("X-length %d is out of bounds. Must be less than %d", maxX, t.X))
	}

	if maxY > t.Y {
		return errors.New(fmt.Sprintf("Y-length %d is out of bounds. Must be less than %d", maxY, t.Y))
	}

	for x := uint16(0); x < destination.X; x++ {
		for y := uint16(0); y < destination.Y; y++ {
			sourceX := x + xOffset
			sourceY := y + yOffset
			sourceIndex := t.getIndex(sourceX, sourceY)
			destIndex := destination.getIndex(x, y)

			destination.layout[destIndex] = t.layout[sourceIndex]
		}
	}

	return nil
}

func (t *Terrain) Iterate(fn func(uint16, uint16, uint16)) {
	for x := uint16(0); x < t.X; x++ {
		for y := uint16(0); y < t.Y; y++ {
			fn(x, y, t.layout[t.getIndex(x, y)])
		}
	}
}

func (t *Terrain) Flush(writer io.Writer) error {
	err := binary.Write(writer, binary.LittleEndian, t.layout)
	return err
}
