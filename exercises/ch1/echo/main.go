package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func test1() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args, " "), time.Since(start).Seconds())
}

func test2() {
	start := time.Now()
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
	fmt.Println(time.Since(start).Seconds())
}

func main() {
	test1()
	test2()
}
