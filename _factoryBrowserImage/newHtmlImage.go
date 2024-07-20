package factoryBrowserImage

import (
	"github.com/helmutkemper/webassembly/browser/html"
)

// NewTagImg
//
// English:
//
// The <img> HTML element embeds an image into the document.
//
// Português:
//
// O elemento HTML <img> incorpora uma imagem no documento.
func NewTagImg() (ref *html.TagImg) {
	ref = new(html.TagImg)
	ref.Init()
	return
}
