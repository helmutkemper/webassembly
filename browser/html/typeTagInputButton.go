package html

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browser/css"
	"log"
	"strconv"
	"strings"
	"syscall/js"
)

type TagInputButton struct {
	tag         Tag
	id          string
	selfElement js.Value
	cssClass    *css.Class
}

// AccessKey
//
// English:
//
//  Specifies a shortcut key to activate/focus an element.
//
//   Input:
//     character: A single character that specifies the shortcut key to activate/focus the element.
//
//   Note:
//     * The accessKey attribute value must be a single character (a letter or a digit).
//     * Adapting accessKeys to all international languages are difficult.
//     * The accessKey value may not be present on all keyboards.
//
//
//   Warning:
//     Using accessKeys is difficult because they may conflict with other key standards in the
//     browser;
//     To avoid this problem, most browsers will use accessKeys only if pressed together with the Alt
//     key.
//
// Português:
//
//  Especifica uma tecla de atalho para ativar o foco de um elemento.
//
//   Entrada:
//     character: Um único caractere que especifica a tecla de atalho para ativar o foco do elemento.
//
//   Nota:
//     * O valor do atributo accessKey deve ser um único caractere (uma letra ou um dígito).
//     * Adaptar as teclas de acesso a todos os idiomas internacionais é difícil.
//     * O valor accessKey pode não estar presente em todos os teclados.
//
//   Aviso:
//     O uso de accessKeys é difícil porque eles podem entrar em conflito com outros padrões
//     importantes no navegador;
//     Para evitar esse problema, a maioria dos navegadores usará as teclas de acesso somente se
//     pressionadas junto com a tecla Alt.
func (e *TagInputButton) AccessKey(key string) (ref *TagInputButton) {
	e.selfElement.Set("accesskey", key)
	return e
}

// Autofocus
//
// English:
//
//  This Boolean attribute specifies that the button should have input focus when the page loads.
//  Only one element in a document can have this attribute.
//
// Português:
//
//  Este atributo booleano especifica que o botão deve ter foco de entrada quando a página for
//  carregada. Apenas um elemento em um documento pode ter esse atributo.
func (e *TagInputButton) Autofocus(autofocus bool) (ref *TagInputButton) {
	e.selfElement.Set("autofocus", autofocus)
	return e
}

// Class
//
// English:
//
//  The class attribute specifies one or more class names for an element.
//
//   Input:
//     classname: Specifies one or more class names for an element. To specify multiple classes,
//                separate the class names with a space, e.g. <span class="left important">.
//                This allows you to combine several CSS classes for one HTML element.
//
//                Naming rules:
//                  Must begin with a letter A-Z or a-z;
//                  Can be followed by: letters (A-Za-z), digits (0-9), hyphens ("-"), and
//                  underscores ("_").
//
// The class attribute is mostly used to point to a class in a style sheet. However, it can also be
// used by a JavaScript (via the HTML DOM) to make changes to HTML elements with a specified class.
//
// Português:
//
//  O atributo class especifica um ou mais nomes de classe para um elemento.
//
//   Entrada:
//     classname: Especifica um ou mais nomes de classe para um elemento. Para especificar várias
//                classes, separe os nomes das classes com um espaço, por exemplo <span class="left
//                important">.
//                Isso permite combinar várias classes CSS para um elemento HTML.
//
//                Regras de nomenclatura:
//                  Deve começar com uma letra A-Z ou a-z;
//                  Pode ser seguido por: letras (A-Za-z), dígitos (0-9), hífens ("-") e
//                  sublinhados ("_").
//
// O atributo class é usado principalmente para apontar para uma classe em uma folha de estilo.
// No entanto, também pode ser usado por um JavaScript (através do HTML DOM) para fazer alterações
// em elementos HTML com uma classe especificada.
func (e *TagInputButton) Class(class ...string) (ref *TagInputButton) {
	e.selfElement.Set("classList", strings.Join(class, " "))
	return e
}

