package justify

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func Justify(args, arr []string, result string, input, target, ansiCode string) string {
	output := ""
	pos, array := helper(result, args)
	width := getWidth()
	for i := 0; i < len(array); i++ {
		if array[i] == "\n" {
			output += "\n"
			continue
		}
		spaces := width - getLen(array[i])
		lines := strings.Split(array[i], "\n")
		switch pos {
		case "right":
			for j := 0; j < 8; j++ {
				lines[j] = strings.Repeat(" ", spaces) + lines[j]
			}
		case "left":
			//
			continue
		case "center":
			for j := 0; j < 8; j++ {
				lines[j] = strings.Repeat(" ", spaces/2) + lines[j] + strings.Repeat(" ", spaces/2)
			}
		case "justify":
			return JustifyCase(input, target, ansiCode, arr)
		}
		output += strings.Join(lines, "\n")
		if i != len(array)-1 {
			output += "\n"
		}
	}

	return output
}

func getWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return width
}

func getLen(s string) int {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	max := 0
	lines := strings.Split(s, "\n")
	for i := 0; i < len(lines); i++ {
		lines[i] = re.ReplaceAllString(lines[i], "")
		if len(lines[i]) > max {
			max = len(lines[i])
		}
	}
	return max
}

func helper(result string, args []string) (string, []string) {
	result = strings.ReplaceAll(result, `\n`, "\n")
	array := []string{}
	lines := strings.Split(result, "\n")
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) != 0 {
			array = append(array, strings.Join(lines[i:i+8], "\n"))
			i += 7
		} else {
			array = append(array, "\n")
		}
	}
	pos := ""

	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(strings.ToLower(args[i]), "--align=") {
			pos = strings.ToLower(args[i][8:])
			break
		}
	}
	return pos, array
}
