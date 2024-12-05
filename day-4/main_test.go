package main

import (
	"testing"
)

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

	got := mat.At(CoordinatePair{0, 0})
	want := "M"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
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

	t.Run("DIR_N", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, DIR_N)
		want := "XMSM"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_NE", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, DIR_NE)
		want := "XXSM"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_E", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, DIR_E)
		want := "XXAM"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_SE", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, DIR_SE)
		want := "XSAM"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_S", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, DIR_S)
		want := "XAAX"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_SW", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, DIR_SW)
		want := "XSAM"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_W", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, DIR_W)
		want := "XMMA"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
	t.Run("DIR_NW", func(t *testing.T) {
		got := mat.DirectionalSubstr(start, size, DIR_NW)
		want := "XAMX"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

//func TestDirectionalSlice(t *testing.T) {
//	text := `MMMSXXMASM
//MSAMXMSMSA
//AMXSXMAAMM
//MSAMASMSMX
//XMASAMXAMM
//XXAMMXXAMA
//SMSMSASXSS
//SAXAMASAAA
//MAMMMXMMMM
//MXMXAXMASX
//`
//	textLen := len(text)
//	textWidth := strings.Index(text, "\n")
//
//	t.Run("DIR_N", func(t *testing.T) {
//		want := "XMAS"
//		got := DirectionalSlice(text, textLen, textWidth, DIR_N, 108, 4)
//
//		if got != want {
//			t.Errorf("got: %v, want: %v", got, want)
//		}
//	})
//
//	t.Run("DIR_NE", func(t *testing.T) {
//		want := "XMAS"
//		got := DirectionalSlice(text, textLen, textWidth, DIR_NE, 108, 4)
//
//		if got != want {
//			t.Errorf("got: %v, want: %v", got, want)
//		}
//	})
//
//	//	t.Run("DIR_E", func(t *testing.T) {
//	//		want := "XMAS"
//	//		got := DirectionalSlice(text, textLen, textWidth, DIR_E, 0, 4)
//	//
//	//		if got != want {
//	//			t.Errorf("got: %v, want: %v", got, want)
//	//		}
//	//	})
//	//
//	//	t.Run("DIR_SE", func(t *testing.T) {
//	//		want := "XMAS"
//	//		got := DirectionalSlice(text, textLen, textWidth, DIR_SE, 0, 4)
//	//
//	//		if got != want {
//	//			t.Errorf("got: %v, want: %v", got, want)
//	//		}
//	//	})
//	//
//	//	t.Run("DIR_S", func(t *testing.T) {
//	//		want := "XMAS"
//	//		got := DirectionalSlice(text, textLen, textWidth, DIR_S, 0, 4)
//	//
//	//		if got != want {
//	//			t.Errorf("got: %v, want: %v", got, want)
//	//		}
//	//	})
//	//
//	//	t.Run("DIR_SW", func(t *testing.T) {
//	//		want := "XMAS"
//	//		got := DirectionalSlice(text, textLen, textWidth, DIR_SW, 0, 4)
//	//
//	//		if got != want {
//	//			t.Errorf("got: %v, want: %v", got, want)
//	//		}
//	//	})
//	//
//	//	t.Run("DIR_W", func(t *testing.T) {
//	//		want := "XMAS"
//	//		got := DirectionalSlice(text, textLen, textWidth, DIR_W, 0, 4)
//	//
//	//		if got != want {
//	//			t.Errorf("got: %v, want: %v", got, want)
//	//		}
//	//	})
//	//
//	//	t.Run("DIR_NW", func(t *testing.T) {
//	//		want := "XMAS"
//	//		got := DirectionalSlice(text, textLen, textWidth, DIR_NW, 0, 4)
//	//
//	//		if got != want {
//	//			t.Errorf("got: %v, want: %v", got, want)
//	//		}
//	//	})
//}
