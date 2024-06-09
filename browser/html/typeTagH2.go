package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
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

// TagH2
//
// English:
//
// The <h1> to <h6> HTML elements represent six levels of section headings. <h1> is the highest
// section level and <h6> is the lowest.
//
// Multiple <h1> elements on one page
//
// Using more than one <h1> is allowed by the HTML specification, but is not considered a best
// practice. Using only one <h1> is beneficial for screenreader users.
//
// The HTML specification includes the concept of an outline formed by the use of <section> elements.
// If this were implemented it would enable the use of multiple <h1> elements, giving user
// agents—including screen readers—a way to understand that an <h1> nested inside a defined section is
// a subheading. This functionality has never been implemented; therefore it is important to use your
// headings to describe the outline of your document.
//
//	Notes:
//	  * Heading information can be used by user agents to construct a table of contents for a
//	    document automatically.
//	  * Avoid using heading elements to resize text. Instead, use the CSS font-size property.
//	  * Avoid skipping heading levels: always start from <h1>, followed by <h2> and so on.
//	  * Use only one <h1> per page or view. It should concisely describe the overall purpose of the
//	    content.
//	  * The align attribute is obsolete; don't use it.
//
// Português:
//
// Os elementos HTML <h1> a <h6> representam seis níveis de cabeçalho, onde, <h1> é o nível mais alto
// e <h6> o nível mais baixo.
//
// Múltiplos elementos <h1> em uma página
//
// O uso de mais de um <h1> é permitido pela especificação HTML, mas não é considerado uma prática
// recomendada. Usar apenas um <h1> é benéfico para usuários de leitores de tela.
//
// A especificação HTML inclui o conceito de contorno formado pelo uso de elementos <section>.
// Se isso fosse implementado, permitiria o uso de vários elementos <h1>, dando aos agentes do usuário
// – incluindo leitores de tela – uma maneira de entender que um <h1> aninhado dentro de uma seção
// definida é um subtítulo. Essa funcionalidade nunca foi implementada; portanto, é importante usar
// seus títulos para descrever o esboço do seu documento.
//
//	Notas:
//	  * As informações de cabeçalho podem ser usadas por agentes de usuário para construir
//	    automaticamente um índice para um documento.
//	  * Evite usar elementos de título para redimensionar o texto. Em vez disso, use a propriedade
//	    CSS font-size.
//	  * Evite pular níveis de título: sempre comece de <h1>, seguido de <h2> e assim por diante.
//	  * Use apenas um <h1> por página ou visualização. Deve descrever de forma concisa o propósito
//	    geral do conteúdo.
//	  * O atributo align está obsoleto; não o use.
type TagH2 struct {
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
func (e *TagH2) Reference(reference **TagH2) (ref *TagH2) {
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
func (e *TagH2) AccessKey(key string) (ref *TagH2) {
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
func (e *TagH2) Autofocus(autofocus bool) (ref *TagH2) {
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
func (e *TagH2) Class(class ...string) (ref *TagH2) {
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
func (e *TagH2) ContentEditable(editable bool) (ref *TagH2) {
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
func (e *TagH2) Data(data map[string]string) (ref *TagH2) {
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
func (e *TagH2) Dir(dir Dir) (ref *TagH2) {
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
func (e *TagH2) Draggable(draggable Draggable) (ref *TagH2) {
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
func (e *TagH2) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagH2) {
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
func (e *TagH2) Hidden() (ref *TagH2) {
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
func (e *TagH2) Id(id string) (ref *TagH2) {
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
func (e *TagH2) InputMode(inputMode InputMode) (ref *TagH2) {
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
func (e *TagH2) Is(is string) (ref *TagH2) {
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
func (e *TagH2) ItemId(id string) (ref *TagH2) {
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
func (e *TagH2) ItemProp(itemprop string) (ref *TagH2) {
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
func (e *TagH2) ItemRef(itemref string) (ref *TagH2) {
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
func (e *TagH2) ItemType(itemType string) (ref *TagH2) {
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
func (e *TagH2) Lang(language Language) (ref *TagH2) {
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
func (e *TagH2) Part(part ...string) (ref *TagH2) {
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
func (e *TagH2) Nonce(nonce string) (ref *TagH2) {
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
func (e *TagH2) Slot(slot string) (ref *TagH2) {
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
func (e *TagH2) Spellcheck(spell bool) (ref *TagH2) {
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
func (e *TagH2) Style(style string) (ref *TagH2) {
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
func (e *TagH2) TabIndex(index int) (ref *TagH2) {
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
func (e *TagH2) Title(title string) (ref *TagH2) {
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
func (e *TagH2) Translate(translate Translate) (ref *TagH2) {
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
func (e *TagH2) CreateElement() (ref *TagH2) {
	e.selfElement = js.Global().Get("document").Call("createElement", "h2")
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
func (e *TagH2) AppendById(appendId string) (ref *TagH2) {

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
func (e *TagH2) Append(elements ...Compatible) (ref *TagH2) {
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
func (e *TagH2) AppendToStage() (ref *TagH2) {
	e.stage.Call("appendChild", e.selfElement)
	return e
}

func (e *TagH2) Get() (el js.Value) {
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
func (e *TagH2) Rotate(angle float64) (ref *TagH2) {
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
func (e *TagH2) RotateDelta(delta float64) (ref *TagH2) {
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
func (e *TagH2) GetRotateDelta() (delta float64) {
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
//func (e *TagH2) AddListener(eventType interface{}, manager mouse.SimpleManager) (ref *TagH2) {
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
//func (e *TagH2) RemoveListener(eventType interface{}) (ref *TagH2) {
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
func (e *TagH2) Mouse(value mouse.CursorType) (ref *TagH2) {
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
func (e *TagH2) Init() (ref *TagH2) {
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
func (e *TagH2) prepareStageReference() {
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
//func (e *TagH2) DragStart() (ref *TagH2) {
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
//func (e *TagH2) DragStop() (ref *TagH2) {
//	e.dragNormalStop()
//	return e
//}

//func (e *TagH2) dragNormalStart() {
//	e.AddListener(mouse.KEventMouseDown, e.onStartDragNormal)
//	e.stage.Call("addEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
//	e.stage.Call("addEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
//}

//func (e *TagH2) dragNormalStop() {
//	e.RemoveListener(mouse.KEventMouseDown)
//	e.stage.Call("removeEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
//	e.stage.Call("removeEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
//	e.isDragging = false
//}

func (e *TagH2) onStopDragNormal(_ js.Value, _ []js.Value) interface{} {
	e.isDragging = false
	return nil
}

func (e *TagH2) onStartDragNormal(event mouse.Event) {
	var screenX = int(event.GetScreenX())
	var screenY = int(event.GetScreenY())

	e.dragDifX = screenX - e.x
	e.dragDifY = screenY - e.y

	e.isDragging = true
}

func (e *TagH2) onMouseDraggingNormal(_ js.Value, args []js.Value) interface{} {
	if e.isDragging == false {
		return nil
	}

	var mouseEvent = mouse.Event{}
	if len(args) > 0 {
		mouseEvent.Object = args[0]

		var x = int(mouseEvent.GetScreenX()) - e.dragDifX
		var y = int(mouseEvent.GetScreenY()) - e.dragDifY

		e.SetXY(x, y)
	}

	return nil
}

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
func (e *TagH2) AddPointsToEasingTween(algorithmRef algorithm.CurveInterface) (ref *TagH2) {
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
func (e *TagH2) EasingTweenWalkingIntoPoints() (function func(percent, p float64, args interface{})) {

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
func (e *TagH2) EasingTweenWalkingAndRotateIntoPoints() (function func(forTenThousand, percent float64, args interface{})) {

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

// Text
//
// English:
//
// Adds plain text to the tag's content.
//
// Text:
//
// Adiciona um texto simples ao conteúdo da tag.
func (e *TagH2) Text(value string) (ref *TagH2) {
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
func (e *TagH2) Html(value string) (ref *TagH2) {
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
func (e *TagH2) GetXY() (x, y int) {
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
func (e *TagH2) GetX() (x int) {
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
func (e *TagH2) GetY() (y int) {
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
func (e *TagH2) GetTop() (top int) {
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
func (e *TagH2) GetRight() (right int) {
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
func (e *TagH2) GetBottom() (bottom int) {
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
func (e *TagH2) GetLeft() (left int) {
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
func (e *TagH2) GetBoundingBox() (x, y, width, height int) {
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
func (e *TagH2) CollisionBoundingBox(elemnt CollisionBoundingBox) (collision bool) {
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
func (e *TagH2) UpdateBoundingClientRect() (ref *TagH2) {
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
func (e *TagH2) SetXY(x, y int) (ref *TagH2) {

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
func (e *TagH2) SetDeltaX(delta int) (ref *TagH2) {
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
func (e *TagH2) SetDeltaY(delta int) (ref *TagH2) {
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
func (e *TagH2) SetX(x int) (ref *TagH2) {

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
func (e *TagH2) SetY(y int) (ref *TagH2) {

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

func (e *TagH2) AddListenerAbort(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerAbort() (ref *TagH2) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagH2) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerAuxclick() (ref *TagH2) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagH2) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerBeforeinput() (ref *TagH2) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagH2) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerBeforematch() (ref *TagH2) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagH2) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerBeforetoggle() (ref *TagH2) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagH2) AddListenerCancel(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerCancel() (ref *TagH2) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagH2) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerCanplay() (ref *TagH2) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagH2) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerCanplaythrough() (ref *TagH2) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagH2) AddListenerChange(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerChange() (ref *TagH2) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagH2) AddListenerClick(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerClick() (ref *TagH2) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagH2) AddListenerClose(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerClose() (ref *TagH2) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagH2) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerContextlost() (ref *TagH2) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagH2) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerContextmenu() (ref *TagH2) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagH2) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerContextrestored() (ref *TagH2) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagH2) AddListenerCopy(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerCopy() (ref *TagH2) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagH2) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerCuechange() (ref *TagH2) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagH2) AddListenerCut(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerCut() (ref *TagH2) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagH2) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerDblclick() (ref *TagH2) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagH2) AddListenerDrag(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerDrag() (ref *TagH2) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagH2) AddListenerDragend(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerDragend() (ref *TagH2) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagH2) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerDragenter() (ref *TagH2) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagH2) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerDragleave() (ref *TagH2) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagH2) AddListenerDragover(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerDragover() (ref *TagH2) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagH2) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerDragstart() (ref *TagH2) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagH2) AddListenerDrop(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerDrop() (ref *TagH2) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagH2) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerDurationchange() (ref *TagH2) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagH2) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerEmptied() (ref *TagH2) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagH2) AddListenerEnded(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerEnded() (ref *TagH2) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagH2) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerFormdata() (ref *TagH2) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagH2) AddListenerInput(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerInput() (ref *TagH2) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagH2) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerInvalid() (ref *TagH2) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagH2) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerKeydown() (ref *TagH2) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagH2) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerKeypress() (ref *TagH2) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagH2) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerKeyup() (ref *TagH2) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagH2) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerLoadeddata() (ref *TagH2) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagH2) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerLoadedmetadata() (ref *TagH2) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagH2) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerLoadstart() (ref *TagH2) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagH2) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerMousedown() (ref *TagH2) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagH2) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerMouseenter() (ref *TagH2) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagH2) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerMouseleave() (ref *TagH2) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagH2) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerMousemove() (ref *TagH2) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagH2) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerMouseout() (ref *TagH2) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagH2) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerMouseover() (ref *TagH2) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagH2) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerMouseup() (ref *TagH2) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagH2) AddListenerPaste(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerPaste() (ref *TagH2) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagH2) AddListenerPause(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerPause() (ref *TagH2) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagH2) AddListenerPlay(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerPlay() (ref *TagH2) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagH2) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerPlaying() (ref *TagH2) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagH2) AddListenerProgress(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerProgress() (ref *TagH2) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagH2) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerRatechange() (ref *TagH2) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagH2) AddListenerReset(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerReset() (ref *TagH2) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagH2) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerScrollend() (ref *TagH2) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagH2) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerSecuritypolicyviolation() (ref *TagH2) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagH2) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerSeeked() (ref *TagH2) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagH2) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerSeeking() (ref *TagH2) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagH2) AddListenerSelect(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerSelect() (ref *TagH2) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagH2) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerSlotchange() (ref *TagH2) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagH2) AddListenerStalled(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerStalled() (ref *TagH2) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagH2) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerSubmit() (ref *TagH2) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagH2) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerSuspend() (ref *TagH2) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagH2) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerTimeupdate() (ref *TagH2) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagH2) AddListenerToggle(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerToggle() (ref *TagH2) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagH2) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerVolumechange() (ref *TagH2) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagH2) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerWaiting() (ref *TagH2) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagH2) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerWebkitanimationend() (ref *TagH2) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagH2) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerWebkitanimationiteration() (ref *TagH2) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagH2) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerWebkitanimationstart() (ref *TagH2) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagH2) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerWebkittransitionend() (ref *TagH2) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagH2) AddListenerWheel(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerWheel() (ref *TagH2) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagH2) AddListenerBlur(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerBlur() (ref *TagH2) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagH2) AddListenerError(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerError() (ref *TagH2) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagH2) AddListenerFocus(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerFocus() (ref *TagH2) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagH2) AddListenerLoad(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerLoad() (ref *TagH2) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagH2) AddListenerResize(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerResize() (ref *TagH2) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagH2) AddListenerScroll(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerScroll() (ref *TagH2) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagH2) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerAfterprint() (ref *TagH2) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagH2) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerBeforeprint() (ref *TagH2) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagH2) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerBeforeunload() (ref *TagH2) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagH2) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerHashchange() (ref *TagH2) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagH2) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerLanguagechange() (ref *TagH2) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagH2) AddListenerMessage(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerMessage() (ref *TagH2) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagH2) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerMessageerror() (ref *TagH2) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagH2) AddListenerOffline(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerOffline() (ref *TagH2) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagH2) AddListenerOnline(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerOnline() (ref *TagH2) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagH2) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerPageswap() (ref *TagH2) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagH2) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerPagehide() (ref *TagH2) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagH2) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerPagereveal() (ref *TagH2) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagH2) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerPageshow() (ref *TagH2) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagH2) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerPopstate() (ref *TagH2) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagH2) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerRejectionhandled() (ref *TagH2) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagH2) AddListenerStorage(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerStorage() (ref *TagH2) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagH2) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerUnhandledrejection() (ref *TagH2) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagH2) AddListenerUnload(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerUnload() (ref *TagH2) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagH2) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerReadystatechange() (ref *TagH2) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagH2) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagH2) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagH2) RemoveListenerVisibilitychange() (ref *TagH2) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
