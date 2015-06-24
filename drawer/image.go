package drawer

import (
	"github.com/lawrencecraft/terrainmodel"
	"image"
	"image/color"
)

func toGrayscale16Image(t *terrain.Terrain) *image.Gray16 {
	i := image.NewGray16(image.Rect(0, 0, int(t.X), int(t.Y)))
	t.Iterate(func(x uint16, y uint16, val uint16) {
		i.SetGray16(int(x), int(y), color.Gray16{Y: val})
	})
	return i
}

func toGrayscale8Image(t *terrain.Terrain) *image.Gray {
	i := image.NewGray(image.Rect(0, 0, int(t.X), int(t.Y)))
	t.Iterate(func(x uint16, y uint16, val uint16) {
		v := val / 256
		i.SetGray(int(x), int(y), color.Gray{Y: uint8(v)})
	})
	return i
}
