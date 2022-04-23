package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

// NewTagInputTel
//
// English:
//
//  Create the element input.
//
// The <input> HTML element is used to create interactive controls for web-based forms in order to
// accept data from the user; a wide variety of types of input data and control widgets are
// available, depending on the device and user agent.
//
// The <input> element is one of the most powerful and complex in all of HTML due to the sheer
// number of combinations of input types and attributes.
//
// Português
//
//  Cria o elemento input.
//
// O elemento HTML <input> é usado para criar controles interativos para formulários baseados na
// web para aceitar dados do usuário; uma ampla variedade de tipos de dados de entrada e widgets
// de controle estão disponíveis, dependendo do dispositivo e do agente do usuário.
//
// O elemento <input> é um dos mais poderosos e complexos dentro do HTML, devido ao grande número
// de combinações de tipos de entrada e atributos.
func NewTagInputTel(id string) (ref *html.TagInputTel) {
	ref = &html.TagInputTel{}
	ref.CreateElement(html.KTagInput)
	ref.Type(html.KInputTypeTel)
	ref.Id(id)

	return ref
}