package main

import (
	_ "embed"
	"fmt"
	//"math"
	"strings"
)

//go:embed resc/input.txt
var input string

const (
	DIR_N  = iota
	DIR_NE = iota
	DIR_E  = iota
	DIR_SE = iota
	DIR_S  = iota
	DIR_SW = iota
	DIR_W  = iota
	DIR_NW = iota
)

type SearchResult struct {
	Position  int
	Direction int
}

func isValidSliceIndex(textLen int, textWidth int, sliceDir int, sliceStart int, index int) bool {
	var valid bool
	switch sliceDir {
	case DIR_N, DIR_S:
		valid = index >= 0 && index < textLen
	case DIR_NE:
		valid = index < textLen && index >= 0 &&
			index/textWidth < sliceStart/textWidth &&
			index%textWidth > sliceStart%textWidth
	case DIR_E:
		valid = index < textLen &&
			index/textWidth == sliceStart/textWidth
	case DIR_W:
		valid = index >= 0 &&
			index/textWidth == sliceStart/textWidth
	}
	return valid
}

func DirectionalSlice(text string, textLen int, textWidth int, dir int, start int, size int) string {
	var step int
	switch dir {
	case DIR_N:
		step = -(textWidth + 1) // +1 because of the newline character
	case DIR_NE:
		step = -(textWidth)
	case DIR_E:
		step = 1
	case DIR_W:
		step = -1
	case DIR_S:
		step = textWidth + 1
	default:
		return ""
	}

	var slice string
	for i := range size {
		index := start + i*step
		if isValidSliceIndex(textLen, textWidth, dir, start, index) {
			slice += string(text[index])
		} else {
			break
		}
	}

	return slice
}

func search(text string, textLen int, textWidth int, term string) []SearchResult {
	var results []SearchResult
	for i, char := range text {
		if char != rune(term[0]) {
			continue
		}

		for dir := range DIR_NW {
			slice := DirectionalSlice(text, textLen, textWidth, dir, i, len(term))
			if slice == term {
				results = append(results, SearchResult{
					Position:  i,
					Direction: dir,
				})
			}
		}
	}
	return results
}

func main() {
	searchTerm := "XMAS"

	width := strings.Index(input, "\n")
	results := search(input, len(input), width, searchTerm)

	fmt.Printf("Pt. 1: instances found: %v\n", len(results))
}
