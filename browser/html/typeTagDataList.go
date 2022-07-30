package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"log"
	"strconv"
	"strings"
	"syscall/js"
)

// TagDataList
//
// English:
//
//	The <datalist> HTML element contains a set of <option> elements that represent the permissible
//	or recommended options available to choose from within other controls.
//
// Português:
//
//	O elemento HTML <datalist> contém um conjunto de elementos <option> que representam as opções
//	permitidas ou recomendadas disponíveis para escolha em outros controles.
type TagDataList struct {
	tag         Tag
	id          string
	selfElement js.Value
	cssClass    *css.Class
}

// Reference
//
// English:
//
// Pass the object reference to an external variable.
//
// Português:
//
// Passa a referencia do objeto para uma variável externa.
//
//	Example: / Exemplo:
//	  var circle *html.TagSvgCircle
//	  factoryBrowser.NewTagSvgCircle().Reference(&circle).R(5).Fill(factoryColor.NewRed())
//	  log.Printf("x: %v, y: %v", circle.GetX(), circle.GetY())
func (e *TagDataList) Reference(reference **TagDataList) (ref *TagDataList) {
	*reference = e
	return e
}

// AccessKey
//
// English:
//
//	Specifies a shortcut key to activate/focus an element.
//
//	 Input:
//	   character: A single character that specifies the shortcut key to activate/focus the element.
//
//	 Note:
//	   * The accessKey attribute value must be a single character (a letter or a digit).
//	   * Adapting accessKeys to all international languages are difficult.
//	   * The accessKey value may not be present on all keyboards.
//
//
//	 Warning:
//	   Using accessKeys is difficult because they may conflict with other key standards in the
//	   browser;
//	   To avoid this problem, most browsers will use accessKeys only if pressed together with the Alt
//	   key.
//
// Português:
//
//	Especifica uma tecla de atalho para ativar o foco de um elemento.
//
//	 Entrada:
//	   character: Um único caractere que especifica a tecla de atalho para ativar o foco do elemento.
//
//	 Nota:
//	   * O valor do atributo accessKey deve ser um único caractere (uma letra ou um dígito).
//	   * Adaptar as teclas de acesso a todos os idiomas internacionais é difícil.
//	   * O valor accessKey pode não estar presente em todos os teclados.
//
//	 Aviso:
//	   O uso de accessKeys é difícil porque eles podem entrar em conflito com outros padrões
//	   importantes no navegador;
//	   Para evitar esse problema, a maioria dos navegadores usará as teclas de acesso somente se
//	   pressionadas junto com a tecla Alt.
func (e *TagDataList) AccessKey(key string) (ref *TagDataList) {
	e.selfElement.Set("accesskey", key)
	return e
}

// Autofocus
//
// English:
//
//	This Boolean attribute specifies that the button should have input focus when the page loads.
//	Only one element in a document can have this attribute.
//
// Português:
//
//	Este atributo booleano especifica que o botão deve ter foco de entrada quando a página for
//	carregada. Apenas um elemento em um documento pode ter esse atributo.
func (e *TagDataList) Autofocus(autofocus bool) (ref *TagDataList) {
	e.selfElement.Set("autofocus", autofocus)
	return e
}

// Class
//
// English:
//
//	The class attribute specifies one or more class names for an element.
//
//	 Input:
//	   classname: Specifies one or more class names for an element. To specify multiple classes,
//	              separate the class names with a space, e.g. <span class="left important">.
//	              This allows you to combine several CSS classes for one HTML element.
//
//	              Naming rules:
//	                Must begin with a letter A-Z or a-z;
//	                Can be followed by: letters (A-Za-z), digits (0-9), hyphens ("-"), and
//	                underscores ("_").
//
// The class attribute is mostly used to point to a class in a style sheet. However, it can also be
// used by a JavaScript (via the HTML DOM) to make changes to HTML elements with a specified class.
//
// Português:
//
//	O atributo class especifica um ou mais nomes de classe para um elemento.
//
//	 Entrada:
//	   classname: Especifica um ou mais nomes de classe para um elemento. Para especificar várias
//	              classes, separe os nomes das classes com um espaço, por exemplo <span class="left
//	              important">.
//	              Isso permite combinar várias classes CSS para um elemento HTML.
//
//	              Regras de nomenclatura:
//	                Deve começar com uma letra A-Z ou a-z;
//	                Pode ser seguido por: letras (A-Za-z), dígitos (0-9), hífens ("-") e
//	                sublinhados ("_").
//
// The class attribute is mostly used to point to a class in a style sheet. However, it can also be
// used by a JavaScript (via the HTML DOM) to make changes to HTML elements with a specified class.
func (e *TagDataList) Class(class ...string) (ref *TagDataList) {
	e.selfElement.Set("classList", strings.Join(class, " "))
	return e
}

