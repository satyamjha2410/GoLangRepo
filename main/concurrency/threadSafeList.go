package main

import (
	"fmt"
	"sync"
)

// SafeList is a thread-safe list implementation
type SafeList struct {
	sync.Mutex
	data []int
}

// Append adds an element to the thread-safe list
func (s *SafeList) Append(value int) {
	s.Lock()
	defer s.Unlock()
	s.data = append(s.data, value)
}

// Print prints the elements of the thread-safe list
func (s *SafeList) Print() {
	s.Lock()
	defer s.Unlock()
	fmt.Println("Thread-safe list:", s.data)
}

func main() {
	// Create a thread-safe list
	threadSafeList := &SafeList{}

	// Add elements concurrently
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			threadSafeList.Append(index)
		}(i)
	}
	wg.Wait()

	// Print the thread-safe list
	threadSafeList.Print()
}
