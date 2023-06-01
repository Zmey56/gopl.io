//Exercise 6.3: (*IntSet).UnionWith computes the union of two sets using |, the
//word-parallel bitwise OR operator. Implement methods for IntersectWith, DifferenceWith, and
//SymmetricDifference for the corresponding set operations. (The symmetric difference of two sets contains the elements present
//in one set or the other but not both.)”

package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &^= t.words[i]
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popCount(word)
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	newSet := &IntSet{}
	newSet.words = make([]uint64, len(s.words))
	copy(newSet.words, s.words)
	return newSet
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func popCount(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)

	y.Add(2)
	y.Add(3)
	y.Add(4)

	fmt.Println("x:", x.String()) // "{1 2 3}"
	fmt.Println("y:", y.String()) // "{2 3 4}"

	x.IntersectWith(&y)
	fmt.Println("Intersection:", x.String()) // "{2 3}"

	x, y = IntSet{}, IntSet{}
	x.Add(1)
	x.Add(2)
	x.Add(3)

	y.Add(2)
	y.Add(3)
	y.Add(4)

	x.DifferenceWith(&y)
	fmt.Println("Difference:", x.String()) // "{1}"

	x, y = IntSet{}, IntSet{}
	x.Add(1)
	x.Add(2)
	x.Add(3)

	y.Add(2)
	y.Add(3)
	y.Add(4)

	x.SymmetricDifference(&y)
	fmt.Println("Symmetric Difference:", x.String()) // "{1 4}"
}
