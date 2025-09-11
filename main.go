package main

import (
	"asciiart/fs"
	"asciiart/output"
	"asciiart/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if strings.ToLower(args[0]) == "--help" {
		fmt.Println("Usage:go run . [option] <string> [fontStyle]")
		return
	}
	fontStyle := fs.FontStyle(args)
	input := args[len(args)-1]

	if len(input) == 0 {
		return
	}
	//check if all characers are printable characters
	for _, char := range input {
		if !(char > 31 && char < 127) {
			fmt.Println("only printable characters allowed")
			os.Exit(0)
		}
	}
	array := utils.Splice(fontStyle)
	result := utils.PrintAscii(input, array)
	output.Output(result, args)
	fmt.Println(result)

}
