package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var sitesThatAreBlocked []string

func read() []string {
	var lines []string

	file, err := os.Open("C:\\Windows\\System32\\drivers\\etc\\hosts")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "127.0.0.1") {
			lines = append(lines, line)
		}
	}

	return lines
}

func write(lines []string) {
	blacklisted := 0

	file, err := os.OpenFile("C:\\Windows\\System32\\drivers\\etc\\hosts", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	for _, site := range sitesThatAreBlocked {
		entry := "127.0.0.1     " + site
		if !contains(lines, entry) {
			_, err := file.WriteString("\n" + entry)
			if err != nil {
				panic(err)
			}
			fmt.Println("Blacklisted", site)
			blacklisted++
		}
	}

	fmt.Println("Done \nBlacklisted", blacklisted, " Sites\nPress any keys to close this application.")
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}

	return false
}

func fetchShit() {
	cmd := exec.Command("cmd", "/c", "C:\\Windows\\System32\\curl https://raw.githubusercontent.com/emirsassan/HostsModRepost/main/domains.txt > domains.txt")

	err := cmd.Run()

	if err != nil {
		return
	}

	data, err := os.ReadFile("domains.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		sitesThatAreBlocked = append(sitesThatAreBlocked, line)
	}
}

func main() {
	fetchShit()

	current := read()

	write(current)

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
