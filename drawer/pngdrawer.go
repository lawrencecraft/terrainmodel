package drawer

import (
	"github.com/lawrencecraft/terrainmodel"
	"image/png"
	"io"
)

type PngDrawer struct {
	Writer io.Writer
}

func NewPngDrawer(writer io.Writer) *PngDrawer {
	return &PngDrawer{Writer: writer}
}

func (pngd *PngDrawer) Draw(t *terrain.Terrain) error {
	return png.Encode(pngd.Writer, toGrayscale16Image(t))
}
