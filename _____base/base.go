package base

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/mouse"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/easingTween"
	"log"
	"strconv"
	"strings"
	"sync"
	"syscall/js"
	"time"
)

type TagBaseGlobal struct {
	id          string
	selfElement js.Value
	cssClass    *css.Class

	x int
	y int

	// listener
	//
	// English:
	//
	//  The javascript function removeEventListener needs to receive the function passed in addEventListener
	//
	// Português:
	//
	//  A função javascript removeEventListener necessitam receber a função passada em addEventListener
	listener *sync.Map

	// drag

	// stage
	//
	// English:
	//
	//  Browser main document reference captured at startup.
	//
	// Português:
	//
	//  Referencia do documento principal do navegador capturado na inicialização.
	stage js.Value

	// isDragging
	//
	// English:
	//
	//  Indicates the process of dragging the element.
	//
	// Português:
	//
	//  Indica o processo de arrasto do elemento.
	isDragging bool

	// dragDifX
	//
	// English:
	//
	//  Used in calculating element drag.
	//
	// Português:
	//
	//  Usado no cálculo do arrasto de elemento.
	dragDifX int

	// dragDifX
	//
	// English:
	//
	//  Used in calculating element drag.
	//
	// Português:
	//
	//  Usado no cálculo do arrasto de elemento.
	dragDifY int

	// deltaMovieX
	//
	// English:
	//
	//  Additional value added in the SetX() function: (x = x + deltaMovieX) and subtracted in the
	//  GetX() function: (x = x - deltaMovieX).
	//
	// Português:
	//
	//  Valor adicional adicionado na função SetX(): (x = x + deltaMovieX)  e subtraído na função
	//  GetX(): (x = x - deltaMovieX).
	deltaMovieX int

	// deltaMovieY
	//
	// English:
	//
	//  Additional value added in the SetY() function: (y = y + deltaMovieY) and subtracted in the
	//  GetY() function: (y = y - deltaMovieY).
	//
	// Português:
	//
	//  Valor adicional adicionado na função SetY(): (y = y + deltaMovieY)  e subtraído na função
	//  GetY(): (y = y - deltaMovieY).
	deltaMovieY int

	// tween
	//
	// English:
	//
	//  Easing tween.
	//
	// Receives an identifier and a pointer of the tween object to be used in case of multiple
	// functions.
	//
	// Português:
	//
	//  Facilitador de interpolação.
	//
	// Recebe um identificador e um ponteiro do objeto tween para ser usado em caso de múltiplas
	// funções.
	tween map[string]interfaces.TweenInterface
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
func (e *TagBaseGlobal) AccessKey(key string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Autofocus(autofocus bool) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Class(class ...string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) ContentEditable(editable bool) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Data(data map[string]string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Dir(dir Dir) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Draggable(draggable Draggable) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Hidden() (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Id(id string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) InputMode(inputMode InputMode) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Is(is string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) ItemId(id string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) ItemDrop(itemprop string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) ItemRef(itemref string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) ItemType(itemType string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Lang(language Language) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Nonce(part ...string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Slot(slot string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Spellcheck(spell bool) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Style(style string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) TabIndex(index int) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Title(title string) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) Translate(translate Translate) (ref *TagBaseGlobal) {
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
func (e *TagBaseGlobal) CreateElement(tag Tag) (ref *TagBaseGlobal) {
	e.selfElement = js.Global().Get("document").Call("createElement", tag.String())
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

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
func (e *TagBaseGlobal) AppendById(appendId string) (ref *TagBaseGlobal) {

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
func (e *TagBaseGlobal) Append(append interface{}) (ref *TagBaseGlobal) {
	switch append.(type) {
	case *TagBaseGlobal:
		e.selfElement.Call("appendChild", append.(*TagBaseGlobal).selfElement)
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

// AppendToStage
//
// English:
//
//  Adds a node to the end of the list of children in the main document body. If the node already
//  exists somewhere in the document, it is removed from its current parent node before being added
//  to the main document.
//
// Português:
//
//  Adiciona um nó ao final da lista de filhos do corpo do documento principal. Se o nó já existir
//  em alguma parte do documento, ele é removido de seu nó pai atual antes de ser adicionado ao
//  documento principal.
//
// todo:https://developer.mozilla.org/en-US/docs/Web/API/Document/createDocumentFragment
// todo: appendMany()
func (e *TagBaseGlobal) AppendToStage() (ref *TagBaseGlobal) {
	e.stage.Call("appendChild", e.selfElement)
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
func (e *TagBaseGlobal) SetXY(x, y int) (ref *TagBaseGlobal) {
	e.x = x
	e.y = y

	px := strconv.FormatInt(int64(x), 10) + "px"
	py := strconv.FormatInt(int64(y), 10) + "px"

	e.selfElement.Get("style").Set("left", px)
	e.selfElement.Get("style").Set("top", py)

	return e
}

// SetDeltaX
//
// English:
//
//  Additional value added in the SetX() function: (x = x + deltaMovieX) and subtracted in the
//  GetX() function: (x = x - deltaMovieX).
//
// Português:
//
//  Valor adicional adicionado na função SetX(): (x = x + deltaMovieX)  e subtraído na função
//  GetX(): (x = x - deltaMovieX).
func (e *TagBaseGlobal) SetDeltaX(delta int) (ref *TagBaseGlobal) {
	e.deltaMovieX = delta
	return
}

// SetDeltaY
//
// English:
//
//  Additional value added in the SetY() function: (y = y + deltaMovieY) and subtracted in the
//  GetY() function: (y = y - deltaMovieY).
//
// Português:
//
//  Valor adicional adicionado na função SetY(): (y = y + deltaMovieY)  e subtraído na função
//  GetX(): (y = y - deltaMovieY).
func (e *TagBaseGlobal) SetDeltaY(delta int) (ref *TagBaseGlobal) {
	e.deltaMovieY = delta
	return
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
func (e *TagBaseGlobal) SetX(x int) (ref *TagBaseGlobal) {
	e.x = x
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
func (e *TagBaseGlobal) SetY(y int) (ref *TagBaseGlobal) {
	e.y = y

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
func (e *TagBaseGlobal) GetXY() (x, y int) {
	x = e.GetX()
	y = e.GetY()

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
func (e *TagBaseGlobal) GetX() (x int) {
	//rect.top, rect.right, rect.bottom, rect.left
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	x = coordinate.Get("left").Int()
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
func (e *TagBaseGlobal) GetY() (y int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	y = coordinate.Get("top").Int()
	return
}

// GetTop
//
// English:
//
//  Same as GetX() function, returns the x position of the element.
//
// Português:
//
//  O mesmo que a função GetX(), retorna a posição x do elemento.
func (e *TagBaseGlobal) GetTop() (top int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	top = coordinate.Get("top").Int()
	return
}

// GetRight
//
// English:
//
//  It is the same as x + width.
//
// Português:
//
//  É o mesmo que x + width.
func (e *TagBaseGlobal) GetRight() (right int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	right = coordinate.Get("right").Int()
	return
}

// GetBotton
//
// English:
//
//  It is the same as y + height.
//
// Português:
//
//  É o mesmo que y + Heught.
func (e *TagBaseGlobal) GetBottom() (bottom int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	bottom = coordinate.Get("bottom").Int()
	return
}

// GetLeft
//
// English:
//
//  Same as GetY() function, returns the y position of the element.
//
// Português:
//
//  O mesmo que a função GetY(), retorna a posição y do elemento.
func (e *TagBaseGlobal) GetLeft() (left int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	left = coordinate.Get("left").Int()
	return
}

// AddListener
//
// English:
//
//  Associates a function with an event.
//
//   Example:
//
//     stage.AddListener(browserMouse.KEventMouseOver, onMouseEvent)
//     timer := time.NewTimer(10 * time.Second)
//     go func() {
//       select {
//         case <-timer.C:
//         stage.RemoveListener(mouse.KEventMouseOver)
//       }
//     }()
//
//     func onMouseEvent(event browserMouse.MouseEvent) {
//       isNull, target := event.GetRelatedTarget()
//       if isNull == false {
//         log.Print("id: ", target.Get("id"))
//         log.Print("tagName: ", target.Get("tagName"))
//       }
//       log.Print(event.GetScreenX())
//       log.Print(event.GetScreenY())
//     }
//
// Português:
//
//  Associa uma função a um evento.
//
//   Exemplo:
//
//     stage.AddListener(browserMouse.KEventMouseOver, onMouseEvent)
//     timer := time.NewTimer(10 * time.Second)
//     go func() {
//       select {
//         case <-timer.C:
//         stage.RemoveListener(mouse.KEventMouseOver)
//       }
//     }()
//
//     func onMouseEvent(event browserMouse.MouseEvent) {
//       isNull, target := event.GetRelatedTarget()
//       if isNull == false {
//         log.Print("id: ", target.Get("id"))
//         log.Print("tagName: ", target.Get("tagName"))
//       }
//       log.Print(event.GetScreenX())
//       log.Print(event.GetScreenY())
//     }
func (e *TagBaseGlobal) AddListener(eventType interface{}, manager mouse.SimpleManager) (ref *TagBaseGlobal) {

	mouseMoveEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var mouseEvent = mouse.MouseEvent{}

		if len(args) > 0 {
			mouseEvent.Object = args[0]
		}

		if manager != nil {
			manager(mouseEvent)
		}

		return nil
	})

	switch converted := eventType.(type) {
	case event.Event:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventAnimation.EventAnimation:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventClipBoard.EventClipBoard:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventDrag.EventDrag:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventFocus.EventFocus:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventHashChange.EventHashChange:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventInput.EventInput:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventKeyboard.EventKeyboard:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	case mouse.Event:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventPageTransition.EventPageTransition:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventUi.EventUi:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	case eventWheel.EventWheel:
		e.listener.Store(converted.String(), mouseMoveEvt)
		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)

	default:
		log.Fatalf("event must be a event type")
	}

	return e
}

// RemoveListener
//
// English:
//
//  Remove the function associated with the event
//
//   Example:
//
//     stage.AddListener(browserMouse.KEventMouseOver, onMouseEvent)
//     timer := time.NewTimer(10 * time.Second)
//     go func() {
//       select {
//         case <-timer.C:
//         stage.RemoveListener(mouse.KEventMouseOver)
//       }
//     }()
//
//     func onMouseEvent(event browserMouse.MouseEvent) {
//       isNull, target := event.GetRelatedTarget()
//       if isNull == false {
//         log.Print("id: ", target.Get("id"))
//         log.Print("tagName: ", target.Get("tagName"))
//       }
//       log.Print(event.GetScreenX())
//       log.Print(event.GetScreenY())
//     }
//
// Português:
//
//  Remove a função associada com o evento.
//
//   Exemplo:
//
//     stage.AddListener(browserMouse.KEventMouseOver, onMouseEvent)
//     timer := time.NewTimer(10 * time.Second)
//     go func() {
//       select {
//         case <-timer.C:
//         stage.RemoveListener(mouse.KEventMouseOver)
//       }
//     }()
//
//     func onMouseEvent(event browserMouse.MouseEvent) {
//       isNull, target := event.GetRelatedTarget()
//       if isNull == false {
//         log.Print("id: ", target.Get("id"))
//         log.Print("tagName: ", target.Get("tagName"))
//       }
//       log.Print(event.GetScreenX())
//       log.Print(event.GetScreenY())
//     }
func (e *TagBaseGlobal) RemoveListener(eventType interface{}) (ref *TagBaseGlobal) {
	switch converted := eventType.(type) {
	case event.Event:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	case eventAnimation.EventAnimation:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	case eventClipBoard.EventClipBoard:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	case eventDrag.EventDrag:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	case eventFocus.EventFocus:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	case eventHashChange.EventHashChange:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	case eventInput.EventInput:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	case eventKeyboard.EventKeyboard:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	case mouse.Event:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	case eventPageTransition.EventPageTransition:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	case eventUi.EventUi:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	case eventWheel.EventWheel:
		f, _ := e.listener.Load(converted.String())
		e.selfElement.Call("removeEventListener", converted.String(), f)

	default:
		log.Fatalf("event must be a event type")
	}

	return e
}

// Mouse
//
// English:
//
//  Defines the shape of the mouse pointer.
//
//   Input:
//     value: mouse pointer shape.
//       Example: SetMouse(mouse.KCursorCell) // Use mouse.K... and let autocomplete do the
//                rest
//
// Português:
//
//  Define o formato do ponteiro do mouse.
//
//   Entrada:
//     value: formato do ponteiro do mouse.
//       Exemplo: SetMouse(mouse.KCursorCell) // Use mouse.K... e deixe o autocompletar fazer
//                o resto
func (e *TagBaseGlobal) Mouse(value mouse.CursorType) (ref *TagBaseGlobal) {
	e.selfElement.Get("style").Set("cursor", value.String())
	return e
}

// Init
//
// English:
//
//  Initializes the object correctly.
//
// Português:
//
//  Inicializa o objeto corretamente.
func (e *TagBaseGlobal) Init() (ref *TagBaseGlobal) {
	e.listener = new(sync.Map)

	e.CreateElement(KTagDiv)
	e.prepareStageReference()

	return e
}

// prepareStageReference
//
// English:
//
//  Prepares the stage reference at initialization.
//
// Português:
//
//  Prepara à referencia do stage na inicialização.
func (e *TagBaseGlobal) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

// DragStart
//
// English:
//
//  Mouse drag function.
//
//   Example:
//
//     factoryBrowser.NewTagDiv("div_0").
//       Class("animate").
//       DragStart().
//       AppendById("stage")
//
// Português:
//
//  Função de arrastar com o mouse.
//
//   Exemplo:
//
//     factoryBrowser.NewTagDiv("div_0").
//       Class("animate").
//       DragStart().
//       AppendById("stage")
func (e *TagBaseGlobal) DragStart() (ref *TagBaseGlobal) {
	e.dragNormalStart()
	return e
}

func (e *TagBaseGlobal) DragStop() (ref *TagBaseGlobal) {
	e.dragNormalStop()
	return e
}

func (e *TagBaseGlobal) dragNormalStart() {
	e.AddListener(mouse.KEventMouseDown, e.onStartDragNormal)
	e.stage.Call("addEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
	e.stage.Call("addEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
}

func (e *TagBaseGlobal) dragNormalStop() {
	e.RemoveListener(mouse.KEventMouseDown)
	e.stage.Call("removeEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
	e.stage.Call("removeEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
	e.isDragging = false
}

func (e *TagBaseGlobal) onStopDragNormal(_ js.Value, _ []js.Value) any {
	e.isDragging = false
	return nil
}

func (e *TagBaseGlobal) onStartDragNormal(event mouse.MouseEvent) {
	var screenX = int(event.GetScreenX())
	var screenY = int(event.GetScreenY())

	e.dragDifX = screenX - e.x
	e.dragDifY = screenY - e.y

	e.isDragging = true
}

func (e *TagBaseGlobal) onMouseDraggingNormal(_ js.Value, args []js.Value) interface{} {
	if e.isDragging == false {
		return nil
	}

	var mouseEvent = mouse.MouseEvent{}
	if len(args) > 0 {
		mouseEvent.Object = args[0]

		var x = int(mouseEvent.GetScreenX()) - e.dragDifX
		var y = int(mouseEvent.GetScreenY()) - e.dragDifY

		e.SetXY(x, y)
	}

	return nil
}

// EasingTweenFunc
//
// English:
//
//  Defines the tween math function to control the loop of interactions
//
//   Input:
//     id: tween identifier for multiple tween functions.
//     value: tween math function.
//       currentTime:   current time, int64(time.Duration);
//       duration:      total time, int64(time.Duration);
//       startValue:    initial value;
//       endValue:      final value;
//       changeInValue: startValue - endValue
//
//   Output:
//     object: reference to the current Tween object.
//
//   Note:
//     * To create a new function, base it on the linear function, where:
//         return changeInValue * currentTime / duration + startValue
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Define a função matemática tween para controle do ciclo de interações
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//     value: função matemática tween.
//       currentTime:   tempo atual, int64(time.Duration);
//       duration:      tempo total, int64(time.Duration);
//       startValue:    valor inicial;
//       endValue:      valor final;
//       changeInValue: startValue - endValue
//
//   Saída:
//     object: referência para o objeto Tween corrente.
//
//   Nota:
//     * Para criar uma nova função, tenha como base a função linear, onde:
//         return changeInValue * currentTime / duration + startValue
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) EasingTweenFunc(id string, value func(currentTime, duration, currentPercentage, startValue, endValue, changeInValue float64) (percent float64)) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = &easingTween.Tween{}
	}

	e.tween[id].SetTweenFunc(value)
	return e
}

// EasingTweenValues
//
// English:
//
//  Defines the initial and final values of the interactions cycle.
//
//   Input:
//     id: tween identifier for multiple tween functions.
//     start: initial value for the beginning of the cycle of interactions;
//     end:   final value for the end of the iteration cycle.
//
//   Output:
//     object: reference to the current Tween object.
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Defines os valores inicial e final do ciclo de interações.
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//     start: valor inicial para o início do ciclo de interações;
//     end:   valor final para o fim do ciclo de interações.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) EasingTweenValues(id string, start, end float64) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = &easingTween.Tween{}
	}

	e.tween[id].SetValues(start, end)
	return e
}

// EasingTweenDuration
//
// English:
//
//  Defines the total cycle time of interactions.
//
//   Input:
//     id: tween identifier for multiple tween functions.
//     value: time.Duration contendo o tempo do ciclo de interações.
//
//   Output:
//     object: reference to the current Tween object.
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Define o tempo total do ciclo de interações.
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//     value: time.Duration contendo o tempo do ciclo de interações.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) EasingTweenDuration(id string, value time.Duration) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].SetDuration(value)
	return e
}

// EasingTweenDoNotReverseMotion
//
// English:
//
//  Defines the option of reversing values at the end of each cycle.
//
//   Input:
//     id: tween identifier for multiple tween functions.
//     value: true to not revert the values at the end of each cycle.
//
//   Output:
//     object: reference to the current Tween object.
//
//   Notas:
//     * In case of loop, the order of event functions are: SetOnStartFunc(), SetOnCycleStartFunc(),
//       SetOnCycleEndFunc(), SetOnInvertFunc(), SetOnCycleStartFunc(), SetOnCycleEndFunc(),
//       SetOnInvertFunc() ...
//     * SetOnEndFunc() will only be called at the end of all interactions;
//     * This function prevents inversion of values, but the SetOnInvertFunc() event function
//       continues to be called.
//
// Português:
//
//  Define a opção de reversão de valores ao final de cada ciclo.
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//     value: true para não reverter os valores ao final de cada ciclo.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
//
//   Notas:
//     * Em caso de laço, a ordem das funções de eventos são: SetOnStartFunc(), SetOnCycleStartFunc(),
//       SetOnCycleEndFunc(), SetOnInvertFunc(), SetOnCycleStartFunc(), SetOnCycleEndFunc(),
//       SetOnInvertFunc() ...
//     * SetOnEndFunc() só será chamada ao final de todas as interações.
//     * Esta função impede a inversão de valores, mas, a função de evento SetOnInvertFunc() continua
//       sendo chamada.
func (e *TagBaseGlobal) EasingTweenDoNotReverseMotion(id string, value bool) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].SetDoNotReverseMotion(value)
	return e
}

// EasingTweenLoops
//
// English:
//
//  Defines the number of loops before the end of the function.
//
//   Input:
//     id: tween identifier for multiple tween functions.
//
//   Notes:
//     * At each new iteration of the loop, a movement inversion will occur, unless the
//       SetDoNotReverseMotion(true) function is used;
//     * For infinite loops, set the value to -1;
//     * In case of loop, the order of event functions are: SetOnStartFunc(), SetOnCycleStartFunc(),
//       SetOnCycleEndFunc(), SetOnInvertFunc(), SetOnCycleStartFunc(), SetOnCycleEndFunc(),
//       SetOnInvertFunc() ...
//     * SetOnEndFunc() will only be called at the end of all interactions.
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Define a quantidade de laços antes do fim da função.
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//
//   Notas:
//     * A cada nova interação do laço ocorrerá uma inversão de movimento, a não ser que seja usada a
//       função SetDoNotReverseMotion(true);
//     * Para laços infinitos, defina o valor como sendo -1;
//     * Em caso de laço, a ordem das funções de eventos são: SetOnStartFunc(), SetOnCycleStartFunc(),
//       SetOnCycleEndFunc(), SetOnInvertFunc(), SetOnCycleStartFunc(), SetOnCycleEndFunc(),
//       SetOnInvertFunc() ...
//     * SetOnEndFunc() só será chamada ao final de todas as interações.
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) EasingTweenLoops(id string, value int) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].SetLoops(value)
	return e
}

// EasingTweenOnStartFunc
//
// English:
//
//  Add the function to be called when the animation starts.
//
//   Input:
//     id: tween identifier for multiple tween functions;
//     function: func(value float64, arguments ...interface{})
//       value: initial value defined in startValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada quando a animação inicia.
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//     function: func(value float64, arguments ...interface{})
//       value: valor inicial definido em startValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (e *TagBaseGlobal) EasingTweenOnStartFunc(id string, function func(value float64, arguments interface{})) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].SetOnStartFunc(function)
	return e
}

// EasingTweenOnEndFunc
//
// English:
//
//  Add the function to be called when the animation ends.
//
//   Input:
//     id: tween identifier for multiple tween functions.
//     function: func(value float64, arguments ...interface{})
//       value: final value defined in endValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada quando a animação inicia.
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//     function: func(value float64, arguments ...interface{})
//       value: valor final definido em endValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (e *TagBaseGlobal) EasingTweenOnEndFunc(id string, function func(value float64, arguments interface{})) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].SetOnEndFunc(function)
	return e
}

// EasingTweenOnCycleStartFunc
//
// English:
//
//  Adds the function to be called at the beginning of the interpolation cycle
//
//   Input:
//     id: tween identifier for multiple tween functions.
//     function: func(value float64, arguments ...interface{})
//       value: initial value defined in startValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada no início do ciclo de interpolação
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//     function: func(value float64, arguments ...interface{})
//       value: valor inicial definido em startValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (e *TagBaseGlobal) EasingTweenOnCycleStartFunc(id string, function func(value float64, arguments interface{})) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].SetOnCycleStartFunc(function)
	return e
}

// EasingTweenOnCycleEndFunc
//
// English:
//
//  Adds the function to be called at the ending of the interpolation cycle
//
//   Input:
//     id: tween identifier for multiple tween functions.
//     function: func(value float64, arguments ...interface{})
//       value: final value defined in endValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada no fim do ciclo de interpolação
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//     function: func(value float64, arguments ...interface{})
//       value: valor final definido em endValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (e *TagBaseGlobal) EasingTweenOnCycleEndFunc(id string, function func(value float64, arguments interface{})) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].SetOnCycleEndFunc(function)
	return e
}

// EasingTweenOnStepFunc
//
// English:
//
//  Adds the function to be called for each iteration.
//
//   Input:
//     id: tween identifier for multiple tween functions.
//     function: func(value float64, arguments ...interface{})
//       value: current value
//       percentToComplete: value between 0.0 and 1.0 indicating the percentage of the process
//       arguments: list of values passed to event functions, defined in SetArguments()
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Adiciona a função a ser chamada a cada interação
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//     function: func(value float64, arguments ...interface{})
//       value: valor corrente
//       percentToComplete: valor entre 0.0 e 1.0 indicando o percentual do processo
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) EasingTweenOnStepFunc(id string, function func(value, percentToComplete float64, arguments interface{})) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].SetOnStepFunc(function)
	return e
}

