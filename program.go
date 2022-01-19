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

const monitoringTurns = 3
const delayBetweenTurns = 5

func main() {

	showIntroduction()
	readSitesFromFile()
	// showNamesInSlice()

	for { // blank for loop works like "while true"
		showMenu()
		inferedCommand := readInput()

		switch inferedCommand {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
			printLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0) // safe exit to OS

		default:
			fmt.Println("I don't know this command!")
			os.Exit(-1) // indicates the OS an unexpected error occurred
		}

	}

}

func showIntroduction() {
	name := "Pedro"
	version := 1.1
	fmt.Println("Hello, mr.", name)
	fmt.Println("This program is running on version", version)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func readInput() int {
	var inferedCommand int
	// fmt.Scanf("%d", &userCommand) // "%d" makes us expect an integer, and the "&" points to the address of a variable, binding that value to the "command" variable
	fmt.Scan(&inferedCommand) // using .Scan instead of .Scanf, you don't need to specify the variable type
	fmt.Println("You choose command", inferedCommand, "Its address is", &inferedCommand)

	return inferedCommand
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	// sites := []string{ // creation of a Slice, which works like an ArrayList in Java/Kotlin
	// 	"https://random-status-code.herokuapp.com",
	// 	"https://www.alura.com.br",
	// 	"https://www.caelum.com.br"}

	sites := readSitesFromFile()

	for j := 0; j < monitoringTurns; j++ {
		for i, site := range sites { // here, site becomes a range of the site indexes, so it will be sites[0], sites[1] and sites[2]

			fmt.Println(site, "in position", i, "of our slice is being tested.")
			testWebsite(site)

		}
		time.Sleep(delayBetweenTurns * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func returnsNameAndAge() (string, int) {
	name := "Pedro"
	age := 28
	return name, age
}

func testMultipleReturn() {
	name, age := returnsNameAndAge()
	fmt.Println("I'm", name, "and I'm", age, "years old.")

	_, newAge := returnsNameAndAge()
	fmt.Println("I just ignored the function's first variable return. Also, I'm", newAge, "years old.")

	newName, _ := returnsNameAndAge()
	fmt.Println("I just ignored the function's first variable return. Also, I'm", newName)
}

func showNamesInSlice() {
	names := []string{"Pedro", "Amanda", "Jane", "Maru"}
	fmt.Println(names)
}

func testWebsite(site string) {

	resp, err := http.Get(site)

	if resp.StatusCode == 200 {

		fmt.Println("Site:", site, "was successfully loaded!")
		registerLogs(site, true)

	} else if err != nil {

		fmt.Println("An error occurred:", err)
		registerLogs(site, false)

	} else {

		fmt.Println("Site:", site, "had issues. Status code:", err)
		registerLogs(site, false)

	}
}

func readSitesFromFile() []string {

	var sites []string

	file, err := os.Open("sites.txt") // returns the file's address in memory
	// arquivo, err := ioutil.ReadFile("sites.txt") // returns the whole file as an array of bytes, can be printed if cast to string

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	fileReader := bufio.NewReader(file)

	for {
		line, err := fileReader.ReadString('\n') // reads the file until the line break
		line = strings.TrimSpace(line)           // trims spaces and newlines
		fmt.Println(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("An error occurred:", err)
		}

	}

	fmt.Println(sites)
	file.Close()

	return sites
}

func registerLogs(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()

}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
