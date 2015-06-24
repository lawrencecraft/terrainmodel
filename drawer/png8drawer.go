package drawer

import (
	"github.com/lawrencecraft/terrainmodel"
	"image/png"
	"io"
)

type Png8Drawer struct {
	Writer io.Writer
}

func NewPng8Drawer(writer io.Writer) *PngDrawer {
	return &PngDrawer{Writer: writer}
}

func (pngd *Png8Drawer) Draw(t *terrain.Terrain) error {
	return png.Encode(pngd.Writer, toGrayscale8Image(t))
}
