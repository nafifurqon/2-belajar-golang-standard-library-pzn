package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Println("line", string(line), "isPrefix", isPrefix)
	}
}
