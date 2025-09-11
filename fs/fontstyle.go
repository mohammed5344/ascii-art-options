package fs

import (
	"fmt"
	"os"
	"strings"
)

func FontStyle(args []string) string {
	if len(args) == 1 {
		return "standard"
	}
	font := strings.ToLower(args[len(args)-1])
	if font != "standard" && font != "shadow" && font != "thinkertoy" {
		return "standard"
	} else {
		str := args[len(args)-2]
		if checkFlag(str) {
			fmt.Println("wrong format!")
			fmt.Println("Usage:go run . [option] <string> [fontStyle]")
			os.Exit(1)
		}
		return font
	}

	return "standard"
}

func checkFlag(s string) bool {
	flags := []string{"--color", "--algin", "--output", "--reverse"}
	s = strings.ToLower(s)
	for i := 0; i < len(flags); i++ {
		if strings.HasPrefix(s, flags[i]) {
			return true
		}
	}
	return false
}