// ContentEditable
//
// English:
//
//	The contentEditable attribute specifies whether the content of an element is editable or not.
//
//	 Input:
//	   contentEditable: specifies whether the content of an element is editable or not
//
//	 Note:
//	   When the contentEditable attribute is not set on an element, the element will inherit it from
//	   its parent.
//
// Português:
//
//	O atributo contentEditable especifica se o conteúdo de um elemento é editável ou não.
//
//	 Entrada:
//	   contentEditable: especifica se o conteúdo de um elemento é editável ou não.
//
//	 Nota:
//	   Quando o atributo contentEditable não está definido em um elemento, o elemento o herdará de
//	   seu pai.
func (e *TagDataList) ContentEditable(editable bool) (ref *TagDataList) {
	e.selfElement.Set("contenteditable", editable)
	return e
}

// Data
//
// English:
//
//	Used to store custom data private to the page or application.
//
//	 Input:
//	   data: custom data private to the page or application.
//
// The data-* attributes is used to store custom data private to the page or application.
// The data-* attributes gives us the ability to embed custom data attributes on all HTML elements.
// The stored (custom) data can then be used in the page's JavaScript to create a more engaging user
// experience (without any Ajax calls or server-side database queries).
//
// The data-* attributes consist of two parts:
//
//	The attribute name should not contain any uppercase letters, and must be at least one character
//	long after the prefix "data-";
//	The attribute value can be any string.
//
//	Note:
//	  * Custom attributes prefixed with "data-" will be completely ignored by the user agent.
//
// Português:
//
//	Usado para armazenar dados personalizados privados para a página ou aplicativo.
//
//	 Entrada:
//	   data: dados personalizados privados para a página ou aplicativo.
//
// Os atributos de dados são usados para armazenar dados personalizados privados para a página ou
// aplicativo;
// Os atributos de dados nos dão a capacidade de incorporar atributos de dados personalizados em todos
// os elementos HTML;
// Os dados armazenados (personalizados) podem ser usados no JavaScript da página para criar uma
// experiência de usuário mais envolvente (sem chamadas Ajax ou consultas de banco de dados do lado do
// servidor).
//
// Os atributos de dados consistem em duas partes:
//
//	O nome do atributo não deve conter letras maiúsculas e deve ter pelo menos um caractere após o
//	prefixo "data-";
//	O valor do atributo pode ser qualquer string.
//
//	Nota:
//	  * Atributos personalizados prefixados com "data-" serão completamente ignorados pelo agente do
//	    usuário.
func (e *TagDataList) Data(data map[string]string) (ref *TagDataList) {
	for k, v := range data {
		e.selfElement.Set(" data-"+k, v)
	}
	return e
}

// Dir
//
// English:
//
//	Specifies the text direction for the content in an element.
//
//	 Input:
//	   dir: direction for the content in an element. [ KDirLeftToRight | KDirRightToLeft | KDirAuto ]
//
// Português:
//
//	Especifica a direção do texto para o conteúdo em um elemento.
//
//	 Entrada:
//	   dir: direção do texto para o conteúdo em um elemento. [ KDirLeftToRight | KDirRightToLeft |
//	        KDirAuto ]
func (e *TagDataList) Dir(dir Dir) (ref *TagDataList) {
	e.selfElement.Set("dir", dir.String())
	return e
}

