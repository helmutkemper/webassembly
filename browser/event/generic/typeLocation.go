package generic

// Location
//
// English:
//
// Represente the location of the key on the keyboard or other input device.
//
// Português:
//
// Representa a localização da tecla no teclado ou outro dispositivo de entrada.
type Location int

const (
	// LocationStandard
	//
	// English:
	//
	// The key has only one version, or can't be distinguished between the left and right versions of the key, and was not
	// pressed on the numeric keypad or a key that is considered to be part of the keypad.
	//
	// Português:
	//
	// A tecla tem apenas uma versão, ou não pode ser distinguida entre as versões esquerda e direita da tecla, e não foi
	// pressionada no teclado numérico ou em uma tecla que seja considerada parte do teclado.
	LocationStandard Location = 0

	// LocationLeft
	//
	// English:
	//
	// The key was the left-hand version of the key; for example, the left-hand Control key was pressed on a standard 101
	// key US keyboard. This value is only used for keys that have more than one possible location on the keyboard.
	//
	// Português:
	//
	// A chave era a versão esquerda da chave; por exemplo, a tecla Control do lado esquerdo foi pressionada em um teclado
	// americano padrão de 101 teclas. Este valor é usado apenas para teclas que possuem mais de uma localização possível
	// no teclado.
	LocationLeft Location = 1

	// LocationRight
	//
	// English:
	//
	// The key was the right-hand version of the key; for example, the right-hand Control key is pressed on a standard 101
	// key US keyboard. This value is only used for keys that have more than one possible location on the keyboard.
	//
	// Português:
	//
	// A chave era a versão direita da chave; por exemplo, a tecla Control do lado direito é pressionada em um teclado
	// americano padrão de 101 teclas. Este valor é usado apenas para teclas que possuem mais de uma localização possível
	// no teclado.
	LocationRight Location = 2

	// LocationNumPad
	//
	// English:
	//
	// The key was on the numeric keypad, or has a virtual key code that corresponds to the numeric keypad.
	//
	// Português:
	//
	// A tecla estava no teclado numérico ou possui um código de tecla virtual que corresponde ao teclado numérico.
	LocationNumPad Location = 3
)
