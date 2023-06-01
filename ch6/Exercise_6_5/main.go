//Exercise 6.5:
//The type of each word used by IntSet is uint64,
//but 64-bit arithmetic may
//be inefficient on a 32-bit platform.  Modify the program to use the
//uint type, which is the most efficient unsigned integer type for
//the platform.
//
//Instead of dividing by 64, define a constant holding the effective
//size of uint in bits, 32 or 64.
//
//You can use the perhaps too-clever expression
//32 << (^uint(0) >> 63) for this purpose.

package main

import (
	"bytes"
	"fmt"
)

const (
	// WordSize represents the effective size of uint in bits
	WordSize = 32 << (^uint(0) >> 63)
)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/WordSize, uint(x%WordSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/WordSize, uint(x%WordSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(values ...int) {
	for _, value := range values {
		s.Add(value)
	}
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

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popCount(word)
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/WordSize, uint(x%WordSize)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	newSet := &IntSet{}
	newSet.words = make([]uint, len(s.words))
	copy(newSet.words, s.words)
	return newSet
}

func (s *IntSet) Elems() []int {
	elems := make([]int, 0, s.Len())
	for i, word := range s.words {
		for j := 0; j < WordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, WordSize*i+j)
			}
		}
	}
	return elems
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", WordSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func popCount(x uint) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

func main() {
	var x IntSet
	x.AddAll(1, 2, 3)

	elems := x.Elems()
	fmt.Println("Elements:", elems) // [1 2 3]
}
