package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"log"
	"strconv"
	"strings"
	"sync"
	"syscall/js"
)

// TagImg
//
// English:
//
// The <img> HTML element embeds an image into the document.
//
// Português:
//
// O elemento HTML <img> incorpora uma imagem no documento.
type TagImg struct {
	commonEvents commonEvents

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

	x          int //#replicar
	y          int //#replicar
	width      int //#replicar
	height     int //#replicar
	heightBBox int //#replicar
	bottom     int //#replicar

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
func (e *TagImg) Reference(reference **TagImg) (ref *TagImg) {
	*reference = e
	return e
}

// Init
//
// English:
//
// Initializes the object correctly.
//
// Português:
//
// Inicializa o objeto corretamente.
func (e *TagImg) Init() (ref *TagImg) {
	e.listener = new(sync.Map)
	e.tween = make(map[string]interfaces.TweenInterface)

	e.CreateElement()
	e.prepareStageReference()

	return e
}

// prepareStageReference
//
// English:
//
// Prepares the stage reference at initialization.
//
// Português:
//
// Prepara à referencia do stage na inicialização.
func (e *TagImg) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}

func (e *TagImg) Get() (el js.Value) {
	return e.selfElement
}

// AccessKey
//
// English:
//
// Specifies a shortcut key to activate/focus an element.
//
//	Input:
//	  character: A single character that specifies the shortcut key to activate/focus the element.
//
//	Note:
//	  * The accessKey attribute value must be a single character (a letter or a digit).
//	  * Adapting accessKeys to all international languages are difficult.
//	  * The accessKey value may not be present on all keyboards.
//
//
//	Warning:
//	  Using accessKeys is difficult because they may conflict with other key standards in the
//	  browser;
//	  To avoid this problem, most browsers will use accessKeys only if pressed together with the Alt
//	  key.
//
// Português:
//
// Especifica uma tecla de atalho para ativar o foco de um elemento.
//
//	 Entrada:
//		   character: Um único caractere que especifica a tecla de atalho para ativar o foco do elemento.
//
//		 Nota:
//		   * O valor do atributo accessKey deve ser um único caractere (uma letra ou um dígito).
//		   * Adaptar as teclas de acesso a todos os idiomas internacionais é difícil.
//		   * O valor accessKey pode não estar presente em todos os teclados.
//
//		 Aviso:
//		   O uso de accessKeys é difícil porque eles podem entrar em conflito com outros padrões
//		   importantes no navegador;
//		   Para evitar esse problema, a maioria dos navegadores usará as teclas de acesso somente se
//		   pressionadas junto com a tecla Alt.
func (e *TagImg) AccessKey(key string) (ref *TagImg) {
	e.selfElement.Set("accesskey", key)
	return e
}

// Autofocus
//
// English:
//
// This Boolean attribute specifies that the button should have input focus when the page loads.
// Only one element in a document can have this attribute.
//
// Português:
//
// Este atributo booleano especifica que o botão deve ter foco de entrada quando a página for
// carregada. Apenas um elemento em um documento pode ter esse atributo.
func (e *TagImg) Autofocus(autofocus bool) (ref *TagImg) {
	e.selfElement.Set("autofocus", autofocus)
	return e
}

// Class
//
// English:
//
// The class attribute specifies one or more class names for an element.
//
//	 Input:
//		   classname: Specifies one or more class names for an element. To specify multiple classes,
//		              separate the class names with a space, e.g. <span class="left important">.
//		              This allows you to combine several CSS classes for one HTML element.
//
//		              Naming rules:
//		                Must begin with a letter A-Z or a-z;
//		                Can be followed by: letters (A-Za-z), digits (0-9), hyphens ("-"), and
//		                underscores ("_").
//
// The class attribute is mostly used to point to a class in a style sheet. However, it can also be
// used by a JavaScript (via the HTML DOM) to make changes to HTML elements with a specified class.
//
// Português:
//
// O atributo class especifica um ou mais nomes de classe para um elemento.
//
//	Entrada:
//	  classname: Especifica um ou mais nomes de classe para um elemento. Para especificar várias
//	             classes, separe os nomes das classes com um espaço, por exemplo <span class="left
//	             important">.
//	             Isso permite combinar várias classes CSS para um elemento HTML.
//
//	             Regras de nomenclatura:
//	               Deve começar com uma letra A-Z ou a-z;
//	               Pode ser seguido por: letras (A-Za-z), dígitos (0-9), hífens ("-") e
//	               sublinhados ("_").
//
// O atributo class é usado principalmente para apontar para uma classe em uma folha de estilo.
// No entanto, também pode ser usado por um JavaScript (através do HTML DOM) para fazer alterações
// em elementos HTML com uma classe especificada.
func (e *TagImg) Class(class ...string) (ref *TagImg) {
	e.selfElement.Set("classList", strings.Join(class, " "))
	return e
}

