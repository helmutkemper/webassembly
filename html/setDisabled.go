package html

import "log"

// SetDisabled
//
// English:
//
//  This Boolean attribute prevents the user from interacting with the button: it cannot be pressed
//  or focused.
//
// Firefox, unlike other browsers, persists the dynamic disabled state of a <button> across page
// loads.
//
// Português:
//
//  Este atributo booleano impede que o usuário interaja com o botão: ele não pode ser pressionado
//  ou focalizado.
//
// O Firefox, ao contrário de outros navegadores, mantém o estado dinâmico desabilitado de um
// <button> entre carregamentos de página.
func (e *GlobalAttributes) SetDisabled(disabled bool) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagButton:
	default:
		log.Printf("tag " + e.tag.String() + " does not support disabled property")
	}

	e.selfElement.Set("disabled", disabled)
	return e
}
