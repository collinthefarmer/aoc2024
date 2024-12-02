package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed resc/input.txt
var input string

func absDiff(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func SplitReports(text string) []string {
	return strings.Split(text, "\n")
}

func SplitLevels(report string) (levels []int) {
	for _, s := range strings.Split(report, " ") {
		v, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		levels = append(levels, v)
	}
	return
}

func AreSafeLevels(report []int) bool {
	asc := true
	last := 0
	for i, v := range report {
		diff := absDiff(last, v)
		safeDiff := diff <= 3 && diff >= 1

		if !safeDiff && i > 0 {
			return false
		} else if asc && v > last {
			last = v
			continue
		} else if !asc && v < last {
			last = v
			continue
		} else if asc && i == 1 {
			asc = false
			last = v
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	var safeCount int
	for _, report := range SplitReports(input) {
		if len(report) == 0 {
			continue
		}

		levels := SplitLevels(report)
		if AreSafeLevels(levels) {
			safeCount += 1
		}
	}

	log.Printf("Pt. 1: number of safe reports: %v", safeCount)
}
