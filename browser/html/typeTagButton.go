package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"log"
	"strconv"
	"strings"
	"syscall/js"
)

// TagButton
//
// English:
//
//  The <button> HTML element is an interactive element activated by a user with a mouse, keyboard,
//  finger, voice command, or other assistive technology. Once activated, it then performs a
//  programmable action, such as submitting a form or opening a dialog.
//
// By default, HTML buttons are presented in a style resembling the platform the user agent runs on,
// but you can change buttons' appearance with CSS.
//
// Português
//
//  O elemento HTML <button> é um elemento interativo ativado por um usuário com mouse, teclado,
//  dedo, comando de voz ou outra tecnologia assistiva. Uma vez ativado, ele executa uma ação
//  programável, como enviar um formulário ou abrir uma caixa de diálogo.
//
// Por padrão, os botões HTML são apresentados em um estilo semelhante à plataforma na qual o agente
// do usuário é executado, mas você pode alterar a aparência dos botões com CSS.
type TagButton struct {
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
func (e *TagButton) AccessKey(key string) (ref *TagButton) {
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
func (e *TagButton) Autofocus(autofocus bool) (ref *TagButton) {
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
func (e *TagButton) Class(class ...string) (ref *TagButton) {
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
func (e *TagButton) ContentEditable(editable bool) (ref *TagButton) {
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
func (e *TagButton) Data(data map[string]string) (ref *TagButton) {
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
func (e *TagButton) Dir(dir Dir) (ref *TagButton) {
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
func (e *TagButton) Draggable(draggable Draggable) (ref *TagButton) {
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
func (e *TagButton) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagButton) {
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
func (e *TagButton) Hidden() (ref *TagButton) {
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
func (e *TagButton) Id(id string) (ref *TagButton) {
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
func (e *TagButton) InputMode(inputMode InputMode) (ref *TagButton) {
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
func (e *TagButton) Is(is string) (ref *TagButton) {
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
func (e *TagButton) ItemId(id string) (ref *TagButton) {
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
func (e *TagButton) ItemDrop(itemprop string) (ref *TagButton) {
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
func (e *TagButton) ItemRef(itemref string) (ref *TagButton) {
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
func (e *TagButton) ItemType(itemType string) (ref *TagButton) {
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
func (e *TagButton) Lang(language Language) (ref *TagButton) {
	e.selfElement.Set("lang", language.String())
	return e
}

// Nonce
//
// English:
//
//  A cryptographic nonce ("number used once") which can be used by Content Security Policy to
//  determine whether or not a given fetch will be allowed to proceed.
//
// Português:
//
//  Um nonce criptográfico ("número usado uma vez") que pode ser usado pela Política de Segurança de
//  Conteúdo para determinar se uma determinada busca terá permissão para prosseguir.
func (e *TagButton) Nonce(nonce int64) (ref *TagButton) {
	e.selfElement.Set("nonce", nonce)
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
func (e *TagButton) Slot(slot string) (ref *TagButton) {
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
func (e *TagButton) Spellcheck(spell bool) (ref *TagButton) {
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
func (e *TagButton) Style(style string) (ref *TagButton) {
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
func (e *TagButton) TabIndex(index int) (ref *TagButton) {
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
func (e *TagButton) Title(title string) (ref *TagButton) {
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
func (e *TagButton) Translate(translate Translate) (ref *TagButton) {
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
func (e *TagButton) CreateElement(tag Tag) (ref *TagButton) {
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
func (e *TagButton) AppendById(appendId string) (ref *TagButton) {

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
func (e *TagButton) Append(append interface{}) (ref *TagButton) {
	switch append.(type) {
	case *TagButton:
		e.selfElement.Call("appendChild", append.(*TagButton).selfElement)
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
func (e *TagButton) Autocomplete(autocomplete Autocomplete) (ref *TagButton) {
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
func (e *TagButton) Disabled(disabled bool) (ref *TagButton) {
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
func (e *TagButton) Form(form string) (ref *TagButton) {
	e.selfElement.Set("form", form)
	return e
}

// FromAction
//
// English:
//
//  The URL that processes the information submitted by the button. Overrides the action attribute of
//  the button's form owner. Does nothing if there is no form owner.
//
//   Input:
//     action: The URL that processes the form submission. This value can be overridden by a
//             formaction attribute on a <button>, <input type="submit">, or <input type="image">
//             element. This attribute is ignored when method="dialog" is set.
//
// Português:
//
//  A URL que processa as informações enviadas pelo botão. Substitui o atributo de ação do
//  proprietário do formulário do botão. Não faz nada se não houver um proprietário de formulário.
//
//   Entrada:
//     action: A URL que processa o envio do formulário. Esse valor pode ser substituído por um
//             atributo formaction em um elemento <button>, <input type="submit"> ou
//             <input type="image">. Este atributo é ignorado quando method="dialog" é definido.
func (e *TagButton) FromAction(action string) (ref *TagButton) {
	e.selfElement.Set("formaction", action)
	return e
}

// FormEncType
//
// English:
//
//  If the button is a submit button (it's inside/associated with a <form> and doesn't have
//  type="button"), specifies how to encode the form data that is submitted. Possible values:
//
//   Input:
//     formenctype: specifies how to encode the form data
//
//       application/x-www-form-urlencoded: The default if the attribute is not used.
//       multipart/form-data: Use to submit <input> elements with their type attributes set to file.
//       text/plain: Specified as a debugging aid; shouldn't be used for real form submission.
//
//   Note:
//     * If this attribute is specified, it overrides the enctype attribute of the button's form
//       owner.
//
// Português:
//
//  Se o botão for um botão de envio (está associado a um <form> e não possui type="button"),
//  especifica como codificar os dados do formulário que são enviados. Valores possíveis:
//
//   Entrada:
//     formenctype: especifica como codificar os dados do formulário
//
//       KFormEncTypeApplication: O padrão se o atributo não for usado.
//       KFormEncTypeMultiPart: Use para enviar elementos <input> com seus atributos de tipo definidos
//         para arquivo.
//       KFormEncTypeText: Especificado como auxiliar de depuração; não deve ser usado para envio de
//         formulário real.
//
//   Note:
//     * Se este atributo for especificado, ele substituirá o atributo enctype do proprietário do
//       formulário do botão.
func (e *TagButton) FormEncType(formenctype FormEncType) (ref *TagButton) {
	e.selfElement.Set("formenctype", formenctype.String())
	return e
}

// FormMethod
//
// English:
//
//  If the button is a submit button (it's inside/associated with a <form> and doesn't have
//  type="button"), this attribute specifies the HTTP method used to submit the form.
//
//   Input:
//     method: specifies the HTTP method used to submit
//       KFormMethodPost: The data from the form are included in the body of the HTTP request when
//         sent to the server. Use when the form contains information that shouldn't be public, like
//         login credentials.
//       KFormMethodGet: The form data are appended to the form's action URL, with a ? as a separator,
//         and the resulting URL is sent to the server. Use this method when the form has no side
//         effects, like search forms.
//
//   Note:
//     * If specified, this attribute overrides the method attribute of the button's form owner.
//
// Português:
//
//  Se o botão for um botão de envio (está associado a um <form> e não possui type="button"),
//  este atributo especifica o método HTTP usado para enviar o formulário.
//
//   Input:
//     method: especifica o método HTTP usado para enviar
//       KFormMethodPost: Os dados do formulário são incluídos no corpo da solicitação HTTP quando
//         enviados ao servidor. Use quando o formulário contém informações que não devem ser
//         públicas, como credenciais de login.
//       KFormMethodGet: Os dados do formulário são anexados à URL de ação do formulário, com um ?
//         como separador e a URL resultante é enviada ao servidor. Use este método quando o
//         formulário não tiver efeitos colaterais, como formulários de pesquisa.
//
//   Nota:
//     * Se especificado, este atributo substitui o atributo method do proprietário do formulário do
//       botão.
func (e *TagButton) FormMethod(method FormMethod) (ref *TagButton) {
	e.selfElement.Set("formmethod", method.String())
	return e
}

// FormValidate
//
// English:
//
//  If the button is a submit button, this Boolean attribute specifies that the form is not to be
//  validated when it is submitted.
//
// If this attribute is specified, it overrides the novalidate attribute of the button's form owner.
//
// Português:
//
//  Se o botão for um botão de envio, este atributo booleano especifica que o formulário não deve ser
//  validado quando for enviado.
//
// Se este atributo for especificado, ele substituirá o atributo novalidate do proprietário do
// formulário do botão.
func (e *TagButton) FormValidate(validate bool) (ref *TagButton) {
	e.selfElement.Set("formnovalidate", validate)
	return e
}

// FormTarget
//
// English:
//
//  If the button is a submit button, this attribute is an author-defined name or standardized,
//  underscore-prefixed keyword indicating where to display the response from submitting the form.
//
// This is the name of, or keyword for, a browsing context (a tab, window, or <iframe>).
//
// If this attribute is specified, it overrides the target attribute of the button's form owner.
// The following keywords have special meanings:
//
//   KTargetSelf: the current browsing context; (Default)
//   KTargetBlank: usually a new tab, but users can configure browsers to open a new window instead;
//   KTargetParent: the parent browsing context of the current one. If no parent, behaves as _self;
//   KTargetTop: the topmost browsing context (the "highest" context that's an ancestor of the current
//     one). If no ancestors, behaves as _self.
//
// Português:
//
//  Se o botão for um botão de envio, esse atributo será um nome definido pelo autor ou uma
//  palavra-chave padronizada com prefixo de sublinhado indicando onde exibir a resposta do envio do
//  formulário.
//
// Este é o nome ou a palavra-chave de um contexto de navegação (uma guia, janela ou <iframe>).
// Se este atributo for especificado, ele substituirá o atributo de destino do proprietário do
// formulário do botão.
// As seguintes palavras-chave têm significados especiais:
//
//   KTargetSelf: o contexto de navegação atual; (padrão)
//   KTargetBlank: geralmente uma nova guia, mas os usuários podem configurar os navegadores para
//     abrir uma nova janela;
//   KTargetParent: o contexto de navegação pai do atual. Se nenhum pai, se comporta como _self;
//   KTargetTop: o contexto de navegação mais alto (o contexto "mais alto" que é um ancestral do
//     atual). Se não houver ancestrais, se comporta como _self.
func (e *TagButton) FormTarget(formtarget Target) (ref *TagButton) {
	e.selfElement.Set("formtarget", formtarget.String())
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
func (e *TagButton) Name(name string) (ref *TagButton) {
	e.selfElement.Set("name", name)
	return e
}

// ButtonType
//
// English:
//
//  The default behavior of the button.
//
//   Input:
//     value: default behavior of the button.
//       KButtonTypeSubmit: The button submits the form data to the server. This is the default if
//         the attribute is not specified for buttons associated with a <form>, or if the attribute
//         is an empty or invalid value.
//       KButtonTypeReset:  The button resets all the controls to their initial values, like
//         <input type="reset">. (This behavior tends to annoy users.)
//       KButtonTypeButton: The button has no default behavior, and does nothing when pressed by
//         default. It can have client-side scripts listen to the element's events, which are
//         triggered when the events occur.
//
// Português:
//
//  O comportamento padrão do botão. Os valores possíveis são:
//
//   Entrada:
//     value: comportamento padrão do botão
//       KButtonTypeSubmit: O botão envia os dados do formulário para o servidor. Este é o padrão se
//         o atributo não for especificado para botões associados a um <form> ou se o atributo for um
//         valor vazio ou inválido.
//       KButtonTypeReset:  O botão redefine todos os controles para seus valores iniciais, como
//         <input type="reset">. (Esse comportamento tende a incomodar os usuários.)
//       KButtonTypeButton: O botão não tem comportamento padrão e não faz nada quando pressionado por
//         padrão. Ele pode fazer com que os scripts do lado do cliente escutem os eventos do
//         elemento, que são acionados quando os eventos ocorrem.
func (e *TagButton) ButtonType(value ButtonType) (ref *TagButton) {
	e.selfElement.Set("type", value.String())
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
func (e *TagButton) Value(value string) (ref *TagButton) {
	e.selfElement.Set("value", value)
	e.selfElement.Set("innerText", value)
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
func (e *TagButton) SetXY(x, y int) (ref *TagButton) {
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
func (e *TagButton) SetX(x int) (ref *TagButton) {
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
func (e *TagButton) SetY(y int) (ref *TagButton) {
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
func (e *TagButton) GetXY() (x, y int) {
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
func (e *TagButton) GetX() (x int) {
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
func (e *TagButton) GetY() (y int) {
	y = e.selfElement.Get("style").Get("top").Int()

	return
}
