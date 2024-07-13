package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	urls := os.Args[1:]
	file, err := os.OpenFile("/etc/hosts", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new writer
	writer := bufio.NewWriter(file)
	for _, line := range urls {
		fmt.Println("Blocking " + line + "...")
		_, err := writer.WriteString("127.0.0.1 " + line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Flush the writer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing the writer:", err)
		return
	}

	fmt.Println("URLs blocked successfully!")
}
