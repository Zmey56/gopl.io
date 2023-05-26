//“Exercise 5.11:
//The instructor of the linear algebra course decides that calculus is now a
//prerequisite.  Extend the topoSort function to report cycles.”

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
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Println("Cycle detected in prerequisites:", err)
		return
	}

	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	path := make(map[string]bool)

	var visitAll func(string) error

	visitAll = func(item string) error {
		if path[item] {
			return fmt.Errorf("cycle detected for course: %s", item)
		}

		if !seen[item] {
			seen[item] = true
			path[item] = true

			for _, dependency := range m[item] {
				err := visitAll(dependency)
				if err != nil {
					return err
				}
			}

			delete(path, item)
			order = append(order, item)
		}

		return nil
	}

	keys := make(map[string]bool)
	for _, dependencies := range m {
		for _, dep := range dependencies {
			keys[dep] = true
		}
	}

	for key := range m {
		if !keys[key] {
			err := visitAll(key)
			if err != nil {
				return nil, err
			}
		}
	}

	return order, nil
}
