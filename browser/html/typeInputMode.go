package html

type InputMode string

func (e InputMode) String() string {
	return string(e)
}

const (
	// KInputModeNone
	//
	// English:
	//
	//  No virtual keyboard. For when the page implements its own keyboard input control.
	//
	// Português:
	//
	//  Sem teclado virtual. Para quando a página implementa seu próprio controle de entrada de teclado.
	KInputModeNone InputMode = "none"

	// KInputModeText
	//
	// English:
	//
	//   Standard input keyboard for the user's current locale. (default)
	//
	// Português:
	//
	//  Teclado de entrada padrão para a localidade atual do usuário. (Padrão)
	KInputModeText InputMode = "text"

	// KInputModeDecimal
	//
	// English:
	//
	//  Fractional numeric input keyboard containing the digits and decimal separator for the user's
	//  locale (typically . or ,). Devices may or may not show a minus key (-).
	//
	// Português:
	//
	//  Teclado de entrada numérica fracionária contendo os dígitos e o separador decimal para a
	//  localidade do usuário (normalmente . ou ,). Os dispositivos podem ou não mostrar uma tecla
	//  de menos (-).
	KInputModeDecimal InputMode = "decimal"

	// KInputModeNumeric
	//
	// English:
	//
	//  Numeric input keyboard, but only requires the digits 0–9. Devices may or may not show a minus
	//  key.
	//
	// Português:
	//
	//  Teclado de entrada numérica, mas requer apenas os dígitos de 0 a 9. Os dispositivos podem ou
	//  não mostrar uma tecla de menos.
	KInputModeNumeric InputMode = "numeric"

	// KInputModeTel
	//
	// English:
	//
	//  A telephone keypad input, including the digits 0–9, the asterisk (*), and the pound (#) key.
	//  Inputs that *require* a telephone number should typically use <input type="tel"> instead.
	//
	// Português:
	//
	//  Uma entrada do teclado do telefone, incluindo os dígitos de 0 a 9, o asterisco (*) e a tecla
	//  sustenido (#). As entradas que exigem um número de telefone normalmente devem usar
	//  <input type="tel">.
	KInputModeTel InputMode = "tel"

	// KInputModeSearch
	//
	// English:
	//
	//  A virtual keyboard optimized for search input. For instance, the return/submit key may be
	//  labeled "Search", along with possible other optimizations. Inputs that require a search query
	//  should typically use <input type="search"> instead.
	//
	// Português:
	//
	//  Um teclado virtual otimizado para entrada de pesquisa. Por exemplo, a chave returnsubmit pode
	//  ser rotulada como "Search", juntamente com possíveis outras otimizações. As entradas que exigem
	//  uma consulta de pesquisa normalmente devem usar <input type="search">.
	KInputModeSearch InputMode = "search"

	// KInputModeEmail
	//
	// English:
	//
	//  A virtual keyboard optimized for entering email addresses. Typically includes the @character
	//  as well as other optimizations. Inputs that require email addresses should typically use
	//  <input type="email"> instead.
	//
	// Português:
	//
	//  Um teclado virtual otimizado para inserir endereços de e-mail. Normalmente inclui o @character,
	//  bem como outras otimizações. As entradas que exigem endereços de e-mail normalmente devem usar
	//  <input type="email">.
	KInputModeEmail InputMode = "email"

	// KInputModeUrl
	//
	// English:
	//
	//  A keypad optimized for entering URLs. This may have the / key more prominent, for example.
	//  Enhanced features could include history access and so on. Inputs that require a URL should
	//  typically use <input type="url"> instead.
	//
	// Português:
	//
	//  Um teclado otimizado para inserir URLs. Isso pode ter a chave mais proeminente, por exemplo.
	//  Recursos aprimorados podem incluir acesso ao histórico e assim por diante. As entradas que
	//  exigem um URL normalmente devem usar <input type="url">.
	KInputModeUrl InputMode = "url"
)
