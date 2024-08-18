// hello.go
package main

import (
	"log"
	"syscall/js"
)

func main() {
	log.Printf("está vivo!")
	js.Global().Set("hello", js.FuncOf(hello))
	select {}
}

func hello(this js.Value, p []js.Value) interface{} {
	return "Hello, " + p[0].String() + "!"
}
