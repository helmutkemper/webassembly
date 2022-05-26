package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventAnimation"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventClipBoard"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventDrag"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventFocus"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventHashChange"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventInput"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventKeyboard"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventPageTransition"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventUi"
	"github.com/helmutkemper/iotmaker.webassembly/browser/eventWheel"
	"github.com/helmutkemper/iotmaker.webassembly/browser/mouse"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"image/color"
	"log"
	"strconv"
	"strings"
	"sync"
	"syscall/js"
	"time"
)

// TagSvgAnimate
//
// English:
//
//  The SVG <animate> element provides a way to animate an attribute of an element over time.
//
// Português:
//
//  O elemento SVG <animate> fornece uma maneira de animar um atributo de um elemento ao longo do tempo.
type TagSvgAnimate struct {

	// id
	//
	// English:
	//
	//  Unique id, standard html id property.
	//
	// Português:
	//
	//  Id único, propriedade padrão id do html.
	id string

	// selfElement
	//
	// English:
	//
	//  Reference to self element as js.Value.
	//
	// Português:
	//
	//  Referencia ao próprio elemento na forma de js.Value.
	selfElement js.Value

	cssClass *css.Class

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

	points    *[]algorithm.Point
	pointsLen int

	rotateDelta float64
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
func (e *TagSvgAnimate) AccessKey(key string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Autofocus(autofocus bool) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Class(class ...string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) ContentEditable(editable bool) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Data(data map[string]string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Dir(dir Dir) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Draggable(draggable Draggable) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Hidden() (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Id(id string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) InputMode(inputMode InputMode) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Is(is string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) ItemId(id string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) ItemDrop(itemprop string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) ItemRef(itemref string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) ItemType(itemType string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Lang(language Language) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Nonce(part ...string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Slot(slot string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Spellcheck(spell bool) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Style(style string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) TabIndex(index int) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Title(title string) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Translate(translate Translate) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) CreateElement(tag Tag) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) AppendById(appendId string) (ref *TagSvgAnimate) {

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
func (e *TagSvgAnimate) Append(append interface{}) (ref *TagSvgAnimate) {
	switch append.(type) {
	case *TagSvgAnimate:
		e.selfElement.Call("appendChild", append.(*TagSvgAnimate).selfElement)
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
func (e *TagSvgAnimate) AppendToStage() (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) SetXY(x, y int) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) SetDeltaX(delta int) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) SetDeltaY(delta int) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) SetX(x int) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) SetY(y int) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) GetXY() (x, y int) {
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
func (e *TagSvgAnimate) GetX() (x int) {
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
func (e *TagSvgAnimate) GetY() (y int) {
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
func (e *TagSvgAnimate) GetTop() (top int) {
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
func (e *TagSvgAnimate) GetRight() (right int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	right = coordinate.Get("right").Int()
	return
}

// GetBottom
//
// English:
//
//  It is the same as y + height.
//
// Português:
//
//  É o mesmo que y + Height.
func (e *TagSvgAnimate) GetBottom() (bottom int) {
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
func (e *TagSvgAnimate) GetLeft() (left int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	left = coordinate.Get("left").Int()
	return
}

// Rotate
//
// English:
//
//  Defines a transformation that rotates an element around a fixed point on the 2D plane, without deforming it.
//
//   Input:
//     angle: representing the angle of the rotation. The direction of rotation depends on the writing direction.
//     In a left-to-right context, a positive angle denotes a clockwise rotation, a negative angle a counter-clockwise
//     one.
//     In a right-to-left context, a positive angle denotes a counter-clockwise rotation, a negative angle a clockwise
//     one.
//
// Português:
//
//  Define uma transformação que gira um elemento em torno de um ponto fixo no plano 2D, sem deformá-lo.
//
//   Entrada:
//     angle: representando o ângulo de rotação. O sentido de rotação depende do sentido de escrita.
//     Em um contexto da esquerda para a direita, um ângulo positivo denota uma rotação no sentido horário, um ângulo
//     negativo no sentido anti-horário.
//     Em um contexto da direita para a esquerda, um ângulo positivo denota uma rotação no sentido anti-horário, um
//     ângulo negativo denota uma rotação no sentido horário.
func (e *TagSvgAnimate) Rotate(angle float64) (ref *TagSvgAnimate) {
	angleAsString := strconv.FormatFloat(angle+e.rotateDelta, 'E', -1, 64)
	e.selfElement.Get("style").Set("transform", "rotate("+angleAsString+"rad)")
	return e
}

// RotateDelta
//
// English:
//
//  Used in conjunction with the Rotate() function, sets the rotation adjustment angle, ie Rotate() = angle + delta.
//
//   Input:
//     angle: delta, object rotation adjustment angle.
//
// Português:
//
//  Usada em conjunto com a função Rotate(), define o ângulo de ajuste da rotação, ou seja, Rotate() = angle + delta.
//
//   Entrada:
//     angle: delta, ângulo de ajuste da rotação do objeto.
func (e *TagSvgAnimate) RotateDelta(delta float64) (ref *TagSvgAnimate) {
	e.rotateDelta = delta
	return e
}

// GetRotateDelta
//
// English:
//
//  Returns the rotation adjustment angle, i.e. Rotate() = angle + delta.
//
//   Output:
//     angle: delta, object rotation adjustment angle.
//
// Português:
//
//  Retorna o ângulo de ajuste da rotação, ou seja, Rotate() = angle + delta.
//
//   Saída:
//     angle: delta, ângulo de ajuste da rotação do objeto.
func (e *TagSvgAnimate) GetRotateDelta() (delta float64) {
	return e.rotateDelta
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
func (e *TagSvgAnimate) AddListener(eventType interface{}, manager mouse.SimpleManager) (ref *TagSvgAnimate) {

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
func (e *TagSvgAnimate) RemoveListener(eventType interface{}) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Mouse(value mouse.CursorType) (ref *TagSvgAnimate) {
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
func (e *TagSvgAnimate) Init(id string) (ref *TagSvgAnimate) {
	e.listener = new(sync.Map)

	e.CreateElement(KTagDiv)
	e.prepareStageReference()
	e.Id(id)

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
func (e *TagSvgAnimate) prepareStageReference() {
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
func (e *TagSvgAnimate) DragStart() (ref *TagSvgAnimate) {
	e.dragNormalStart()
	return e
}

func (e *TagSvgAnimate) DragStop() (ref *TagSvgAnimate) {
	e.dragNormalStop()
	return e
}

func (e *TagSvgAnimate) dragNormalStart() {
	e.AddListener(mouse.KEventMouseDown, e.onStartDragNormal)
	e.stage.Call("addEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
	e.stage.Call("addEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
}

func (e *TagSvgAnimate) dragNormalStop() {
	e.RemoveListener(mouse.KEventMouseDown)
	e.stage.Call("removeEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
	e.stage.Call("removeEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
	e.isDragging = false
}

func (e *TagSvgAnimate) onStopDragNormal(_ js.Value, _ []js.Value) any {
	e.isDragging = false
	return nil
}

func (e *TagSvgAnimate) onStartDragNormal(event mouse.MouseEvent) {
	var screenX = int(event.GetScreenX())
	var screenY = int(event.GetScreenY())

	e.dragDifX = screenX - e.x
	e.dragDifY = screenY - e.y

	e.isDragging = true
}

func (e *TagSvgAnimate) onMouseDraggingNormal(_ js.Value, args []js.Value) interface{} {
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

// Begin
//
// English:
//
//  The begin attribute defines when an animation should begin or when an element should be discarded.
//
// Português:
//
//  O atributo begin define quando uma animação deve começar ou quando um elemento deve ser descartado.
func (e *TagSvgAnimate) Begin(begin time.Duration) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "begin", begin.String())
	return e
}

// Dur
//
// English:
//
//  The dur attribute indicates the simple duration of an animation.
//
// Português:
//
//  O atributo dur indica a duração simples de uma animação.
func (e *TagSvgAnimate) Dur(dur time.Duration) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "dur", dur.String())
	return e
}

// End
//
// English:
//
//  The end attribute defines an end value for the animation that can constrain the active duration.
//
// Português:
//
//  O atributo final define um valor final para a animação que pode restringir a duração ativa.
func (e *TagSvgAnimate) End(end interface{}) (ref *TagSvgAnimate) {
	if converted, ok := end.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "end", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "end", end)
	return e
}

// Min
//
// English:
//
//  The min attribute specifies the minimum value of the active animation duration.
//
// Português:
//
//  O atributo min especifica o valor mínimo da duração da animação ativa.
func (e *TagSvgAnimate) Min(min time.Duration) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "min", min.String())
	return e
}

// Max
//
// English:
//
//  The max attribute specifies the maximum value of the active animation duration.
//
// Português:
//
//  O atributo max especifica o valor máximo da duração da animação ativa.
func (e *TagSvgAnimate) Max(max time.Duration) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "max", max.String())
	return e
}

// Restart
//
// English:
//
//  The restart attribute specifies whether or not an animation can restart.
//
//   Input:
//     restart: specifies whether or not an animation can restart
//       KSvgAnimationRestartAlways:This value indicates that the animation can be restarted at any time.
//       KSvgAnimationRestartWhenNotActive: This value indicates that the animation can only be restarted when it is not
//         active (i.e. after the active end).
//       KSvgAnimationRestartNever: This value indicates that the animation cannot be restarted for the time the
//         document is loaded.
//
// Português:
//
//  O atributo restart especifica se uma animação pode ou não reiniciar.
//
//   Entrada:
//     restart: especifica se uma animação pode ou não reiniciar.
//       KSvgAnimationRestartAlways: Este valor indica que a animação pode ser reiniciada a qualquer momento.
//       KSvgAnimationRestartWhenNotActive: Este valor indica que a animação só pode ser reiniciada quando não estiver
//         ativa (ou seja, após o término ativo).
//       KSvgAnimationRestartNever: Esse valor indica que a animação não pode ser reiniciada durante o carregamento do
//         documento.
func (e *TagSvgAnimate) Restart(restart SvgAnimationRestart) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "restart", restart.String())
	return e
}

// RepeatCount
//
// English:
//
//  The repeatCount attribute indicates the number of times an animation will take place.
//
//   Input:
//     repeatCount: indicates the number of times an animation will occur. (-1 for infinity)
//
// Português:
//
//  O atributo repeatCount indica o número de vezes que uma animação ocorrerá.
//
//   Entrada:
//     repeatCount: indica o número de vezes que uma animação ocorrerá. (-1 para infinito)
func (e *TagSvgAnimate) RepeatCount(repeatCount int) (ref *TagSvgAnimate) {
	if repeatCount >= 0 {
		e.selfElement.Call("setAttribute", "end", repeatCount)
		return e
	}

	e.selfElement.Call("setAttribute", "repeatCount", "indefinite")
	return e
}

// RepeatDur
//
// English:
//
//  The repeatDur attribute specifies the total duration for repeating an animation.
//
//   Input:
//     repeatDur: specifies the total duration for repeating an animation, (-1 for undefined)
//
// Português:
//
//  O atributo repeatDur especifica a duração total para repetir uma animação.
//
//   Entrada:
//     repeatDur: especifica a duração total para repetir uma animação. (-1 para indefinido)
func (e *TagSvgAnimate) RepeatDur(repeatDur interface{}) (ref *TagSvgAnimate) {
	if converted, ok := repeatDur.(time.Duration); ok {
		e.selfElement.Call("setAttribute", "repeatDur", converted.String())
		return e
	}

	e.selfElement.Call("setAttribute", "repeatDur", "indefinite")
	return e
}

// Fill
//
// English:
//
//  The fill attribute has two different meanings. For shapes and text it's a presentation attribute that defines the
//  color (or any SVG paint servers like gradients or patterns) used to paint the element;
//
// for animation it defines the final state of the animation.
//
// Português:
//
//  O atributo fill tem dois significados diferentes. Para formas e texto, é um atributo de apresentação que define a
//  cor (ou qualquer servidor de pintura SVG, como gradientes ou padrões) usado para pintar o elemento;
//
// para animação, define o estado final da animação.
func (e *TagSvgAnimate) Fill(value interface{}) (ref *TagSvgAnimate) {
	if converted, ok := value.(color.RGBA); ok {
		e.selfElement.Call("setAttribute", "fill", RGBAToJs(converted))
		return e
	}

	e.selfElement.Call("setAttribute", "fill", value)
	return e
}

// HRef
//
// English:
//
//  The href attribute defines a link to a resource as a reference URL. The exact meaning of that link depends on the
//  context of each element using it.
//
//   Notes:
//     * Specifications before SVG 2 defined an xlink:href attribute, which is now rendered obsolete by the href
//       attribute.
//       If you need to support earlier browser versions, the deprecated xlink:href attribute can be used as a fallback
//       in addition to the href attribute, e.g. <use href="some-id" xlink:href="some-id x="5" y="5" />.
//
// Português
//
//  O atributo href define um link para um recurso como um URL de referência. O significado exato desse link depende do
//  contexto de cada elemento que o utiliza.
//
//   Notas:
//     * As especificações anteriores ao SVG 2 definiam um atributo xlink:href, que agora se torna obsoleto pelo
//       atributo href.
//       Se você precisar oferecer suporte a versões anteriores do navegador, o atributo obsoleto xlink:href pode ser
//       usado como um substituto além do atributo href, por exemplo,
//       <use href="some-id" xlink:href="some-id x="5" y="5" />.
func (e *TagSvgAnimate) HRef(href string) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "href", href)
	return e
}

// AttributeName
//
// English:
//
//  The attributeName attribute indicates the name of the CSS property or attribute of the target element that is going
//  to be changed during an animation.
//
// Português
//
//  O atributo attributeName indica o nome da propriedade CSS ou atributo do elemento de destino que será alterado
//  durante uma animação.
func (e *TagSvgAnimate) AttributeName(attributeName string) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "attributeName", attributeName)
	return e
}

// CalcMode
//
// English:
//
//  The calcMode attribute specifies the interpolation mode for the animation.
//
//   Input:
//     KSvgCalcModeDiscrete: This specifies that the animation function will jump from one value to the next without
//       any interpolation.
//     KSvgCalcModeLinear: Simple linear interpolation between values is used to calculate the animation function.
//       Except for <animateMotion>, this is the default value.
//     KSvgCalcModePaced: Defines interpolation to produce an even pace of change across the animation.
//     KSvgCalcModeSpline: Interpolates from one value in the values list to the next according to a time function
//       defined by a cubic Bézier spline. The points of the spline are defined in the keyTimes attribute, and the
//       control points for each interval are defined in the keySplines attribute.
//
// The default mode is linear, however if the attribute does not support linear interpolation (e.g. for strings), the
// calcMode attribute is ignored and discrete interpolation is used.
//
//   Notes:
//     Default value: KSvgCalcModePaced
//
// Português
//
//  O atributo calcMode especifica o modo de interpolação para a animação.
//
//   Entrada:
//     KSvgCalcModeDiscrete: Isso especifica que a função de animação saltará de um valor para o próximo sem qualquer
//       interpolação.
//     KSvgCalcModeLinear: A interpolação linear simples entre valores é usada para calcular a função de animação.
//       Exceto para <animateMotion>, este é o valor padrão.
//     KSvgCalcModePaced: Define a interpolação para produzir um ritmo uniforme de mudança na animação.
//     KSvgCalcModeSpline: Interpola de um valor na lista de valores para o próximo de acordo com uma função de tempo
//       definida por uma spline de Bézier cúbica. Os pontos do spline são definidos no atributo keyTimes e os pontos
//       de controle para cada intervalo são definidos no atributo keySplines.
//
// O modo padrão é linear, no entanto, se o atributo não suportar interpolação linear (por exemplo, para strings), o
// atributo calcMode será ignorado e a interpolação discreta será usada.
//
//   Notas:
//     * Valor padrão: KSvgCalcModePaced
func (e *TagSvgAnimate) CalcMode(calcMode SvgCalcMode) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "calcMode", calcMode)
	return e
}

// Values
//
// English:
//
//  The values attribute has different meanings, depending upon the context where it's used, either it defines a
//  sequence of values used over the course of an animation, or it's a list of numbers for a color matrix, which is
//  interpreted differently depending on the type of color change to be performed.
//
// Português
//
//  O atributo values tem significados diferentes, dependendo do contexto em que é usado, ou define uma sequência de
//  valores usados ao longo de uma animação, ou é uma lista de números para uma matriz de cores, que é interpretada de
//  forma diferente dependendo da mudança de cor a ser executada.
func (e *TagSvgAnimate) Values(values string) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "values", values)
	return e
}

