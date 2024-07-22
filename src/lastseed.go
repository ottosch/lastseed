package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ottosch/lastseed/src/seed"
	"github.com/ottosch/lastseed/src/table"
)

var (
	spacesRegex = regexp.MustCompile(`\s{2,}`)
	reader      = bufio.NewReader(os.Stdin)
)

func main() {
	words := readWordsFromCliArgs()
	for words == "" {
		fmt.Println("Enter seed words: ")
		words, _ = reader.ReadString('\n')
	}

	words = spacesRegex.ReplaceAllString(strings.TrimSpace(strings.ToLower(words)), " ")
	s := seed.NewSeed(words)

	fmt.Println()

	table.DrawSummary(s)
	table.DrawResults(s)
}

func readWordsFromCliArgs() string {
	args := os.Args
	if len(args) == 1 {
		return ""
	}

	return strings.Join(args[1:], " ")
}
