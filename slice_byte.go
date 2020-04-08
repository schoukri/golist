// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2020-04-08 00:13:31.066195 +0000 UTC.

package golist

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// SliceByte is a slice of type byte.
type SliceByte struct {
	data []byte
}

// NewSliceByte returns a pointer to a new SliceByte initialized with the specified elements.
func NewSliceByte(elems ...byte) *SliceByte {
	s := new(SliceByte)
	s.data = make([]byte, len(elems))
	for i := 0; i < len(elems); i++ {
		s.data[i] = elems[i]
	}
	return s
}

// Append adds the elements to the end of SliceByte.
func (s *SliceByte) Append(elems ...byte) *SliceByte {
	if s == nil {
		return nil
	}
	s.data = append(s.data, elems...)
	return s
}

// Prepend adds the elements to the beginning of SliceByte.
func (s *SliceByte) Prepend(elems ...byte) *SliceByte {
	if s == nil {
		return nil
	}
	s.data = append(elems, s.data...)
	return s
}

// At returns the element in SliceByte at the specified index.
func (s *SliceByte) At(index int) byte {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceByte does not contain any elements")
	}

	if index >= len(s.data) || index < 0 {
		panic(fmt.Sprintf("index %d outside the range of SliceByte", index))
	}

	return s.data[index]
}

// Set sets the element of SliceByte at the specified index.
func (s *SliceByte) Set(index int, elem byte) *SliceByte {
	if s == nil {
		return nil
	}
	s.data[index] = elem
	return s
}

// Insert inserts the elements into SliceByte at the specified index.
func (s *SliceByte) Insert(index int, elems ...byte) *SliceByte {
	if s == nil {
		return nil
	}

	// Grow the slice by the number of elements (using the zero value)
	var zero byte
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

// Remove removes the element from SliceByte at the specified index.
func (s *SliceByte) Remove(index int) *SliceByte {
	if s == nil {
		return nil
	}
	s.data = append(s.data[:index], s.data[index+1:]...)
	return s
}

// Filter removes elements from SliceByte that do not satisfy the filter function.
func (s *SliceByte) Filter(fn func(elem byte) bool) *SliceByte {
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

// Transform modifies each element of SliceByte according to the specified function.
func (s *SliceByte) Transform(fn func(elem byte) byte) *SliceByte {
	if s == nil {
		return nil
	}
	for i, elem := range s.data {
		s.data[i] = fn(elem)
	}
	return s
}

// Unique modifies SliceByte to keep only the first occurrence of each element (removing any duplicates).
func (s *SliceByte) Unique() *SliceByte {
	if s == nil {
		return nil
	}
	seen := make(map[byte]struct{})
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

// Reverse reverses the order of the elements of SliceByte.
func (s *SliceByte) Reverse() *SliceByte {
	if s == nil {
		return nil
	}

	for i := len(s.data)/2 - 1; i >= 0; i-- {
		opp := len(s.data) - 1 - i
		s.Swap(i, opp)
	}

	return s
}

// Shuffle randomly shuffles the order of the elements in SliceByte.
func (s *SliceByte) Shuffle(seed int64) *SliceByte {
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

// Data returns the raw elements of SliceByte.
func (s *SliceByte) Data() []byte {
	if s == nil {
		return nil
	}
	return s.data
}

// Count returns the number of elements in SliceByte.
func (s *SliceByte) Count() int {
	return len(s.data)
}

// Len returns the number of elements in SliceByte (alias for Count).
func (s *SliceByte) Len() int {
	return s.Count()
}

// Swap swaps the elements in SliceByte specified by the indices i and j.
func (s *SliceByte) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}
// Less returns true if the SliceByte element at index i is less than the element at index j.
func (s *SliceByte) Less(i, j int) bool {
	return s.data[i] < s.data[j]
}

// Sort sorts the elements of SliceByte in increasing order.
func (s *SliceByte) Sort() *SliceByte {
	if s == nil {
		return nil
	}
	sort.Sort(s)
	return s
}

// Min returns the smallest (least ordered) element in SliceByte.
func (s *SliceByte) Min() byte {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceByte does not contain any elements")
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

// Max returns the largest (greatest ordered) element in SliceByte.
func (s *SliceByte) Max() byte {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceByte does not contain any elements")
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

// Clone performs a deep copy of SliceByte and returns it
func (s *SliceByte) Clone() *SliceByte {
	if s == nil {
		return nil
	}
	s2 := new(SliceByte)
	s2.data = make([]byte, len(s.data))
	copy(s2.data, s.data)
	return s2
}

// Equal returns true if the SliceByte is logically equivalent to the specified SliceByte.
func (s *SliceByte) Equal(s2 *SliceByte) bool {
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