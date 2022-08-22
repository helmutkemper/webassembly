package keyboard

import (
	"syscall/js"
)

type Data struct {
	// This
	//
	// English:
	//
	// This is the equivalent property of JavaScript's 'this'.
	//
	// The way to use it is This.Get(property string name). E.g. chan.This.Get("id")
	//
	// Português:
	//
	// Esta é a propriedade equivalente ao 'this' do JavaScript.
	//
	// A forma de usar é This.Get(property string name). Ex. chan.This.Get("id")
	This js.Value

	// AltKey
	//
	// English:
	//
	// Returns a boolean value that is true if the Alt (Option or ⌥ on macOS) key was active when the key event was
	// generated.
	//
	// Português:
	//
	// Retorna um valor booleano que é verdadeiro se a tecla Alt (Option ou ⌥ no macOS) estiver ativa quando o evento de
	// chave foi gerado.
	AltKey bool

	// Code
	//
	// English:
	//
	// Returns a string with the code value of the physical key represented by the event.
	//
	//  Warning:
	//    This ignores the user's keyboard layout, so that if the user presses the key at the "Y" position in a QWERTY
	//    keyboard layout (near the middle of the row above the home row), this will always return "KeyY", even if the
	//    user has a QWERTZ keyboard (which would mean the user expects a "Z" and all the other properties would indicate
	//    a "Z") or a Dvorak keyboard layout (where the user would expect an "F"). If you want to display the correct
	//    keystrokes to the user, you can use Keyboard.getLayoutMap().
	//
	// Português:
	//
	// Retorna uma string com o valor do código da chave física representada pelo evento.
	//
	//  Cuidado:
	//    Ele ignora o layout do teclado do usuário, de modo que, se o usuário pressionar a tecla na posição "Y" em um
	//    layout de teclado QWERTY (perto do meio da linha acima da linha inicial), isso sempre retornará "KeyY", mesmo
	//    se o usuário tiver um teclado QWERTZ (o que significa que o usuário espera um "Z" e todas as outras propriedades
	//    indicariam um "Z") ou um layout de teclado Dvorak (onde o usuário espera um "F"). Se você deseja exibir as
	//    teclas corretas para o usuário, você pode usar Keyboard.getLayoutMap().
	Code string

	// CtrlKey
	//
	// English:
	//
	// Returns a boolean value that is true if the Ctrl key was active when the key event was generated.
	//
	// Português:
	//
	// Retorna um valor booleano que é true se a tecla Ctrl estava ativa quando o evento de chave foi gerado.
	CtrlKey bool

	// IsComposing
	//
	// English:
	//
	// Returns a boolean value that is true if the event is fired between after compositionstart and before
	// compositionend.
	//
	// Português:
	//
	// Retorna um valor booleano que é verdadeiro se o evento for acionado entre após o início da composição e antes da
	// conclusão da composição.
	IsComposing bool

	// Key
	//
	// English:
	//
	// Returns a string representing the key value of the key represented by the event.
	//
	// Português:
	//
	// Retorna uma string representando o valor da chave representada pelo evento.
	Key string

	// Location
	//
	// English:
	//
	// Returns a number representing the location of the key on the keyboard or other input device.
	//
	// Português:
	//
	// Retorna um número que representa a localização da tecla no teclado ou outro dispositivo de entrada.
	Location Location

	// MetaKey
	//
	// English:
	//
	// Returns a boolean value that is true if the Meta key (on Mac keyboards, the ⌘ Command key; on Windows keyboards,
	// the Windows key (⊞)) was active when the key event was generated.
	//
	// Português:
	//
	// Returns a boolean value that is true if the Meta key (on Mac keyboards, the ⌘ Command key; on Windows keyboards,
	// the Windows key (⊞)) was active when the key event was generated.
	MetaKey bool

	// Repeat
	//
	// English:
	//
	// Returns a boolean value that is true if the key is being held down such that it is automatically repeating.
	//
	// Português:
	//
	// Retorna um valor booleano que é true se a chave estiver sendo mantida pressionada de forma que ela se repita
	// automaticamente.
	Repeat bool

	// ShiftKey
	//
	// English:
	//
	// Returns a boolean value that is true if the Shift key was active when the key event was generated.
	//
	// Português:
	//
	// Retorna um valor booleano que é true se a tecla Shift estava ativa quando o evento de chave foi gerado.
	ShiftKey bool

	Object js.Value

	// EventName
	//
	// English:
	//
	// Name of event
	//
	// Português:
	//
	// Nome do evento
	EventName EventName
}

