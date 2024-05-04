package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type Counts struct {
	words int
	lines int
	bytes int
	chars int
}

func count(scanner *bufio.Scanner) (Counts, error) {
	c := Counts{}
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Fields(line)

		c.lines++
		c.words += len(words)
		c.bytes += len(line)
		c.chars += utf8.RuneCountInString(line)
	}

	if err := scanner.Err(); err != nil {
		return Counts{}, err
	}

	return c, nil
}

func getFileScanner(filename string) (*bufio.Scanner, io.Closer, error) {
	f, err :=
		os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	return bufio.NewScanner(f), f, nil
}

func formatCounts(c *Counts, lines, words, bytes, chars *bool) string {
	formattedCounts := ""

	if !(*lines || *words || *bytes || *chars) {
		*lines, *words, *bytes = true, true, true
	}

	if *lines {
		formattedCounts += fmt.Sprintf("%d  ", c.lines)
	}
	if *words {
		formattedCounts += fmt.Sprintf("%d  ", c.words)
	}
	if *bytes {
		formattedCounts += fmt.Sprintf("%d  ", c.bytes)
	}
	if *chars {
		formattedCounts += fmt.Sprintf("%d  ", c.chars)
	}

	return formattedCounts
}

func main() {
	lines := flag.Bool("l", false, "print the newline counts")
	bytes := flag.Bool("c", false, "print the byte counts")
	chars := flag.Bool("m", false, "print the charater counts")
	words := flag.Bool("w", false, "print the word counts")

	flag.Parse()

	var scanner *bufio.Scanner
	var filename string
	var file io.Closer

	if len(flag.Args()) > 0 {
		filename = flag.Arg(0)
		var err error
		scanner, file, err = getFileScanner(filename)
		defer file.Close()

		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	c, err := count(scanner)
	if err != nil {
		log.Fatalf("Failed to count: %v", err)
	}

	formattedCounts := formatCounts(&c, lines, words, bytes, chars)

	if filename != "" {
		formattedCounts += filename
	}
	fmt.Println(formattedCounts)

}
