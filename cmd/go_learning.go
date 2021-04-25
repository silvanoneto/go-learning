package main

import (
	"flag"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/silvanoneto/go-learning/pkg/chapter01"
	"github.com/silvanoneto/go-learning/pkg/util"
)

func main() {
	functionNamePtr := flag.String(
		"name", "chapter01.HelloWorld", "Name of function to be called")
	flag.Parse()

	os.Args = append([]string{os.Args[0]}, flag.Args()...)

	if functionNamePtr == nil || len(*functionNamePtr) == 0 {
		log.Fatal("You need to enter at least the name of a function." +
			" Use -h for more information.")
	}

	isValid, err := regexp.MatchString(`chapter[0-9]+\..*`, *functionNamePtr)
	if err != nil {
		log.Fatalln(err)
	}
	if !isValid {
		log.Fatalln("The function name is invalid." +
			" Use -h for more information.")
	}

	functions := map[string]func(){
		"chapter01.HelloWorld": chapter01.HelloWorld,
		"chapter01.Echo1":      chapter01.Echo1,
		"chapter01.Echo2":      chapter01.Echo2,
		"chapter01.Echo3":      chapter01.Echo3,
		"chapter01.Exercise01": chapter01.Exercise01,
		"chapter01.Exercise02": chapter01.Exercise02,
		"chapter01.Exercise03": chapter01.Exercise03,
		"chapter01.Dup1":       chapter01.Dup1,
		"chapter01.Dup2":       chapter01.Dup2,
		"chapter01.Dup3":       chapter01.Dup3,
		"chapter01.Exercise04": chapter01.Exercise04,
		"chapter01.Lissajous":  chapter01.Lissajous,
		"chapter01.Exercise05": chapter01.Exercise05,
		"chapter01.Exercise06": chapter01.Exercise06,
		"chapter01.Fetch":      chapter01.Fetch,
		"chapter01.Exercise07": chapter01.Exercise07,
		"chapter01.Exercise08": chapter01.Exercise08,
		"chapter01.Exercise09": chapter01.Exercise09,
	}

	f := functions[*functionNamePtr]
	if f == nil {
		log.Fatalln("The function has not been encountered." +
			" Use -h for more information.")
	}

	now := time.Now()
	f()
	util.TimeTrack(now, *functionNamePtr)
}