// EasingTweenOnInvertFunc
//
// English:
//
//  Adds the function to be called on inversion of the interpolation cycle
//
//   Input:
//     id: tween identifier for multiple tween functions.
//     function: func(value float64, arguments ...interface{})
//       value: current value
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada a cada interação
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//     function: func(value, percentToComplete float64, arguments ...interface{})
//       value: valor corrente
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (e *TagBaseGlobal) EasingTweenOnInvertFunc(id string, function func(value float64, arguments interface{})) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].SetOnInvertFunc(function)
	return e
}

// EasingTweenArgumentsFunc
//
// English:
//
//  Determines the arguments passed to event functions.
//
//   Input:
//     id: tween identifier for multiple tween functions;
//     arguments: list of interfaces{} passed to all event functions when they are invoked.
//
//   Output:
//     object: reference to the current Tween object.
//
//   Note:
//     * If you need complex functions, remember to use pointers to data in the arguments.
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Determina os argumentos passados para as funções de eventos.
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//     arguments: lista de interfaces{} passadas para todas as funções de eventos quando elas são invocadas.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
//
//   Nota:
//     * Caso necessite de funções complexas, lembre-se de usar ponteiros para dados nos argumentos.
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) EasingTweenArgumentsFunc(id string, arguments interface{}) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].SetArgumentsFunc(arguments)
	return e
}