// Draggable
//
// English:
//
//	Specifies whether an element is draggable or not.
//
//	 Input:
//	   draggable: element is draggable or not. [ KDraggableYes | KDraggableNo | KDraggableAuto ]
//
// The draggable attribute specifies whether an element is draggable or not.
//
//	Note:
//	  * Links and images are draggable by default;
//	  * The draggable attribute is often used in drag and drop operations.
//	  * Read our HTML Drag and Drop tutorial to learn more.
//	    https://www.w3schools.com/html/html5_draganddrop.asp
//
// Português:
//
//	Especifica se um elemento pode ser arrastado ou não. [ KDraggableYes | KDraggableNo |
//	KDraggableAuto ]
//
//	 Entrada:
//	   draggable: elemento é arrastável ou não.
//
// O atributo arrastável especifica se um elemento é arrastável ou não.
//
//	Nota:
//	  * Links e imagens podem ser arrastados por padrão;
//	  * O atributo arrastável é frequentemente usado em operações de arrastar e soltar.
//	  * Leia nosso tutorial de arrastar e soltar HTML para saber mais.
//	    https://www.w3schools.com/html/html5_draganddrop.asp
func (e *TagDataList) Draggable(draggable Draggable) (ref *TagDataList) {
	e.selfElement.Set("draggable", draggable.String())
	return e
}

// EnterKeyHint
//
// English:
//
//	The enterKeyHint property is an enumerated property defining what action label (or icon) to
//	present for the enter key on virtual keyboards. It reflects the enterkeyhint HTML global attribute
//	and is an enumerated property, only accepting the following values as a DOMString:
//
//	 Input:
//	   enterKeyHint: defining what action label (or icon) to present for the enter key on virtual
//	     keyboards
//	     KEnterKeyHintEnter: typically indicating inserting a new line.
//	     KEnterKeyHintDone: typically meaning there is nothing more to input and the input method
//	      editor (IME) will be closed.
//	     KEnterKeyHintGo: typically meaning to take the user to the target of the text they typed.
//	     KEnterKeyHintNext: typically taking the user to the next field that will accept text.
//	     KEnterKeyHintPrevious: typically taking the user to the previous field that will accept text.
//	     KEnterKeyHintSearch: typically taking the user to the results of searching for the text they
//	       have typed.
//	     KEnterKeyHintSend: typically delivering the text to its target.
//
// If no enterKeyHint value has been specified or if it was set to a different value than the allowed
// ones, it will return an empty string.
//
// Português:
//
//	A propriedade enterKeyHint é uma propriedade enumerada que define qual rótulo de ação (ou ícone)
//	apresentar para a tecla Enter em teclados virtuais. Ele reflete o atributo global enterkeyhint
//	HTML e é uma propriedade enumerada, aceitando apenas os seguintes valores como DOMString:
//
//	 Entrada:
//	   enterKeyHint: definindo qual rótulo de ação (ou ícone) apresentar para a tecla Enter em
//	     teclados virtuais
//	     KEnterKeyHintEnter: normalmente indicando a inserção de uma nova linha.
//	     KEnterKeyHintDone: normalmente significa que não há mais nada para inserir e o editor de
//	       método de entrada (IME) será fechado.
//	     KEnterKeyHintGo: normalmente significando levar o usuário ao destino do texto digitado.
//	     KEnterKeyHintNext: normalmente levando o usuário para o próximo campo que aceitará texto.
//	     KEnterKeyHintPrevious: normalmente levando o usuário ao campo anterior que aceitará texto.
//	     KEnterKeyHintSearch: normalmente levando o usuário aos resultados da pesquisa do texto que
//	       digitou.
//	     KEnterKeyHintSend: normalmente entregando o texto ao seu destino.
//
// Se nenhum valor enterKeyHint foi especificado ou se foi definido com um valor diferente dos
// permitidos, ele retornará uma string vazia.
func (e *TagDataList) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagDataList) {
	e.selfElement.Set("enterKeyHint", enterKeyHint.String())
	return e
}

