package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
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

// TagH3
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
type TagH3 struct {

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
func (e *TagH3) Reference(reference **TagH3) (ref *TagH3) {
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
func (e *TagH3) AccessKey(key string) (ref *TagH3) {
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
func (e *TagH3) Autofocus(autofocus bool) (ref *TagH3) {
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
func (e *TagH3) Class(class ...string) (ref *TagH3) {
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
func (e *TagH3) ContentEditable(editable bool) (ref *TagH3) {
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
func (e *TagH3) Data(data map[string]string) (ref *TagH3) {
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
func (e *TagH3) Dir(dir Dir) (ref *TagH3) {
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
func (e *TagH3) Draggable(draggable Draggable) (ref *TagH3) {
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
func (e *TagH3) EnterKeyHint(enterKeyHint EnterKeyHint) (ref *TagH3) {
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
func (e *TagH3) Hidden() (ref *TagH3) {
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
func (e *TagH3) Id(id string) (ref *TagH3) {
	e.id = id
	e.selfElement.Set("id", id)

	// Saves the element reference with ID for later use.
	// Salva a referência dos elementos com ID para uso posterior.
	htmlGlobalAllElementsList.Delete(id)
	htmlGlobalAllElementsList.Store(id, e)
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
func (e *TagH3) InputMode(inputMode InputMode) (ref *TagH3) {
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
func (e *TagH3) Is(is string) (ref *TagH3) {
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
func (e *TagH3) ItemId(id string) (ref *TagH3) {
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
func (e *TagH3) ItemProp(itemprop string) (ref *TagH3) {
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
func (e *TagH3) ItemRef(itemref string) (ref *TagH3) {
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
func (e *TagH3) ItemType(itemType string) (ref *TagH3) {
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
func (e *TagH3) Lang(language Language) (ref *TagH3) {
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
func (e *TagH3) Part(part ...string) (ref *TagH3) {
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
func (e *TagH3) Nonce(nonce string) (ref *TagH3) {
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
func (e *TagH3) Slot(slot string) (ref *TagH3) {
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
func (e *TagH3) Spellcheck(spell bool) (ref *TagH3) {
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
func (e *TagH3) Style(style string) (ref *TagH3) {
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
func (e *TagH3) TabIndex(index int) (ref *TagH3) {
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
func (e *TagH3) Title(title string) (ref *TagH3) {
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
func (e *TagH3) Translate(translate Translate) (ref *TagH3) {
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
func (e *TagH3) CreateElement() (ref *TagH3) {
	e.selfElement = js.Global().Get("document").Call("createElement", "h3")
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
func (e *TagH3) AppendById(appendId string) (ref *TagH3) {

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
func (e *TagH3) Append(elements ...Compatible) (ref *TagH3) {
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
func (e *TagH3) AppendToStage() (ref *TagH3) {
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
func (e *TagH3) SetXY(x, y int) (ref *TagH3) {

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
func (e *TagH3) SetDeltaX(delta int) (ref *TagH3) {
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
func (e *TagH3) SetDeltaY(delta int) (ref *TagH3) {
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
func (e *TagH3) SetX(x int) (ref *TagH3) {

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
func (e *TagH3) SetY(y int) (ref *TagH3) {

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

func (e *TagH3) Get() (el js.Value) {
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
func (e *TagH3) GetXY() (x, y int) {
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
func (e *TagH3) GetX() (x int) {
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
func (e *TagH3) GetY() (y int) {
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
func (e *TagH3) GetTop() (top int) {
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
func (e *TagH3) GetRight() (right int) {
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
func (e *TagH3) GetBottom() (bottom int) {
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
func (e *TagH3) GetLeft() (left int) {
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
func (e *TagH3) Rotate(angle float64) (ref *TagH3) {
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
func (e *TagH3) RotateDelta(delta float64) (ref *TagH3) {
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
func (e *TagH3) GetRotateDelta() (delta float64) {
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
//func (e *TagH3) AddListener(eventType interface{}, manager mouse.SimpleManager) (ref *TagH3) {
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
//func (e *TagH3) RemoveListener(eventType interface{}) (ref *TagH3) {
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
func (e *TagH3) Mouse(value mouse.CursorType) (ref *TagH3) {
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
func (e *TagH3) Init() (ref *TagH3) {
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
func (e *TagH3) prepareStageReference() {
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
//func (e *TagH3) DragStart() (ref *TagH3) {
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
//func (e *TagH3) DragStop() (ref *TagH3) {
//	e.dragNormalStop()
//	return e
//}

//func (e *TagH3) dragNormalStart() {
//	e.AddListener(mouse.KEventMouseDown, e.onStartDragNormal)
//	e.stage.Call("addEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
//	e.stage.Call("addEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
//}

//func (e *TagH3) dragNormalStop() {
//	e.RemoveListener(mouse.KEventMouseDown)
//	e.stage.Call("removeEventListener", mouse.KEventMouseUp.String(), js.FuncOf(e.onStopDragNormal))
//	e.stage.Call("removeEventListener", mouse.KEventMouseMove.String(), js.FuncOf(e.onMouseDraggingNormal))
//	e.isDragging = false
//}

func (e *TagH3) onStopDragNormal(_ js.Value, _ []js.Value) interface{} {
	e.isDragging = false
	return nil
}

func (e *TagH3) onStartDragNormal(event mouse.Event) {
	var screenX = int(event.GetScreenX())
	var screenY = int(event.GetScreenY())

	e.dragDifX = screenX - e.x
	e.dragDifY = screenY - e.y

	e.isDragging = true
}

func (e *TagH3) onMouseDraggingNormal(_ js.Value, args []js.Value) interface{} {
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
func (e *TagH3) AddPointsToEasingTween(algorithmRef algorithm.CurveInterface) (ref *TagH3) {
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
func (e *TagH3) EasingTweenWalkingIntoPoints() (function func(percent, p float64, args interface{})) {

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
func (e *TagH3) EasingTweenWalkingAndRotateIntoPoints() (function func(forTenThousand, percent float64, args interface{})) {

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
func (e *TagH3) Text(value string) (ref *TagH3) {
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
func (e *TagH3) Html(value string) (ref *TagH3) {
	e.selfElement.Set("innerHTML", value)
	return e
}
