package main

import (
	"fmt"
	"strings"
)

func pointers() {
	i, j := 42, 2701
	p := &i         // point i
	fmt.Println(*p) // read i through pointer
	*p = 21         // Set i through pointer
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(*p)
}

type Vertex struct {
	X int
	Y int
}

func structExample() {
	fmt.Println(Vertex{1, 2})
}

func structFields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X, v)
}

func pointerToStruct() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func structLiterals() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)
	fmt.Println(v1, v2, v3, p)
}

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}

func slicesAsReference() {
	names := [4]string{
		"Sameer",
		"Vilay",
		"Sonali",
		"Mayur",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	r := []bool{true, false, true, true, false, true}
	s := []struct {
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
	fmt.Println(q, r, s)
}

func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[:0]
	printSlice(s)
	s = s[:4]
	printSlice(s)
	s = s[2:]
	printSlice(s)
}

func nilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func sliceWithMake() {
	a := make([]int, 5)
	printSlice1("a", a)
	b := make([]int, 0, 5)
	printSlice1("b", b)
	c := b[:2]
	printSlice1("c", c)
	d := c[2:5]
	printSlice1("d", d)
}

func slicesOfSlices() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendingToSlice() {
	var s []int
	printSlice(s)
	s = append(s, 0)
	printSlice(s)
	s = append(s, 1)
	printSlice(s)
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func rangeExample() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func rangeContinued() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func sliceExcersize(dx, dy int) [][]uint8 {
	//r := make([][]uint8, dx)
	var r [][]uint8
	for i := 0; i < dx; i++ {
		//r[i] = make([]uint8, dy)
		var r1 []uint8
		for j := 0; j < dy; j++ {
			//r[i][j] = uint8(i+j) / 2
			r1 = append(r1, uint8(i+j)/2)
		}
		r = append(r, r1)
	}
	fmt.Println(r)
	return r
}

func main() {
	fmt.Println("Welcome to learn more types")
	// Learn pointers
	pointers()

	// struct example
	structExample()
	structFields()
	pointerToStruct()
	structLiterals()

	// Arrays
	arrays()

	// Slices
	slices()
	slicesAsReference()
	sliceLiterals()
	sliceDefaults()
	sliceLengthAndCapacity()
	nilSlice()
	sliceWithMake()
	slicesOfSlices()
	appendingToSlice()

	// Range example
	rangeExample()
	rangeContinued()

	// excersize
	sliceExcersize(3, 4)

}