// Hidden
//
// English:
//
//	Specifies that an element is not yet, or is no longer, relevant.
//
//	 Input:
//	   hidden:
//
// The hidden attribute is a boolean attribute.
//
// When present, it specifies that an element is not yet, or is no longer, relevant.
//
// Browsers should not display elements that have the hidden attribute specified.
//
// The hidden attribute can also be used to keep a user from seeing an element until some other
// condition has been met (like selecting a checkbox, etc.). Then, a JavaScript could remove the
// hidden attribute, and make the element visible.
//
// Português:
//
//	Especifica que um elemento ainda não é ou não é mais relevante.
//
// O atributo oculto é um atributo booleano.
//
// Quando presente, especifica que um elemento ainda não é ou não é mais relevante.
//
// Os navegadores não devem exibir elementos que tenham o atributo oculto especificado.
//
// O atributo oculto também pode ser usado para impedir que um usuário veja um elemento até que alguma
// outra condição seja atendida (como marcar uma caixa de seleção etc.). Então, um JavaScript pode
// remover o atributo oculto e tornar o elemento visível.
func (e *TagDataList) Hidden() (ref *TagDataList) {
	e.selfElement.Get("style").Set("visibility", "hidden")
	return e
}

// Id
//
// English:
//
//	Specifies a unique id for an element
//
// The id attribute specifies a unique id for an HTML element (the value must be unique within the
// HTML document).
//
// The id attribute is most used to point to a style in a style sheet, and by JavaScript (via the HTML
// DOM) to manipulate the element with the specific id.
//
// Português:
//
//	Especifica um ID exclusivo para um elemento
//
// O atributo id especifica um id exclusivo para um elemento HTML (o valor deve ser exclusivo no
// documento HTML).
//
// O atributo id é mais usado para apontar para um estilo em uma folha de estilo, e por JavaScript
// (através do HTML DOM) para manipular o elemento com o id específico.
func (e *TagDataList) Id(id string) (ref *TagDataList) {
	e.id = id
	e.selfElement.Set("id", id)

	// Saves the element reference with ID for later use.
	// Salva a referência dos elementos com ID para uso posterior.
	htmlGlobalAllElementsList.Delete(id)
	htmlGlobalAllElementsList.Store(id, e)
	return e
}

// InputMode
//
// English:
//
//	The inputmode global attribute is an enumerated attribute that hints at the type of data that
//	might be entered by the user while editing the element or its contents. This allows a browser to
//	display an appropriate virtual keyboard.
//
// It is used primarily on <input> elements, but is usable on any element in contenteditable mode.
//
// It's important to understand that the inputmode attribute doesn't cause any validity requirements
// to be enforced on input. To require that input conforms to a particular data type, choose an
// appropriate <input> element type. For specific guidance on choosing <input> types, see the Values
// section.
//
// Português:
//
//	O atributo global inputmode é um atributo enumerado que indica o tipo de dados que pode ser
//	inserido pelo usuário ao editar o elemento ou seu conteúdo. Isso permite que um navegador exiba
//	um teclado virtual apropriado.
//
// Ele é usado principalmente em elementos <input>, mas pode ser usado em qualquer elemento no modo
// contenteditable.
//
// É importante entender que o atributo inputmode não faz com que nenhum requisito de validade seja
// imposto na entrada. Para exigir que a entrada esteja em conformidade com um tipo de dados
// específico, escolha um tipo de elemento <input> apropriado. Para obter orientações específicas
// sobre como escolher os tipos de <input>, consulte a seção Valores.
func (e *TagDataList) InputMode(inputMode InputMode) (ref *TagDataList) {
	e.selfElement.Set("inputmode", inputMode.String())
	return e
}

// Is
//
// English:
//
//	Allows you to specify that a standard HTML element should behave like a registered custom
//	built-in element.
//
// Português:
//
//	Permite especificar que um elemento HTML padrão deve se comportar como um elemento interno
//	personalizado registrado.
func (e *TagDataList) Is(is string) (ref *TagDataList) {
	e.selfElement.Set("is", is)
	return e
}

// ItemId
//
// English:
//
//	The unique, global identifier of an item.
//
// Português:
//
//	O identificador global exclusivo de um item.
func (e *TagDataList) ItemId(id string) (ref *TagDataList) {
	e.selfElement.Set("itemid", id)
	return e
}

