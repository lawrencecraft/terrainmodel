package generator

import "github.com/lawrencecraft/terrainmodel"

type Generator interface {
	Generate() (*terrain.Terrain, error)
}
