package main

import (
	"fmt"
	"math"

	"github.com/rpezo/go-learning/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(3.9))
	fmt.Println(math.Ceil(3.9))
	fmt.Println(strutil.Reverse("aloH"))
}