// ItemProp
//
// English:
//
//	Used to add properties to an item. Every HTML element may have an itemprop attribute specified,
//	where an itemprop consists of a name and value pair.
//
// Português:
//
//	Usado para adicionar propriedades a um item. Cada elemento HTML pode ter um atributo itemprop
//	especificado, onde um itemprop consiste em um par de nome e valor.
func (e *TagDataList) ItemProp(itemprop string) (ref *TagDataList) {
	e.selfElement.Set("itemprop", itemprop)
	return e
}

// ItemRef
//
// English:
//
//	Properties that are not descendants of an element with the itemscope attribute can be associated
//	with the item using an itemref. It provides a list of element ids (not itemids) with additional
//	properties elsewhere in the document.
//
// Português:
//
//	Propriedades que não são descendentes de um elemento com o atributo itemscope podem ser
//	associadas ao item usando um itemref. Ele fornece uma lista de IDs de elementos (não IDs de itens)
//	com propriedades adicionais em outras partes do documento.
func (e *TagDataList) ItemRef(itemref string) (ref *TagDataList) {
	e.selfElement.Set("itemref", itemref)
	return e
}

// ItemType
//
// English:
//
//	Specifies the URL of the vocabulary that will be used to define itemprops (item properties) in
//	the data structure. itemscope is used to set the scope of where in the data structure the
//	vocabulary set by itemtype will be active.
//
// Português:
//
//	Especifica a URL do vocabulário que será usado para definir itemprops (propriedades do item) na
//	estrutura de dados. itemscope é usado para definir o escopo de onde na estrutura de dados o
//	vocabulário definido por tipo de item estará ativo.
func (e *TagDataList) ItemType(itemType string) (ref *TagDataList) {
	e.selfElement.Set("itemtype", itemType)
	return e
}

// Lang
//
// English:
//
//	Specifies the language of the element's content.
//
// The lang attribute specifies the language of the element's content.
//
// Common examples are KLanguageEnglish for English, KLanguageSpanish for Spanish, KLanguageFrench
// for French, and so on.
//
// Português:
//
//	Especifica o idioma do conteúdo do elemento.
//
// O atributo lang especifica o idioma do conteúdo do elemento.
//
// Exemplos comuns são KLanguageEnglish para inglês, KLanguageSpanish para espanhol, KLanguageFrench
// para francês e assim por diante.
func (e *TagDataList) Lang(language Language) (ref *TagDataList) {
	e.selfElement.Set("lang", language.String())
	return e
}

// Nonce
//
// English:
//
//	A cryptographic nonce ("number used once") which can be used by Content Security Policy to
//	determine whether or not a given fetch will be allowed to proceed.
//
// Português:
//
//	Um nonce criptográfico ("número usado uma vez") que pode ser usado pela Política de Segurança de
//	Conteúdo para determinar se uma determinada busca terá permissão para prosseguir.
func (e *TagDataList) Nonce(nonce int64) (ref *TagDataList) {
	e.selfElement.Set("nonce", nonce)
	return e
}

// Slot
//
// English:
//
//	Assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is
//	assigned to the slot created by the <slot> element whose name attribute's value matches that slot
//	attribute's value.
//
// Português:
//
//	Atribui um slot em uma shadow DOM shadow tree a um elemento: Um elemento com um atributo slot é
//	atribuído ao slot criado pelo elemento <slot> cujo valor do atributo name corresponde ao valor
//	desse atributo slot.
func (e *TagDataList) Slot(slot string) (ref *TagDataList) {
	e.selfElement.Set("slot", slot)
	return e
}

// Spellcheck
//
// English:
//
//	Specifies whether the element is to have its spelling and grammar checked or not
//
//	 Note:
//	   * The following can be spellchecked:
//	       Text values in input elements (not password)
//	       Text in <textarea> elements
//	       Text in editable elements
//
// Português:
//
//	Especifica se o elemento deve ter sua ortografia e gramática verificadas ou não
//
// O seguinte pode ser verificado ortográfico:
//
//	Nota:
//	  * O seguinte pode ser verificado ortográfico:
//	      Valores de texto em elementos de entrada (não senha)
//	      Texto em elementos <textarea>
//	      Texto em elementos editáveis
func (e *TagDataList) Spellcheck(spell bool) (ref *TagDataList) {
	e.selfElement.Set("spellcheck", spell)

	return e
}

