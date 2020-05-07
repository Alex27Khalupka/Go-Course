package main

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	testCases := []struct {
		in  Figure
		out float64
	}{
		{Square{10}, 100},
		{Circle{10}, 100 * math.Pi},
		{Square{5}, 25},
		{Circle{5}, 25 * math.Pi},
	}

	for _, c := range testCases {
		res := c.in.area()
		if res != c.out {
			t.Errorf("area(%d) returns %v expected %v", c.in, res, c.out)
		}
	}
}

func TestPerimeter(t *testing.T) {
	testCases := []struct {
		in  Figure
		out float64
	}{
		{Square{10}, 40},
		{Circle{10}, 20 * math.Pi},
		{Square{5}, 20},
		{Circle{5}, 10 * math.Pi},
	}

	for _, c := range testCases {
		res := c.in.perimeter()
		if res != c.out {
			t.Errorf("area(%f) returns %f expected %f", c.in, res, c.out)
		}
	}
}
