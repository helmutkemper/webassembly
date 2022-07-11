package factoryBrowser

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

// NewTagFigCaption
//
// English:
//
// The <figcaption> HTML element represents a caption or legend describing the rest of the contents
// of its parent <figure> element.
//
// Português:
//
// O elemento HTML <figcaption> representa uma legenda ou legenda descrevendo o restante do conteúdo
// de seu elemento pai <figure>.
func NewTagFigCaption() (ref *html.TagFigCaption) {
	ref = &html.TagFigCaption{}
	ref.Init()

	return ref
}
