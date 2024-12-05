package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type PageOrderRule [2]int

type Update []int

func (u *Update) IsValid(rules []PageOrderRule) bool {
	ruleMap := map[int][]int{}
	for _, rule := range rules {
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}

	update := *u
	for i, num := range update {
		mayNotPreceed := ruleMap[num]

		for p := range i {
			for _, n := range mayNotPreceed {
				if update[p] == n {
					return false
				}
			}
		}
	}

	return true
}

func (u *Update) SortToValid(rules []PageOrderRule) {
	ruleMap := map[int][]int{}
	for _, rule := range rules {
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}

	var changed bool
	update := *u
	for i, num := range update {
		mayNotPreceed := ruleMap[num]

		for p := range i {
			for _, n := range mayNotPreceed {
				if update[p] == n {
					update[i], update[p] = update[p], update[i]
					changed = true
				}
			}
		}
	}

	if changed {
		u.SortToValid(rules)
	}
}

func (update *Update) MiddlePage() int {
	return (*update)[len(*update)/2]
}

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

func ToUpdates(text string) []Update {
	var updates []Update
	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			continue
		}

		var update Update
		for _, s := range strings.Split(line, ",") {
			updateI, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			update = append(update, updateI)
		}
		updates = append(updates, update)
	}
	return updates
}

//go:embed resc/input.txt
var input string

func main() {
	components := strings.Split(input, "\n\n")

	rules := ToPageOrderRules(components[0])
	updates := ToUpdates(components[1])

	// pt.1

	var middlePageSum int
	for _, update := range updates {
		if update.IsValid(rules) {
			middlePageSum += update.MiddlePage()
		}
	}

	fmt.Printf("Pt. 1: Sum of middle pages: %v\n", middlePageSum)

	// pt.2

	var partTwoMiddlePageSum int
	for _, update := range updates {
		if update.IsValid(rules) {
			continue
		}

		update.SortToValid(rules)
		partTwoMiddlePageSum += update.MiddlePage()
	}

	fmt.Printf("Pt. 2: Sum of middle pages: %v\n", partTwoMiddlePageSum)
}
