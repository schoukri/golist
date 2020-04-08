// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2020-04-08 01:37:44.184766 +0000 UTC.

package golist

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// SliceString is a slice of type string.
type SliceString struct {
	data []string
}

// NewSliceString returns a pointer to a new SliceString initialized with the specified elements.
func NewSliceString(elems ...string) *SliceString {
	s := new(SliceString)
	s.data = make([]string, len(elems))
	for i := 0; i < len(elems); i++ {
		s.data[i] = elems[i]
	}
	return s
}

// Append adds the elements to the end of SliceString.
func (s *SliceString) Append(elems ...string) *SliceString {
	if s == nil {
		return nil
	}
	s.data = append(s.data, elems...)
	return s
}

// Prepend adds the elements to the beginning of SliceString.
func (s *SliceString) Prepend(elems ...string) *SliceString {
	if s == nil {
		return nil
	}
	s.data = append(elems, s.data...)
	return s
}

// At returns the element in SliceString at the specified index.
func (s *SliceString) At(index int) string {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceString does not contain any elements")
	}

	if index >= len(s.data) || index < 0 {
		panic(fmt.Sprintf("index %d outside the range of SliceString", index))
	}

	return s.data[index]
}

// Set sets the element of SliceString at the specified index.
func (s *SliceString) Set(index int, elem string) *SliceString {
	if s == nil {
		return nil
	}
	s.data[index] = elem
	return s
}

// Insert inserts the elements into SliceString at the specified index.
func (s *SliceString) Insert(index int, elems ...string) *SliceString {
	if s == nil {
		return nil
	}

	// Grow the slice by the number of elements (using the zero value)
	var zero string
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

// Remove removes the element from SliceString at the specified index.
func (s *SliceString) Remove(index int) *SliceString {
	if s == nil {
		return nil
	}
	s.data = append(s.data[:index], s.data[index+1:]...)
	return s
}

// Filter removes elements from SliceString that do not satisfy the filter function.
func (s *SliceString) Filter(fn func(elem string) bool) *SliceString {
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

// Transform modifies each element of SliceString according to the specified function.
func (s *SliceString) Transform(fn func(elem string) string) *SliceString {
	if s == nil {
		return nil
	}
	for i, elem := range s.data {
		s.data[i] = fn(elem)
	}
	return s
}

// Unique modifies SliceString to keep only the first occurrence of each element (removing any duplicates).
func (s *SliceString) Unique() *SliceString {
	if s == nil {
		return nil
	}
	seen := make(map[string]struct{})
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

// Reverse reverses the order of the elements of SliceString.
func (s *SliceString) Reverse() *SliceString {
	if s == nil {
		return nil
	}

	for i := len(s.data)/2 - 1; i >= 0; i-- {
		opp := len(s.data) - 1 - i
		s.Swap(i, opp)
	}

	return s
}

// Shuffle randomly shuffles the order of the elements in SliceString.
func (s *SliceString) Shuffle(seed int64) *SliceString {
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

// Data returns the raw elements of SliceString.
func (s *SliceString) Data() []string {
	if s == nil {
		return nil
	}
	return s.data
}

// Count returns the number of elements in SliceString.
func (s *SliceString) Count() int {
	return len(s.data)
}

// Len returns the number of elements in SliceString (alias for Count).
func (s *SliceString) Len() int {
	return s.Count()
}

// Swap swaps the elements in SliceString specified by the indices i and j.
func (s *SliceString) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

// Less returns true if the SliceString element at index i is less than the element at index j.
func (s *SliceString) Less(i, j int) bool {
	return s.data[i] < s.data[j]
}

// Sort sorts the elements of SliceString in increasing order.
func (s *SliceString) Sort() *SliceString {
	if s == nil {
		return nil
	}
	sort.Sort(s)
	return s
}

// Min returns the smallest (least ordered) element in SliceString.
func (s *SliceString) Min() string {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceString does not contain any elements")
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

// Max returns the largest (greatest ordered) element in SliceString.
func (s *SliceString) Max() string {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceString does not contain any elements")
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

// Clone performs a deep copy of SliceString and returns it
func (s *SliceString) Clone() *SliceString {
	if s == nil {
		return nil
	}
	s2 := new(SliceString)
	s2.data = make([]string, len(s.data))
	copy(s2.data, s.data)
	return s2
}

// Equal returns true if the SliceString is logically equivalent to the specified SliceString.
func (s *SliceString) Equal(s2 *SliceString) bool {
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
