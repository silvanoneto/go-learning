package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/silvanoneto/go-learning/pkg/chsutil"
	"github.com/silvanoneto/go-learning/pkg/util"
)

func main() {
	functionNamePtr := flag.String("name", chsutil.Chapter01HelloWorld,
		"Name of function to be called")
	flag.Parse()

	if functionNamePtr == nil || len(*functionNamePtr) == 0 {
		log.Fatal("You need to enter at least the name of a function." +
			" Use -h for more information.")
	}

	f, err := chsutil.GetFunction(*functionNamePtr)
	if err != nil {
		log.Fatalln(err)
	}

	os.Args = append([]string{os.Args[0]}, flag.Args()...)

	now := time.Now()
	f()
	util.TimeTrack(now, *functionNamePtr)
}
