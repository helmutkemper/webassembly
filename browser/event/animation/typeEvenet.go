// Package animation
//
// English:
//
// This package returns the variables available for animation, however, they are experimental and have an error in
// 07/2022. For this reason, only the getCurrentTime() function was used. All others are unstable.
//
// Português:
//
// Este pacore retorna as variáveis disponíveis para animação, porém, as mesmas são experimentais e apresentam erro em
// 07/2022. Por isto, apenas a função getCurrentTime() foi usada. Todas as outras são instáveis.
//
// Sorce: / Fonte:
// https://developer.mozilla.org/en-US/docs/Web/API/SVGAnimationElement
package animation

import (
	"syscall/js"
)

type Event struct {
	Object js.Value
}

// GetCurrentTime
//
// English:
//
// Returns a float representing the current time in seconds relative to time zero for the given time container.
//
// Português:
//
// Retorna um float representando o tempo atual em segundos em relação ao tempo zero para o contêiner de tempo
// fornecido.
func (e Event) GetCurrentTime() (currentTime float64) {
	return e.Object.Call("getCurrentTime").Float()
}

func (e Event) GetClientX() (clientX float64) {
	clientX = e.Object.Get("clientX").Float()
	return
}

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
func EventManager(this js.Value) (data Data) {
	var event = Event{}
	event.Object = this

	data.CurrentTime = event.GetCurrentTime()
	data.This = this
	return
}

type Data struct {

	// CurrentTime
	//
	// English:
	//
	// Returns a float representing the current time in seconds relative to time zero for the given time container.
	//
	// Português:
	//
	// Retorna um float representando o tempo atual em segundos em relação ao tempo zero para o contêiner de tempo
	// fornecido.
	CurrentTime float64

	// This
	//
	// English:
	//
	// This is the equivalent property of JavaScript's 'this'.
	//
	// The way to use it is This.Get(property string name). E.g. chan.This.Get("id")
	//
	// Português:
	//
	// Esta é a propriedade equivalente ao 'this' do JavaScript.
	//
	// A forma de usar é This.Get(property string name). Ex. chan.This.Get("id")
	This js.Value
}
