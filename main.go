package main

import (
	"bufio"
	"fmt"
	hah "hah/functions"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	Input := os.Args[1]
	inputFile, err := os.Open(Input)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	var lines []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// Print the lines stored in the array

	modifiedLine := hah.ModifyText(lines)
	fmt.Println(modifiedLine)

}