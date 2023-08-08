package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readStdinInputText(query string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", query)
	text, _ := reader.ReadString('\n')
	return strings.TrimRight(text, "\n")
}
