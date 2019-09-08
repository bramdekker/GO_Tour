package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	complexNum  complex64 = 8 + 10i
	unsignedNum uint8     = 1<<8 - 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(20))
	fmt.Printf("Math package exported variable is Pi not pi: %g.\n", math.Pi)
	var age = 18
	fmt.Println("A legal age in Holland is", age)
	fmt.Println("Nice fact: one can use := declaration only inside a function!")
	fmt.Printf("Complex numbers in Go of type %T and value: %v!\n", complexNum, complexNum)
	fmt.Printf("Also the basic variable types like uint: %v.\n", unsignedNum)
	fmt.Print("Constants in Go have the const keyword and cant be declared with := signs.")
}
