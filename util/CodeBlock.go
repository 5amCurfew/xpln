package xpln

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type CodeBlock struct {
	file    string
	lang    string
	comment string
	code    string
}

func NewCodeBlock(f, s, e string) CodeBlock {
	return CodeBlock{
		file:    f,
		lang:    determineLang(f),
		comment: determineComment(f),
		code:    readFile(f, s, e),
	}
}

func determineLang(file string) string {
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

func determineComment(file string) string {
	switch strings.Split(file, ".")[1] {
	case "py":
		return "#"
	case "R":
		return "#"
	default:
		return "//"
	}
}

func readFile(file, start, end string) string {

	content, err := os.Open(file)

	if err != nil {
		log.Fatal("Error: Could not find file")
	}

	defer content.Close()

	raw, err := ioutil.ReadAll(content)

	var lines = strings.Split(string(raw), "\n")

	var s, startNotProvided = strconv.Atoi(start)
	if startNotProvided != nil {
		fmt.Println("Starting line not provided: Defaulting to line 1")
		s = 0
	} else if s > len(lines)-1 {
		fmt.Println("Starting line is greater than file length: Defaulting to line 1")
		s = 0
	}

	var e, endNotProvided = strconv.Atoi(end)
	if endNotProvided != nil {
		fmt.Println("Ending line not provided: Defaulting to end of file")
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

func (c CodeBlock) GetLang() string {
	return c.lang
}

func (c CodeBlock) GetComment() string {
	return c.comment
}

func (c CodeBlock) GetCode() string {
	return c.code
}

func (c CodeBlock) FormatBlock() string {

	var lines = strings.Split(string(c.code), "\n")

	var formatted string
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) > 200 {
			formatted += string(lines[i][:200]) + "...\n"
		} else {
			formatted += string(lines[i]) + "\n"
		}
	}

	return formatted
}
