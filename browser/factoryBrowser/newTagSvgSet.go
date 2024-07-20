package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

// NewTagSvgSet
//
// English:
//
// The SVG <set> element provides a simple means of just setting the value of an attribute for a specified duration.
//
// It supports all attribute types, including those that cannot reasonably be interpolated, such as string and boolean
// values. For attributes that can be reasonably be interpolated, the <animate> is usually preferred.
//
//	Notes:
//	  * The <set> element is non-additive. The additive and accumulate attributes are not allowed, and will be
//	    ignored if specified.
//
// Português:
//
// O elemento SVG <set> fornece um meio simples de apenas definir o valor de um atributo para uma duração especificada.
//
// Ele suporta todos os tipos de atributos, incluindo aqueles que não podem ser interpolados de maneira razoável, como
// valores de string e booleanos. Para atributos que podem ser razoavelmente interpolados, o <animate> geralmente é
// preferido.
//
//	Notas:
//	  * O elemento <set> não é aditivo. Os atributos aditivo e acumular não são permitidos e serão ignorados se
//	    especificados.
func NewTagSvgSet() (ref *html.TagSvgSet) {
	ref = &html.TagSvgSet{}
	ref.Init()

	return ref
}
