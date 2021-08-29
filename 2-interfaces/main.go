package main

import (
	"fmt"
)

type Moving interface {
	Start()
	Stop()
}

type Car struct{}

func (c Car) Start() {
	fmt.Println("Car is started")
}

func (c Car) Stop() {
	fmt.Println("Car is stopping")
}

func main() {
	var car Moving = Car{}
	car.Start()
	car.Stop()
}
