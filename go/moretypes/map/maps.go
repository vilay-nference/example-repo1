package main

import (
	"fmt"
	"math"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

func mapExample() {
	//var m map[string]Vertex
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m)
}

func mapLiterals() {
	m := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func mapMutation() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println(m["Answer"])
	m["Answer"] = 48
	fmt.Println(m["Answer"])
	delete(m, "Answer")
	fmt.Println(m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("The value ", v, "Present?", ok)
}

func mapExercise(s string) map[string]int {
	chars := strings.Split(s, " ")
	count := make(map[string]int)
	for _, v := range chars {
		c, ok := count[v]
		if ok {
			count[v] = c + 1
		} else {
			count[v] = 1
		}
	}
	fmt.Println(count)
	return count
}

// Function as value
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionAsValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// Function as closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func functionClosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func fibonacci() func() int {
	f0 := 0
	f1 := 1
	return func() int {
		temp := f0
		f0 = f1
		f1 += temp
		return temp
	}
}

func main() {
	fmt.Println("Welcome to learn map")
	mapExample()
	mapLiterals()
	mapMutation()
	mapExercise("This exersice is to count the character in the sentense")

	// function as value
	functionAsValues()
	functionClosure()

	// exercise with fibonacci number
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
