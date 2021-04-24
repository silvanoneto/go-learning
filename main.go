package main

import (
	"github.com/silvanoneto/go-learning/util"
)

func main() {
	results := []util.Result{
		{
			FunctionsList: []func(){
				// chapter01.HelloWorld,
				// chapter01.Echo1,
				// chapter01.Echo2,
				// chapter01.Echo3,
				// chapter01.Exercise01,
				// chapter01.Exercise02,
				// chapter01.Exercise03,
				// chapter01.Dup1,
				// chapter01.Dup2,
				// chapter01.Dup3,
				// chapter01.Exercise04,
			},
			Name: "Chapter 01",
		},
	}

	for _, result := range results {
		err := result.PrintResults()
		if err != nil {
			panic(err)
		}
	}
}