// EasingTweenStart
//
// English:
//
//  Starts the interaction according to the chosen tween function.
//
//   Input:
//     id: tween identifier for multiple tween functions.
//
//   Output:
//     object: reference to the current Tween object.
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Inicia a interação conforme a função tween escolhida.
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//
//   Saída:
//     object: referência para o objeto Tween corrente.
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) EasingTweenStart(id string) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].Start()
	return e
}

// EasingTweenEnd
//
// English:
//
//  Terminates all interactions of the chosen Tween function, without invoking the onCycleEnd and
//  onEnd functions.
//
//   Entrada:
//     id: tween identifier for multiple tween functions.
//
//   Saída:
//     object: reference to the current Tween object.
//
// Português:
//
// Termina todas as interações da função Tween escolhida, sem invocar as funções onCycleEnd e onEnd.
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (e *TagBaseGlobal) EasingTweenEnd(id string) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].End()
	return e
}

// EasingTweenStop
//
// English:
//
//  Ends all interactions of the chosen Tween function, interacting with the onCycleEnd and onEnd
//  functions, respectively, in that order, if they have been defined.
//
//  Input:
//     id: tween identifier for multiple tween functions.
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Termina todas as interações da função Tween escolhida, interagindo com as funções onCycleEnd e
//  onEnd, respectivamente nessa ordem, se elas tiverem sido definidas.
//
//   Entrada:
//     id: identificador de tween para múltiplas funções tween;
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (e *TagBaseGlobal) EasingTweenStop(id string) (ref *TagBaseGlobal) {
	if e.tween[id] == nil {
		e.tween[id] = new(easingTween.Tween)
	}

	e.tween[id].Stop()
	return e
}

