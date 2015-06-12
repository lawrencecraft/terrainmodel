package generator

import (
	log "github.com/Sirupsen/logrus"
	terrain "github.com/lawrencecraft/terrainmodel"
	"testing"
)

func TestDiamondStep(t *testing.T) {
	ter := terrain.New(3, 3, 65535)

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
	ter := terrain.New(3, 3, 65535)

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
	ter := terrain.New(3, 3, 65535)

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
	ter := terrain.New(3, 3, 65535)

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
	ter := terrain.New(3, 3, 65535)

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
	ter := terrain.New(5, 5, 65535)

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

func BenchmarkSC4DiamondSquare(b *testing.B) {
	log.SetLevel(log.InfoLevel)
	d := NewDiamondSquareGenerator(0.5, 1025, 1025)
	for i := 0; i < b.N; i++ {
		d.Generate()
	}
}

func BenchmarkCSDiamondSquare(b *testing.B) {
	log.SetLevel(log.InfoLevel)
	d := NewDiamondSquareGenerator(0.5, 1028, 1028)
	for i := 0; i < b.N; i++ {
		d.Generate()
	}
}

func TestScaleIsSetCorrectly(t *testing.T) {
	scale := getScale(1000, 1000)
	if scale != 10 {
		t.Error("Expected 10 but got ", scale)
	}
}

func TestScaleWorksAtBoundaries(t *testing.T) {
	scale := getScale(1025, 1025)
	if scale != 10 {
		t.Error("Exptected 10 but got ", scale)
	}
}

func TestCanGenerateAtBounds(t *testing.T) {
	d := NewDiamondSquareGenerator(0.5, 5, 5)
	_, err := d.Generate()

	if err != nil {
		t.Error("Got error generating: ", err)
	}
}
