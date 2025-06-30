package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape   Shape
		hasArea float64
		title   string
	}{
		{shape: Rectangle{12, 6}, hasArea: 72.0, title: "Rectangle"},
		{shape: Circle{10}, hasArea: 314.1592653589793, title: "Circle"},
		{shape: Triangle{12, 6}, hasArea: 36.0, title: "Triangle"},
	}

	for _, tt := range areaTests {
		t.Run(tt.title, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
