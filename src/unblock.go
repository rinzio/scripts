package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func loadURLs() []string {
	lines := []string{}
	for _, url := range os.Args[1:] {
		lines = append(lines, "127.0.0.1 "+url)
	}
	return lines
}

func main() {
	filename := "/etc/hosts"
	urls := loadURLs()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !slices.Contains(urls, line) {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = os.WriteFile(filename, []byte(strings.Join(lines, "\n")+"\n"), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("URLs unblocked successfully!")
}
