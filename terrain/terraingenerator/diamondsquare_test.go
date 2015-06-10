package terraingenerator

import (
	log "github.com/Sirupsen/logrus"
	"github.com/lawrencecraft/terrainmodel/terrain"
	"testing"
)

func TestDiamondStep(t *testing.T) {
	ter := terrain.New(1, 65535)

	ter.SetHeight(0, 1, 1)
	ter.SetHeight(1, 0, 3)
	ter.SetHeight(2, 1, 0)
	ter.SetHeight(1, 2, 4)

	setDiamond(ter, 1, 1, 0, 1)

	v, e := ter.GetHeight(1, 1)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 2 {
		t.Error("Expected 2, got instead", v)
	}
}

func TestDiamondStepAtBounds(t *testing.T) {
	ter := terrain.New(1, 65535)

	ter.SetHeight(0, 0, 2)
	ter.SetHeight(0, 1, 0)
	ter.SetHeight(1, 1, 4)

	setDiamond(ter, 1, 0, 0, 1)

	v, e := ter.GetHeight(1, 0)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 2 {
		t.Error("Expected 2, got instead", v)
	}
}
func TestDiamondStepWithOffset(t *testing.T) {
	ter := terrain.New(1, 65535)

	ter.SetHeight(0, 0, 2)
	ter.SetHeight(0, 1, 0)
	ter.SetHeight(1, 1, 4)

	setDiamond(ter, 1, 0, 2, 1)

	v, e := ter.GetHeight(1, 0)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 4 {
		t.Error("Expected 4, got instead", v)
	}
}

func TestSquareStep(t *testing.T) {
	ter := terrain.New(1, 65535)

	ter.SetHeight(0, 0, 1)
	ter.SetHeight(2, 0, 3)
	ter.SetHeight(0, 2, 0)
	ter.SetHeight(2, 2, 4)

	setSquare(ter, 1, 1, 0, 1)

	v, e := ter.GetHeight(1, 1)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 2 {
		t.Error("Expected 2, got instead", v)
	}
}

func TestSquareStepWithOffset(t *testing.T) {
	ter := terrain.New(1, 65535)

	ter.SetHeight(0, 0, 1)
	ter.SetHeight(2, 0, 3)
	ter.SetHeight(0, 2, 0)
	ter.SetHeight(2, 2, 4)

	setSquare(ter, 1, 1, 2, 1)

	v, e := ter.GetHeight(1, 1)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 4 {
		t.Error("Expected 4, got instead", v)
	}
}

func TestSquareStepAtBounds(t *testing.T) {
	ter := terrain.New(2, 65535)

	ter.SetHeight(2, 1, 2)
	ter.SetHeight(0, 1, 0)

	setSquare(ter, 1, 0, 0, 1)

	v, e := ter.GetHeight(1, 0)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 1 {
		t.Error("Expected 2, got instead", v)
	}
}

func BenchmarkDiamondSquare(b *testing.B) {
	log.SetLevel(log.InfoLevel)
	for i := 0; i < b.N; i++ {
		ter := terrain.New(10, 65535)

		ter.SetHeight(0, 0, 30000)
		ter.SetHeight(0, 1024, 30000)
		ter.SetHeight(1024, 0, 30000)
		ter.SetHeight(1024, 1024, 30000)

		GenerateTerrain(ter, 0.5, 30000, 30000, 30000, 30000)
	}
}
