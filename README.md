"# first-go-module" 

# How to use Go modules and packages

We want to create a simple Go module `first-go-module` that consists of one package, named `hello`. We develop the module and the package on the local PC and we follow the [Golang.org](https://golang.org/doc/code.html) conventions for naming modules paths to prepare for later delivery of that module in a repository. 

This cookbook describes how to do this.

## Go Environment

When Go is installed on your PC, it creates a default environment for you. The author used Go version 1.15 on Windows. The environment variables that matter are `GOPATH` and `GOBIN`.

`GOPATH` tells the Go compiler on your PC, where it can find modules and packages, other than the standard packages that are provided by the Go language itself and that are located using the `GOROOT` environment variable. Unless `GOBIN` is set, executables created by `go install` are also stored in a directory determined by `GOPATH`, `%GOPATH%\bin`.

`GOBIN` can be set to tell the Go compiler where executables are stored, when the defaults shouldn't be used.

The author's settings are:
```
> go env
set GOBIN=
set GOPATH=%USERPROFILE%\LearningGo
```

## Setup

The module and package resides on the local PC. All you need is a directory whose name matches the module name, here `first-go-module`. This is the working directory.
```
> mkdir first-go-module
> cd first-go-module
```
Next, we have to declare the *module path*, i.e. to tell where Go can find the module and package when you want to import it in your program. 

Since Go modules typically reside in a Github repository, the module path is constructed according to the following scheme `github.com/<user>/<module-name>`. It doesn't matter if such repository does not yet exist.

To declare the module path, use the `go mod` command which produces a file called `go.mod`:
```
> go mod init github.com/jmh-git/first-go-module
> cat go.mod
module github.com/jmh-git/first-go-module

go 1.15
```
Now, we have a proper module skeleton. In the next section, we create some content.

## First program

In the working directory, we create a `main.go`. It is an executable program and hence specifies the `package main` in the first line.
```
package main

import "fmt"

func main() {
	fmt.Println("first-go-module/main.go :: main()")
	fmt.Println("(executable module)")
}
``` 

## Install first program

In the working directory, we create an executable program using `go install`. This command produces an executable and stores it in a location that depends on the setting of the `GOBIN` and `GOPATH` environment variables. With the settings described above, the command
```
go install github.com/jmh-git/first-go-module
``` 
creates `first-go-module.exe` in the `bin` subdirectory underneath `GOPATH`.