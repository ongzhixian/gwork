package main

import (
	"example/module1"
	"fmt"
)

func main() {
	fmt.Println("Hello world, Zhixian")

	message := module1.Hello("Zhixian")
	fmt.Println(message)
}