type Event struct {
	Object js.Value
}

// GetAltKey
//
// English:
//
// Returns a boolean value that is true if the Alt (Option or ⌥ on macOS) key was active when the key event was
// generated.
//
// Português:
//
// Retorna um valor booleano que é verdadeiro se a tecla Alt (Option ou ⌥ no macOS) estiver ativa quando o evento de
// chave foi gerado.
func (e Event) GetAltKey() (AltKey bool) {
	return e.Object.Get("altKey").Bool()
}

// GetCode
//
// English:
//
// Returns a string with the code value of the physical key represented by the event.
//
//	Warning:
//	  This ignores the user's keyboard layout, so that if the user presses the key at the "Y" position in a QWERTY
//	  keyboard layout (near the middle of the row above the home row), this will always return "KeyY", even if the
//	  user has a QWERTZ keyboard (which would mean the user expects a "Z" and all the other properties would indicate
//	  a "Z") or a Dvorak keyboard layout (where the user would expect an "F"). If you want to display the correct
//	  keystrokes to the user, you can use Keyboard.getLayoutMap().
//
// Português:
//
// Retorna uma string com o valor do código da chave física representada pelo evento.
//
//	Cuidado:
//	  Ele ignora o layout do teclado do usuário, de modo que, se o usuário pressionar a tecla na posição "Y" em um
//	  layout de teclado QWERTY (perto do meio da linha acima da linha inicial), isso sempre retornará "KeyY", mesmo
//	  se o usuário tiver um teclado QWERTZ (o que significa que o usuário espera um "Z" e todas as outras propriedades
//	  indicariam um "Z") ou um layout de teclado Dvorak (onde o usuário espera um "F"). Se você deseja exibir as
//	  teclas corretas para o usuário, você pode usar Keyboard.getLayoutMap().
func (e Event) GetCode() (code string) {
	return e.Object.Get("code").String()
}

// GetCtrlKey
//
// English:
//
// Returns a boolean value that is true if the Ctrl key was active when the key event was generated.
//
// Português:
//
// Retorna um valor booleano que é true se a tecla Ctrl estava ativa quando o evento de chave foi gerado.
func (e Event) GetCtrlKey() (ctrlKey bool) {
	return e.Object.Get("ctrlKey").Bool()
}

// GetIsComposing
//
// English:
//
// Returns a boolean value that is true if the event is fired between after compositionstart and before
// compositionend.
//
// Português:
//
// Retorna um valor booleano que é verdadeiro se o evento for acionado entre após o início da composição e antes da
// conclusão da composição.
func (e Event) GetIsComposing() (isComposing bool) {
	return e.Object.Get("isComposing").Bool()
}

// GetKey
//
// English:
//
// Returns a string representing the key value of the key represented by the event.
//
// Português:
//
// Retorna uma string representando o valor da chave representada pelo evento.
func (e Event) GetKey() (key string) {
	return e.Object.Get("key").String()
}

// GetLocation
//
// English:
//
// Returns a number representing the location of the key on the keyboard or other input device.
//
// Português:
//
// Retorna um número que representa a localização da tecla no teclado ou outro dispositivo de entrada.
func (e Event) GetLocation() (location Location) {
	return Location(e.Object.Get("location").Int())
}

// GetMetaKey
//
// English:
//
// Returns a boolean value that is true if the Meta key (on Mac keyboards, the ⌘ Command key; on Windows keyboards,
// the Windows key (⊞)) was active when the key event was generated.
//
// Português:
//
// Returns a boolean value that is true if the Meta key (on Mac keyboards, the ⌘ Command key; on Windows keyboards,
// the Windows key (⊞)) was active when the key event was generated.
func (e Event) GetMetaKey() (metaKey bool) {
	return e.Object.Get("metaKey").Bool()
}

// GetRepeat
//
// English:
//
// Returns a boolean value that is true if the key is being held down such that it is automatically repeating.
//
// Português:
//
// Retorna um valor booleano que é true se a chave estiver sendo mantida pressionada de forma que ela se repita
// automaticamente.
func (e Event) GetRepeat() (repeat bool) {
	return e.Object.Get("repeat").Bool()
}

// GetShiftKey
//
// English:
//
// Returns a boolean value that is true if the Shift key was active when the key event was generated.
//
// Português:
//
// Retorna um valor booleano que é true se a tecla Shift estava ativa quando o evento de chave foi gerado.
func (e Event) GetShiftKey() (shiftKey bool) {
	return e.Object.Get("shiftKey").Bool()
}
