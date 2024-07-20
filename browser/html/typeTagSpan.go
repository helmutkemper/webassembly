package html

import (
	"github.com/helmutkemper/webassembly/browser/css"
	"github.com/helmutkemper/webassembly/browser/event/mouse"
	"github.com/helmutkemper/webassembly/interfaces"
	"github.com/helmutkemper/webassembly/platform/algorithm"
	"log"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"syscall/js"
)

// TagSpan
//
// English:
//
//	The <span> HTML element is a generic inline container for phrasing content, which does not inherently represent anything.
//
//	It can be used to group elements for styling purposes (using the class or id attributes), or because they share attribute values, such as lang. It should be used only when no other semantic element is appropriate. <span> is very much like a <div> element, but <div> is a block-level element whereas a <span> is an inline-level element.
//
// Português:
//
//	O elemento HTML <span> é um contêiner embutido genérico para frasear conteúdo, que não representa nada inerentemente.
//
//	Ele pode ser usado para agrupar elementos para fins de estilo (usando os atributos class ou id) ou porque eles compartilham valores de atributos, como lang. Deve ser usado somente quando nenhum outro elemento semântico for apropriado. <span> é muito parecido com um elemento <div>, mas <div> é um elemento de nível de bloco, enquanto <span> é um elemento de nível embutido.
type TagSpan struct {
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
func (e *TagSpan) Reference(reference **TagSpan) (ref *TagSpan) {
	*reference = e
	return e
}

// #global - start -----------------------------------------------------------------------------------------------------

// AccessKey #global
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
func (e *TagSpan) AccessKey(key string) (ref *TagSpan) {
	e.selfElement.Set("accesskey", key)
	return e
}

// AutoCapitalize #global
//
// English:
//
// Controls whether and how text input is automatically capitalized as it is entered/edited by the user.
//
// Português:
//
// Controla se e como a entrada de texto é colocada em maiúsculas automaticamente à medida que é inserida e editada
// pelo usuário.
func (e *TagSpan) AutoCapitalize(value interface{}) (ref *TagSpan) {
	if converted, ok := value.(AutoCapitalize); ok {
		e.selfElement.Set("autocapitalize", converted.String())
		return e
	}

	e.selfElement.Set("autocapitalize", value)
	return e
}

// AutoFocus #global
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
func (e *TagSpan) AutoFocus(autofocus bool) (ref *TagSpan) {
	e.selfElement.Set("autofocus", autofocus)
	return e
}

// Class #global
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
func (e *TagSpan) Class(class ...string) (ref *TagSpan) {
	e.selfElement.Set("classList", strings.Join(class, " "))
	return e
}

// ContentEditable #global
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
func (e *TagSpan) ContentEditable(editable bool) (ref *TagSpan) {
	e.selfElement.Set("contenteditable", editable)
	return e
}

// Data #global #replicar
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
func (e *TagSpan) Data(data map[string]string) (ref *TagSpan) {
	for k, v := range data {
		e.selfElement.Set("data-"+k, v)
	}
	return e
}

// GetData #global #replicar
//
// English:
//
//		Used to get custom data private to the page or application.
//
//		 Input:
//		   key: custom key of data to get
//
//	  Output:
//	    value: value of custom data
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
//		Usado para armazenar dados personalizados privados para a página ou aplicativo.
//
//		 Entrada:
//		   key: chave personalizada de dados para obter
//
//	  Saída:
//	    value: valor do dado personalizado
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
func (e *TagSpan) GetData(key string) (value string) {
	return e.selfElement.Get("data-" + key).String()
}

// Dir #global
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
func (e *TagSpan) Dir(dir Dir) (ref *TagSpan) {
	e.selfElement.Set("dir", dir.String())
	return e
}

// Draggable #global
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
func (e *TagSpan) Draggable(draggable Draggable) (ref *TagSpan) {
	e.selfElement.Set("draggable", draggable.String())
	return e
}

// EnterKeyHint #global
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
func (e *TagSpan) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagSpan) {
	e.selfElement.Set("enterKeyHint", enterKeyHint.String())
	return e
}

// Hidden #global
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
func (e *TagSpan) Hidden() (ref *TagSpan) {
	e.selfElement.Get("style").Set("visibility", "hidden")
	return e
}

// Id #global
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
func (e *TagSpan) Id(id string) (ref *TagSpan) {
	e.id = id
	e.selfElement.Set("id", id)
	return e
}

// ItemProp #global
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
func (e *TagSpan) ItemProp(itemprop string) (ref *TagSpan) {
	e.selfElement.Set("itemprop", itemprop)
	return e
}

// ItemRef #global
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
func (e *TagSpan) ItemRef(itemref string) (ref *TagSpan) {
	e.selfElement.Set("itemref", itemref)
	return e
}

// ItemScope #global
//
// English:
//
// itemscope (usually) works along with itemtype to specify that the HTML contained in a block is about a particular
// item.
//
// itemscope creates the Item and defines the scope of the itemtype associated with it. itemtype is a valid URL of a
// vocabulary (such as schema.org) that describes the item and its properties context.
//
// Português:
//
// itemscope (geralmente) funciona junto com itemtype para especificar que o HTML contido em um bloco é sobre um item
// específico.
//
// itemscope cria o Item e define o escopo do tipo de item associado a ele. itemtype é uma URL válida de um vocabulário
// (como schema.org) que descreve o item e seu contexto de propriedades.
func (e *TagSpan) ItemScope(itemscope bool) (ref *TagSpan) {
	e.selfElement.Set("itemscope", itemscope)
	return e
}

