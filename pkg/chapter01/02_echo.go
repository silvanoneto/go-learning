// Package chapter01 02_echo contains the code studied on section 02
package chapter01

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/silvanoneto/go-learning/pkg/util"
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
