package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/mouse"
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"
	"syscall/js"
)

// TagVideo
//
// English:
//
// The <video> HTML element embeds a media player which supports video playback into the document.
//
// You can use <video> for audio content as well, but the <audio> element may provide a more appropriate user
// experience.
//
// Português:
//
// O elemento HTML <video> incorpora um media player que suporta a reprodução de vídeo no documento.
//
// Você também pode usar <video> para conteúdo de áudio, mas o elemento <audio> pode fornecer uma experiência de usuário
// mais apropriada.
type TagVideo struct {

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

	x int
	y int

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

	fnCanplay        *js.Func
	fnCanPlayThrough *js.Func
	fnComplete       *js.Func
	fnDurationChange *js.Func
	fnEmptied        *js.Func
	fnEnded          *js.Func
	fnLoadedData     *js.Func
	fnLoadedMetadata *js.Func
	fnPause          *js.Func
	fnPlay           *js.Func
	fnPlaying        *js.Func
	fnProgress       *js.Func
	fnRateChange     *js.Func
	fnSeeked         *js.Func
	fnSeeking        *js.Func
	fnStalled        *js.Func
	fnSuspend        *js.Func
	fnTimeUpdate     *js.Func
	fnVolumeChange   *js.Func
	fnWaiting        *js.Func
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
func (e *TagVideo) Reference(reference **TagVideo) (ref *TagVideo) {
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
func (e *TagVideo) AccessKey(key string) (ref *TagVideo) {
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
func (e *TagVideo) Autofocus(autofocus bool) (ref *TagVideo) {
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
func (e *TagVideo) Class(class ...string) (ref *TagVideo) {
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
func (e *TagVideo) ContentEditable(editable bool) (ref *TagVideo) {
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
func (e *TagVideo) Data(data map[string]string) (ref *TagVideo) {
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
func (e *TagVideo) Dir(dir Dir) (ref *TagVideo) {
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
func (e *TagVideo) Draggable(draggable Draggable) (ref *TagVideo) {
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
func (e *TagVideo) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagVideo) {
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
func (e *TagVideo) Hidden() (ref *TagVideo) {
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
func (e *TagVideo) Id(id string) (ref *TagVideo) {
	e.id = id
	e.selfElement.Set("id", id)
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
func (e *TagVideo) ItemProp(itemprop string) (ref *TagVideo) {
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
func (e *TagVideo) ItemRef(itemref string) (ref *TagVideo) {
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
func (e *TagVideo) ItemType(itemType string) (ref *TagVideo) {
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
func (e *TagVideo) Lang(language Language) (ref *TagVideo) {
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
func (e *TagVideo) Part(part ...string) (ref *TagVideo) {
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
func (e *TagVideo) Nonce(nonce string) (ref *TagVideo) {
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
func (e *TagVideo) Slot(slot string) (ref *TagVideo) {
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
func (e *TagVideo) Spellcheck(spell bool) (ref *TagVideo) {
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
func (e *TagVideo) Style(style string) (ref *TagVideo) {
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
func (e *TagVideo) TabIndex(index int) (ref *TagVideo) {
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
func (e *TagVideo) Title(title string) (ref *TagVideo) {
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
func (e *TagVideo) Translate(translate Translate) (ref *TagVideo) {
	e.selfElement.Set("translate", translate.String())
	return e
}

// AutoPlay
//
// English:
//
// If specified, the video automatically begins to play back as soon as it can do so without stopping to finish loading
// the data.
//
//	Input:
//	  autoplay: the video automatically begins to play back
//
// To disable video autoplay, autoplay="false" will not work; the video will autoplay if the attribute is there in the
// <video> tag at all. To remove autoplay, the attribute needs to be removed altogether.
//
// In some browsers (e.g. Chrome 70.0) autoplay doesn't work if no muted attribute is present.
//
//	Notes:
//	 * Sites that automatically play audio (or videos with an audio track) can be an unpleasant experience for users,
//	   so should be avoided when possible. If you must offer autoplay functionality, you should make it opt-in
//	   (requiring a user to specifically enable it). However, this can be useful when creating media elements whose
//	   source will be set at a later time, under user control. See our autoplay guide for additional information about
//	   how to properly use autoplay.
//
// Português:
//
// Se especificado, o vídeo começará a ser reproduzido automaticamente assim que puder, sem parar para concluir o
// carregamento dos dados.
//
//	Entrada:
//	  autoplay: o vídeo começa a ser reproduzido automaticamente
//
// Para desabilitar a reprodução automática de vídeo, autoplay="false" não funcionará; o vídeo será reproduzido
// automaticamente se o atributo estiver na tag <video>. Para remover a reprodução automática, o atributo precisa ser
// removido completamente.
//
// Em alguns navegadores (por exemplo, Chrome 70.0), a reprodução automática não funciona se nenhum atributo mudo
// estiver presente.
//
//	Notas:
//	 * Sites que reproduzem áudio automaticamente (ou vídeos com faixa de áudio) podem ser uma experiência desagradável
//	   para os usuários, portanto, devem ser evitados sempre que possível. Se você deve oferecer a funcionalidade de
//	   reprodução automática, você deve ativá-la (exigindo que um usuário a habilite especificamente). No entanto,
//	   isso pode ser útil ao criar elementos de mídia cuja origem será definida posteriormente, sob controle do
//	   usuário. Consulte nosso guia de reprodução automática para obter informações adicionais sobre como usar a
//	   reprodução automática corretamente.
func (e *TagVideo) AutoPlay(autoplay bool) (ref *TagVideo) {
	e.selfElement.Set("autoplay", autoplay)
	return e
}

// Controls
//
// English:
//
// If this attribute is present, the browser will offer controls to allow the user to control video playback,
// including volume, seeking, and pause/resume playback.
//
//	Input:
//	  controls: the browser will offer controls to allow the user to control video playback
//
// Português:
//
// Se esse atributo estiver presente, o navegador oferecerá controles para permitir que o usuário controle a reprodução
// do vídeo, incluindo volume, busca e pausar e retomar a reprodução.
//
//	Entrada:
//	  controls: o navegador oferecerá controles para permitir que o usuário controle a reprodução do vídeo
func (e *TagVideo) Controls(controls bool) (ref *TagVideo) {
	e.selfElement.Set("controls", controls)
	return e
}

// Crossorigin
//
// English:
//
// This enumerated attribute indicates whether to use CORS to fetch the related video.
//
//	Input:
//	  value: indicates whether to use CORS to fetch the related video
//	    const: KCrossOrigin... (e.g. KCrossOriginAnonymous)
//	    any other type: interface{}
//
// CORS-enabled resources can be reused in the <canvas> element without being tainted.
//
// When not present, the resource is fetched without a CORS request (i.e. without sending the Origin: HTTP header),
// preventing its non-tainted used in <canvas> elements. If invalid, it is handled as if the enumerated keyword
// anonymous was used. See CORS settings attributes for additional information.
//
// Português:
//
// This enumerated attribute indicates whether to use CORS to fetch the related video.
//
//	Entrada:
//	  value: indica se deve usar CORS para buscar o vídeo relacionado
//	    const: KCrossOrigin... (ex. KCrossOriginAnonymous)
//	    qualquer outro tipo: interface{}
//
// CORS-enabled resources can be reused in the <canvas> element without being tainted.
//
// Quando não está presente, o recurso é buscado sem uma solicitação CORS (ou seja, sem enviar o cabeçalho
// Origin: HTTP), evitando que seu uso não contaminado em elementos <canvas>. Se inválido, é tratado como se a
// palavra-chave enumerada anônimo fosse usada. Consulte Atributos de configurações do CORS para obter informações
// adicionais.
func (e *TagVideo) Crossorigin(value interface{}) (ref *TagVideo) {
	if converted, ok := value.(CrossOrigin); ok {
		e.selfElement.Set("crossorigin", converted.String())
		return e
	}

	e.selfElement.Set("crossorigin", value)
	return e
}

// Height
//
// English:
//
// The height of the video's display area, in CSS pixels (absolute values only; no percentages).
//
//	Input:
//	  value: the height of the video's display area
//
// Português:
//
// A altura da área de exibição do vídeo, em pixels CSS (somente valores absolutos; sem porcentagens).
//
//	Entrada:
//	  value: a altura da área de exibição do vídeo
func (e *TagVideo) Height(value float64) (ref *TagVideo) {
	e.selfElement.Set("height", value)
	return e
}

// Loop
//
// English:
//
// If specified, the browser will automatically seek back to the start upon reaching the end of the video.
//
//	Input:
//	  value: the browser will automatically seek back to the start upon reaching the end of the video.
//
// Português:
//
// Se especificado, o navegador retornará automaticamente ao início ao chegar ao final do vídeo.
//
//	Entrada:
//	  value: o navegador retornará automaticamente ao início ao chegar ao final do vídeo.
func (e *TagVideo) Loop(value bool) (ref *TagVideo) {
	e.selfElement.Set("loop", value)
	return e
}

// Muted
//
// English:
//
// Indicates the default setting of the audio contained in the video. If set, the audio will be initially silenced.
//
//	Input:
//	  value: If true, the audio will be initially silenced.
//
// Its default value is false, meaning that the audio will be played when the video is played.
//
// Português:
//
// Indica o padrão do áudio contido no vídeo. Se definido, o áudio será silenciado inicialmente.
//
//	Entrada:
//	  value: Se true, o áudio será silenciado inicialmente.
//
// O valor padrão é false, significando que o áudio será tocado quando o vídeo for tocado.
func (e *TagVideo) Muted(value bool) (ref *TagVideo) {
	e.selfElement.Set("muted", value)
	return e
}

// PlaySinLine
//
// English:
//
// Indicating that the video is to be played "inline", that is within the element's playback area.
//
//	Input:
//	  value: If true, the video is to be played "inline", that is within the element's playback area.
//
// Note that the absence of this attribute does not imply that the video will always be played in fullscreen.
//
// Português:
//
// Indicando que o vídeo deve ser reproduzido "inline", ou seja, dentro da área de reprodução do elemento.
//
//	Entrada:
//	  value: Se true, o vídeo deve ser reproduzido "inline", ou seja, dentro da área de reprodução do elemento.
//
// Observe que a ausência deste atributo não implica que o vídeo seja sempre reproduzido em tela cheia.
func (e *TagVideo) PlaySinLine(value bool) (ref *TagVideo) {
	e.selfElement.Set("playsinline", value)
	return e
}

// Poster
//
// English:
//
// A URL for an image to be shown while the video is downloading.
//
//	Input:
//	  value: A URL for an image to be shown while the video is downloading.
//
// If this attribute isn't specified, nothing is displayed until the first frame is available, then the first frame is
// shown as the poster frame.
//
// Português:
//
// Um URL para uma imagem a ser exibida durante o download do vídeo.
//
//	Entrada:
//	  value: Um URL para uma imagem a ser exibida durante o download do vídeo.
//
// Se este atributo não for especificado, nada será exibido até que o primeiro quadro esteja disponível, então o
// primeiro quadro será mostrado como o quadro de pôster.
func (e *TagVideo) Poster(value string) (ref *TagVideo) {
	e.selfElement.Set("poster", value)
	return e
}

// Preload
//
// English:
//
// This enumerated attribute is intended to provide a hint to the browser about what the author thinks will lead to the
// best user experience with regards to what content is loaded before the video is played.
//
//	Input:
//	  value: The preload attribute specifies what the user agent should do before the video is played.
//	    Const: KPreload... (e.g. KPreloadMetadata)
//	    any other type: interface{}
//
//	Notes:
//	 * The autoplay attribute has precedence over preload. If autoplay is specified, the browser would obviously need
//	   to start downloading the video for playback.
//	 * The specification does not force the browser to follow the value of this attribute; it is a mere hint.
//
// Português:
//
// Esse atributo enumerado destina-se a fornecer uma dica ao navegador sobre o que o autor acha que levará à melhor
// experiência do usuário em relação ao conteúdo carregado antes da reprodução do vídeo.
//
//	Entrada:
//	  value: O atributo preload especifica o que o navegador deve fazer antes da reprodução do vídeo.
//	    Const: KPreload... (ex. KPreloadMetadata)
//	    qualquer outro tipo: interface{}
//
//	Notas:
//	 * O atributo de reprodução automática tem precedência sobre o pré-carregamento. Se a reprodução automática for
//	   especificada, o navegador obviamente precisaria iniciar o download do vídeo para reprodução.
//	 * A especificação não força o navegador a seguir o valor deste atributo; é uma mera dica.
func (e *TagVideo) Preload(value interface{}) (ref *TagVideo) {
	if converted, ok := value.(Preload); ok {
		e.selfElement.Set("preload", converted.String())
		return e
	}

	e.selfElement.Set("preload", value)
	return e
}

// Src
//
// English:
//
// The URL of the video to embed. This is optional; you may instead use the <source> element within the video block to
// specify the video to embed.
//
//	Input:
//	  value: The URL of the video to embed.
//
// Português:
//
// O URL do vídeo a ser incorporado. Isso é opcional; você pode usar o elemento <source> dentro do bloco de vídeo para
// especificar o vídeo a ser incorporado.
//
//	Entrada:
//	  value: O URL do vídeo a ser incorporado.
func (e *TagVideo) Src(value string) (ref *TagVideo) {
	e.selfElement.Set("src", value)
	return e
}

// Width
//
// English:
//
// The width of the video's display area, in CSS pixels (absolute values only; no percentages).
//
// Português:
//
// A largura da área de exibição do vídeo, em pixels CSS (somente valores absolutos; sem porcentagens).
func (e *TagVideo) Width(value float64) (ref *TagVideo) {
	e.selfElement.Set("width", value)
	return e
}

// AddListenerCanPlay
//
// Enclish:
//
// Adds a video event litener can play, equivalent to the JavaScript command addEventListener('canplay',fn).
//
//	Input:
//	  mouseEvet: pointer to channel event.Data
//
// Português:
//
// Adiciona um ouvinte de evento de vídeo pode tocar, equivalente ao comando JavaScript addEventListener('canplay',fn).
//
//	Entrada:
//	  mouseEvet: ponteiro para o channel event.Data
//
//	Example: / Exemplo:
//	  videoEvent := make(chan event.Data)
//
//	  s1 := factoryBrowser.NewTagVideo().AddListenerCanPlay(&videoEvent).Controls(true).Width(250).Append(
//	    factoryBrowser.NewTagSource().Src("https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm").Type("video/webm"),
//	    factoryBrowser.NewTagSource().Src("https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.mp4").Type("video/mp4"),
//	  )
//
//	  stage.Append(s1)
//
//	  go func() {
//	    for {
//	      select {
//	        case converted := <-videoEvent:
//	        log.Printf("%+v", converted)
//	      }
//	    }
//	  }()
func (e *TagVideo) AddListenerCanPlay(evet *chan event.Data) (ref *TagVideo) {
	var fn js.Func

	if e.fnCanplay == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if len(args) == 0 {
				return nil
			}

			*evet <- event.EventManager(event.KEventCanPlay, this, args)
			return nil
		})
		e.fnCanplay = &fn
	}

	e.selfElement.Call(
		"addEventListener",
		"canplay",
		*e.fnCanplay,
	)
	return e
}

// RemoveListenerCanPlay
//
// English:
//
// Removes a video event litener can play, equivalent to the JavaScript command RemoveEventListener('canplay',fn).
//
// Português:
//
// Remove um ouvinte de evento de vídeo pode tocar, equivalente ao comando JavaScript RemoveEventListener('canplay',fn).
func (e *TagVideo) RemoveListenerCanPlay() (ref *TagVideo) {
	if e.fnCanplay == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"canplay",
		*e.fnCanplay,
	)
	e.fnCanplay = nil
	return e
}

func (e *TagVideo) AddListenerCanPlayThrough(evet *chan event.Data) (ref *TagVideo) {
	var fn js.Func

	if e.fnCanPlayThrough == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if len(args) == 0 {
				return nil
			}

			*evet <- event.EventManager(event.KEventCanPlayThrough, this, args)
			return nil
		})
		e.fnCanPlayThrough = &fn
	}

	e.selfElement.Call(
		"addEventListener",
		"canplaythrough",
		*e.fnCanPlayThrough,
	)
	return e
}

func (e *TagVideo) RemoveListenerCanPlayThrough() (ref *TagVideo) {
	if e.fnCanPlayThrough == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"canplaythrough",
		*e.fnCanPlayThrough,
	)
	e.fnCanPlayThrough = nil
	return e
}

func (e *TagVideo) AddListenerComplete(evet *chan event.Data) (ref *TagVideo) {
	var fn js.Func

	if e.fnComplete == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if len(args) == 0 {
				return nil
			}

			*evet <- event.EventManager(event.KEventComplete, this, args)
			return nil
		})
		e.fnComplete = &fn
	}

	e.selfElement.Call(
		"addEventListener",
		"complete",
		*e.fnComplete,
	)
	return e
}

func (e *TagVideo) RemoveListenerComplete() (ref *TagVideo) {
	if e.fnComplete == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"complete",
		*e.fnComplete,
	)
	e.fnComplete = nil
	return e
}

func (e *TagVideo) AddListenerDurationChange(evet *chan event.Data) (ref *TagVideo) {
	var fn js.Func

	if e.fnDurationChange == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if len(args) == 0 {
				return nil
			}

			*evet <- event.EventManager(event.KEventDurationChange, this, args)
			return nil
		})
		e.fnDurationChange = &fn
	}

	e.selfElement.Call(
		"addEventListener",
		"durationchange",
		*e.fnDurationChange,
	)
	return e
}

func (e *TagVideo) RemoveListenerDurationChange() (ref *TagVideo) {
	if e.fnDurationChange == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"durationchange",
		*e.fnDurationChange,
	)
	e.fnDurationChange = nil
	return e
}

func (e *TagVideo) AddListenerEmptied(evet *chan event.Data) (ref *TagVideo) {
	var fn js.Func

	if e.fnEmptied == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if len(args) == 0 {
				return nil
			}

			*evet <- event.EventManager(event.KEventEmptied, this, args)
			return nil
		})
		e.fnEmptied = &fn
	}

	e.selfElement.Call(
		"addEventListener",
		"emptied",
		*e.fnEmptied,
	)
	return e
}

func (e *TagVideo) RemoveListenerEmptied() (ref *TagVideo) {
	if e.fnEmptied == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"emptied",
		*e.fnEmptied,
	)
	e.fnEmptied = nil
	return e
}

func (e *TagVideo) AddListenerEnded(evet *chan event.Data) (ref *TagVideo) {
	var fn js.Func

	if e.fnEnded == nil {
		fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if len(args) == 0 {
				return nil
			}

			*evet <- event.EventManager(event.KEventEnded, this, args)
			return nil
		})
		e.fnEnded = &fn
	}

	e.selfElement.Call(
		"addEventListener",
		"ended",
		*e.fnEnded,
	)
	return e
}

