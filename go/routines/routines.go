package main

import (
	"fmt"
	"time"
)

func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func Routine() {
	go Say("World")
	Say("Hello")
}

func main() {
	Routine()
}