// ContentEditable
//
// English:
//
// The contentEditable attribute specifies whether the content of an element is editable or not.
//
//	Input:
//	  contentEditable: specifies whether the content of an element is editable or not
//
//	Note:
//	  When the contentEditable attribute is not set on an element, the element will inherit it from
//	  its parent.
//
// Português:
//
// O atributo contentEditable especifica se o conteúdo de um elemento é editável ou não.
//
//	Entrada:
//	  contentEditable: especifica se o conteúdo de um elemento é editável ou não.
//
//	Nota:
//	  Quando o atributo contentEditable não está definido em um elemento, o elemento o herdará de
//	  seu pai.
func (e *TagImg) ContentEditable(editable bool) (ref *TagImg) {
	e.selfElement.Set("contenteditable", editable)
	return e
}

// Data
//
// English:
//
// Used to store custom data private to the page or application.
//
//	Input:
//	  data: custom data private to the page or application.
//
// The data-* attributes is used to store custom data private to the page or application.
// The data-* attributes gives us the ability to embed custom data attributes on all HTML elements.
// The stored (custom) data can then be used in the page's JavaScript to create a more engaging user
// experience (without any Ajax calls or server-side database queries).
//
// The data-* attributes consist of two parts:
//
// The attribute name should not contain any uppercase letters, and must be at least one character
// long after the prefix "data-";
// The attribute value can be any string.
//
//	Note:
//	  * Custom attributes prefixed with "data-" will be completely ignored by the user agent.
//
// Português:
//
// Usado para armazenar dados personalizados privados para a página ou aplicativo.
//
//	Entrada:
//	  data: dados personalizados privados para a página ou aplicativo.
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
// O nome do atributo não deve conter letras maiúsculas e deve ter pelo menos um caractere após o
// prefixo "data-";
// O valor do atributo pode ser qualquer string.
//
//	Nota:
//	  * Atributos personalizados prefixados com "data-" serão completamente ignorados pelo agente do
//	    usuário.
func (e *TagImg) Data(data map[string]string) (ref *TagImg) {
	for k, v := range data {
		e.selfElement.Set(" data-"+k, v)
	}
	return e
}

