package main

import (
	"reflect"
	"testing"
)

func TestToPageOrderRules(t *testing.T) {
	input = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13
`
	t.Run("should return correct number of PageOrderRules", func(t *testing.T) {
		got := len(ToPageOrderRules(input))
		want := 21

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("should return correct PageOrderRules", func(t *testing.T) {
		got := ToPageOrderRules(input)[0]
		want := PageOrderRule{47, 53}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("should return correct PageOrderRules 2", func(t *testing.T) {
		rules := ToPageOrderRules(input)
		got := rules[len(rules)-1]
		want := PageOrderRule{53, 13}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}

func TestToUpdates(t *testing.T) {
	input = `75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

	t.Run("should return correct number of Updates", func(t *testing.T) {
		got := len(ToUpdates(input))
		want := 6

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("should return correct Updates", func(t *testing.T) {
		updates := ToUpdates(input)
		got := updates[0]
		want := Update{75, 47, 61, 53, 29}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("should return correct Updates 2", func(t *testing.T) {
		updates := ToUpdates(input)
		got := updates[len(updates)-1]
		want := Update{97, 13, 75, 29, 47}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
