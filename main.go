package main

import (
	"github.com/silvanoneto/go-learning/chapter01"
	"github.com/silvanoneto/go-learning/util"
)

func main() {
	results := []util.Result{
		{
			FunctionsList: []func(){
				chapter01.HelloWorld,
				chapter01.Echo1,
				chapter01.Echo2,
				chapter01.Echo3,
				chapter01.Exercise01,
				chapter01.Exercise02,
				chapter01.Exercise03,
			},
			Name: "Chapter 01",
		},
		{
			FunctionsList: []func(){},
			Name:          "Chapter 02",
		},
		{
			FunctionsList: nil,
			Name:          "Chapter 03",
		},
	}

	for _, result := range results {
		err := result.PrintResults()
		if err != nil {
			panic(err)
		}
	}
}
