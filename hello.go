package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println(greeting("pierre"))
}

func greeting(name string) string {
	return "hello " + name
}