// KeyTimes
//
// English:
//
//  The keyTimes attribute represents a list of time values used to control the pacing of the animation.
//
// Each time in the list corresponds to a value in the values attribute list, and defines when the value is used in the
// animation. Each time value in the keyTimes list is specified as a floating point value between 0 and 1 (inclusive),
// representing a proportional offset into the duration of the animation element.
//
// Português
//
//  O atributo keyTimes representa uma lista de valores de tempo usados para controlar o ritmo da animação.
//
// Cada vez na lista corresponde a um valor na lista de atributos de valores e define quando o valor é usado na
// animação. Cada valor de tempo na lista keyTimes é especificado como um valor de ponto flutuante entre 0 e 1
// (inclusive), representando um deslocamento proporcional à duração do elemento de animação.
func (e *TagSvgAnimate) KeyTimes(keyTimes string) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "keyTimes", keyTimes)
	return e
}

// KeySplines
//
// English:
//
//  The keySplines attribute defines a set of Bézier curve control points associated with the keyTimes list, defining a
//  cubic Bézier function that controls interval pacing.
//
// This attribute is ignored unless the calcMode attribute is set to spline.
//
// If there are any errors in the keySplines specification (bad values, too many or too few values), the animation will
// not occur.
//
// Português
//
//  O atributo keySplines define um conjunto de pontos de controle da curva Bézier associados à lista keyTimes,
//  definindo uma função Bézier cúbica que controla o ritmo do intervalo.
//
// Esse atributo é ignorado, a menos que o atributo calcMode seja definido como spline.
//
// Se houver algum erro na especificação de keySplines (valores incorretos, muitos ou poucos valores), a animação não
// ocorrerá.
func (e *TagSvgAnimate) KeySplines(keySplines string) (ref *TagSvgAnimate) {
	e.CalcMode(KSvgCalcModeSpline)

	e.selfElement.Call("setAttribute", "keySplines", keySplines)
	return e
}