// Style
//
// English:
//
//	Specifies an inline CSS style for an element.
//
// The style attribute will override any style set globally, e.g. styles specified in the <style> tag
// or in an external style sheet.
//
// The style attribute can be used on any HTML element (it will validate on any HTML element.
// However, it is not necessarily useful).
//
// Português:
//
//	Especifica um estilo CSS embutido para um elemento
//
// O atributo style substituirá qualquer conjunto de estilos globalmente, por exemplo estilos
// especificados na tag <style> ou em uma folha de estilo externa.
//
// O atributo style pode ser usado em qualquer elemento HTML (vai validar em qualquer elemento HTML.
// No entanto, não é necessariamente útil).
func (e *TagDataList) Style(style string) (ref *TagDataList) {
	e.selfElement.Set("style", style)
	return e
}

// TabIndex
//
// English:
//
//	Specifies the tabbing order of an element (when the "tab" button is used for navigating).
//
// The tabindex attribute can be used on any HTML element (it will validate on any HTML element.
// However, it is not necessarily useful).
//
// Português:
//
//	Especifica a ordem de tabulação de um elemento (quando o botão "tab" é usado para navegar).
//
// O atributo tabindex pode ser usado em qualquer elemento HTML (vai validar em qualquer elemento
// HTML. No entanto, não é necessariamente útil).
func (e *TagDataList) TabIndex(index int) (ref *TagDataList) {
	e.selfElement.Set("tabindex", index)
	return e
}

// Title
//
// English:
//
//	Specifies extra information about an element.
//
// The information is most often shown as a tooltip text when the mouse moves over the element.
//
// The title attribute can be used on any HTML element (it will validate on any HTML element.
// However, it is not necessarily useful).
//
// Português:
//
//	Especifica informações extras sobre um elemento.
//
// As informações geralmente são mostradas como um texto de dica de ferramenta quando o mouse se move
// sobre o elemento.
//
// O atributo title pode ser usado em qualquer elemento HTML (vai validar em qualquer elemento HTML.
// No entanto, não é necessariamente útil).
func (e *TagDataList) Title(title string) (ref *TagDataList) {
	e.selfElement.Set("title", title)
	return e
}

// Translate
//
// English:
//
//	Specifies whether the content of an element should be translated or not.
//
//	 Input:
//	   translate: element should be translated or not. [ KTranslateYes | KTranslateNo ]
//
// Português:
//
//	Especifica se o conteúdo de um elemento deve ser traduzido ou não.
//
//	 Entrada:
//	   translate: elemento deve ser traduzido ou não. [ KTranslateYes | KTranslateNo ]
func (e *TagDataList) Translate(translate Translate) (ref *TagDataList) {
	e.selfElement.Set("translate", translate.String())
	return e
}

// CreateElement
//
// English:
//
//	In an HTML document, the Document.createElement() method creates the specified HTML element or an
//	HTMLUnknownElement if the given element name is not known.
//
// Português:
//
//	Em um documento HTML, o método Document.createElement() cria o elemento HTML especificado ou um
//	HTMLUnknownElement se o nome do elemento dado não for conhecido.
func (e *TagDataList) CreateElement(tag Tag) (ref *TagDataList) {
	e.selfElement = js.Global().Get("document").Call("createElement", tag.String())
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}
	e.tag = tag

	return e
}

// AppendById
//
// English:
//
//	Adds a node to the end of the list of children of a specified parent node. If the node already
//	exists in the document, it is removed from its current parent node before being added to the
//	new parent.
//
//	 Input:
//	   appendId: id of parent element.
//
//	 Note:
//	   * The equivalent of:
//	       var p = document.createElement("p");
//	       document.body.appendChild(p);
//
// Português:
//
//	Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no
//	documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//	 Entrada:
//	   appendId: id do elemento pai.
//
//	 Nota:
//	   * Equivale a:
//	       var p = document.createElement("p");
//	       document.body.appendChild(p);
func (e *TagDataList) AppendById(appendId string) (ref *TagDataList) {

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return e
	}

	toAppend.Call("appendChild", e.selfElement)
	return e
}

