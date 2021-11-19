package main

import (
	"bytes"
	"fmt"
)

const WORD = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/WORD, uint(x%WORD)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) AddAll(bits ...int) {
	for _, bit := range bits {
		s.Add(bit)
	}
}

func (s *IntSet) Add(x int) {
	word, bit := x/WORD, uint(x%WORD)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Remove(x int) {
	word, bit := x/WORD, x%WORD
	if word >= len(s.words) {
		return
	}
	if bit != 0 {
		s.words[word] &= ^(1 << bit)
	}
}

func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

func (s *IntSet) Copy() *IntSet {
	var c *IntSet = new(IntSet)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WORD; j++ {
			if word&(1<<uint(j)) != 0 {
				c.Add(i*WORD + j)
			}
		}
	}

	return c
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
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
	s.words = s.words[:len(t.words)]
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	l := len(s.words)
	if l > len(t.words) {
		l = len(t.words)
	}

	for i := 0; i < l; i++ {
		if s.words[i] == 0 && t.words[i] == 0 {
			continue
		}
		for j := 0; j < WORD; j++ {
			sbit := s.words[i]&(1<<uint(j)) != 0
			tbit := t.words[i]&(1<<uint(j)) != 0
			if sbit && tbit {
				s.Remove(i*WORD + j)
			}
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	l := len(s.words)
	if l > len(t.words) {
		l = len(t.words)
	}

	for i := 0; i < l; i++ {
		if s.words[i] == 0 && t.words[i] == 0 {
			continue
		}
		for j := 0; j < WORD; j++ {
			sbit := s.words[i]&(1<<uint(j)) != 0
			tbit := t.words[i]&(1<<uint(j)) != 0
			if sbit && tbit {
				s.Remove(i*WORD + j)
			} else if tbit {
				s.Add(i*WORD + j)
			}
		}
	}
	if l != len(t.words) {
		for i := l; i < len(t.words); i++ {
			if t.words[i] == 0 {
				continue
			}
			for j := 0; j < WORD; j++ {
				if t.words[i]&(1<<uint(j)) != 0 {
					s.Add(i*WORD + j)
				}
			}
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WORD; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", WORD*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WORD; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) Elems() []uint {
	var t []uint
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WORD; j++ {
			if word&(1<<uint(j)) != 0 {
				t = append(t, uint(i*WORD+j))
			}
		}
	}
	return t
}

func main() {
	var x, y IntSet
	x.AddAll(1, 9, 144)

	y.Add(9)
	y.Add(42)
	fmt.Println(x.String())
	fmt.Println(y.String())

	x.DifferenceWith(&y)
	fmt.Println(x.String())

	t := x.Elems()
	for _, bit := range t {
		fmt.Println(bit)
	}
}
