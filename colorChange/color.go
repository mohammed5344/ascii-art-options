package colorChange

import (
	"asciiartoptions/utils"
	"fmt"
	"os"
	"strings"
)

func Color(args []string, input string) (string, string) {
	if len(args) == 1 {
		return "", ""
	}

	count := 0
	clr := ""
	for _, arg := range args {
		if strings.HasPrefix(strings.ToLower(arg), "--color=") {
			clr = strings.ToLower(arg[8:])
			count++
		}
	}

	if count == 0 {
		return "", ""
	}

	if count > 1 {
		fmt.Println("Error: Cannot choose more than one color")
		os.Exit(1)
	}

	target := ""
	for i, el := range args {
		if strings.HasPrefix(el, strings.ToLower("--color=")) {
			if !utils.CheckFlags(args[i+1]) && args[i+1] != input {
				target = args[i+1]
				break
			}
		}
	}

	if !(strings.Contains(input, target)) {
		fmt.Printf("The substring '%s' was not found in the input string\n", target)
		os.Exit(1)
	}
	colors := map[string]string{
		"black":  "\033[30m",
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",

		"bright_black":  "\033[90m",
		"bright_red":    "\033[91m",
		"bright_green":  "\033[92m",
		"bright_yellow": "\033[93m",
		"bright_blue":   "\033[94m",
		"bright_purple": "\033[95m",
		"bright_cyan":   "\033[96m",
		"bright_white":  "\033[97m",

		"gray":    "\033[38;5;245m", // mid gray
		"pink":    "\033[38;5;205m",
		"orange":  "\033[38;5;208m", // bright orange
		"teal":    "\033[38;5;30m",
		"lime":    "\033[38;5;118m",
		"gold":    "\033[38;5;220m",
		"violet":  "\033[38;5;177m",
		"brown":   "\033[38;5;94m",
		"skyblue": "\033[38;5;117m",

		"reset": "\033[0m",
	}

	ansiCode, ok := colors[clr]
	if clr != "" {
		if !ok {
			fmt.Printf("Error: wrong or unspported color '%s'", clr)
			os.Exit(1)
		}
	}

	return ansiCode, target
}
