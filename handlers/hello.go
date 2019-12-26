package handlers

type helloHandler struct {
}

func (hh helloHandler) getMessage() interface{} {
	return "Привет!"
}

func NewHelloHandler() *helloHandler {
	return &helloHandler{}
}
