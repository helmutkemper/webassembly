package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"log"
	"strconv"
	"strings"
	"syscall/js"
)

type TagOptionGroup struct {
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
func (e *TagOptionGroup) Reference(reference **TagOptionGroup) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) AccessKey(key string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Autofocus(autofocus bool) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Class(class ...string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) ContentEditable(editable bool) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Data(data map[string]string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Dir(dir Dir) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Draggable(draggable Draggable) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Hidden() (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Id(id string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) InputMode(inputMode InputMode) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Is(is string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) ItemId(id string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) ItemProp(itemprop string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) ItemRef(itemref string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) ItemType(itemType string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Lang(language Language) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Part(part ...string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Nonce(nonce string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Slot(slot string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Spellcheck(spell bool) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Style(style string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) TabIndex(index int) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Title(title string) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) Translate(translate Translate) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) CreateElement(tag Tag) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) AppendById(appendId string) (ref *TagOptionGroup) {

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
func (e *TagOptionGroup) Append(append interface{}) (ref *TagOptionGroup) {
	switch append.(type) {
	case *TagOptionGroup:
		e.selfElement.Call("appendChild", append.(*TagOptionGroup).selfElement)
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

// Label
//
// English:
//
//	The name of the group of options, which the browser can use when labeling the options in the user
//	interface. This attribute is mandatory if this element is used.
//
// Português:
//
//	O nome do grupo de opções, que o navegador pode usar ao rotular as opções na interface do usuário.
//	Este atributo é obrigatório se este elemento for usado.
func (e *TagOptionGroup) Label(label string) (ref *TagOptionGroup) {
	e.selfElement.Set("label", label)
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
func (e *TagOptionGroup) Disabled(disabled bool) (ref *TagOptionGroup) {
	e.selfElement.Set("disabled", disabled)
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
func (e *TagOptionGroup) GetXY() (x, y int) {
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
func (e *TagOptionGroup) GetX() (x int) {
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
func (e *TagOptionGroup) GetY() (y int) {
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
func (e *TagOptionGroup) GetTop() (top int) {
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
func (e *TagOptionGroup) GetRight() (right int) {
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
func (e *TagOptionGroup) GetBottom() (bottom int) {
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
func (e *TagOptionGroup) GetLeft() (left int) {
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
func (e *TagOptionGroup) GetBoundingBox() (x, y, width, height int) {
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
func (e *TagOptionGroup) CollisionBoundingBox(elemnt CollisionBoundingBox) (collision bool) {
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
func (e *TagOptionGroup) UpdateBoundingClientRect() (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) SetXY(x, y int) (ref *TagOptionGroup) {

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
func (e *TagOptionGroup) SetDeltaX(delta int) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) SetDeltaY(delta int) (ref *TagOptionGroup) {
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
func (e *TagOptionGroup) SetX(x int) (ref *TagOptionGroup) {

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
func (e *TagOptionGroup) SetY(y int) (ref *TagOptionGroup) {

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

func (e *TagOptionGroup) AddListenerAbort(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerAbort() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagOptionGroup) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerAuxclick() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagOptionGroup) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerBeforeinput() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagOptionGroup) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerBeforematch() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagOptionGroup) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerBeforetoggle() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagOptionGroup) AddListenerCancel(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerCancel() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagOptionGroup) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerCanplay() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagOptionGroup) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerCanplaythrough() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagOptionGroup) AddListenerChange(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerChange() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagOptionGroup) AddListenerClick(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerClick() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagOptionGroup) AddListenerClose(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerClose() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagOptionGroup) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerContextlost() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagOptionGroup) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerContextmenu() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagOptionGroup) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerContextrestored() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagOptionGroup) AddListenerCopy(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerCopy() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagOptionGroup) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerCuechange() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagOptionGroup) AddListenerCut(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerCut() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagOptionGroup) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerDblclick() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagOptionGroup) AddListenerDrag(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerDrag() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagOptionGroup) AddListenerDragend(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerDragend() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagOptionGroup) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerDragenter() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagOptionGroup) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerDragleave() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagOptionGroup) AddListenerDragover(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerDragover() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagOptionGroup) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerDragstart() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagOptionGroup) AddListenerDrop(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerDrop() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagOptionGroup) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerDurationchange() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagOptionGroup) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerEmptied() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagOptionGroup) AddListenerEnded(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerEnded() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagOptionGroup) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerFormdata() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagOptionGroup) AddListenerInput(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerInput() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagOptionGroup) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerInvalid() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagOptionGroup) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerKeydown() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagOptionGroup) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerKeypress() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagOptionGroup) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerKeyup() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagOptionGroup) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerLoadeddata() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagOptionGroup) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerLoadedmetadata() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagOptionGroup) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerLoadstart() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagOptionGroup) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerMousedown() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagOptionGroup) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerMouseenter() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagOptionGroup) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerMouseleave() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagOptionGroup) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerMousemove() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagOptionGroup) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerMouseout() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagOptionGroup) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerMouseover() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagOptionGroup) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerMouseup() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagOptionGroup) AddListenerPaste(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerPaste() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagOptionGroup) AddListenerPause(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerPause() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagOptionGroup) AddListenerPlay(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerPlay() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagOptionGroup) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerPlaying() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagOptionGroup) AddListenerProgress(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerProgress() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagOptionGroup) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerRatechange() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagOptionGroup) AddListenerReset(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerReset() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagOptionGroup) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerScrollend() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagOptionGroup) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerSecuritypolicyviolation() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagOptionGroup) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerSeeked() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagOptionGroup) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerSeeking() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagOptionGroup) AddListenerSelect(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerSelect() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagOptionGroup) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerSlotchange() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagOptionGroup) AddListenerStalled(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerStalled() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagOptionGroup) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerSubmit() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagOptionGroup) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerSuspend() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagOptionGroup) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerTimeupdate() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagOptionGroup) AddListenerToggle(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerToggle() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagOptionGroup) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerVolumechange() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagOptionGroup) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerWaiting() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagOptionGroup) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerWebkitanimationend() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagOptionGroup) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerWebkitanimationiteration() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagOptionGroup) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerWebkitanimationstart() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagOptionGroup) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerWebkittransitionend() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagOptionGroup) AddListenerWheel(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerWheel() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagOptionGroup) AddListenerBlur(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerBlur() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagOptionGroup) AddListenerError(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerError() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagOptionGroup) AddListenerFocus(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerFocus() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagOptionGroup) AddListenerLoad(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerLoad() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagOptionGroup) AddListenerResize(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerResize() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagOptionGroup) AddListenerScroll(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerScroll() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagOptionGroup) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerAfterprint() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagOptionGroup) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerBeforeprint() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagOptionGroup) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerBeforeunload() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagOptionGroup) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerHashchange() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagOptionGroup) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerLanguagechange() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagOptionGroup) AddListenerMessage(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerMessage() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagOptionGroup) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerMessageerror() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagOptionGroup) AddListenerOffline(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerOffline() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagOptionGroup) AddListenerOnline(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerOnline() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagOptionGroup) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerPageswap() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagOptionGroup) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerPagehide() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagOptionGroup) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerPagereveal() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagOptionGroup) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerPageshow() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagOptionGroup) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerPopstate() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagOptionGroup) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerRejectionhandled() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagOptionGroup) AddListenerStorage(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerStorage() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagOptionGroup) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerUnhandledrejection() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagOptionGroup) AddListenerUnload(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerUnload() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagOptionGroup) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerReadystatechange() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagOptionGroup) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagOptionGroup) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagOptionGroup) RemoveListenerVisibilitychange() (ref *TagOptionGroup) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
