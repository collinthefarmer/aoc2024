package main

import (
	_ "embed"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed resc/input.txt
var input string

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func sum(vs []int) int {
	var sum = 0
	for _, v := range vs {
		sum += v
	}
	return sum
}

func countOccurences(vs []int) map[int]int {
	var occurences = map[int]int{}

	for _, v := range vs {
		occurences[v] = occurences[v] + 1
	}

	return occurences
}

func main() {

	// split input into two int arrays

	var ids0 []int
	var ids1 []int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		ids := strings.Split(line, "   ")
		if len(ids) == 2 {
			int0, err := strconv.Atoi(ids[0])
			if err != nil {
				panic(err)
			}

			int1, err := strconv.Atoi(ids[1])
			if err != nil {
				panic(err)
			}

			ids0 = append(ids0, int0)
			ids1 = append(ids1, int1)
		}
	}

	// sort int arrays into asc. order

	sort.Ints(ids0)
	sort.Ints(ids1)

	// calculate differences

	var diffs []int

	for i, a := range ids0 {
		b := ids1[i]
		diffs = append(diffs, abs(a-b))
	}

	// sum differences for final answer of Pt. 1
	diffSum := sum(diffs)
	log.Printf("Pt. 1:: sum of ID differences: %v", diffSum)

	// occurences of each number in right list

	occurences := countOccurences(ids1)

	// calculate "similiarity score"
	var similiarityScore int
	for _, v := range ids0 {
		similiarityScore += v * occurences[v]
	}

	log.Printf("Pt. 2:: sum of similiarity scores: %v", similiarityScore)
}