// NewEasingTweenRandom
//
// English:
//
//  Ease tween random.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenRandom("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenRandom("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.SelectRandom()).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.SelectRandom()).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação random.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenRandom("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenRandom("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.SelectRandom()).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.SelectRandom()).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenRandom(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.SelectRandom()).
		Start()

	return e
}

// NewEasingTweenLinear
//
// English:
//
//  Ease tween linear.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenLinear("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenLinear("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KLinear).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KLinear).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação linear.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenLinear("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenLinear("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KLinear).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KLinear).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenLinear(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KLinear).
		Start()

	return e
}

// NewEasingTweenOutSine
//
// English:
//
//  Ease tween out sine.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutSine("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutSine("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutSine).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutSine).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação out sine.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutSine("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutSine("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutSine).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutSine).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenOutSine(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseOutSine).
		Start()

	return e
}

// NewEasingTweenOutQuintic
//
// English:
//
//  Ease tween out quintic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutQuintic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutQuintic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutQuintic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutQuintic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação out quintic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutQuintic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutQuintic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutQuintic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutQuintic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenOutQuintic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseOutQuintic).
		Start()

	return e
}

// NewEasingTweenOutQuartic
//
// English:
//
//  Ease tween out quartic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutQuartic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutQuartic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutQuartic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutQuartic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação out quartic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutQuartic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutQuartic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutQuartic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutQuartic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenOutQuartic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseOutQuartic).
		Start()

	return e
}

