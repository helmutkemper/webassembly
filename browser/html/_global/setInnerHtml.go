package _global

// SetInnerHtml
//
// English:
//
//  The Element property GetInnerHtml() / SetInnerHtml() gets or sets the HTML or XML markup contained
//  within the element.
//
// To insert the HTML into the document rather than replace the contents of an element, use the method
// InsertAdjacentHtml().
//
// Value: A DOMString containing the HTML serialization of the element's descendants.
//
//   Note:
//     * Setting the value of SetInnerHtml() removes all of the element's descendants and replaces
//       them with nodes constructed by parsing the HTML given in the string htmlString.
//
// SyntaxError DOMException
//
// Thrown if an attempt was made to set the value of innerHTML using a string which is not
// properly-formed HTML.
//
// NoModificationAllowedError DOMException
//
//  Thrown if an attempt was made to insert the HTML into a node whose parent is a Document.
//
// Usage notes
//
// The innerHTML property can be used to examine the current HTML source of the page, including any
// changes that have been made since the page was initially loaded.
//
// Reading the HTML contents of an element
//
// Reading GetInnerHtml() causes the user agent to serialize the HTML or XML fragment comprised of the
// element's descendants. The resulting string is returned.
//
//   Note:
//     * The returned HTML or XML fragment is generated based on the current contents of the element,
//       so the markup and formatting of the returned fragment is likely not to match the original
//       page markup.
//
// Replacing the contents of an element
//
// Setting the value of SetInnerHtml() lets you easily replace the existing contents of an element
// with new content.
//
//   Note:
//     * This is a security risk if the string to be inserted might contain potentially malicious
//       content. When inserting user-supplied data you should always consider using SetInnerHtml()
//       instead, in order to sanitize the content before it is inserted.
//
// Português:
//
//  A propriedade Element GetInnerHtml() / SetInnerHtml() obtém ou define a marcação HTML ou XML
//  contida no elemento.
//
// Para inserir o HTML no documento em vez de substituir o conteúdo de um elemento, use o método
// InsertAdjacentHtml().
//
// Valor: Um DOMString contendo a serialização HTML dos descendentes do elemento.
//
//   Nota:
//     * Definir o valor de SetInnerHtml() remove todos os descendentes do elemento e os substitui por
//       nós construídos analisando o HTML fornecido na string htmlString.
//
// SyntaxError DOMException
//
// Lançado se foi feita uma tentativa de definir o valor de innerHTML usando uma string que não é
// HTML corretamente formada.
//
// NoModificationAllowedError DOMException
//
//  Lançado se foi feita uma tentativa de inserir o HTML em um nó cujo pai é um Documento.
//
// Notas de uso
//
// A propriedade innerHTML pode ser usada para examinar a fonte HTML atual da página, incluindo
// quaisquer alterações feitas desde que a página foi carregada inicialmente.
//
// Lendo o conteúdo HTML de um elemento
//
// A leitura de GetInnerHtml() faz com que o agente do usuário serialize o fragmento HTML ou XML
// composto pelos descendentes do elemento. A string resultante é retornada.
//
//   Nota:
//     * O fragmento HTML ou XML retornado é gerado com base no conteúdo atual do elemento, portanto,
//       a marcação e a formatação do fragmento retornado provavelmente não corresponderão à marcação
//       da página original.
//
// Substituindo o conteúdo de um elemento
//
// Definir o valor de SetInnerHtml() permite substituir facilmente o conteúdo existente de um elemento
// por um novo conteúdo.
//
//   Nota:
//     * Este é um risco de segurança se a string a ser inserida puder conter conteúdo potencialmente
//       mal-intencionado. Ao inserir dados fornecidos pelo usuário, você deve sempre considerar o uso
//       de SetInnerHtml(), para limpar o conteúdo antes de ser inserido.
func (e *GlobalAttributes) SetInnerHtml(html string) (ref *GlobalAttributes) {
	e.selfElement.Set("innerHTML", html)
	return e
}
