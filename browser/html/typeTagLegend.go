package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"log"
	"strconv"
	"strings"
	"syscall/js"
)

type TagLegend struct {
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
func (e *TagLegend) Reference(reference **TagLegend) (ref *TagLegend) {
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
func (e *TagLegend) AccessKey(key string) (ref *TagLegend) {
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
func (e *TagLegend) Autofocus(autofocus bool) (ref *TagLegend) {
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
func (e *TagLegend) Class(class ...string) (ref *TagLegend) {
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
func (e *TagLegend) ContentEditable(editable bool) (ref *TagLegend) {
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
func (e *TagLegend) Data(data map[string]string) (ref *TagLegend) {
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
func (e *TagLegend) Dir(dir Dir) (ref *TagLegend) {
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
func (e *TagLegend) Draggable(draggable Draggable) (ref *TagLegend) {
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
func (e *TagLegend) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagLegend) {
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
func (e *TagLegend) Hidden() (ref *TagLegend) {
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
func (e *TagLegend) Id(id string) (ref *TagLegend) {
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
func (e *TagLegend) InputMode(inputMode InputMode) (ref *TagLegend) {
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
func (e *TagLegend) Is(is string) (ref *TagLegend) {
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
func (e *TagLegend) ItemId(id string) (ref *TagLegend) {
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
func (e *TagLegend) ItemProp(itemprop string) (ref *TagLegend) {
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
func (e *TagLegend) ItemRef(itemref string) (ref *TagLegend) {
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
func (e *TagLegend) ItemType(itemType string) (ref *TagLegend) {
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
func (e *TagLegend) Lang(language Language) (ref *TagLegend) {
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
func (e *TagLegend) Part(part ...string) (ref *TagLegend) {
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
func (e *TagLegend) Nonce(nonce string) (ref *TagLegend) {
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
func (e *TagLegend) Slot(slot string) (ref *TagLegend) {
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
func (e *TagLegend) Spellcheck(spell bool) (ref *TagLegend) {
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
func (e *TagLegend) Style(style string) (ref *TagLegend) {
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
func (e *TagLegend) TabIndex(index int) (ref *TagLegend) {
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
func (e *TagLegend) Title(title string) (ref *TagLegend) {
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
func (e *TagLegend) Translate(translate Translate) (ref *TagLegend) {
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
func (e *TagLegend) CreateElement(tag Tag) (ref *TagLegend) {
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
func (e *TagLegend) AppendById(appendId string) (ref *TagLegend) {

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
func (e *TagLegend) Append(append interface{}) (ref *TagLegend) {
	switch append.(type) {
	case *TagLegend:
		e.selfElement.Call("appendChild", append.(*TagLegend).selfElement)
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

// GetXY #replicar
//
// English:
//
//	Returns the X and Y axes in pixels.
//
// Português:
//
//	Retorna os eixos X e Y em pixels.
func (e *TagLegend) GetXY() (x, y int) {
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
func (e *TagLegend) GetX() (x int) {
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
func (e *TagLegend) GetY() (y int) {
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
func (e *TagLegend) GetTop() (top int) {
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
func (e *TagLegend) GetRight() (right int) {
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
func (e *TagLegend) GetBottom() (bottom int) {
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
func (e *TagLegend) GetLeft() (left int) {
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
func (e *TagLegend) GetBoundingBox() (x, y, width, height int) {
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
func (e *TagLegend) CollisionBoundingBox(elemnt CollisionBoundingBox) (collision bool) {
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
func (e *TagLegend) UpdateBoundingClientRect() (ref *TagLegend) {
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
func (e *TagLegend) SetXY(x, y int) (ref *TagLegend) {

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
func (e *TagLegend) SetDeltaX(delta int) (ref *TagLegend) {
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
func (e *TagLegend) SetDeltaY(delta int) (ref *TagLegend) {
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
func (e *TagLegend) SetX(x int) (ref *TagLegend) {

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
func (e *TagLegend) SetY(y int) (ref *TagLegend) {

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

func (e *TagLegend) AddListenerAbort(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerAbort() (ref *TagLegend) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagLegend) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerAuxclick() (ref *TagLegend) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagLegend) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerBeforeinput() (ref *TagLegend) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagLegend) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerBeforematch() (ref *TagLegend) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagLegend) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerBeforetoggle() (ref *TagLegend) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagLegend) AddListenerCancel(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerCancel() (ref *TagLegend) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagLegend) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerCanplay() (ref *TagLegend) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagLegend) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerCanplaythrough() (ref *TagLegend) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagLegend) AddListenerChange(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerChange() (ref *TagLegend) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagLegend) AddListenerClick(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerClick() (ref *TagLegend) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagLegend) AddListenerClose(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerClose() (ref *TagLegend) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagLegend) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerContextlost() (ref *TagLegend) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagLegend) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerContextmenu() (ref *TagLegend) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagLegend) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerContextrestored() (ref *TagLegend) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagLegend) AddListenerCopy(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerCopy() (ref *TagLegend) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagLegend) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerCuechange() (ref *TagLegend) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagLegend) AddListenerCut(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerCut() (ref *TagLegend) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagLegend) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerDblclick() (ref *TagLegend) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagLegend) AddListenerDrag(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerDrag() (ref *TagLegend) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagLegend) AddListenerDragend(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerDragend() (ref *TagLegend) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagLegend) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerDragenter() (ref *TagLegend) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagLegend) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerDragleave() (ref *TagLegend) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagLegend) AddListenerDragover(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerDragover() (ref *TagLegend) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagLegend) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerDragstart() (ref *TagLegend) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagLegend) AddListenerDrop(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerDrop() (ref *TagLegend) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagLegend) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerDurationchange() (ref *TagLegend) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagLegend) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerEmptied() (ref *TagLegend) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagLegend) AddListenerEnded(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerEnded() (ref *TagLegend) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagLegend) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerFormdata() (ref *TagLegend) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagLegend) AddListenerInput(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerInput() (ref *TagLegend) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagLegend) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerInvalid() (ref *TagLegend) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagLegend) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerKeydown() (ref *TagLegend) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagLegend) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerKeypress() (ref *TagLegend) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagLegend) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerKeyup() (ref *TagLegend) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagLegend) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerLoadeddata() (ref *TagLegend) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagLegend) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerLoadedmetadata() (ref *TagLegend) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagLegend) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerLoadstart() (ref *TagLegend) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagLegend) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerMousedown() (ref *TagLegend) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagLegend) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerMouseenter() (ref *TagLegend) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagLegend) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerMouseleave() (ref *TagLegend) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagLegend) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerMousemove() (ref *TagLegend) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagLegend) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerMouseout() (ref *TagLegend) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagLegend) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerMouseover() (ref *TagLegend) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagLegend) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerMouseup() (ref *TagLegend) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagLegend) AddListenerPaste(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerPaste() (ref *TagLegend) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagLegend) AddListenerPause(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerPause() (ref *TagLegend) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagLegend) AddListenerPlay(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerPlay() (ref *TagLegend) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagLegend) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerPlaying() (ref *TagLegend) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagLegend) AddListenerProgress(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerProgress() (ref *TagLegend) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagLegend) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerRatechange() (ref *TagLegend) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagLegend) AddListenerReset(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerReset() (ref *TagLegend) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagLegend) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerScrollend() (ref *TagLegend) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagLegend) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerSecuritypolicyviolation() (ref *TagLegend) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagLegend) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerSeeked() (ref *TagLegend) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagLegend) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerSeeking() (ref *TagLegend) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagLegend) AddListenerSelect(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerSelect() (ref *TagLegend) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagLegend) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerSlotchange() (ref *TagLegend) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagLegend) AddListenerStalled(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerStalled() (ref *TagLegend) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagLegend) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerSubmit() (ref *TagLegend) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagLegend) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerSuspend() (ref *TagLegend) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagLegend) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerTimeupdate() (ref *TagLegend) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagLegend) AddListenerToggle(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerToggle() (ref *TagLegend) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagLegend) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerVolumechange() (ref *TagLegend) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagLegend) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerWaiting() (ref *TagLegend) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagLegend) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerWebkitanimationend() (ref *TagLegend) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagLegend) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerWebkitanimationiteration() (ref *TagLegend) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagLegend) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerWebkitanimationstart() (ref *TagLegend) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagLegend) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerWebkittransitionend() (ref *TagLegend) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagLegend) AddListenerWheel(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerWheel() (ref *TagLegend) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagLegend) AddListenerBlur(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerBlur() (ref *TagLegend) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagLegend) AddListenerError(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerError() (ref *TagLegend) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagLegend) AddListenerFocus(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerFocus() (ref *TagLegend) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagLegend) AddListenerLoad(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerLoad() (ref *TagLegend) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagLegend) AddListenerResize(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerResize() (ref *TagLegend) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagLegend) AddListenerScroll(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerScroll() (ref *TagLegend) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagLegend) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerAfterprint() (ref *TagLegend) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagLegend) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerBeforeprint() (ref *TagLegend) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagLegend) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerBeforeunload() (ref *TagLegend) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagLegend) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerHashchange() (ref *TagLegend) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagLegend) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerLanguagechange() (ref *TagLegend) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagLegend) AddListenerMessage(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerMessage() (ref *TagLegend) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagLegend) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerMessageerror() (ref *TagLegend) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagLegend) AddListenerOffline(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerOffline() (ref *TagLegend) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagLegend) AddListenerOnline(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerOnline() (ref *TagLegend) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagLegend) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerPageswap() (ref *TagLegend) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagLegend) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerPagehide() (ref *TagLegend) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagLegend) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerPagereveal() (ref *TagLegend) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagLegend) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerPageshow() (ref *TagLegend) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagLegend) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerPopstate() (ref *TagLegend) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagLegend) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerRejectionhandled() (ref *TagLegend) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagLegend) AddListenerStorage(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerStorage() (ref *TagLegend) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagLegend) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerUnhandledrejection() (ref *TagLegend) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagLegend) AddListenerUnload(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerUnload() (ref *TagLegend) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagLegend) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerReadystatechange() (ref *TagLegend) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagLegend) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagLegend) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagLegend) RemoveListenerVisibilitychange() (ref *TagLegend) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
