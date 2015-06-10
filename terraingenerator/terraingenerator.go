package terraingenerator

import "github.com/lawrencecraft/terrainmodel/terrain"

type TerrainGenerator interface {
	Generate() (*terrain.Terrain, error)
}
