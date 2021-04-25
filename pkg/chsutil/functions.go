// Package chsutil functions has facilities to call chapters functions
package chsutil

import (
	"errors"

	"github.com/silvanoneto/go-learning/pkg/chapter01"
)

const (
	// Chapter01HelloWorld is a function name from chapter01 package
	Chapter01HelloWorld string = "chapter01.HelloWorld"
	Chapter01Echo1      string = "chapter01.Echo1"      // Chapter01Echo1 is a function name from chapter01 package
	Chapter01Echo2      string = "chapter01.Echo2"      // Chapter01Echo2 is a function name from chapter01 package
	Chapter01Echo3      string = "chapter01.Echo3"      // Chapter01Echo3 is a function name from chapter01 package
	Chapter01Exercise01 string = "chapter01.Exercise01" // Chapter01Exercise01 is a function name from chapter01 package
	Chapter01Exercise02 string = "chapter01.Exercise02" // Chapter01Exercise02 is a function name from chapter01 package
	Chapter01Exercise03 string = "chapter01.Exercise03" // Chapter01Exercise03 is a function name from chapter01 package
	Chapter01Dup1       string = "chapter01.Dup1"       // Chapter01Dup1 is a function name from chapter01 package
	Chapter01Dup2       string = "chapter01.Dup2"       // Chapter01Dup2 is a function name from chapter01 package
	Chapter01Dup3       string = "chapter01.Dup3"       // Chapter01Dup3 is a function name from chapter01 package
	Chapter01Exercise04 string = "chapter01.Exercise04" // Chapter01Exercise04 is a function name from chapter01 package
	Chapter01Lissajous  string = "chapter01.Lissajous"  // Chapter01Lissajous is a function name from chapter01 package
	Chapter01Exercise05 string = "chapter01.Exercise05" // Chapter01Exercise05 is a function name from chapter01 package
	Chapter01Exercise06 string = "chapter01.Exercise06" // Chapter01Exercise06 is a function name from chapter01 package
	Chapter01Fetch      string = "chapter01.Fetch"      // Chapter01Fetch is a function name from chapter01 package
	Chapter01Exercise07 string = "chapter01.Exercise07" // Chapter01Exercise07 is a function name from chapter01 package
	Chapter01Exercise08 string = "chapter01.Exercise08" // Chapter01Exercise08 is a function name from chapter01 package
	Chapter01Exercise09 string = "chapter01.Exercise09" // Chapter01Exercise09 is a function name from chapter01 package
)

// GetFunction finds a mapped function from chapters packages
func GetFunction(name string) (func(), error) {
	functions := map[string]func(){
		Chapter01HelloWorld: chapter01.HelloWorld,
		Chapter01Echo1:      chapter01.Echo1,
		Chapter01Echo2:      chapter01.Echo2,
		Chapter01Echo3:      chapter01.Echo3,
		Chapter01Exercise01: chapter01.Exercise01,
		Chapter01Exercise02: chapter01.Exercise02,
		Chapter01Exercise03: chapter01.Exercise03,
		Chapter01Dup1:       chapter01.Dup1,
		Chapter01Dup2:       chapter01.Dup2,
		Chapter01Dup3:       chapter01.Dup3,
		Chapter01Exercise04: chapter01.Exercise04,
		Chapter01Lissajous:  chapter01.Lissajous,
		Chapter01Exercise05: chapter01.Exercise05,
		Chapter01Exercise06: chapter01.Exercise06,
		Chapter01Fetch:      chapter01.Fetch,
		Chapter01Exercise07: chapter01.Exercise07,
		Chapter01Exercise08: chapter01.Exercise08,
		Chapter01Exercise09: chapter01.Exercise09,
	}

	function := functions[name]
	if function == nil {
		return nil, errors.New("the function has not been encountered")
	}
	return function, nil
}