// Dir
//
// English:
//
// Specifies the text direction for the content in an element.
//
//	Input:
//	  dir: direction for the content in an element. [ KDirLeftToRight | KDirRightToLeft | KDirAuto ]
//
// Português:
//
// Especifica a direção do texto para o conteúdo em um elemento.
//
//	Entrada:
//	  dir: direção do texto para o conteúdo em um elemento. [ KDirLeftToRight | KDirRightToLeft |
//	       KDirAuto ]
func (e *TagImg) Dir(dir Dir) (ref *TagImg) {
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
func (e *TagImg) Draggable(draggable Draggable) (ref *TagImg) {
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
func (e *TagImg) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagImg) {
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
func (e *TagImg) Hidden() (ref *TagImg) {
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
func (e *TagImg) Id(id string) (ref *TagImg) {
	e.id = id
	e.selfElement.Set("id", id)
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
func (e *TagImg) InputMode(inputMode InputMode) (ref *TagImg) {
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
func (e *TagImg) Is(is string) (ref *TagImg) {
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
func (e *TagImg) ItemId(id string) (ref *TagImg) {
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
func (e *TagImg) ItemProp(itemprop string) (ref *TagImg) {
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
func (e *TagImg) ItemRef(itemref string) (ref *TagImg) {
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
func (e *TagImg) ItemType(itemType string) (ref *TagImg) {
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
func (e *TagImg) Lang(language Language) (ref *TagImg) {
	e.selfElement.Set("lang", language.String())
	return e
}

// Part
//
// English:
//
//	A space-separated list of the part names of the element. Part names allows CSS to select and style
//	specific elements in a shadow tree via the ::part pseudo-element.
//
// Português:
//
//	Uma lista separada por espaços dos nomes das partes do elemento. Os nomes das partes permitem que
//	o CSS selecione e estilize elementos específicos em uma árvore de sombra por meio do
//	pseudo-elemento ::part.
func (e *TagImg) Part(part ...string) (ref *TagImg) {
	e.selfElement.Set("part", strings.Join(part, " "))
	return e
}

// Nonce
//
// English:
//
// A cryptographic nonce ("number used once") which can be used by Content Security Policy to determine whether or not
// a given fetch will be allowed to proceed.
//
// Português:
//
// Um nonce criptográfico ("número usado uma vez") que pode ser usado pela Política de Segurança de Conteúdo para
// determinar se uma determinada busca terá permissão para prosseguir.
func (e *TagImg) Nonce(nonce string) (ref *TagImg) {
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
func (e *TagImg) Slot(slot string) (ref *TagImg) {
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
func (e *TagImg) Spellcheck(spell bool) (ref *TagImg) {
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
func (e *TagImg) Style(style string) (ref *TagImg) {
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
func (e *TagImg) TabIndex(index int) (ref *TagImg) {
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
func (e *TagImg) Title(title string) (ref *TagImg) {
	e.selfElement.Set("title", title)
	return e
}

func (e *TagImg) Transform(transform string) (ref *TagImg) {
	e.selfElement.Get("style").Set("transform", transform)
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
func (e *TagImg) Translate(translate Translate) (ref *TagImg) {
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
func (e *TagImg) CreateElement() (ref *TagImg) {
	e.selfElement = js.Global().Get("document").Call("createElement", "img")
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
func (e *TagImg) AppendById(appendId string) (ref *TagImg) {

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
func (e *TagImg) Append(append interface{}) (ref *TagImg) {
	switch append.(type) {
	case *TagImg:
		e.selfElement.Call("appendChild", append.(*TagImg).selfElement)
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

// Alt
//
// English:
//
//	The alt attribute provides alternative text for the image, displaying the value of the attribute
//	if the image src is missing or otherwise fails to load.
//
// Português:
//
//	O atributo alt fornece texto alternativo para a imagem, exibindo o valor do atributo se o src da
//	imagem estiver ausente ou falhar ao carregar.
func (e *TagImg) Alt(alt string) (ref *TagImg) {
	e.selfElement.Set("alt", alt)
	return e
}

// CrossOrigin
//
// English:
//
// Indicates if the fetching of the image must be done using a CORS request.
//
//	Input:
//	  value: Indicates if the fetching of the image must be done using a CORS request.
//	    const: KCrossOrigin... (e.g.: KCrossOriginAnonymous)
//	    any other type: interface{}
//
// Image data from a CORS-enabled image returned from a CORS request can be reused in the <canvas>
// element without being marked "tainted".
//
// If the crossorigin attribute is not specified, then a non-CORS request is sent (without the Origin
// request header), and the browser marks the image as tainted and restricts access to its image data,
// preventing its usage in <canvas> elements.
//
// If the crossorigin attribute is specified, then a CORS request is sent (with the Origin request
// header); but if the server does not opt into allowing cross-origin access to the image data by the
// origin site (by not sending any Access-Control-Allow-Origin response header, or by not including
// the site's origin in any Access-Control-Allow-Origin response header it does send), then the
// browser blocks the image from loading, and logs a CORS error to the devtools console.
//
// Português:
//
// Indica se a busca da imagem deve ser feita por meio de uma solicitação CORS.
//
//	Entrada:
//	  value: Indica se a busca da imagem deve ser feita por meio de uma solicitação CORS.
//	    const: KCrossOrigin... (ex.: KCrossOriginAnonymous)
//	    qualquer outro tipo: interface{}
//
// Os dados de imagem de uma imagem habilitada para CORS retornados de uma solicitação CORS podem ser
// reutilizados no elemento <canvas> sem serem marcados como "contaminados".
//
// Se o atributo crossorigin não for especificado, uma solicitação não CORS é enviada (sem o
// cabeçalho da solicitação Origin) e o navegador marca a imagem como contaminada e restringe o
// acesso aos dados da imagem, impedindo seu uso em elementos <canvas>.
//
// Se o atributo crossorigin for especificado, uma solicitação CORS será enviada (com o cabeçalho da
// solicitação Origem); mas se o servidor não permitir o acesso de origem cruzada aos dados da imagem
// pelo site de origem (não enviando nenhum cabeçalho de resposta Access-Control-Allow-Origin ou não
// incluindo a origem do site em qualquer Access-Control-Allow-Origin response header que ele envia),
// o navegador bloqueia o carregamento da imagem e registra um erro CORS no console devtools.
func (e *TagImg) CrossOrigin(value interface{}) (ref *TagImg) {
	if converted, ok := value.(CrossOrigin); ok {
		e.selfElement.Set("crossorigin", converted.String())
		return e
	}

	e.selfElement.Set("crossorigin", value)
	return e
}

// Decoding
//
// English:
//
//	Provides an image decoding hint to the browser.
//
// Português:
//
//	Fornece uma dica de decodificação de imagem para o navegador.
func (e *TagImg) Decoding(decoding Decoding) (ref *TagImg) {
	e.selfElement.Set("decoding", decoding.String())
	return e
}

// FetchPriority
//
// English:
//
//	Provides a hint of the relative priority to use when fetching the image.
//
// Português:
//
//	Fornece uma dica da prioridade relativa a ser usada ao buscar a imagem.
func (e *TagImg) FetchPriority(priority FetchPriority) (ref *TagImg) {
	e.selfElement.Set("fetchpriority", priority.String())
	return e
}

// IsMap
//
// English:
//
//	This Boolean attribute indicates that the image is part of a server-side map. If so, the
//	coordinates where the user clicked on the image are sent to the server.
//
//	 Note:
//	   * This attribute is allowed only if the <img> element is a descendant of an <a> element with a
//	     valid href attribute. This gives users without pointing devices a fallback destination.
//
// Português:
//
//	Este atributo booleano indica que a imagem faz parte de um mapa do lado do servidor.
//	Se sim, as coordenadas onde o usuário clicou na imagem são enviadas para o servidor.
//
//	 Nota:
//	   * Este atributo é permitido somente se o elemento <img> for descendente de um elemento <a> com
//	     um atributo href válido. Isso oferece aos usuários sem dispositivos apontadores um destino
//	     de fallback.
func (e *TagImg) IsMap(isMap bool) (ref *TagImg) {
	e.selfElement.Set("ismap", isMap)
	return e
}

// ReferrerPolicy
//
// English:
//
//	How much of the referrer to send when following the link.
//
//	 KRefPolicyNoReferrer: The Referer header will not be sent.
//	 KRefPolicyNoReferrerWhenDowngrade: The Referer header will not be sent to origins without TLS
//	   (HTTPS).
//	 KRefPolicyOrigin: The sent referrer will be limited to the origin of the referring page: its
//	   scheme, host, and port.
//	 KRefPolicyOriginWhenCrossOrigin: The referrer sent to other origins will be limited to the
//	   scheme, the host, and the port. Navigations on the same origin will still include the path.
//	 KRefPolicySameOrigin: A referrer will be sent for same origin, but cross-origin requests will
//	   contain no referrer information.
//	 KRefPolicyStrictOrigin: Only send the origin of the document as the referrer when the protocol
//	   security level stays the same (HTTPS→HTTPS), but don't send it to a less secure destination
//	   (HTTPS→HTTP).
//	 KRefPolicyStrictOriginWhenCrossOrigin (default): Send a full URL when performing a same-origin
//	   request, only send the origin when the protocol security level stays the same (HTTPS→HTTPS),
//	   and send no header to a less secure destination (HTTPS→HTTP).
//	 KRefPolicyUnsafeUrl: The referrer will include the origin and the path (but not the fragment,
//	   password, or username). This value is unsafe, because it leaks origins and paths from
//	   TLS-protected resources to insecure origins.
//
//	 Note:
//	   * Experimental. Expect behavior to change in the future. (04/2022)
//
// Português:
//
//	Quanto do referenciador enviar ao seguir o link.
//
//	 KRefPolicyNoReferrer: O cabeçalho Referer não será enviado.
//	 KRefPolicyNoReferrerWhenDowngrade: O cabeçalho Referer não será enviado para origens sem TLS
//	   (HTTPS).
//	 KRefPolicyOrigin: O referenciador enviado será limitado à origem da página de referência: seu
//	   esquema, host e porta.
//	 KRefPolicyOriginWhenCrossOrigin: O referenciador enviado para outras origens será limitado ao
//	   esquema, ao host e à porta. As navegações na mesma origem ainda incluirão o caminho.
//	 KRefPolicySameOrigin: Um referenciador será enviado para a mesma origem, mas as solicitações de
//	   origem cruzada não conterão informações de referenciador.
//	 KRefPolicyStrictOrigin: Só envie a origem do documento como referenciador quando o nível de
//	   segurança do protocolo permanecer o mesmo (HTTPS→HTTPS), mas não envie para um destino menos
//	   seguro (HTTPS→HTTP).
//	 KRefPolicyStrictOriginWhenCrossOrigin (padrão): Envie uma URL completa ao realizar uma
//	   solicitação de mesma origem, envie a origem apenas quando o nível de segurança do protocolo
//	   permanecer o mesmo (HTTPS→HTTPS) e não envie nenhum cabeçalho para um destino menos seguro
//	   (HTTPS→HTTP).
//	 KRefPolicyUnsafeUrl: O referenciador incluirá a origem e o caminho (mas não o fragmento, a senha
//	   ou o nome de usuário). Esse valor não é seguro porque vaza origens e caminhos de recursos
//	   protegidos por TLS para origens inseguras.
//
//	 Note:
//	   * Experimental. Expect behavior to change in the future. (04/2022)
func (e *TagImg) ReferrerPolicy(referrerPolicy ReferrerPolicy) (ref *TagImg) {
	e.selfElement.Set("referrerpolicy", referrerPolicy)
	return e
}

// Sizes
//
// English:
//
//	One or more strings separated by commas, indicating a set of source sizes. Each source size
//	consists of:
//	  1 A media condition. This must be omitted for the last item in the list.
//	  2 A source size value.
//
// Media Conditions describe properties of the viewport, not of the image. For example, (max-height:
// 500px) 1000px proposes to use a source of 1000px width, if the viewport is not higher than 500px.
//
// Source size values specify the intended display size of the image. User agents use the current
// source size to select one of the sources supplied by the srcset attribute, when those sources are
// described using width (w) descriptors. The selected source size affects the intrinsic size of the
// image (the image's display size if no CSS styling is applied). If the srcset attribute is absent,
// or contains no values with a width descriptor, then the sizes attribute has no effect.
//
// Português:
//
//	Uma ou mais strings separadas por vírgulas, indicando um conjunto de tamanhos de origem. Cada
//	tamanho de origem consiste em:
//	  1 Uma condição de mídia. Isso deve ser omitido para o último item da lista.
//	  2 Um valor de tamanho de origem.
//
// As condições de mídia descrevem as propriedades da janela de visualização, não da imagem. Por
// exemplo, (max-height: 500px) 1000px propõe usar uma fonte de 1000px de largura, se a janela de
// visualização não for maior que 500px.
//
// Os valores de tamanho de origem especificam o tamanho de exibição pretendido da imagem. Os agentes
// do usuário usam o tamanho da fonte atual para selecionar uma das fontes fornecidas pelo atributo
// srcset, quando essas fontes são descritas usando descritores de largura (w). O tamanho de origem
// selecionado afeta o tamanho intrínseco da imagem (o tamanho de exibição da imagem se nenhum estilo
// CSS for aplicado). Se o atributo srcset estiver ausente ou não contiver valores com um descritor
// de largura, o atributo tamanhos não terá efeito.
func (e *TagImg) Sizes(sizes string) (ref *TagImg) {
	e.selfElement.Set("sizes", sizes)
	return e
}

// Src
//
// English:
//
//	The image URL. Mandatory for the <img> element. On browsers supporting srcset, src is treated
//	like a candidate image with a pixel density descriptor 1x, unless an image with this pixel
//	density descriptor is already defined in srcset, or unless srcset contains w descriptors.
//
// Português:
//
//	O URL da imagem. Obrigatório para o elemento <img>. Em navegadores que suportam srcset, src é
//	tratado como uma imagem candidata com um descritor de densidade de pixels 1x, a menos que uma
//	imagem com esse descritor de densidade de pixels já esteja definida em srcset, ou a menos que
//	srcset contenha descritores w.
func (e *TagImg) Src(src string, waitLoad bool) (ref *TagImg) {
	e.selfElement.Set("src", src)

	if waitLoad == true {
		var waitGroup = new(sync.WaitGroup)

		waitGroup.Add(1)
		e.selfElement.Call(
			"addEventListener",
			"error",
			js.FuncOf(
				func(this js.Value, args []js.Value) interface{} {
					log.Print("image load error", e.id)
					waitGroup.Done()
					return nil
				},
			),
		)

		e.selfElement.Call(
			"addEventListener",
			"load",
			js.FuncOf(
				func(this js.Value, args []js.Value) interface{} {
					waitGroup.Done()
					return nil
				},
			),
		)

		waitGroup.Wait()
	}

	return e
}

// SrcSet
//
// English:
//
//	One or more strings separated by commas, indicating possible image sources for the user agent to
//	use. Each string is composed of:
//	  1 A URL to an image
//	  2 Optionally, whitespace followed by one of:
//	    * A width descriptor (a positive integer directly followed by w). The width descriptor is
//	      divided by the source size given in the sizes attribute to calculate the effective pixel
//	      density.
//	    * A pixel density descriptor (a positive floating point number directly followed by x).
//
// If no descriptor is specified, the source is assigned the default descriptor of 1x.
//
// It is incorrect to mix width descriptors and pixel density descriptors in the same srcset
// attribute. Duplicate descriptors (for instance, two sources in the same srcset which are both
// described with 2x) are also invalid.
//
// The user agent selects any of the available sources at its discretion. This provides them with
// significant leeway to tailor their selection based on things like user preferences or bandwidth
// conditions. See our Responsive images tutorial for an example.
//
// Português:
//
//	Uma ou mais strings separadas por vírgulas, indicando possíveis fontes de imagem para o agente do
//	usuário usar. Cada corda é composta por:
//	  1 Um URL para uma imagem
//	  2 Opcionalmente, espaço em branco seguido por um dos seguintes:
//	    * Um descritor de largura (um inteiro positivo seguido diretamente por w). O descritor de
//	      largura é dividido pelo tamanho da fonte fornecido no atributo de tamanhos para calcular a
//	      densidade de pixels efetiva.
//	    * Um descritor de densidade de pixels (um número de ponto flutuante positivo seguido
//	      diretamente por x).
//
// Se nenhum descritor for especificado, a origem receberá o descritor padrão de 1x.
//
// É incorreto misturar descritores de largura e descritores de densidade de pixels no mesmo atributo
// srcset. Descritores duplicados (por exemplo, duas fontes no mesmo srcset que são ambas descritas
// com 2x) também são inválidas.
//
// O agente do usuário seleciona qualquer uma das fontes disponíveis a seu critério. Isso fornece a
// eles uma margem de manobra significativa para personalizar sua seleção com base em coisas como
// preferências do usuário ou condições de largura de banda. Veja nosso tutorial de imagens
// responsivas para obter um exemplo.
func (e *TagImg) SrcSet(srcSet string) (ref *TagImg) {
	e.selfElement.Set("srcset", srcSet)
	return e
}

// Width
//
// English:
//
//	The intrinsic width of the image in pixels. Must be an integer without a unit.
//
// Português:
//
//	A largura intrínseca da imagem em pixels. Deve ser um número inteiro sem uma unidade.
func (e *TagImg) Width(width int) (ref *TagImg) {
	e.selfElement.Set("width", width)
	return e
}

func (e *TagImg) GetWidth() (width int) {
	return e.selfElement.Get("width").Int()
}

// Height
//
// English:
//
//	The intrinsic height of the image, in pixels. Must be an integer without a unit.
//
// Português:
//
//	A altura intrínseca da imagem, em pixels. Deve ser um número inteiro sem uma unidade.
func (e *TagImg) Height(height int) (ref *TagImg) {
	e.selfElement.Set("height", height)
	return e
}

func (e *TagImg) GetHeight() (height int) {
	return e.selfElement.Get("height").Int()
}

// UseMap
//
// English:
//
//	The partial URL (starting with #) of an image map associated with the element.
//
//	 Note:
//	   * You cannot use this attribute if the <img> element is inside an <a> or <button> element.
//
// Português:
//
//	The partial URL (starting with #) of an image map associated with the element.
//
//	 Note:
//	   * You cannot use this attribute if the <img> element is inside an <a> or <button> element.
func (e *TagImg) UseMap(useMap bool) (ref *TagImg) {
	e.selfElement.Set("usemap", useMap)
	return e
}

// GetXY #replicar
//
// English:
//
//	Returns the X and Y axes in pixels.
//
// Português:
//
//	Retorna os eixos X e Y em pixels.
func (e *TagImg) GetXY() (x, y int) {
	x = e.x
	y = e.y

	x = x - e.deltaMovieX
	y = y - e.deltaMovieY
	return
}

// GetX #replicar
//
// English:
//
//	Returns the X axe in pixels.
//
// Português:
//
//	Retorna o eixo X em pixels.
func (e *TagImg) GetX() (x int) {
	return e.x - e.deltaMovieX
}

// GetY #replicar
//
// English:
//
//	Returns the Y axe in pixels.
//
// Português:
//
//	Retorna o eixo Y em pixels.
func (e *TagImg) GetY() (y int) {
	return e.y - e.deltaMovieY
}

// GetTop #replicar
//
// English:
//
//	Same as GetX() function, returns the x position of the element.
//
// Português:
//
//	O mesmo que a função GetX(), retorna a posição x do elemento.
func (e *TagImg) GetTop() (top int) {
	return e.x - e.deltaMovieX
}

// GetRight #replicar
//
// English:
//
//	It is the same as x + width.
//
// Português:
//
//	É o mesmo que x + width.
func (e *TagImg) GetRight() (right int) {
	return e.x + e.width - e.deltaMovieX
}

// GetBottom #replicar
//
// English:
//
//	It is the same as y + height.
//
// Português:
//
//	É o mesmo que y + Height.
func (e *TagImg) GetBottom() (bottom int) {
	return e.y + e.height - e.deltaMovieY
}

// GetLeft #replicar
//
// English:
//
//	Same as GetY() function, returns the y position of the element.
//
// Português:
//
//	O mesmo que a função GetY(), retorna a posição y do elemento.
func (e *TagImg) GetLeft() (left int) {
	return e.y - e.deltaMovieY
}

// GetBoundingBox #replicar
//
// English:
//
// Returns the last update of the element's bounding box.
//
// Português:
//
// Retorna a última atualização do bounding box do elemnto.
func (e *TagImg) GetBoundingBox() (x, y, width, height int) {
	return e.x - e.deltaMovieX, e.y - e.deltaMovieY, e.width, e.height
}

// CollisionBoundingBox #replicar
//
// English:
//
// Detect collision between two bounding boxes.
//
// Português:
//
// Detecta colisão entre dois bounding box.
func (e *TagImg) CollisionBoundingBox(elemnt CollisionBoundingBox) (collision bool) {
	x, y, width, height := elemnt.GetBoundingBox()
	if e.x-e.deltaMovieX < x+width && e.x-e.deltaMovieX+e.width > x && e.y-e.deltaMovieY < y+height && e.y-e.deltaMovieY+e.height > y {
		return true
	}

	return false
}

// UpdateBoundingClientRect #replicar
//
// English:
//
// Updates the coordinates and dimensions of the element's bounds box.
//
// Português:
//
// Atualiza as coordenadas e as dimeções da caixa de limites do elemento.
func (e *TagImg) UpdateBoundingClientRect() (ref *TagImg) {
	// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
	//
	//                    ⋀                ⋀
	//                    |                |
	//                  y/top            bottom
	//                    |                |
	//                    ⋁                |
	// <---- x/left ----> +--------------+ | ---
	//                    |              | |   ⋀
	//                    |              | | width
	//                    |              | ⋁   ⋁
	//                    +--------------+ -----
	//                    | <- right ->  |
	// <--------- right bbox ----------> |

	bbox := e.selfElement.Call("getBoundingClientRect")
	e.x = bbox.Get("left").Int()
	e.y = bbox.Get("top").Int()
	e.heightBBox = bbox.Get("right").Int()
	e.bottom = bbox.Get("bottom").Int()

	e.height = e.heightBBox - e.x
	e.width = e.bottom - e.y

	return e
}

// SetXY #replicar
//
// English:
//
//	Sets the X and Y axes in pixels.
//
// Português:
//
//	Define os eixos X e Y em pixels.
func (e *TagImg) SetXY(x, y int) (ref *TagImg) {

	// dragging does not move delta(x,y) as the dragging function uses the delta(x,y) of mouse click
	// dragging não move delta (x,y) pois a função dragging usa o delta (x,y) do click do mouse
	if e.isDragging == true {
		e.x = x
		e.y = y
	} else {
		e.x = x + e.deltaMovieX
		e.y = y + e.deltaMovieY
	}

	px := strconv.FormatInt(int64(e.x), 10) + "px"
	py := strconv.FormatInt(int64(e.y), 10) + "px"

	e.selfElement.Get("style").Set("left", px)
	e.selfElement.Get("style").Set("top", py)
	e.selfElement.Get("style").Set("position", "absolute")

	e.UpdateBoundingClientRect()

	return e
}

// SetDeltaX
//
// English:
//
//	Additional value added in the SetX() function: (x = x + deltaMovieX) and subtracted in the
//	GetX() function: (x = x - deltaMovieX).
//
// Português:
//
//	Valor adicional adicionado na função SetX(): (x = x + deltaMovieX)  e subtraído na função
//	GetX(): (x = x - deltaMovieX).
func (e *TagImg) SetDeltaX(delta int) (ref *TagImg) {
	e.deltaMovieX = delta
	return e
}

// SetDeltaY
//
// English:
//
//	Additional value added in the SetY() function: (y = y + deltaMovieY) and subtracted in the
//	GetY() function: (y = y - deltaMovieY).
//
// Português:
//
//	Valor adicional adicionado na função SetY(): (y = y + deltaMovieY)  e subtraído na função
//	GetX(): (y = y - deltaMovieY).
func (e *TagImg) SetDeltaY(delta int) (ref *TagImg) {
	e.deltaMovieY = delta
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
func (e *TagImg) SetX(x int) (ref *TagImg) {

	// dragging does not move delta(x,y) as the dragging function uses the delta(x,y) of mouse click
	// dragging não move delta (x,y) pois a função dragging usa o delta (x,y) do click do mouse
	if e.isDragging == true {
		e.x = x
	} else {
		e.x = x + e.deltaMovieX
	}

	px := strconv.FormatInt(int64(e.x), 10) + "px"
	e.selfElement.Get("style").Set("left", px)
	e.selfElement.Get("style").Set("position", "absolute")

	e.UpdateBoundingClientRect()

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
func (e *TagImg) SetY(y int) (ref *TagImg) {

	// dragging does not move delta(x,y) as the dragging function uses the delta(x,y) of mouse click
	// dragging não move delta (x,y) pois a função dragging usa o delta (x,y) do click do mouse
	if e.isDragging == true {
		e.y = y
	} else {
		e.y = y + e.deltaMovieY
	}

	py := strconv.FormatInt(int64(e.y), 10) + "px"
	e.selfElement.Get("style").Set("top", py)
	e.selfElement.Get("style").Set("position", "absolute")

	e.UpdateBoundingClientRect()

	return e
}

func (e *TagImg) AddListenerAbort(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerAbort() (ref *TagImg) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagImg) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerAuxclick() (ref *TagImg) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagImg) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerBeforeinput() (ref *TagImg) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagImg) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerBeforematch() (ref *TagImg) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagImg) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerBeforetoggle() (ref *TagImg) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagImg) AddListenerCancel(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerCancel() (ref *TagImg) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagImg) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerCanplay() (ref *TagImg) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagImg) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerCanplaythrough() (ref *TagImg) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagImg) AddListenerChange(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerChange() (ref *TagImg) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagImg) AddListenerClick(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerClick() (ref *TagImg) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagImg) AddListenerClose(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerClose() (ref *TagImg) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagImg) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerContextlost() (ref *TagImg) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagImg) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerContextmenu() (ref *TagImg) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagImg) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerContextrestored() (ref *TagImg) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagImg) AddListenerCopy(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerCopy() (ref *TagImg) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagImg) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerCuechange() (ref *TagImg) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagImg) AddListenerCut(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerCut() (ref *TagImg) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagImg) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerDblclick() (ref *TagImg) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagImg) AddListenerDrag(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerDrag() (ref *TagImg) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagImg) AddListenerDragend(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerDragend() (ref *TagImg) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagImg) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerDragenter() (ref *TagImg) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagImg) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerDragleave() (ref *TagImg) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagImg) AddListenerDragover(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerDragover() (ref *TagImg) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagImg) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerDragstart() (ref *TagImg) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagImg) AddListenerDrop(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerDrop() (ref *TagImg) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagImg) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerDurationchange() (ref *TagImg) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagImg) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerEmptied() (ref *TagImg) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagImg) AddListenerEnded(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerEnded() (ref *TagImg) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagImg) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerFormdata() (ref *TagImg) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagImg) AddListenerInput(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerInput() (ref *TagImg) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagImg) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerInvalid() (ref *TagImg) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagImg) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerKeydown() (ref *TagImg) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagImg) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerKeypress() (ref *TagImg) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagImg) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerKeyup() (ref *TagImg) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagImg) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerLoadeddata() (ref *TagImg) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagImg) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerLoadedmetadata() (ref *TagImg) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagImg) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerLoadstart() (ref *TagImg) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagImg) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerMousedown() (ref *TagImg) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagImg) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerMouseenter() (ref *TagImg) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagImg) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerMouseleave() (ref *TagImg) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagImg) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerMousemove() (ref *TagImg) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagImg) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerMouseout() (ref *TagImg) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagImg) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerMouseover() (ref *TagImg) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagImg) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerMouseup() (ref *TagImg) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagImg) AddListenerPaste(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerPaste() (ref *TagImg) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagImg) AddListenerPause(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerPause() (ref *TagImg) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagImg) AddListenerPlay(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerPlay() (ref *TagImg) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagImg) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerPlaying() (ref *TagImg) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagImg) AddListenerProgress(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerProgress() (ref *TagImg) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagImg) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerRatechange() (ref *TagImg) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagImg) AddListenerReset(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerReset() (ref *TagImg) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagImg) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerScrollend() (ref *TagImg) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagImg) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerSecuritypolicyviolation() (ref *TagImg) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagImg) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerSeeked() (ref *TagImg) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagImg) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerSeeking() (ref *TagImg) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagImg) AddListenerSelect(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerSelect() (ref *TagImg) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagImg) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerSlotchange() (ref *TagImg) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagImg) AddListenerStalled(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerStalled() (ref *TagImg) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagImg) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerSubmit() (ref *TagImg) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagImg) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerSuspend() (ref *TagImg) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagImg) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerTimeupdate() (ref *TagImg) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagImg) AddListenerToggle(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerToggle() (ref *TagImg) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagImg) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerVolumechange() (ref *TagImg) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagImg) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerWaiting() (ref *TagImg) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagImg) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerWebkitanimationend() (ref *TagImg) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagImg) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerWebkitanimationiteration() (ref *TagImg) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagImg) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerWebkitanimationstart() (ref *TagImg) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagImg) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerWebkittransitionend() (ref *TagImg) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagImg) AddListenerWheel(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerWheel() (ref *TagImg) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagImg) AddListenerBlur(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerBlur() (ref *TagImg) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagImg) AddListenerError(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerError() (ref *TagImg) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagImg) AddListenerFocus(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerFocus() (ref *TagImg) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagImg) AddListenerLoad(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerLoad() (ref *TagImg) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagImg) AddListenerResize(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerResize() (ref *TagImg) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagImg) AddListenerScroll(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerScroll() (ref *TagImg) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagImg) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerAfterprint() (ref *TagImg) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagImg) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerBeforeprint() (ref *TagImg) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagImg) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerBeforeunload() (ref *TagImg) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagImg) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerHashchange() (ref *TagImg) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagImg) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerLanguagechange() (ref *TagImg) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagImg) AddListenerMessage(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerMessage() (ref *TagImg) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagImg) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerMessageerror() (ref *TagImg) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagImg) AddListenerOffline(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerOffline() (ref *TagImg) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagImg) AddListenerOnline(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerOnline() (ref *TagImg) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagImg) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerPageswap() (ref *TagImg) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagImg) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerPagehide() (ref *TagImg) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagImg) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerPagereveal() (ref *TagImg) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagImg) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerPageshow() (ref *TagImg) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagImg) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerPopstate() (ref *TagImg) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagImg) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerRejectionhandled() (ref *TagImg) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagImg) AddListenerStorage(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerStorage() (ref *TagImg) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagImg) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerUnhandledrejection() (ref *TagImg) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagImg) AddListenerUnload(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerUnload() (ref *TagImg) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagImg) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerReadystatechange() (ref *TagImg) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagImg) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagImg) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagImg) RemoveListenerVisibilitychange() (ref *TagImg) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
