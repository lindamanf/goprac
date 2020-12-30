package main

import (
	"fmt"
)

func _main() {
	c := make(chan string, 4)

	// c <- "1"
	// fmt.Println("1")
	// c <- "2"
	// fmt.Println("2")
	// c <- "3"
	// fmt.Println("3")
	// c <- "4"
	// fmt.Println("4")

	go func(input chan string) {
		fmt.Println("sending 1 to the channel")
		input <- "hello1"

		fmt.Println("sending 2 to the channel")
		input <- "hello2"

		fmt.Println("sending 3 to the channel")
		input <- "hello3"

		fmt.Println("sending 4 to the channel")
		input <- "hello4"
	}(c)

	fmt.Println("receiving from the channel")
	for greeting := range c {
		fmt.Println("greeting received")

		fmt.Println(greeting)
	}
}

func helloWorld() {
	fmt.Println("Hello world")
}