// NewEasingTweenOutQuadratic
//
// English:
//
//  Ease tween out quadratic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutQuadratic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutQuadratic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutQuadratic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutQuadratic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação out quadratic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutQuadratic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutQuadratic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutQuadratic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutQuadratic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenOutQuadratic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseOutQuadratic).
		Start()

	return e
}

// NewEasingTweenOutExponential
//
// English:
//
//  Ease tween out exponential.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutExponential("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutExponential("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutExponential).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutExponential).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação out exponential.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutExponential("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutExponential("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutExponential).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutExponential).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenOutExponential(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseOutExponential).
		Start()

	return e
}

// NewEasingTweenOutElastic
//
// English:
//
//  Ease tween out elastic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutElastic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutElastic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutElastic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutElastic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação out elastic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutElastic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutElastic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutElastic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutElastic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenOutElastic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseOutElastic).
		Start()

	return e
}

// NewEasingTweenOutCubic
//
// English:
//
//  Ease tween out cubic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutCubic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutCubic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutCubic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutCubic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação out cubic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutCubic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutCubic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutCubic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutCubic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenOutCubic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseOutCubic).
		Start()

	return e
}

// NewEasingTweenOutCircular
//
// English:
//
//  Ease tween out circular.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutCircular("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutCircular("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutCircular).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutCircular).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação out circular.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutCircular("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutCircular("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutCircular).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutCircular).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenOutCircular(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseOutCircular).
		Start()

	return e
}

