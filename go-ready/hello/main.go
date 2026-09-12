package main

import (
	"fmt"
	"go-ready/module1/greetings"
)

func main() {
	fmt.Println("Hello, World!")

	message := greetings.Hello("Zhixian")
	fmt.Println(message)
}
