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

type TagLabel struct {
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

	// fnClick
	//
	// English:
	//
	// Fired when the user clicks the primary pointer button.
	//
	// Português:
	//
	// Acionado quando o usuário clica no botão do ponteiro principal.
	fnClick *js.Func

	// fnMouseOver
	//
	// English:
	//
	// Fired when a mouse or other pointing device is moved outside the element.
	//
	// Português:
	//
	// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do elemento.
	fnMouseOver *js.Func

	// fnMouseOut
	//
	// English:
	//
	// Fired when a mouse or other pointing device is moved outside the boundary of the element.
	//
	// Português:
	//
	// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do limite do elemento.
	fnMouseOut *js.Func

	// fnMouseMove
	//
	// English:
	//
	// Fired when a mouse or other pointing device is moved while over an element.
	//
	// Português:
	//
	// Acionado quando um mouse ou outro dispositivo apontador é movido sobre um elemento.
	fnMouseMove *js.Func

	// fnMouseLeave
	//
	// English:
	//
	// Fired when a mouse or other pointing device is moved outside the boundary of the element and all of its descendants.
	//
	// Português:
	//
	// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do limite do elemento e de todos os seus descendentes.
	fnMouseLeave *js.Func

	// fnMouseEnter
	//
	// English:
	//
	// Fired when a mouse or other pointing device is moved inside the boundary of the element or one of its descendants.
	//
	// Português:
	//
	// Acionado quando um mouse ou outro dispositivo apontador é movido para dentro do limite do elemento ou de um de seus descendentes.
	fnMouseEnter *js.Func

	// fnMouseDown
	//
	// English:
	//
	// Fired when the user presses a button on a mouse or other pointing device, while the pointer is over the element.
	//
	// Português:
	//
	// Acionado quando o usuário pressiona um botão em um mouse ou outro dispositivo apontador, enquanto o ponteiro está sobre o elemento.
	fnMouseDown *js.Func

	// fnMouseUp
	//
	// English:
	//
	// Fired when the user releases a button on a mouse or other pointing device, while the pointer is over the element.
	//
	// Português:
	//
	// Acionado quando o usuário libera um botão em um mouse ou outro dispositivo apontador, enquanto o ponteiro está sobre o elemento.
	fnMouseUp *js.Func

	// fnMouseWheel
	//
	// English:
	//
	// Fired when the user rotates a mouse wheel or similar user interface component such as a touchpad.
	//
	// Português:
	//
	// Acionado quando o usuário gira a roda do mouse ou um componente de interface de usuário semelhante, como um touchpad.
	fnMouseWheel *js.Func

	// fnDoubleClick
	//
	// English:
	//
	// Fired when the user double-clicks the primary pointer button.
	//
	// Português:
	//
	// Acionado quando o usuário clica duas vezes no botão do ponteiro principal.
	fnDoubleClick *js.Func
}