// ItemType #global
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
func (e *TagSpan) ItemType(itemType string) (ref *TagSpan) {
	e.selfElement.Set("itemtype", itemType)
	return e
}

// Lang #global
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
func (e *TagSpan) Lang(language Language) (ref *TagSpan) {
	e.selfElement.Set("lang", language.String())
	return e
}

// Part #global
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
func (e *TagSpan) Part(part ...string) (ref *TagSpan) {
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
func (e *TagSpan) Nonce(nonce string) (ref *TagSpan) {
	e.selfElement.Set("nonce", nonce)
	return e
}

// Slot #global
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
func (e *TagSpan) Slot(slot string) (ref *TagSpan) {
	e.selfElement.Set("slot", slot)
	return e
}

// Spellcheck #global
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
func (e *TagSpan) Spellcheck(spell bool) (ref *TagSpan) {
	e.selfElement.Set("spellcheck", spell)

	return e
}

// Style #global
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
func (e *TagSpan) Style(style string) (ref *TagSpan) {
	e.selfElement.Set("style", style)
	return e
}

// TabIndex #global
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
func (e *TagSpan) TabIndex(index int) (ref *TagSpan) {
	e.selfElement.Set("tabindex", index)
	return e
}

// Title #global
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
func (e *TagSpan) Title(title string) (ref *TagSpan) {
	e.selfElement.Set("title", title)
	return e
}

// Translate #global
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
func (e *TagSpan) Translate(translate Translate) (ref *TagSpan) {
	e.selfElement.Set("translate", translate.String())
	return e
}

// #global - end -------------------------------------------------------------------------------------------------------

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
func (e *TagSpan) InputMode(inputMode InputMode) (ref *TagSpan) {
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
func (e *TagSpan) Is(is string) (ref *TagSpan) {
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
func (e *TagSpan) ItemId(id string) (ref *TagSpan) {
	e.selfElement.Set("itemid", id)
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
func (e *TagSpan) CreateElement(tag Tag) (ref *TagSpan) {
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
func (e *TagSpan) AppendById(appendId string) (ref *TagSpan) {

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
//
// fixme: fazer append() assim em todas as tags html, exceto svg
func (e *TagSpan) Append(elements ...Compatible) (ref *TagSpan) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfElement.Call("appendChild", fragment)
	return e
}

// AppendToStage
//
// English:
//
//	Adds a node to the end of the list of children in the main document body. If the node already
//	exists somewhere in the document, it is removed from its current parent node before being added
//	to the main document.
//
// Português:
//
//	Adiciona um nó ao final da lista de filhos do corpo do documento principal. Se o nó já existir
//	em alguma parte do documento, ele é removido de seu nó pai atual antes de ser adicionado ao
//	documento principal.
//
// todo:https://developer.mozilla.org/en-US/docs/Web/API/Document/createDocumentFragment
// todo: appendMany()
func (e *TagSpan) AppendToStage() (ref *TagSpan) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSpan) Get() (el js.Value) {
	return e.selfElement
}

// Rotate
//
// English:
//
//	Defines a transformation that rotates an element around a fixed point on the 2D plane, without deforming it.
//
//	 Input:
//	   angle: representing the angle of the rotation. The direction of rotation depends on the writing direction.
//	   In a left-to-right context, a positive angle denotes a clockwise rotation, a negative angle a counter-clockwise
//	   one.
//	   In a right-to-left context, a positive angle denotes a counter-clockwise rotation, a negative angle a clockwise
//	   one.
//
// Português:
//
//	Define uma transformação que gira um elemento em torno de um ponto fixo no plano 2D, sem deformá-lo.
//
//	 Entrada:
//	   angle: representando o ângulo de rotação. O sentido de rotação depende do sentido de escrita.
//	   Em um contexto da esquerda para a direita, um ângulo positivo denota uma rotação no sentido horário, um ângulo
//	   negativo no sentido anti-horário.
//	   Em um contexto da direita para a esquerda, um ângulo positivo denota uma rotação no sentido anti-horário, um
//	   ângulo negativo denota uma rotação no sentido horário.
func (e *TagSpan) Rotate(angle float64) (ref *TagSpan) {
	angleAsString := strconv.FormatFloat(angle+e.rotateDelta, 'E', -1, 64)
	e.selfElement.Get("style").Set("transform", "rotate("+angleAsString+"rad)")
	return e
}

// RotateDelta
//
// English:
//
//	Used in conjunction with the Rotate() function, sets the rotation adjustment angle, ie Rotate() = angle + delta.
//
//	 Input:
//	   angle: delta, object rotation adjustment angle.
//
// Português:
//
//	Usada em conjunto com a função Rotate(), define o ângulo de ajuste da rotação, ou seja, Rotate() = angle + delta.
//
//	 Entrada:
//	   angle: delta, ângulo de ajuste da rotação do objeto.
func (e *TagSpan) RotateDelta(delta float64) (ref *TagSpan) {
	e.rotateDelta = delta
	return e
}

// GetRotateDelta
//
// English:
//
//	Returns the rotation adjustment angle, i.e. Rotate() = angle + delta.
//
//	 Output:
//	   angle: delta, object rotation adjustment angle.
//
// Português:
//
//	Retorna o ângulo de ajuste da rotação, ou seja, Rotate() = angle + delta.
//
//	 Saída:
//	   angle: delta, ângulo de ajuste da rotação do objeto.
func (e *TagSpan) GetRotateDelta() (delta float64) {
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
//func (e *TagSpan) AddListener(eventType interface{}, manager mouse.SimpleManager) (ref *TagSpan) {
//
//	mouseMoveEvt := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
//		var mouseEvent = mouse.Event{}
//
//		if len(args) > 0 {
//			mouseEvent.Object = args[0]
//		}
//
//		if manager != nil {
//			manager(mouseEvent)
//		}
//
//		return nil
//	})
//
//	switch converted := eventType.(type) {
//	case event.Event:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventAnimation.EventAnimation:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventClipBoard.EventClipBoard:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventDrag.EventDrag:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventFocus.EventFocus:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventHashChange.EventHashChange:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventInput.EventInput:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventKeyboard.EventKeyboard:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	//case mouse.Event:
//	//	e.listener.Store(converted.String(), mouseMoveEvt)
//	//	e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventPageTransition.EventPageTransition:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventUi.EventUi:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventWheel.EventWheel:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfElement.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	default:
//		log.Fatalf("event must be a event type")
//	}
//
//	return e
//}

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
//func (e *TagSpan) RemoveListener(eventType interface{}) (ref *TagSpan) {
//	switch converted := eventType.(type) {
//	case event.Event:
//		f, _ := e.listener.Load(converted.String())
//		e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	case eventAnimation.EventAnimation:
//		f, _ := e.listener.Load(converted.String())
//		e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	case eventClipBoard.EventClipBoard:
//		f, _ := e.listener.Load(converted.String())
//		e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	case eventDrag.EventDrag:
//		f, _ := e.listener.Load(converted.String())
//		e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	case eventFocus.EventFocus:
//		f, _ := e.listener.Load(converted.String())
//		e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	case eventHashChange.EventHashChange:
//		f, _ := e.listener.Load(converted.String())
//		e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	case eventInput.EventInput:
//		f, _ := e.listener.Load(converted.String())
//		e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	case eventKeyboard.EventKeyboard:
//		f, _ := e.listener.Load(converted.String())
//		e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	//case mouse.Event:
//	//	f, _ := e.listener.Load(converted.String())
//	//	e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	case eventPageTransition.EventPageTransition:
//		f, _ := e.listener.Load(converted.String())
//		e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	case eventUi.EventUi:
//		f, _ := e.listener.Load(converted.String())
//		e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	case eventWheel.EventWheel:
//		f, _ := e.listener.Load(converted.String())
//		e.selfElement.Call("removeEventListener", converted.String(), f)
//
//	default:
//		log.Fatalf("event must be a event type")
//	}
//
//	return e
//}

// Mouse
//
// English:
//
//	Defines the shape of the mouse pointer.
//
//	 Input:
//	   value: mouse pointer shape.
//	     Example: SetMouse(mouse.KCursorCell) // Use mouse.K... and let autocomplete do the
//	              rest
//
// Português:
//
//	Define o formato do ponteiro do mouse.
//
//	 Entrada:
//	   value: formato do ponteiro do mouse.
//	     Exemplo: SetMouse(mouse.KCursorCell) // Use mouse.K... e deixe o autocompletar fazer
//	              o resto
func (e *TagSpan) Mouse(value mouse.CursorType) (ref *TagSpan) {
	e.selfElement.Get("style").Set("cursor", value.String())
	return e
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
func (e *TagSpan) Init() (ref *TagSpan) {
	e.listener = new(sync.Map)
	e.tween = make(map[string]interfaces.TweenInterface)

	e.CreateElement(KTagSpan)
	e.prepareStageReference()
	e.id = e.commonEvents.GetUuidStr()

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
func (e *TagSpan) prepareStageReference() {
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
//     factoryBrowser.NewTagSpan("div_0").
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
//     factoryBrowser.NewTagSpan("div_0").
//       Class("animate").
//       DragStart().
//       AppendById("stage")
//func (e *TagSpan) DragStart() (ref *TagSpan) {
//	e.dragNormalStart()
//	return e
//}

// DragStop
//
// English:
//
//  Stop mouse drag functionality.
//
//   Example:
//
//     factoryBrowser.NewTagSpan("div_0").
//       Class("animate").
//       DragStart().
//       AppendById("stage")
//
//     go func() {
//       time.Sleep(10 * time.Second)
//       div.DragStop()
//     }()
//
// Português:
//
//  Para a funcionalidade de arrastar com o mouse.
//
//   Exemplo:
//
//     factoryBrowser.NewTagSpan("div_0").
//       Class("animate").
//       DragStart().
//       AppendById("stage")
//
//     go func() {
//       time.Sleep(10 * time.Second)
//       div.DragStop()
//     }()
//func (e *TagSpan) DragStop() (ref *TagSpan) {
//	e.dragNormalStop()
//	return e
//}

//func (e *TagSpan) dragNormalStart() {
//	e.AddListener(mouse.KEventMouseDown, e.onStartDragNormal)
//	e.stage.Call("addEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
//	e.stage.Call("addEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
//}
//
//func (e *TagSpan) dragNormalStop() {
//	e.RemoveListener(mouse.KEventMouseDown)
//	e.stage.Call("removeEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
//	e.stage.Call("removeEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
//	e.isDragging = false
//}

func (e *TagSpan) onStopDragNormal(_ js.Value, _ []js.Value) interface{} {
	e.isDragging = false
	return nil
}

//func (e *TagSpan) onStartDragNormal(event mouse.MouseEvent) {
//	var screenX = int(event.GetScreenX())
//	var screenY = int(event.GetScreenY())
//
//	e.dragDifX = screenX - e.x
//	e.dragDifY = screenY - e.y
//
//	e.isDragging = true
//}

//func (e *TagSpan) onMouseDraggingNormal(_ js.Value, args []js.Value) interface{} {
//	if e.isDragging == false {
//		return nil
//	}
//
//	var mouseEvent = mouse.MouseEvent{}
//	if len(args) > 0 {
//		mouseEvent.Object = args[0]
//
//		var x = int(mouseEvent.GetScreenX()) - e.dragDifX
//		var y = int(mouseEvent.GetScreenY()) - e.dragDifY
//
//		e.SetXY(x, y)
//	}
//
//	return nil
//}

// AddPointsToEasingTween
//
// English:
//
//	This function returns an easing tween function compatible with the easing onStepFunc() function in order to use the
//	points generated by the line algorithms as a reference to the movement.
//
//	 Note:
//	   * The 'onStartValue' and 'onEndValue' parameters must have the values 0 and 10000.
//	     Example:
//	       factoryEasingTween.NewLinear(5*time.Second, 0, 10000, div.EasingTweenWalkingAndRotateIntoPoints(), 0)
//
// Português:
//
//	Esta função retorna uma função easing tween compatível com a função onStepFunc() do easing de modo a usar os pontos
//	gerados pelos algoritmos de linha como referência ao movimento.
//
//	 Nota:
//	   * O parâmetros 'onStartValue' e 'onEndValue' devem, obrigatoriamente, ter os valores 0 e 10000.
//	     Exemplo:
//	       factoryEasingTween.NewLinear(5*time.Second, 0, 10000, div.EasingTweenWalkingAndRotateIntoPoints(), 0)
func (e *TagSpan) AddPointsToEasingTween(algorithmRef algorithm.CurveInterface) (ref *TagSpan) {
	if algorithmRef == nil {
		return e
	}

	var points = algorithmRef.GetProcessed()

	e.points = points
	e.pointsLen = len(*points)

	return e
}

// EasingTweenWalkingIntoPoints
//
// English:
//
//	Moves the element on the line added by the AddPointsToEasingTween() function.
//
// This function returns a second function compatible with the easing tween's onStepFunc() function.
//
//	Note:
//	  * The 'onStartValue' and 'onEndValue' parameters must have the values 0 and 10000.
//	    Example:
//	      factoryEasingTween.NewLinear(5*time.Second, 0, 10000, div.EasingTweenWalkingAndRotateIntoPoints(), 0)
//
// Português:
//
//	Desloca o elemento na linha adicionada pela função AddPointsToEasingTween().
//
// Esta função retorna uma segunda função compatível com a função onStepFunc() do easing tween.
//
//	Nota:
//	  * O parâmetros 'onStartValue' e 'onEndValue' devem, obrigatoriamente, ter os valores 0 e 10000.
//	    Exemplo:
//	      factoryEasingTween.NewLinear(5*time.Second, 0, 10000, div.EasingTweenWalkingAndRotateIntoPoints(), 0)
func (e *TagSpan) EasingTweenWalkingIntoPoints() (function func(percent, p float64, args interface{})) {

	function = func(forTenThousand, percent float64, args interface{}) {

		if forTenThousand > 10000.0 {
			forTenThousand = forTenThousand - 10000.0
		} else if forTenThousand < 0.0 {
			forTenThousand = 10000.0 + forTenThousand
		}

		pCalc := int(float64(e.pointsLen) * forTenThousand / 10000.0)
		e.SetXY(int((*e.points)[pCalc].X), int((*e.points)[pCalc].Y))
	}

	return
}

// EasingTweenWalkingAndRotateIntoPoints
//
// English:
//
//	Moves the element on the line added by the AddPointsToEasingTween() function and adjusts the rotation of the
//	element with respect to the next point.
//
// This function returns a second function compatible with the easing tween's onStepFunc() function.
//
//	Note:
//	  * Use the RotateDelta() function to adjust the starting angle;
//	  * The 'onStartValue' and 'onEndValue' parameters must have the values 0 and 10000.
//	    Example:
//	      factoryEasingTween.NewLinear(5*time.Second, 0, 10000, div.EasingTweenWalkingAndRotateIntoPoints(), 0)
//
// Português:
//
//	Desloca o elemento na linha adicionada pela função AddPointsToEasingTween() e ajusta a rotação do elemento em relação ao próximo ponto.
//
// Esta função retorna uma segunda função compatível com a função onStepFunc() do easing tween.
//
//	Nota:
//	  * Use a função RotateDelta() para ajustar o ângulo inicial;
//	  * O parâmetros 'onStartValue' e 'onEndValue' devem, obrigatoriamente, ter os valores 0 e 10000.
//	    Exemplo:
//	      factoryEasingTween.NewLinear(5*time.Second, 0, 10000, div.EasingTweenWalkingAndRotateIntoPoints(), 0)
func (e *TagSpan) EasingTweenWalkingAndRotateIntoPoints() (function func(forTenThousand, percent float64, args interface{})) {

	function = func(forTenThousand, percent float64, args interface{}) {

		angleCorrection := false

		if forTenThousand > 10000.0 {
			forTenThousand = forTenThousand - 10000.0
			angleCorrection = true
		} else if forTenThousand < 0.0 {
			forTenThousand = 10000.0 + forTenThousand
			angleCorrection = true
		}

		pCalc := int(float64(e.pointsLen) * forTenThousand / 10000.0)

		var angle float64
		switch pCalc {
		case 0.0:
			if angleCorrection == false {
				angle = math.Atan2((*e.points)[0].Y-(*e.points)[1].Y, (*e.points)[0].X-(*e.points)[1].X)
			} else {
				angle = math.Atan2((*e.points)[1].Y-(*e.points)[0].Y, (*e.points)[1].X-(*e.points)[0].X)
			}

		case 1.0:
			if angleCorrection == true {
				angle = math.Atan2((*e.points)[pCalc].Y-(*e.points)[pCalc-1].Y, (*e.points)[pCalc].X-(*e.points)[pCalc-1].X)
			} else {
				angle = math.Atan2((*e.points)[pCalc-1].Y-(*e.points)[pCalc].Y, (*e.points)[pCalc-1].X-(*e.points)[pCalc].X)
			}

		default:
			if angleCorrection == true {
				angle = math.Atan2((*e.points)[pCalc].Y-(*e.points)[pCalc-1].Y, (*e.points)[pCalc].X-(*e.points)[pCalc-1].X)
			} else {
				angle = math.Atan2((*e.points)[pCalc-1].Y-(*e.points)[pCalc].Y, (*e.points)[pCalc-1].X-(*e.points)[pCalc].X)
			}
		}

		e.Rotate(angle)
		e.SetXY(int((*e.points)[pCalc].X), int((*e.points)[pCalc].Y))
		e.Data(map[string]string{"angle": strconv.FormatFloat(angle, 'g', 10, 64)})
	}

	return
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
func (e *TagSpan) Text(value any) (ref *TagSpan) {
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
func (e *TagSpan) Html(value string) (ref *TagSpan) {
	e.selfElement.Set("innerHTML", value)
	return e
}

// AddListenerMouseOver #replicar
//
// English:
//
// Adds a mouse over event listener equivalent to the JavaScript command addEventListener('mouseover',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
// Fired when a mouse or other pointing device is moved outside the element.
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de mouse sobre, equivalente ao comando JavaScript addEventListener('mouseover',fn).
//
// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do elemento.
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) AddListenerMouseOver(mouseEvent chan mouse.Data) (ref *TagSpan) {
	if e.fnMouseOver != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseOver, this, args)
		return nil
	})
	e.fnMouseOver = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseover",
		*e.fnMouseOver,
	)
	return e
}

// RemoveListenerMouseOver #replicar #pareiaqui
//
// English:
//
// Removes a mouse over event listener, equivalent to the JavaScript command RemoveEventListener('mouseover',fn).
//
// Fired when a mouse or other pointing device is moved outside the element.
//
// Português:
//
// Remove um ouvinte de evento de mouse sobre, equivalente ao comando JavaScript RemoveEventListener('mouseover',fn).
//
// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do elemento.
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) RemoveListenerMouseOver() (ref *TagSpan) {
	if e.fnMouseOver == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseover",
		*e.fnMouseOver,
	)
	e.fnMouseOver = nil
	return e
}

// AddListenerMouseOut #replicar
//
// English:
//
// Adds a mouse out event listener equivalent to the JavaScript command addEventListener('mouseout',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
// Fired when a mouse or other pointing device is moved outside the boundary of the element.
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de mouse fora, equivalente ao comando JavaScript addEventListener('mouseout',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do limite do elemento.
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) AddListenerMouseOut(mouseEvent chan mouse.Data) (ref *TagSpan) {
	if e.fnMouseOut != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseOut, this, args)
		return nil
	})
	e.fnMouseOut = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseout",
		*e.fnMouseOut,
	)
	return e
}

