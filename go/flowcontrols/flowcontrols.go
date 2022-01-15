package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func forLoopExample() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forWithOptional() {
	sum := 1
	for ; sum < 1000; sum += sum {
		sum += sum
	}
	fmt.Println(sum)
}

func whileOfC() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/*
Infinite loop
for{

}
*/

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
	// can't use v here, though
	return lim
}

func sqrtCalculation(x float64) (float64, int) {
	z := 1.0
	i := 1
	for math.Abs(z*z-x) > .001 {
		z -= (z*z - x) / (2 * z)
		i += 1
	}
	return z, i
}

func switchExample() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s\n", os)
	}
}

func switchEvaluationOrder() {
	fmt.Println("When is saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("Day after tomorrow.")
	default:
		fmt.Println("Some other day.")
	}
}

func switchInCaseOfLongIfElse() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func deferExample() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func deferDeep() {
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}

func main() {
	// for loop example
	forLoopExample()
	// for with optional
	forWithOptional()
	// while of C
	whileOfC()
	// If condition example
	fmt.Println(sqrt(2), sqrt(-4))
	// If and else
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	// Sqrt calculation using newton rapson
	fmt.Println(sqrtCalculation(5))
	fmt.Println(math.Sqrt(5))
	// switch statement
	switchExample()
	// switch evaluation order
	switchEvaluationOrder()
	// use swith instead of long if else
	switchInCaseOfLongIfElse()
	// Defer example
	deferExample()
	// Defer in elobarote way
	deferDeep()
}
