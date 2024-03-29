package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagCanvas
//
// English:
//
// Use the HTML <canvas> element with either the canvas scripting API or the WebGL API to draw graphics and animations.
//
// Português:
//
// Use o elemento HTML <canvas> com a API de script de tela ou a API WebGL para desenhar gráficos e animações.
func NewTagCanvas(width, height int) (ref *html.TagCanvas) {
	ref = &html.TagCanvas{}
	ref.Init(width, height)

	return ref
}
