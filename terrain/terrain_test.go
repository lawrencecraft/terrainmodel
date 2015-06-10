package terrain

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestTerrainMaxSize(t *testing.T) {
	ter := New(10, 65535)

	if ter.Max != 1025 {
		t.Error("Expected 1025, got", ter.Max)
	}
}

func TestTerrainLayout(t *testing.T) {
	ter := New(10, 65535)

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
	ter := New(10, 65535)
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
	ter := New(10, 65535)

	_, err := ter.GetHeight(22222, 22)
	if err == nil {
		t.Errorf("Should have produced an error")
	}
}

func TestGetWhenYAboveBounds(t *testing.T) {
	ter := New(10, 65535)

	_, err := ter.GetHeight(22, 22222)
	if err == nil {
		t.Errorf("Should have produced an error")
	}
}

func TestGetWhenYAtBounds(t *testing.T) {
	ter := New(10, 65535)

	_, err := ter.GetHeight(22, 1024)
	if err != nil {
		t.Errorf("Should not have produced an error")
	}
}
