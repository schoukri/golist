// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2020-04-08 01:37:45.518152 +0000 UTC.

package golist

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// SliceUint64 is a slice of type uint64.
type SliceUint64 struct {
	data []uint64
}

// NewSliceUint64 returns a pointer to a new SliceUint64 initialized with the specified elements.
func NewSliceUint64(elems ...uint64) *SliceUint64 {
	s := new(SliceUint64)
	s.data = make([]uint64, len(elems))
	for i := 0; i < len(elems); i++ {
		s.data[i] = elems[i]
	}
	return s
}

// Append adds the elements to the end of SliceUint64.
func (s *SliceUint64) Append(elems ...uint64) *SliceUint64 {
	if s == nil {
		return nil
	}
	s.data = append(s.data, elems...)
	return s
}

// Prepend adds the elements to the beginning of SliceUint64.
func (s *SliceUint64) Prepend(elems ...uint64) *SliceUint64 {
	if s == nil {
		return nil
	}
	s.data = append(elems, s.data...)
	return s
}

// At returns the element in SliceUint64 at the specified index.
func (s *SliceUint64) At(index int) uint64 {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceUint64 does not contain any elements")
	}

	if index >= len(s.data) || index < 0 {
		panic(fmt.Sprintf("index %d outside the range of SliceUint64", index))
	}

	return s.data[index]
}

// Set sets the element of SliceUint64 at the specified index.
func (s *SliceUint64) Set(index int, elem uint64) *SliceUint64 {
	if s == nil {
		return nil
	}
	s.data[index] = elem
	return s
}

// Insert inserts the elements into SliceUint64 at the specified index.
func (s *SliceUint64) Insert(index int, elems ...uint64) *SliceUint64 {
	if s == nil {
		return nil
	}

	// Grow the slice by the number of elements (using the zero value)
	var zero uint64
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

// Remove removes the element from SliceUint64 at the specified index.
func (s *SliceUint64) Remove(index int) *SliceUint64 {
	if s == nil {
		return nil
	}
	s.data = append(s.data[:index], s.data[index+1:]...)
	return s
}

// Filter removes elements from SliceUint64 that do not satisfy the filter function.
func (s *SliceUint64) Filter(fn func(elem uint64) bool) *SliceUint64 {
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

// Transform modifies each element of SliceUint64 according to the specified function.
func (s *SliceUint64) Transform(fn func(elem uint64) uint64) *SliceUint64 {
	if s == nil {
		return nil
	}
	for i, elem := range s.data {
		s.data[i] = fn(elem)
	}
	return s
}

// Unique modifies SliceUint64 to keep only the first occurrence of each element (removing any duplicates).
func (s *SliceUint64) Unique() *SliceUint64 {
	if s == nil {
		return nil
	}
	seen := make(map[uint64]struct{})
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

// Reverse reverses the order of the elements of SliceUint64.
func (s *SliceUint64) Reverse() *SliceUint64 {
	if s == nil {
		return nil
	}

	for i := len(s.data)/2 - 1; i >= 0; i-- {
		opp := len(s.data) - 1 - i
		s.Swap(i, opp)
	}

	return s
}

// Shuffle randomly shuffles the order of the elements in SliceUint64.
func (s *SliceUint64) Shuffle(seed int64) *SliceUint64 {
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

// Data returns the raw elements of SliceUint64.
func (s *SliceUint64) Data() []uint64 {
	if s == nil {
		return nil
	}
	return s.data
}

// Count returns the number of elements in SliceUint64.
func (s *SliceUint64) Count() int {
	return len(s.data)
}

// Len returns the number of elements in SliceUint64 (alias for Count).
func (s *SliceUint64) Len() int {
	return s.Count()
}

// Swap swaps the elements in SliceUint64 specified by the indices i and j.
func (s *SliceUint64) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

// Less returns true if the SliceUint64 element at index i is less than the element at index j.
func (s *SliceUint64) Less(i, j int) bool {
	return s.data[i] < s.data[j]
}

// Sort sorts the elements of SliceUint64 in increasing order.
func (s *SliceUint64) Sort() *SliceUint64 {
	if s == nil {
		return nil
	}
	sort.Sort(s)
	return s
}

// Min returns the smallest (least ordered) element in SliceUint64.
func (s *SliceUint64) Min() uint64 {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceUint64 does not contain any elements")
	}
	// start with the first value
	min := s.data[0]
	for _, elem := range s.data[1:] {
		if elem < min {
			min = elem
		}
	}
	return min
}

// Max returns the largest (greatest ordered) element in SliceUint64.
func (s *SliceUint64) Max() uint64 {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceUint64 does not contain any elements")
	}
	// start with the first value
	max := s.data[0]
	for _, elem := range s.data[1:] {
		if elem > max {
			max = elem
		}
	}
	return max
}

// Clone performs a deep copy of SliceUint64 and returns it
func (s *SliceUint64) Clone() *SliceUint64 {
	if s == nil {
		return nil
	}
	s2 := new(SliceUint64)
	s2.data = make([]uint64, len(s.data))
	copy(s2.data, s.data)
	return s2
}

// Equal returns true if the SliceUint64 is logically equivalent to the specified SliceUint64.
func (s *SliceUint64) Equal(s2 *SliceUint64) bool {
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
