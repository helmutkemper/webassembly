package document

import (
	"syscall/js"
)

//https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent

type Event struct {
	Object js.Value
}

// GetName
//
// English:
//
// Gets the name of the window's browsing context.
//
// Português:
//
// Obtém o nome do contexto de navegação da janela.
func (e Event) GetName() (name string) {
	return e.Object.Get("name").String()
}

// GetWidth
//
// English:
//
// The read-only Window property innerWidth returns the interior width of the window in pixels. This includes the
// width of the vertical scroll bar, if one is present.
//
// More precisely, innerWidth returns the width of the window's layout viewport. The interior height of the window—the
// height of the layout viewport—can be obtained from the innerHeight property.
//
// Português:
//
// A propriedade de janela somente leitura innerWidth retorna a largura interna da janela em pixels. Isso inclui a
// largura da barra de rolagem vertical, se houver.
//
// Mais precisamente, innerWidth retorna a largura da viewport de layout da janela. A altura interior da janela—a
// altura da viewport de layout—pode ser obtida da propriedade innerHeight.
func (e Event) GetWidth() (width float64) {
	return e.Object.Get("innerWidth").Float()
}

// GetHeight
//
// English:
//
// The read-only innerHeight property of the Window interface returns the interior height of the window in pixels,
// including the height of the horizontal scroll bar, if present.
//
// The value of innerHeight is taken from the height of the window's layout viewport. The width can be obtained using
// the innerWidth property.
//
// Português:
//
// A propriedade innerHeight somente leitura da interface Window retorna a altura interna da janela em pixels,
// incluindo a altura da barra de rolagem horizontal, se presente.
//
// O valor de innerHeight é obtido da altura da viewport de layout da janela. A largura pode ser obtida usando a
// propriedade innerWidth.
func (e Event) GetHeight() (height float64) {
	return e.Object.Get("innerHeight").Float()
}

// GetFrameLength
//
// English:
//
// Returns the number of frames (either <frame> or <iframe> elements) in the window.
//
// Português:
//
// Retorna o número de quadros (elementos <frame> ou <iframe>) na janela.
func (e Event) GetFrameLength() (length int) {
	return e.Object.Get("length").Int()
}

// GetClosed
//
// English:
//
// The Window.closed read-only property indicates whether the referenced window is closed or not
//
// Português:
//
// A propriedade somente leitura Window.closed indica se a janela referenciada está fechada ou não
func (e Event) GetClosed() (closed bool) {
	return e.Object.Get("closed").Bool()
}

// GetOuterHeight
//
// English:
//
// The Window.outerHeight read-only property returns the height in pixels of the whole browser window, including any
// sidebar, window chrome, and window-resizing borders/handles.
//
//   Notes:
//     * To change the size of a window, see window.resizeBy() and window.resizeTo().
//
// Português:
//
// A propriedade somente leitura Window.outerHeight retorna a altura em pixels de toda a janela do navegador, incluindo
// qualquer barra lateral, cromo de janela e alças de borda de redimensionamento de janela.
//
//   Notas:
//     * Para alterar o tamanho de uma janela, consulte window.resizeBy() e window.resizeTo().
func (e Event) GetOuterHeight() (outerHeight float64) {
	return e.Object.Get("outerHeight").Float()
}

// GetOuterWidth
//
// English:
//
// Window.outerWidth read-only property returns the width of the outside of the browser window. It represents the width
// of the whole browser window including sidebar (if expanded), window chrome and window resizing borders/handles.
//
//   Notes:
//    * To change the size of a window, see window.resizeBy() and window.resizeTo().
//
// Português:
//
// A propriedade somente leitura Window.outerWidth retorna a largura da parte externa da janela do navegador.
// Ele representa a largura de toda a janela do navegador, incluindo barra lateral (se expandida), cromo da janela e
// alças de bordas de redimensionamento de janela.
//
//   Notas:
//    * Para alterar o tamanho de uma janela, consulte window.resizeBy() e window.resizeTo().
func (e Event) GetOuterWidth() (outerWidth float64) {
	return e.Object.Get("outerWidth").Float()
}

