package html

import (
	"log"
	"math"
	"reflect"
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

// GetId #global
//
// English:
//
//	Return a unique id for an element
//
// The id attribute specifies a unique id for an HTML element (the value must be unique within the
// HTML document).
//
// The id attribute is most used to point to a style in a style sheet, and by JavaScript (via the HTML
// DOM) to manipulate the element with the specific id.
//
// Português:
//
//	Retorna um ID exclusivo para um elemento
//
// O atributo id especifica um id exclusivo para um elemento HTML (o valor deve ser exclusivo no
// documento HTML).
//
// O atributo id é mais usado para apontar para um estilo em uma folha de estilo, e por JavaScript
// (através do HTML DOM) para manipular o elemento com o id específico.
func (e *TagSource) GetId() (id string) {
	return e.id
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
func (e *TagSource) ListenerAddReflect(event string, params []interface{}, functions []reflect.Value, reference any) (ref *TagSource) {
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
func (e *TagSource) ListenerRemove(event string) (ref *TagSource) {
	e.commonEvents.ListenerRemove(event)
	return e
}

// SetAttribute
//
// English:
//
//	Sets an attribute on the tag element with the specified key and value.
//
// Português:
//
//	Define um atributo no elemento da tag com a chave e valor especificada.
func (e *TagSource) SetAttribute(key string, value any) (ref *TagSource) {
	e.selfElement.Call("setAttribute", key, value)
	return e
}

// GetAttribute
//
// English:
//
//	Gets an attribute on the tag element with the specified key and value.
//
// Português:
//
//	Retorna um atributo no elemento da tag com a chave e valor especificada.
func (e *TagSource) GetAttribute(key string) (value string) {
	return e.selfElement.Call("getAttribute", key).String()
}

// RemoveAttribute
//
// English:
//
//	Remove an attribute from the tag element.
//
// Português:
//
//	Remove um atributo no elemento da tag.
func (e *TagSource) RemoveAttribute(key string) (value string) {
	return e.selfElement.Call("removeAttribute", key).String()
}

// Remove
//
// English:
//
//	Removes a child node from the DOM and returns the removed node.
//
// Português:
//
//	Remove um nó filho do DOM e retorna o nó removido.
func (e *TagSource) Remove(elements ...Compatible) (ref *TagSource) {
	for _, element := range elements {
		e.selfElement.Call("removeChild", element.Get())
	}

	return e
}

// SetZIndex
//
// English:
//
//	The z-index CSS property sets the z-order of a positioned element and its descendants or flex and grid items.
//
// Overlapping elements with a larger z-index cover those with a smaller one.
//
// Português:
//
//	A propriedade CSS z-index define a ordem z de um elemento posicionado e seus descendentes ou itens flexíveis e de
//	grade.
//
// Elementos sobrepostos com um z-index maior cobrem aqueles com um z-index menor.
func (e *TagSource) SetZIndex(index int) (ref *TagSource) {
	e.selfElement.Call("setAttribute", "zIndex", strconv.FormatInt(int64(index), 10))
	return e
}

// GetZIndex
//
// English:
//
//	The z-index CSS property sets the z-order of a positioned element and its descendants or flex and grid items.
//
// Overlapping elements with a larger z-index cover those with a smaller one.
//
// Português:
//
//	A propriedade CSS z-index define a ordem z de um elemento posicionado e seus descendentes ou itens flexíveis e de
//	grade.
//
// Elementos sobrepostos com um z-index maior cobrem aqueles com um z-index menor.
func (e *TagSource) GetZIndex() (index int) {
	z := e.selfElement.Call("getAttribute", "zIndex").String()
	if z == "auto" {
		return math.MinInt
	}

	index, _ = strconv.Atoi(z)
	return
}

// RemoveZIndex
//
// English:
//
//	The z-index CSS property sets the z-order of a positioned element and its descendants or flex and grid items.
//
// Overlapping elements with a larger z-index cover those with a smaller one.
//
// Português:
//
//	A propriedade CSS z-index define a ordem z de um elemento posicionado e seus descendentes ou itens flexíveis e de
//	grade.
//
// Elementos sobrepostos com um z-index maior cobrem aqueles com um z-index menor.
func (e *TagSource) RemoveZIndex() (ref *TagSource) {
	e.selfElement.Call("removeAttribute", "zIndex")
	return e
}