// From
//
// English:
//
//  The from attribute indicates the initial value of the attribute that will be modified during the animation.
//
// When used with the to attribute, the animation will change the modified attribute from the from value to the to
// value. When used with the by attribute, the animation will change the attribute relatively from the from value by
// the value specified in by.
//
// Português
//
//  O atributo from indica o valor inicial do atributo que será modificado durante a animação.
//
// Quando usado com o atributo to, a animação mudará o atributo modificado do valor from para o valor to. Quando usado
// com o atributo by, a animação mudará o atributo relativamente do valor from pelo valor especificado em by.
func (e *TagSvgAnimate) From(from float64) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "from", from)
	return e
}

// To
//
// English:
//
//  The to attribute indicates the final value of the attribute that will be modified during the animation.
//
// The value of the attribute will change between the from attribute value and this value.
//
// Português
//
//  O atributo to indica o valor final do atributo que será modificado durante a animação.
//
// O valor do atributo mudará entre o valor do atributo from e este valor.
func (e *TagSvgAnimate) To(to float64) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "to", to)
	return e
}

// By
//
// English:
//
//  The by attribute specifies a relative offset value for an attribute that will be modified during an animation.
//
// The starting value for the attribute is either indicated by specifying it as value for the attribute given in the
// attributeName or the from attribute.
//
// Português
//
//  O atributo by especifica um valor de deslocamento relativo para um atributo que será modificado durante uma
//  animação.
//
// O valor inicial para o atributo é indicado especificando-o como valor para o atributo fornecido no attributeName ou
// no atributo from.
func (e *TagSvgAnimate) By(by float64) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "by", by)
	return e
}

