// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2020-04-08 00:13:31.723233 +0000 UTC.

package golist

import (
	"fmt"
	"math/rand"
	"time"
)

// SliceComplex64 is a slice of type complex64.
type SliceComplex64 struct {
	data []complex64
}

// NewSliceComplex64 returns a pointer to a new SliceComplex64 initialized with the specified elements.
func NewSliceComplex64(elems ...complex64) *SliceComplex64 {
	s := new(SliceComplex64)
	s.data = make([]complex64, len(elems))
	for i := 0; i < len(elems); i++ {
		s.data[i] = elems[i]
	}
	return s
}

// Append adds the elements to the end of SliceComplex64.
func (s *SliceComplex64) Append(elems ...complex64) *SliceComplex64 {
	if s == nil {
		return nil
	}
	s.data = append(s.data, elems...)
	return s
}

// Prepend adds the elements to the beginning of SliceComplex64.
func (s *SliceComplex64) Prepend(elems ...complex64) *SliceComplex64 {
	if s == nil {
		return nil
	}
	s.data = append(elems, s.data...)
	return s
}

// At returns the element in SliceComplex64 at the specified index.
func (s *SliceComplex64) At(index int) complex64 {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceComplex64 does not contain any elements")
	}

	if index >= len(s.data) || index < 0 {
		panic(fmt.Sprintf("index %d outside the range of SliceComplex64", index))
	}

	return s.data[index]
}

// Set sets the element of SliceComplex64 at the specified index.
func (s *SliceComplex64) Set(index int, elem complex64) *SliceComplex64 {
	if s == nil {
		return nil
	}
	s.data[index] = elem
	return s
}

// Insert inserts the elements into SliceComplex64 at the specified index.
func (s *SliceComplex64) Insert(index int, elems ...complex64) *SliceComplex64 {
	if s == nil {
		return nil
	}

	// Grow the slice by the number of elements (using the zero value)
	var zero complex64
	for i := 0; i < len(elems); i++ {
		s.data = append(s.data, zero)
	}

	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(s.data[index+len(elems):], s.data[index:])

	// Store the new values
	for i := 0; i < len(elems); i++ {
		s.data[index+i] = elems[i]
	}

	// Return the result.
	return s
}

// Remove removes the element from SliceComplex64 at the specified index.
func (s *SliceComplex64) Remove(index int) *SliceComplex64 {
	if s == nil {
		return nil
	}
	s.data = append(s.data[:index], s.data[index+1:]...)
	return s
}

// Filter removes elements from SliceComplex64 that do not satisfy the filter function.
func (s *SliceComplex64) Filter(fn func(elem complex64) bool) *SliceComplex64 {
	if s == nil {
		return nil
	}
	data := s.data[:0]
	for _, elem := range s.data {
		if fn(elem) {
			data = append(data, elem)
		}
	}
	s.data = data
	return s
}

// Transform modifies each element of SliceComplex64 according to the specified function.
func (s *SliceComplex64) Transform(fn func(elem complex64) complex64) *SliceComplex64 {
	if s == nil {
		return nil
	}
	for i, elem := range s.data {
		s.data[i] = fn(elem)
	}
	return s
}

// Unique modifies SliceComplex64 to keep only the first occurrence of each element (removing any duplicates).
func (s *SliceComplex64) Unique() *SliceComplex64 {
	if s == nil {
		return nil
	}
	seen := make(map[complex64]struct{})
	data := s.data[:0]
	for _, elem := range s.data {
		if _, ok := seen[elem]; !ok {
			data = append(data, elem)
			seen[elem] = struct{}{}
		}
	}
	s.data = data
	return s
}

// Reverse reverses the order of the elements of SliceComplex64.
func (s *SliceComplex64) Reverse() *SliceComplex64 {
	if s == nil {
		return nil
	}

	for i := len(s.data)/2 - 1; i >= 0; i-- {
		opp := len(s.data) - 1 - i
		s.Swap(i, opp)
	}

	return s
}

// Shuffle randomly shuffles the order of the elements in SliceComplex64.
func (s *SliceComplex64) Shuffle(seed int64) *SliceComplex64 {
	if s == nil {
		return nil
	}

	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	r := rand.New(rand.NewSource(seed))
	r.Shuffle(s.Count(), s.Swap)

	return s
}

// Data returns the raw elements of SliceComplex64.
func (s *SliceComplex64) Data() []complex64 {
	if s == nil {
		return nil
	}
	return s.data
}

// Count returns the number of elements in SliceComplex64.
func (s *SliceComplex64) Count() int {
	return len(s.data)
}

// Len returns the number of elements in SliceComplex64 (alias for Count).
func (s *SliceComplex64) Len() int {
	return s.Count()
}

// Swap swaps the elements in SliceComplex64 specified by the indices i and j.
func (s *SliceComplex64) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

// Clone performs a deep copy of SliceComplex64 and returns it
func (s *SliceComplex64) Clone() *SliceComplex64 {
	if s == nil {
		return nil
	}
	s2 := new(SliceComplex64)
	s2.data = make([]complex64, len(s.data))
	copy(s2.data, s.data)
	return s2
}

// Equal returns true if the SliceComplex64 is logically equivalent to the specified SliceComplex64.
func (s *SliceComplex64) Equal(s2 *SliceComplex64) bool {
	if s == s2 {
		return true
	}

	if s == nil || s2 == nil {
		return false // has to be false because s == s2 tested earlier
	}

	if len(s.data) != len(s2.data) {
		return false
	}

	for i, elem := range s.data {
		if elem != s2.data[i] {
			return false
		}
	}

	return true
}