// RemoveListenerMouseOut #replicar
//
// English:
//
// Removes a mouse out event listener, equivalent to the JavaScript command RemoveEventListener('mouseout',fn).
//
// Fired when a mouse or other pointing device is moved outside the boundary of the element.
//
// Português:
//
// Remove um ouvinte de evento de mouse fora, equivalente ao comando JavaScript RemoveEventListener('mouseout',fn).
//
// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do limite do elemento.
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) RemoveListenerMouseOut() (ref *TagSpan) {
	if e.fnMouseOut == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseout",
		*e.fnMouseOut,
	)
	e.fnMouseOut = nil
	return e
}

// AddListenerMouseMove #replicar
//
// English:
//
// Adds a mouse move event listener equivalent to the JavaScript command addEventListener('mousemove',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
// Fired when a mouse or other pointing device is moved while over an element.
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de mouse move, equivalente ao comando JavaScript addEventListener('mousemove',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
// Acionado quando um mouse ou outro dispositivo apontador é movido sobre um elemento.
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) AddListenerMouseMove(mouseEvent chan mouse.Data) (ref *TagSpan) {
	if e.fnMouseMove != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseMove, this, args)
		return nil
	})
	e.fnMouseMove = &fn

	e.selfElement.Call(
		"addEventListener",
		"mousemove",
		*e.fnMouseMove,
	)
	return e
}

