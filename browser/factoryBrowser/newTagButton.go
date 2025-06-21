package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/utilsMath"
)

// NewTagButton
//
// English:
//
//	Create a new tag html button.
//
// The <button> HTML element is an interactive element activated by a user with a mouse, keyboard,
// finger, voice command, or other assistive technology. Once activated, it then performs a
// programmable action, such as submitting a form or opening a dialog.
//
// By default, HTML buttons are presented in a style resembling the platform the user agent runs on,
// but you can change buttons' appearance with CSS.
//
// Português:
//
//	Cria uma nova tag html button
//
// O elemento HTML <button> é um elemento interativo ativado por um usuário com mouse, teclado,
// dedo, comando de voz ou outra tecnologia assistiva. Uma vez ativado, ele executa uma ação
// programável, como enviar um formulário ou abrir uma caixa de diálogo.
//
// Por padrão, os botões HTML são apresentados em um estilo semelhante à plataforma na qual o agente
// do usuário é executado, mas você pode alterar a aparência dos botões com CSS.
func NewTagButton() (ref *html.TagButton) {
	ref = &html.TagButton{}
	ref.CreateElement(html.KTagButton)
	ref.Id(utilsMath.GetUID())

	return ref
}
