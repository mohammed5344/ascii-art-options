package justify

import (
	"fmt"
	"os"
	"strings"
)

func JustifyCase(input, target, ansiCode string, array []string) string {
	output := ""
	processed := strings.ReplaceAll(input, `\n`, "\n")
	var arr []string
	str := ""

	// make an array where newlines are separators and also elements
	for i := 0; i < len(processed); i++ {
		if processed[i] == '\n' {
			if str != "" {
				arr = append(arr, strings.TrimSpace(str))
				str = ""
			}
			arr = append(arr, strings.TrimSpace(str))
		} else {
			str += string(processed[i])
		}
	}
	if str != "" {
		arr = append(arr, strings.TrimSpace(str))
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == "\n" || arr[i] == "" {
			continue
		}
		arr[i] = spaces(arr[i], array)
	}
	input = ""
	for i := 0; i < len(arr); i++ {
		if arr[i] == "\n" {
			input += "\n"
			continue
		}
		input += arr[i]
	}
	output = printJustify(input, array, target, ansiCode)
	return output
}

func spaces(s string, array []string) string {
	words := strings.Fields(s)
	if len(words) == 1 {
		return s
	}
	width := getWidth()
	length := length(strings.ReplaceAll(s, " ", ""), array)
	spaces := (width - length) / (len(words) - 1)
	output := ""
	for i := 0; i < len(words)-1; i++ {
		output += words[i]
		output += strings.Repeat(" ", spaces)
	}
	output += words[len(words)-1]
	return output
}

func length(s string, array []string) int {
	length := 0
	for i := 0; i < len(s); i++ {
		max := 0
		lines := strings.Split(array[rune(s[i])-32], "\n")
		for j := 0; j < len(lines); j++ {
			if len(lines[j]) > max {
				max = len(lines[j])
			}
		}
		length += max
	}
	return length
}

func printJustify(input string, array []string, target string, ansiCode string) string {
	processed := strings.ReplaceAll(input, `\n`, "\n")
	var arr []string
	str := ""
	// make an array where newlines are separators and also elements
	for i := 0; i < len(processed); i++ {
		if processed[i] == '\n' {
			if str != "" {
				arr = append(arr, strings.TrimSpace(str))
				str = ""
			}
			arr = append(arr, strings.TrimSpace(str))
		} else {
			str += string(processed[i])
		}
	}

	if str != "" {
		arr = append(arr, strings.TrimSpace(str))
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
			os.Exit(1)
		}
	}
	//taking the correct index of target
	var index int
	count := -1 //counting in which element we are
	for _, el := range arr {
		index = strings.Index(el, target)
		count++
		if index >= 0 {
			break
		}
	}

	result := ""
	for i := 0; i < len(arr); i++ {
		if arr[i] == "\n"  {
			result += "\n"
			continue
		}
		// loop into n lines of each char then add it to the result
		for j := 0; j < 8; j++ {
			for x := 0; x < len(arr[i]); x++ {
				if rune(arr[i][x]) == ' ' {
					result += " "
					continue
				}
				lines := strings.Split(array[rune(arr[i][x]-32)], "\n")
				if (i == count && x >= index && x < index+len(target)) || (target == "" && ansiCode != "") {
					result += ansiCode + lines[j] + "\033[0m"
				} else {
					result += lines[j]
				}
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