// RemoveListenerMouseMove #replicar
//
// English:
//
// Removes a mouse move event listener, equivalent to the JavaScript command RemoveEventListener('mousemove',fn).
//
// Fired when a mouse or other pointing device is moved while over an element.
//
// Português:
//
// Remove um ouvinte de evento de mouse move, equivalente ao comando JavaScript RemoveEventListener('mousemove',fn).
//
// Acionado quando um mouse ou outro dispositivo apontador é movido sobre um elemento.
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) RemoveListenerMouseMove() (ref *TagSpan) {
	if e.fnMouseMove == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mousemove",
		*e.fnMouseMove,
	)
	e.fnMouseMove = nil
	return e
}

// AddListenerMouseLeave #replicar
//
// English:
//
// Adds a mouse leave event listener equivalent to the JavaScript command addEventListener('mouseleave',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
// Fired when a mouse or other pointing device is moved outside the boundary of the element and all of its descendants.
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de mouse saiu, equivalente ao comando JavaScript addEventListener('mouseleave',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do limite do elemento e de todos os seus descendentes.
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) AddListenerMouseLeave(mouseEvent chan mouse.Data) (ref *TagSpan) {
	if e.fnMouseLeave != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseLeave, this, args)
		return nil
	})
	e.fnMouseLeave = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseleave",
		e.fnMouseLeave,
	)
	return e
}

