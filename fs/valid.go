package utils

import (
	"fmt"
	"os"
	"strings"
)

func ValidInput(args []string) (bool, string, string, int) {
	valid := false
	inputIndex := -1
	input := args[len(args)-1]
	//there is only one argument and it must be the input string
	if len(args) == 1 {
		return true, input, "standard", 0
	} else {
		valid = true
	}
	//check if there is a banner
	banner := strings.ToLower(args[len(args)-1])
	bannerIncluded := false
	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
		input = args[len(args)-2]
		inputIndex = len(args) - 2
		if CheckFlags(input) {
			banner = "standard"
			input = args[len(args)-1]
			inputIndex = len(args) - 1
		} else {
			bannerIncluded = true
		}
	} else {
		banner = "standard"
		input = args[len(args)-1]
		inputIndex = len(args) - 1
	}
	if inputIndex == -1 {
		return false, "", "", -1
	}
	if bannerIncluded {
		for i := 0; i < inputIndex; i++ {
			if !CheckFlags(args[i]) {
				if i > 0 {
					if strings.HasPrefix(strings.ToLower(args[i-1]), "--color=") {
						continue
					}
				}
				return false, "", "", -1
			}
		}
	} else {
		for i := 0; i < inputIndex; i++ {
			if !CheckFlags(args[i]) {
				if i > 0 {
					if strings.HasPrefix(strings.ToLower(args[i-1]), "--color=") {
						continue
					}
				}
				return false, "", "", -1
			}
		}
	}
	count := 0
	for i := 0; i < inputIndex; i++ {
		if strings.HasPrefix(strings.ToLower(args[i]), "--align=") {
			pos := strings.ToLower(args[i][8:])
			if pos != "right" && pos != "left" && pos != "center" && pos != "justify" {
				fmt.Println("Error: --align input is invalid!")
				os.Exit(0)
			}
			count++
		}
	}
	if count > 1 {
		fmt.Println("Error: The flag * --align * cannot be used more than once")
		os.Exit(0)
	}
	return valid, input, banner, inputIndex
}

// function to check if the input is a flag
func CheckFlags(s string) bool {
	s = strings.ToLower(s)
	flags := []string{"--color=", "--align=", "--output=", "--reverse="}
	for i := 0; i < len(flags); i++ {
		if strings.HasPrefix(s, flags[i]) {
			return true
		}
	}
	return false
}
