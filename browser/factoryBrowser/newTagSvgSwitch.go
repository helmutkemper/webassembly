package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgSwitch
//
// English:
//
// The <switch> SVG element evaluates any requiredFeatures, requiredExtensions and systemLanguage attributes on its
// direct child elements in order, and then renders the first child where these attributes evaluate to true.
//
// Other direct children will be bypassed and therefore not rendered. If a child element is a container element,
// like <g>, then its subtree is also processed/rendered or bypassed/not rendered.
//
//   Notes:
//     * The display and visibility properties have no effect on <switch> element processing.
//       In particular, setting display:none on a child has no effect on the true/false testing for <switch> processing.
//
// Português:
//
// O elemento SVG <switch> avalia todos os atributos requiredFeatures, requiredExtensions e systemLanguage em seus
// elementos filho diretos em ordem e, em seguida, renderiza o primeiro filho em que esses atributos são avaliados
// como true.
//
// Outros filhos diretos serão ignorados e, portanto, não renderizados. Se um elemento filho for um elemento contêiner,
// como <g>, sua subárvore também será processada, renderizada ou ignorada, não renderizada.
//
//   Notas:
//     * As propriedades de exibição e visibilidade não têm efeito no processamento do elemento <switch>.
//       Em particular, configurar display:none em um filho não tem efeito no teste truefalse para processamento de <switch>.
func NewTagSvgSwitch() (ref *html.TagSvgSwitch) {
	ref = &html.TagSvgSwitch{}
	ref.Init()

	return ref
}
