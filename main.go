package main

import "fmt"
import "github.com/jmh-git/first-go-module/hello"

func main() {
	fmt.Println("first-go-module/main.go :: main()")
	fmt.Println("(executable module)")

	hello.Hello()
}