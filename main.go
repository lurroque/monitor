package main

import "fmt"

func main() {
	name := "Roque"
	version := 1.1
	fmt.Printf("Hi %+v!\n", name)
	fmt.Printf("Version: %+v\n", version)

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

}
