package output

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

func Output(result string, args []string) {
	count := 0
	file := ""
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(strings.ToLower(args[i]), "--output=") {
			if count == 0 {
				file = args[i][9:]
			} 
			count++
		}
	}
	if count == 0 {
		return
	} else if count > 1 {
		fmt.Println("The flag --output can only be used once")
		os.Exit(0)
	}
	if len(file) == 0 {
		fmt.Println("Error: wrong format")
		fmt.Println("Correct flag usage: go run . --output=file.txt <string> ")
		os.Exit(0)
	}
	if filepath.Ext(file) != ".txt" {
		fmt.Println("none valid file: " + file)
		fmt.Println("only .txt files can be used")
		os.Exit(0)
	}

	err := os.WriteFile(file, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(0)
	}
	if err == nil {
		fmt.Println("successfully written file!")
		os.Exit(0)
	}
	return
}