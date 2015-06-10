package terraingenerator

import "github.com/lawrencecraft/terrainmodel"

type TerrainGenerator interface {
	Generate() (*terrain.Terrain, error)
}
