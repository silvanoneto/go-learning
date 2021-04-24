// Package chapter01 03_dup contains the code studied on section 03
package chapter01

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Dup1 shows all duplicated text lines from the standart input (stdin),
// and their respective counts
func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Dup2 shows all duplicated text lines from the standart input (stdin) or
// file(s), and their respective counts
func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

// Dup3 shows all duplicated text lines from file(s), and their respective
// counts
func Dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// CountList receives a text stream and count their unique lines
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
