package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"log"
	"strconv"
	"strings"
	"syscall/js"
)

type TagInputUrl struct {
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
func (e *TagInputUrl) Reference(reference **TagInputUrl) (ref *TagInputUrl) {
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
func (e *TagInputUrl) AccessKey(key string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Autofocus(autofocus bool) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Class(class ...string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) ContentEditable(editable bool) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Data(data map[string]string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Dir(dir Dir) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Draggable(draggable Draggable) (ref *TagInputUrl) {
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
func (e *TagInputUrl) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Hidden() (ref *TagInputUrl) {
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
func (e *TagInputUrl) Id(id string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) InputMode(inputMode InputMode) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Is(is string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) ItemId(id string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) ItemProp(itemprop string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) ItemRef(itemref string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) ItemType(itemType string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Lang(language Language) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Part(part ...string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Nonce(nonce string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Slot(slot string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Spellcheck(spell bool) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Style(style string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) TabIndex(index int) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Title(title string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Translate(translate Translate) (ref *TagInputUrl) {
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
func (e *TagInputUrl) CreateElement(tag Tag) (ref *TagInputUrl) {
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
func (e *TagInputUrl) AppendById(appendId string) (ref *TagInputUrl) {

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
func (e *TagInputUrl) Append(append interface{}) (ref *TagInputUrl) {
	switch append.(type) {
	case *TagInputUrl:
		e.selfElement.Call("appendChild", append.(*TagInputUrl).selfElement)
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
func (e *TagInputUrl) Autocomplete(autocomplete Autocomplete) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Disabled(disabled bool) (ref *TagInputUrl) {
	e.selfElement.Set("disabled", disabled)
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
func (e *TagInputUrl) Form(form string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) List(list string) (ref *TagInputUrl) {
	e.selfElement.Set("list", list)
	return e
}

// MaxLength
//
// English:
//
//	Valid for text, search, url, tel, email, and password, it defines the maximum number of characters
//	(as UTF-16 code units) the user can enter into the field. This must be an integer value 0 or
//	higher. If no maxlength is specified, or an invalid value is specified, the field has no maximum
//	length. This value must also be greater than or equal to the value of minlength.
//
// The input will fail constraint validation if the length of the text entered into the field is
// greater than maxlength UTF-16 code units long. By default, browsers prevent users from entering
// more characters than allowed by the maxlength attribute.
//
// Português:
//
//	Válido para texto, pesquisa, url, tel, email e senha, define o número máximo de caracteres
//	(como unidades de código UTF-16) que o usuário pode inserir no campo.
//
// Este deve ser um valor inteiro 0 ou superior. Se nenhum comprimento máximo for especificado ou um
// valor inválido for especificado, o campo não terá comprimento máximo. Esse valor também deve ser
// maior ou igual ao valor de minlength.
//
// A entrada falhará na validação de restrição se o comprimento do texto inserido no campo for maior
// que o comprimento máximo das unidades de código UTF-16. Por padrão, os navegadores impedem que os
// usuários insiram mais caracteres do que o permitido pelo atributo maxlength.
func (e *TagInputUrl) MaxLength(maxlength int) (ref *TagInputUrl) {
	e.selfElement.Set("maxlength", maxlength)
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
func (e *TagInputUrl) MinLength(minlength int) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Name(name string) (ref *TagInputUrl) {
	e.selfElement.Set("name", name)
	return e
}

// Placeholder
//
// English:
//
//	The placeholder attribute is a string that provides a brief hint to the user as to what kind of
//	information is expected in the field. It should be a word or short phrase that provides a hint
//	as to the expected type of data, rather than an explanation or prompt. The text must not include
//	carriage returns or line feeds. So for example if a field is expected to capture a user's first
//	name, and its label is "First Name", a suitable placeholder might be "e.g. Mustafa".
//
//	 Note:
//	   * The placeholder attribute is not as semantically useful as other ways to explain your form,
//	     and can cause unexpected technical issues with your content. See Labels for more information.
//
// Português:
//
//	O atributo placeholder é uma string que fornece uma breve dica ao usuário sobre que tipo de
//	informação é esperada no campo. Deve ser uma palavra ou frase curta que forneça uma dica sobre o
//	tipo de dados esperado, em vez de uma explicação ou prompt. O texto não deve incluir retornos de
//	carro ou feeds de linha. Assim, por exemplo, se espera-se que um campo capture o primeiro nome de
//	um usuário e seu rótulo for "Nome", um espaço reservado adequado pode ser "por exemplo, Mustafa".
//
//	 Nota:
//	   * O atributo placeholder não é tão semanticamente útil quanto outras formas de explicar seu
//	     formulário e pode causar problemas técnicos inesperados com seu conteúdo. Consulte Rótulos
//	     para obter mais informações.
func (e *TagInputUrl) Placeholder(placeholder string) (ref *TagInputUrl) {
	e.selfElement.Set("placeholder", placeholder)
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
func (e *TagInputUrl) ReadOnly(readonly bool) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Required(required bool) (ref *TagInputUrl) {
	e.selfElement.Set("required", required)
	return e
}

// Size
//
// English:
//
//	If the control is presented as a scrolling list box (e.g. when multiple is specified), this
//	attribute represents the number of rows in the list that should be visible at one time.
//	Browsers are not required to present a select element as a scrolled list box. The default value
//	is 0.
//
//	 Note:
//	   * According to the HTML5 specification, the default value for size should be 1; however, in
//	     practice, this has been found to break some web sites, and no other browser currently does
//	     that, so Mozilla has opted to continue to return 0 for the time being with Firefox.
//
// Português:
//
//	Se o controle for apresentado como uma caixa de listagem de rolagem (por exemplo, quando múltiplo
//	é especificado), esse atributo representa o número de linhas na lista que devem estar visíveis ao
//	mesmo tempo. Os navegadores não precisam apresentar um elemento de seleção como uma caixa de
//	listagem rolada. O valor padrão é 0.
//
//	 Nota:
//	   * De acordo com a especificação HTML5, o valor padrão para tamanho deve ser 1; no entanto, na
//	     prática, descobriu-se que isso quebra alguns sites, e nenhum outro navegador atualmente faz
//	     isso, então a Mozilla optou por continuar retornando 0 por enquanto com o Firefox.
func (e *TagInputUrl) Size(size int) (ref *TagInputUrl) {
	e.selfElement.Set("size", size)
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
func (e *TagInputUrl) Type(inputType InputType) (ref *TagInputUrl) {
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
func (e *TagInputUrl) Value(value string) (ref *TagInputUrl) {
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
func (e *TagInputUrl) GetXY() (x, y int) {
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
func (e *TagInputUrl) GetX() (x int) {
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
func (e *TagInputUrl) GetY() (y int) {
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
func (e *TagInputUrl) GetTop() (top int) {
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
func (e *TagInputUrl) GetRight() (right int) {
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
func (e *TagInputUrl) GetBottom() (bottom int) {
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
func (e *TagInputUrl) GetLeft() (left int) {
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
func (e *TagInputUrl) GetBoundingBox() (x, y, width, height int) {
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
func (e *TagInputUrl) CollisionBoundingBox(elemnt CollisionBoundingBox) (collision bool) {
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
func (e *TagInputUrl) UpdateBoundingClientRect() (ref *TagInputUrl) {
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
func (e *TagInputUrl) SetXY(x, y int) (ref *TagInputUrl) {

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
func (e *TagInputUrl) SetDeltaX(delta int) (ref *TagInputUrl) {
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
func (e *TagInputUrl) SetDeltaY(delta int) (ref *TagInputUrl) {
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
func (e *TagInputUrl) SetX(x int) (ref *TagInputUrl) {

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
func (e *TagInputUrl) SetY(y int) (ref *TagInputUrl) {

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

func (e *TagInputUrl) Get() (el js.Value) {
	return e.selfElement
}

func (e *TagInputUrl) AddListenerAbort(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerAbort() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagInputUrl) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerAuxclick() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagInputUrl) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerBeforeinput() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagInputUrl) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerBeforematch() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagInputUrl) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerBeforetoggle() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagInputUrl) AddListenerCancel(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerCancel() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagInputUrl) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerCanplay() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagInputUrl) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerCanplaythrough() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagInputUrl) AddListenerChange(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerChange() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagInputUrl) AddListenerClick(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerClick() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagInputUrl) AddListenerClose(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerClose() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagInputUrl) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerContextlost() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagInputUrl) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerContextmenu() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagInputUrl) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerContextrestored() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagInputUrl) AddListenerCopy(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerCopy() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagInputUrl) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerCuechange() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagInputUrl) AddListenerCut(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerCut() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagInputUrl) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerDblclick() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagInputUrl) AddListenerDrag(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerDrag() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagInputUrl) AddListenerDragend(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerDragend() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagInputUrl) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerDragenter() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagInputUrl) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerDragleave() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagInputUrl) AddListenerDragover(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerDragover() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagInputUrl) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerDragstart() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagInputUrl) AddListenerDrop(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerDrop() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagInputUrl) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerDurationchange() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagInputUrl) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerEmptied() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagInputUrl) AddListenerEnded(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerEnded() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagInputUrl) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerFormdata() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagInputUrl) AddListenerInput(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerInput() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagInputUrl) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerInvalid() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagInputUrl) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerKeydown() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagInputUrl) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerKeypress() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagInputUrl) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerKeyup() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagInputUrl) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerLoadeddata() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagInputUrl) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerLoadedmetadata() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagInputUrl) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerLoadstart() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagInputUrl) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerMousedown() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagInputUrl) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerMouseenter() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagInputUrl) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerMouseleave() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagInputUrl) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerMousemove() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagInputUrl) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerMouseout() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagInputUrl) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerMouseover() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagInputUrl) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerMouseup() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagInputUrl) AddListenerPaste(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerPaste() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagInputUrl) AddListenerPause(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerPause() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagInputUrl) AddListenerPlay(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerPlay() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagInputUrl) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerPlaying() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagInputUrl) AddListenerProgress(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerProgress() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagInputUrl) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerRatechange() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagInputUrl) AddListenerReset(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerReset() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagInputUrl) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerScrollend() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagInputUrl) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerSecuritypolicyviolation() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagInputUrl) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerSeeked() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagInputUrl) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerSeeking() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagInputUrl) AddListenerSelect(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerSelect() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagInputUrl) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerSlotchange() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagInputUrl) AddListenerStalled(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerStalled() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagInputUrl) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerSubmit() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagInputUrl) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerSuspend() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagInputUrl) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerTimeupdate() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagInputUrl) AddListenerToggle(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerToggle() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagInputUrl) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerVolumechange() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagInputUrl) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerWaiting() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagInputUrl) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerWebkitanimationend() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagInputUrl) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerWebkitanimationiteration() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagInputUrl) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerWebkitanimationstart() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagInputUrl) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerWebkittransitionend() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagInputUrl) AddListenerWheel(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerWheel() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagInputUrl) AddListenerBlur(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerBlur() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagInputUrl) AddListenerError(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerError() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagInputUrl) AddListenerFocus(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerFocus() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagInputUrl) AddListenerLoad(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerLoad() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagInputUrl) AddListenerResize(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerResize() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagInputUrl) AddListenerScroll(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerScroll() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagInputUrl) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerAfterprint() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagInputUrl) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerBeforeprint() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagInputUrl) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerBeforeunload() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagInputUrl) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerHashchange() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagInputUrl) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerLanguagechange() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagInputUrl) AddListenerMessage(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerMessage() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagInputUrl) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerMessageerror() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagInputUrl) AddListenerOffline(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerOffline() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagInputUrl) AddListenerOnline(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerOnline() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagInputUrl) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerPageswap() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagInputUrl) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerPagehide() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagInputUrl) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerPagereveal() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagInputUrl) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerPageshow() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagInputUrl) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerPopstate() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagInputUrl) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerRejectionhandled() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagInputUrl) AddListenerStorage(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerStorage() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagInputUrl) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerUnhandledrejection() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagInputUrl) AddListenerUnload(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerUnload() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagInputUrl) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerReadystatechange() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagInputUrl) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagInputUrl) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagInputUrl) RemoveListenerVisibilitychange() (ref *TagInputUrl) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
