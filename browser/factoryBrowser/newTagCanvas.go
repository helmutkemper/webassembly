package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
)

// NewTagCanvas
//
// English:
//
// Use the HTML <canvas> element with either the canvas scripting API or the WebGL API to draw graphics and animations.
//
// Português:
//
// Use o elemento HTML <canvas> com a API de script de tela ou a API WebGL para desenhar gráficos e animações.
func NewTagCanvas(size ...int) (ref *html.TagCanvas) {
	width := 0
	height := 0

	if len(size) == 2 {
		width = size[0]
		height = size[1]
	}

	ref = &html.TagCanvas{}
	ref.Init(width, height)
	ref.Id(mathUtil.GetUID())

	return ref
}
