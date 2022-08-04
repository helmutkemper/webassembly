package media

// ConstrainULong
//
// English:
//
// The ConstrainULong constraint type is used to specify a constraint for a property whose value is an integer.
//
// Português:
//
// O tipo de restrição ConstrainULong é usado para especificar uma restrição para uma propriedade cujo valor é um
// inteiro.
type ConstrainULong struct {

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
	// An integer specifying the largest permissible value of the property it describes.
	// If the value cannot remain equal to or less than this value, matching will fail.
	//
	// Português:
	//
	// Um inteiro que especifica o maior valor permitido da propriedade que descreve.
	// Se o valor não puder permanecer igual ou menor que esse valor, a correspondência falhará.
	Max interface{} `js:"max"`

	// Min
	//
	// English:
	//
	// An integer specifying the smallest permissible value of the property it describes.
	// If the value cannot remain equal to or greater than this value, matching will fail.
	//
	// Português:
	//
	// Um inteiro que especifica o menor valor permitido da propriedade que descreve.
	// Se o valor não puder permanecer igual ou maior que esse valor, a correspondência falhará.
	Min interface{} `js:"min"`

	// Exact
	//
	// English:
	//
	// An integer specifying a specific, required, value the property must have to be considered acceptable.
	//
	// Português:
	//
	// Um número inteiro que especifica um valor específico e obrigatório que a propriedade deve ter para ser
	// considerada aceitável.
	Exact interface{} `js:"exact"`

	// Ideal
	//
	// English:
	//
	// An integer specifying an ideal value for the property.
	// If possible, this value will be used, but if it's not possible, the user agent will use the closest possible match.
	//
	// Português:
	//
	// Um inteiro que especifica um valor ideal para a propriedade.
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
func (e *ConstrainULong) SetValue(value int) {
	e.Value = value
}

// SetMax
//
// English:
//
// An integer specifying the largest permissible value of the property it describes.
// If the value cannot remain equal to or less than this value, matching will fail.
//
// Português:
//
// Um inteiro que especifica o maior valor permitido da propriedade que descreve.
// Se o valor não puder permanecer igual ou menor que esse valor, a correspondência falhará.
func (e *ConstrainULong) SetMax(value int) {
	e.Max = value
}

// SetMin
//
// English:
//
// An integer specifying the smallest permissible value of the property it describes.
// If the value cannot remain equal to or greater than this value, matching will fail.
//
// Português:
//
// Um inteiro que especifica o menor valor permitido da propriedade que descreve.
// Se o valor não puder permanecer igual ou maior que esse valor, a correspondência falhará.
func (e *ConstrainULong) SetMin(value int) {
	e.Min = value
}

// SetExact
//
// English:
//
// An integer specifying a specific, required, value the property must have to be considered acceptable.
//
// Português:
//
// Um número inteiro que especifica um valor específico e obrigatório que a propriedade deve ter para ser
// considerada aceitável.
func (e *ConstrainULong) SetExact(value int) {
	e.Exact = value
}

// SetIdeal
//
// English:
//
// An integer specifying an ideal value for the property.
// If possible, this value will be used, but if it's not possible, the user agent will use the closest possible match.
//
// Português:
//
// Um inteiro que especifica um valor ideal para a propriedade.
// Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a correspondência mais
// próxima possível.
func (e *ConstrainULong) SetIdeal(value int) {
	e.Ideal = value
}
