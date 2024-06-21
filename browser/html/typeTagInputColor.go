package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"log"
	"strconv"
	"strings"
	"syscall/js"
)

type TagInputColor struct {
	commonEvents commonEvents

	tag         Tag
	id          string
	selfElement js.Value
	cssClass    *css.Class

	x          int //#replicar
	y          int //#replicar
	width      int //#replicar
	height     int //#replicar
	heightBBox int //#replicar
	bottom     int //#replicar

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
func (e *TagInputColor) Reference(reference **TagInputColor) (ref *TagInputColor) {
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
func (e *TagInputColor) AccessKey(key string) (ref *TagInputColor) {
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
func (e *TagInputColor) Autofocus(autofocus bool) (ref *TagInputColor) {
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
func (e *TagInputColor) Class(class ...string) (ref *TagInputColor) {
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
func (e *TagInputColor) ContentEditable(editable bool) (ref *TagInputColor) {
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
func (e *TagInputColor) Data(data map[string]string) (ref *TagInputColor) {
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
func (e *TagInputColor) Dir(dir Dir) (ref *TagInputColor) {
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
func (e *TagInputColor) Draggable(draggable Draggable) (ref *TagInputColor) {
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
func (e *TagInputColor) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagInputColor) {
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
func (e *TagInputColor) Hidden() (ref *TagInputColor) {
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
func (e *TagInputColor) Id(id string) (ref *TagInputColor) {
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
func (e *TagInputColor) InputMode(inputMode InputMode) (ref *TagInputColor) {
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
func (e *TagInputColor) Is(is string) (ref *TagInputColor) {
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
func (e *TagInputColor) ItemId(id string) (ref *TagInputColor) {
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
func (e *TagInputColor) ItemProp(itemprop string) (ref *TagInputColor) {
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
func (e *TagInputColor) ItemRef(itemref string) (ref *TagInputColor) {
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
func (e *TagInputColor) ItemType(itemType string) (ref *TagInputColor) {
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
func (e *TagInputColor) Lang(language Language) (ref *TagInputColor) {
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
func (e *TagInputColor) Part(part ...string) (ref *TagInputColor) {
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
func (e *TagInputColor) Nonce(nonce string) (ref *TagInputColor) {
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
func (e *TagInputColor) Slot(slot string) (ref *TagInputColor) {
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
func (e *TagInputColor) Spellcheck(spell bool) (ref *TagInputColor) {
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
func (e *TagInputColor) Style(style string) (ref *TagInputColor) {
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
func (e *TagInputColor) TabIndex(index int) (ref *TagInputColor) {
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
func (e *TagInputColor) Title(title string) (ref *TagInputColor) {
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
func (e *TagInputColor) Translate(translate Translate) (ref *TagInputColor) {
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
func (e *TagInputColor) CreateElement(tag Tag) (ref *TagInputColor) {
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
func (e *TagInputColor) AppendById(appendId string) (ref *TagInputColor) {

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
func (e *TagInputColor) Append(append interface{}) (ref *TagInputColor) {
	switch append.(type) {
	case *TagInputColor:
		e.selfElement.Call("appendChild", append.(*TagInputColor).selfElement)
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
//	The HTML autocomplete attribute lets web developers specify what if any permission the user agent
//	has to provide automated assistance in filling out form field values, as well as guidance to the
//	browser as to the type of information expected in the field.
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
//	Note:
//	  * In order to provide autocompletion, user-agents might require <input>/<select>/<textarea>
//	    elements to:
//	      Have a name and/or id attribute;
//	      Be descendants of a <form> element;
//	      The form to have a submit button.
//
// Português:
//
//	O atributo autocomplete HTML permite que os desenvolvedores da Web especifiquem se existe alguma
//	permissão que o agente do usuário tenha para fornecer assistência automatizada no preenchimento
//	dos valores dos campos do formulário, bem como orientação ao navegador quanto ao tipo de
//	informação esperado no campo.
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
//	Nota:
//	  * Para fornecer preenchimento automático, os agentes do usuário podem exigir elementos
//	    <input> / <select> / <textarea> para:
//	      Ter um atributo name e ou id;
//	      Ser descendentes de um elemento <form>;
//	      O formulário para ter um botão de envio.
func (e *TagInputColor) Autocomplete(autocomplete Autocomplete) (ref *TagInputColor) {
	e.selfElement.Set("autocomplete", autocomplete.String())
	return e
}

// Disabled
//
// English:
//
//	Este atributo booleano impede que o usuário interaja com o elemento.
//
// Português:
//
//	Este atributo booleano impede que o usuário interaja com o elemento.
func (e *TagInputColor) Disabled(disabled bool) (ref *TagInputColor) {
	if disabled {
		e.selfElement.Set("disabled", "disabled")
	}
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
func (e *TagInputColor) Form(form string) (ref *TagInputColor) {
	e.selfElement.Set("form", form)
	return e
}

// List
//
// English:
//
//	The value given to the list attribute should be the id of a <datalist> element located in the same
//	document.
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
//	O valor dado ao atributo list deve ser o id de um elemento <datalist> localizado no mesmo
//	documento.
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
func (e *TagInputColor) List(list string) (ref *TagInputColor) {
	e.selfElement.Set("list", list)
	return e
}

// Max
//
// English:
//
//	Valid for date, month, week, time, datetime-local, number, and range, it defines the greatest
//	value in the range of permitted values.
//	If the value entered into the element exceeds this, the element fails constraint validation.
//	If the value of the max attribute isn't a number, then the element has no maximum value.
//
// There is a special case: if the data type is periodic (such as for dates or times), the value of
// max may be lower than the value of min, which indicates that the range may wrap around;
// for example, this allows you to specify a time range from 10 PM to 4 AM.
//
// Português:
//
//	Válido para data, mês, semana, hora, datetime-local, número e intervalo, define o maior valor no
//	intervalo de valores permitidos. Se o valor inserido no elemento exceder isso, o elemento falhará
//	na validação de restrição. Se o valor do atributo max não for um número, o elemento não terá
//	valor máximo.
//
// Há um caso especial: se o tipo de dado for periódico (como para datas ou horas), o valor de max
// pode ser menor que o valor de min, o que indica que o intervalo pode ser contornado;
// por exemplo, isso permite que você especifique um intervalo de tempo das 22h às 4h.
func (e *TagInputColor) Max(max int) (ref *TagInputColor) {
	e.selfElement.Set("max", max)
	return e
}

// Min
//
// English:
//
//	Valid for date, month, week, time, datetime-local, number, and range, it defines the most negative
//	value in the range of permitted values.
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
//	Válido para data, mês, semana, hora, data e hora local, número e intervalo, define o valor mais
//	negativo no intervalo de valores permitidos.
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
func (e *TagInputColor) Min(min int) (ref *TagInputColor) {
	e.selfElement.Set("max", min)
	return e
}

// MinLength
//
// English:
//
//	Valid for text, search, url, tel, email, and password, it defines the minimum number of
//	characters (as UTF-16 code units) the user can enter into the entry field.
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
//	Válido para texto, pesquisa, url, tel, email e senha, define o número mínimo de caracteres
//	(como unidades de código UTF-16) que o usuário pode inserir no campo de entrada.
//
// Este deve ser um valor inteiro não negativo menor ou igual ao valor especificado por maxlength.
// Se nenhum comprimento mínimo for especificado ou um valor inválido for especificado, a entrada não
// terá comprimento mínimo.
//
// A entrada falhará na validação de restrição se o comprimento do texto inserido no campo for
// inferior a unidades de código UTF-16 de comprimento mínimo, impedindo o envio do formulário.
func (e *TagInputColor) MinLength(minlength int) (ref *TagInputColor) {
	e.selfElement.Set("minlength", minlength)
	return e
}

// Name
//
// English:
//
//	The name of the button, submitted as a pair with the button's value as part of the form data,
//	when that button is used to submit the form.
//
// Português:
//
//	O nome do botão, enviado como um par com o valor do botão como parte dos dados do formulário,
//	quando esse botão é usado para enviar o formulário.
func (e *TagInputColor) Name(name string) (ref *TagInputColor) {
	e.selfElement.Set("name", name)
	return e
}

// ReadOnly
//
// English:
//
//	A Boolean attribute which, if present, indicates that the user should not be able to edit the
//	value of the input.
//
// The readonly attribute is supported by the text, search, url, tel, email, date, month, week, time,
// datetime-local, number, and password input types.
//
// Português:
//
//	A Boolean attribute which, if present, indicates that the user should not be able to edit the value of the input. The readonly attribute is supported by the text, search, url, tel, email, date, month, week, time, datetime-local, number, and password input types.
func (e *TagInputColor) ReadOnly(readonly bool) (ref *TagInputColor) {
	e.selfElement.Set("readonly", readonly)
	return e
}

// Required
//
// English:
//
//	A Boolean attribute indicating that an option with a non-empty string value must be selected.
//
// Português:
//
//	Um atributo booleano que indica que uma opção com um valor de string não vazio deve ser
//	selecionada.
func (e *TagInputColor) Required(required bool) (ref *TagInputColor) {
	e.selfElement.Set("required", required)
	return e
}

// Step
//
// English:
//
//	Valid for the numeric input types, including number, date/time input types, and range, the step
//	attribute is a number that specifies the granularity that the value must adhere to.
//
//	 If not explicitly included:
//	   * step defaults to 1 for number and range;
//	   * For the date/time input types, step is expressed in seconds, with the default step being 60
//	     seconds. The step scale factor is 1000 (which converts the seconds to milliseconds, as used
//	     in other algorithms);
//	   * The value must be a positive number—integer or float—or the special value any, which means
//	     no stepping is implied, and any value is allowed (barring other constraints, such as min and
//	     max).
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
//	Note:
//	  * When the data entered by the user doesn't adhere to the stepping configuration, the value is
//	    considered invalid in constraint validation and will match the :invalid pseudoclass.
//
// Português:
//
//	Válido para os tipos de entrada numérica, incluindo número, tipos de entrada de data e hora e
//	intervalo, o atributo step é um número que especifica a granularidade à qual o valor deve aderir.
//
//	 Se não estiver explicitamente incluído:
//	   * step padroniza para 1 para número e intervalo.
//	   * Para os tipos de entrada de data e hora, a etapa é expressa em segundos, com a etapa padrão
//	     sendo 60 segundos. O fator de escala de passo é 1000 (que converte os segundos em
//	     milissegundos, conforme usado em outros algoritmos).
//	   * O valor deve ser um número positivo — inteiro ou flutuante — ou o valor especial any, o que
//	     significa que nenhuma depuração está implícita e qualquer valor é permitido (exceto outras
//	     restrições, como min e max).
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
//	Nota:
//	  * Quando os dados inseridos pelo usuário não estão de acordo com a configuração de stepping,
//	    o valor é considerado inválido na validação da restrição e corresponderá à
//	    :invalid pseudoclass.
func (e *TagInputColor) Step(step int) (ref *TagInputColor) {
	e.selfElement.Set("step", step)
	return e
}

// Type
//
// English:
//
//	How an <input> works varies considerably depending on the value of its type attribute, hence the
//	different types are covered in their own separate reference pages.
//
// If this attribute is not specified, the default type adopted is text.
//
// Português:
//
//	Como um <input> funciona varia consideravelmente dependendo do valor de seu atributo type,
//	portanto, os diferentes tipos são abordados em suas próprias páginas de referência separadas.
//
// Se este atributo não for especificado, o tipo padrão adotado é texto.
func (e *TagInputColor) Type(inputType InputType) (ref *TagInputColor) {
	e.selfElement.Set("type", inputType.String())
	return e
}

// Value
//
// English:
//
//	Defines the value associated with the element.
//
// Português:
//
//	Define o valor associado ao elemento.
func (e *TagInputColor) Value(value string) (ref *TagInputColor) {
	e.selfElement.Set("value", value)
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
func (e *TagInputColor) GetXY() (x, y int) {
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
func (e *TagInputColor) GetX() (x int) {
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
func (e *TagInputColor) GetY() (y int) {
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
func (e *TagInputColor) GetTop() (top int) {
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
func (e *TagInputColor) GetRight() (right int) {
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
func (e *TagInputColor) GetBottom() (bottom int) {
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
func (e *TagInputColor) GetLeft() (left int) {
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
func (e *TagInputColor) GetBoundingBox() (x, y, width, height int) {
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
func (e *TagInputColor) CollisionBoundingBox(elemnt CollisionBoundingBox) (collision bool) {
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
func (e *TagInputColor) UpdateBoundingClientRect() (ref *TagInputColor) {
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
func (e *TagInputColor) SetXY(x, y int) (ref *TagInputColor) {

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
func (e *TagInputColor) SetDeltaX(delta int) (ref *TagInputColor) {
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
func (e *TagInputColor) SetDeltaY(delta int) (ref *TagInputColor) {
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
func (e *TagInputColor) SetX(x int) (ref *TagInputColor) {

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
func (e *TagInputColor) SetY(y int) (ref *TagInputColor) {

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

func (e *TagInputColor) Get() (el js.Value) {
	return e.selfElement
}

func (e *TagInputColor) AddListenerAbort(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerAbort() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagInputColor) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerAuxclick() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagInputColor) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerBeforeinput() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagInputColor) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerBeforematch() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagInputColor) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerBeforetoggle() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagInputColor) AddListenerCancel(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerCancel() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagInputColor) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerCanplay() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagInputColor) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerCanplaythrough() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagInputColor) AddListenerChange(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerChange() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagInputColor) AddListenerClick(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerClick() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagInputColor) AddListenerClose(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerClose() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagInputColor) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerContextlost() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagInputColor) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerContextmenu() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagInputColor) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerContextrestored() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagInputColor) AddListenerCopy(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerCopy() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagInputColor) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerCuechange() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagInputColor) AddListenerCut(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerCut() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagInputColor) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerDblclick() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagInputColor) AddListenerDrag(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerDrag() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagInputColor) AddListenerDragend(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerDragend() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagInputColor) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerDragenter() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagInputColor) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerDragleave() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagInputColor) AddListenerDragover(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerDragover() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagInputColor) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerDragstart() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagInputColor) AddListenerDrop(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerDrop() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagInputColor) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerDurationchange() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagInputColor) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerEmptied() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagInputColor) AddListenerEnded(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerEnded() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagInputColor) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerFormdata() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagInputColor) AddListenerInput(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerInput() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagInputColor) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerInvalid() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagInputColor) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerKeydown() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagInputColor) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerKeypress() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagInputColor) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerKeyup() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagInputColor) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerLoadeddata() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagInputColor) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerLoadedmetadata() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagInputColor) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerLoadstart() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagInputColor) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerMousedown() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagInputColor) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerMouseenter() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagInputColor) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerMouseleave() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagInputColor) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerMousemove() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagInputColor) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerMouseout() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagInputColor) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerMouseover() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagInputColor) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerMouseup() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagInputColor) AddListenerPaste(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerPaste() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagInputColor) AddListenerPause(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerPause() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagInputColor) AddListenerPlay(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerPlay() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagInputColor) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerPlaying() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagInputColor) AddListenerProgress(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerProgress() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagInputColor) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerRatechange() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagInputColor) AddListenerReset(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerReset() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagInputColor) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerScrollend() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagInputColor) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerSecuritypolicyviolation() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagInputColor) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerSeeked() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagInputColor) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerSeeking() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagInputColor) AddListenerSelect(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerSelect() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagInputColor) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerSlotchange() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagInputColor) AddListenerStalled(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerStalled() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagInputColor) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerSubmit() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagInputColor) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerSuspend() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagInputColor) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerTimeupdate() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagInputColor) AddListenerToggle(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerToggle() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagInputColor) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerVolumechange() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagInputColor) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerWaiting() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagInputColor) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerWebkitanimationend() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagInputColor) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerWebkitanimationiteration() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagInputColor) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerWebkitanimationstart() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagInputColor) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerWebkittransitionend() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagInputColor) AddListenerWheel(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerWheel() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagInputColor) AddListenerBlur(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerBlur() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagInputColor) AddListenerError(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerError() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagInputColor) AddListenerFocus(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerFocus() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagInputColor) AddListenerLoad(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerLoad() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagInputColor) AddListenerResize(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerResize() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagInputColor) AddListenerScroll(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerScroll() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagInputColor) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerAfterprint() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagInputColor) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerBeforeprint() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagInputColor) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerBeforeunload() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagInputColor) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerHashchange() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagInputColor) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerLanguagechange() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagInputColor) AddListenerMessage(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerMessage() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagInputColor) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerMessageerror() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagInputColor) AddListenerOffline(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerOffline() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagInputColor) AddListenerOnline(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerOnline() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagInputColor) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerPageswap() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagInputColor) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerPagehide() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagInputColor) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerPagereveal() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagInputColor) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerPageshow() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagInputColor) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerPopstate() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagInputColor) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerRejectionhandled() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagInputColor) AddListenerStorage(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerStorage() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagInputColor) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerUnhandledrejection() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagInputColor) AddListenerUnload(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerUnload() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagInputColor) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerReadystatechange() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagInputColor) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagInputColor) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagInputColor) RemoveListenerVisibilitychange() (ref *TagInputColor) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