// Additive
//
// English:
//
//  The additive attribute controls whether or not an animation is additive.
//
//   Input:
//     KSvgAdditiveSum: Specifies that the animation will add to the underlying value of the attribute and other
//       lower priority animations.
//     KSvgAdditiveReplace: (default) Specifies that the animation will override the underlying value of the attribute
//       and other lower priority animations.
//
// It is frequently useful to define animation as an offset or delta to an attribute's value, rather than as
// absolute values.
//
// Português
//
//  O atributo aditivo controla se uma animação é ou não aditiva.
//
//   Input:
//     KSvgAdditiveSum: Especifica que a animação será adicionada ao valor subjacente do atributo e outras animações de
//       prioridade mais baixa.
//     KSvgAdditiveReplace: (default) Especifica que a animação substituirá o valor subjacente do atributo e outras
//       animações de prioridade mais baixa.
//
// É frequentemente útil definir a animação como um deslocamento ou delta para o valor de um atributo, em vez de
// valores absolutos.
func (e *TagSvgAnimate) Additive(additive SvgAdditive) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "additive", additive.String())
	return e
}

// Accumulate
//
// English:
//
//  The accumulate attribute controls whether or not an animation is cumulative.
//
//   Input:
//     KSvgAccumulateSum: Specifies that each repeat iteration after the first builds upon the last value of the
//       previous iteration;
//     KSvgAccumulateNone: Specifies that repeat iterations are not cumulative.
//
// It is frequently useful for repeated animations to build upon the previous results, accumulating with each iteration.
// This attribute said to the animation if the value is added to the previous animated attribute's value on each
// iteration.
//
//   Notes:
//     * This attribute is ignored if the target attribute value does not support addition, or if the animation element
//       does not repeat;
//     * This attribute will be ignored if the animation function is specified with only the to attribute.
//
// Português
//
//  O atributo acumular controla se uma animação é cumulativa ou não.
//
//   Input:
//     KSvgAccumulateSum: Especifica que cada iteração repetida após a primeira se baseia no último valor da iteração
//       anterior;
//     KSvgAccumulateNone: Especifica que as iterações repetidas não são cumulativas.
//
// Frequentemente, é útil que as animações repetidas se baseiem nos resultados anteriores, acumulando a cada iteração.
// Este atributo é dito à animação se o valor for adicionado ao valor do atributo animado anterior em cada iteração.
//
//   Notas:
//     * Esse atributo será ignorado se o valor do atributo de destino não suportar adição ou se o elemento de animação
//       não se repetir;
//     * Este atributo será ignorado se a função de animação for especificada apenas com o atributo to.
func (e *TagSvgAnimate) Accumulate(accumulate SvgAccumulate) (ref *TagSvgAnimate) {
	e.selfElement.Call("setAttribute", "accumulate", accumulate)
	return e
}
