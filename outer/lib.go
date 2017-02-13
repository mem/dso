package outer

import "github.com/mem/dso/outer/inner"

type HelloWorld struct{}

func (HelloWorld) Msg() string {
	return inner.Msg()
}
