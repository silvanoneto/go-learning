# go-learning

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/silvanoneto/go-learning.svg)](https://github.com/silvanoneto/go-learning)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/silvanoneto/go-learning)
[![GoReportCard example](https://goreportcard.com/badge/github.com/silvanoneto/go-learning)](https://goreportcard.com/report/github.com/silvanoneto/go-learning)
[![GitHub license](https://img.shields.io/github/license/silvanoneto/go-learning.svg)](https://github.com/silvanoneto/go-learning/blob/develop/LICENSE)

This repository contains my progress (re-)studying Go programming language based on "The Go Programming Language" book, written by Alan A. A. Donovan and Brian W. Kernighan ([gopl.io](https://www.gopl.io/ "The Go Programming Language")).

## CMD Application

The cmd app executes functions supported by [chsutil](https://github.com/silvanoneto/go-learning/tree/develop/pkg/chsutil) package.

### How to run it locally (Docker)

```sh
docker pull silvanoneto/go-learning-cmd:develop
docker run --rm silvanoneto/go-learning-cmd:develop
```

A command example that set a function and file names as parameters:

```sh
docker run --rm go-learning-cmd:develop /bin/sh -c "./go-learning -name chapter01.Exercise04 examples/data/01_03_dup_file1.txt examples/data/01_03_dup_file2.txt"
```

### How to run it locally (Manual)

You'll need to install Go in your machine. Follow the instructions in [https://golang.org/doc/install](https://golang.org/doc/install) or use the package manager of your preference.

After this:
1. Download the source code;
2. Open your terminal/prompt within the project folder path;
3. Run the following commands:
```sh
go mod download
go run cmd/go_learning.go
```

If you want to set the function to run and other parameters, follow the example:
```sh
go run cmd/go_learning.go -name chapter01.Exercise04 examples/data/01_03_dup_file1.txt examples/data/01_03_dup_file2.txt
```
