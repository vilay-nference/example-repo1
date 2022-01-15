package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func Welcome() {
	fmt.Println("Welcome to playground")
}

func timeDemo() {
	fmt.Println("The time is ", time.Now())
}

func random() {
	rand.Seed(time.Now().UnixMicro())
	fmt.Println("The random number is : ", rand.Intn(10))
}

func mathDemo() {
	fmt.Printf("Now you have %g Problems . \n", math.Sqrt(7))
	fmt.Println(math.Pi)
}

func add(x int, y int) int {
	return x + y
}

func addWithSameParamType(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func varExample() {
	var i int
	fmt.Println(i, c, python, java)
}

var i, j int = 1, 2

func varInitializerExample() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func variableWithinFunction() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

/**
Basic types in the go are as follow
bool
string
int, int8, int16, int32, int64
uint, unit8, uint16, uint32, uint64, uintptr
byte // alias for uint8
rune // Alias for int32 represent unicode code point
float32, float64
complex64, complex128

int, uint, unitptr are 32 bit on 32-bit machine while 64 bit on 64-bit machine.
When you need int use int unless you have specific reason to use sized or unsigned integer type
*/
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func dataTypes() {
	fmt.Printf("Type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value: %v\n", z, z)
}

func defaultValue() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

/*
Type casting
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	or put simply
	i := 42
	f := float64(i)
	u := uint(f)
*/
func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}

// Type inference
/*
var i int
j := i // This is int

But when right hand side is contain untyped numeric constant.
New variable may be int, float32, or complex128 based on precision.

i := 42            // int
f := 3.142         // float64
c := 0.867 + 0.5i  // complex128
*/
func typeInference() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}

const Pi = 3.14

func constant() {
	const World = "world"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstant() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func main() {
	// Example to just illustrate the go code
	Welcome()
	// Example to demonstrate the time
	timeDemo()
	// Random function use
	random()
	// Math library basics
	mathDemo()
	// Function with addition
	fmt.Println(add(4, 7))
	// Function with addition of same type
	fmt.Println(addWithSameParamType(4, 7))
	// Swapping of the string and send multiple parameter from the function
	x, y := swap("hello", "world")
	fmt.Println(x, y)
	// Naked return example
	fmt.Println(split(17))
	//Local and global variable example
	varExample()
	// variable example with initializer
	varInitializerExample()
	// variable declaration ways in function
	variableWithinFunction()
	// Basic types example
	dataTypes()
	// Default value of the variables
	defaultValue()
	// Type conversion
	typeConversion()
	// Type inference
	typeInference()
	// Use of constant
	constant()
	// Numeric constant
	numericConstant()
}
