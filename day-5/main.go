package main

import (
	_ "embed"
	"strconv"
	"strings"
)

type PageOrderRule [2]int

type Update []int

func (update *Update) IsValid(rules []PageOrderRule) bool {
	ruleMap := map[int][]int{}
	for _, rule := range rules {
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}

	for i, num := range *update {
		mayNotPreceed := ruleMap[num]
		for p := range i {
			for _, n := range mayNotPreceed {
				if (*update)[p] == n {
					return false
				}
			}
		}
	}

	return true
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

	print(rules)
	print(updates)
}
