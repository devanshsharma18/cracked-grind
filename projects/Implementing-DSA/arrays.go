package main

import "fmt"

// MyArray is our simple, fixed-size array structure.
// It holds a slice of a generic type T.
type MyArray[T any] struct {
	data []T
}

// NewArray is the constructor. It creates our array with a fixed size.
func NewArray[T any](size int) *MyArray[T] {
	// We create a slice with both its length and capacity set to the desired size.
	// This means all the slots are created and ready to be used.
	return &MyArray[T]{
		data: make([]T, size),
	}
}

// Set places a value at a specific index.
func (a *MyArray[T]) Set(index int, value T) error {
	// Bounds Check: We must ensure the index is within the valid range.
	if index < 0 || index >= len(a.data) {
		return fmt.Errorf("index %d is out of bounds for array with size %d", index, len(a.data))
	}

	a.data[index] = value
	return nil
}

// Get retrieves a value from a specific index.
func (a *MyArray[T]) Get(index int) (T, error) {
	// Bounds Check: Same as in Set.
	if index < 0 || index >= len(a.data) {
		var zero T // The default "zero" value for the type T (e.g., 0, "", nil)
		return zero, fmt.Errorf("index %d is out of bounds for array with size %d", index, len(a.data))
	}

	return a.data[index], nil
}

// Size returns the total size (capacity) of the array.
func (a *MyArray[T]) Size() int {
	return len(a.data)
}

func main() {
	// Create a new array of strings with a fixed size of 3.
	fmt.Println("Creating an array of size 3...")
	names := NewArray[string](3)
	fmt.Printf("Array created. Size: %d\n\n", names.Size())

	// Set some values.
	fmt.Println("Setting values...")
	names.Set(0, "Alice")
	names.Set(1, "Bob")
	names.Set(2, "Charlie")

	// Get a value and print it.
	fmt.Println("Getting value at index 1...")
	name, err := names.Get(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value at index 1 is:", name)
	}
	fmt.Println("Current data in array:", names.data)
	fmt.Println()

	// --- Test Error Handling ---
	fmt.Println("--- Testing out of bounds access ---")

	// Try to get a value from an index that doesn't exist.
	_, err = names.Get(3)
	if err != nil {
		fmt.Println("SUCCESS! Got expected error:", err)
	}

	// Try to set a value at an index that doesn't exist.
	err = names.Set(-1, "David")
	if err != nil {
		fmt.Println("SUCCESS! Got expected error:", err)
	}
}
