package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
)

// NewTagH5
//
// English:
//
// The <h1> to <h6> HTML elements represent six levels of section headings. <h1> is the highest
// section level and <h6> is the lowest.
//
// Multiple <h1> elements on one page
//
// Using more than one <h1> is allowed by the HTML specification, but is not considered a best
// practice. Using only one <h1> is beneficial for screenreader users.
//
// The HTML specification includes the concept of an outline formed by the use of <section> elements.
// If this were implemented it would enable the use of multiple <h1> elements, giving user
// agents—including screen readers—a way to understand that an <h1> nested inside a defined section is
// a subheading. This functionality has never been implemented; therefore it is important to use your
// headings to describe the outline of your document.
//
//	Notes:
//	  * Heading information can be used by user agents to construct a table of contents for a
//	    document automatically.
//	  * Avoid using heading elements to resize text. Instead, use the CSS font-size property.
//	  * Avoid skipping heading levels: always start from <h1>, followed by <h2> and so on.
//	  * Use only one <h1> per page or view. It should concisely describe the overall purpose of the
//	    content.
//	  * The align attribute is obsolete; don't use it.
//
// Português:
//
// Os elementos HTML <h1> a <h6> representam seis níveis de cabeçalho, onde, <h1> é o nível mais alto
// e <h6> o nível mais baixo.
//
// Múltiplos elementos <h1> em uma página
//
// O uso de mais de um <h1> é permitido pela especificação HTML, mas não é considerado uma prática
// recomendada. Usar apenas um <h1> é benéfico para usuários de leitores de tela.
//
// A especificação HTML inclui o conceito de contorno formado pelo uso de elementos <section>.
// Se isso fosse implementado, permitiria o uso de vários elementos <h1>, dando aos agentes do usuário
// – incluindo leitores de tela – uma maneira de entender que um <h1> aninhado dentro de uma seção
// definida é um subtítulo. Essa funcionalidade nunca foi implementada; portanto, é importante usar
// seus títulos para descrever o esboço do seu documento.
//
//	Notas:
//	  * As informações de cabeçalho podem ser usadas por agentes de usuário para construir
//	    automaticamente um índice para um documento.
//	  * Evite usar elementos de título para redimensionar o texto. Em vez disso, use a propriedade
//	    CSS font-size.
//	  * Evite pular níveis de título: sempre comece de <h1>, seguido de <h2> e assim por diante.
//	  * Use apenas um <h1> por página ou visualização. Deve descrever de forma concisa o propósito
//	    geral do conteúdo.
//	  * O atributo align está obsoleto; não o use.
func NewTagH5() (ref *html.TagH5) {
	ref = &html.TagH5{}
	ref.Init()

	return ref
}
