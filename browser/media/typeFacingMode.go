package media

type FacingMode struct {

	// Exact
	//
	// English:
	//
	// A string specifying a specific, required, value the property must have to be considered acceptable.
	//
	// Português:
	//
	// Uma string que especifica um valor específico e obrigatório que a propriedade deve ter para ser
	// considerada aceitável.
	Exact FacingModeValue //exact

	// Ideal
	//
	// English:
	//
	// A string specifying an ideal value for the property.
	// If possible, this value will be used, but if it's not possible, the user agent will use the closest possible match.
	//
	// Português:
	//
	// Uma string que especifica um valor ideal para a propriedade.
	// Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a correspondência mais
	// próxima possível.
	Ideal FacingModeValue //ideal
}

// mount
//
// English:
//
// Assemble a native golang object with the facing mode rules defined in:
// https://developer.mozilla.org/en-US/docs/Web/API/MediaTrackConstraints/facingMode
//
// Português:
//
// Monta um objeto nativo golang com as regras de facing mode definidas em:
// https://developer.mozilla.org/en-US/docs/Web/API/MediaTrackConstraints/facingMode
func (e *FacingMode) mount(facingMode *map[string]interface{}) {
	if *facingMode == nil {
		*facingMode = make(map[string]interface{})
	}

	if e.Exact != "" {
		(*facingMode)["exact"] = e.Exact.String()

		return
	}

	if e.Ideal != "" {
		(*facingMode)["ideal"] = e.Ideal.String()
	}

	return
}

// SetExact
//
// English:
//
// A string specifying a specific, required, value the property must have to be considered acceptable.
//
// Português:
//
// Uma string que especifica um valor específico e obrigatório que a propriedade deve ter para ser
// considerada aceitável.
func (e *FacingMode) SetExact(value FacingModeValue) {
	e.Exact = value
}

// SetIdeal
//
// English:
//
// A string specifying an ideal value for the property.
// If possible, this value will be used, but if it's not possible, the user agent will use the closest possible match.
//
// Português:
//
// Uma string que especifica um valor ideal para a propriedade.
// Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a correspondência mais
// próxima possível.
func (e *FacingMode) SetIdeal(value FacingModeValue) {
	e.Ideal = value
}
