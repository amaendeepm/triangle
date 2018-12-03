package trianglen

import (
	"testing"
)

func TestIsTrianglePossible ( t *testing.T) {
	if IsTriangle(5,1,6) {
		t.Error("ERR(t1): Triangle was not possible here")
	}

	if !IsTriangle(5,2,6) {
		t.Error("ERR(t2): Triangle was possible here")
	}

	if !IsEquilateral(5,5,5) {
		t.Error("ERR(t3): Triangle was equilateral here")
	}

	if IsEquilateral(5,6,5) {
		t.Error("ERR(t4): Triangle was not equilateral here")
	}

	if !IsIsoceles(5,5,1) {
		t.Error("ERR(t5): Triangle was isoceles here")
	}

	if CheckTriangleType(5,6,1.5) != "Scalene" {
		t.Error("ERR(t6): Triangle Type should return Scalene")
	}

	if CheckTriangleType(15,1,15) != "Isoceles" {
		t.Error("ERR(t7): Triangle Type should return Isoceles")
	}

	if CheckTriangleType(15,0,15) != "Impossible" {
		t.Error("ERR(t8): Triangle Type should be Impossible i.e. Trianle not possible should have returned")
	}

}
