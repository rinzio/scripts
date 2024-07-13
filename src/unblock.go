package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "/etc/hosts"
	url := "127.0.0.1 " + os.Args[1]

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
		if line != url {
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

	fmt.Println("URL unblocked successfully!")
}
