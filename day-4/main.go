package main

import (
	_ "embed"
	"fmt"
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

type CoordinatePair struct {
	X int
	Y int
}

func (cp *CoordinatePair) Add(v Vector) CoordinatePair {
	return CoordinatePair{
		X: cp.X + v.X,
		Y: cp.Y + v.Y,
	}
}

type Vector struct {
	X int
	Y int
}

func (v *Vector) Mult(value int) Vector {
	return Vector{
		X: v.X * value,
		Y: v.Y * value,
	}
}

type SearchResult struct {
	Position  CoordinatePair
	Direction Direction
}

type Matrix [][]string

func (m *Matrix) Rows() int {
	return len(*m)
}

func (m *Matrix) Columns() int {
	return len((*m)[0])
}

func (m *Matrix) At(coordinates CoordinatePair) (string, error) {
	mat := *m
	if coordinates.X >= mat.Columns() ||
		coordinates.X < 0 ||
		coordinates.Y >= mat.Rows() ||
		coordinates.Y < 0 {
		return "", nil
	}
	return mat[coordinates.Y][coordinates.X], nil
}

func (m *Matrix) DirectionalSubstr(start CoordinatePair, size int, step Vector) string {
	var substr string
	for i := range size {
		char, err := m.At(start.Add(step.Mult(i)))
		if err != nil {
			break
		}
		substr += char
	}
	return substr
}

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

func DirectionVectors() map[Direction]Vector {
	return map[Direction]Vector{
		DIR_N:  {Y: -1},
		DIR_NE: {X: 1, Y: -1},
		DIR_E:  {X: 1},
		DIR_SE: {X: 1, Y: 1},
		DIR_S:  {Y: 1},
		DIR_SW: {X: -1, Y: 1},
		DIR_W:  {X: -1},
		DIR_NW: {X: -1, Y: -1},
	}
}

func WordSearch(text string, term string) []SearchResult {
	matrix := ToMatrix(text)
	termLen := len(term)

	var results []SearchResult
	for dir, vector := range DirectionVectors() {
		for iy := range matrix.Rows() {
			for ix := range matrix.Columns() {
				start := CoordinatePair{ix, iy}
				substr := matrix.DirectionalSubstr(start, termLen, vector)
				if substr == term {
					results = append(results, SearchResult{
						Direction: dir,
						Position:  start,
					})
				}
			}
		}
	}

	return results
}

func main() {
	searchText := input
	searchTerm := "XMAS"

	results := WordSearch(searchText, searchTerm)
	fmt.Printf("Pt. 1: instances found: %v\n", len(results))
}
