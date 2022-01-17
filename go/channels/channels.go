package main

import (
	"fmt"
	"time"

	"golang.org/x/tour/tree"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func ChannelExample() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func BufferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func RangeAndClose() {
	fmt.Println("Range and close example")
	c := make(chan int, 10)
	go fibonacci1(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func SelectExample() {
	fmt.Println("Select example")
	c := make(chan int)
	q := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		q <- 0
	}()
	fibonacci(c, q)
}

// Default case in selection will be excuted if nothing is ready
func DefaultSelection() {
	tick := time.Tick(100 * time.Microsecond)
	boom := time.After(500 * time.Microsecond)
	for i := 0; i < 100; i++ {
		select {
		case <-tick:
			fmt.Println("Tick.")
		case <-boom:
			fmt.Println("Boom.")
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Microsecond)
		}
	}
}

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func WalkUtil(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go WalkUtil(t1, ch1)
	go WalkUtil(t2, ch2)
	for k := range ch1 {
		select {
		case g := <-ch2:
			if k != g {
				return false
			}
		default:
			break
		}
	}
	return true
}

func EquivalentTree() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

func main() {
	ChannelExample()
	BufferedChannels()
	RangeAndClose()
	SelectExample()
	//DefaultSelection()
	EquivalentTree()
}
