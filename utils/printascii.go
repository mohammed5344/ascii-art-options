package utils

import (
	"fmt"
	"os"
	"strings"
)

func PrintAscii(input string, array []string) string {
	processed := strings.ReplaceAll(input, `\n`, "\n")
	var arr []string
	str := ""

	// make an array where newlines are separators and also elements
	for i := 0; i < len(processed); i++ {
		if processed[i] == '\n' {
			if str != "" {
				arr = append(arr, str)
				str = ""
			}
			arr = append(arr, "\n")
		} else {
			str += string(processed[i])
		}
	}

	if str != "" {
		arr = append(arr, str)
	}

	// handle if only new lines included
	newline := ""
	for i := 0; i < len(arr); i++ {
		if arr[i] != "\n" {
			break
		}
		newline += "\n"
		if i == len(arr)-1 {
			fmt.Print(newline)
			os.Exit(0)
		}
	}

	result := ""
	for i := 0; i < len(arr); i++ {
		if arr[i] == "\n" {
			result += "\n"
			continue
		}
		// loop into n lines of each char then add it to the result
		for j := 0; j < 8; j++ {
			for x := 0; x < len(arr[i]); x++ {
				lines := strings.Split(array[rune(arr[i][x]-32)], "\n")
				result += lines[j]
			}
			if j != 7 {
				result += "\n"
			}
		}
	}
	result += " "
	// remove the last extra line
	return result[:len(result)-1]
}
