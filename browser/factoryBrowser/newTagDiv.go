package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagDiv
//
// English:
//
//	Creates a new html DIV element.
//
//	 Note:
//	   * Div Extends GlobalAttributes
//	   * By default, browsers always place a line break before and after the <div> element;
//	   * The <div> tag is used as a container for HTML elements - which is then styled with CSS or
//	     manipulated with JavaScript;
//	   * The <div> tag is easily styled by using the class or id attribute;
//	   * Any sort of content can be put inside the <div> tag.
//
//	The <div> tag defines a division or a section in an HTML document.
//
// Português:
//
//	Cria um novo elemento html DIV.
//
//	 Nota:
//	   * Div estende GlobalAttributes
//	   * Por padrão, os navegadores sempre colocam uma quebra de linha antes e depois do elemento
//	     <div>;
//	   * A tag <div> é usada como um contêiner para elementos HTML - que são estilizados com CSS ou
//	     manipulados com JavaScript
//	   * A tag <div> é facilmente estilizada usando o atributo class ou id;
//	   * Qualquer tipo de conteúdo pode ser colocado dentro da tag <div>.
//
//	A tag <div> define uma divisão ou uma seção em um documento HTML.
func NewTagDiv() (ref *html.TagDiv) {
	ref = new(html.TagDiv)
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}

func NewTagDivWithDelta(deltaX, deltaY int) (ref *html.TagDiv) {
	ref = &html.TagDiv{}
	ref.Init().
		SetDeltaX(deltaX).
		SetDeltaY(deltaY).
		Id(utilsMath.GetUID())

	return ref
}
