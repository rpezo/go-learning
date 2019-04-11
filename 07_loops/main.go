package main

import "fmt"

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		//  i = i+1
		i++
	}

	// Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Numero %d\n", i)
	}

	// TangananicaTanganana
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("TangananicaTanganana")
		} else if i%3 == 0 {
			fmt.Println("Tangananica")
		} else if i%5 == 0 {
			fmt.Println("Tanganana")
		} else {
			fmt.Println(i)
		}
	}

}