// ContentEditable
//
// English:
//
//  The contentEditable attribute specifies whether the content of an element is editable or not.
//
//   Input:
//     contentEditable: specifies whether the content of an element is editable or not
//
//   Note:
//     When the contentEditable attribute is not set on an element, the element will inherit it from
//     its parent.
//
// Português:
//
//  O atributo contentEditable especifica se o conteúdo de um elemento é editável ou não.
//
//   Entrada:
//     contentEditable: especifica se o conteúdo de um elemento é editável ou não.
//
//   Nota:
//     Quando o atributo contentEditable não está definido em um elemento, o elemento o herdará de
//     seu pai.
func (e *TagInputButton) ContentEditable(editable bool) (ref *TagInputButton) {
	e.selfElement.Set("contenteditable", editable)
	return e
}

// Data
//
// English:
//
//  Used to store custom data private to the page or application.
//
//   Input:
//     data: custom data private to the page or application.
//
// The data-* attributes is used to store custom data private to the page or application.
// The data-* attributes gives us the ability to embed custom data attributes on all HTML elements.
// The stored (custom) data can then be used in the page's JavaScript to create a more engaging user
// experience (without any Ajax calls or server-side database queries).
//
// The data-* attributes consist of two parts:
//   The attribute name should not contain any uppercase letters, and must be at least one character
//   long after the prefix "data-";
//   The attribute value can be any string.
//
//   Note:
//     * Custom attributes prefixed with "data-" will be completely ignored by the user agent.
//
// Português:
//
//  Usado para armazenar dados personalizados privados para a página ou aplicativo.
//
//   Entrada:
//     data: dados personalizados privados para a página ou aplicativo.
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
//   O nome do atributo não deve conter letras maiúsculas e deve ter pelo menos um caractere após o
//   prefixo "data-";
//   O valor do atributo pode ser qualquer string.
//
//   Nota:
//     * Atributos personalizados prefixados com "data-" serão completamente ignorados pelo agente do
//       usuário.
func (e *TagInputButton) Data(data map[string]string) (ref *TagInputButton) {
	for k, v := range data {
		e.selfElement.Set(" data-"+k, v)
	}
	return e
}

// Dir
//
// English:
//
//  Specifies the text direction for the content in an element.
//
//   Input:
//     dir: direction for the content in an element. [ KDirLeftToRight | KDirRightToLeft | KDirAuto ]
//
// Português:
//
//  Especifica a direção do texto para o conteúdo em um elemento.
//
//   Entrada:
//     dir: direção do texto para o conteúdo em um elemento. [ KDirLeftToRight | KDirRightToLeft |
//          KDirAuto ]
func (e *TagInputButton) Dir(dir Dir) (ref *TagInputButton) {
	e.selfElement.Set("dir", dir.String())
	return e
}

// Draggable
//
// English:
//
//  Specifies whether an element is draggable or not.
//
//   Input:
//     draggable: element is draggable or not. [ KDraggableYes | KDraggableNo | KDraggableAuto ]
//
// The draggable attribute specifies whether an element is draggable or not.
//
//   Note:
//     * Links and images are draggable by default;
//     * The draggable attribute is often used in drag and drop operations.
//     * Read our HTML Drag and Drop tutorial to learn more.
//       https://www.w3schools.com/html/html5_draganddrop.asp
//
// Português:
//
//  Especifica se um elemento pode ser arrastado ou não. [ KDraggableYes | KDraggableNo |
//  KDraggableAuto ]
//
//   Entrada:
//     draggable: elemento é arrastável ou não.
//
// O atributo arrastável especifica se um elemento é arrastável ou não.
//
//   Nota:
//     * Links e imagens podem ser arrastados por padrão;
//     * O atributo arrastável é frequentemente usado em operações de arrastar e soltar.
//     * Leia nosso tutorial de arrastar e soltar HTML para saber mais.
//       https://www.w3schools.com/html/html5_draganddrop.asp
func (e *TagInputButton) Draggable(draggable Draggable) (ref *TagInputButton) {
	e.selfElement.Set("draggable", draggable.String())
	return e
}

