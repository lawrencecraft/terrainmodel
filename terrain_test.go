package terrain

import (
	log "github.com/Sirupsen/logrus"
	"math"
	"testing"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestTerrainMaxSize(t *testing.T) {
	ter := New(1025, 100, math.MaxUint16)

	if ter.X != 1025 {
		t.Error("Expected X 1025, got ", ter.X)
	}

	if ter.Y != 100 {
		t.Error("Expected Y 100, got ", ter.Y)
	}
}

func TestTerrainLayout(t *testing.T) {
	ter := New(1025, 1025, math.MaxUint16)

	if len(ter.layout) != 1025*1025 {
		t.Error("Expected length 1025 but got", len(ter.layout))
	}
}

func TestGetWhenInRange(t *testing.T) {
	ter := New(1025, 1025, math.MaxUint16)
	ter.SetHeight(22, 22, 3252)
	num, err := ter.GetHeight(22, 22)
	if err != nil {
		t.Errorf("Should not have produced an error")
	}

	if num != 3252 {
		t.Errorf("Returned incorrect number")
	}
}

func TestGetWhenXAboveBounds(t *testing.T) {
	ter := New(1025, 1025, math.MaxUint16)

	_, err := ter.GetHeight(22222, 22)
	if err == nil {
		t.Errorf("Should have produced an error")
	}
}

func TestGetWhenYAboveBounds(t *testing.T) {
	ter := New(1025, 1025, math.MaxUint16)

	_, err := ter.GetHeight(22, 22222)
	if err == nil {
		t.Errorf("Should have produced an error")
	}
}

func TestGetWhenYAtBounds(t *testing.T) {
	ter := New(1025, 1025, math.MaxUint16)

	_, err := ter.GetHeight(22, 1024)
	if err != nil {
		t.Errorf("Should not have produced an error")
	}
}

func TestCopyOfTerrainWithNoOffsetThrowsNoError(t *testing.T) {
	src := New(10, 10, math.MaxUint16)
	dst := New(5, 5, math.MaxUint16)

	src.SetHeight(1, 1, 100)
	src.SetHeight(2, 1, 200)
	src.SetHeight(4, 4, 300)

	err := src.CopyTo(dst, 0, 0)

	if err != nil {
		t.Error("Got error when none expected: ", err)
	}
}

func TestCopyOfTerrainWithNoOffsetAssignsValues(t *testing.T) {
	src := New(10, 10, math.MaxUint16)
	dst := New(5, 5, math.MaxUint16)

	src.SetHeight(1, 1, 100)
	src.SetHeight(2, 1, 200)
	src.SetHeight(4, 4, 300)

	src.CopyTo(dst, 0, 0)

	firstHeight, _ := dst.GetHeight(1, 1)
	secondHeight, _ := dst.GetHeight(2, 1)
	thirdHeight, _ := dst.GetHeight(4, 4)

	if firstHeight != 100 {
		t.Error("Got incorrect height: expected", 100, "but got", firstHeight)
	}

	if secondHeight != 200 {
		t.Error("Got incorrect height: expected", 200, "but got", secondHeight)
	}

	if thirdHeight != 300 {
		t.Error("Got incorrect height: expected", 300, "but got", secondHeight)
	}
}

func TestCopyOfTerrainWithOffsetAssignsValues(t *testing.T) {
	src := New(10, 10, math.MaxUint16)
	dst := New(5, 5, math.MaxUint16)

	src.SetHeight(1, 1, 100)
	src.SetHeight(2, 1, 200)
	src.SetHeight(4, 4, 300)

	src.CopyTo(dst, 1, 1)

	firstHeight, _ := dst.GetHeight(0, 0)
	secondHeight, _ := dst.GetHeight(1, 0)
	thirdHeight, _ := dst.GetHeight(3, 3)

	if firstHeight != 100 {
		t.Error("Got incorrect height: expected", 100, "but got", firstHeight)
	}

	if secondHeight != 200 {
		t.Error("Got incorrect height: expected", 200, "but got", secondHeight)
	}

	if thirdHeight != 300 {
		t.Error("Got incorrect height: expected", 300, "but got", secondHeight)
	}
}

func TestCopyOfTerrainWithErrorCondition(t *testing.T) {
	src := New(10, 10, math.MaxUint16)
	dst := New(5, 5, math.MaxUint16)

	src.SetHeight(1, 1, 100)
	src.SetHeight(2, 1, 200)
	src.SetHeight(4, 4, 300)

	err := src.CopyTo(dst, 10, 10)

	if err == nil {
		t.Error("Should have got an error, but received nil")
	}
}