// RemoveListenerMouseLeave #replicar
//
// English:
//
// Removes a mouse leave event listener, equivalent to the JavaScript command RemoveEventListener('mouseleave',fn).
//
// Fired when a mouse or other pointing device is moved outside the boundary of the element and all of its descendants.
//
// Português:
//
// Remove um ouvinte de evento de mouse saiu, equivalente ao comando JavaScript RemoveEventListener('mouseleave',fn).
//
// Acionado quando um mouse ou outro dispositivo apontador é movido para fora do limite do elemento e de todos os seus descendentes.
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) RemoveListenerMouseLeave() (ref *TagSpan) {
	if e.fnMouseLeave == nil {
		return e
	}

	e.selfElement.Call(
		"addEventListener",
		"mouseleave",
		e.fnMouseLeave,
	)
	e.fnMouseLeave = nil
	return e
}

// AddListenerMouseEnter #replicar
//
// English:
//
// Adds a mouse enter event listener equivalent to the JavaScript command addEventListener('mouseenter',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
// Fired when a mouse or other pointing device is moved inside the boundary of the element or one of its descendants.
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de mouse entrou, equivalente ao comando JavaScript addEventListener('mouseenter',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
// Acionado quando um mouse ou outro dispositivo apontador é movido para dentro do limite do elemento ou de um de seus descendentes.
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) AddListenerMouseEnter(mouseEvent chan mouse.Data) (ref *TagSpan) {
	if e.fnMouseEnter != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseEnter, this, args)
		return nil
	})
	e.fnMouseEnter = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseenter",
		*e.fnMouseEnter,
	)
	return e
}

