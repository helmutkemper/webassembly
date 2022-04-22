package html

import "log"

// SetDirName
//
// English:
//
//  Valid for text and search input types only, the dirname attribute enables the submission of the
//  directionality of the element. When included, the form control will submit with two name/value
//  pairs: the first being the name and value, the second being the value of the dirname as the name
//  with the value of ltr or rtl being set by the browser.
//
// Português:
//
//  Válido apenas para tipos de entrada de texto e pesquisa, o atributo dirname permite o envio da
//  direcionalidade do elemento. Quando incluído, o controle de formulário será enviado com dois pares
//  nomevalor: o primeiro sendo o nome e o valor, o segundo sendo o valor do dirname como o nome com o
//  valor de ltr ou rtl sendo definido pelo navegador.
func (e *GlobalAttributes) SetDirName(dirname string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support dirname property")
	}

	e.selfElement.Set("dirname", dirname)
	return e
}