// Append
//
// English:
//
//	Adds a node to the end of the list of children of a specified parent node. If the node already
//	exists in the document, it is removed from its current parent node before being added to the new
//	parent.
//
//	 Input:
//	   append: element in js.Value format.
//
//	 Note:
//	   * The equivalent of:
//	       var p = document.createElement("p");
//	       document.body.appendChild(p);
//
// Português:
//
//	Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no
//	documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//	 Entrada:
//	   appendId: elemento no formato js.Value.
//
//	 Nota:
//	   * Equivale a:
//	       var p = document.createElement("p");
//	       document.body.appendChild(p);
func (e *TagDataList) Append(append interface{}) (ref *TagDataList) {
	switch append.(type) {
	case *TagDataList:
		e.selfElement.Call("appendChild", append.(*TagDataList).selfElement)
	case js.Value:
		e.selfElement.Call("appendChild", append)
	case string:
		toAppend := js.Global().Get("document").Call("getElementById", append.(string))
		if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
			log.Print(KIdToAppendNotFound, append.(string))
			return e
		}

		toAppend.Call("appendChild", e.selfElement)
	}

	return e
}

// NewOption
//
// English:
//
//	The <option> HTML element is used to define an item contained in a <select>, an <optgroup>, or
//	a <datalist> element. As such, <option> can represent menu items in popups and other lists of
//	items in an HTML document.
//
//	 Input:
//	   id: a unique id for an element;
//	   label: This attribute is text for the label indicating the meaning of the option. If the label
//	     attribute isn't defined, its value is that of the element text content;
//	   value: The content of this attribute represents the value to be submitted with the form, should
//	     this option be selected. If this attribute is omitted, the value is taken from the text
//	     content of the option element;
//	   disabled: If this Boolean attribute is set, this option is not checkable. Often browsers grey
//	     out such control and it won't receive any browsing event, like mouse clicks or focus-related
//	     ones. If this attribute is not set, the element can still be disabled if one of its ancestors
//	     is a disabled <optgroup> element;
//	   selected: If present, this Boolean attribute indicates that the option is initially selected.
//	     If the <option> element is the descendant of a <select> element whose multiple attribute is
//	     not set, only one single <option> of this <select> element may have the selected attribute.
//
// Português:
//
//	O elemento HTML <option> é usado para definir um item contido em um elemento <select>, <optgroup>
//	ou <datalist>. Como tal, <option> pode representar itens de menu em pop-ups e outras listas de
//	itens em um documento HTML.
//
//	 Entrada:
//	   id: um id exclusivo para um elemento;
//	   label: Este atributo é um texto para o rótulo que indica o significado da opção. Se o atributo
//	     label não estiver definido, seu valor será o do conteúdo do texto do elemento;
//	   value: O conteúdo deste atributo representa o valor a ser enviado com o formulário, caso esta
//	     opção seja selecionada. Se este atributo for omitido, o valor será obtido do conteúdo de
//	     texto do elemento de opção;
//	   disabled: Se este atributo booleano estiver definido, esta opção não poderá ser marcada.
//	     Muitas vezes, os navegadores desativam esse controle e não recebem nenhum evento de
//	     navegação, como cliques do mouse ou relacionados ao foco. Se este atributo não for definido,
//	     o elemento ainda poderá ser desabilitado se um de seus ancestrais for um elemento <optgroup>
//	     desabilitado;
//	   selected: Se presente, este atributo booleano indica que a opção foi selecionada inicialmente.
//	     Se o elemento <option> é descendente de um elemento <select> cujo atributo múltiplo não está
//	     definido, apenas um único <option> deste elemento <select> pode ter o atributo selecionado.
func (e *TagDataList) NewOption(id, label, value string, disabled, selected bool) (ref *TagDataList) {

	ref = &TagDataList{}
	ref.CreateElement(KTagOption)
	ref.Id(id)
	ref.value(value)
	ref.textContent(label)

	if disabled == true {
		ref.disabled(disabled)
	}

	if selected == true {
		ref.selected(selected)
	}

	e.Append(ref.selfElement)

	return e
}

