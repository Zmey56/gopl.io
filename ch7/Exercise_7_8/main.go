//Exercise 7.8:
//Many GUIs provide a table widget with a stateful multi-tier sort: the
//primary sort key is the most recently clicked column head, the
//secondary sort key is the second-most recently clicked column head,
//and so on.  Define an implementation of sort.Interface for
//use by such a table. Compare that approach with repeated sorting using sort.Stable.

package main

import (
	"fmt"
	"sort"
)

type Table struct {
	data       []sort.Interface
	sortOrders [][]int
}

func (t *Table) Len() int {
	return t.data[0].Len()
}

func (t *Table) Less(i, j int) bool {
	a, b := t.getData(i), t.getData(j)

	for _, sortOrder := range t.sortOrders {
		for _, col := range sortOrder {
			if a[col] != b[col] {
				return a[col] < b[col]
			}
		}
	}

	return false
}

func (t *Table) Swap(i, j int) {
	for _, data := range t.data {
		data.Swap(i, j)
	}
}

func (t *Table) UpdateSortOrder(column int) {
	t.sortOrders = append(t.sortOrders, []int{column})
}

func (t *Table) getData(index int) []int {
	data := make([]int, len(t.data))
	for i, d := range t.data {
		data[i] = d.(sort.IntSlice)[index]
	}
	return data
}

// Example usage
func main() {
	data := [][]int{
		{1, 5, 3},
		{2, 4, 1},
		{3, 3, 2},
		{4, 2, 4},
		{5, 1, 5},
	}

	var intSlices []sort.Interface
	for _, row := range data {
		intSlices = append(intSlices, sort.IntSlice(row))
	}

	t := &Table{
		data:       intSlices,
		sortOrders: [][]int{},
	}

	// Initial sorting
	sort.Sort(t)

	// Clicked column headers
	t.UpdateSortOrder(2)
	t.UpdateSortOrder(1)

	// Perform sorting
	sort.Stable(t)

	// Print sorted data
	for _, row := range data {
		fmt.Println(row)
	}
}
