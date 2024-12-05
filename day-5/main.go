package main

import (
	_ "embed"
	"strconv"
	"strings"
)

type PageOrderRule [2]int

type Update []int

func ToPageOrderRules(text string) []PageOrderRule {
	var rules []PageOrderRule
	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			continue
		}
		split := strings.Split(line, "|")

		rule0, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		rule1, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		rules = append(rules, [2]int{rule0, rule1})
	}
	return rules
}

//go:embed resc/input.txt
var input string

func main() {
}
