// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2020-04-08 01:37:41.559513 +0000 UTC.

package golist

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// SliceFloat32 is a slice of type float32.
type SliceFloat32 struct {
	data []float32
}

// NewSliceFloat32 returns a pointer to a new SliceFloat32 initialized with the specified elements.
func NewSliceFloat32(elems ...float32) *SliceFloat32 {
	s := new(SliceFloat32)
	s.data = make([]float32, len(elems))
	for i := 0; i < len(elems); i++ {
		s.data[i] = elems[i]
	}
	return s
}

// Append adds the elements to the end of SliceFloat32.
func (s *SliceFloat32) Append(elems ...float32) *SliceFloat32 {
	if s == nil {
		return nil
	}
	s.data = append(s.data, elems...)
	return s
}

// Prepend adds the elements to the beginning of SliceFloat32.
func (s *SliceFloat32) Prepend(elems ...float32) *SliceFloat32 {
	if s == nil {
		return nil
	}
	s.data = append(elems, s.data...)
	return s
}

// At returns the element in SliceFloat32 at the specified index.
func (s *SliceFloat32) At(index int) float32 {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceFloat32 does not contain any elements")
	}

	if index >= len(s.data) || index < 0 {
		panic(fmt.Sprintf("index %d outside the range of SliceFloat32", index))
	}

	return s.data[index]
}

// Set sets the element of SliceFloat32 at the specified index.
func (s *SliceFloat32) Set(index int, elem float32) *SliceFloat32 {
	if s == nil {
		return nil
	}
	s.data[index] = elem
	return s
}

// Insert inserts the elements into SliceFloat32 at the specified index.
func (s *SliceFloat32) Insert(index int, elems ...float32) *SliceFloat32 {
	if s == nil {
		return nil
	}

	// Grow the slice by the number of elements (using the zero value)
	var zero float32
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

// Remove removes the element from SliceFloat32 at the specified index.
func (s *SliceFloat32) Remove(index int) *SliceFloat32 {
	if s == nil {
		return nil
	}
	s.data = append(s.data[:index], s.data[index+1:]...)
	return s
}

// Filter removes elements from SliceFloat32 that do not satisfy the filter function.
func (s *SliceFloat32) Filter(fn func(elem float32) bool) *SliceFloat32 {
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

// Transform modifies each element of SliceFloat32 according to the specified function.
func (s *SliceFloat32) Transform(fn func(elem float32) float32) *SliceFloat32 {
	if s == nil {
		return nil
	}
	for i, elem := range s.data {
		s.data[i] = fn(elem)
	}
	return s
}

// Unique modifies SliceFloat32 to keep only the first occurrence of each element (removing any duplicates).
func (s *SliceFloat32) Unique() *SliceFloat32 {
	if s == nil {
		return nil
	}
	seen := make(map[float32]struct{})
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

// Reverse reverses the order of the elements of SliceFloat32.
func (s *SliceFloat32) Reverse() *SliceFloat32 {
	if s == nil {
		return nil
	}

	for i := len(s.data)/2 - 1; i >= 0; i-- {
		opp := len(s.data) - 1 - i
		s.Swap(i, opp)
	}

	return s
}

// Shuffle randomly shuffles the order of the elements in SliceFloat32.
func (s *SliceFloat32) Shuffle(seed int64) *SliceFloat32 {
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

// Data returns the raw elements of SliceFloat32.
func (s *SliceFloat32) Data() []float32 {
	if s == nil {
		return nil
	}
	return s.data
}

// Count returns the number of elements in SliceFloat32.
func (s *SliceFloat32) Count() int {
	return len(s.data)
}

// Len returns the number of elements in SliceFloat32 (alias for Count).
func (s *SliceFloat32) Len() int {
	return s.Count()
}

// Swap swaps the elements in SliceFloat32 specified by the indices i and j.
func (s *SliceFloat32) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

// Less returns true if the SliceFloat32 element at index i is less than the element at index j.
func (s *SliceFloat32) Less(i, j int) bool {
	return s.data[i] < s.data[j]
}

// Sort sorts the elements of SliceFloat32 in increasing order.
func (s *SliceFloat32) Sort() *SliceFloat32 {
	if s == nil {
		return nil
	}
	sort.Sort(s)
	return s
}

// Min returns the smallest (least ordered) element in SliceFloat32.
func (s *SliceFloat32) Min() float32 {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceFloat32 does not contain any elements")
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

// Max returns the largest (greatest ordered) element in SliceFloat32.
func (s *SliceFloat32) Max() float32 {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceFloat32 does not contain any elements")
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

// Clone performs a deep copy of SliceFloat32 and returns it
func (s *SliceFloat32) Clone() *SliceFloat32 {
	if s == nil {
		return nil
	}
	s2 := new(SliceFloat32)
	s2.data = make([]float32, len(s.data))
	copy(s2.data, s.data)
	return s2
}

// Equal returns true if the SliceFloat32 is logically equivalent to the specified SliceFloat32.
func (s *SliceFloat32) Equal(s2 *SliceFloat32) bool {
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
