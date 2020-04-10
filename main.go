package main

import "fmt"

func main() {
	showIntro()
	showMenu()
	option := readOption()
	switch option {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Logs...")
	case 0:
		fmt.Println("Bye!")
	default:
		fmt.Println("Unknown option!")
	}

}

func showIntro() {
	name := "Roque"
	version := 1.1
	fmt.Printf("Hi %+v!\n", name)
	fmt.Printf("Version: %+v\n", version)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func readOption() int {
	var option int
	fmt.Scan(&option)
	return option
}

func startMonitoring() {
	fmt.Println("Monitoring...")
}
