package main

import "C"

type EnglishMsger struct{}

func (EnglishMsger) Msg() string {
	return "hello, world! (from plugin)"
}

var Hello EnglishMsger
