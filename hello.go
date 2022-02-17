package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(greeting("pierre"))
	showTime()
}

func showTime() {
	fmt.Println("the time is", time.Now())
}

func greeting(name string) string {
	return "hello " + name
}
