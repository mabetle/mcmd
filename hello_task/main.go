package main

import (
	"fmt"
)

func main() {
	Hello()
}

func Hello() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Hello, World!\n")
	}
}