// GetScrollX
//
// English:
//
// The scrollX property of the Window interface returns the number of pixels that the document is currently scrolled
// horizontally. This value is subpixel precise in modern browsers, meaning that it isn't necessarily a whole number.
// You can get the number of pixels the document is scrolled vertically from the scrollY property.
//
// In more technical terms, scrollX returns the X coordinate of the left edge of the current viewport. If there is no
// viewport, the returned value is 0.
//
// Português:
//
// A propriedade scrollX da interface Window retorna o número de pixels que o documento está atualmente rolado
// horizontalmente. Esse valor é preciso em subpixels em navegadores modernos, o que significa que não é necessariamente
// um número inteiro. Você pode obter o número de pixels em que o documento é rolado verticalmente na propriedade
// scrollY.
//
// Em termos mais técnicos, scrollX retorna a coordenada X da borda esquerda da viewport atual. Se não houver viewport,
// o valor retornado será 0.
func (e Event) GetScrollX() (scrollX float64) {
	return e.Object.Get("scrollX").Float()
}

// GetScrollY
//
// English:
//
// The scrollY property of the Window interface returns the number of pixels that the document is currently scrolled
// vertically.
//
// This value is subpixel precise in modern browsers, meaning that it isn't necessarily a whole number. You can get the
// number of pixels the document is scrolled horizontally from the scrollX property.
//
// In more technical terms, scrollY returns the Y coordinate of the top edge of the current viewport. If there is no
// viewport, the returned value is 0.
//
// Português:
//
// A propriedade scrollY da interface Window retorna o número de pixels que o documento está atualmente rolado
// verticalmente.
//
// Esse valor é preciso em subpixels em navegadores modernos, o que significa que não é necessariamente um número
// inteiro. Você pode obter o número de pixels em que o documento é rolado horizontalmente na propriedade scrollX.
//
// Em termos mais técnicos, scrollY retorna a coordenada Y da borda superior da viewport atual. Se não houver viewport,
// o valor retornado será 0.
func (e Event) GetScrollY() (scrollY float64) {
	return e.Object.Get("scrollY").Float()
}

// GetScreenX
//
// English:
//
// The Window.screenX read-only property returns the horizontal distance, in CSS pixels, of the left border of the
// user's browser viewport to the left side of the screen.
//
// Português:
//
// A propriedade somente leitura Window.screenX retorna a distância horizontal, em pixels CSS, da borda esquerda da
// janela de visualização do navegador do usuário para o lado esquerdo da tela.
func (e Event) GetScreenX() (screenX float64) {
	return e.Object.Get("screenX").Float()
}

// GetScreenY
//
// English:
//
// The screenY property returns the vertical distance, in CSS pixels, of the top border of the user's browser viewport
// to the top edge of the screen.
//
// Português:
//
// A propriedade screenY retorna a distância vertical, em pixels CSS, da borda superior da janela de visualização do
// navegador do usuário até a borda superior da tela.
func (e Event) GetScreenY() (screenY float64) {
	return e.Object.Get("screenY").Float()
}

