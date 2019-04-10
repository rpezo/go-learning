package main

import "fmt"

func main() {
	// Main Types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte -alias for uint32
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	//Using var
	//	var name string = "Rod"
	var age int32 = 39
	var isCool bool = true
	var size float32 = 2.3
	isCool = false
	//Shorthand
	name := "Rodrigo"
	email := "rod@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
}
