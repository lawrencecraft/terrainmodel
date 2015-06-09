package main

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestTerrainMaxSize(t *testing.T) {
	ter := InitializeTerrain(10, 65535)

	if ter.max != 1025 {
		t.Error("Expected 1025, got", ter.max)
	}
}

func TestTerrainLayout(t *testing.T) {
	ter := InitializeTerrain(10, 65535)

	if len(ter.layout) != 1025 {
		t.Error("Expected length 1025 but got", len(ter.layout))
	}

	for i := range ter.layout {
		if len(ter.layout[i]) != 1025 {
			t.Error("Expected length 1025 but got", len(ter.layout[i]), "on iteration", i)
		}
	}
}

func TestGetWhenInRange(t *testing.T) {
	ter := InitializeTerrain(10, 65535)
	ter.layout[22][22] = 3252

	num, err := ter.GetHeight(22, 22)
	if err != nil {
		t.Errorf("Should not have produced an error")
	}

	if num != 3252 {
		t.Errorf("Returned incorrect number")
	}
}

func TestGetWhenXAboveBounds(t *testing.T) {
	ter := InitializeTerrain(10, 65535)

	_, err := ter.GetHeight(22222, 22)
	if err == nil {
		t.Errorf("Should have produced an error")
	}
}

func TestGetWhenYAboveBounds(t *testing.T) {
	ter := InitializeTerrain(10, 65535)

	_, err := ter.GetHeight(22, 22222)
	if err == nil {
		t.Errorf("Should have produced an error")
	}
}

func TestGetWhenYAtBounds(t *testing.T) {
	ter := InitializeTerrain(10, 65535)

	_, err := ter.GetHeight(22, 1024)
	if err != nil {
		t.Errorf("Should not have produced an error")
	}
}

func TestDiamondStep(t *testing.T) {
	ter := InitializeTerrain(1, 65535)

	ter.SetHeight(0, 1, 1)
	ter.SetHeight(1, 0, 3)
	ter.SetHeight(2, 1, 0)
	ter.SetHeight(1, 2, 4)

	ter.SetDiamond(1, 1, 0, 1)

	v, e := ter.GetHeight(1, 1)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 2 {
		t.Error("Expected 2, got instead", v)
	}
}

func TestDiamondStepAtBounds(t *testing.T) {
	ter := InitializeTerrain(1, 65535)

	ter.SetHeight(0, 0, 2)
	ter.SetHeight(0, 1, 0)
	ter.SetHeight(1, 1, 4)

	ter.SetDiamond(1, 0, 0, 1)

	v, e := ter.GetHeight(1, 0)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 2 {
		t.Error("Expected 2, got instead", v)
	}
}
func TestDiamondStepWithOffset(t *testing.T) {
	ter := InitializeTerrain(1, 65535)

	ter.SetHeight(0, 0, 2)
	ter.SetHeight(0, 1, 0)
	ter.SetHeight(1, 1, 4)

	ter.SetDiamond(1, 0, 2, 1)

	v, e := ter.GetHeight(1, 0)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 4 {
		t.Error("Expected 4, got instead", v)
	}
}

func TestSquareStep(t *testing.T) {
	ter := InitializeTerrain(1, 65535)

	ter.SetHeight(0, 0, 1)
	ter.SetHeight(2, 0, 3)
	ter.SetHeight(0, 2, 0)
	ter.SetHeight(2, 2, 4)

	ter.SetSquare(1, 1, 0, 1)

	v, e := ter.GetHeight(1, 1)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 2 {
		t.Error("Expected 2, got instead", v)
	}
}

func TestSquareStepWithOffset(t *testing.T) {
	ter := InitializeTerrain(1, 65535)

	ter.SetHeight(0, 0, 1)
	ter.SetHeight(2, 0, 3)
	ter.SetHeight(0, 2, 0)
	ter.SetHeight(2, 2, 4)

	ter.SetSquare(1, 1, 2, 1)

	v, e := ter.GetHeight(1, 1)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 4 {
		t.Error("Expected 4, got instead", v)
	}
}

func TestSquareStepAtBounds(t *testing.T) {
	ter := InitializeTerrain(2, 65535)

	ter.SetHeight(2, 1, 2)
	ter.SetHeight(0, 1, 0)

	ter.SetSquare(1, 0, 0, 1)

	v, e := ter.GetHeight(1, 0)

	if e != nil {
		t.Error("Should not have gotten an error, got", e)
	}

	if v != 1 {
		t.Error("Expected 2, got instead", v)
	}
}

func BenchmarkDiamondSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ter := InitializeTerrain(10, 65535)

		ter.SetHeight(0, 0, 30000)
		ter.SetHeight(0, 1024, 30000)
		ter.SetHeight(1024, 0, 30000)
		ter.SetHeight(1024, 1024, 30000)

		ter.Generate(0.5, 30000, 30000, 30000, 30000)
	}
}