// EnterKeyHint
//
// English:
//
//  The enterKeyHint property is an enumerated property defining what action label (or icon) to
//  present for the enter key on virtual keyboards. It reflects the enterkeyhint HTML global attribute
//  and is an enumerated property, only accepting the following values as a DOMString:
//
//   Input:
//     enterKeyHint: defining what action label (or icon) to present for the enter key on virtual
//       keyboards
//       KEnterKeyHintEnter: typically indicating inserting a new line.
//       KEnterKeyHintDone: typically meaning there is nothing more to input and the input method
//        editor (IME) will be closed.
//       KEnterKeyHintGo: typically meaning to take the user to the target of the text they typed.
//       KEnterKeyHintNext: typically taking the user to the next field that will accept text.
//       KEnterKeyHintPrevious: typically taking the user to the previous field that will accept text.
//       KEnterKeyHintSearch: typically taking the user to the results of searching for the text they
//         have typed.
//       KEnterKeyHintSend: typically delivering the text to its target.
//
// If no enterKeyHint value has been specified or if it was set to a different value than the allowed
// ones, it will return an empty string.
//
// Português:
//
//  A propriedade enterKeyHint é uma propriedade enumerada que define qual rótulo de ação (ou ícone)
//  apresentar para a tecla Enter em teclados virtuais. Ele reflete o atributo global enterkeyhint
//  HTML e é uma propriedade enumerada, aceitando apenas os seguintes valores como DOMString:
//
//   Entrada:
//     enterKeyHint: definindo qual rótulo de ação (ou ícone) apresentar para a tecla Enter em
//       teclados virtuais
//       KEnterKeyHintEnter: normalmente indicando a inserção de uma nova linha.
//       KEnterKeyHintDone: normalmente significa que não há mais nada para inserir e o editor de
//         método de entrada (IME) será fechado.
//       KEnterKeyHintGo: normalmente significando levar o usuário ao destino do texto digitado.
//       KEnterKeyHintNext: normalmente levando o usuário para o próximo campo que aceitará texto.
//       KEnterKeyHintPrevious: normalmente levando o usuário ao campo anterior que aceitará texto.
//       KEnterKeyHintSearch: normalmente levando o usuário aos resultados da pesquisa do texto que
//         digitou.
//       KEnterKeyHintSend: normalmente entregando o texto ao seu destino.
//
// Se nenhum valor enterKeyHint foi especificado ou se foi definido com um valor diferente dos
// permitidos, ele retornará uma string vazia.
func (e *TagInputButton) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagInputButton) {
	e.selfElement.Set("enterKeyHint", enterKeyHint.String())
	return e
}

// Hidden
//
// English:
//
//  Specifies that an element is not yet, or is no longer, relevant.
//
//   Input:
//     hidden:
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
//  Especifica que um elemento ainda não é ou não é mais relevante.
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
func (e *TagInputButton) Hidden() (ref *TagInputButton) {
	e.selfElement.Get("style").Set("visibility", "hidden")
	return e
}

// Id
//
// English:
//
//  Specifies a unique id for an element
//
// The id attribute specifies a unique id for an HTML element (the value must be unique within the
// HTML document).
//
// The id attribute is most used to point to a style in a style sheet, and by JavaScript (via the HTML
// DOM) to manipulate the element with the specific id.
//
// Português:
//
//  Especifica um ID exclusivo para um elemento
//
// O atributo id especifica um id exclusivo para um elemento HTML (o valor deve ser exclusivo no
// documento HTML).
//
// O atributo id é mais usado para apontar para um estilo em uma folha de estilo, e por JavaScript
// (através do HTML DOM) para manipular o elemento com o id específico.
func (e *TagInputButton) Id(id string) (ref *TagInputButton) {
	e.id = id
	e.selfElement.Set("id", id)
	return e
}

