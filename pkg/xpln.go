package xpln

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadCodeBlock(file, start, end string) string {

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

	return (string(selected))
}
