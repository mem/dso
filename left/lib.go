package left

import "github.com/mem/dso/right"

type HelloWorld struct{}

func (HelloWorld) Msg() string {
	return right.Msg()
}
