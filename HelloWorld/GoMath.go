package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Printf("Now you have %g problems.\n", rand.Intn(30))

	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

}
