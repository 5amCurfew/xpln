package xpln

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type CodeBlock struct {
	File    string
	Lang    string
	Comment string
	Block   string
}

func DetermineLang(file string) string {
	switch strings.Split(file, ".")[1] {
	case "py":
		return "Python"
	case "js":
		return "Javascript"
	case "R":
		return "R Programming"
	default:
		return "Go"
	}
}

func DetermineComment(file string) string {
	switch strings.Split(file, ".")[1] {
	case "py":
		return "#"
	case "R":
		return "#"
	default:
		return "//"
	}
}

func ReadFile(file, start, end string) string {

	content, _ := os.Open(file)

	defer content.Close()

	raw, _ := ioutil.ReadAll(content)

	var lines = strings.Split(string(raw), "\n")

	var s, startNotProvided = strconv.Atoi(start)
	if startNotProvided != nil {
		s = 0
	} else if s > len(lines)-1 {
		fmt.Println("Starting line is greater than file length: Defaulting to line 1")
		s = 0
	}

	var e, endNotProvided = strconv.Atoi(end)
	if endNotProvided != nil {
		e = len(lines) - 1
	} else if e < s {
		fmt.Println("Ending line is greater than starting line: Defaulting to end of file")
		e = len(lines) - 1
	} else if e > len(lines)-1 {
		fmt.Println("Ending line is greater than file length: Defaulting to end of file")
		e = len(lines) - 1
	}

	// Selected lines
	var selected string
	for i := s; i <= e; i++ {
		if i >= s && i <= e {
			selected += string(lines[i]) + "\n"
		}
	}

	return selected
}

func (c CodeBlock) FormatBlockOutput(w int) string {

	var lines = strings.Split(string(c.Block), "\n")
	const maxOutput = 25

	var formatted string
	for i := 0; i < len(lines) && i <= maxOutput; i++ {
		if i == maxOutput {
			formatted += "..."
			break
		}
		if len(lines[i]) > w {
			formatted += string(lines[i][:w]) + "...\n"
		} else {
			formatted += string(lines[i]) + "\n"
		}
	}

	return formatted
}
