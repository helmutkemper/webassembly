package keyboard

import (
	"syscall/js"
)

type EventName string

func (e EventName) String() string {
	return string(e)
}

const (
	KEventKeyDown EventName = "keydown"
	KEventKeyUp   EventName = "keyup"
)

// EventManager
//
// English:
//
// Capture event information and format to Golang
//
//	Output:
//	  data: list with all the information provided by the browser.
//
// Português:
//
// Captura as informações do evento e formata para o Golang
//
//	Saída:
//	  data: lista com todas as informações fornecidas pelo navegador.
func EventManager(eventName EventName, this js.Value, args []js.Value) (data Data) {
	var event = Event{}
	event.Object = args[0]

	data.EventName = eventName
	data.AltKey = event.GetAltKey()
	data.Code = event.GetCode()
	data.CtrlKey = event.GetCtrlKey()
	data.IsComposing = event.GetIsComposing()
	data.Key = event.GetKey()
	data.Location = event.GetLocation()
	data.MetaKey = event.GetMetaKey()
	data.Repeat = event.GetRepeat()
	data.ShiftKey = event.GetShiftKey()

	data.This = this
	return
}
