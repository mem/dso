package lib

type HelloWorld struct{}

func (HelloWorld) Msg() string {
	return "hello, world! (from lib)"
}
