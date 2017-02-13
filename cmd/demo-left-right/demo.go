package main

import (
	"fmt"

	"github.com/mem/dso/left"
)

func libDemo() {
	var l left.HelloWorld

	fmt.Printf("%s\n", l.Msg())
}

func main() {
	libDemo()
}
