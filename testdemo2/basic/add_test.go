package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{1, 2, 0},
		{0, 2, 0},
		{0, 0, 0},
		{-1, 1, 0},
		//{math.MaxInt32, 1, math.MinInt32},
	}

	for _, test := range tests {
		if actual := Add(test.a, test.b); actual != test.c {
			t.Errorf("actual:%d,test is %d", actual, test.c)
		}
	}
}

func TestCalcTriangele(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
	}

	for _, test := range tests {
		if actual := CalcTriangele(test.a, test.b); actual != test.c {
			t.Errorf("actual:%d,test is %d", actual, test.c)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	x := 1
	y := 1
	z := 0
	for i := 0; i < b.N; i++ {
		actual := Add(x, y)
		if actual != z {
			b.Errorf("actual:%d,test is %d", actual, z)
		}
	}
}
