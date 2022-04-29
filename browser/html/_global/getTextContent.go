package _global

// GetTextContent
//
// English:
//
//  The textContent property of the Node interface represents the text content of the node and its
//  descendants.
//
//   Note:
//     * SetTextContent() and SetInnerText() are easily confused, but the two properties are different
//       in important ways.
//     * Setting SetTextContent() on a node removes all of the node's children and replaces them with
//       a single text node with the given string value.
//
// Differences from SetInnerText()
//
// Don't get confused by the differences between GetTextContent() / SetTextContent() and
// GetInnerText() / SetInnerText(). Although the names seem similar, there are important differences:
//
// GetTextContent() / SetTextContent() gets the content of all elements, including <script> and
// <style> elements. In contrast, GetInnerText() / SetInnerText() only shows "human-readable"
// elements.
//
// GetTextContent() returns every element in the node. In contrast, innerText is aware of styling and
// won't return the text of "hidden" elements.
//
// Moreover, since GetInnerText() / SetInnerText() takes CSS styles into account, reading the value of
// innerText triggers a reflow to ensure up-to-date computed styles.
// (Reflows can be computationally expensive, and thus should be avoided when possible.)
//
// Both SetTextContent() and SetInnerText() remove child nodes when altered, but altering innerText in
// Internet Explorer (version 11 and below) also permanently destroys all descendant text nodes.
// It is impossible to insert the nodes again into any other element or into the same element after
// doing so.
//
// Differences from SetInnerHtml()
//
// GetInnerHtml() returns HTML, as its name indicates. Sometimes people use GetInnerHtml() /
// SetInnerHtml() to retrieve or write text inside an element, but GetTextContent() / SetTextContent()
// has better performance because its value is not parsed as HTML.
//
// Moreover, using GetTextContent() / SetTextContent() can prevent XSS attacks.
//
// Português:
//
//  A propriedade textContent da interface Node representa o conteúdo de texto do nó e seus
//  descendentes.
//
//   Nota:
//     * SetTextContent() e SetInnerText() são facilmente confundidos, mas as duas propriedades são
//       diferentes em aspectos importantes;
//     * Definir SetTextContent() em um nó remove todos os filhos do nó e os substitui por um único nó
//       de texto com o valor de string fornecido.
//
// Diferenças de SetInnerText()
//
// Não se confunda com as diferenças entre GetTextContent() / SetTextContent() e GetInnerText() /
// SetInnerText(). Embora os nomes pareçam semelhantes, existem diferenças importantes:
//
// GetTextContent() / SetTextContent() obtém o conteúdo de todos os elementos, incluindo os elementos
// <script> e <style>. Em contraste, GetInnerText() SetInnerText() mostra apenas elementos
// "legíveis para humanos".
//
// GetTextContent() retorna todos os elementos no nó. Em contraste, innerText está ciente do estilo e
// não retornará o texto de elementos "ocultos".
//
// Além disso, como GetInnerText() / SetInnerText() leva em consideração os estilos CSS, a leitura do
// valor de innerText aciona um refluxo para garantir estilos computados atualizados.
// (Os refluxos podem ser computacionalmente caros e, portanto, devem ser evitados quando possível.)
//
// Ambos SetTextContent() e SetInnerText() removem nós filho quando alterados, mas alterar innerText
// no Internet Explorer (versão 11 e inferior) também destrói permanentemente todos os nós de texto
// descendentes.
// É impossível inserir os nós novamente em qualquer outro elemento ou no mesmo elemento depois de
// fazê-lo.
//
// Diferenças de SetInnerHtml()
//
// GetInnerHtml() retorna HTML, como seu nome indica. Às vezes, as pessoas usam GetInnerHtml() /
// SetInnerHtml() para recuperar ou escrever texto dentro de um elemento, mas GetTextContent() /
// SetTextContent() tem melhor desempenho porque seu valor não é analisado como HTML.
//
// Além disso, usar GetTextContent() / SetTextContent() pode prevenir ataques XSS.
func (e *GlobalAttributes) GetTextContent() (text string) {
	return e.selfElement.Get("textContent").String()
}
