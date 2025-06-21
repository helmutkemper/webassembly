package utilsWindow

import (
	"sync"
	"syscall/js"
)

var onceInjectWindowBody sync.Once

// InjectBodyNoMargin
//
// English:
//
//	Ensures the global body element has no margin, no padding, and disables touch-action using CSS styles.
//
//	It uses a sync.Once to guarantee that the injection occurs only once during the program's lifetime.
//
//	The function modifies the DOM by appending a style element containing the defined CSS rules to the document head.
//
// Português:
//
//	Garante que o elemento body não tenha margem, nem preenchimento, e desativa touch-action automático usando estilos CSS.
//
//	Ele usa um sync.Once para garantir que à injeção ocorra apenas uma vez durante a vida útil do programa.
//
//	A função modifica o DOM, anexando um elemento de estilo que contém as regras CSS definidas no document head.
func InjectBodyNoMargin() {
	onceInjectWindowBody.Do(func() {
		document := js.Global().Get("document")

		css := `
			html, body {
			  margin: 0;
			  padding: 0;
			  width: 100vw;
			  height: 100vh;
			  overflow: hidden;
			  touch-action: none;
			}
		`

		styleEl := document.Call("createElement", "style")
		styleEl.Set("type", "text/css")
		styleEl.Set("textContent", css)

		document.Get("head").Call("appendChild", styleEl)

		//	js.Global().Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) any {
		//		width := js.Global().Get("innerWidth").Int()
		//		height := js.Global().Get("innerHeight").Int()
		//		svg := js.Global().Get("document").Call("getElementById", "mysvg")
		//		svg.Call("setAttribute", "width", fmt.Sprintf("%d", width))
		//		svg.Call("setAttribute", "height", fmt.Sprintf("%d", height))
		//		return nil
		//	}))
	})
}
