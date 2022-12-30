package xpln

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type CodeBlock struct {
	File      string
	Lang      string
	Comment   string
	Block     string
	Explained string
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
	} else {
		s = s - 1
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
