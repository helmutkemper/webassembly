package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagMeter
//
// English:
//
//	Create the Meter element.
//
//	The <meter> HTML element represents either a scalar value within a known range or a fractional
//	value.
//
// Português:
//
//	Crie o elemento Medidor.
//
//	O elemento HTML <meter> representa um valor escalar dentro de um intervalo conhecido ou um
//	valor fracionário.
func NewTagMeter() (ref *html.TagMeter) {
	ref = &html.TagMeter{}
	ref.CreateElement(html.KTagMeter)
	ref.Id(utilsMath.GetUID())

	return ref
}
