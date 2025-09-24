// hashmap
package main

import "fmt"

func main() {
	// 1. Create a map using make()
	// This map stores string keys and int values.
	userAges := make(map[string]int)

	// 2. Add key-value pairs
	userAges["Alice"] = 30
	userAges["Bob"] = 25
	userAges["Charlie"] = 35

	fmt.Println("Map contents:", userAges) // map[Alice:30 Bob:25 Charlie:35]

	// 3. Retrieve a value
	bobsAge := userAges["Bob"]
	fmt.Println("Bob's age is:", bobsAge) // Bob's age is: 25

	// 4. Check if a key exists (the "comma ok" idiom)
	// This is the best way to check for a key.
	age, ok := userAges["David"]
	if !ok {
		fmt.Println("David's age not found!")
	} else {
		fmt.Println("David's age is:", age)
	}

	// 5. Delete a key-value pair
	delete(userAges, "Charlie")
	fmt.Println("Map after deleting Charlie:", userAges) // map[Alice:30 Bob:25]

	// 6. Iterate over a map (order is not guaranteed!)
	fmt.Println("\nIterating over users:")
	for name, currentAge := range userAges {
		fmt.Printf("- %s is %d years old\n", name, currentAge)
	}
}
