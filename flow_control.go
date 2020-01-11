package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func guessSqrt(x float64) float64 {
	z := x / 2
	for math.Abs(math.Sqrt(x)-z) > 0.0001 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	defer fmt.Println("Im from a defer statement declared at the very beginning of this function!")
	defer fmt.Println("Im directly evaluated but not executed until this function ends.")
	defer fmt.Println("Defers are pushed onto a stack and thus are executed in reversed order!")

	sum := 1
	for sum < 1000 { // Just like a while in C
		sum += sum
	}

	fmt.Println(sum)
	fmt.Println("If condition of for loop is omitted, you end up with an infinite loop.")
	fmt.Println(sqrt(16), sqrt(-36))
	fmt.Println(pow(2, 5, 40), pow(2, 6, 40))
	fmt.Println(guessSqrt(5), guessSqrt(9))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
