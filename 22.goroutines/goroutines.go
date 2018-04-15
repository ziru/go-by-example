package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