// RemoveListenerMouseEnter #replicar
//
// English:
//
// Removes a mouse enter event listener, equivalent to the JavaScript command RemoveEventListener('mouseenter',fn).
//
// Fired when a mouse or other pointing device is moved inside the boundary of the element or one of its descendants.
//
// Português:
//
// Remove um ouvinte de evento de mouse entrou, equivalente ao comando JavaScript RemoveEventListener('mouseenter',fn).
//
// Acionado quando um mouse ou outro dispositivo apontador é movido para dentro do limite do elemento ou de um de seus descendentes.
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) RemoveListenerMouseEnter() (ref *TagSpan) {
	if e.fnMouseEnter == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseenter",
		*e.fnMouseEnter,
	)
	e.fnMouseEnter = nil
	return e
}

// AddListenerMouseDown #replicar
//
// English:
//
// Adds a mouse down event listener equivalent to the JavaScript command addEventListener('mousedown',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
// Fired when the user presses a button on a mouse or other pointing device, while the pointer is over the element.
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de botão do mouse precionado, equivalente ao comando JavaScript
// addEventListener('mousedown',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
// Acionado quando o usuário pressiona um botão em um mouse ou outro dispositivo apontador, enquanto o ponteiro está sobre o elemento.
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) AddListenerMouseDown(mouseEvent chan mouse.Data) (ref *TagSpan) {
	if e.fnMouseDown != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseDown, this, args)
		return nil
	})
	e.fnMouseDown = &fn

	e.selfElement.Call(
		"addEventListener",
		"mousedown",
		e.fnMouseDown,
	)
	return e
}

// RemoveListenerMouseDown #replicar
//
// English:
//
// Removes a mouse down event listener, equivalent to the JavaScript command RemoveEventListener('mousedown',fn).
//
// Fired when the user presses a button on a mouse or other pointing device, while the pointer is over the element.
//
// Português:
//
// Remove um ouvinte de evento de botão do mouse precionado, equivalente ao comando JavaScript RemoveEventListener('mousedown',fn).
//
// Acionado quando o usuário pressiona um botão em um mouse ou outro dispositivo apontador, enquanto o ponteiro está sobre o elemento.
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) RemoveListenerMouseDown() (ref *TagSpan) {
	if e.fnMouseDown == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mousedown",
		e.fnMouseDown,
	)
	e.fnMouseDown = nil
	return e
}

// AddListenerMouseUp #replicar
//
// English:
//
// Adds a mouse uo event listener equivalent to the JavaScript command addEventListener('mouseup',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
// Fired when the user releases a button on a mouse or other pointing device, while the pointer is over the element.
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de botão do mouse liberado, equivalente ao comando JavaScript
// addEventListener('mouseup',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
// Acionado quando o usuário libera um botão em um mouse ou outro dispositivo apontador, enquanto o ponteiro está sobre o elemento.
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) AddListenerMouseUp(mouseEvent chan mouse.Data) (ref *TagSpan) {
	if e.fnMouseUp != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseUp, this, args)
		return nil
	})
	e.fnMouseUp = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseup",
		*e.fnMouseUp,
	)
	return e
}

