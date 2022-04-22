package html

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/css"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/globalDocument"
	"log"
	"strings"
	"syscall/js"
)

type TagInput struct {
	tag         Tag
	id          string
	selfElement js.Value
	document    globalDocument.Document
	cursor      browserMouse.CursorType
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
func (e *TagInput) AccessKey(key string) (ref *TagInput) {
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
func (e *TagInput) Autofocus(autofocus bool) (ref *TagInput) {
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
func (e *TagInput) Class(class ...string) (ref *TagInput) {
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
func (e *TagInput) ContentEditable(editable bool) (ref *TagInput) {
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
func (e *TagInput) Data(data map[string]string) (ref *TagInput) {
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
func (e *TagInput) Dir(dir Dir) (ref *TagInput) {
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
func (e *TagInput) Draggable(draggable Draggable) (ref *TagInput) {
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
func (e *TagInput) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagInput) {
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
func (e *TagInput) Hidden() (ref *TagInput) {
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
func (e *TagInput) Id(id string) (ref *TagInput) {
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
func (e *TagInput) InputMode(inputMode InputMode) (ref *TagInput) {
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
func (e *TagInput) Is(is string) (ref *TagInput) {
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
func (e *TagInput) ItemId(id string) (ref *TagInput) {
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
func (e *TagInput) ItemDrop(itemprop string) (ref *TagInput) {
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
func (e *TagInput) ItemRef(itemref string) (ref *TagInput) {
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
func (e *TagInput) ItemType(itemType string) (ref *TagInput) {
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
func (e *TagInput) Lang(language Language) (ref *TagInput) {
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
func (e *TagInput) Nonce(part ...string) (ref *TagInput) {
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
func (e *TagInput) Slot(slot string) (ref *TagInput) {
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
func (e *TagInput) Spellcheck(spell bool) (ref *TagInput) {
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
func (e *TagInput) Style(style string) (ref *TagInput) {
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
func (e *TagInput) TabIndex(index int) (ref *TagInput) {
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
func (e *TagInput) Title(title string) (ref *TagInput) {
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
func (e *TagInput) Translate(translate Translate) (ref *TagInput) {
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
func (e *TagInput) CreateElement(tag Tag) (ref *TagInput) {
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
func (e *TagInput) AppendById(appendId string) (ref *TagInput) {
	
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
func (e *TagInput) Append(append interface{}) (ref *TagInput) {
	switch append.(type) {
	case *TagInput:
		e.selfElement.Call("appendChild", append.(*TagInput).selfElement)
	case js.Value:
		e.selfElement.Call("appendChild", append)
	}
	
	return e
}

// Accept
//
// English:
//
//  Valid for the file input type only, the accept attribute defines which file types are selectable
//  in a file upload control. See the file input type.
//
// Português:
//
//  Válido apenas para o tipo de entrada de arquivo, o atributo accept define quais tipos de arquivo
//  são selecionáveis em um controle de upload de arquivo. Consulte o tipo de entrada do arquivo.
func (e *TagInput) Accept(accept string) (ref *TagInput) {
	e.selfElement.Set("accept", accept)
	return e
}

// Alt
//
// English:
//
//  The alt attribute provides alternative text for the image, displaying the value of the attribute
//  if the image src is missing or otherwise fails to load.
//
// Português:
//
//  O atributo alt fornece texto alternativo para a imagem, exibindo o valor do atributo se o src da
//  imagem estiver ausente ou falhar ao carregar.
func (e *TagInput) Alt(alt string) (ref *TagInput) {
	e.selfElement.Set("alt", alt)
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
func (e *TagInput) Autocomplete(autocomplete Autocomplete) (ref *TagInput) {
	e.selfElement.Set("autocomplete", autocomplete.String())
	return e
}

// Capture
//
// English:
//
//  Introduced in the HTML Media Capture specification and valid for the file input type only, the
//  capture attribute defines which media—microphone, video, or camera—should be used to capture a
//  new file for upload with file upload control in supporting scenarios.
//
// Português:
//
//  Introduzido na especificação HTML Media Capture e válido apenas para o tipo de entrada de arquivo,
//  o atributo capture define qual mídia—microfone, vídeo ou câmera—deve ser usada para capturar um
//  novo arquivo para upload com controle de upload de arquivo em cenários de suporte.
func (e *TagInput) Capture(capture string) (ref *TagInput) {
	e.selfElement.Set("capture", capture)
	return e
}

// Checked
//
// English:
//
//  Valid for both radio and checkbox types, checked is a Boolean attribute. If present on a radio
//  type, it indicates that the radio button is the currently selected one in the group of same-named
//  radio buttons. If present on a checkbox type, it indicates that the checkbox is checked by default
//  (when the page loads).
//  It does not indicate whether this checkbox is currently checked: if the checkbox's state is
//  changed, this content attribute does not reflect the change.
//  (Only the HTMLInputElement's checked IDL attribute is updated.)
//
//   Note:
//     * Unlike other input controls, a checkboxes and radio buttons value are only included in the
//       submitted data if they are currently checked. If they are, the name and the value(s) of the
//       checked controls are submitted.
//       For example, if a checkbox whose name is fruit has a value of cherry, and the checkbox is
//       checked, the form data submitted will include fruit=cherry. If the checkbox isn't active,
//       it isn't listed in the form data at all. The default value for checkboxes and radio buttons
//       is on.
//
// Português:
//
//  Válido para os tipos de rádio e caixa de seleção, marcado é um atributo booleano. Se estiver
//  presente em um tipo de rádio, indica que o botão de opção é o selecionado atualmente no grupo de
//  botões de opção com o mesmo nome. Se estiver presente em um tipo de caixa de seleção, indica que
//  a caixa de seleção está marcada por padrão (quando a página é carregada). Não indica se esta caixa
//  de seleção está marcada no momento: se o estado da caixa de seleção for alterado, esse atributo
//  de conteúdo não reflete a alteração.
//  (Apenas o atributo IDL verificado do HTMLInputElement é atualizado.)
//
//   Nota:
//     * Ao contrário de outros controles de entrada, um valor de caixas de seleção e botões de opção
//       só são incluídos nos dados enviados se estiverem marcados no momento. Se estiverem, o nome e
//       o(s) valor(es) dos controles verificados são enviados.
//       Por exemplo, se uma caixa de seleção cujo nome é fruta tiver o valor cereja e a caixa de
//       seleção estiver marcada, os dados do formulário enviados incluirão fruta=cereja.
//       Se a caixa de seleção não estiver ativa, ela não está listada nos dados do formulário.
//       O valor padrão para caixas de seleção e botões de opção é ativado.
func (e *TagInput) Checked(checked bool) (ref *TagInput) {
	e.selfElement.Set("checked", checked)
	return e
}

// DirName
//
// English:
//
//  Valid for text and search input types only, the dirname attribute enables the submission of the
//  directionality of the element. When included, the form control will submit with two name/value
//  pairs: the first being the name and value, the second being the value of the dirname as the name
//  with the value of ltr or rtl being set by the browser.
//
// Português:
//
//  Válido apenas para tipos de entrada de texto e pesquisa, o atributo dirname permite o envio da
//  direcionalidade do elemento. Quando incluído, o controle de formulário será enviado com dois pares
//  nomevalor: o primeiro sendo o nome e o valor, o segundo sendo o valor do dirname como o nome com o
//  valor de ltr ou rtl sendo definido pelo navegador.
func (e *TagInput) DirName(dirname string) (ref *TagInput) {
	e.selfElement.Set("dirname", dirname)
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
func (e *TagInput) Disabled(disabled bool) (ref *TagInput) {
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
func (e *TagInput) Form(form string) (ref *TagInput) {
	e.selfElement.Set("form", form)
	return e
}

// FormAction
//
// English:
//
//  A string indicating the URL to which to submit the data. This takes precedence over the action
//  attribute on the <form> element that owns the <input>.
//
// This attribute is also available on <input type="image"> and <button> elements.
//
// Português:
//
//  Uma string indicando o URL para o qual enviar os dados. Isso tem precedência sobre o atributo
//  action no elemento <form> que possui o <input>.
//
// Este atributo também está disponível nos elementos <input type="image"> e <button>.
func (e *TagInput) FormAction(action string) (ref *TagInput) {
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
func (e *TagInput) FormEncType(formenctype FormEncType) (ref *TagInput) {
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
func (e *TagInput) FormMethod(method FormMethod) (ref *TagInput) {
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
func (e *TagInput) FormValidate(validate bool) (ref *TagInput) {
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
func (e *TagInput) FormTarget(formtarget Target) (ref *TagInput) {
	e.selfElement.Set("formtarget", formtarget.String())
	return e
}

// Height
//
// English:
//
//  The height is the height of the image file to display to represent the graphical submit button.
//
// Português:
//
//  A altura é a altura do arquivo de imagem a ser exibido para representar o botão de envio gráfico.
func (e *TagInput) Height(height int) (ref *TagInput) {
	e.selfElement.Set("height", height)
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
func (e *TagInput) List(list string) (ref *TagInput) {
	e.selfElement.Set("list", list)
	return e
}

// Max
//
// English:
//
//  Valid for date, month, week, time, datetime-local, number, and range, it defines the greatest
//  value in the range of permitted values.
//  If the value entered into the element exceeds this, the element fails constraint validation.
//  If the value of the max attribute isn't a number, then the element has no maximum value.
//
// There is a special case: if the data type is periodic (such as for dates or times), the value of
// max may be lower than the value of min, which indicates that the range may wrap around;
// for example, this allows you to specify a time range from 10 PM to 4 AM.
//
// Português:
//
//  Válido para data, mês, semana, hora, datetime-local, número e intervalo, define o maior valor no
//  intervalo de valores permitidos. Se o valor inserido no elemento exceder isso, o elemento falhará
//  na validação de restrição. Se o valor do atributo max não for um número, o elemento não terá
//  valor máximo.
//
// Há um caso especial: se o tipo de dado for periódico (como para datas ou horas), o valor de max
// pode ser menor que o valor de min, o que indica que o intervalo pode ser contornado;
// por exemplo, isso permite que você especifique um intervalo de tempo das 22h às 4h.
func (e *TagInput) Max(max int) (ref *TagInput) {
	e.selfElement.Set("max", max)
	return e
}

// MaxLength
//
// English:
//
//  Valid for text, search, url, tel, email, and password, it defines the maximum number of characters
//  (as UTF-16 code units) the user can enter into the field. This must be an integer value 0 or
//  higher. If no maxlength is specified, or an invalid value is specified, the field has no maximum
//  length. This value must also be greater than or equal to the value of minlength.
//
// The input will fail constraint validation if the length of the text entered into the field is
// greater than maxlength UTF-16 code units long. By default, browsers prevent users from entering
// more characters than allowed by the maxlength attribute.
//
// Português:
//
//  Válido para texto, pesquisa, url, tel, email e senha, define o número máximo de caracteres
//  (como unidades de código UTF-16) que o usuário pode inserir no campo.
//
// Este deve ser um valor inteiro 0 ou superior. Se nenhum comprimento máximo for especificado ou um
// valor inválido for especificado, o campo não terá comprimento máximo. Esse valor também deve ser
// maior ou igual ao valor de minlength.
//
// A entrada falhará na validação de restrição se o comprimento do texto inserido no campo for maior
// que o comprimento máximo das unidades de código UTF-16. Por padrão, os navegadores impedem que os
// usuários insiram mais caracteres do que o permitido pelo atributo maxlength.
func (e *TagInput) MaxLength(maxlength int) (ref *TagInput) {
	e.selfElement.Set("maxlength", maxlength)
	return e
}

// Min
//
// English:
//
//  Valid for date, month, week, time, datetime-local, number, and range, it defines the most negative
//  value in the range of permitted values.
//
// If the value entered into the element is less than this, the element fails constraint validation.
// If the value of the min attribute isn't a number, then the element has no minimum value.
//
// This value must be less than or equal to the value of the max attribute. If the min attribute is
// present but is not specified or is invalid, no min value is applied. If the min attribute is valid
// and a non-empty value is less than the minimum allowed by the min attribute, constraint validation
// will prevent form submission. See Client-side validation for more information.
//
// There is a special case: if the data type is periodic (such as for dates or times), the value of
// max may be lower than the value of min, which indicates that the range may wrap around; for
// example, this allows you to specify a time range from 10 PM to 4 AM.
//
// Português:
//
//  Válido para data, mês, semana, hora, data e hora local, número e intervalo, define o valor mais
//  negativo no intervalo de valores permitidos.
//
// Se o valor inserido no elemento for menor que isso, o elemento falhará na validação de restrição.
// Se o valor do atributo min não for um número, o elemento não terá valor mínimo.
//
// Esse valor deve ser menor ou igual ao valor do atributo max. Se o atributo min estiver presente,
// mas não for especificado ou for inválido, nenhum valor min será aplicado. Se o atributo min for
// válido e um valor não vazio for menor que o mínimo permitido pelo atributo min, a validação de
// restrição impedirá o envio do formulário. Consulte Validação do lado do cliente para obter mais
// informações.
//
// Há um caso especial: se o tipo de dado for periódico (como para datas ou horas), o valor de max
// pode ser menor que o valor de min, o que indica que o intervalo pode ser contornado; por exemplo,
// isso permite que você especifique um intervalo de tempo das 22h às 4h.
func (e *TagInput) Min(min int) (ref *TagInput) {
	e.selfElement.Set("max", min)
	return e
}

// MinLength
//
// English:
//
//  Valid for text, search, url, tel, email, and password, it defines the minimum number of
//  characters (as UTF-16 code units) the user can enter into the entry field.
//
// This must be an non-negative integer value smaller than or equal to the value specified by
// maxlength. If no minlength is specified, or an invalid value is specified, the input has no
// minimum length.
//
// The input will fail constraint validation if the length of the text entered into the field is
// fewer than minlength UTF-16 code units long, preventing form submission.
//
// Português:
//
//  Válido para texto, pesquisa, url, tel, email e senha, define o número mínimo de caracteres
//  (como unidades de código UTF-16) que o usuário pode inserir no campo de entrada.
//
// Este deve ser um valor inteiro não negativo menor ou igual ao valor especificado por maxlength.
// Se nenhum comprimento mínimo for especificado ou um valor inválido for especificado, a entrada não
// terá comprimento mínimo.
//
// A entrada falhará na validação de restrição se o comprimento do texto inserido no campo for
// inferior a unidades de código UTF-16 de comprimento mínimo, impedindo o envio do formulário.
func (e *TagInput) MinLength(minlength int) (ref *TagInput) {
	e.selfElement.Set("minlength", minlength)
	return e
}

// Multiple
//
// English:
//
//  This Boolean attribute indicates that multiple options can be selected in the list. If it is not
//  specified, then only one option can be selected at a time. When multiple is specified, most
//  browsers will show a scrolling list box instead of a single line dropdown.
//
// Português:
//
//  Este atributo booleano indica que várias opções podem ser selecionadas na lista. Se não for
//  especificado, apenas uma opção pode ser selecionada por vez. Quando vários são especificados, a
//  maioria dos navegadores mostrará uma caixa de listagem de rolagem em vez de uma lista suspensa
//  de uma única linha.
func (e *TagInput) Multiple(multiple bool) (ref *TagInput) {
	e.selfElement.Set("multiple", multiple)
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
func (e *TagInput) Name(name string) (ref *TagInput) {
	e.selfElement.Set("name", name)
	return e
}

// Pattern
//
// English:
//
//  The pattern attribute, when specified, is a regular expression that the input's value must match
//  in order for the value to pass constraint validation. It must be a valid JavaScript regular
//  expression, as used by the RegExp type, and as documented in our guide on regular expressions;
//  the 'u' flag is specified when compiling the regular expression, so that the pattern is treated
//  as a sequence of Unicode code points, instead of as ASCII. No forward slashes should be specified
//  around the pattern text.
//
// If the pattern attribute is present but is not specified or is invalid, no regular expression is
// applied and this attribute is ignored completely. If the pattern attribute is valid and a non-empty
// value does not match the pattern, constraint validation will prevent form submission.
//
//   Note:
//     * If using the pattern attribute, inform the user about the expected format by including
//       explanatory text nearby. You can also include a title attribute to explain what the
//       requirements are to match the pattern; most browsers will display this title as a tooltip.
//       The visible explanation is required for accessibility. The tooltip is an enhancement.
//
// Português:
//
//  O atributo pattern, quando especificado, é uma expressão regular que o valor da entrada deve
//  corresponder para que o valor passe na validação de restrição. Deve ser uma expressão regular
//  JavaScript válida, conforme usada pelo tipo RegExp e conforme documentado em nosso guia sobre
//  expressões regulares; o sinalizador 'u' é especificado ao compilar a expressão regular, para que
//  o padrão seja tratado como uma sequência de pontos de código Unicode, em vez de como ASCII.
//  Nenhuma barra deve ser especificada ao redor do texto do padrão.
//
// Se o atributo pattern estiver presente, mas não for especificado ou for inválido, nenhuma
// expressão regular será aplicada e esse atributo será completamente ignorado. Se o atributo de
// padrão for válido e um valor não vazio não corresponder ao padrão, a validação de restrição
// impedirá o envio do formulário.
//
//   Nota:
//     * Se estiver usando o atributo pattern, informe o usuário sobre o formato esperado incluindo
//       um texto explicativo próximo. Você também pode incluir um atributo title para explicar quais
//       são os requisitos para corresponder ao padrão; a maioria dos navegadores exibirá este título
//       como uma dica de ferramenta. A explicação visível é necessária para acessibilidade. A dica
//       de ferramenta é um aprimoramento.
func (e *TagInput) Pattern(pattern string) (ref *TagInput) {
	e.selfElement.Set("pattern", pattern)
	return e
}

// Placeholder
//
// English:
//
//  The placeholder attribute is a string that provides a brief hint to the user as to what kind of
//  information is expected in the field. It should be a word or short phrase that provides a hint
//  as to the expected type of data, rather than an explanation or prompt. The text must not include
//  carriage returns or line feeds. So for example if a field is expected to capture a user's first
//  name, and its label is "First Name", a suitable placeholder might be "e.g. Mustafa".
//
//   Note:
//     * The placeholder attribute is not as semantically useful as other ways to explain your form,
//       and can cause unexpected technical issues with your content. See Labels for more information.
//
// Português:
//
//  O atributo placeholder é uma string que fornece uma breve dica ao usuário sobre que tipo de
//  informação é esperada no campo. Deve ser uma palavra ou frase curta que forneça uma dica sobre o
//  tipo de dados esperado, em vez de uma explicação ou prompt. O texto não deve incluir retornos de
//  carro ou feeds de linha. Assim, por exemplo, se espera-se que um campo capture o primeiro nome de
//  um usuário e seu rótulo for "Nome", um espaço reservado adequado pode ser "por exemplo, Mustafa".
//
//   Nota:
//     * O atributo placeholder não é tão semanticamente útil quanto outras formas de explicar seu
//       formulário e pode causar problemas técnicos inesperados com seu conteúdo. Consulte Rótulos
//       para obter mais informações.
func (e *TagInput) Placeholder(placeholder string) (ref *TagInput) {
	e.selfElement.Set("placeholder", placeholder)
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
func (e *TagInput) ReadOnly(readonly bool) (ref *TagInput) {
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
func (e *TagInput) Required(required bool) (ref *TagInput) {
	e.selfElement.Set("required", required)
	return e
}

// Size
//
// English:
//
//  If the control is presented as a scrolling list box (e.g. when multiple is specified), this
//  attribute represents the number of rows in the list that should be visible at one time.
//  Browsers are not required to present a select element as a scrolled list box. The default value
//  is 0.
//
//   Note:
//     * According to the HTML5 specification, the default value for size should be 1; however, in
//       practice, this has been found to break some web sites, and no other browser currently does
//       that, so Mozilla has opted to continue to return 0 for the time being with Firefox.
//
// Português:
//
//  Se o controle for apresentado como uma caixa de listagem de rolagem (por exemplo, quando múltiplo
//  é especificado), esse atributo representa o número de linhas na lista que devem estar visíveis ao
//  mesmo tempo. Os navegadores não precisam apresentar um elemento de seleção como uma caixa de
//  listagem rolada. O valor padrão é 0.
//
//   Nota:
//     * De acordo com a especificação HTML5, o valor padrão para tamanho deve ser 1; no entanto, na
//       prática, descobriu-se que isso quebra alguns sites, e nenhum outro navegador atualmente faz
//       isso, então a Mozilla optou por continuar retornando 0 por enquanto com o Firefox.
func (e *TagInput) Size(size int) (ref *TagInput) {
	e.selfElement.Set("size", size)
	return e
}

// Src
//
// English:
//
//  Valid for the image input button only, the src is string specifying the URL of the image file to
//  display to represent the graphical submit button.
//
// Português:
//
//  Válido apenas para o botão de entrada de imagem, o src é uma string que especifica a URL do
//  arquivo de imagem a ser exibido para representar o botão de envio gráfico.
func (e *TagInput) Src(src string) (ref *TagInput) {
	e.selfElement.Set("src", src)
	return e
}

// Step
//
// English:
//
//  Valid for the numeric input types, including number, date/time input types, and range, the step
//  attribute is a number that specifies the granularity that the value must adhere to.
//
//   If not explicitly included:
//     * step defaults to 1 for number and range;
//     * For the date/time input types, step is expressed in seconds, with the default step being 60
//       seconds. The step scale factor is 1000 (which converts the seconds to milliseconds, as used
//       in other algorithms);
//     * The value must be a positive number—integer or float—or the special value any, which means
//       no stepping is implied, and any value is allowed (barring other constraints, such as min and
//       max).
//
// If any is not explicitly set, valid values for the number, date/time input types, and range input
// types are equal to the basis for stepping — the min value and increments of the step value, up to
// the max value, if specified.
//
// For example, if you have <input type="number" min="10" step="2">, then any even integer, 10 or
// greater, is valid. If omitted, <input type="number">, any integer is valid, but floats (like 4.2)
// are not valid, because step defaults to 1. For 4.2 to be valid, step would have had to be set to
// any, 0.1, 0.2, or any the min value would have had to be a number ending in .2, such as
// <input type="number" min="-5.2">
//
//   Note:
//     * When the data entered by the user doesn't adhere to the stepping configuration, the value is
//       considered invalid in constraint validation and will match the :invalid pseudoclass.
//
// Português:
//
//  Válido para os tipos de entrada numérica, incluindo número, tipos de entrada de data e hora e
//  intervalo, o atributo step é um número que especifica a granularidade à qual o valor deve aderir.
//
//   Se não estiver explicitamente incluído:
//     * step padroniza para 1 para número e intervalo.
//     * Para os tipos de entrada de data e hora, a etapa é expressa em segundos, com a etapa padrão
//       sendo 60 segundos. O fator de escala de passo é 1000 (que converte os segundos em
//       milissegundos, conforme usado em outros algoritmos).
//     * O valor deve ser um número positivo — inteiro ou flutuante — ou o valor especial any, o que
//       significa que nenhuma depuração está implícita e qualquer valor é permitido (exceto outras
//       restrições, como min e max).
//
// Se algum não for definido explicitamente, os valores válidos para o número, tipos de entrada de
// data e hora e tipos de entrada de intervalo são iguais à base para a depuração — o valor mínimo e
// os incrementos do valor da etapa, até o valor máximo, se especificado.
//
// Por exemplo, se você tiver <input type="number" min="10" step="2">, qualquer número inteiro par,
// 10 ou maior, é válido. Se omitido, <input type="number">, qualquer inteiro é válido, mas floats
// (como 4.2) não são válidos, porque step é padronizado como 1. Para 4.2 ser válido, step teria que
// ser definido como any, 0.1 , 0.2 ou qualquer valor mínimo teria que ser um número que terminasse
// em .2, como <input type="number" min="-5.2">
//
//   Nota:
//     * Quando os dados inseridos pelo usuário não estão de acordo com a configuração de stepping,
//       o valor é considerado inválido na validação da restrição e corresponderá à
//       :invalid pseudoclass.
func (e *TagInput) Step(step int) (ref *TagInput) {
	e.selfElement.Set("step", step)
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
func (e *TagInput) Type(inputType InputType) (ref *TagInput) {
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
func (e *TagInput) Value(value string) (ref *TagInput) {
	e.selfElement.Set("value", value)
	return e
}

// Width
//
// English:
//
//  Valid for the image input button only, the width is the width of the image file to display to
//  represent the graphical submit button. See the image input type.
//
// Português:
//
//  Válido apenas para o botão de entrada de imagem, a largura é a largura do arquivo de imagem a
//  ser exibido para representar o botão de envio gráfico. Consulte o tipo de entrada de imagem.
func (e *TagInput) Width(width int) (ref *TagInput) {
	e.selfElement.Set("width", width)
	return e
}
