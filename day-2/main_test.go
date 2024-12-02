package main

import (
	"reflect"
	"testing"
)

func TestSplitReports(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

		want := []string{
			"7 6 4 2 1",
			"1 2 7 8 9",
			"9 7 6 2 1",
			"1 3 2 4 5",
			"8 6 4 4 1",
			"1 3 6 7 9",
		}
		got := SplitReports(input)

		if !reflect.DeepEqual(want, got) {
			t.Error()
		}
	})
}

func TestSplitLevels(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := "7 6 4 2 1"

		want := []int{7, 6, 4, 2, 1}
		got := SplitLevels(input)

		if !reflect.DeepEqual(want, got) {
			t.Error()
		}
	})
}

func TestAreSafeLevels(t *testing.T) {
	t.Run("7 6 4 2 1", func(t *testing.T) {
		want := true
		got := AreSafeLevels([]int{7, 6, 4, 2, 1}, false)

		if want != got {
			t.Error()
		}
	})

	t.Run("1 2 7 8 9", func(t *testing.T) {
		want := false
		got := AreSafeLevels([]int{1, 2, 7, 8, 8}, false)

		if want != got {
			t.Error()
		}
	})

	t.Run("9 7 6 2 1", func(t *testing.T) {
		want := false
		got := AreSafeLevels([]int{9, 7, 6, 2, 1}, false)

		if want != got {
			t.Error()
		}
	})

	t.Run("1 3 2 4 5", func(t *testing.T) {
		want := false
		got := AreSafeLevels([]int{1, 3, 2, 4, 5}, false)

		if want != got {
			t.Error()
		}
	})

	t.Run("1 3 2 4 5 (w/ prob. damp)", func(t *testing.T) {
		want := true
		got := AreSafeLevels([]int{1, 3, 2, 4, 5}, true)

		if want != got {
			t.Error()
		}
	})

	t.Run("8 6 4 4 1", func(t *testing.T) {
		want := false
		got := AreSafeLevels([]int{8, 6, 4, 4, 1}, false)

		if want != got {
			t.Error()
		}
	})

	t.Run("8 6 4 4 1 (w/ prob. damp)", func(t *testing.T) {
		want := true
		got := AreSafeLevels([]int{8, 6, 4, 4, 1}, true)

		if want != got {
			t.Error()
		}
	})

	t.Run("1 3 6 7 9", func(t *testing.T) {
		want := true
		got := AreSafeLevels([]int{1, 3, 6, 7, 9}, false)
		if want != got {
			t.Error()
		}
	})
}