// InputMode
//
// English:
//
//  The inputmode global attribute is an enumerated attribute that hints at the type of data that
//  might be entered by the user while editing the element or its contents. This allows a browser to
//  display an appropriate virtual keyboard.
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
//  O atributo global inputmode é um atributo enumerado que indica o tipo de dados que pode ser
//  inserido pelo usuário ao editar o elemento ou seu conteúdo. Isso permite que um navegador exiba
//  um teclado virtual apropriado.
//
// Ele é usado principalmente em elementos <input>, mas pode ser usado em qualquer elemento no modo
// contenteditable.
//
// É importante entender que o atributo inputmode não faz com que nenhum requisito de validade seja
// imposto na entrada. Para exigir que a entrada esteja em conformidade com um tipo de dados
// específico, escolha um tipo de elemento <input> apropriado. Para obter orientações específicas
// sobre como escolher os tipos de <input>, consulte a seção Valores.
func (e *TagInputButton) InputMode(inputMode InputMode) (ref *TagInputButton) {
	e.selfElement.Set("inputmode", inputMode.String())
	return e
}

// Is
//
// English:
//
//  Allows you to specify that a standard HTML element should behave like a registered custom
//  built-in element.
//
// Português:
//
//  Permite especificar que um elemento HTML padrão deve se comportar como um elemento interno
//  personalizado registrado.
func (e *TagInputButton) Is(is string) (ref *TagInputButton) {
	e.selfElement.Set("is", is)
	return e
}

// ItemId
//
// English:
//
//  The unique, global identifier of an item.
//
// Português:
//
//  O identificador global exclusivo de um item.
func (e *TagInputButton) ItemId(id string) (ref *TagInputButton) {
	e.selfElement.Set("itemid", id)
	return e
}

// ItemDrop
//
// English:
//
//  Used to add properties to an item. Every HTML element may have an itemprop attribute specified,
//  where an itemprop consists of a name and value pair.
//
// Português:
//
//  Usado para adicionar propriedades a um item. Cada elemento HTML pode ter um atributo itemprop
//  especificado, onde um itemprop consiste em um par de nome e valor.
func (e *TagInputButton) ItemDrop(itemprop string) (ref *TagInputButton) {
	e.selfElement.Set("itemprop", itemprop)
	return e
}

// ItemRef
//
// English:
//
//  Properties that are not descendants of an element with the itemscope attribute can be associated
//  with the item using an itemref. It provides a list of element ids (not itemids) with additional
//  properties elsewhere in the document.
//
// Português:
//
//  Propriedades que não são descendentes de um elemento com o atributo itemscope podem ser
//  associadas ao item usando um itemref. Ele fornece uma lista de IDs de elementos (não IDs de itens)
//  com propriedades adicionais em outras partes do documento.
func (e *TagInputButton) ItemRef(itemref string) (ref *TagInputButton) {
	e.selfElement.Set("itemref", itemref)
	return e
}

// ItemType
//
// English:
//
//  Specifies the URL of the vocabulary that will be used to define itemprops (item properties) in
//  the data structure. itemscope is used to set the scope of where in the data structure the
//  vocabulary set by itemtype will be active.
//
// Português:
//
//  Especifica a URL do vocabulário que será usado para definir itemprops (propriedades do item) na
//  estrutura de dados. itemscope é usado para definir o escopo de onde na estrutura de dados o
//  vocabulário definido por tipo de item estará ativo.
func (e *TagInputButton) ItemType(itemType string) (ref *TagInputButton) {
	e.selfElement.Set("itemtype", itemType)
	return e
}

// Lang
//
// English:
//
//  Specifies the language of the element's content.
//
// The lang attribute specifies the language of the element's content.
//
// Common examples are KLanguageEnglish for English, KLanguageSpanish for Spanish, KLanguageFrench
// for French, and so on.
//
// Português:
//
//  Especifica o idioma do conteúdo do elemento.
//
// O atributo lang especifica o idioma do conteúdo do elemento.
//
// Exemplos comuns são KLanguageEnglish para inglês, KLanguageSpanish para espanhol, KLanguageFrench
// para francês e assim por diante.
func (e *TagInputButton) Lang(language Language) (ref *TagInputButton) {
	e.selfElement.Set("lang", language.String())
	return e
}

// Nonce
//
// English:
//
//  A space-separated list of the part names of the element. Part names allows CSS to select and style
//  specific elements in a shadow tree via the ::part pseudo-element.
//
// Português:
//
//  Uma lista separada por espaços dos nomes das partes do elemento. Os nomes das partes permitem que
//  o CSS selecione e estilize elementos específicos em uma árvore de sombra por meio do
//  pseudo-elemento ::part.
func (e *TagInputButton) Nonce(part ...string) (ref *TagInputButton) {
	e.selfElement.Set("part", strings.Join(part, " "))
	return e
}

