package media

type ConstrainDouble struct {

	// Value
	//
	// English:
	//
	// Property value, when Exact or Ideal are not required.
	//
	// Português:
	//
	// Valor da propriedade, quando Exact ou Ideal não são necessários.
	Value interface{}

	// Max
	//
	// English:
	//
	// An decimal specifying the largest permissible value of the property it describes.
	// If the value cannot remain equal to or less than this value, matching will fail.
	//
	// Português:
	//
	// Um número decimal que especifica o maior valor permitido da propriedade que descreve.
	// Se o valor não puder permanecer igual ou menor que esse valor, a correspondência falhará.
	Max interface{} `js:"max"`

	// Min
	//
	// English:
	//
	// An decimal specifying the smallest permissible value of the property it describes.
	// If the value cannot remain equal to or greater than this value, matching will fail.
	//
	// Português:
	//
	// Um número decimal que especifica o menor valor permitido da propriedade que descreve.
	// Se o valor não puder permanecer igual ou maior que esse valor, a correspondência falhará.
	Min interface{} `js:"min"`

	// Exact
	//
	// English:
	//
	// An decimal specifying a specific, required, value the property must have to be considered acceptable.
	//
	// Português:
	//
	// Um número decimal que especifica um valor específico e obrigatório que a propriedade deve ter para ser
	// considerada aceitável.
	Exact interface{} `js:"exact"`

	// Ideal
	//
	// English:
	//
	// An decimal specifying an ideal value for the property.
	// If possible, this value will be used, but if it's not possible, the user agent will use the closest possible match.
	//
	// Português:
	//
	// Um número decimal que especifica um valor ideal para a propriedade.
	// Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a correspondência mais
	// próxima possível.
	Ideal interface{} `js:"ideal"`
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
func (e *ConstrainDouble) SetValue(value float64) {
	e.Value = value
}

// SetMax
//
// English:
//
// An decimal specifying the largest permissible value of the property it describes.
// If the value cannot remain equal to or less than this value, matching will fail.
//
// Português:
//
// Um número decimal que especifica o maior valor permitido da propriedade que descreve.
// Se o valor não puder permanecer igual ou menor que esse valor, a correspondência falhará.
func (e *ConstrainDouble) SetMax(value float64) {
	e.Max = value
}

// SetMin
//
// English:
//
// An decimal specifying the smallest permissible value of the property it describes.
// If the value cannot remain equal to or greater than this value, matching will fail.
//
// Português:
//
// Um número decimal que especifica o menor valor permitido da propriedade que descreve.
// Se o valor não puder permanecer igual ou maior que esse valor, a correspondência falhará.
func (e *ConstrainDouble) SetMin(value float64) {
	e.Min = value
}

// SetExact
//
// English:
//
// An decimal specifying a specific, required, value the property must have to be considered acceptable.
//
// Português:
//
// Um número decimal que especifica um valor específico e obrigatório que a propriedade deve ter para ser
// considerada aceitável.
func (e *ConstrainDouble) SetExact(value float64) {
	e.Exact = value
}

// SetIdeal
//
// English:
//
// An decimal specifying an ideal value for the property.
// If possible, this value will be used, but if it's not possible, the user agent will use the closest possible match.
//
// Português:
//
// Um número decimal que especifica um valor ideal para a propriedade.
// Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a correspondência mais
// próxima possível.
func (e *ConstrainDouble) SetIdeal(value float64) {
	e.Ideal = value
}
