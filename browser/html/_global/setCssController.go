package _global

import "github.com/helmutkemper/webassembly/browser/css"

// SetCssController
//
// English:
//
//	Add the css classes to the created element.
//
//	 Input:
//	   value: object pointer to css.Class initialized
//
//	 Note:
//	   * This function is equivalent to css.SetList("current", classes...)
//	   * Css has a feature that allows you to easily change the list of css classes of an html tag,
//	     with the functions SetList(), CssToggle() and CssToggleTime();
//	   * Is the equivalent of <... css="name1 name2 nameN">
//
// Português:
//
//	Adiciona as classes css ao elemento criado.
//
//	 Entrada:
//	   classes: lista de classes css.
//
//	 Nota:
//	   * Esta função equivale a SetList("current", classes...);
//	   * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//	     fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//	   * Equivale a <... css="name1 name2 nameN">
func (e *GlobalAttributes) SetCssController(value *css.Class) (ref *GlobalAttributes) {
	e.cssClass = value
	e.cssClass.SetRef(e.id, &e.selfElement)

	return e
}
