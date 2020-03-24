package main

// https://golangbyexample.com/go-iterator-design-pattern/
// https://groups.google.com/forum/#!topic/golang-nuts/v6m86sTRbqA
//https://sunkware.com/blog/2019-06-05-implementing-an-iterator-in-go/

// Iterable represents any structure for which
// it is possible to obtain an iterator
type Iterable interface {
	Iterator() Iterator
}

// Iterator represents any structure that can be traversed
// following the Iterator design pattern
type Iterator interface {
	HasNext() bool
	Next() int
}

// IntSlice represents a slice of int (alias of []int)
type IntSlice []int

// IntSliceIterator represents an iterator on a slice of int
type IntSliceIterator struct {
}

func (s *IntSlice) Iterator() Iterator {
	return &IntSliceIterator{}
}

func (it *IntSliceIterator) HasNext() bool {
	return false
}

func (it *IntSliceIterator) Next() int {
	return 0
}

func foo(data Iterable) {
}

func main() {
	var data Iterable
	data = &IntSlice{1, 2, 3}
	foo(data)
}