package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSquareGetArea(t *testing.T) {
	s := Square{sideLength: 2}

	want := 4.0
	got := s.getArea()
	if got != want {
		t.Errorf("square.GetArea() = %v, want %v", got, want)
	}
}

func TestTriangleGetArea(t *testing.T) {
	tr := Triangle{base: 2, height: 2}

	want := 2.0
	got := tr.getArea()
	if got != want {
		t.Errorf("square.GetArea() = %v, want %v", got, want)
	}
}
