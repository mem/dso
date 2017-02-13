package main

import "C"

type SpanishMsger struct{}

func (SpanishMsger) Msg() string {
	return "¡hola, mundo! (desde un plugin)"
}

var Hello SpanishMsger