// Slot
//
// English:
//
//  Assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is
//  assigned to the slot created by the <slot> element whose name attribute's value matches that slot
//  attribute's value.
//
// Português:
//
//  Atribui um slot em uma shadow DOM shadow tree a um elemento: Um elemento com um atributo slot é
//  atribuído ao slot criado pelo elemento <slot> cujo valor do atributo name corresponde ao valor
//  desse atributo slot.
func (e *TagInputButton) Slot(slot string) (ref *TagInputButton) {
	e.selfElement.Set("slot", slot)
	return e
}

// Spellcheck
//
// English:
//
//  Specifies whether the element is to have its spelling and grammar checked or not
//
//   Note:
//     * The following can be spellchecked:
//         Text values in input elements (not password)
//         Text in <textarea> elements
//         Text in editable elements
//
// Português:
//
//  Especifica se o elemento deve ter sua ortografia e gramática verificadas ou não
//
// O seguinte pode ser verificado ortográfico:
//
//   Nota:
//     * O seguinte pode ser verificado ortográfico:
//         Valores de texto em elementos de entrada (não senha)
//         Texto em elementos <textarea>
//         Texto em elementos editáveis
func (e *TagInputButton) Spellcheck(spell bool) (ref *TagInputButton) {
	e.selfElement.Set("spellcheck", spell)

	return e
}

// Style
//
// English:
//
//  Specifies an inline CSS style for an element.
//
// The style attribute will override any style set globally, e.g. styles specified in the <style> tag
// or in an external style sheet.
//
// The style attribute can be used on any HTML element (it will validate on any HTML element.
// However, it is not necessarily useful).
//
// Português:
//
//  Especifica um estilo CSS embutido para um elemento
//
// O atributo style substituirá qualquer conjunto de estilos globalmente, por exemplo estilos
// especificados na tag <style> ou em uma folha de estilo externa.
//
// O atributo style pode ser usado em qualquer elemento HTML (vai validar em qualquer elemento HTML.
// No entanto, não é necessariamente útil).
func (e *TagInputButton) Style(style string) (ref *TagInputButton) {
	e.selfElement.Set("style", style)
	return e
}

// TabIndex
//
// English:
//
//  Specifies the tabbing order of an element (when the "tab" button is used for navigating).
//
// The tabindex attribute can be used on any HTML element (it will validate on any HTML element.
// However, it is not necessarily useful).
//
// Português:
//
//  Especifica a ordem de tabulação de um elemento (quando o botão "tab" é usado para navegar).
//
// O atributo tabindex pode ser usado em qualquer elemento HTML (vai validar em qualquer elemento
// HTML. No entanto, não é necessariamente útil).
func (e *TagInputButton) TabIndex(index int) (ref *TagInputButton) {
	e.selfElement.Set("tabindex", index)
	return e
}

// Title
//
// English:
//
//  Specifies extra information about an element.
//
// The information is most often shown as a tooltip text when the mouse moves over the element.
//
// The title attribute can be used on any HTML element (it will validate on any HTML element.
// However, it is not necessarily useful).
//
// Português:
//
//  Especifica informações extras sobre um elemento.
//
// As informações geralmente são mostradas como um texto de dica de ferramenta quando o mouse se move
// sobre o elemento.
//
// O atributo title pode ser usado em qualquer elemento HTML (vai validar em qualquer elemento HTML.
// No entanto, não é necessariamente útil).
func (e *TagInputButton) Title(title string) (ref *TagInputButton) {
	e.selfElement.Set("title", title)
	return e
}

// Translate
//
// English:
//
//  Specifies whether the content of an element should be translated or not.
//
//   Input:
//     translate: element should be translated or not. [ KTranslateYes | KTranslateNo ]
//
// English:
//
//  Especifica se o conteúdo de um elemento deve ser traduzido ou não.
//
//   Entrada:
//     translate: elemento deve ser traduzido ou não. [ KTranslateYes | KTranslateNo ]
func (e *TagInputButton) Translate(translate Translate) (ref *TagInputButton) {
	e.selfElement.Set("translate", translate.String())
	return e
}

