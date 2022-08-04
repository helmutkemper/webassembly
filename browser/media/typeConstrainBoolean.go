package media

// ConstrainBoolean
//
// English:
//
// The ConstrainBoolean constraint type is used to specify a constraint for a property whose value is a Boolean value.
//
// Português:
//
// O tipo de restrição ConstrainBoolean é usado para especificar uma restrição para uma propriedade cujo valor é um
// valor booleano.
type ConstrainBoolean struct {
	// Value
	//
	// English:
	//
	// Property value, when Exact or Ideal are not required.
	//
	// Português:
	//
	// Valor da propriedade, quando Exact ou Ideal não são necessários.
	Value BOOLEAN

	// Exact
	//
	// English:
	//
	// A Boolean which must be the value of the property. If the property can't be set to this value, matching will fail.
	//
	// Português:
	//
	// Um booleano que deve ser o valor da propriedade. Se a propriedade não puder ser definida com esse valor, a
	// correspondência falhará.
	Exact BOOLEAN `js:"exact"`

	// Ideal
	//
	// English:
	//
	// A Boolean specifying an ideal value for the property. If possible, this value will be used, but if it's not
	// possible, the user agent will use the closest possible match.
	//
	// Português:
	//
	// Um booleano que especifica um valor ideal para a propriedade. Se possível, esse valor será usado, mas se não for
	// possível, o agente do usuário usará a correspondência mais próxima possível.
	Ideal BOOLEAN `js:"ideal"`
}

// SetValue
//
// English:
//
// Property value, when Exact or Ideal are not required.
//
// Português:
//
// Valor da propriedade, quando Exact ou Ideal não são necessários.
func (e *ConstrainBoolean) SetValue(value bool) {
	if value == true {
		e.Value = 1
	} else {
		e.Value = -1
	}
}

// SetExact
//
// English:
//
// A Boolean which must be the value of the property. If the property can't be set to this value, matching will fail.
//
// Português:
//
// Um booleano que deve ser o valor da propriedade. Se a propriedade não puder ser definida com esse valor, a
// correspondência falhará.
func (e *ConstrainBoolean) SetExact(value bool) {
	if value == true {
		e.Exact = 1
	} else {
		e.Exact = -1
	}
}

// SetIdeal
//
// English:
//
// A Boolean specifying an ideal value for the property. If possible, this value will be used, but if it's not
// possible, the user agent will use the closest possible match.
//
// Português:
//
// Um booleano que especifica um valor ideal para a propriedade. Se possível, esse valor será usado, mas se não for
// possível, o agente do usuário usará a correspondência mais próxima possível.
func (e *ConstrainBoolean) SetIdeal(value bool) {
	if value == true {
		e.Ideal = 1
	} else {
		e.Ideal = -1
	}
}
