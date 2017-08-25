package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	if result := GetTriangleType(1, 2, 3); result != Error {
		t.Errorf("GetTriangleType was incorrect, got: %d, want: %d.", result, Error)
	}
	if result := GetTriangleType(1, 2, 0); result != Error {
		t.Errorf("GetTriangleType was incorrect, got: %d, want: %d.", result, Error)
	}
	if result := GetTriangleType(3, 3, 3); result != Equilateral {
		t.Errorf("GetTriangleType was incorrect, got: %d, want: %d.", result, Equilateral)
	}
	if result := GetTriangleType(2, 3, 3); result != Isosceles {
		t.Errorf("GetTriangleType was incorrect, got: %d, want: %d.", result, Isosceles)
	}
	if result := GetTriangleType(2, 3, 4); result != Scalene {
		t.Errorf("GetTriangleType was incorrect, got: %d, want: %d.", result, Scalene)
	}
}
