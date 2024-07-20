package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgA
//
// English:
//
//	The <a> SVG element creates a hyperlink to other web pages, files, locations in the same page, email addresses, or
//	any other URL. It is very similar to HTML's <a> element.
//
// SVG's <a> element is a container, which means you can create a link around text (like in HTML) but also around any
// shape.
//
//	Notes:
//	  * Since this element shares its tag name with HTML's <a> element, selecting a with CSS or querySelector may apply
//	    to the wrong kind of element. Try the @namespace rule to distinguish the two.
//
// Português:
//
//	O elemento SVG <a> cria um hiperlink para outras páginas da web, arquivos, locais na mesma página, endereços de
//	e-mail ou qualquer outro URL. É muito semelhante ao elemento <a> do HTML.
//
// O elemento SVGs <a> é um contêiner, o que significa que você pode criar um link em torno do texto (como em HTML),
// mas também em torno de qualquer forma.
//
//	Notes:
//	  * Como esse elemento compartilha seu nome de tag com o elemento <a> do HTML, selecionar a com CSS ou
//	    querySelector pode se aplicar ao tipo errado de elemento. Experimente a regra @namespace para distinguir os
//	    dois.
func NewTagSvgA() (ref *html.TagSvgA) {
	ref = &html.TagSvgA{}
	ref.Init()

	return ref
}
