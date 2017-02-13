package main

import (
	"fmt"
	"os"
	"plugin"
)

type Msger interface {
	Msg() string
}

func pluginDemo(fn string) {
	p, err := plugin.Open(fn)
	if err != nil {
		fmt.Printf("plugin.Open: %s\n", err)
		return
	}

	h, err := p.Lookup("Hello")
	if err != nil {
		fmt.Printf("p.Lookup: %s\n", err)
		return
	}

	m, ok := h.(Msger)
	if !ok {
		fmt.Println("E: Expecting Msger interface, but got something else.")
		return
	}

	fmt.Printf("%s: %s\n", fn, m.Msg())
}

func main() {
	if len(os.Args) >= 2 {
		for _, fn := range os.Args[1:] {
			pluginDemo(fn)
		}
	}
}
