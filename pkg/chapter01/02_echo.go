// Package chapter01 02_echo contains the code studied on section 02
package chapter01

import (
	"fmt"
	"os"
	"strings"
)

// Echo1 prints command line arguments, separated by blank spaces
func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Echo2 prints command line arguments, separated by blank spaces
func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Echo3 prints command line arguments, separated by blank spaces
func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