// value
//
// English:
//
//	Defines the value associated with the element.
//
// Português:
//
//	Define o valor associado ao elemento.
func (e *TagDataList) value(value string) (ref *TagDataList) {
	e.selfElement.Set("value", value)
	return e
}

// textContent
//
// English:
//
//	The textContent property of the Node interface represents the text content of the node and its
//	descendants.
//
//	 Note:
//	   * SetTextContent() and SetInnerText() are easily confused, but the two properties are different
//	     in important ways.
//	   * Setting SetTextContent() on a node removes all of the node's children and replaces them with
//	     a single text node with the given string value.
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
//	A propriedade textContent da interface Node representa o conteúdo de texto do nó e seus
//	descendentes.
//
//	 Nota:
//	   * SetTextContent() e SetInnerText() são facilmente confundidos, mas as duas propriedades são
//	     diferentes em aspectos importantes;
//	   * Definir SetTextContent() em um nó remove todos os filhos do nó e os substitui por um único nó
//	     de texto com o valor de string fornecido.
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
func (e *TagDataList) textContent(text string) (ref *TagDataList) {
	e.selfElement.Set("textContent", text)
	return e
}

// disabled
//
// English:
//
//	Este atributo booleano impede que o usuário interaja com o elemento.
//
// Português:
//
//	Este atributo booleano impede que o usuário interaja com o elemento.
func (e *TagDataList) disabled(disabled bool) (ref *TagDataList) {
	e.selfElement.Set("disabled", disabled)
	return e
}

// SetSelected
//
// English:
//
//	If present, this Boolean attribute indicates that the option is initially selected. If the
//	<option> element is the descendant of a <select> element whose multiple attribute is not set,
//	only one single <option> of this <select> element may have the selected attribute.
//
// Português:
//
//	Se presente, este atributo booleano indica que a opção foi selecionada inicialmente. Se o elemento
//	<option> é descendente de um elemento <select> cujo atributo múltiplo não está definido, apenas um
//	único <option> deste elemento <select> pode ter o atributo selecionado.
func (e *TagDataList) selected(selected bool) (ref *TagDataList) {
	e.selfElement.Set("selected", selected)
	return e
}

// SetXY
//
// English:
//
//	Sets the X and Y axes in pixels.
//
// Português:
//
//	Define os eixos X e Y em pixels.
func (e *TagDataList) SetXY(x, y int) (ref *TagDataList) {
	px := strconv.FormatInt(int64(x), 10) + "px"
	py := strconv.FormatInt(int64(y), 10) + "px"

	e.selfElement.Get("style").Set("left", px)
	e.selfElement.Get("style").Set("top", py)

	return e
}

// SetX
//
// English:
//
//	Sets the X axe in pixels.
//
// Português:
//
//	Define o eixo X em pixels.
func (e *TagDataList) SetX(x int) (ref *TagDataList) {
	px := strconv.FormatInt(int64(x), 10) + "px"
	e.selfElement.Get("style").Set("left", px)

	return e
}

// SetY
//
// English:
//
//	Sets the Y axe in pixels.
//
// Português:
//
//	Define o eixo Y em pixels.
func (e *TagDataList) SetY(y int) (ref *TagDataList) {
	py := strconv.FormatInt(int64(y), 10) + "px"
	e.selfElement.Get("style").Set("top", py)

	return e
}

// GetXY
//
// English:
//
//	Returns the X and Y axes in pixels.
//
// Português:
//
//	Retorna os eixos X e Y em pixels.
func (e *TagDataList) GetXY() (x, y int) {
	x = e.selfElement.Get("style").Get("left").Int()
	y = e.selfElement.Get("style").Get("top").Int()

	return
}

// GetX
//
// English:
//
//	Returns the X axe in pixels.
//
// Português:
//
//	Retorna o eixo X em pixels.
func (e *TagDataList) GetX() (x int) {
	x = e.selfElement.Get("style").Get("left").Int()

	return
}

// GetY
//
// English:
//
//	Returns the Y axe in pixels.
//
// Português:
//
//	Retorna o eixo Y em pixels.
func (e *TagDataList) GetY() (y int) {
	y = e.selfElement.Get("style").Get("top").Int()

	return
}
