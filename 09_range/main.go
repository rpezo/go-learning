package main

import "fmt"

func main() {
	ids := []int{10, 20, 30, 40, 50, 60}

	// Loop through PrintDefaults
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Rod": "rod@gmail.com", "Capito": "capito@gmai.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