// NewEasingTweenOutBounce
//
// English:
//
//  Ease tween out bounce.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutBounce("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutBounce("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutBounce).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutBounce).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação out bounce.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutBounce("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutBounce("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutBounce).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutBounce).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenOutBounce(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseOutBounce).
		Start()

	return e
}

// NewEasingTweenOutBack
//
// English:
//
//  Ease tween out back.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação out back.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenOutBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenOutBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseOutBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseOutBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenOutBack(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseOutBack).
		Start()

	return e
}

// NewEasingTweenInSine
//
// English:
//
//  Ease tween in sine.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInSine("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInSine("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInSine).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInSine).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in sine.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInSine("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInSine("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInSine).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInSine).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInSine(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInSine).
		Start()

	return e
}

// NewEasingTweenInQuintic
//
// English:
//
//  Ease tween in quintic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInQuintic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInQuintic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInQuintic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInQuintic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in quintic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInQuintic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInQuintic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInQuintic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInQuintic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInQuintic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInQuintic).
		Start()

	return e
}

// NewEasingTweenInQuartic
//
// English:
//
//  Ease tween in quartic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInQuartic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInQuartic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInQuartic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInQuartic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in quartic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInQuartic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInQuartic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInQuartic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInQuartic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInQuartic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInQuartic).
		Start()

	return e
}

