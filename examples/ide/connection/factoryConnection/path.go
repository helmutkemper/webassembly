package factoryConnection

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/rulesConnection"
	"github.com/helmutkemper/webassembly/mathUtil"
)

// NewConnection
//
// English:
//
// The <path> SVG element is the generic element to define a shape. All the basic shapes can be created with a path
// element.
//
// Português:
//
// O elemento SVG <path> é o elemento genérico para definir uma forma. Todas as formas básicas podem ser criadas com
// um elemento de caminho.
func NewConnection(dataType, markEnd string) (ref *html.TagSvgPath) {
	ref = &html.TagSvgPath{}
	ref.Init()
	ref.Id(mathUtil.GetUID())
	ref.Fill(rulesConnection.TypeToColor(dataType))
	ref.Stroke("none")
	ref.MarkerEnd(markEnd)

	return ref
}
