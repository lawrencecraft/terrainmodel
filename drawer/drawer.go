package drawer

import (
	"github.com/lawrencecraft/terrainmodel"
)

type Drawer interface {
	Draw(t *terrain.Terrain) error
}
