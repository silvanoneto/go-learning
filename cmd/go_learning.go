package main

import (
	"flag"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/silvanoneto/go-learning/pkg/chsutil"
	"github.com/silvanoneto/go-learning/pkg/util"
)

func main() {
	functionNamePtr := flag.String("name", chsutil.Chapter01HelloWorld,
		"Name of function to be called")
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

	f, err := chsutil.GetFunction(*functionNamePtr)
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now()
	f()
	util.TimeTrack(now, *functionNamePtr)
}
