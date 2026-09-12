package main

import (
	"fmt"
	"sample/module1/greetings"
)

func main() {
	fmt.Println("Hello, World!")

	message := greetings.Hello("Zhixian")
	fmt.Println(message)
}
