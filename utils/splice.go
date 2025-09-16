package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Splice reads a file and returns them as a slice of strings
// Each string in the slice represents one ASCII character
func Splice(fontStyle string) []string {
	file := ""
	if fontStyle == "standard" {
		file = "assets/standard.txt"
	} else if fontStyle == "shadow" {
		file = "assets/shadow.txt"
	} else if fontStyle == "thinkertoy" {
		file = "assets/thinkertoy.txt"
	}
	input, err := os.Open(file)
	if err != nil {
		fmt.Println("Error: opening file")
		os.Exit(1)
	}
	defer input.Close()

	var arr []string
	var group []string
	count := 0
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			arr = append(arr, line)
			count++
		}
		if count == 8 {
			group = append(group, strings.Join(arr, "\n"))
			arr = []string{}
			count = 0
		}
	}
	return group
}
