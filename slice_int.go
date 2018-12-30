package golist

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// SliceInt is a slice of integers.
type SliceInt struct {
	data []int
}

// NewSliceInt returns a pointer to a new SliceInt initialized with the specified elements.
func NewSliceInt(elems ...int) *SliceInt {
	s := new(SliceInt)
	s.data = make([]int, len(elems))
	for i := 0; i < len(elems); i++ {
		s.data[i] = elems[i]
	}
	return s
}

// Append adds the elements to the end of SliceInt.
func (s *SliceInt) Append(elems ...int) *SliceInt {
	if s == nil {
		return nil
	}
	s.data = append(s.data, elems...)
	return s
}

// Prepend adds the elements to the beginning of SliceInt.
func (s *SliceInt) Prepend(elems ...int) *SliceInt {
	if s == nil {
		return nil
	}
	s.data = append(elems, s.data...)
	return s
}

// At returns the element in SliceInt at the specified index.
func (s *SliceInt) At(index int) int {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceInt does not contain any elements")
	}

	if index >= len(s.data) || index < 0 {
		panic(fmt.Sprintf("index %d outside the range of SliceInt", index))
	}

	return s.data[index]
}

// Set sets the element of SliceInt at the specified index.
func (s *SliceInt) Set(index, elem int) *SliceInt {
	if s == nil {
		return nil
	}
	s.data[index] = elem
	return s
}

// Insert inserts the elements into SliceInt at the specified index.
func (s *SliceInt) Insert(index int, elems ...int) *SliceInt {
	if s == nil {
		return nil
	}

	// Grow the slice by the number of elements (using the zero value)
	var zero int
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

// Filter removes elements from SliceInt that do not satisfy the filter function.
func (s *SliceInt) Filter(fn func(elem int) bool) *SliceInt {
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

// Transform modifies each element of SliceInt according to the specified function.
func (s *SliceInt) Transform(fn func(elem int) int) *SliceInt {
	if s == nil {
		return nil
	}
	for i, elem := range s.data {
		s.data[i] = fn(elem)
	}
	return s
}

// Unique modifies SliceInt to keep only the first occurrence of each element (removing any duplicates).
func (s *SliceInt) Unique() *SliceInt {
	if s == nil {
		return nil
	}
	seen := make(map[int]struct{})
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

// Sort sorts the elements of SliceInt in increasing order.
func (s *SliceInt) Sort() *SliceInt {
	if s == nil {
		return nil
	}
	sort.Ints(s.data)
	return s
}

// Reverse reverses the order of the elements of SliceInt.
func (s *SliceInt) Reverse() *SliceInt {
	if s == nil {
		return nil
	}

	for i := len(s.data)/2 - 1; i >= 0; i-- {
		opp := len(s.data) - 1 - i
		s.data[i], s.data[opp] = s.data[opp], s.data[i]
	}

	return s
}

// Shuffle randomly shuffles the order of the elements in SliceInt.
func (s *SliceInt) Shuffle(seed int64) *SliceInt {
	if s == nil {
		return nil
	}

	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(s.data), func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	})

	return s
}

// Data returns the raw elements of SliceInt.
func (s *SliceInt) Data() []int {
	if s == nil {
		return nil
	}
	return s.data
}

// Len returns the number of elements in SliceInt.
func (s *SliceInt) Len() int {
	return len(s.data)
}

// Count is an alias for Len()
func (s *SliceInt) Count() int {
	return s.Len()
}

// Min returns the smallest element in SliceInt.
func (s *SliceInt) Min() int {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceInt does not contain any elements")
	}
	// start with the largest int value possible
	min := math.MaxInt32
	for _, elem := range s.data {
		if elem < min {
			min = elem
		}
	}
	return min
}

// Max returns the largest element in SliceInt.
func (s *SliceInt) Max() int {
	if s.data == nil || len(s.data) == 0 {
		panic("SliceInt does not contain any elements")
	}
	// start with the smallest int value possible
	max := math.MinInt32
	for _, elem := range s.data {
		if elem > max {
			max = elem
		}
	}
	return max
}

// Equal returns true if the SliceInt is logically equivalent to the specified SliceInt.
func (s *SliceInt) Equal(x *SliceInt) bool {
	if s == x {
		return true
	}

	if s == nil || x == nil {
		return false // has to be false because s == x tested earlier
	}

	if len(s.data) != len(x.data) {
		return false
	}

	for i, elem := range s.data {
		if elem != x.data[i] {
			return false
		}
	}

	return true
}
