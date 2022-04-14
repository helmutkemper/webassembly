package html

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/css"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/globalDocument"
	"log"
	"strconv"
	"strings"
	"syscall/js"
)

type Translate int

func (e Translate) String() string {
	return translateString[e]
}

const (
	// KTranslateYes
	//
	// English:
	//
	//  The translate attribute specifies whether the content of an element should be translated.
	//
	// Português:
	//
	//  O atributo translate especifica se o conteúdo de um elemento deve ser traduzido.
	KTranslateYes Translate = iota + 1

	// KTranslateNo
	//
	// English:
	//
	//  The translate attribute specifies whether the content of an element should not be translated.
	//
	// Português:
	//
	//  O atributo translate especifica se o conteúdo de um elemento não deve ser traduzido.
	KTranslateNo
)

var translateString = [...]string{
	"",
	"yes",
	"no",
}

// Dir
//
// English:
//
//  Specifies the text direction for the content in an element.
//
// Português:
//
//  Especifica a direção do texto para o conteúdo em um elemento.
type Dir int

func (e Dir) String() string {
	return dirString[e]
}

const (
	// KDirLeftToRight
	//
	// English:
	//
	//  Default. Left-to-right text direction.
	//
	// Português:
	//
	//  Padrão. Direção do texto da esquerda para a direita.
	KDirLeftToRight Dir = iota + 1

	// KDirRightToLeft
	//
	// English:
	//
	//  Right-to-left text direction.
	//
	// Português:
	//
	//  Direção do texto da direita para a esquerda.
	KDirRightToLeft

	// KDirAuto
	//
	// English:
	//
	//  Let the browser figure out the text direction, based on the content (only recommended if the
	//  text direction is unknown)
	//
	// Português:
	//
	//  Deixe o navegador descobrir a direção do texto, com base no conteúdo (recomendado apenas se a
	//  direção do texto for desconhecida)
	KDirAuto
)

var dirString = [...]string{
	"",
	"ltr",
	"rtl",
	"auto",
}

type GlobalAttributes struct {
	accessKey       string
	class           string
	contentEditable bool
	dataX           map[string]string
	dir             string
	draggable       bool
	hidden          bool
	id              string
	lang            string
	spellcheck      bool
	style           string
	tabIndex        int
	title           string
	translate       Translate
}

// SetAccessKey
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
func (e *GlobalAttributes) SetAccessKey(accessKey string) {
	e.accessKey = accessKey
}

// SetClass
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
// The class attribute is mostly used to point to a class in a style sheet. However, it can also be
// used by a JavaScript (via the HTML DOM) to make changes to HTML elements with a specified class.
func (e *GlobalAttributes) SetClass(classname ...string) {
	e.class = strings.Join(classname, " ")
}

// SetContentEditable
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
func (e *GlobalAttributes) SetContentEditable(contentEditable bool) {
	e.contentEditable = contentEditable
}

// SetDataX
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
func (e *GlobalAttributes) SetDataX(data map[string]string) {
	e.dataX = data
}

// SetDir
//
// English:
//
//  Specifies the text direction for the content in an element.
//
// Português:
//
//  Specifies the text direction for the content in an element.
func (e *GlobalAttributes) SetDir(dir Dir) {
	e.dir = dir.String()
}

// SetDraggable
//
// English:
//
//  Specifies whether an element is draggable or not.
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
//  Especifica se um elemento pode ser arrastado ou não.
//
// O atributo arrastável especifica se um elemento é arrastável ou não.
//
//   Nota:
//     * Links e imagens podem ser arrastados por padrão;
//     * O atributo arrastável é frequentemente usado em operações de arrastar e soltar.
//     * Leia nosso tutorial de arrastar e soltar HTML para saber mais.
//       https://www.w3schools.com/html/html5_draganddrop.asp
func (e *GlobalAttributes) SetDraggable(draggable bool) {
	e.draggable = draggable
}

// SetHidden
//
// English:
//
//  Specifies that an element is not yet, or is no longer, relevant.
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
func (e *GlobalAttributes) SetHidden(hidden bool) {
	e.hidden = hidden
}

// SetId
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
func (e *GlobalAttributes) SetId(id string) {
	e.id = id
}

// SetLang
//
// English:
//
//  Specifies the language of the element's content.
//
// The lang attribute specifies the language of the element's content.
//
// Common examples are "en" for English, "es" for Spanish, "fr" for French, and so on.
//
// Português:
//
//  Especifica o idioma do conteúdo do elemento.
//
// O atributo lang especifica o idioma do conteúdo do elemento.
//
// Exemplos comuns são "en" para inglês, "es" para espanhol, "fr" para francês e assim por diante.
func (e *GlobalAttributes) SetLang(lang string) {
	e.lang = lang
}

