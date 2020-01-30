package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
)

type Vertex struct {
	X, Y float64
}

// Method has special receiver, e.g. v Vertex between func and name.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// A method can be esily converted to a normal function:
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Stringer Exercise
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[2])
}

// Error Exercise
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.1f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

// Exercise Reader
type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, e error) {
	b[0] = 'A'
	b[1] = 'A'
	return 2, nil
}

// Exercise rot13reader
type rot13Reader struct {
	r io.Reader
}

func rot13(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		char += 13
		if char > 'Z' {
			char -= 26
		}
	} else if char >= 'a' && char <= 'z' {
		char += 13
		if char > 'z' {
			char -= 26
		}
	}

	return char
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	m, e := rot.r.Read(b)
	for i := 0; i < m; i++ {
		b[i] = rot13(b[i])
	}
	return m, e
}

// Exercise Images
type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8((x ^ y)), uint8((x ^ y)), 255, 255}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs(), AbsFunc(v))
	fmt.Println("You can only declare a receiver whose type is defined in the same package as the method.")
	fmt.Println("So you cant make methods with built-in types like int or float64!")
	fmt.Println("But you can use this type: type MyFloat float64.")
	fmt.Println("Pointer receivers are more common than value receivers because they can change the actual value to which the pointer points.")
	fmt.Println("Value receivers are just copies of the 'original' struct/type and so they cant modify the value.")
	fmt.Println("As a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.")
	fmt.Println("Same for value receivers. In this case, the method call p.Abs() is interpreted as (*p).Abs().")
	fmt.Println("Pointer receivers are in general faster than value pointers because they dont need to copy the whole struct.")
	fmt.Println("In general, use in 1 package either only pointer receivers or value receivers.")
	fmt.Println("An interface type is defined as a set of method signatures. A value of interface type can hold any value that implements those methods.")
	fmt.Println("A type implements an interface by implementing its methods. This happens implicitly.")

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	// Exercise stringers
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	reader.Validate(MyReader{})

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	m := Image{}
	pic.ShowImage(m)
}
