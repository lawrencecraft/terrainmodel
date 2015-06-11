package generator

import (
	log "github.com/Sirupsen/logrus"
	terrain "github.com/lawrencecraft/terrainmodel"
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

func incrementAverage(x uint16, y uint16, t *terrain.Terrain, currentCount int, currentSum int) (int, int) {
	v, e := t.GetHeight(x, y)
	if e == nil {
		return currentCount + 1, currentSum + int(v)
	}
	return currentCount, currentSum
}

func normalizeOffset(avg uint16, offset int) uint16 {
	norm := int(avg) + offset
	if norm < 0 {
		return 0
	}
	return uint16(norm)
}

func setSquare(t *terrain.Terrain, x uint16, y uint16, offset int, scale uint16) {
	c, s := incrementAverage(x+scale, y+scale, t, 0, 0)
	c, s = incrementAverage(x-scale, y-scale, t, c, s)
	c, s = incrementAverage(x+scale, y-scale, t, c, s)
	c, s = incrementAverage(x-scale, y+scale, t, c, s)

	avg := uint16(s / c)

	n := normalizeOffset(avg, offset)
	t.SetHeight(x, y, n)
}

func setDiamond(t *terrain.Terrain, x uint16, y uint16, offset int, scale uint16) {
	c, s := incrementAverage(x, y+scale, t, 0, 0)
	c, s = incrementAverage(x, y-scale, t, c, s)
	c, s = incrementAverage(x+scale, y, t, c, s)
	c, s = incrementAverage(x-scale, y, t, c, s)

	avg := uint16(s / c)
	n := normalizeOffset(avg, offset)

	t.SetHeight(x, y, n)
}

func generateTerrain(t *terrain.Terrain, roughness float32, x0y0 uint16, xmaxy0 uint16, x0ymax uint16, xmaxymax uint16) {
	maxDimension := t.Max - 1
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
		for x := halfSize; x < t.Max; x += sideLength {
			for y := halfSize; y < t.Max; y += sideLength {
				offset := int(rand.Float32()*float32(offsetMultiplier)*2.0) - offsetMultiplier
				log.Debug("Setting square: ", x, ",", y, "with offset", offset)
				setSquare(t, x, y, offset, halfSize)
			}
		}

		// Diamond
		for y := uint16(0); y < t.Max; y += halfSize {
			for x := (y + halfSize) % sideLength; x < t.Max; x += sideLength {
				offset := int(rand.Float32()*float32(offsetMultiplier)*2.0) - offsetMultiplier
				log.Debug("Setting diamond: ", x, ",", y, " with offset ", offset)
				setDiamond(t, x, y, offset, halfSize)
			}
		}

		iteration += 1
		offsetMultiplier /= 2
		sideLength = halfSize
	}
}

func (d *DiamondSquareGenerator) Generate() (*terrain.Terrain, error) {
	t := terrain.New(uint16(d.scale), math.MaxUint16)
	generateTerrain(t, d.roughness, uint16(math.MaxUint16*rand.Float32()), uint16(math.MaxUint16*rand.Float32()), uint16(math.MaxUint16*rand.Float32()), uint16(math.MaxUint16*rand.Float32()))
	// Copy to a new terrain of the requested size
	return t, nil
}
