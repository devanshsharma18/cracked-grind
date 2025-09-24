//hashsets

package main

import "fmt"

func main() {
	// Create a "set" to store unique fruit names
	fruitSet := make(map[string]struct{})

	// Add items to the set.
	// The empty struct{} is just a placeholder value.
	fruitSet["apple"] = struct{}{}
	fruitSet["banana"] = struct{}{}
	fruitSet["orange"] = struct{}{}

	// Try adding a duplicate. It won't cause an error;
	// it simply overwrites the existing entry with the same value.
	// The set of keys remains unique.
	fruitSet["apple"] = struct{}{}

	fmt.Println("Current items in set:", fruitSet) // map[apple:{} banana:{} orange:{}]

	// Check for existence in the set
	_, exists := fruitSet["banana"]
	if exists {
		fmt.Println("Yes, 'banana' is in the set.") // This will print
	}

	_, exists = fruitSet["grape"]
	if !exists {
		fmt.Println("No, 'grape' is not in the set.") // This will print
	}

	// Remove an item from the set
	delete(fruitSet, "orange")

	// Print the final unique items
	fmt.Println("\nFinal unique fruits:")
	for fruit := range fruitSet {
		fmt.Println("-", fruit)
	}
}
