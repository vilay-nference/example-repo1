package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

// Remember method is just a function with receiver argument
/*
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
*/

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// This can be donw with simple function as well
func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Method over non-struct type as well
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Pointer receiver
func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func PointerReceiverExample() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

func ScaleFunc(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func PointerInFunction() {
	v := Vertex{3, 4}
	ScaleFunc(&v, 10)
	fmt.Println(v.Abs())
}

// Method and pointer indirection
func PointerIndirection() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 5}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

/*
There are two reason to choose pointer
1. To modify the value
2. Avoid copying to the value on each method
3. In general given type should have pointer or value receiver but not both
*/
func pointerOverValue() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

func main() {
	fmt.Println("Welcome to learn methods in GO")
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	mf := MyFloat(-math.Sqrt2)
	fmt.Println(mf.Abs())

	PointerReceiverExample()
	PointerInFunction()
	PointerIndirection()
	pointerOverValue()
}
