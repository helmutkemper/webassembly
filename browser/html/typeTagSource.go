package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"log"
	"strconv"
	"strings"
	"syscall/js"
)

// TagSource
//
// English:
//
// The <source> HTML element specifies multiple media resources for the <picture>, the <audio> element, or the <video>
// element.
//
// It is an empty element, meaning that it has no content and does not have a closing tag. It is commonly used to offer
// the same media content in multiple file formats in order to provide compatibility with a broad range of browsers
// given their differing support for image file formats and media file formats.
//
// Português:
//
// O elemento HTML <source> especifica vários recursos de mídia para o elemento <picture>, o elemento <audio> ou o
// elemento <video>. É um elemento vazio, o que significa que não possui conteúdo e não possui uma tag de fechamento.
// É comumente usado para oferecer o mesmo conteúdo de mídia em vários formatos de arquivo para fornecer compatibilidade
// com uma ampla variedade de navegadores, devido ao suporte diferente para formatos de arquivo de imagem e formatos de
// arquivo de mídia.
type TagSource struct {
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
func (e *TagSource) Reference(reference **TagSource) (ref *TagSource) {
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
func (e *TagSource) AccessKey(key string) (ref *TagSource) {
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
func (e *TagSource) AutoCapitalize(value interface{}) (ref *TagSource) {
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
func (e *TagSource) AutoFocus(autofocus bool) (ref *TagSource) {
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
func (e *TagSource) Class(class ...string) (ref *TagSource) {
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
func (e *TagSource) ContentEditable(editable bool) (ref *TagSource) {
	e.selfElement.Set("contenteditable", editable)
	return e
}

// Data #global
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
func (e *TagSource) Data(data map[string]string) (ref *TagSource) {
	for k, v := range data {
		e.selfElement.Set(" data-"+k, v)
	}
	return e
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
func (e *TagSource) Dir(dir Dir) (ref *TagSource) {
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
func (e *TagSource) Draggable(draggable Draggable) (ref *TagSource) {
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
func (e *TagSource) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagSource) {
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
func (e *TagSource) Hidden() (ref *TagSource) {
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
func (e *TagSource) Id(id string) (ref *TagSource) {
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
func (e *TagSource) ItemProp(itemprop string) (ref *TagSource) {
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
func (e *TagSource) ItemRef(itemref string) (ref *TagSource) {
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
func (e *TagSource) ItemScope(itemscope bool) (ref *TagSource) {
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
func (e *TagSource) ItemType(itemType string) (ref *TagSource) {
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
func (e *TagSource) Lang(language Language) (ref *TagSource) {
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
func (e *TagSource) Part(part ...string) (ref *TagSource) {
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
func (e *TagSource) Nonce(nonce string) (ref *TagSource) {
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
func (e *TagSource) Slot(slot string) (ref *TagSource) {
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
func (e *TagSource) Spellcheck(spell bool) (ref *TagSource) {
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
func (e *TagSource) Style(style string) (ref *TagSource) {
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
func (e *TagSource) TabIndex(index int) (ref *TagSource) {
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
func (e *TagSource) Title(title string) (ref *TagSource) {
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
func (e *TagSource) Translate(translate Translate) (ref *TagSource) {
	e.selfElement.Set("translate", translate.String())
	return e
}

// #global - end -------------------------------------------------------------------------------------------------------

// Type
//
// English:
//
// The MIME media type of the resource, optionally with a codecs parameter.
//
// Português:
//
// O tipo de mídia MIME do recurso, opcionalmente com um parâmetro codecs.
func (e *TagSource) Type(value interface{}) (ref *TagSource) {
	if converted, ok := value.(Mime); ok {
		e.selfElement.Set("type", converted.String())
		return e
	}

	e.selfElement.Set("type", value)
	return e
}

// Src
//
// English:
//
// Required if the source element's parent is an <audio> and <video> element, but not allowed if the source element's
// parent is a <picture> element.
//
// Address of the media resource.
//
// Português:
//
// Obrigatório se o pai do elemento de origem for um elemento <audio> e <video>, mas não permitido se o pai do elemento
// de origem for um elemento <picture>.
//
// Endereço do recurso de mídia.
func (e *TagSource) Src(value string) (ref *TagSource) {
	e.selfElement.Set("src", value)
	return e
}

// SrcSet
//
// English:
//
// Required if the source element's parent is a <picture> element, but not allowed if the source element's parent is an
// <audio> or <video> element.
//
// A list of one or more strings, separated by commas, indicating a set of possible images represented by the source for
// the browser to use. Each string is composed of:
//   - One URL specifying an image.
//   - A width descriptor, which consists of a string containing a positive integer directly followed by "w", such as
//     300w. The default value, if missing, is the infinity.
//   - A pixel density descriptor, that is a positive floating number directly followed by "x". The default value, if
//     missing, is 1x.
//
// Each string in the list must have at least a width descriptor or a pixel density descriptor to be valid. Among the
// list, there must be only one string containing the same tuple of width descriptor and pixel density descriptor.
// The browser chooses the most adequate image to display at a given point of time. If width descriptors are used,
// the sizes attribute must also be present, or the srcset value will be ignored.
//
// Português:
//
// Obrigatório se o pai do elemento de origem for um elemento <picture>, mas não permitido se o pai do elemento de
// origem for um elemento <audio> ou <video>.
//
// Uma lista de uma ou mais strings, separadas por vírgulas, indicando um conjunto de imagens possíveis representadas
// pela fonte para o navegador usar. Cada corda é composta por:
//   - Um URL especificando uma imagem.
//   - Um descritor de largura, que consiste em uma string contendo um inteiro positivo seguido diretamente por "w",
//     como 300w. O valor padrão, se ausente, é o infinito.
//   - Um descritor de densidade de pixels, que é um número flutuante positivo seguido diretamente por "x". O valor
//     padrão, se ausente, é 1x.
//
// Cada string na lista deve ter pelo menos um descritor de largura ou um descritor de densidade de pixels para ser
// válido. Entre a lista, deve haver apenas uma string contendo a mesma tupla de descritor de largura e descritor de
// densidade de pixels. O navegador escolhe a imagem mais adequada para exibir em um determinado momento. Se forem
// usados descritores de largura, o atributo de tamanhos também deve estar presente, ou o valor srcset será ignorado.
func (e *TagSource) SrcSet(value string) (ref *TagSource) {
	e.selfElement.Set("srcset", value)
	return e
}

// Sizes
//
// English:
//
// Allowed if the source element's parent is a <picture> element, but not allowed if the source element's parent is an
// <audio> or <video> element.
//
// A list of source sizes that describes the final rendered width of the image represented by the source. Each source
// size consists of a comma-separated list of media condition-length pairs. This information is used by the browser to
// determine, before laying the page out, which image defined in srcset to use. Please note that sizes will have its
// effect only if width dimension descriptors are provided with srcset instead of pixel ratio values (200w instead of
// 2x for example).
//
// Português:
//
// Permitido se o pai do elemento de origem for um elemento <picture>, mas não permitido se o pai do elemento de origem
// for um elemento <audio> ou <video>.
//
// Uma lista de tamanhos de origem que descreve a largura final renderizada da imagem representada pela origem. Cada
// tamanho de origem consiste em uma lista separada por vírgulas de pares de condição-comprimento de mídia. Essas
// informações são usadas pelo navegador para determinar, antes de colocar a página, qual imagem definida em srcset
// usar. Observe que os tamanhos terão seu efeito somente se os descritores de dimensão de largura forem fornecidos com
// srcset em vez de valores de proporção de pixel (200w em vez de 2x, por exemplo).
func (e *TagSource) Sizes(value string) (ref *TagSource) {
	e.selfElement.Set("sizes", value)
	return e
}

// Media
//
// English:
//
// Allowed if the source element's parent is a <picture> element, but not allowed if the source element's parent is an
// <audio> or <video> element.
//
// Media query of the resource's intended media.
//
// Português:
//
// Permitido se o pai do elemento de origem for um elemento <picture>, mas não permitido se o pai do elemento de origem
// for um elemento <audio> ou <video>.
//
// Consulta de mídia da mídia pretendida do recurso.
func (e *TagSource) Media(value string) (ref *TagSource) {
	e.selfElement.Set("media", value)
	return e
}

// Height
//
// English:
//
// Allowed if the source element's parent is a <picture> element, but not allowed if the source element's parent is an
// <audio> or <video> element.
//
// The intrinsic height of the image, in pixels. Must be an integer without a unit.
//
// Português:
//
// Permitido se o pai do elemento de origem for um elemento <picture>, mas não permitido se o pai do elemento de origem
// for um elemento <audio> ou <video>.
//
// A altura intrínseca da imagem, em pixels. Deve ser um número inteiro sem uma unidade.
func (e *TagSource) Height(value interface{}) (ref *TagSource) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Set("height", p)
		return e
	}

	e.selfElement.Set("height", value)
	return e
}

// Width
//
// English:
//
// Allowed if the source element's parent is a <picture> element, but not allowed if the source element's parent is an
// <audio> or <video> element.
//
// The intrinsic width of the image in pixels. Must be an integer without a unit.
//
// Português:
//
// Permitido se o pai do elemento de origem for um elemento <picture>, mas não permitido se o pai do elemento de origem
// for um elemento <audio> ou <video>.
//
// A largura intrínseca da imagem em pixels. Deve ser um número inteiro sem uma unidade.
func (e *TagSource) Width(value interface{}) (ref *TagSource) {
	if converted, ok := value.(float32); ok {
		p := strconv.FormatFloat(100.0*float64(converted), 'g', -1, 64) + "%"
		e.selfElement.Set("width", p)
		return e
	}

	e.selfElement.Set("width", value)
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
func (e *TagSource) CreateElement() (ref *TagSource) {
	e.selfElement = js.Global().Get("document").Call("createElement", "source")
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
func (e *TagSource) AppendById(appendId string) (ref *TagSource) {

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
func (e *TagSource) Append(elements ...Compatible) (ref *TagSource) {
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
func (e *TagSource) AppendToStage() (ref *TagSource) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagSource) Get() (el js.Value) {
	return e.selfElement
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
func (e *TagSource) Init() (ref *TagSource) {
	e.CreateElement()
	e.prepareStageReference()

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
func (e *TagSource) prepareStageReference() {
	e.stage = js.Global().Get("document").Get("body")
}
func (e *TagSource) AddListenerAbort(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerAbort() (ref *TagSource) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagSource) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerAuxclick() (ref *TagSource) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagSource) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerBeforeinput() (ref *TagSource) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagSource) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerBeforematch() (ref *TagSource) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagSource) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerBeforetoggle() (ref *TagSource) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagSource) AddListenerCancel(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerCancel() (ref *TagSource) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagSource) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerCanplay() (ref *TagSource) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagSource) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerCanplaythrough() (ref *TagSource) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagSource) AddListenerChange(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerChange() (ref *TagSource) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagSource) AddListenerClick(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerClick() (ref *TagSource) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagSource) AddListenerClose(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerClose() (ref *TagSource) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagSource) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerContextlost() (ref *TagSource) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagSource) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerContextmenu() (ref *TagSource) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagSource) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerContextrestored() (ref *TagSource) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagSource) AddListenerCopy(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerCopy() (ref *TagSource) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagSource) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerCuechange() (ref *TagSource) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagSource) AddListenerCut(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerCut() (ref *TagSource) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagSource) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerDblclick() (ref *TagSource) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagSource) AddListenerDrag(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerDrag() (ref *TagSource) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagSource) AddListenerDragend(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerDragend() (ref *TagSource) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagSource) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerDragenter() (ref *TagSource) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagSource) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerDragleave() (ref *TagSource) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagSource) AddListenerDragover(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerDragover() (ref *TagSource) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagSource) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerDragstart() (ref *TagSource) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagSource) AddListenerDrop(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerDrop() (ref *TagSource) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagSource) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerDurationchange() (ref *TagSource) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagSource) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerEmptied() (ref *TagSource) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagSource) AddListenerEnded(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerEnded() (ref *TagSource) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagSource) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerFormdata() (ref *TagSource) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagSource) AddListenerInput(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerInput() (ref *TagSource) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagSource) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerInvalid() (ref *TagSource) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagSource) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerKeydown() (ref *TagSource) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagSource) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerKeypress() (ref *TagSource) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagSource) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerKeyup() (ref *TagSource) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagSource) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerLoadeddata() (ref *TagSource) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagSource) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerLoadedmetadata() (ref *TagSource) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagSource) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerLoadstart() (ref *TagSource) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagSource) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerMousedown() (ref *TagSource) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagSource) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerMouseenter() (ref *TagSource) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagSource) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerMouseleave() (ref *TagSource) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagSource) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerMousemove() (ref *TagSource) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagSource) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerMouseout() (ref *TagSource) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagSource) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerMouseover() (ref *TagSource) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagSource) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerMouseup() (ref *TagSource) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagSource) AddListenerPaste(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerPaste() (ref *TagSource) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagSource) AddListenerPause(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerPause() (ref *TagSource) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagSource) AddListenerPlay(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerPlay() (ref *TagSource) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagSource) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerPlaying() (ref *TagSource) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagSource) AddListenerProgress(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerProgress() (ref *TagSource) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagSource) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerRatechange() (ref *TagSource) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagSource) AddListenerReset(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerReset() (ref *TagSource) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagSource) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerScrollend() (ref *TagSource) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagSource) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerSecuritypolicyviolation() (ref *TagSource) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagSource) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerSeeked() (ref *TagSource) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagSource) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerSeeking() (ref *TagSource) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagSource) AddListenerSelect(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerSelect() (ref *TagSource) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagSource) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerSlotchange() (ref *TagSource) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagSource) AddListenerStalled(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerStalled() (ref *TagSource) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagSource) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerSubmit() (ref *TagSource) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagSource) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerSuspend() (ref *TagSource) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagSource) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerTimeupdate() (ref *TagSource) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagSource) AddListenerToggle(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerToggle() (ref *TagSource) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagSource) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerVolumechange() (ref *TagSource) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagSource) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerWaiting() (ref *TagSource) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagSource) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerWebkitanimationend() (ref *TagSource) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagSource) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerWebkitanimationiteration() (ref *TagSource) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagSource) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerWebkitanimationstart() (ref *TagSource) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagSource) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerWebkittransitionend() (ref *TagSource) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagSource) AddListenerWheel(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerWheel() (ref *TagSource) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagSource) AddListenerBlur(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerBlur() (ref *TagSource) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagSource) AddListenerError(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerError() (ref *TagSource) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagSource) AddListenerFocus(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerFocus() (ref *TagSource) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagSource) AddListenerLoad(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerLoad() (ref *TagSource) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagSource) AddListenerResize(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerResize() (ref *TagSource) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagSource) AddListenerScroll(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerScroll() (ref *TagSource) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagSource) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerAfterprint() (ref *TagSource) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagSource) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerBeforeprint() (ref *TagSource) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagSource) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerBeforeunload() (ref *TagSource) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagSource) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerHashchange() (ref *TagSource) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagSource) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerLanguagechange() (ref *TagSource) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagSource) AddListenerMessage(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerMessage() (ref *TagSource) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagSource) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerMessageerror() (ref *TagSource) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagSource) AddListenerOffline(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerOffline() (ref *TagSource) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagSource) AddListenerOnline(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerOnline() (ref *TagSource) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagSource) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerPageswap() (ref *TagSource) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagSource) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerPagehide() (ref *TagSource) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagSource) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerPagereveal() (ref *TagSource) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagSource) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerPageshow() (ref *TagSource) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagSource) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerPopstate() (ref *TagSource) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagSource) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerRejectionhandled() (ref *TagSource) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagSource) AddListenerStorage(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerStorage() (ref *TagSource) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagSource) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerUnhandledrejection() (ref *TagSource) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagSource) AddListenerUnload(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerUnload() (ref *TagSource) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagSource) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerReadystatechange() (ref *TagSource) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagSource) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagSource) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagSource) RemoveListenerVisibilitychange() (ref *TagSource) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