// GetOpener
//
// English:
//
// The Window interface's opener property returns a reference to the window that opened the window, either with
// open(), or by navigating a link with a target attribute.
//
// In other words, if window A opens window B, B.opener returns A.
//
// If the opener is not on the same origin as the current page, functionality of the opener object is limited.
// For example, variables and functions on the window object are not accessible. However, navigation of the opener
// window is possible, which means that the opened page can open a URL in the original tab or window.
// In some cases, this makes phishing attacks possible, where a trusted page that is opened in the original window is
// replaced by a phishing page by the newly opened page.
//
// In the following cases, the browser does not populate window.opener, but leaves it null:
//  * The opener can be omitted by specifying rel=noopener on a link, or passing noopener in the windowFeatures
//    parameter.
//  * Windows opened because of links with a target of _blank don't get an opener, unless explicitly requested with
//    rel=opener.
//  * Having a Cross-Origin-Opener-Policy header with a value of same-origin prevents setting opener.
//    Since the new window is loaded in a different browsing context, it won't have a reference to the opening window.
//
// Português:
//
// A propriedade opener da interface Window retorna uma referência à janela que abriu a janela, seja com open(), ou
// navegando em um link com um atributo target.
//
// Em outras palavras, se a janela A abrir a janela B, B.opener retornará A.
//
// Se o abridor não estiver na mesma origem da página atual, a funcionalidade do objeto abridor será limitada.
// Por exemplo, variáveis e funções no objeto de janela não são acessíveis. No entanto, a navegação da janela de
// abertura é possível, o que significa que a página aberta pode abrir um URL na guia ou janela original.
// Em alguns casos, isso possibilita ataques de phishing, em que uma página confiável aberta na janela original é
// substituída por uma página de phishing pela página recém-aberta.
//
// Nos casos a seguir, o navegador não preenche window.opener, mas o deixa nulo:
//  * O opener pode ser omitido especificando rel=noopener em um link ou passando noopener no parâmetro
//    windowFeatures.
//  * O Windows aberto devido a links com um destino de _blank não obtém um abridor, a menos que solicitado
//    explicitamente com rel=opener.
//  * Ter um cabeçalho Cross-Opener-Policy com um valor de mesma origem impede a configuração do abridor.
//    Como a nova janela é carregada em um contexto de navegação diferente, ela não terá uma referência à janela de
//    abertura.
func (e Event) GetOpener() (opener js.Value) {
	return e.Object.Get("opener")
}

// GetParent
//
// English:
//
// The Window.parent property is a reference to the parent of the current window or subframe.
//
// If a window does not have a parent, its parent property is a reference to itself.
//
// When a window is loaded in an <iframe>, <object>, or <frame>, its parent is the window with the element embedding
// the window.
//
// Português:
//
// A propriedade Window.parent é uma referência ao pai da janela ou subquadro atual.
//
// Se uma janela não tiver um pai, sua propriedade pai será uma referência a si mesma.
//
// Quando uma janela é carregada em um <iframe>, <object> ou <frame>, seu pai é a janela com o elemento incorporado
// à janela.
func (e Event) GetParent() (parent js.Value) {
	return e.Object.Get("parent")
}

// GetScreen
//
// English:
//
// The Window property screen returns a reference to the screen object associated with the window. The screen object,
// implementing the Screen interface, is a special object for inspecting properties of the screen on which the current
// window is being rendered.
//
// Português:
//
// A tela de propriedades da janela retorna uma referência ao objeto de tela associado à janela. O objeto de tela,
// implementando a interface Tela, é um objeto especial para inspecionar as propriedades da tela na qual a janela
// atual está sendo renderizada.
func (e Event) GetScreen() (parent js.Value) {
	return e.Object.Get("screen")
}

// GetScrollBars
//
// English:
//
// The Window.scrollbars property returns the scrollbars object, whose visibility can be checked.
//
// Português:
//
// A propriedade Window.scrollbars retorna o objeto scrollbars, cuja visibilidade pode ser verificada.
func (e Event) GetScrollBars() (scrollbars js.Value) {
	return e.Object.Get("scrollbars")
}

// GetStatusBar
//
// English:
//
// The Window.statusbar property returns the statusbar object, whose visibility can be toggled in the window.
//
// Português:
//
// A propriedade Window.statusbar retorna o objeto statusbar, cuja visibilidade pode ser alternada na janela.
func (e Event) GetStatusBar() (statusbar js.Value) {
	return e.Object.Get("statusbar")
}

// GetTop
//
// English:
//
// Returns a reference to the topmost window in the window hierarchy.
//
// Português:
//
// Retorna uma referência à janela superior na hierarquia de janelas.
func (e Event) GetTop() (top js.Value) {
	return e.Object.Get("top")
}

// GetNavigator
//
// English:
//
// Returns a reference to the Navigator object, which has methods and properties about the application running
// the script.
//
// Português:
//
// Retorna uma referência ao objeto Navigator, que possui métodos e propriedades sobre o aplicativo que executa
// o script.
func (e Event) GetNavigator() (navigator js.Value) {
	return e.Object.Get("navigator")
}
