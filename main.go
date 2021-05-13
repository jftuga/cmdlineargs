package main

import (
	"fmt"
	"os"
)

const version string = "1.0.0"

func main() {
	fmt.Println()
	fmt.Println("length:", len(os.Args))
	fmt.Println()
	for i, arg := range os.Args {
		fmt.Printf("[%2d] %s\n", i, arg)
	}
}