// RemoveListenerMouseUp #replicar
//
// English:
//
// Removes a mouse up event listener, equivalent to the JavaScript command RemoveEventListener('mouseup',fn).
//
// Fired when the user releases a button on a mouse or other pointing device, while the pointer is over the element.
//
// Português:
//
// Remove um ouvinte de evento de botão do mouse liberado, equivalente ao comando JavaScript RemoveEventListener('mouseup',fn).
//
// Acionado quando o usuário libera um botão em um mouse ou outro dispositivo apontador, enquanto o ponteiro está sobre o elemento.
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) RemoveListenerMouseUp() (ref *TagSpan) {
	if e.fnMouseUp == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseup",
		*e.fnMouseUp,
	)
	e.fnMouseUp = nil
	return e
}

// AddListenerMouseWheel #replicar
//
// English:
//
// Adds a mouse wheel event listener equivalent to the JavaScript command addEventListener('mousewheel',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
// Fired when the user rotates a mouse wheel or similar user interface component such as a touchpad.
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de roda do mouse equivalente ao comando JavaScript addEventListener('mousewheel',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
// Acionado quando o usuário gira a roda do mouse ou um componente de interface de usuário semelhante, como um touchpad.
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) AddListenerMouseWheel(mouseEvent chan mouse.Data) (ref *TagSpan) {
	if e.fnMouseWheel != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventMouseWheel, this, args)
		return nil
	})
	e.fnMouseWheel = &fn

	e.selfElement.Call(
		"addEventListener",
		"mousewheel",
		*e.fnMouseWheel,
	)
	return e
}

// RemoveListenerMouseWheel #replicar
//
// English:
//
// Removes a mouse wheel event listener, equivalent to the JavaScript command RemoveEventListener('mousewheel',fn).
//
// Fired when the user rotates a mouse wheel or similar user interface component such as a touchpad.
//
// Português:
//
// Remove um ouvinte de evento de roda do mouse, equivalente ao comando JavaScript RemoveEventListener('mousewheel',fn).
//
// Acionado quando o usuário gira a roda do mouse ou um componente de interface de usuário semelhante, como um touchpad.
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) RemoveListenerMouseWheel() (ref *TagSpan) {
	if e.fnMouseWheel == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"mousewheel",
		*e.fnMouseWheel,
	)
	e.fnMouseWheel = nil
	return e
}

// AddListenerDoubleClick #replicar
//
// English:
//
// Adds a mouse double click event listener equivalent to the JavaScript command addEventListener('dblclick',fn).
//
//	Input:
//	  mouseEvent: pointer to channel mouse.Data
//
// Fired when the user double-clicks the primary pointer button.
//
//	Notes:
//	  * For more information see the website https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
// Português:
//
// Adiciona um ouvinte de evento de click duplo do mouse equivalente ao comando JavaScript
// addEventListener('dblclick',fn).
//
//	Entrada:
//	  mouseEvent: ponteiro para o channel mouse.Data
//
// Acionado quando o usuário clica duas vezes no botão do ponteiro principal.
//
//	Notas:
//	  * Para mais informações veja o site https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) AddListenerDoubleClick(mouseEvent chan mouse.Data) (ref *TagSpan) {
	if e.fnDoubleClick != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		mouseEvent <- mouse.EventManager(mouse.KEventDoubleClick, this, args)
		return nil
	})
	e.fnDoubleClick = &fn

	e.selfElement.Call(
		"addEventListener",
		"dblclick",
		*e.fnDoubleClick,
	)
	return e
}

