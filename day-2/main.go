package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed resc/input.txt
var input string

//func absDiff(a int, b int) int {
//	if a > b {
//		return a - b
//	} else {
//		return b - a
//	}
//}

func remove(slice []int, index int) []int {
	removed := make([]int, 0, len(slice)-1)
	removed = append(removed, slice[:index]...)
	return append(removed, slice[index+1:]...)
}

func SplitReports(text string) []string {
	return strings.Split(text, "\n")
}

func SplitLevels(report string) (levels []int) {
	for _, s := range strings.Split(report, " ") {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		levels = append(levels, v)
	}
	return
}

func AreSafeLevels(report []int, useProblemDampener bool) bool {
	dir := 1
	var prev int
	for i, v := range report {
		diff := (v - prev) * dir
		if diff < 0 && i == 1 {
			if useProblemDampener && AreSafeLevels(remove(report, i), false) {
				return true
			} else {
				dir = -1
				diff *= dir
			}
		}

		if (diff > 3 || diff <= 0) && i != 0 {
			return useProblemDampener && AreSafeLevels(remove(report, i), false)
		}

		prev = v
	}
	return true
}

func main() {
	var safeCount int
	var safeCountProblemDampened int
	for _, report := range SplitReports(input) {
		if len(report) == 0 {
			continue
		}

		levels := SplitLevels(report)
		if AreSafeLevels(levels, false) {
			safeCount += 1
		}

		if AreSafeLevels(levels, true) {
			safeCountProblemDampened += 1
		}
		log.Printf("%v", levels)
	}

	log.Printf("Pt. 1: number of safe reports: %v", safeCount)
	log.Printf("Pt. 2: number of safe reports (problem dampened): %v", safeCountProblemDampened)
}
