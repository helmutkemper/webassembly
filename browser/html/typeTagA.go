package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"log"
	"strconv"
	"strings"
	"syscall/js"
)

// TagA
//
// English:
//
//	The Anchor element.
//
// The <a> HTML element (or anchor element), with its href attribute, creates a hyperlink to web
// pages, files, email addresses, locations in the same page, or anything else a URL can address.
//
// Content within each <a> should indicate the link's destination. If the href attribute is present,
// pressing the enter key while focused on the <a> element will activate it.
//
// Português:
//
//	O elemento Âncora.
//
// O elemento HTML <a> (ou elemento âncora), com seu atributo href, cria um hiperlink para páginas
// da web, arquivos, endereços de e-mail, locais na mesma página ou qualquer outra coisa que um URL
// possa endereçar.
//
// O conteúdo de cada <a> deve indicar o destino do link. Se o atributo href estiver presente,
// pressionar a tecla enter enquanto estiver focado no elemento <a> irá ativá-lo.
type TagA struct {
	commonEvents commonEvents

	tag         Tag
	id          string
	selfElement js.Value
	cssClass    *css.Class

	x          int
	y          int
	width      int
	height     int
	heightBBox int
	bottom     int

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
func (e *TagA) Reference(reference **TagA) (ref *TagA) {
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
func (e *TagA) AccessKey(key string) (ref *TagA) {
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
func (e *TagA) Autofocus(autofocus bool) (ref *TagA) {
	switch e.tag {
	case KTagButton:
	case KTagSelect:
	default:
		log.Printf("tag " + e.tag.String() + " does not support autofocus property")
	}

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
func (e *TagA) Class(class ...string) (ref *TagA) {
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
func (e *TagA) ContentEditable(editable bool) (ref *TagA) {
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
func (e *TagA) Data(data map[string]string) (ref *TagA) {
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
func (e *TagA) Dir(dir Dir) (ref *TagA) {
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
func (e *TagA) Draggable(draggable Draggable) (ref *TagA) {
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
//	       editor (IME) will be closed.
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
//	The enterKeyHint property is an enumerated property defining what action label (or icon) to
//	present for the enter key on virtual keyboards. It reflects the enterkeyhint HTML global attribute
//	and is an enumerated property, only accepting the following values as a DOMString:
//
//	 Input:
//	   enterKeyHint: defining what action label (or icon) to present for the enter key on virtual
//	     keyboards
//	     KEnterKeyHintEnter: typically indicating inserting a new line.
//	     KEnterKeyHintDone: typically meaning there is nothing more to input and the input method
//	       editor (IME) will be closed.
//	     KEnterKeyHintGo: typically meaning to take the user to the target of the text they typed.
//	     KEnterKeyHintNext: typically taking the user to the next field that will accept text.
//	     KEnterKeyHintPrevious: typically taking the user to the previous field that will accept text.
//	     KEnterKeyHintSearch: typically taking the user to the results of searching for the text they
//	       have typed.
//	     KEnterKeyHintSend: typically delivering the text to its target.
//
// If no enterKeyHint value has been specified or if it was set to a different value than the allowed
// ones, it will return an empty string.
func (e *TagA) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagA) {
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
func (e *TagA) Hidden() (ref *TagA) {
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
func (e *TagA) Id(id string) (ref *TagA) {
	e.id = id
	e.selfElement.Set("id", id)
	return e
}

// InputMode
//
// English:
//
//	The inputmode global attribute is an enumerated attribute that hints at the type of data that
//	might be entered by the user while editing the element or its contents.
//
// This allows a browser to display an appropriate virtual keyboard.
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
func (e *TagA) InputMode(inputMode InputMode) (ref *TagA) {
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
func (e *TagA) Is(is string) (ref *TagA) {
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
func (e *TagA) ItemId(id string) (ref *TagA) {
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
func (e *TagA) ItemProp(itemprop string) (ref *TagA) {
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
func (e *TagA) ItemRef(itemref string) (ref *TagA) {
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
func (e *TagA) ItemType(itemType string) (ref *TagA) {
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
func (e *TagA) Lang(language Language) (ref *TagA) {
	e.selfElement.Set("lang", language.String())
	return e
}

// Part
//
// English:
//
//	A space-separated list of the part names of the element. Part names allows CSS to select and
//	style specific elements in a shadow tree via the ::part pseudo-element.
//
// Português:
//
//	Uma lista separada por espaços dos nomes das partes do elemento. Os nomes das partes permitem
//	que o CSS selecione e estilize elementos específicos em uma árvore de sombra por meio do
//	pseudo-elemento ::part.
func (e *TagA) Part(part ...string) (ref *TagA) {
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
func (e *TagA) Nonce(nonce string) (ref *TagA) {
	e.selfElement.Set("nonce", nonce)
	return e
}

// Slot
//
// English:
//
//	Assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is
//	assigned to the slot created by the <slot> element whose name attribute's value matches that
//	slot attribute's value.
//
// Português:
//
//	Atribui um slot em uma shadow DOM shadow tree a um elemento: Um elemento com um atributo slot é
//	atribuído ao slot criado pelo elemento <slot> cujo valor do atributo name corresponde ao valor
//	desse atributo slot.
func (e *TagA) Slot(slot string) (ref *TagA) {
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
func (e *TagA) Spellcheck(spell bool) (ref *TagA) {
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
func (e *TagA) Style(style string) (ref *TagA) {
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
func (e *TagA) TabIndex(index int) (ref *TagA) {
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
func (e *TagA) Title(title string) (ref *TagA) {
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
func (e *TagA) Translate(translate Translate) (ref *TagA) {
	e.selfElement.Set("translate", translate.String())
	return e
}

// Download
//
// English:
//
//	Causes the browser to treat the linked URL as a download. Can be used with or without a value
//
//	 Note:
//	   * Without a value, the browser will suggest a filename/extension, generated from various
//	     sources:
//	       The Content-Disposition HTTP header;
//	       The final segment in the URL path;
//	       The media type (from the Content-Type header, the start of a data: URL, or Blob.type for a
//	       blob: URL).
//	   * Defining a value suggests it as the filename. / and \ characters are converted to
//	     underscores (_). Filesystems may forbid other characters in filenames, so browsers will
//	     adjust the suggested name if necessary;
//	   * Download only works for same-origin URLs, or the blob: and data: schemes;
//	   * How browsers treat downloads varies by browser, user settings, and other factors. The user
//	     may be prompted before a download starts, or the file may be saved automatically, or it may
//	     open automatically, either in an external application or in the browser itself;
//	   * If the Content-Disposition header has different information from the download attribute,
//	     resulting behavior may differ:
//	       * If the header specifies a filename, it takes priority over a filename specified in the
//	         download attribute;
//	       * If the header specifies a disposition of inline, Chrome and Firefox prioritize the
//	         attribute and treat it as a download. Old Firefox versions (before 82) prioritize the
//	         header and will display the content inline.
//
// Português:
//
//	Faz com que o navegador trate a URL vinculada como um download. Pode ser usado com ou sem valor
//
//	 Nota:
//	   * Sem um valor, o navegador sugerirá uma extensão de nome de arquivo, gerada a partir de várias
//	     fontes:
//	       O cabeçalho HTTP Content-Disposition;
//	       O segmento final no caminho do URL;
//	       O tipo de mídia (do cabeçalho Content-Type, o início de um data: URL ou Blob.type para um
//	       blob: URL).
//	   * Definir um valor sugere-o como o nome do arquivo. / e \ caracteres são convertidos em
//	     sublinhados (_). Os sistemas de arquivos podem proibir outros caracteres em nomes de
//	     arquivos, portanto, os navegadores ajustarão o nome sugerido, se necessário;
//	   * O download funciona apenas para URLs de mesma origem, ou os esquemas blob: e data: schemes;
//	   * A forma como os navegadores tratam os downloads varia de acordo com o navegador, as
//	     configurações do usuário e outros fatores. O usuário pode ser avisado antes do início de um
//	     download, ou o arquivo pode ser salvo automaticamente, ou pode ser aberto automaticamente,
//	     seja em um aplicativo externo ou no próprio navegador;
//	   * Se o cabeçalho Content-Disposition tiver informações diferentes do atributo download, o
//	     comportamento resultante pode ser diferente:
//	       * Se o cabeçalho especificar um nome de arquivo, ele terá prioridade sobre um nome de
//	         arquivo especificado no atributo download;
//	       * Se o cabeçalho especificar uma disposição de inline, o Chrome e o Firefox priorizarão o
//	         atributo e o tratarão como um download. Versões antigas do Firefox (antes de 82)
//	         priorizam o cabeçalho e exibirão o conteúdo inline.
func (e *TagA) Download(download string) (ref *TagA) {
	e.selfElement.Set("download", download)
	return e
}

// HRef
//
// English:
//
//	The URL that the hyperlink points to. Links are not restricted to HTTP-based URLs — they can use
//	any URL scheme supported by browsers:
//	  * Sections of a page with fragment URLs;
//	  * Pieces of media files with media fragments;
//	  * Telephone numbers with tel: URLs;
//	  * Email addresses with mailto: URLs;
//	  * While web browsers may not support other URL schemes, web sites can with
//	    registerProtocolHandler().
//
// Português:
//
//	A URL para a qual o hiperlink aponta. Os links não são restritos a URLs baseados em HTTP — eles
//	podem usar qualquer esquema de URL suportado pelos navegadores:
//	  * Seções de uma página com URLs de fragmento;
//	  * Pedaços de arquivos de mídia com fragmentos de mídia;
//	  * Números de telefone com tel: URLs;
//	  * Endereços de e-mail com mailto: URLs;
//	  * Embora os navegadores da Web possam não suportar outros esquemas de URL, os sites da Web podem
//	    com registerProtocolHandler().
func (e *TagA) HRef(href interface{}) (ref *TagA) {
	e.selfElement.Set("href", href)
	return e
}

// HRefLang
//
// English:
//
//	Hints at the human language of the linked URL. No built-in functionality. Allowed values are the
//	same as the global lang attribute.
//
// Português:
//
//	Dicas para a linguagem humana da URL vinculada. Nenhuma funcionalidade embutida. Os valores
//	permitidos são os mesmos do atributo lang global.
func (e *TagA) HRefLang(hreflang string) (ref *TagA) {
	e.selfElement.Set("hreflang", hreflang)
	return e
}

// Ping
//
// English:
//
//	A space-separated list of URLs. When the link is followed, the browser will send POST requests
//	with the body PING to the URLs. Typically for tracking.
//
// Português:
//
//	Uma lista de URLs separados por espaços. Quando o link for seguido, o navegador enviará
//	solicitações POST com o corpo PING para as URLs. Normalmente para rastreamento.
func (e *TagA) Ping(ping ...string) (ref *TagA) {
	switch e.tag {
	case KTagA:
	default:
		log.Printf("tag " + e.tag.String() + " does not support ping property")
	}

	e.selfElement.Set("hreflang", strings.Join(ping, " "))
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
//	 KRefPolicyNoReferrerWhenDowngrade: O cabeçalho Referer não será enviado para origens sem
//	   TLS (HTTPS).
//	 KRefPolicyOrigin: O referenciador enviado será limitado à origem da página de referência: seu
//	   esquema, host e porta.
//	 KRefPolicyOriginWhenCrossOrigin: O referenciador enviado para outras origens será limitado ao
//	   esquema, ao host e à porta. As navegações na mesma origem ainda incluirão o caminho.
//	 KRefPolicySameOrigin: Um referenciador será enviado para a mesma origem, mas as solicitações
//	   de origem cruzada não conterão informações de referenciador.
//	 KRefPolicyStrictOrigin: Só envie a origem do documento como referenciador quando o nível de
//	   segurança do protocolo permanecer o mesmo (HTTPS→HTTPS), mas não envie para um destino menos
//	   seguro (HTTPS→HTTP).
//	 KRefPolicyStrictOriginWhenCrossOrigin (padrão): Envie uma URL completa ao realizar uma
//	   solicitação de mesma origem, envie a origem apenas quando o nível de segurança do protocolo
//	   permanecer o mesmo (HTTPS→HTTPS) e não envie nenhum cabeçalho para um destino menos seguro
//	   (HTTPS→HTTP).
//	 KRefPolicyUnsafeUrl: O referenciador incluirá a origem e o caminho (mas não o fragmento, a
//	   senha ou o nome de usuário). Esse valor não é seguro porque vaza origens e caminhos de
//	   recursos protegidos por TLS para origens inseguras.
//
//	 Note:
//	   * Experimental. Expect behavior to change in the future. (04/2022)
func (e *TagA) ReferrerPolicy(referrerPolicy ReferrerPolicy) (ref *TagA) {
	e.selfElement.Set("referrerpolicy", referrerPolicy)
	return e
}

// Rel
//
// English:
//
//	The relationship of the linked URL as space-separated link types.
//
// Português:
//
//	O relacionamento da URL vinculada como tipos de link separados por espaço.
func (e *TagA) Rel(rel string) (ref *TagA) {
	switch e.tag {
	case KTagA:
	case KTagForm:
	default:
		log.Printf("tag " + e.tag.String() + " does not support rel property")
	}

	e.selfElement.Set("rel", rel)
	return e
}

// Target
//
// English:
//
// Where to display the linked URL, as the name for a browsing context (a tab, window, or <iframe>).
// The following keywords have special meanings for where to load the URL:
//
//	KTargetSelf: the current browsing context; (Default)
//	KTargetBlank: usually a new tab, but users can configure browsers to open a new window instead;
//	KTargetParent: the parent browsing context of the current one. If no parent, behaves as _self;
//	KTargetTop: the topmost browsing context (the "highest" context that's an ancestor of the current
//	  one). If no ancestors, behaves as _self.
//
//	Note:
//	  * Setting target="_blank" on <a> elements implicitly provides the same rel behavior as setting
//	    rel="noopener" which does not set window.opener. See browser compatibility for support
//	    status.
//
// Português:
//
// Onde exibir a URL vinculada, como o nome de um contexto de navegação (uma guia, janela ou
// <iframe>). As seguintes palavras-chave têm significados especiais para onde carregar o URL:
//
//	KTargetSelf: o contexto de navegação atual; (padrão)
//	KTargetBlank: geralmente uma nova guia, mas os usuários podem configurar os navegadores para
//	  abrir uma nova janela;
//	KTargetParent: o contexto de navegação pai do atual. Se nenhum pai, se comporta como _self;
//	KTargetTop: o contexto de navegação mais alto (o contexto "mais alto" que é um ancestral do
//	  atual). Se não houver ancestrais, se comporta como _self.
//
//	Nota:
//	  * Definir target="_blank" em elementos <a> fornece implicitamente o mesmo comportamento rel
//	    que definir rel="noopener" que não define window.opener. Consulte a compatibilidade do
//	    navegador para obter o status do suporte.
func (e *TagA) Target(target Target) (ref *TagA) {
	switch e.tag {
	case KTagA:
	case KTagForm:
	default:
		log.Printf("tag " + e.tag.String() + " does not support target property")
	}

	e.selfElement.Set("target", target.String())
	return e
}

// Type
//
// English:
//
// Hints at the linked URL's format with a MIME type. No built-in functionality.
//
// Português:
//
// Dicas no formato do URL vinculado com um tipo MIME. Nenhuma funcionalidade embutida.
func (e *TagA) Type(typeProperty Mime) (ref *TagA) {
	e.selfElement.Set("type", typeProperty)
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
func (e *TagA) CreateElement(tag Tag) (ref *TagA) {
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
func (e *TagA) AppendById(appendId string) (ref *TagA) {

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
func (e *TagA) Append(append interface{}) (ref *TagA) {
	switch append.(type) {
	case *TagA:
		e.selfElement.Call("appendChild", append.(*TagA).selfElement)
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
func (e *TagA) GetXY() (x, y int) {
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
func (e *TagA) GetX() (x int) {
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
func (e *TagA) GetY() (y int) {
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
func (e *TagA) GetTop() (top int) {
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
func (e *TagA) GetRight() (right int) {
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
func (e *TagA) GetBottom() (bottom int) {
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
func (e *TagA) GetLeft() (left int) {
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
func (e *TagA) GetBoundingBox() (x, y, width, height int) {
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
func (e *TagA) CollisionBoundingBox(elemnt CollisionBoundingBox) (collision bool) {
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
func (e *TagA) UpdateBoundingClientRect() (ref *TagA) {
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
func (e *TagA) SetXY(x, y int) (ref *TagA) {

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
func (e *TagA) SetDeltaX(delta int) (ref *TagA) {
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
func (e *TagA) SetDeltaY(delta int) (ref *TagA) {
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
func (e *TagA) SetX(x int) (ref *TagA) {

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
func (e *TagA) SetY(y int) (ref *TagA) {

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

func (e *TagA) Get() (el js.Value) {
	return e.selfElement
}

// Text
//
// English:
//
// Adds plain text to the tag's content.
//
// Text:
//
// Adiciona um texto simples ao conteúdo da tag.
func (e *TagA) Text(value string) (ref *TagA) {
	e.selfElement.Set("textContent", value)
	return e
}

// Html
//
// English:
//
// Adds HTML to the tag's content.
//
// Text:
//
// Adiciona HTML ao conteúdo da tag.
func (e *TagA) Html(value string) (ref *TagA) {
	e.selfElement.Set("innerHTML", value)
	return e
}

func (e *TagA) AddListenerAbort(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagA) RemoveListenerAbort() (ref *TagA) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagA) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagA) RemoveListenerAuxclick() (ref *TagA) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagA) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagA) RemoveListenerBeforeinput() (ref *TagA) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagA) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagA) RemoveListenerBeforematch() (ref *TagA) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagA) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagA) RemoveListenerBeforetoggle() (ref *TagA) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagA) AddListenerCancel(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagA) RemoveListenerCancel() (ref *TagA) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagA) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagA) RemoveListenerCanplay() (ref *TagA) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagA) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagA) RemoveListenerCanplaythrough() (ref *TagA) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagA) AddListenerChange(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagA) RemoveListenerChange() (ref *TagA) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagA) AddListenerClick(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagA) RemoveListenerClick() (ref *TagA) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagA) AddListenerClose(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagA) RemoveListenerClose() (ref *TagA) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagA) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagA) RemoveListenerContextlost() (ref *TagA) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagA) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagA) RemoveListenerContextmenu() (ref *TagA) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagA) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagA) RemoveListenerContextrestored() (ref *TagA) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagA) AddListenerCopy(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagA) RemoveListenerCopy() (ref *TagA) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagA) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagA) RemoveListenerCuechange() (ref *TagA) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagA) AddListenerCut(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagA) RemoveListenerCut() (ref *TagA) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagA) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagA) RemoveListenerDblclick() (ref *TagA) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagA) AddListenerDrag(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagA) RemoveListenerDrag() (ref *TagA) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagA) AddListenerDragend(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagA) RemoveListenerDragend() (ref *TagA) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagA) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagA) RemoveListenerDragenter() (ref *TagA) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagA) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagA) RemoveListenerDragleave() (ref *TagA) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagA) AddListenerDragover(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagA) RemoveListenerDragover() (ref *TagA) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagA) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagA) RemoveListenerDragstart() (ref *TagA) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagA) AddListenerDrop(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagA) RemoveListenerDrop() (ref *TagA) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagA) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagA) RemoveListenerDurationchange() (ref *TagA) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagA) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagA) RemoveListenerEmptied() (ref *TagA) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagA) AddListenerEnded(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagA) RemoveListenerEnded() (ref *TagA) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagA) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagA) RemoveListenerFormdata() (ref *TagA) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagA) AddListenerInput(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagA) RemoveListenerInput() (ref *TagA) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagA) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagA) RemoveListenerInvalid() (ref *TagA) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagA) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagA) RemoveListenerKeydown() (ref *TagA) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagA) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagA) RemoveListenerKeypress() (ref *TagA) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagA) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagA) RemoveListenerKeyup() (ref *TagA) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagA) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagA) RemoveListenerLoadeddata() (ref *TagA) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagA) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagA) RemoveListenerLoadedmetadata() (ref *TagA) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagA) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagA) RemoveListenerLoadstart() (ref *TagA) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagA) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagA) RemoveListenerMousedown() (ref *TagA) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagA) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagA) RemoveListenerMouseenter() (ref *TagA) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagA) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagA) RemoveListenerMouseleave() (ref *TagA) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagA) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagA) RemoveListenerMousemove() (ref *TagA) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagA) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagA) RemoveListenerMouseout() (ref *TagA) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagA) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagA) RemoveListenerMouseover() (ref *TagA) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagA) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagA) RemoveListenerMouseup() (ref *TagA) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagA) AddListenerPaste(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagA) RemoveListenerPaste() (ref *TagA) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagA) AddListenerPause(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagA) RemoveListenerPause() (ref *TagA) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagA) AddListenerPlay(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagA) RemoveListenerPlay() (ref *TagA) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagA) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagA) RemoveListenerPlaying() (ref *TagA) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagA) AddListenerProgress(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagA) RemoveListenerProgress() (ref *TagA) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagA) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagA) RemoveListenerRatechange() (ref *TagA) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagA) AddListenerReset(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagA) RemoveListenerReset() (ref *TagA) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagA) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagA) RemoveListenerScrollend() (ref *TagA) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagA) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagA) RemoveListenerSecuritypolicyviolation() (ref *TagA) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagA) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagA) RemoveListenerSeeked() (ref *TagA) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagA) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagA) RemoveListenerSeeking() (ref *TagA) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagA) AddListenerSelect(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagA) RemoveListenerSelect() (ref *TagA) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagA) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagA) RemoveListenerSlotchange() (ref *TagA) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagA) AddListenerStalled(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagA) RemoveListenerStalled() (ref *TagA) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagA) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagA) RemoveListenerSubmit() (ref *TagA) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagA) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagA) RemoveListenerSuspend() (ref *TagA) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagA) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagA) RemoveListenerTimeupdate() (ref *TagA) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagA) AddListenerToggle(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagA) RemoveListenerToggle() (ref *TagA) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagA) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagA) RemoveListenerVolumechange() (ref *TagA) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagA) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagA) RemoveListenerWaiting() (ref *TagA) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagA) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagA) RemoveListenerWebkitanimationend() (ref *TagA) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagA) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagA) RemoveListenerWebkitanimationiteration() (ref *TagA) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagA) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagA) RemoveListenerWebkitanimationstart() (ref *TagA) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagA) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagA) RemoveListenerWebkittransitionend() (ref *TagA) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagA) AddListenerWheel(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagA) RemoveListenerWheel() (ref *TagA) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagA) AddListenerBlur(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagA) RemoveListenerBlur() (ref *TagA) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagA) AddListenerError(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagA) RemoveListenerError() (ref *TagA) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagA) AddListenerFocus(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagA) RemoveListenerFocus() (ref *TagA) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagA) AddListenerLoad(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagA) RemoveListenerLoad() (ref *TagA) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagA) AddListenerResize(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagA) RemoveListenerResize() (ref *TagA) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagA) AddListenerScroll(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagA) RemoveListenerScroll() (ref *TagA) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagA) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagA) RemoveListenerAfterprint() (ref *TagA) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagA) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagA) RemoveListenerBeforeprint() (ref *TagA) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagA) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagA) RemoveListenerBeforeunload() (ref *TagA) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagA) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagA) RemoveListenerHashchange() (ref *TagA) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagA) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagA) RemoveListenerLanguagechange() (ref *TagA) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagA) AddListenerMessage(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagA) RemoveListenerMessage() (ref *TagA) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagA) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagA) RemoveListenerMessageerror() (ref *TagA) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagA) AddListenerOffline(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagA) RemoveListenerOffline() (ref *TagA) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagA) AddListenerOnline(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagA) RemoveListenerOnline() (ref *TagA) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagA) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagA) RemoveListenerPageswap() (ref *TagA) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagA) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagA) RemoveListenerPagehide() (ref *TagA) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagA) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagA) RemoveListenerPagereveal() (ref *TagA) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagA) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagA) RemoveListenerPageshow() (ref *TagA) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagA) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagA) RemoveListenerPopstate() (ref *TagA) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagA) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagA) RemoveListenerRejectionhandled() (ref *TagA) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagA) AddListenerStorage(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagA) RemoveListenerStorage() (ref *TagA) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagA) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagA) RemoveListenerUnhandledrejection() (ref *TagA) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagA) AddListenerUnload(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagA) RemoveListenerUnload() (ref *TagA) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagA) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagA) RemoveListenerReadystatechange() (ref *TagA) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagA) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagA) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagA) RemoveListenerVisibilitychange() (ref *TagA) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
