package set

import (
	"bytes"
	"fmt"
)

// IntSet a set of small non-negative integers
type IntSet struct {
	words []uint64
}

// Add add to set
func (s *IntSet) Add(n int) {
	word, bit := n/64, uint(n%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

// Remove remove the element in the set
func (s *IntSet) Remove(n int) {
	if s.Has(n) {
		word, bit := n/64, uint(n%64)
		s.words[word] ^= (1 << bit)
	}
}

// Has integer is in the set
func (s *IntSet) Has(n int) bool {
	word, bits := n/64, uint(n%64)
	return (word < len(s.words)) && (s.words[word]&(1<<bits) != 0)
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

// String returns set as a string of the form "{1 2 3}"
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

// Len return set length
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		for word != 0 {
			count += int((word & 1))
			word = (word >> 1)
		}
	}
	return count
}

// Clear clear set
func (s *IntSet) Clear()  {
    s.words = nil
}

//AddAll add all the elements in the parameter
func (s *IntSet) AddAll(vals ...int)  {
    for _, val := range vals {
        s.Add(val)
    }
}
