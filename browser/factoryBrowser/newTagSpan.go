package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
)

// NewTagSpan
//
// English:
//
//	The <span> HTML element is a generic inline container for phrasing content, which does not inherently represent anything.
//
//	It can be used to group elements for styling purposes (using the class or id attributes), or because they share attribute values, such as lang. It should be used only when no other semantic element is appropriate. <span> is very much like a <div> element, but <div> is a block-level element whereas a <span> is an inline-level element.
//
// Português:
//
//	O elemento HTML <span> é um contêiner embutido genérico para frasear conteúdo, que não representa nada inerentemente.
//
//	Ele pode ser usado para agrupar elementos para fins de estilo (usando os atributos class ou id) ou porque eles compartilham valores de atributos, como lang. Deve ser usado somente quando nenhum outro elemento semântico for apropriado. <span> é muito parecido com um elemento <div>, mas <div> é um elemento de nível de bloco, enquanto <span> é um elemento de nível embutido.
func NewTagSpan() (ref *html.TagSpan) {
	ref = new(html.TagSpan)
	ref.Init()
	ref.Id(mathUtil.GetUID())

	return ref
}

func NewTagSpanWithDelta(deltaX, deltaY int) (ref *html.TagSpan) {
	ref = &html.TagSpan{}
	ref.Init().
		SetDeltaX(deltaX).
		SetDeltaY(deltaY).
		Id(mathUtil.GetUID())

	return ref
}
