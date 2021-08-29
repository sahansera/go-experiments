package main

import (
	"fmt"
	"time"
)

func iterate(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go iterate("goroutine")
	iterate("direct")

	fmt.Println("running")

	time.Sleep(time.Second) // this lets main thread to wait for a bit until the goroutine is done
	fmt.Println("done")
}
