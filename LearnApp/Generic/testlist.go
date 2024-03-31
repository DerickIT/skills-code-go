package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Push adds a new element with the given value to the front of the list.
func (l *List[T]) Push(val T) {
	newNode := &List[T]{val: val}
	newNode.next = l.next
	l.next = newNode
}

// Pop removes and returns the first element of the list.
// It returns the zero value of type T if the list is empty.
func (l *List[T]) Pop() T {
	if l.next == nil {
		var zero T
		return zero
	}
	val := l.next.val
	l.next = l.next.next
	return val
}

// Peek returns the value of the first element of the list without removing it.
// It returns the zero value of type T if the list is empty.
func (l *List[T]) Peek() T {
	if l.next == nil {
		var zero T
		return zero
	}
	return l.next.val
}

// IsEmpty returns true if the list is empty, false otherwise.
func (l *List[T]) IsEmpty() bool {
	return l.next == nil
}

// Len returns the number of elements in the list.
func (l *List[T]) Len() int {
	count := 0
	for current := l.next; current != nil; current = current.next {
		count++
	}
	return count
}

// Print prints the values of all elements in the list.
func (l *List[T]) Print() {
	for current := l.next; current != nil; current = current.next {
		fmt.Print(current.val, " ")
	}
	fmt.Println()
}

func main() {
	var list List[int]

	list.Push(3)
	list.Push(2)
	list.Push(1)

	fmt.Println("List elements:")
	list.Print()

	fmt.Println("List length:", list.Len())
	fmt.Println("Is list empty?", list.IsEmpty())

	fmt.Println("First element:", list.Peek())

	fmt.Println("Pop elements:")
	for !list.IsEmpty() {
		fmt.Print(list.Pop(), " ")
	}
	fmt.Println()

	fmt.Println("Is list empty?", list.IsEmpty())
}
