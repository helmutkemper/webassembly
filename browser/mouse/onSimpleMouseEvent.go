package mouse

import (
	"syscall/js"
)

//https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent

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
}

// Button
//
// English:
//
// Archive the pressed mouse button, e.g. KMouseButtonMain to the main button.
//
// Português:
//
// Arquiva o botão do mouse pressionado, por exeplo, KMouseButtonMain para o botão principal.
type Button int

const (
	// KMouseButtonNoButton
	//
	// English:
	//
	//  No button or un-initialized.
	//
	// Português:
	//
	//  Sem botão ou não inicializado.
	KMouseButtonNoButton Button = 0

	// KMouseButtonMain
	//
	// English:
	//
	//  Main button pressed, usually the left button or the un-initialized state.
	//
	// Português:
	//
	//  Botão principal pressionado, geralmente o botão esquerdo ou o estado não inicializado.
	KMouseButtonMain Button = 1

	// KMouseButtonAuxiliary
	//
	// English:
	//
	//  Auxiliary button pressed, usually the wheel button or the middle button (if present).
	//
	// Português:
	//
	//  Botão auxiliar pressionado, geralmente o botão da roda ou o botão do meio (se houver).
	KMouseButtonAuxiliary Button = 2

	// KMouseButtonSecondary
	//
	// English:
	//
	//  Secondary button pressed, usually the right button.
	//
	// Português:
	//
	//  Botão secundário pressionado, geralmente o botão direito.
	KMouseButtonSecondary Button = 4

	// KMouseButtonFourth
	//
	// English:
	//
	//  Fourth button, typically the Browser Back button.
	//
	// Português:
	//
	//  Quarto botão, normalmente o botão Voltar do navegador.
	KMouseButtonFourth Button = 8

	// KMouseButtonFifth
	//
	// English:
	//
	//  Fifth button, typically the Browser Forward button.
	//
	// Português:
	//
	//  Quinto botão, normalmente o botão Browser Forward.
	KMouseButtonFifth Button = 16
)

type Event struct {
	Object js.Value
}

// GetClientX
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
func (e Event) GetClientX() (clientX float64) {
	clientX = e.Object.Get("clientX").Float()
	return
}

// GetClientY
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
func (e Event) GetClientY() (clientY float64) {
	clientY = e.Object.Get("clientY").Float()
	return
}

// GetMovementX
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
func (e Event) GetMovementX() (movementX float64) {
	movementX = e.Object.Get("movementX").Float()
	return
}

// GetMovementY
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
func (e Event) GetMovementY() (movementY float64) {
	movementY = e.Object.Get("movementY").Float()
	return
}

// GetOffsetX
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
func (e Event) GetOffsetX() (offsetX float64) {
	offsetX = e.Object.Get("offsetX").Float()
	return
}

// GetOffsetY
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
func (e Event) GetOffsetY() (offsetY float64) {
	offsetY = e.Object.Get("offsetY").Float()
	return
}

// GetPageX
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
func (e Event) GetPageX() (pageX float64) {
	pageX = e.Object.Get("pageX").Float()
	return
}

// GetPageY
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
func (e Event) GetPageY() (pageY float64) {
	pageY = e.Object.Get("pageY").Float()
	return
}

// GetScreenX
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
func (e Event) GetScreenX() (screenX float64) {
	screenX = e.Object.Get("screenX").Float()
	return
}

// GetScreenY
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
func (e Event) GetScreenY() (screenY float64) {
	screenY = e.Object.Get("screenY").Float()
	return
}

// GetX
//
// English:
//
// The MouseEvent.x property is an alias for the MouseEvent.clientX property.
//
// Português:
//
// A propriedade MouseEvent.x é um alias para a propriedade MouseEvent.clientX.
func (e Event) GetX() (x float64) {
	x = e.Object.Get("x").Float()
	return
}

// GetY
//
// English:
//
// The MouseEvent.y property is an alias for the MouseEvent.clientY property.
//
// Português:
//
// A propriedade MouseEvent.y é um alias para a propriedade MouseEvent.clientY.
func (e Event) GetY() (y float64) {
	y = e.Object.Get("y").Float()
	return
}

// GetRelatedTarget
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
func (e Event) GetRelatedTarget() (object js.Value) {
	object = e.Object.Get("relatedTarget")
	return
}

// GetRegion
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
func (e Event) GetRegion() (region string) {
	region = e.Object.Get("region").String()
	return
}

// GetButton
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
func (e Event) GetButton() (mouseButton Button) {
	return Button(e.Object.Get("buttons").Int())
}

// GetAltKey
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
func (e Event) GetAltKey() (altKey bool) {
	return e.Object.Get("altKey").Bool()
}

// GetShiftKey
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
func (e Event) GetShiftKey() (shiftKey bool) {
	return e.Object.Get("shiftKey").Bool()
}

// GetMetaKey
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
func (e Event) GetMetaKey() (metaKey bool) {
	return e.Object.Get("metaKey").Bool()
}

// GetCtrlKey
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
func (e Event) GetCtrlKey() (ctrlKey bool) {
	return e.Object.Get("ctrlKey").Bool()
}

type SimpleManager func(event Event)

// EventManager
//
// English:
//
// Registers a Golang function to fire when a mouse event happens.
//
//   Input:
//     manager: Golang function like func().
//
// Português:
//
// Registra uma função Golang para ser disparada quando um evento do mouse acontece.
//
//   Entrada:
//     manager: função Golang tipo func().
func EventManager(_ js.Value, args []js.Value) (data Data) {
	var event = Event{}
	event.Object = args[0]

	data.ClientX = event.GetClientX()
	data.ClientY = event.GetClientY()
	data.MovementX = event.GetMovementX()
	data.MovementY = event.GetMovementY()
	data.OffsetX = event.GetOffsetX()
	data.OffsetY = event.GetOffsetY()
	data.PageX = event.GetPageX()
	data.PageY = event.GetPageY()
	data.ScreenX = event.GetScreenX()
	data.ScreenY = event.GetScreenY()
	data.X = event.GetX()
	data.Y = event.GetY()
	data.RelatedTarget = event.GetRelatedTarget()
	data.Region = event.GetRegion()
	data.Button = event.GetButton()
	data.AltKey = event.GetAltKey()
	data.ShiftKey = event.GetShiftKey()
	data.MetaKey = event.GetMetaKey()
	data.CtrlKey = event.GetCtrlKey()

	return
}