func (e *TagVideo) RemoveListenerEnded() (ref *TagVideo) {
	if e.fnEnded == nil {
		return e
	}

	e.selfElement.Call(
		"removeEventListener",
		"ended",
		*e.fnEnded,
	)
	e.fnEnded = nil
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
func (e *TagVideo) CreateElement() (ref *TagVideo) {
	e.selfElement = js.Global().Get("document").Call("createElement", "video")
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
func (e *TagVideo) AppendById(appendId string) (ref *TagVideo) {

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
func (e *TagVideo) Append(elements ...Compatible) (ref *TagVideo) {
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
func (e *TagVideo) AppendToStage() (ref *TagVideo) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

// SetXY
//
// English:
//
//	Sets the X and Y axes in pixels.
//
// Português:
//
//	Define os eixos X e Y em pixels.
func (e *TagVideo) SetXY(x, y int) (ref *TagVideo) {

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
func (e *TagVideo) SetDeltaX(delta int) (ref *TagVideo) {
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
func (e *TagVideo) SetDeltaY(delta int) (ref *TagVideo) {
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
func (e *TagVideo) SetX(x int) (ref *TagVideo) {

	// dragging does not move delta(x,y) as the dragging function uses the delta(x,y) of mouse click
	// dragging não move delta (x,y) pois a função dragging usa o delta (x,y) do click do mouse
	if e.isDragging == true {
		e.x = x
	} else {
		e.x = x + e.deltaMovieX
	}

	px := strconv.FormatInt(int64(e.x), 10) + "px"
	e.selfElement.Get("style").Set("left", px)

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
func (e *TagVideo) SetY(y int) (ref *TagVideo) {

	// dragging does not move delta(x,y) as the dragging function uses the delta(x,y) of mouse click
	// dragging não move delta (x,y) pois a função dragging usa o delta (x,y) do click do mouse
	if e.isDragging == true {
		e.y = y
	} else {
		e.y = y + e.deltaMovieY
	}

	py := strconv.FormatInt(int64(e.y), 10) + "px"
	e.selfElement.Get("style").Set("top", py)

	return e
}

func (e *TagVideo) Get() (el js.Value) {
	return e.selfElement
}

// GetXY
//
// English:
//
//	Returns the X and Y axes in pixels.
//
// Português:
//
//	Retorna os eixos X e Y em pixels.
func (e *TagVideo) GetXY() (x, y int) {
	x = e.GetX()
	y = e.GetY()

	return
}

// GetX
//
// English:
//
//	Returns the X axe in pixels.
//
// Português:
//
//	Retorna o eixo X em pixels.
func (e *TagVideo) GetX() (x int) {
	//rect.top, rect.right, rect.bottom, rect.left
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	x = coordinate.Get("left").Int()
	return
}

// GetY
//
// English:
//
//	Returns the Y axe in pixels.
//
// Português:
//
//	Retorna o eixo Y em pixels.
func (e *TagVideo) GetY() (y int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	y = coordinate.Get("top").Int()
	return
}

// GetTop
//
// English:
//
//	Same as GetX() function, returns the x position of the element.
//
// Português:
//
//	O mesmo que a função GetX(), retorna a posição x do elemento.
func (e *TagVideo) GetTop() (top int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	top = coordinate.Get("top").Int()
	return
}

// GetRight
//
// English:
//
//	It is the same as x + width.
//
// Português:
//
//	É o mesmo que x + width.
func (e *TagVideo) GetRight() (right int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	right = coordinate.Get("right").Int()
	return
}

// GetBottom
//
// English:
//
//	It is the same as y + height.
//
// Português:
//
//	É o mesmo que y + Height.
func (e *TagVideo) GetBottom() (bottom int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	bottom = coordinate.Get("bottom").Int()
	return
}

// GetLeft
//
// English:
//
//	Same as GetY() function, returns the y position of the element.
//
// Português:
//
//	O mesmo que a função GetY(), retorna a posição y do elemento.
func (e *TagVideo) GetLeft() (left int) {
	var coordinate = e.selfElement.Call("getBoundingClientRect")
	left = coordinate.Get("left").Int()
	return
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
func (e *TagVideo) Rotate(angle float64) (ref *TagVideo) {
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
func (e *TagVideo) RotateDelta(delta float64) (ref *TagVideo) {
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
func (e *TagVideo) GetRotateDelta() (delta float64) {
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
//func (e *TagVideo) AddListener(eventType interface{}, manager mouse.SimpleManager) (ref *TagVideo) {
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
//func (e *TagVideo) RemoveListener(eventType interface{}) (ref *TagVideo) {
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
func (e *TagVideo) Mouse(value mouse.CursorType) (ref *TagVideo) {
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
func (e *TagVideo) Init() (ref *TagVideo) {
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
//	Prepares the stage reference at initialization.
//
// Português:
//
//	Prepara à referencia do stage na inicialização.
func (e *TagVideo) prepareStageReference() {
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
//     factoryBrowser.NewTagDiv("div_0").
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
//     factoryBrowser.NewTagDiv("div_0").
//       Class("animate").
//       DragStart().
//       AppendById("stage")
//func (e *TagVideo) DragStart() (ref *TagVideo) {
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
//     factoryBrowser.NewTagDiv("div_0").
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
//     factoryBrowser.NewTagDiv("div_0").
//       Class("animate").
//       DragStart().
//       AppendById("stage")
//
//     go func() {
//       time.Sleep(10 * time.Second)
//       div.DragStop()
//     }()
//func (e *TagVideo) DragStop() (ref *TagVideo) {
//	e.dragNormalStop()
//	return e
//}

//func (e *TagVideo) dragNormalStart() {
//	e.AddListener(mouse.KEventMouseDown, e.onStartDragNormal)
//	e.stage.Call("addEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
//	e.stage.Call("addEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
//}
//
//func (e *TagVideo) dragNormalStop() {
//	e.RemoveListener(mouse.KEventMouseDown)
//	e.stage.Call("removeEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
//	e.stage.Call("removeEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
//	e.isDragging = false
//}

func (e *TagVideo) onStopDragNormal(_ js.Value, _ []js.Value) interface{} {
	e.isDragging = false
	return nil
}

//func (e *TagVideo) onStartDragNormal(event mouse.MouseEvent) {
//	var screenX = int(event.GetScreenX())
//	var screenY = int(event.GetScreenY())
//
//	e.dragDifX = screenX - e.x
//	e.dragDifY = screenY - e.y
//
//	e.isDragging = true
//}

//func (e *TagVideo) onMouseDraggingNormal(_ js.Value, args []js.Value) interface{} {
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
func (e *TagVideo) AddPointsToEasingTween(algorithmRef algorithm.CurveInterface) (ref *TagVideo) {
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
func (e *TagVideo) EasingTweenWalkingIntoPoints() (function func(percent, p float64, args interface{})) {

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
func (e *TagVideo) EasingTweenWalkingAndRotateIntoPoints() (function func(forTenThousand, percent float64, args interface{})) {

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
	}

	return
}