// Init
//
// English:
//
//	Initializes the object correctly.
//
// Português:
//
//	Inicializa o objeto corretamente.
func (e *TagLabel) Init() (ref *TagLabel) {
	e.listener = new(sync.Map)
	e.tween = make(map[string]interfaces.TweenInterface)

	e.CreateElement(KTagLabel)
	e.prepareStageReference()
	e.id = e.commonEvents.GetUuidStr()

	return e
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
func (e *TagLabel) Reference(reference **TagLabel) (ref *TagLabel) {
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
func (e *TagLabel) AccessKey(key string) (ref *TagLabel) {
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
func (e *TagLabel) Autofocus(autofocus bool) (ref *TagLabel) {
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
// O atributo class é usado principalmente para apontar para uma classe em uma folha de estilo.
// No entanto, também pode ser usado por um JavaScript (através do HTML DOM) para fazer alterações
// em elementos HTML com uma classe especificada.
func (e *TagLabel) Class(class ...string) (ref *TagLabel) {
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
func (e *TagLabel) ContentEditable(editable bool) (ref *TagLabel) {
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
func (e *TagLabel) Data(data map[string]string) (ref *TagLabel) {
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
func (e *TagLabel) Dir(dir Dir) (ref *TagLabel) {
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
func (e *TagLabel) Draggable(draggable Draggable) (ref *TagLabel) {
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
func (e *TagLabel) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagLabel) {
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
func (e *TagLabel) Hidden() (ref *TagLabel) {
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
func (e *TagLabel) Id(id string) (ref *TagLabel) {
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
func (e *TagLabel) InputMode(inputMode InputMode) (ref *TagLabel) {
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
func (e *TagLabel) Is(is string) (ref *TagLabel) {
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
func (e *TagLabel) ItemId(id string) (ref *TagLabel) {
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
func (e *TagLabel) ItemProp(itemprop string) (ref *TagLabel) {
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
func (e *TagLabel) ItemRef(itemref string) (ref *TagLabel) {
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
func (e *TagLabel) ItemType(itemType string) (ref *TagLabel) {
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
func (e *TagLabel) Lang(language Language) (ref *TagLabel) {
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
func (e *TagLabel) Part(part ...string) (ref *TagLabel) {
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
func (e *TagLabel) Nonce(nonce string) (ref *TagLabel) {
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
func (e *TagLabel) Slot(slot string) (ref *TagLabel) {
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
func (e *TagLabel) Spellcheck(spell bool) (ref *TagLabel) {
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
func (e *TagLabel) Style(style string) (ref *TagLabel) {
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
func (e *TagLabel) TabIndex(index int) (ref *TagLabel) {
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
func (e *TagLabel) Title(title string) (ref *TagLabel) {
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
func (e *TagLabel) Translate(translate Translate) (ref *TagLabel) {
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
func (e *TagLabel) CreateElement(tag Tag) (ref *TagLabel) {
	e.selfElement = js.Global().Get("document").Call("createElement", tag.String())
	if e.selfElement.IsUndefined() == true || e.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}

	return e
}

// prepareStageReference
//
// English:
//
//	Prepares the stage reference at initialization.
//
// Português:
//
//	Prepara à referencia do stage na inicialização.
func (e *TagLabel) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
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
func (e *TagLabel) AppendById(appendId string) (ref *TagLabel) {

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
func (e *TagLabel) Append(elements ...Compatible) (ref *TagLabel) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

// Form
//
// English:
//
//	The <form> element to associate the button with (its form owner). The value of this attribute must
//	be the id of a <form> in the same document. (If this attribute is not set, the <button> is
//	associated with its ancestor <form> element, if any.)
//
// This attribute lets you associate <button> elements to <form>s anywhere in the document, not just
// inside a <form>. It can also override an ancestor <form> element.
//
// Português:
//
//	O elemento <form> ao qual associar o botão (seu proprietário do formulário). O valor deste
//	atributo deve ser o id de um <form> no mesmo documento. (Se esse atributo não for definido, o
//	<button> será associado ao elemento <form> ancestral, se houver.)
//
// Este atributo permite associar elementos <button> a <form>s em qualquer lugar do documento, não
// apenas dentro de um <form>. Ele também pode substituir um elemento <form> ancestral.
func (e *TagLabel) Form(form string) (ref *TagLabel) {
	e.selfElement.Set("form", form)
	return e
}

// Text #replicar
//
// English:
//
// Adds plain text to the tag's content.
//
// Text:
//
// Adiciona um texto simples ao conteúdo da tag.
func (e *TagLabel) Text(value any) (ref *TagLabel) {
	e.selfElement.Set("textContent", value)
	return e
}

// Html #replicar
//
// English:
//
// Adds HTML to the tag's content.
//
// Text:
//
// Adiciona HTML ao conteúdo da tag.
func (e *TagLabel) Html(value string) (ref *TagLabel) {
	e.selfElement.Set("innerHTML", value)
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
func (e *TagLabel) GetXY() (x, y int) {
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
func (e *TagLabel) GetX() (x int) {
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
func (e *TagLabel) GetY() (y int) {
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
func (e *TagLabel) GetTop() (top int) {
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
func (e *TagLabel) GetRight() (right int) {
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
func (e *TagLabel) GetBottom() (bottom int) {
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
func (e *TagLabel) GetLeft() (left int) {
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
func (e *TagLabel) GetBoundingBox() (x, y, width, height int) {
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
func (e *TagLabel) CollisionBoundingBox(elemnt CollisionBoundingBox) (collision bool) {
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
func (e *TagLabel) UpdateBoundingClientRect() (ref *TagLabel) {
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
func (e *TagLabel) SetXY(x, y int) (ref *TagLabel) {

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
func (e *TagLabel) SetDeltaX(delta int) (ref *TagLabel) {
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
func (e *TagLabel) SetDeltaY(delta int) (ref *TagLabel) {
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
func (e *TagLabel) SetX(x int) (ref *TagLabel) {

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
func (e *TagLabel) SetY(y int) (ref *TagLabel) {

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

func (e *TagLabel) Get() (el js.Value) {
	return e.selfElement
}

func (e *TagLabel) AddListenerAbort(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerAbort() (ref *TagLabel) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagLabel) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerAuxclick() (ref *TagLabel) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagLabel) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerBeforeinput() (ref *TagLabel) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagLabel) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerBeforematch() (ref *TagLabel) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagLabel) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerBeforetoggle() (ref *TagLabel) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagLabel) AddListenerCancel(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerCancel() (ref *TagLabel) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagLabel) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerCanplay() (ref *TagLabel) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagLabel) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerCanplaythrough() (ref *TagLabel) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagLabel) AddListenerChange(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerChange() (ref *TagLabel) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagLabel) AddListenerClick(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerClick() (ref *TagLabel) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagLabel) AddListenerClose(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerClose() (ref *TagLabel) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagLabel) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerContextlost() (ref *TagLabel) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagLabel) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerContextmenu() (ref *TagLabel) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagLabel) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerContextrestored() (ref *TagLabel) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagLabel) AddListenerCopy(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerCopy() (ref *TagLabel) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagLabel) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerCuechange() (ref *TagLabel) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagLabel) AddListenerCut(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerCut() (ref *TagLabel) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagLabel) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerDblclick() (ref *TagLabel) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagLabel) AddListenerDrag(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerDrag() (ref *TagLabel) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagLabel) AddListenerDragend(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerDragend() (ref *TagLabel) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagLabel) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerDragenter() (ref *TagLabel) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagLabel) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerDragleave() (ref *TagLabel) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagLabel) AddListenerDragover(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerDragover() (ref *TagLabel) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagLabel) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerDragstart() (ref *TagLabel) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagLabel) AddListenerDrop(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerDrop() (ref *TagLabel) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagLabel) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerDurationchange() (ref *TagLabel) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagLabel) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerEmptied() (ref *TagLabel) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagLabel) AddListenerEnded(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerEnded() (ref *TagLabel) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagLabel) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerFormdata() (ref *TagLabel) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagLabel) AddListenerInput(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerInput() (ref *TagLabel) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagLabel) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerInvalid() (ref *TagLabel) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagLabel) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerKeydown() (ref *TagLabel) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagLabel) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerKeypress() (ref *TagLabel) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagLabel) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerKeyup() (ref *TagLabel) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagLabel) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerLoadeddata() (ref *TagLabel) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagLabel) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerLoadedmetadata() (ref *TagLabel) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagLabel) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerLoadstart() (ref *TagLabel) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagLabel) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerMousedown() (ref *TagLabel) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagLabel) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerMouseenter() (ref *TagLabel) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagLabel) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerMouseleave() (ref *TagLabel) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagLabel) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerMousemove() (ref *TagLabel) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagLabel) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerMouseout() (ref *TagLabel) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagLabel) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerMouseover() (ref *TagLabel) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagLabel) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerMouseup() (ref *TagLabel) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagLabel) AddListenerPaste(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerPaste() (ref *TagLabel) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagLabel) AddListenerPause(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerPause() (ref *TagLabel) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagLabel) AddListenerPlay(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerPlay() (ref *TagLabel) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagLabel) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerPlaying() (ref *TagLabel) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagLabel) AddListenerProgress(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerProgress() (ref *TagLabel) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagLabel) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerRatechange() (ref *TagLabel) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagLabel) AddListenerReset(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerReset() (ref *TagLabel) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagLabel) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerScrollend() (ref *TagLabel) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagLabel) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerSecuritypolicyviolation() (ref *TagLabel) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagLabel) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerSeeked() (ref *TagLabel) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagLabel) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerSeeking() (ref *TagLabel) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagLabel) AddListenerSelect(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerSelect() (ref *TagLabel) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagLabel) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerSlotchange() (ref *TagLabel) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagLabel) AddListenerStalled(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerStalled() (ref *TagLabel) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagLabel) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerSubmit() (ref *TagLabel) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagLabel) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerSuspend() (ref *TagLabel) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagLabel) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerTimeupdate() (ref *TagLabel) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagLabel) AddListenerToggle(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerToggle() (ref *TagLabel) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagLabel) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerVolumechange() (ref *TagLabel) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagLabel) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerWaiting() (ref *TagLabel) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagLabel) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerWebkitanimationend() (ref *TagLabel) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagLabel) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerWebkitanimationiteration() (ref *TagLabel) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagLabel) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerWebkitanimationstart() (ref *TagLabel) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagLabel) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerWebkittransitionend() (ref *TagLabel) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagLabel) AddListenerWheel(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerWheel() (ref *TagLabel) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagLabel) AddListenerBlur(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerBlur() (ref *TagLabel) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagLabel) AddListenerError(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerError() (ref *TagLabel) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagLabel) AddListenerFocus(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerFocus() (ref *TagLabel) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagLabel) AddListenerLoad(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerLoad() (ref *TagLabel) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagLabel) AddListenerResize(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerResize() (ref *TagLabel) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagLabel) AddListenerScroll(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerScroll() (ref *TagLabel) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagLabel) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerAfterprint() (ref *TagLabel) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagLabel) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerBeforeprint() (ref *TagLabel) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagLabel) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerBeforeunload() (ref *TagLabel) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagLabel) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerHashchange() (ref *TagLabel) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagLabel) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerLanguagechange() (ref *TagLabel) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagLabel) AddListenerMessage(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerMessage() (ref *TagLabel) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagLabel) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerMessageerror() (ref *TagLabel) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagLabel) AddListenerOffline(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerOffline() (ref *TagLabel) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagLabel) AddListenerOnline(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerOnline() (ref *TagLabel) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagLabel) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerPageswap() (ref *TagLabel) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagLabel) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerPagehide() (ref *TagLabel) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagLabel) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerPagereveal() (ref *TagLabel) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagLabel) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerPageshow() (ref *TagLabel) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagLabel) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerPopstate() (ref *TagLabel) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagLabel) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerRejectionhandled() (ref *TagLabel) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagLabel) AddListenerStorage(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerStorage() (ref *TagLabel) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagLabel) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerUnhandledrejection() (ref *TagLabel) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagLabel) AddListenerUnload(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerUnload() (ref *TagLabel) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagLabel) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerReadystatechange() (ref *TagLabel) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagLabel) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagLabel) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagLabel) RemoveListenerVisibilitychange() (ref *TagLabel) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