// NewEasingTweenInQuadratic
//
// English:
//
//  Ease tween in quadratic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInQuadratic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInQuadratic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInQuadratic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInQuadratic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in quadratic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInQuadratic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInQuadratic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInQuadratic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInQuadratic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInQuadratic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInQuadratic).
		Start()

	return e
}

// NewEasingTweenInOutSine
//
// English:
//
//  Ease tween in out sine.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutSine("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutSine("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutSine).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutSine).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in out sine.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutSine("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutSine("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutSine).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutSine).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInOutSine(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInOutSine).
		Start()

	return e
}

// NewEasingTweenInOutQuintic
//
// English:
//
//  Ease tween in out quintic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutQuintic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutQuintic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutQuintic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutQuintic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in out quintic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutQuintic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutQuintic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutQuintic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutQuintic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInOutQuintic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInOutQuintic).
		Start()

	return e
}

// NewEasingTweenInOutQuartic
//
// English:
//
//  Ease tween in out quartic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutQuartic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutQuartic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutQuartic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutQuartic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in out quartic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutQuartic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutQuartic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutQuartic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutQuartic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInOutQuartic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInOutQuartic).
		Start()

	return e
}

// NewEasingTweenInOutQuadratic
//
// English:
//
//  Ease tween in out quadratic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutQuadratic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutQuadratic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutQuadratic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutQuadratic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in out quadratic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutQuadratic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutQuadratic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutQuadratic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutQuadratic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInOutQuadratic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInOutQuadratic).
		Start()

	return e
}

// NewEasingTweenInOutExponential
//
// English:
//
//  Ease tween in out exponential.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutExponential("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutExponential("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutExponential).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutExponential).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in out exponential.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutExponential("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutExponential("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutExponential).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutExponential).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInOutExponential(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInOutExponential).
		Start()

	return e
}

// NewEasingTweenInOutElastic
//
// English:
//
//  Ease tween in out elastic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutElastic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutElastic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutElastic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutElastic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in out elastic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutElastic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutElastic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutElastic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutElastic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInOutElastic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInOutElastic).
		Start()

	return e
}

