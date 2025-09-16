package main

import (
	"asciiartoptions/colorChange"
	"asciiartoptions/justify"
	"asciiartoptions/output"
	"asciiartoptions/utils"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	/*ValidInput function will check that the input is valid and if so will return
	the input string and the banner */
	isValid, input, banner, inputIndex := utils.ValidInput(args)
	if !isValid {
		fmt.Println("input is invalid!")
		fmt.Println("Correct Usage: go run . [OPTION] <string> [BANNER]")
		os.Exit(0)
	}
	if len(input) == 0 {
		fmt.Println("input is invalid!")
		fmt.Println("Correct Usage: go run . [OPTION] <string> [BANNER]")
		os.Exit(0)
	}
	//check if all characers are printable characters
	for _, char := range input {
		if !(char > 31 && char < 127) {
			fmt.Println("only printable characters allowed")
			os.Exit(0)
		}
	}

	array := utils.Splice(banner)
	ansiCode, target := colorChange.Color(args, input)
	result := utils.PrintAscii(input, array, target, ansiCode)
	result = justify.Justify(args[0:inputIndex],array, result, input, target, ansiCode)
	output.Output(result, args, inputIndex)
	fmt.Println(result)
}
