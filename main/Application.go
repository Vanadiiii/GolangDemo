package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	for i := 0; i < 10; i++ {
		fmt.Print(i)
		if i == 5 {
			fmt.Println(" - this is my favourite number - ")
		} else {
			fmt.Println(" - number")
		}
	}
}
