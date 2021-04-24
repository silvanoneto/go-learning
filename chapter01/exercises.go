// Package chapter01 exercises contains all the chapter exercises resolutions
package chapter01

import (
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
