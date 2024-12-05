package main

import (
	_ "embed"
	//"fmt"
	//"math"
	"strings"
)

//go:embed resc/input.txt
var input string

type Direction int

const (
	DIR_N  Direction = iota
	DIR_NE Direction = iota
	DIR_E  Direction = iota
	DIR_SE Direction = iota
	DIR_S  Direction = iota
	DIR_SW Direction = iota
	DIR_W  Direction = iota
	DIR_NW Direction = iota
)

type Matrix [][]string

type CoordinatePair struct {
	X int
	Y int
}

type SearchResult struct {
	Position  CoordinatePair
	Direction Direction
}

func (m *Matrix) Rows() int {
	return len(*m)
}

func (m *Matrix) Columns() int {
	return len((*m)[0])
}

func (m *Matrix) At(coordinates CoordinatePair) string {
	return (*m)[coordinates.X][coordinates.Y]
}

func (m *Matrix) DirectionalSubstr(start CoordinatePair, size int, dir Direction) string {
	return ""
}

//
//func isValidSliceIndex(textLen int, textWidth int, sliceDir int, sliceStart int, index int) bool {
//	var valid bool
//	switch sliceDir {
//	case DIR_N, DIR_S:
//		valid = index >= 0 && index < textLen
//	case DIR_NE:
//		fmt.Printf("start:: index: %v, x: %v, y: %v\n", sliceStart, sliceStart%textWidth, sliceStart/textWidth)
//		fmt.Printf("end:: index: %v, x: %v, y: %v\n\n", index, index%textWidth, index/textWidth)
//		valid = index == sliceStart ||
//			index < textLen && index >= 0 &&
//				index/textWidth < sliceStart/textWidth &&
//				index%textWidth > sliceStart%textWidth
//	case DIR_E:
//		valid = index < textLen &&
//			index/textWidth == sliceStart/textWidth
//	case DIR_W:
//		valid = index >= 0 &&
//			index/textWidth == sliceStart/textWidth
//	}
//	return valid
//}
//
//func DirectionalSlice(text string, textLen int, textWidth int, dir int, start int, size int) string {
//	var step int
//	switch dir {
//	case DIR_N:
//		step = -(textWidth + 1) // +1 because of the newline character
//	case DIR_NE:
//		step = -(textWidth)
//	case DIR_E:
//		step = 1
//	case DIR_W:
//		step = -1
//	case DIR_S:
//		step = textWidth + 1
//	default:
//		return ""
//	}
//
//	var slice string
//	for i := range size {
//		index := start + i*step
//		if isValidSliceIndex(textLen, textWidth, dir, start, index) {
//			slice += string(text[index])
//		} else {
//			break
//		}
//	}
//
//	return slice
//}
//
//func directionalSearch(text string, textLen int, textWidth int, term string) []SearchResult {
//	var results []SearchResult
//	for i, char := range text {
//		if char != rune(term[0]) {
//			continue
//		}
//
//		for dir := range DIR_NW {
//			slice := DirectionalSlice(text, textLen, textWidth, dir, i, len(term))
//			if slice == term {
//				results = append(results, SearchResult{
//					Position:  i,
//					Direction: dir,
//				})
//			}
//		}
//	}
//	return results
//}

func ToMatrix(text string) Matrix {
	var matrix Matrix

	rows := strings.Split(text, "\n")
	for _, row := range rows {
		if len(row) != 0 {
			matrix = append(matrix, strings.Split(row, ""))
		}
	}

	return matrix
}

func main() {
	//inputMatrix := ToMatrix(input)
	//
	//	searchTerm := "XMAS"
	//
	//	width := strings.Index(input, "\n")
	//	results := directionalSearch(input, len(input), width, searchTerm)
	//
	//	fmt.Printf("Pt. 1: instances found: %v\n", len(results))
}
