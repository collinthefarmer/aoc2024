package main

import (
	"strings"
	"testing"
)

func TestDirectionalSlice(t *testing.T) {
	text := `MMMSXXMASM
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
	textLen := len(text)
	textWidth := strings.Index(text, "\n")

	t.Run("DIR_N", func(t *testing.T) {
		want := "XMAS"
		got := DirectionalSlice(text, textLen, textWidth, DIR_N, 108, 4)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("DIR_NE", func(t *testing.T) {
		want := "XMAS"
		got := DirectionalSlice(text, textLen, textWidth, DIR_NE, 104, 4)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("DIR_E", func(t *testing.T) {
		want := "XMAS"
		got := DirectionalSlice(text, textLen, textWidth, DIR_E, 0, 4)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("DIR_SE", func(t *testing.T) {
		want := "XMAS"
		got := DirectionalSlice(text, textLen, textWidth, DIR_SE, 0, 4)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("DIR_S", func(t *testing.T) {
		want := "XMAS"
		got := DirectionalSlice(text, textLen, textWidth, DIR_S, 0, 4)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("DIR_SW", func(t *testing.T) {
		want := "XMAS"
		got := DirectionalSlice(text, textLen, textWidth, DIR_SW, 0, 4)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("DIR_W", func(t *testing.T) {
		want := "XMAS"
		got := DirectionalSlice(text, textLen, textWidth, DIR_W, 0, 4)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("DIR_NW", func(t *testing.T) {
		want := "XMAS"
		got := DirectionalSlice(text, textLen, textWidth, DIR_NW, 0, 4)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
