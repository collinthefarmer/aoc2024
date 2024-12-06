package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

type Vector struct {
	X int
	Y int
}

func (v Vector) Rotated(mat [][]int) Vector {
	return Vector{
		X: v.X*mat[0][0] + v.Y*mat[0][1],
		Y: v.X*mat[1][0] + v.Y*mat[1][1],
	}
}

type LabMapPosition struct {
	X int
	Y int
}

func (pos LabMapPosition) Add(vec Vector) LabMapPosition {
	return LabMapPosition{
		X: pos.X + vec.X,
		Y: pos.Y + vec.Y,
	}
}

type LabMap struct {
	Layout        [][]bool
	GuardLocation LabMapPosition
}

func (lab *LabMap) ReadText(text, obstruction, guard string) {
	for iy, line := range strings.Split(text, "\n") {
		if line == "" {
			continue
		}

		var row []bool
		for ix, char := range line {
			row = append(row, string(char) == obstruction)
			if string(char) == guard {
				lab.GuardLocation.X, lab.GuardLocation.Y = ix, iy
			}
		}

		lab.Layout = append(lab.Layout, row)
	}
}

func (lab *LabMap) IsValidPosition(pos LabMapPosition) bool {
	return !(pos.X < 0 ||
		pos.Y < 0 ||
		pos.X >= len(lab.Layout[0]) ||
		pos.Y >= len(lab.Layout))
}

func (lab *LabMap) IsObstructed(pos LabMapPosition) bool {
	return lab.Layout[pos.Y][pos.X]
}

func (lab *LabMap) StepGuard(dir Vector, obstructionRotation [][]int) (Vector, error) {
	nextPos := lab.GuardLocation.Add(dir)

	if !lab.IsValidPosition(nextPos) {
		return dir, fmt.Errorf("invalid position %v - guard has left the map!", nextPos)
	}

	if lab.IsObstructed(nextPos) {
		return dir.Rotated(obstructionRotation), nil
	} else {
		lab.GuardLocation = nextPos
		return dir, nil
	}
}

func (lab *LabMap) stringify(guardDir Vector) string {
	var out string
	for iy, row := range lab.Layout {
		var srow string
		for ix, cell := range row {
			if lab.GuardLocation.X == ix && lab.GuardLocation.Y == iy {
				var guardSymbol string
				switch guardDir {
				case Vector{0, 1}:
					guardSymbol = "v"
				case Vector{1, 0}:
					guardSymbol = ">"
				case Vector{-1, 0}:
					guardSymbol = "<"
				case Vector{0, -1}:
					guardSymbol = "^"
				}
				srow += "\033[91m" + guardSymbol + "\033[0m"
			} else if cell {
				srow += "#"
			} else {
				srow += "."
			}
		}
		out += srow + "\n"
	}
	return out
}

func (lab *LabMap) Display(guardDir Vector) {
	fmt.Printf("\033[?25l\033[1J\033[H%v", lab.stringify(guardDir))
	time.Sleep(17 * time.Millisecond)
}

//go:embed resc/input.txt
var input string

func main() {
	rightRotMat := [][]int{{0, -1}, {1, 0}}

	labMap := LabMap{}
	labMap.ReadText(input, "#", "^")

	var positions = map[LabMapPosition]bool{}
	positions[labMap.GuardLocation] = true // we count the guard's first position

	dir := Vector{X: 0, Y: -1}
	var err error
	for err == nil {
		dir, err = labMap.StepGuard(dir, rightRotMat)
		positions[labMap.GuardLocation] = true
		labMap.Display(dir)
	}

	fmt.Printf("pt. 1: distinct positions: %v", len(positions))
}
