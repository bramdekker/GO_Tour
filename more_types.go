package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/wc"
)

// Point represents a 2d point in Euclidian space.
type Point struct {
	X, Y int
}

// Vertex is an example coordinate struct.
type Vertex struct {
	Lat, Long float64
}

// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

// Initialize structs as an example.
var (
	v1 = Point{3, 2}
	v2 = Point{Y: 4}
	v3 = Point{}
	p2 = &Point{5, 2}
)

// Pic generates a small image.
func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for i := range img {
		img[i] = make([]uint8, dx)
		for j := range img[i] {
			img[i][j] = uint8((i + j) / 2)
		}
	}
	return img
}

// WordCount counts words in string and returns a map without string.Fields.
func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	curWord := ""
	for i, v := range s {
		if v == ' ' || i == len(s)-1 {
			if i == len(s)-1 {
				curWord += string(v)
			}

			_, ok := wordMap[curWord]
			if ok {
				wordMap[curWord]++
			} else {
				wordMap[curWord] = 1
			}
			curWord = ""
		} else {
			curWord += string(v)
		}
	}
	return wordMap
}

// compute computes a given function with two args with (3, 4).
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// adder is a closure function which returns a function that adds to the sum.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	a := 1
	b := 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	q := Point{4, 6}
	pq := &q
	q.Y = -9
	pq.X = 100
	fmt.Println(q.Y)
	fmt.Println(q)
	fmt.Println(v1, v2, v3, p2)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [10]int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println(primes)

	var s = primes[2:4]
	fmt.Println(s)
	s[1] = 300
	fmt.Println(s, primes)

	// Slice literals!!!
	q5 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q5)

	r5 := []bool{true, false, true, true, false, true}
	fmt.Println(r5)

	s5 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s5)
	fmt.Println("Slices have default 0 and length of array as bounds.")
	fmt.Println("Slices length is how many elements the slice has.")
	fmt.Println("Slices capacity is how many elements the underlying array has.")
	printSlice(s)
	fmt.Println("A nil slice has a length and capacity of 0 and has no underlying array.")
	var nilll []int
	printSlice(nilll)
	if nilll == nil {
		fmt.Println("Slice is equal to nil!")
	}

	// Make function can create slices (dynamically sized arrays).
	a6 := make([]int, 5)
	a6[2] = 9
	printSlice(a6)
	b6 := make([]int, 0, 5)
	printSlice(b6)
	c6 := b6[:2]
	printSlice(c6)
	d6 := c6[2:5]
	printSlice(d6)

	// Append function can append elements to slices.
	d6 = append(d6, 2, 3, 4, 5, 6)
	printSlice(d6)

	tictactoeBoard := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	tictactoeBoard[0][0] = "X"
	tictactoeBoard[1][1] = "O"
	tictactoeBoard[0][1] = "X"
	tictactoeBoard[0][2] = "O"
	tictactoeBoard[2][0] = "X"
	tictactoeBoard[1][0] = "O"
	tictactoeBoard[1][2] = "X"
	tictactoeBoard[2][1] = "O"

	for i := 0; i < len(tictactoeBoard); i++ {
		fmt.Printf("%s\n", strings.Join(tictactoeBoard[i], " "))
	}

	// Range function to get index and value of array/slice.
	var example2 = []int{1, 3, 9, 27, 81}
	for i, v := range example2 {
		fmt.Printf("Index %d and value %d\n", i, v)
	}
	fmt.Println("Omit index or value by using _ instead of i or v.")

	// pic.Show(Pic)

	// map is like a dict in Python.
	m2 := make(map[string]Vertex)
	m2["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m2["Bell Labs"])
	fmt.Println(m)

	m4 := make(map[string]int)

	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])

	delete(m4, "Answer")
	fmt.Println("The value:", m4["Answer"])

	v8, ok := m4["Answer"]
	fmt.Println("The value:", v8, "Present?", ok)

	wc.Test(WordCount)

	fmt.Println("Functions are values too. They can be passed around just like other values.")
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