// SetSpellcheck
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
func (e *GlobalAttributes) SetSpellcheck(spellcheck bool) {
	e.spellcheck = spellcheck
}

// SetStyle
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
func (e *GlobalAttributes) SetStyle(style string) {
	e.style = style
}

// SetTabIndex
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
func (e *GlobalAttributes) SetTabIndex(tabIndex int) {
	e.tabIndex = tabIndex
}

// SetTitle
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
func (e *GlobalAttributes) SetTitle(tittle string) {
	e.title = tittle
}

// SetTranslate
//
// English:
//
// Specifies whether the content of an element should be translated or not.
//
func (e *GlobalAttributes) SetTranslate(translate Translate) {
	e.translate = translate
}

func (e *Tag) Move(x, y int) (ref *Tag) {
	px := strconv.FormatInt(int64(x), 10) + "px"
	py := strconv.FormatInt(int64(y), 10) + "py"

	e.selfElement.Get("style").Set("left", px)
	e.selfElement.Get("style").Set("top", py)

	return e
}

func (e *Tag) GetXY() (x, y int) {
	x = e.selfElement.Get("style").Get("left").Int()
	y = e.selfElement.Get("style").Get("top").Int()

	return
}

const (
	KIdToAppendNotFound    = "html.AppendById().error: id to append not found:"
	KNewElementIsUndefined = "div.NewDiv().error: new element is undefined:"
)

type Tag struct {
	id          string
	selfElement js.Value
	document    globalDocument.Document

	cursor browserMouse.CursorType
	// css recebe os nomes das classes css do elemento quando o mesmo é criado, pois, css.Class só passa
	// a funcionar quando o elemento é criado.

	css      string
	cssClass *css.Class
}

// Css
//
// English:
//
//  Is the equivalent of <... css="name1 name2 nameN">
//
//   Input:
//     classes: list of css class.
//
// Português:
//
//  Equivale a <... css="name1 name2 nameN">
//
//   Entrada:
//     classes: lista de classes css.
func (e *Tag) Css(classes ...string) (ref *Tag) {
	e.css = strings.Join(classes, " ")

	// English:
	// When a tag is called, for example NewDiv(), it already assembles the data. Therefore, this
	// function must be called whenever an html tag changes.
	//
	// Português
	// Quando uma tag é chamada, por exemplo, NewDiv(), ela já monta os dados. Por isto, esta função
	// deve ser chamada sempre que uma tag html muda.
	e.mount()

	return e
}

// mount
//
// English:
//
//  Assemble the HTML tags of each element.
//
//  Note:
//	  * When a tag is called, for example NewDiv(), it already assembles the data. Therefore, this
//	    function must be called whenever an html tag changes.
//
// Português:
//
//  Monta as tags HTML de cada elemento.
//
//  Nota:
//	  * Quando uma tag é chamada, por exemplo, NewDiv(), ela já monta os dados. Por isto, esta função
//	    deve ser chamada sempre que uma tag html muda.
func (e *Tag) mount() {

	if e.css != "" {
		e.selfElement.Set("classList", e.css)
	}

	if e.selfElement.IsUndefined() == false {

		e.selfElement.Set("style", e.cursor.String())

		if e.selfElement.Get("body").IsUndefined() == false {
			e.selfElement.Get("body").Set("style", browserMouse.KCursorAuto.String())
		}
	}
}

// SetCssController
//
// English:
//
//  Add the css classes to the created element.
//
//   Input:
//     value: object pointer to css.Class initialized
//
//   Note:
//     * This function is equivalent to css.SetList("current", classes...)
//     * Css has a feature that allows you to easily change the list of css classes of an html tag,
//       with the functions SetList(), CssToggle() and CssToggleTime();
//     * Is the equivalent of <... css="name1 name2 nameN">
//
//
// Português:
//
//  Adiciona as classes css ao elemento criado.
//
//   Entrada:
//     classes: lista de classes css.
//
//   Nota:
//     * Esta função equivale a SetList("current", classes...);
//     * Css tem uma funcionalidade que permite trocar a lista de classes css de uma tag html de forma
//       fácil, com as funções SetList(), CssToggle() e CssToggleTime();
//     * Equivale a <... css="name1 name2 nameN">
func (e *Tag) SetCssController(value *css.Class) (ref *Tag) {
	e.cssClass = value
	e.cssClass.SetRef(e.id, &e.selfElement)

	return e
}

