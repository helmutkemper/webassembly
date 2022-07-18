package mouse

import "syscall/js"

// Data
//
// English:
//
// Archives the values collected during the mouse event.
//
// Português:
//
// Arquiva os valores coletados durante o evento do mouse.
type Data struct {
	// ClientX
	//
	// English:
	//
	// The clientX property of the MouseEvent interface provides the horizontal coordinate within the
	// application's viewport at which the event occurred (as opposed to the coordinate within the page).
	//
	// For example, clicking on the left edge of the viewport will always result in a mouse event with a
	// clientX value of 0, regardless of whether the page is scrolled horizontally.
	//
	// Português:
	//
	// A propriedade clientX da interface MouseEvent fornece a coordenada horizontal na janela de
	// visualização do aplicativo na qual o evento ocorreu (em oposição à coordenada na página).
	//
	// Por exemplo, clicar na borda esquerda da janela de visualização sempre resultará em um evento de
	// mouse com um valor clientX de 0, independentemente de a página ser rolada horizontalmente.
	ClientX float64

	// ClientY
	//
	// English:
	//
	// The clientY read-only property of the MouseEvent interface provides the vertical coordinate within
	// the application's viewport at which the event occurred (as opposed to the coordinate within the
	// page).
	//
	// For example, clicking on the top edge of the viewport will always result in a mouse event with a
	// clientY value of 0, regardless of whether the page is scrolled vertically.
	//
	// Português:
	//
	// A propriedade somente leitura clientY da interface MouseEvent fornece a coordenada vertical na
	// janela de visualização do aplicativo na qual o evento ocorreu (em oposição à coordenada na
	// página).
	//
	// Por exemplo, clicar na borda superior da janela de visualização sempre resultará em um evento de
	// mouse com um valor clientY de 0, independentemente de a página ser rolada verticalmente.
	ClientY float64

	// MovementX
	//
	// English:
	//
	// The movementX read-only property of the MouseEvent interface provides the difference in the X
	// coordinate of the mouse pointer between the given event and the previous mousemove event.
	// In other words, the value of the property is computed like this:
	// currentEvent.movementX = currentEvent.screenX - previousEvent.screenX.
	//
	// Português:
	//
	// A propriedade somente leitura movementX da interface MouseEvent fornece a diferença na coordenada
	// X do ponteiro do mouse entre o evento fornecido e o evento mousemove anterior. Em outras palavras,
	// o valor da propriedade é calculado assim:
	// currentEvent.movementX = currentEvent.screenX - previousEvent.screenX.
	MovementX float64

	// MovementY
	//
	// English:
	//
	// The movementY read-only property of the MouseEvent interface provides the difference in the Y
	// coordinate of the mouse pointer between the given event and the previous mousemove event.
	// In other words, the value of the property is computed like this:
	// currentEvent.movementY = currentEvent.screenY - previousEvent.screenY.
	//
	// Português:
	//
	// A propriedade somente leitura movementY da interface MouseEvent fornece a diferença na coordenada
	// Y do ponteiro do mouse entre o evento fornecido e o evento mousemove anterior.
	// Em outras palavras, o valor da propriedade é calculado assim:
	// currentEvent.movementY = currentEvent.screenY - previousEvent.screenY.
	MovementY float64

	// OffsetX
	//
	// English:
	//
	// The offsetX read-only property of the MouseEvent interface provides the offset in the X coordinate
	// of the mouse pointer between that event and the padding edge of the target node.
	//
	// Português:
	//
	// A propriedade somente leitura offsetX da interface MouseEvent fornece o deslocamento na coordenada
	// X do ponteiro do mouse entre esse evento e a borda de preenchimento do nó de destino.
	OffsetX float64

	// OffsetY
	//
	// English:
	//
	// The offsetY read-only property of the MouseEvent interface provides the offset in the Y coordinate
	// of the mouse pointer between that event and the padding edge of the target node.
	//
	// Português:
	//
	// A propriedade somente leitura offsetY da interface MouseEvent fornece o deslocamento na coordenada
	// Y do ponteiro do mouse entre esse evento e a borda de preenchimento do nó de destino.
	OffsetY float64

	// PageX
	//
	// English:
	//
	// The pageX read-only property of the MouseEvent interface returns the X (horizontal) coordinate
	// (in pixels) at which the mouse was clicked, relative to the left edge of the entire document.
	// This includes any portion of the document not currently visible.
	//
	//   Output:
	//     pageX: A floating-point number of pixels from the left edge of the document at which the mouse
	//     was clicked, regardless of any scrolling or viewport positioning that may be in effect.
	//
	// Being based on the edge of the document as it is, this property takes into account any horizontal
	// scrolling of the page. For example, if the page is scrolled such that 200 pixels of the left side
	// of the document are scrolled out of view, and the mouse is clicked 100 pixels inward from the left
	// edge of the view, the value returned by pageX will be 300.
	//
	// Originally, this property was defined as a long integer. The CSSOM View Module redefined it as a
	// double float. See the Browser compatibility section for details.
	//
	// See Page in Coordinate systems for some additional information about coordinates specified in this
	// fashion.
	//
	// This property was originally specified in the Touch Events specification as a long integer, but
	// was redefined in the CSSOM View Module to be a double-precision floating-point number to allow
	// for subpixel precision. Even though numeric types both are represented by Number in JavaScript,
	// they may be handled differently internally in the browser's code, resulting in potential
	// behavior differences.
	//
	// See Browser compatibility to learn which browsers have been updated to use the revised data
	// type.
	//
	// Português:
	//
	// A propriedade somente leitura pageX da interface MouseEvent retorna a coordenada X (horizontal)
	// (em pixels) na qual o mouse foi clicado, em relação à borda esquerda do documento inteiro. Isso
	// inclui qualquer parte do documento não visível no momento.
	//
	//   Saída:
	//     pageX: Um número de pixels de ponto flutuante da borda esquerda do documento em que o mouse foi
	//     clicado, independentemente de qualquer rolagem ou posicionamento da janela de visualização que
	//     possa estar em vigor.
	//
	// Baseando-se na borda do documento como está, essa propriedade leva em consideração qualquer rolagem
	// horizontal da página. Por exemplo, se a página for rolada de forma que 200 pixels do lado esquerdo
	// do documento sejam rolados para fora da visualização e o mouse for clicado 100 pixels para dentro
	// da borda esquerda da visualização, o valor retornado por pageX será 300.
	//
	// Originalmente, essa propriedade foi definida como um inteiro longo. O CSSOM View Module o redefiniu
	// como um float duplo. Consulte a seção Compatibilidade do navegador para obter detalhes.
	//
	// Consulte Página em sistemas de coordenadas para obter algumas informações adicionais sobre as
	// coordenadas especificadas desta forma.
	//
	// Essa propriedade foi originalmente especificada na especificação Touch Events como um inteiro
	// longo, mas foi redefinida no CSSOM View Module para ser um número de ponto flutuante de
	// precisão dupla para permitir a precisão de subpixel. Embora ambos os tipos numéricos sejam
	// representados por Number em JavaScript, eles podem ser tratados de forma diferente internamente
	// no código do navegador, resultando em possíveis diferenças de comportamento.
	//
	// Consulte Compatibilidade do navegador para saber quais navegadores foram atualizados para usar
	// o tipo de dados revisado.
	PageX float64

	// PageY
	//
	// English:
	//
	// The pageY read-only property of the MouseEvent interface returns the Y (vertical) coordinate in
	// pixels of the event relative to the whole document. This property takes into account any vertical
	// scrolling of the page.
	//
	//   Output:
	//     pageY: A double floating point value.
	//
	// Português:
	//
	// A propriedade somente leitura pageY da interface MouseEvent retorna a coordenada Y (vertical) em
	// pixels do evento em relação do documento inteiro.
	//
	// Esta propriedade leva em consideração qualquer rolagem vertical da página.
	//
	//   Saída:
	//     pageY: Um valor de ponto flutuante duplo.
	PageY float64

	// ScreenX
	//
	// English:
	//
	// The screenX read-only property of the MouseEvent interface provides the horizontal coordinate
	// (offset) of the mouse pointer in global (screen) coordinates.
	//
	//   Note:
	//     * In a multiscreen environment, screens aligned horizontally will be treated as a single
	//       device, and so the range of the screenX value will increase to the combined width of the
	//       screens.
	//
	// Português:
	//
	// A propriedade somente leitura screenX da interface MouseEvent fornece a coordenada horizontal
	// (deslocamento) do ponteiro do mouse em coordenadas globais (tela).
	//
	//   Nota:
	//     * Em um ambiente de várias telas, as telas alinhadas horizontalmente serão tratadas como um
	//       único dispositivo e, portanto, o intervalo do valor screenX aumentará para a largura
	//       combinada das telas.
	ScreenX float64

	// ScreenY
	//
	// English:
	//
	// The screenY read-only property of the MouseEvent interface provides the vertical coordinate
	// (offset) of the mouse pointer in global (screen) coordinates.
	//
	// Português:
	//
	// A propriedade somente leitura screenY da interface MouseEvent fornece a coordenada vertical
	// (deslocamento) do ponteiro do mouse em coordenadas globais (tela).
	ScreenY float64

	// X
	//
	// English:
	//
	// The MouseEvent.x property is an alias for the MouseEvent.clientX property.
	//
	// Português:
	//
	// A propriedade MouseEvent.x é um alias para a propriedade MouseEvent.clientX.
	X float64

	// Y
	//
	// English:
	//
	// The MouseEvent.y property is an alias for the MouseEvent.clientY property.
	//
	// Português:
	//
	// A propriedade MouseEvent.y é um alias para a propriedade MouseEvent.clientY.
	Y float64

	// RelatedTarget
	//
	// English:
	//
	// The MouseEvent.relatedTarget read-only property is the secondary target for the mouse event,
	// if there is one.
	//
	// In practice, it returns the html element of the target. Like, if the target is a div,
	// it returns the div with all its attributes.
	//
	//   Example:
	//     if RelatedTarget.IsNull() == false {
	//       log.Print("id: ", target.Get("id"))
	//     }
	//
	// Português:
	//
	// A propriedade somente leitura MouseEvent.relatedTarget é o destino secundário para o evento
	// de mouse, se houver.
	//
	// Na prática, ele retorna o elemento html do destino. Tipo, se o alvo for um div, ele retorna o
	// div com todos os seus atributos.
	//
	//   Exemplo:
	//     if RelatedTarget.IsNull() == false {
	//       log.Print("id: ", target.Get("id"))
	//     }
	RelatedTarget js.Value

	// Region
	//
	// English:
	//
	// The MouseEvent.region read-only property returns the id of the canvas hit region affected by the
	// event. If no hit region is affected, null is returned.
	//
	//   Output:
	//     region: A DOMString representing the id of the hit region.
	//
	// Português:
	//
	// A propriedade somente leitura MouseEvent.region retorna o id da região de acerto da tela afetada
	// pelo evento. Se nenhuma região de hit for afetada, null será retornado.
	//
	//   Saída:
	//     region: Um DOMString que representa o id da região do hit.
	Region string

	// Button
	//
	// English:
	//
	// Indicates which buttons are pressed on the mouse (or other input device) when a mouse event is
	// triggered.
	//
	// Português:
	//
	// Indica quais botões são pressionados no mouse (ou outro dispositivo de entrada) quando um evento
	// de mouse é acionado.
	Button Button

	// AltKey
	//
	// English:
	//
	// is a boolean value that indicates whether the alt key was pressed or not when a given mouse event
	// occurs.
	//
	// Be aware that the browser can't always detect the alt key on some operating systems.
	//
	// On some Linux variants, for example, a left mouse click combined with the alt key is used to move
	// or resize windows.
	//
	//   Note:
	//     * On Macintosh keyboards, this key is also known as the option key.
	//
	// Português:
	//
	// É um valor booleano que indica se a tecla alt foi pressionada ou não quando ocorre um determinado
	// evento de mouse.
	//
	// Esteja ciente de que o navegador nem sempre pode detectar a tecla alt em alguns sistemas
	// operacionais.
	//
	// Em algumas variantes do Linux, por exemplo, um clique esquerdo do mouse combinado com a tecla alt
	// é usado para mover ou redimensionar janelas.
	//
	//   Nota:
	//     * Em teclados Macintosh, essa tecla também é conhecida como tecla de opção.
	AltKey bool

	// ShiftKey
	//
	// English:
	//
	// The MouseEvent.shiftKey read-only property is a boolean value that indicates whether the shift key
	// was pressed or not when a given mouse event occurs.
	//
	// Português:
	//
	// A propriedade somente leitura MouseEvent.shiftKey é um valor booleano que indica se a tecla shift
	// foi pressionada ou não quando ocorre um determinado evento de mouse.
	ShiftKey bool

	// MetaKey
	//
	// English:
	//
	// Is a boolean value that indicates whether the meta key was pressed or not when a given mouse event
	// occurs.
	//
	// Be aware that many operating systems bind special functionality to the meta key, so this property
	// may be false even when the key is actually pressed. On Windows, for example, this key may open the
	// Start menu.
	//
	//   Note:
	//     * On Macintosh keyboards, this key is the command key (⌘). On Windows keyboards, this key is
	//       the Windows key (⊞).
	//
	// Português:
	//
	// É um valor booleano que indica se a tecla meta foi pressionada ou não quando ocorre um determinado
	// evento de mouse.
	//
	// Esteja ciente de que muitos sistemas operacionais vinculam funcionalidades especiais à meta-chave,
	// portanto, essa propriedade pode ser falsa mesmo quando a tecla é realmente pressionada. No Windows,
	// por exemplo, essa tecla pode abrir o menu Iniciar.
	//
	//   Nota:
	//     * Em teclados Macintosh, esta tecla é a tecla de comando (⌘). Em teclados Windows, esta tecla
	//       é a tecla Windows (⊞).
	MetaKey bool

	// CtrlKey
	//
	// English:
	//
	// The MouseEvent.ctrlKey read-only property is a boolean value that indicates whether the ctrl key
	// was pressed or not when a given mouse event occurs.
	//
	//   Note:
	//     * On Macintosh keyboards, this key is the control key.
	//
	// Português:
	//
	// A propriedade somente leitura MouseEvent.ctrlKey é um valor booleano que indica se a tecla ctrl
	// foi pressionada ou não quando ocorre um determinado evento de mouse.
	//
	//   Nota:
	//     * Em teclados Macintosh, esta tecla é a tecla de controle.
	CtrlKey bool

	// This
	//
	// English:
	//
	// This is the equivalent property of JavaScript's 'this'.
	//
	// The way to use it is This.Get(property string name). E.g. chan.This.Get("id")
	//
	// Português:
	//
	// Esta é a propriedade equivalente ao 'this' do JavaScript.
	//
	// A forma de usar é This.Get(property string name). Ex. chan.This.Get("id")
	This js.Value
}
