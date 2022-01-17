package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"strings"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}

func ErrorExample() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// Excersice
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("at %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func SqrtExercise() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

// Readers
func readerExample() {
	fmt.Println("Reader example")
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Println(n, err, b)
		fmt.Println(b[:n])
		if err == io.EOF {
			break
		}
	}
}

// Image example
func imageExample() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func main() {
	// Error example
	ErrorExample()
	SqrtExercise()
	readerExample()
	imageExample()
}
