package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntro()
	for {
		showMenu()
		option := readOption()
		switch option {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Logs...")
		case 0:
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Unknown option!")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Roque"
	version := 1.1
	fmt.Printf("Hi %+v!\n", name)
	fmt.Printf("Version: %+v\n", version)
	fmt.Println("-----------------------------------------")
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
	sites := []string{"https://www.globo.com", "https://www.terra.com.br"}
	for _, s := range sites {
		testConnection(s)
	}
	fmt.Println("-----------------------------------------")
}

func testConnection(site string) {
	r, err := http.Get(site)
	if err != nil {
		fmt.Println(err)
	}
	if r.StatusCode == 200 {
		fmt.Printf("Site ON: %+v\n", site)
	}
	if r.StatusCode != 200 {
		fmt.Printf("Site connection with problems: %+v\n", site)
		fmt.Printf("StatusCode: %+v\n", r.StatusCode)
	}
}
