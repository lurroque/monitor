package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 5
const delay = 3

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
			showLogs()
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
	sites := readFile()
	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			testConnection(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("-----------------------------------------")
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
		writeLog(site, true)
	}
	if r.StatusCode != 200 {
		fmt.Printf("Site connection with problems: %+v\n", site)
		fmt.Printf("StatusCode: %+v\n", r.StatusCode)
		writeLog(site, false)
	}
}

func readFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			break
		}
		sites = append(sites, line)
	}
	file.Close()
	return sites
}

func writeLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(time.Now().Format(
		"02/01/2006 15:04:05") + " - " + site + "- Online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	// O ioutil já fecha o arquivo
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
