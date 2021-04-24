// Package util result_printer automates the execution and print of functions results
package util

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// Result collects functions to be performed at once
type Result struct {
	Name          string
	FunctionsList []func()
}

// PrintResults runs all the function collection and print their execution times
func (result *Result) PrintResults() error {
	if len(result.Name) == 0 {
		return errors.New("result name is required")
	}

	fmt.Println(result.Name)
	fmt.Println("------------")
	fmt.Println()

	if result.FunctionsList == nil || len(result.FunctionsList) == 0 {
		fmt.Println("No functions available to run")
		fmt.Println()
	} else {
		for i, f := range result.FunctionsList {
			funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
			fmt.Printf("(%v) %v:\n", i+1, funcName)
			now := time.Now()
			f()
			TimeTrack(now, "Function")
			fmt.Println()
		}
	}

	fmt.Println()

	return nil
}
