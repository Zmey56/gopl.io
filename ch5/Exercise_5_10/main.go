//“Exercise 5.10:
//Rewrite topoSort to use maps instead of slices and eliminate the
//initial sort.
//
//Verify that the results, though nondeterministic, are valid
//topological orderings.”

package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(string)

	visitAll = func(item string) {
		if !seen[item] {
			seen[item] = true
			for _, dependency := range m[item] {
				visitAll(dependency)
			}
			order = append(order, item)
		}
	}

	keys := make(map[string]bool)
	for _, dependencies := range m {
		for _, dep := range dependencies {
			keys[dep] = true
		}
	}

	for key := range m {
		if !keys[key] {
			visitAll(key)
		}
	}

	return order
}