// CreateElement
//
// English:
//
//  In an HTML document, the Document.createElement() method creates the specified HTML element or an
//  HTMLUnknownElement if the given element name is not known.
//
// Português:
//
//  Em um documento HTML, o método Document.createElement() cria o elemento HTML especificado ou um
//  HTMLUnknownElement se o nome do elemento dado não for conhecido.
func (e *TagInputButton) CreateElement(tag Tag) (ref *TagInputButton) {
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
//  Adds a node to the end of the list of children of a specified parent node. If the node already
//  exists in the document, it is removed from its current parent node before being added to the
//  new parent.
//
//   Input:
//     appendId: id of parent element.
//
//   Note:
//     * The equivalent of:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
//
// Português:
//
//  Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no
//  documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//   Entrada:
//     appendId: id do elemento pai.
//
//   Nota:
//     * Equivale a:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
func (e *TagInputButton) AppendById(appendId string) (ref *TagInputButton) {

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
//  Adds a node to the end of the list of children of a specified parent node. If the node already
//  exists in the document, it is removed from its current parent node before being added to the new
//  parent.
//
//   Input:
//     append: element in js.Value format.
//
//   Note:
//     * The equivalent of:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
//
// Português:
//
//  Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no
//  documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//   Entrada:
//     appendId: elemento no formato js.Value.
//
//   Nota:
//     * Equivale a:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
func (e *TagInputButton) Append(append interface{}) (ref *TagInputButton) {
	switch append.(type) {
	case *TagInputButton:
		e.selfElement.Call("appendChild", append.(*TagInputButton).selfElement)
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

// Autocomplete
//
// English:
//
//  The HTML autocomplete attribute lets web developers specify what if any permission the user agent
//  has to provide automated assistance in filling out form field values, as well as guidance to the
//  browser as to the type of information expected in the field.
//
// It is available on <input> elements that take a text or numeric value as input, <textarea>
// elements, <select> elements, and <form> elements.
//
// The source of the suggested values is generally up to the browser; typically values come from past
// values entered by the user, but they may also come from pre-configured values. For instance, a
// browser might let the user save their name, address, phone number, and email addresses for
// autocomplete purposes. Perhaps the browser offers the ability to save encrypted credit card
// information, for autocompletion following an authentication procedure.
//
// If an <input>, <select> or <textarea> element has no autocomplete attribute, then browsers use the
// autocomplete attribute of the element's form owner, which is either the <form> element that the
// element is a descendant of, or the <form> whose id is specified by the form attribute of the
// element.
//
//   Note:
//     * In order to provide autocompletion, user-agents might require <input>/<select>/<textarea>
//       elements to:
//         Have a name and/or id attribute;
//         Be descendants of a <form> element;
//         The form to have a submit button.
//
// Português:
//
//  O atributo autocomplete HTML permite que os desenvolvedores da Web especifiquem se existe alguma
//  permissão que o agente do usuário tenha para fornecer assistência automatizada no preenchimento
//  dos valores dos campos do formulário, bem como orientação ao navegador quanto ao tipo de
//  informação esperado no campo.
//
// Ele está disponível em elementos <input> que recebem um texto ou valor numérico como entrada,
// elementos <textarea>, elementos <select> e elementos <form>.
//
// A origem dos valores sugeridos geralmente depende do navegador; normalmente os valores vêm de
// valores passados inseridos pelo usuário, mas também podem vir de valores pré-configurados.
// Por exemplo, um navegador pode permitir que o usuário salve seu nome, endereço, número de telefone
// e endereços de e-mail para fins de preenchimento automático. Talvez o navegador ofereça a
// capacidade de salvar informações de cartão de crédito criptografadas, para preenchimento automático
// após um procedimento de autenticação.
//
// Se um elemento <input>, <select> ou <textarea> não tiver um atributo autocomplete, os navegadores
// usarão o atributo autocomplete do proprietário do formulário do elemento, que é o elemento <form>
// do qual o elemento é descendente ou o < form> cujo id é especificado pelo atributo form do
// elemento.
//
//   Nota:
//     * Para fornecer preenchimento automático, os agentes do usuário podem exigir elementos
//       <input> / <select> / <textarea> para:
//         Ter um atributo name e ou id;
//         Ser descendentes de um elemento <form>;
//         O formulário para ter um botão de envio.
func (e *TagInputButton) Autocomplete(autocomplete Autocomplete) (ref *TagInputButton) {
	e.selfElement.Set("autocomplete", autocomplete.String())
	return e
}

// Disabled
//
// English:
//
//  Este atributo booleano impede que o usuário interaja com o elemento.
//
// Português:
//
//  Este atributo booleano impede que o usuário interaja com o elemento.
func (e *TagInputButton) Disabled(disabled bool) (ref *TagInputButton) {
	e.selfElement.Set("disabled", disabled)
	return e
}

// Form
//
// English:
//
//  The <form> element to associate the button with (its form owner). The value of this attribute must
//  be the id of a <form> in the same document. (If this attribute is not set, the <button> is
//  associated with its ancestor <form> element, if any.)
//
// This attribute lets you associate <button> elements to <form>s anywhere in the document, not just
// inside a <form>. It can also override an ancestor <form> element.
//
// Português:
//
//  O elemento <form> ao qual associar o botão (seu proprietário do formulário). O valor deste
//  atributo deve ser o id de um <form> no mesmo documento. (Se esse atributo não for definido, o
//  <button> será associado ao elemento <form> ancestral, se houver.)
//
// Este atributo permite associar elementos <button> a <form>s em qualquer lugar do documento, não
// apenas dentro de um <form>. Ele também pode substituir um elemento <form> ancestral.
func (e *TagInputButton) Form(form string) (ref *TagInputButton) {
	e.selfElement.Set("form", form)
	return e
}

// List
//
// English:
//
//  The value given to the list attribute should be the id of a <datalist> element located in the same
//  document.
//
// The <datalist> provides a list of predefined values to suggest to the user for this input.
// Any values in the list that are not compatible with the type are not included in the suggested
// options. The values provided are suggestions, not requirements: users can select from this
// predefined list or provide a different value.
//
// It is valid on text, search, url, tel, email, date, month, week, time, datetime-local, number,
// range, and color.
//
// Per the specifications, the list attribute is not supported by the hidden, password, checkbox,
// radio, file, or any of the button types.
//
// Depending on the browser, the user may see a custom color palette suggested, tic marks along a
// range, or even a input that opens like a <select> but allows for non-listed values.
// Check out the browser compatibility table for the other input types.
//
// See factoryBrowser.NewTagDataList()
//
// Português:
//
//  O valor dado ao atributo list deve ser o id de um elemento <datalist> localizado no mesmo
//  documento.
//
// O <datalist> fornece uma lista de valores predefinidos para sugerir ao usuário para esta entrada.
// Quaisquer valores na lista que não sejam compatíveis com o tipo não são incluídos nas opções
// sugeridas.
// Os valores fornecidos são sugestões, não requisitos: os usuários podem selecionar dessa lista
// predefinida ou fornecer um valor diferente.
//
// É válido em texto, pesquisa, url, telefone, email, data, mês, semana, hora, data e hora local,
// número, intervalo e cor.
//
// De acordo com as especificações, o atributo de lista não é suportado pelo oculto, senha, caixa de
// seleção, rádio, arquivo ou qualquer um dos tipos de botão.
//
// Dependendo do navegador, o usuário pode ver uma paleta de cores personalizada sugerida, marcas de
// tique ao longo de um intervalo ou até mesmo uma entrada que abre como um <select>, mas permite
// valores não listados. Confira a tabela de compatibilidade do navegador para os outros tipos de
// entrada.
//
// Veja factoryBrowser.NewTagDataList()
func (e *TagInputButton) List(list string) (ref *TagInputButton) {
	e.selfElement.Set("list", list)
	return e
}

// Name
//
// English:
//
//  The name of the button, submitted as a pair with the button's value as part of the form data,
//  when that button is used to submit the form.
//
// Português:
//
//  O nome do botão, enviado como um par com o valor do botão como parte dos dados do formulário,
//  quando esse botão é usado para enviar o formulário.
func (e *TagInputButton) Name(name string) (ref *TagInputButton) {
	e.selfElement.Set("name", name)
	return e
}

// ReadOnly
//
// English:
//
//  A Boolean attribute which, if present, indicates that the user should not be able to edit the
//  value of the input.
//
// The readonly attribute is supported by the text, search, url, tel, email, date, month, week, time,
// datetime-local, number, and password input types.
//
// Português:
//
//  A Boolean attribute which, if present, indicates that the user should not be able to edit the value of the input. The readonly attribute is supported by the text, search, url, tel, email, date, month, week, time, datetime-local, number, and password input types.
func (e *TagInputButton) ReadOnly(readonly bool) (ref *TagInputButton) {
	e.selfElement.Set("readonly", readonly)
	return e
}

// Required
//
// English:
//
//  A Boolean attribute indicating that an option with a non-empty string value must be selected.
//
// Português:
//
//  Um atributo booleano que indica que uma opção com um valor de string não vazio deve ser
//  selecionada.
func (e *TagInputButton) Required(required bool) (ref *TagInputButton) {
	e.selfElement.Set("required", required)
	return e
}

// Type
//
// English:
//
//  How an <input> works varies considerably depending on the value of its type attribute, hence the
//  different types are covered in their own separate reference pages.
//
// If this attribute is not specified, the default type adopted is text.
//
// Português:
//
//  Como um <input> funciona varia consideravelmente dependendo do valor de seu atributo type,
//  portanto, os diferentes tipos são abordados em suas próprias páginas de referência separadas.
//
// Se este atributo não for especificado, o tipo padrão adotado é texto.
func (e *TagInputButton) Type(inputType InputType) (ref *TagInputButton) {
	e.selfElement.Set("type", inputType.String())
	return e
}

// Value
//
// English:
//
//  Defines the value associated with the element.
//
// Português:
//
//  Define o valor associado ao elemento.
func (e *TagInputButton) Value(value string) (ref *TagInputButton) {
	e.selfElement.Set("value", value)
	return e
}

// SetXY
//
// English:
//
//  Sets the X and Y axes in pixels.
//
// Português:
//
//  Define os eixos X e Y em pixels.
func (e *TagInputButton) SetXY(x, y int) (ref *TagInputButton) {
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
//  Sets the X axe in pixels.
//
// Português:
//
//  Define o eixo X em pixels.
func (e *TagInputButton) SetX(x int) (ref *TagInputButton) {
	px := strconv.FormatInt(int64(x), 10) + "px"
	e.selfElement.Get("style").Set("left", px)

	return e
}

// SetY
//
// English:
//
//  Sets the Y axe in pixels.
//
// Português:
//
//  Define o eixo Y em pixels.
func (e *TagInputButton) SetY(y int) (ref *TagInputButton) {
	py := strconv.FormatInt(int64(y), 10) + "px"
	e.selfElement.Get("style").Set("top", py)

	return e
}

// GetXY
//
// English:
//
//  Returns the X and Y axes in pixels.
//
// Português:
//
//  Retorna os eixos X e Y em pixels.
func (e *TagInputButton) GetXY() (x, y int) {
	x = e.selfElement.Get("style").Get("left").Int()
	y = e.selfElement.Get("style").Get("top").Int()

	return
}

// GetX
//
// English:
//
//  Returns the X axe in pixels.
//
// Português:
//
//  Retorna o eixo X em pixels.
func (e *TagInputButton) GetX() (x int) {
	x = e.selfElement.Get("style").Get("left").Int()

	return
}

// GetY
//
// English:
//
//  Returns the Y axe in pixels.
//
// Português:
//
//  Retorna o eixo Y em pixels.
func (e *TagInputButton) GetY() (y int) {
	y = e.selfElement.Get("style").Get("top").Int()

	return
}
