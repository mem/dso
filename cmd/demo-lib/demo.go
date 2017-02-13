package main

import (
	"fmt"

	"github.com/mem/dso/lib"
)

func libDemo() {
	var l lib.HelloWorld

	fmt.Printf("%s\n", l.Msg())
}

func main() {
	libDemo()
}
