package main

import (
	"fmt"
	"syscall/js"
)

func main() {

	document := js.Global().Get("document")

	createCustomContextMenu := func(this js.Value, args []js.Value) any {
		event := args[0]
		event.Call("preventDefault")

		pageX := event.Get("pageX").Float()
		pageY := event.Get("pageY").Float()

		contextMenu := document.Call("getElementById", "contextMenu")
		style := contextMenu.Get("style")
		style.Set("display", "block")
		style.Set("left", fmt.Sprintf("%vpx", pageX))
		style.Set("top", fmt.Sprintf("%vpx", pageY))
		return nil
	}
	document.Call("addEventListener", "contextmenu", js.FuncOf(createCustomContextMenu))

	// Hide custom context menu when clicking anywhere
	hideCustomContextMenu := func(this js.Value, args []js.Value) any {
		contextMenu := document.Call("getElementById", "contextMenu")
		style := contextMenu.Get("style")
		style.Set("display", "none")
		return nil
	}
	document.Call("addEventListener", "click", js.FuncOf(hideCustomContextMenu))

	option1Event := func(this js.Value, args []js.Value) any {
		//js.Global().Call("alert", "1")
		return nil
	}
	option1 := document.Call("getElementById", "option1")
	option1.Call("addEventListener", "click", js.FuncOf(option1Event))

	option2Event := func(this js.Value, args []js.Value) any {
		//js.Global().Call("alert", "2")
		return nil
	}
	option2 := document.Call("getElementById", "option2")
	option2.Call("addEventListener", "click", js.FuncOf(option2Event))

	option3Event := func(this js.Value, args []js.Value) any {
		//js.Global().Call("alert", "3")
		return nil
	}
	option3 := document.Call("getElementById", "option3")
	option3.Call("addEventListener", "click", js.FuncOf(option3Event))

	done := make(chan struct{})
	done <- struct{}{}
}
