package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgLinearGradient
//
// English:
//
// The <defs> element is used to store graphical objects that will be used at a later time.
//
// Objects created inside a <defs> element are not rendered directly. To display them you have to reference them
// (with a <use> element for example).
//
// Graphical objects can be referenced from anywhere, however, defining these objects inside of a <defs> element
// promotes understandability of the SVG content and is beneficial to the overall accessibility of the document.
//
// Português:
//
// O elemento <defs> é usado para armazenar objetos gráficos que serão usados posteriormente.
//
// Objetos criados dentro de um elemento <defs> não são renderizados diretamente. Para exibi-los, você deve
// referenciá-los (com um elemento <use>, por exemplo).
//
// Graphical objects can be referenced from anywhere, however, defining these objects inside of a <defs> element
// promotes understandability of the SVG content and is beneficial to the overall accessibility of the document.
func NewTagSvgLinearGradient() (ref *html.TagSvgLinearGradient) {
	ref = &html.TagSvgLinearGradient{}
	ref.Init()

	return ref
}
