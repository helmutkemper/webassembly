package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/globalEngine"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagSvgAnimateMotion
//
// English:
//
// The SVG <animateMotion> element provides a way to define how an element moves along a motion path.
//
//	Notes:
//	  * To reuse an existing path, it will be necessary to use an <mpath> element inside the <animateMotion> element
//	    instead of the path attribute.
//
// Português:
//
// O elemento SVG <animateMotion> fornece uma maneira de definir como um elemento se move ao longo de um caminho
// de movimento.
//
//	Notas:
//	  * Para reutilizar um caminho existente, será necessário usar um elemento <mpath> dentro do elemento
//	    <animateMotion> ao invés do atributo path.
func NewTagSvgAnimateMotion() (ref *html.TagSvgAnimateMotion) {
	ref = &html.TagSvgAnimateMotion{}
	ref.Engine(globalEngine.Engine) //todo: fazer em todos
	ref.Init()
	ref.Id(utilsMath.GetUID())

	return ref
}
