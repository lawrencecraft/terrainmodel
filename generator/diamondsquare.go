package generator

import (
	log "github.com/Sirupsen/logrus"
	terrain "github.com/lawrencecraft/terrainmodel"
	"math"
	"math/rand"
	"time"
)

type DiamondSquareGenerator struct {
	Roughness float32
	X         int
	Y         int
}

func NewDiamondSquareGenerator(roughness float32, x int, y int) *DiamondSquareGenerator {
	return &DiamondSquareGenerator{X: x, Y: y, Roughness: roughness}
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
	maxDimension := t.X - 1 // At this point, we are 100% sure the terrain will be square
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
		for x := halfSize; x < t.X; x += sideLength {
			for y := halfSize; y < t.Y; y += sideLength {
				offset := int(float32(int(rand.Float32()*float32(offsetMultiplier)*2.0)-offsetMultiplier) * roughness)
				log.Debug("Setting square: ", x, ",", y, "with offset", offset)
				setSquare(t, x, y, offset, halfSize)
			}
		}

		// Diamond
		for y := uint16(0); y < t.Y; y += halfSize {
			for x := (y + halfSize) % sideLength; x < t.X; x += sideLength {
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

func getScale(x int, y int) int {
	maxDimension := x
	if y > x {
		maxDimension = y
	}
	return int(math.Ceil(math.Log2(float64(maxDimension - 1))))
}

func (d *DiamondSquareGenerator) Generate() (*terrain.Terrain, error) {
	scale := getScale(d.X, d.Y)
	squareDimension := uint16(math.Exp2(float64(scale))) + 1 // Must be 2^n + 1 -- use the minimum size we calculated previously
	log.Debug("Generating square of dimension ", squareDimension)
	operatingTerrain := terrain.New(squareDimension, squareDimension, math.MaxUint16)
	generateTerrain(operatingTerrain, d.Roughness, uint16(math.MaxUint16*rand.Float32()), uint16(math.MaxUint16*rand.Float32()), uint16(math.MaxUint16*rand.Float32()), uint16(math.MaxUint16*rand.Float32()))
	// Special case: if the generated terrain is the same as what is requested, return it
	if d.X == int(squareDimension) && d.Y == int(squareDimension) {
		log.Debug("Shortcutting copy - no need as requested terrain is square")
		return operatingTerrain, nil
	}
	// Copy to a new terrain of the requested size
	returningTerrain := terrain.New(uint16(d.X), uint16(d.Y), operatingTerrain.MaxHeight)
	err := operatingTerrain.CopyTo(returningTerrain, 0, 0)
	if err != nil {
		log.Debug("diamondsquare.go - encountered error on terrain copy: ", err)
		return nil, err
	}
	return returningTerrain, nil
}
