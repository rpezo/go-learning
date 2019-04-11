package main

import "fmt"

func main() {
	//Arrays
	var fruitArr [3]string

	//Assign values
	fruitArr[0] = "Manzana"
	fruitArr[1] = "Naranja"
	fruitArr[2] = "Platano"
	fmt.Println(fruitArr)
	fmt.Println(fruitArr[2])

	fruitSlice := []string{"Apple", "Orange", "Banana"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:3])
}