// AppendById
//
// English:
//
// Adds a node to the end of the list of children of a specified parent node. If the node already exists in the document, it is removed from its current parent node before being added to the new parent.
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
//  Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//   Entrada:
//     appendId: id do elemento pai.
//
//   Nota:
//     * Equivale a:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
func (e Tag) AppendById(appendId string) (ref *Tag) {

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return
	}

	toAppend.Call("appendChild", e.selfElement)
	return
}

// Append
//
// English:
//
// Adds a node to the end of the list of children of a specified parent node. If the node already exists in the document, it is removed from its current parent node before being added to the new parent.
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
//  Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//   Entrada:
//     appendId: elemento no formato js.Value.
//
//   Nota:
//     * Equivale a:
//         var p = document.createElement("p");
//         document.body.appendChild(p);
func (e *Tag) Append(append interface{}) (ref *Tag) {
	append.(Tag).selfElement.Call("appendChild", e.selfElement)
	return e
}

// MousePointerAuto
//
// English:
//
//  Sets the mouse pointer to auto.
//
// Português:
//
//  Define o ponteiro do mouse como automático.
func (e *Tag) MousePointerAuto() (ref *Tag) {
	//e.selfElement.Get("body").Set("style", mouse.KCursorAuto.String())
	e.cursor = browserMouse.KCursorAuto
	e.mount()

	return e
}

// MousePointerHide
//
// English:
//
//  Sets the mouse pointer to hide.
//
// Português:
//
//  Define o ponteiro do mouse como oculto.
func (e *Tag) MousePointerHide() (ref *Tag) {
	//e.selfElement.Get("body").Set("style", mouse.KCursorNone.String())
	e.cursor = browserMouse.KCursorNone
	e.mount()

	return e
}

// SetMousePointer
//
// English:
//
//  Defines the shape of the mouse pointer.
//
//   Input:
//     value: mouse pointer shape.
//       Example: SetMousePointer(mouse.KCursorCell) // Use mouse.K... and let autocomplete do the
//                rest
//
// Português:
//
//  Define o formato do ponteiro do mouse.
//
//   Entrada:
//     V: formato do ponteiro do mouse.
//       Exemplo: SetMousePointer(mouse.KCursorCell) // Use mouse.K... e deixe o autocompletar fazer
//                o resto
func (e *Tag) SetMousePointer(value browserMouse.CursorType) (ref *Tag) {
	//e.selfElement.Get("body").Set("style", value.String())
	e.cursor = value
	e.mount()

	return e
}

// Div
//
// English:
//
//  Tag html div
//
// Português:
//
//  Tag html div
type Div struct {
	Tag
}

// NewDiv
//
// English:
//
//  Defines a section in a document, div tag.
//
//   Note:
//     * Div inherits Tag, so see Tag documentation for all functions.
//
//   Example:
//       // basic
//       var a html.Div
//       a.NewDiv("example1").
//         Css("user").
//         AppendById("stage")
//
//       // css
//       var b html.Div
//       b.NewDiv("example2").
//         Css("user").
//         SetList("red", "user", "red").
//         SetList("yellow", "user", "yellow").
//         SetList("normal", "user").
//         CssToggleTime(time.Second, "red", "yellow").
//         CssToggleLoop(10).
//         CssOnLoopEnd("normal").
//         CssToggleStart().
//         AppendById("stage")
//
// Português:
//
//  Define uma sessão do documento, uma tag div.
//
//   Nota:
//     * Div herda Tag, por isto, veja a documentação de Tag para vê todas as funções.
//
//   Exemplo:
//       // basic
//       var a html.Div
//       a.NewDiv("example1").
//         Css("user").
//         AppendById("stage")
//
//       // css
//       var b html.Div
//       b.NewDiv("example2").
//         Css("user").
//         SetList("red", "user", "red").
//         SetList("yellow", "user", "yellow").
//         SetList("normal", "user").
//         CssToggleTime(time.Second, "red", "yellow").
//         CssToggleLoop(10).
//         CssOnLoopEnd("normal").
//         CssToggleStart().
//         AppendById("stage")
func (e *Div) NewDiv(id string) (ref *Div) {
	e.id = id
	e.selfElement = js.Global().Get("document").Call("createElement", "div")
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined, id)
		return
	}

	e.selfElement.Set("id", id)
	e.mount()
	return e
}
