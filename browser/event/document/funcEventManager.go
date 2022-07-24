package document

import (
	"syscall/js"
)

type EventName string

func (e EventName) String() string {
	return string(e)
}

const (
	KEventResize EventName = "resize"
)

// EventManager
//
// English:
//
// Capture event information and format to Golang
//
//   Output:
//     data: list with all the information provided by the browser.
//
// Português:
//
// Captura as informações do evento e formata para o Golang
//
//   Saída:
//     data: lista com todas as informações fornecidas pelo navegador.
func EventManager(name EventName, this js.Value, args []js.Value) (data Data) {
	var event = Event{}
	event.Object = this

	data.Width = event.GetWidth()
	data.Height = event.GetHeight()
	data.Name = event.GetName()
	data.Length = event.GetFrameLength()
	data.Closed = event.GetClosed()
	data.OuterHeight = event.GetOuterHeight()
	data.OuterWidth = event.GetOuterWidth()
	data.ScrollX = event.GetScrollX()
	data.ScrollY = event.GetScrollY()
	data.ScreenX = event.GetScreenX()
	data.ScreenY = event.GetScreenY()
	data.Opener = event.GetOpener()
	data.Parent = event.GetParent()
	data.Screen = event.GetScreen()
	data.ScrollBars = event.GetScrollBars()
	data.StatusBar = event.GetStatusBar()
	data.Top = event.GetTop()
	data.This = this

	return
}
