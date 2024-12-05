package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	coords := CoordinatePair{0, 0}

	got := coords.Add(Vector{1, 1})
	want := CoordinatePair{1, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, wanted: %v", got, want)
	}
}

func TestMult(t *testing.T) {
	vec := Vector{1, 2}

	got := vec.Mult(2)
	want := Vector{2, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, wanted: %v", got, want)
	}
}

func TestRows(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	mat := ToMatrix(input)

	got := mat.Rows()
	want := 10

	if got != want {
		t.Errorf("got %v rows, wanted %v", got, want)
	}
}
func TestColumns(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	mat := ToMatrix(input)

	got := mat.Columns()
	want := 10

	if got != want {
		t.Errorf("got %v columns, wanted %v", got, want)
	}
}

func TestAt(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	mat := ToMatrix(input)

	t.Run("should return string at valid coordinates", func(t *testing.T) {
		got, _ := mat.At(CoordinatePair{0, 0})
		want := "M"

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("should return err at invalid coordinates", func(t *testing.T) {
		_, err := mat.At(CoordinatePair{-1, 0})

		if err != nil {
			t.Error("expected error")
		}
	})

	t.Run("should return err at invalid coordinates", func(t *testing.T) {
		_, err := mat.At(CoordinatePair{11, 0})

		if err != nil {
			t.Error("expected error")
		}
	})
}

func TestDirectionalSubstr(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	mat := ToMatrix(input)
	start := CoordinatePair{X: 5, Y: 5}
	size := 4

	dirs := DirectionVectors()

	t.Run("DIR_N", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, dirs[DIR_N])
		want := "XMSM"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_NE", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, dirs[DIR_NE])
		want := "XXSM"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_E", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, dirs[DIR_E])
		want := "XXAM"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_SE", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, dirs[DIR_SE])
		want := "XSAM"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_S", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, dirs[DIR_S])
		want := "XAAX"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_SW", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, dirs[DIR_SW])
		want := "XSAM"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_W", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, dirs[DIR_W])
		want := "XMMA"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_NW", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, dirs[DIR_NW])
		want := "XAMX"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestWordSearch(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`
	term := "XMAS"

	results := WordSearch(input, term)
	got := len(results)
	want := 18

	if got != want {
		t.Errorf("found %v results, %v wanted", got, want)
	}
}