// NewEasingTweenInOutCubic
//
// English:
//
//  Ease tween in out cubic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutCubic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutCubic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutCubic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutCubic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in out cubic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutCubic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutCubic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutCubic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutCubic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInOutCubic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInOutCubic).
		Start()

	return e
}

// NewEasingTweenInOutCircular
//
// English:
//
//  Ease tween in out circular.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutCircular("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutCircular("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutCircular).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutCircular).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in out circular.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutCircular("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutCircular("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutCircular).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutCircular).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInOutCircular(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInOutCircular).
		Start()

	return e
}

// NewEasingTweenInOutBounce
//
// English:
//
//  Ease tween in out bounce.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutBounce("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutBounce("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutBounce).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutBounce).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in out bounce.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutBounce("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutBounce("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutBounce).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutBounce).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInOutBounce(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInOutBounce).
		Start()

	return e
}

// NewEasingTweenInOutBack
//
// English:
//
//  Ease tween in out back.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in out back.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInOutBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInOutBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInOutBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInOutBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInOutBack(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInOutBack).
		Start()

	return e
}

// NewEasingTweenInExponential
//
// English:
//
//  Ease tween in exponential.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInExponential("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInExponential("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInExponential).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInExponential).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in exponential.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInExponential("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInExponential("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInExponential).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInExponential).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInExponential(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInExponential).
		Start()

	return e
}

// NewEasingTweenInElastic
//
// English:
//
//  Ease tween in elastic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInElastic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInElastic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInElastic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInElastic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in elastic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInElastic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInElastic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInElastic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInElastic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInElastic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInElastic).
		Start()

	return e
}

// NewEasingTweenInCubic
//
// English:
//
//  Ease tween in cubic.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInCubic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInCubic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInCubic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInCubic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in cubic.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInCubic("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInCubic("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInCubic).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInCubic).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInCubic(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInCubic).
		Start()

	return e
}

// NewEasingTweenInCircular
//
// English:
//
//  Ease tween in circular.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInCircular("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInCircular("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInCircular).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInCircular).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in circular.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInCircular("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInCircular("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInCircular).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInCircular).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInCircular(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInCircular).
		Start()

	return e
}

// NewEasingTweenInBack
//
// English:
//
//  Ease tween in back.
//
//   Input:
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//        Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Example 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Example 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // Caution: the pointer has been populated in here. Hence a code break in two parts.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
// Português:
//
//  Facilitador de interpolação in back.
//
//   Entrada:
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//        Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
//   Exemplo 1:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         NewEasingTweenInBack("x", 3*time.Second, 50, 300, onUpdateX, -1).
//         NewEasingTweenInBack("y", 3*time.Second, 50, 300, onUpdateY, -1).
//         AppendToStage()
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
//
//   Exemplo 2:
//
//     //go:build js
//     // +build js
//     package main
//
//     import (
//       "github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/html"
//       "github.com/helmutkemper/iotmaker.webassembly/browser/stage"
//       "time"
//     )
//
//     //.animate {
//     //  width: 29px;
//     //  height: 50px;
//     //  position: absolute;
//     //  background-image: url("./small.png");
//     //}
//
//     func main() {
//       done := make(chan struct{}, 0)
//       var stage = stage.Stage{}
//       stage.Init()
//       var div *html.TagDiv
//       div = factoryBrowser.NewTagDiv("div_0").
//         Class("animate").
//         AppendToStage()
//
//       // cuidado: o ponteiro foi preenchido aqui. Por isto, uma quebra do código em duas partes.
//
//       // easing tween for x
//       div.EasingTweenDuration("x", durationX).
//         EasingTweenValues("x", xStart, xEnd).
//         EasingTweenFunc("x", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("x", onUpdateX).
//         EasingTweenLoops("x", -1).
//         EasingTweenArgumentsFunc("x", []interface{}{div}).
//         EasingTweenStart("x").
//
//         // easing tween for y
//         EasingTweenDuration("y", durationY).
//         EasingTweenValues("y", yStart, yEnd).
//         EasingTweenFunc("y", easingTween.KEaseInBack).
//         EasingTweenOnStepFunc("y", onUpdateY).
//         EasingTweenLoops("y", -1).
//         EasingTweenStart("y").
//         EasingTweenArgumentsFunc("y", []interface{}{div})
//
//       <-done
//     }
//
//     func onUpdateX(x, _ float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetX(int(x))
//     }
//
//     func onUpdateY(y, p float64, args interface{}) {
//       this := args.([]interface{})[0].(*html.TagDiv)
//       this.SetY(int(y))
//     }
func (e *TagBaseGlobal) NewEasingTweenInBack(
	id string,
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) (ref *TagBaseGlobal) {

	arguments = append([]interface{}{e}, arguments...)

	e.tween[id] = new(easingTween.Tween)
	e.tween[id].SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInBack).
		Start()

	return e
}