// RemoveListenerDoubleClick #replicar
//
// English:
//
// Removes a double click event listener, equivalent to the JavaScript command RemoveEventListener('dblclick',fn).
//
// Fired when the user double-clicks the primary pointer button.
//
// Português:
//
// Remove um ouvinte de evento de click duplo, equivalente ao comando JavaScript RemoveEventListener('dblclick',fn).
//
// Acionado quando o usuário clica duas vezes no botão do ponteiro principal.
//
//	Example: / Exemplo:
//	  tagCircle := &html.TagSvgCircle{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().Reference(&tagCircle).AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: Remove the addEventListener('click') from the three elements
//	        // Português: Remove o addEventListener('click') dos três elementos
//	        tagCircle.RemoveListenerClick()
//	      }
//	    }
//	  }()
//
//	Example: / Exemplo:
//	  tagUse := &html.TagSvgUse{}
//	  mouseEvent := make(chan mouse.Data)
//
//	  stage := factoryBrowser.NewStage()
//
//	  s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, 0, 30, 10}).Append(
//	    factoryBrowser.NewTagSvgCircle().AddListenerClick(&mouseEvent).Id("myCircle").Cx(5).Cy(5).R(4).Stroke(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().Reference(&tagUse).HRef("#myCircle").X(10).Fill(factoryColor.NewBlue()),
//	    factoryBrowser.NewTagSvgUse().HRef("#myCircle").X(20).Fill(factoryColor.NewWhite()).Stroke(factoryColor.NewRed()),
//	  )
//
//	  go func() {
//	    for {
//	      select {
//	      case <-mouseEvent:
//	        log.Printf("click")
//	        // English: addEventListener('click') was created on the <circle> element, so the reference is invalid and
//	        //   the command does not work.
//	        // Português: addEventListener('click') foi criado no elemento <circle>, por isto, a refereência é
//	        //   inválida e o comando não funciona.
//	        tagUse.RemoveListenerClick()
//	      }
//	    }
//	  }()
func (e *TagSpan) RemoveListenerDoubleClick() (ref *TagSpan) {
	if e.fnDoubleClick == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"dblclick",
		*e.fnDoubleClick,
	)
	e.fnDoubleClick = nil
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
func (e *TagSpan) GetXY() (x, y int) {
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
func (e *TagSpan) GetX() (x int) {
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
func (e *TagSpan) GetY() (y int) {
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
func (e *TagSpan) GetTop() (top int) {
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
func (e *TagSpan) GetRight() (right int) {
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
func (e *TagSpan) GetBottom() (bottom int) {
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
func (e *TagSpan) GetLeft() (left int) {
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
func (e *TagSpan) GetBoundingBox() (x, y, width, height int) {
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
func (e *TagSpan) CollisionBoundingBox(elemnt CollisionBoundingBox) (collision bool) {
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
func (e *TagSpan) UpdateBoundingClientRect() (ref *TagSpan) {
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
func (e *TagSpan) SetXY(x, y int) (ref *TagSpan) {

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
func (e *TagSpan) SetDeltaX(delta int) (ref *TagSpan) {
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
func (e *TagSpan) SetDeltaY(delta int) (ref *TagSpan) {
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
func (e *TagSpan) SetX(x int) (ref *TagSpan) {

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
func (e *TagSpan) SetY(y int) (ref *TagSpan) {

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

// ListenerAddReflect
//
// English:
//
//	Add event listener
//
//	Events:
//	  cancel: Fired for <input> and <dialog> elements when the user cancels the currently open dialog by closing it with the Esc key.
//	  change: Fired when the value of an <input>, <select>, or <textarea> element has been changed and committed by the user. Unlike the input event, the change event is not necessarily fired for each alteration to an element's value.
//	  error:  Fired when a resource failed to load, or can't be used.
//	  load:   Fires for elements containing a resource when the resource has successfully loaded.
//
//	Clipboard events
//	  copy:   Fired when the user initiates a copy action through the browser's user interface.
//	  cut:    Fired when the user initiates a cut action through the browser's user interface.
//	  paste:  Fired when the user initiates a paste action through the browser's user interface.
//
//	Drag & drop events
//	  drag:       This event is fired when an element or text selection is being dragged.
//	  dragend:    This event is fired when a drag operation is being ended (by releasing a mouse button or hitting the escape key).
//	  dragenter:  This event is fired when a dragged element or text selection enters a valid drop target.
//	  dragleave:  This event is fired when a dragged element or text selection leaves a valid drop target.
//	  dragover:   This event is fired continuously when an element or text selection is being dragged and the mouse pointer is over a valid drop target (every 50 ms WHEN mouse is not moving ELSE much faster between 5 ms (slow movement) and 1ms (fast movement) approximately. This firing pattern is different than mouseover ).
//	  dragstart:  This event is fired when the user starts dragging an element or text selection.
//	  drop:       This event is fired when an element or text selection is dropped on a valid drop target.
//
//	Popover events
//	  beforetoggle: Fired when the element is a popover, before it is hidden or shown.
//	  toggle:       Fired when the element is a popover, just after it is hidden or shown.
func (e *TagSpan) ListenerAddReflect(event string, params []interface{}, functions []reflect.Value, reference any) (ref *TagSpan) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.ListenerAddReflect(event, params, functions, reference)
	return e
}

// ListenerRemove
//
// English:
//
//	Remove event listener
//
//	Events:
//	  cancel: Fired for <input> and <dialog> elements when the user cancels the currently open dialog by closing it with the Esc key.
//	  change: Fired when the value of an <input>, <select>, or <textarea> element has been changed and committed by the user. Unlike the input event, the change event is not necessarily fired for each alteration to an element's value.
//	  error:  Fired when a resource failed to load, or can't be used.
//	  load:   Fires for elements containing a resource when the resource has successfully loaded.
//
//	Clipboard events
//	  copy:   Fired when the user initiates a copy action through the browser's user interface.
//	  cut:    Fired when the user initiates a cut action through the browser's user interface.
//	  paste:  Fired when the user initiates a paste action through the browser's user interface.
//
//	Drag & drop events
//	  drag:       This event is fired when an element or text selection is being dragged.
//	  dragend:    This event is fired when a drag operation is being ended (by releasing a mouse button or hitting the escape key).
//	  dragenter:  This event is fired when a dragged element or text selection enters a valid drop target.
//	  dragleave:  This event is fired when a dragged element or text selection leaves a valid drop target.
//	  dragover:   This event is fired continuously when an element or text selection is being dragged and the mouse pointer is over a valid drop target (every 50 ms WHEN mouse is not moving ELSE much faster between 5 ms (slow movement) and 1ms (fast movement) approximately. This firing pattern is different than mouseover ).
//	  dragstart:  This event is fired when the user starts dragging an element or text selection.
//	  drop:       This event is fired when an element or text selection is dropped on a valid drop target.
//
//	Popover events
//	  beforetoggle: Fired when the element is a popover, before it is hidden or shown.
//	  toggle:       Fired when the element is a popover, just after it is hidden or shown.
func (e *TagSpan) ListenerRemove(event string) (ref *TagSpan) {
	e.commonEvents.ListenerRemove(event)
	return e
}
