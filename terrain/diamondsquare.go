package terrain

import (
	log "github.com/Sirupsen/logrus"
	"math"
	"math/rand"
	"time"
)

type DiamondSquareGenerator struct {
	roughness float32
	scale     int
	x         int
	y         int
}

func NewDiamondSquareGenerator(roughness float32, x int, y int) *DiamondSquareGenerator {
	scale := int(math.Ceil(math.Log2(float64(x - 1))))
	return &DiamondSquareGenerator{x: x, y: y, scale: scale, roughness: roughness}
}

func IncrementAverage(x uint16, y uint16, t *Terrain, currentCount int, currentSum int) (int, int) {
	v, e := t.GetHeight(x, y)
	if e == nil {
		return currentCount + 1, currentSum + int(v)
	}
	return currentCount, currentSum
}

func NormalizeOffset(avg uint16, offset int) uint16 {
	norm := int(avg) + offset
	if norm < 0 {
		return 0
	}
	return uint16(norm)
}

func SetSquare(t *Terrain, x uint16, y uint16, offset int, scale uint16) {
	c, s := IncrementAverage(x+scale, y+scale, t, 0, 0)
	c, s = IncrementAverage(x-scale, y-scale, t, c, s)
	c, s = IncrementAverage(x+scale, y-scale, t, c, s)
	c, s = IncrementAverage(x-scale, y+scale, t, c, s)

	avg := uint16(s / c)

	n := NormalizeOffset(avg, offset)
	t.SetHeight(x, y, n)
}

func SetDiamond(t *Terrain, x uint16, y uint16, offset int, scale uint16) {
	c, s := IncrementAverage(x, y+scale, t, 0, 0)
	c, s = IncrementAverage(x, y-scale, t, c, s)
	c, s = IncrementAverage(x+scale, y, t, c, s)
	c, s = IncrementAverage(x-scale, y, t, c, s)

	avg := uint16(s / c)
	n := NormalizeOffset(avg, offset)

	t.SetHeight(x, y, n)
}

func Generate(t *Terrain, roughness float32, x0y0 uint16, xmaxy0 uint16, x0ymax uint16, xmaxymax uint16) {
	maxDimension := t.max - 1
	offsetMultiplier := int(math.MaxUint16)
	t.SetHeight(0, 0, x0y0)
	t.SetHeight(maxDimension, 0, xmaxy0)
	t.SetHeight(0, maxDimension, x0ymax)
	t.SetHeight(maxDimension, maxDimension, xmaxymax)

	rand.Seed(time.Now().UnixNano())

	iteration := 1

	sideLength := maxDimension
	for sideLength > 1 {

		halfSize := uint16(sideLength / 2)
		log.Debug("Size ", sideLength, " halfsize:", halfSize)

		// Square
		for x := halfSize; x < t.max; x += sideLength {
			for y := halfSize; y < t.max; y += sideLength {
				offset := int(rand.Float32()*float32(offsetMultiplier)*2.0) - offsetMultiplier
				log.Debug("Setting square: ", x, ",", y, "with offset", offset)
				SetSquare(t, x, y, offset, halfSize)
			}
		}

		// Diamond
		for y := uint16(0); y < t.max; y += halfSize {
			for x := (y + halfSize) % sideLength; x < t.max; x += sideLength {
				offset := int(rand.Float32()*float32(offsetMultiplier)*2.0) - offsetMultiplier
				log.Debug("Setting diamond: ", x, ",", y, " with offset ", offset)
				SetDiamond(t, x, y, offset, halfSize)
			}
		}

		iteration += 1
		offsetMultiplier /= 2
		sideLength = halfSize
	}
}
