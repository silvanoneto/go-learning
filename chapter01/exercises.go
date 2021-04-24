// Package chapter01 exercises contains all the chapter exercises resolutions
package chapter01

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/silvanoneto/go-learning/util"
)

// Exercise01 prints all command line arguments, separated by blank spaces
func Exercise01() {
	fmt.Println(strings.Join(os.Args, " "))
}

// Exercise02 prints all command line arguments and an index, one per line
func Exercise02() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

// Exercise03 prints the time execution for each implemented Echo function
func Exercise03() {
	echoFuncs := []func(){Echo1, Echo2, Echo3}

	for i, echo := range echoFuncs {
		funcName := fmt.Sprintf("Echo%v", i+1)
		now := time.Now()
		echo()
		util.TimeTrack(now, funcName)
	}
}

// Exercise04 shows all duplicated text lines from the standart input (stdin) or
// file(s), their respective counts and files where they was encountered
func Exercise04() {
	type count struct {
		CountNumber int
		Filenames   map[string]struct{}
	}

	getFileNames := func(c *count) []string {
		filenames := make([]string, 0, len(c.Filenames))
		for filename := range c.Filenames {
			filenames = append(filenames, filename)
		}
		return filenames
	}

	countLines := func(f *os.File, counts map[string]count) {
		input := bufio.NewScanner(f)
		for input.Scan() {
			val, ok := counts[input.Text()]
			if ok {
				val.CountNumber++
				val.Filenames[f.Name()] = struct{}{}
			} else {
				val = count{
					CountNumber: 1,
					Filenames: map[string]struct{}{
						f.Name(): {},
					},
				}
			}
			counts[input.Text()] = val
		}
	}

	counts := make(map[string]count)
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
			if n.CountNumber > 1 {
				fmt.Printf(
					"%d\t%s\t%v\n", n.CountNumber, line, getFileNames(&n))
			}
		}
	}
}
