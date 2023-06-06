//“Exercise 7.3:
//Write a String method for the *tree type in  gopl.io/ch4/treesort (§4.4) that reveals the
//sequence of values in the tree.”

package main

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var builder strings.Builder
	t.traverseInOrder(&builder)
	return builder.String()
}

func (t *tree) traverseInOrder(builder *strings.Builder) {
	if t != nil {
		t.left.traverseInOrder(builder)
		builder.WriteString(fmt.Sprintf("%d ", t.value))
		t.right.traverseInOrder(builder)
	}
}

func main() {
	values := []int{5, 2, 7, 1, 9, 3}
	Sort(values)

	t := &tree{}
	for _, value := range values {
		t = add(t, value)
	}

	fmt.Println(t) // Output: "1 2 3 5 7 9 "
}
