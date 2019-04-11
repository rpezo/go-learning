package main

import "fmt"

func main() {
	x := 10
	y := 10

	if x <= y {
		fmt.Printf("%d es menor o igual %d\n", x, y)
	} else {
		fmt.Printf("%d es menor que %d\n", y, x)
	}

	// else if
	color := "negro"

	if color == "rojo" {
		fmt.Println("color es rojo")
	} else if color == "azul" {
		fmt.Println("color es azul")
	} else {
		fmt.Println("color no es Rojo ni Azul")
	}

	// switch
	switch color {
	case "rojo":
		fmt.Println("color es rojo")
	case "azul":
		fmt.Println("color es azul")
	default:
		fmt.Println("color no es rojo ni azul")

	}

}
