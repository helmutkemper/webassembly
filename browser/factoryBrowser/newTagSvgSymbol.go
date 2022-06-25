package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgSymbol
//
// English:
//
// The <symbol> element is used to define graphical template objects which can be instantiated by a <use> element.
//
// The use of <symbol> elements for graphics that are used multiple times in the same document adds structure and
// semantics.
//
// Documents that are rich in structure may be rendered graphically, as speech, or as Braille, and thus promote
// accessibility.
//
// Português:
//
// O elemento <symbol> é usado para definir objetos de template gráficos que podem ser instanciados por um elemento
// <use>.
//
// O uso de elementos <symbol> para gráficos que são usados várias vezes no mesmo documento adiciona estrutura e
// semântica.
//
// Documentos ricos em estrutura podem ser renderizados graficamente, como fala, ou como Braille, promovendo assim a
// acessibilidade.
func NewTagSvgSymbol(id string) (ref *html.TagSvgSymbol) {
	ref = &html.TagSvgSymbol{}
	ref.Init(id)

	return ref
}
