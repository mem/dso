package main

import (
	"fmt"

	"github.com/mem/dso/outer"
)

func libDemo() {
	var l outer.HelloWorld

	fmt.Printf("%s\n", l.Msg())
}

func main() {
	libDemo()
}
