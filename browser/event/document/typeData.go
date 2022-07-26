package document

import (
	"syscall/js"
)

type Navigator struct {
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

	// CookieEnabled
	//
	// English:
	//
	// Boolean value that indicates whether cookies are enabled or not.
	//
	//   Notes:
	//    * When the browser is configured to block third-party cookies, and navigator.cookieEnabled is invoked inside a
	//      third-party iframe, it returns true in Safari, Edge Spartan and IE (while trying to set a cookie in such
	//      scenario would fail). It returns false in Firefox and Chromium-based browsers.
	//    * Web browsers may prevent writing certain cookies in certain scenarios. For example, Chrome 80+ does not allow
	//      creating cookies with SameSite=None attribute, unless they are created over HTTPS and with Secure attribute.
	//
	// Português:
	//
	// Valor booleano que indica se os cookies estão habilitados ou não.
	//
	//   Notas:
	//    * Quando o navegador está configurado para bloquear cookies de terceiros e o navigator.cookieEnabled é invocado
	//      dentro de um iframe de terceiros, ele retorna true no Safari, Edge Spartan e IE (enquanto tentar definir um
	//      cookie nesse cenário falharia). Ele retorna false em navegadores baseados em Firefox e Chromium.
	//    * Os navegadores da Web podem impedir a gravação de determinados cookies em determinados cenários. Por exemplo,
	//      o Chrome 80+ não permite a criação de cookies com o atributo SameSite=None, a menos que sejam criados em HTTPS
	//      e com o atributo Secure.
	CookieEnabled bool

	Object js.Value
}

// GetCookieEnabled
//
// English:
//
// Boolean value that indicates whether cookies are enabled or not.
//
//   Notes:
//    * When the browser is configured to block third-party cookies, and navigator.cookieEnabled is invoked inside a
//      third-party iframe, it returns true in Safari, Edge Spartan and IE (while trying to set a cookie in such
//      scenario would fail). It returns false in Firefox and Chromium-based browsers.
//    * Web browsers may prevent writing certain cookies in certain scenarios. For example, Chrome 80+ does not allow
//      creating cookies with SameSite=None attribute, unless they are created over HTTPS and with Secure attribute.
//
// Português:
//
// Valor booleano que indica se os cookies estão habilitados ou não.
//
//   Notas:
//    * Quando o navegador está configurado para bloquear cookies de terceiros e o navigator.cookieEnabled é invocado
//      dentro de um iframe de terceiros, ele retorna true no Safari, Edge Spartan e IE (enquanto tentar definir um
//      cookie nesse cenário falharia). Ele retorna false em navegadores baseados em Firefox e Chromium.
//    * Os navegadores da Web podem impedir a gravação de determinados cookies em determinados cenários. Por exemplo,
//      o Chrome 80+ não permite a criação de cookies com o atributo SameSite=None, a menos que sejam criados em HTTPS
//      e com o atributo Secure.
func (e Navigator) GetCookieEnabled() (enabled bool) {
	return e.Object.Get("cookieEnabled").Bool()
}

// Geolocation
//
// English:
//
// The Geolocation interface represents an object able to obtain the position of the device programmatically. It gives
// Web content access to the location of the device. This allows a website or app to offer customized results based on
// the user's location.
//
// An object with this interface is obtained using the navigator.geolocation property implemented by the Navigator
// object.
//
//   Notes:
//    * For security reasons, when a web page tries to access location information, the user is notified and asked to
//      grant permission. Be aware that each browser has its own policies and methods for requesting this permission.
//    * This feature is available only in secure contexts (HTTPS), in some or all supporting browsers.
//
// Português:
//
// A interface Geolocation representa um objeto capaz de obter a posição do dispositivo programaticamente.
// Dá acesso ao conteúdo da Web à localização do dispositivo. Isso permite que um site ou aplicativo ofereça resultados
// personalizados com base na localização do usuário.
//
// An object with this interface is obtained using the navigator.geolocation property implemented by the Navigator
// object.
//
//   Notes:
//    * Por motivos de segurança, quando uma página da Web tenta acessar informações de localização, o usuário é
//      notificado e solicitado a conceder permissão. Esteja ciente de que cada navegador tem suas próprias políticas e
//      métodos para solicitar essa permissão.
//    * Este recurso está disponível apenas em contextos seguros (HTTPS), em alguns ou em todos os navegadores compatíveis.
type Geolocation struct {

	// enableHighAccuracy
	//
	// English:
	//
	// A boolean value that indicates the application would like to receive the best possible results.
	// If true and if the device is able to provide a more accurate position, it will do so. Note that this can result in
	// slower response times or increased power consumption (with a GPS chip on a mobile device for example). On the other
	// hand, if false, the device can take the liberty to save resources by responding more quickly and/or using less
	// power.
	//
	// Default: false.
	//
	// Português:
	//
	// Um valor booleano que indica que o aplicativo gostaria de receber os melhores resultados possíveis. Se for
	// verdadeiro e se o dispositivo for capaz de fornecer uma posição mais precisa, ele o fará. Observe que isso pode
	// resultar em tempos de resposta mais lentos ou maior consumo de energia (com um chip GPS em um dispositivo móvel,
	// por exemplo). Por outro lado, se falso, o dispositivo pode ter a liberdade de economizar recursos respondendo mais
	// rapidamente e ou usando menos energia.
	//
	// Padrão: falso.
	enableHighAccuracy bool

	// timeout
	//
	// English:
	//
	// A positive long value representing the maximum length of time (in milliseconds) the device is allowed to take in
	// order to return a position.
	// The default value is Infinity, meaning that getCurrentPosition() won't return until the position is available.
	//
	// Português:
	//
	// Um valor longo positivo que representa o tempo máximo (em milissegundos) que o dispositivo pode levar para retornar
	// uma posição.
	// O valor padrão é Infinity, o que significa que getCurrentPosition() não retornará até que a posição esteja
	// disponível.
	timeout int

	// maximumAge
	//
	// English:
	//
	// A positive long value indicating the maximum age in milliseconds of a possible cached position that is acceptable
	// to return.
	// If set to 0, it means that the device cannot use a cached position and must attempt to retrieve the real current
	// position.
	// If set to -1 the device must return a cached position regardless of its age.
	// Default: 0.
	//
	// Português:
	//
	// Um valor longo positivo que indica a idade máxima em milissegundos de uma possível posição em cache que é aceitável
	// retornar.
	// Se definido como 0, significa que o dispositivo não pode usar uma posição em cache e deve tentar recuperar a
	// posição atual real.
	// Se definido como -1, o dispositivo deve retornar uma posição em cache, independentemente de sua idade.
	// Padrão: 0.
	maximumAge int

	Object js.Value
}

// MaximumAge
//
// English:
//
// A positive long value indicating the maximum age in milliseconds of a possible cached position that is acceptable
// to return.
//
//   Input:
//     maximumAge: value indicating the maximum age in milliseconds.
//
// If set to 0, it means that the device cannot use a cached position and must attempt to retrieve the real current
// position.
// If set to -1 the device must return a cached position regardless of its age.
// Default: 0.
//
// Português:
//
// Um valor longo positivo que indica a idade máxima em milissegundos de uma possível posição em cache que é aceitável
// retornar.
//
//   Entrada:
//     maximumAge: valor que indica a idade máxima em milissegundos.
//
// Se definido como 0, significa que o dispositivo não pode usar uma posição em cache e deve tentar recuperar a
// posição atual real.
// Se definido como -1, o dispositivo deve retornar uma posição em cache, independentemente de sua idade.
// Padrão: 0.
func (e *Geolocation) MaximumAge(maximumAge int) {
	e.maximumAge = maximumAge
}

// Timeout
//
// English:
//
// A positive long value representing the maximum length of time (in milliseconds) the device is allowed to take in
// order to return a position.
//
//   Input:
//     timeout: value representing the maximum length of time (in milliseconds)
//
// The default value is Infinity, meaning that getCurrentPosition() won't return until the position is available.
//
// Português:
//
// Um valor longo positivo que representa o tempo máximo (em milissegundos) que o dispositivo pode levar para retornar
// uma posição.
//
//   Entrada:
//     timeout: valor que representa a duração máxima de tempo (em milissegundos)
//
// O valor padrão é Infinity, o que significa que getCurrentPosition() não retornará até que a posição esteja
// disponível.
func (e *Geolocation) Timeout(timeout int) {
	e.timeout = timeout
}

// EnableHighAccuracy
//
// English:
//
// A boolean value that indicates the application would like to receive the best possible results.
//
//   Input:
//     enableHighAccuracy: indicates the application would like to receive the best possible results.
//
// If true and if the device is able to provide a more accurate position, it will do so. Note that this can result in
// slower response times or increased power consumption (with a GPS chip on a mobile device for example). On the other
// hand, if false, the device can take the liberty to save resources by responding more quickly and/or using less
// power.
//
// Default: false.
//
// Português:
//
// Um valor booleano que indica que o aplicativo gostaria de receber os melhores resultados possíveis.
//
//   Entrada:
//     enableHighAccuracy: indica que o aplicativo gostaria de receber os melhores resultados possíveis.
//
// Se for verdadeiro e se o dispositivo for capaz de fornecer uma posição mais precisa, ele o fará. Observe que isso
// pode resultar em tempos de resposta mais lentos ou maior consumo de energia (com um chip GPS em um dispositivo móvel,
// por exemplo). Por outro lado, se falso, o dispositivo pode ter a liberdade de economizar recursos respondendo mais
// rapidamente e ou usando menos energia.
//
// Padrão: falso.
func (e *Geolocation) EnableHighAccuracy(enableHighAccuracy bool) {
	e.enableHighAccuracy = enableHighAccuracy
}

type Coordinate struct {
	Latitude  float64
	Longitude float64

	Accuracy     float64
	ErrorCode    int
	ErrorMessage string
}

// GetGeolocation
//
// English:
//
// Get the current position of the device.
//
//   Output:
//     coordinate: coordinate object
//
//   Notes:
//     * Accuracy is in meters
//     * This feature is available only in secure contexts (HTTPS).
//
// Português:
//
// Retorna a posição atual do dispositivo.
//
//   Saída:
//     coordinate: objeto de coordenadas
//
//   Notas:
//     * Accuracy é em metros.
//     * Esse recurso está disponível apenas em contextos seguros (HTTPS).
// todo: not tested
func (e *Geolocation) GetGeolocation() (coordinate Coordinate) {
	var options = e.prepareOptions()

	onError := js.FuncOf(func(errJs js.Value, _ []js.Value) interface{} {
		coordinate.ErrorCode = errJs.Get("code").Int()
		coordinate.ErrorMessage = errJs.Get("message").String()
		return nil
	})

	onSuccess := js.FuncOf(func(pos js.Value, _ []js.Value) interface{} {
		if pos.IsUndefined() || pos.IsNull() {
			return nil
		}

		coordinates := pos.Get("coords")
		coordinate.Latitude = coordinates.Get("latitude").Float()
		coordinate.Longitude = coordinates.Get("longitude").Float()
		coordinate.Accuracy = coordinates.Get("accuracy").Float()

		return nil
	})

	js.Global().Get("navigator").Get("geolocation").Call("getCurrentPosition", onSuccess, onError, options)
	return
}

// prepareOptions
//
// English:
//
// An optional object from Geolocation.getCurrentPosition()
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Geolocation/getCurrentPosition
//
// Português:
//
// Um objeto opcional de Geolocation.getCurrentPosition()
//
// Veja https://developer.mozilla.org/en-US/docs/Web/API/Geolocation/getCurrentPosition
func (e Geolocation) prepareOptions() (options js.Value) {
	options = js.Global().Get("Object")

	if e.maximumAge < 0 {
		options.Set("maximumAge", "Infinity")
	} else {
		options.Set("maximumAge", e.maximumAge)
	}

	if e.timeout > 0 {
		options.Set("timeout", e.timeout)
	} else {
		options.Set("timeout", "Infinity")
	}

	options.Set("enableHighAccuracy", e.enableHighAccuracy)
	return
}

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

	// EventName
	//
	// English:
	//
	// Name o event
	//
	// Português:
	//
	// Nome do evento
	EventName EventName

	// Width
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
	Width float64

	// Height
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
	Height float64

	// Name
	//
	// English:
	//
	// Gets the name of the window's browsing context.
	//
	// Português:
	//
	// Obtém o nome do contexto de navegação da janela.
	Name string

	// Length
	//
	// English:
	//
	// Returns the number of frames (either <frame> or <iframe> elements) in the window.
	//
	// Português:
	//
	// Retorna o número de quadros (elementos <frame> ou <iframe>) na janela.
	Length int

	// Closed
	//
	// English:
	//
	// The Window.closed read-only property indicates whether the referenced window is closed or not
	//
	// Português:
	//
	// A propriedade somente leitura Window.closed indica se a janela referenciada está fechada ou não
	Closed bool

	// OuterHeight
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
	OuterHeight float64

	// OuterWidth
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
	OuterWidth float64

	// ScrollX
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
	ScrollX float64

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
	ScrollY float64

	// ScreenX
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
	ScreenX float64

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
	ScreenY float64

	// Opener
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
	Opener js.Value

	// Parent
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
	Parent js.Value

	// Screen
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
	Screen js.Value

	// ScrollBars
	//
	// English:
	//
	// The Window.scrollbars property returns the scrollbars object, whose visibility can be checked.
	//
	// Português:
	//
	// A propriedade Window.scrollbars retorna o objeto scrollbars, cuja visibilidade pode ser verificada.
	ScrollBars js.Value

	// StatusBar
	//
	// English:
	//
	// The Window.statusbar property returns the statusbar object, whose visibility can be toggled in the window.
	//
	// Português:
	//
	// A propriedade Window.statusbar retorna o objeto statusbar, cuja visibilidade pode ser alternada na janela.
	StatusBar js.Value

	// Top
	//
	// English:
	//
	// Returns a reference to the topmost window in the window hierarchy.
	//
	// Português:
	//
	// Retorna uma referência à janela superior na hierarquia de janelas.
	Top js.Value

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

	Object js.Value
}

// Blur
//
// English:
//
// Shifts focus away from the window.
//
// Português:
//
// Desvia o foco da janela.
func (e Data) Blur() {
	e.Object.Call("blur")
}

// Close
//
// English:
//
// The Close() method closes the current window, or the window on which it was called.
//
// This method can only be called on windows that were opened by a script using the Window.open() method. If the window
// was not opened by a script, an error similar to this one appears in the console: Scripts may not close windows that
// were not opened by script.
//
//   Notes
//    * Close() has no effect when called on Window objects returned by HTMLIFrameElement.contentWindow.
//
// Português:
//
// O método Close() fecha a janela atual ou a janela na qual foi chamado.
//
// Este método só pode ser chamado em janelas que foram abertas por um script usando o método Window.open(). Se a janela
// não foi aberta por um script, um erro semelhante a este aparece no console: Scripts não podem fechar janelas que não
// foram abertas por script.
//
//   Notas
//    * Close() não tem efeito quando chamado em objetos Window retornados por HTMLIFrameElement.contentWindow.
func (e Data) Close() {
	e.Object.Call("close")
}

// Focus
//
// English:
//
// Makes a request to bring the window to the front. It may fail due to user settings and the window isn't guaranteed to
// be frontmost before this method returns.
//
// Português:
//
// Faz um pedido para trazer a janela para a frente. Pode falhar devido às configurações do usuário e não é garantido
// que a janela esteja na frente antes que esse método retorne.
func (e Data) Focus() {
	e.Object.Call("focus")
}

// MoveTo
//
// English:
//
// Moves the window to the specified coordinates.
//
//   Input:
//     x: is the horizontal coordinate to be moved to.
//     y: is the vertical coordinate to be moved to.
//
//   Notes:
//    * This function moves the window to an absolute location. In contrast, window.moveBy() moves the window relative to its current location.
//
// Português:
//
// Move a janela para as coordenadas especificadas.
//
//   Entrada:
//     x: é a coordenada horizontal a ser movida.
//     y: é a coordenada vertical a ser movida.
//
//   Notes:
//    * This function moves the window to an absolute location. In contrast, window.moveBy() moves the window relative to its current location.
func (e Data) MoveTo(x, y float64) {
	e.Object.Call("moveTo", x, y)
}

// MoveBy
//
// English:
//
// The moveBy() method of the Window interface moves the current window by a specified amount.
//
//   Input:
//     deltaX: is the amount of pixels to move the window horizontally. Positive values are to the right, while negative
//       values are to the left.
//     deltaY: is the amount of pixels to move the window vertically. Positive values are down, while negative values
//       are up.
//
//   Notes:
//    * This function moves the window relative to its current location. In contrast, window.moveTo() moves the window
//      to an absolute location.
//
// Português:
//
// O método moveBy() da interface Window move a janela atual por um valor especificado.
//
//   Entrada:
//     deltaX: é a quantidade de pixels para mover a janela horizontalmente. Os valores positivos estão à direita,
//       enquanto os valores negativos estão à esquerda.
//     deltaY: é a quantidade de pixels para mover a janela verticalmente. Os valores positivos estão em baixa, enquanto
//       os valores negativos estão em alta.
//
//   Notas:
//    * Esta função move a janela em relação à sua localização atual. Em contraste, window.moveTo() move a janela para
//      um local absoluto.
func (e Data) MoveBy(deltaX, deltaY float64) {
	e.Object.Call("moveBy", deltaX, deltaY)
}

// ResizeBy
//
// English:
//
// The Window.resizeBy() method resizes the current window by a specified amount.
//
//   Input:
//     xDelta: Number of pixels to grow the window horizontally.
//     yDelta: Number of pixels to grow the window vertically.
//
// Português:
//
// O método Window.resizeBy() redimensiona a janela atual em um valor especificado.
//
//   Entrada:
//     xDelta: Número de pixels para aumentar a janela horizontalmente.
//     yDelta: Número de pixels para aumentar a janela verticalmente.
func (e Data) ResizeBy(deltaX, deltaY float64) {
	e.Object.Call("resizeBy", deltaX, deltaY)
}

// ResizeTo
//
// English:
//
// The Window.resizeTo() method dynamically resizes the window.
//
//   Input:
//     width: An integer representing the new outerWidth in pixels (including scroll bars, title bars, etc).
//     height: An integer value representing the new outerHeight in pixels (including scroll bars, title bars, etc).
//
// Português:
//
// O método Window.resizeTo() redimensiona dinamicamente a janela.
//
//   Entrada:
//     width: Um inteiro que representa o novo outerWidth em pixels (incluindo barras de rolagem, barras de título etc.)
//     height: Um valor inteiro que representa o novo outerHeight em pixels (incluindo barras de rolagem, barras de
//       título etc.)
func (e Data) ResizeTo(width, height float64) {
	e.Object.Call("resizeTo", width, height)
}

// Scroll
//
// English:
//
// Scrolls the window to a particular place in the document.
//
//   Input:
//     x: Is the pixel along the horizontal axis of the document that you want displayed in the upper left.
//     y: is the pixel along the vertical axis of the document that you want displayed in the upper left.
//
// Português:
//
// Rola a janela para um local específico no documento.
//
//   Entrada:
//     x: É o pixel ao longo do eixo horizontal do documento que você deseja exibir no canto superior esquerdo.
//     y: é o pixel ao longo do eixo vertical do documento que você deseja exibir no canto superior esquerdo.
func (e Data) Scroll(x, y float64) {
	e.Object.Call("scroll", x, y)
}

// ScrollBy
//
// English:
//
// Scrolls the document in the window by the given amount.
//
//   Input:
//     x: Is the horizontal pixel value that you want to scroll by.
//     y: Is the vertical pixel value that you want to scroll by.
//
// Português:
//
// Rola o documento na janela pela quantidade especificada.
//
//   Entrada:
//     x: É o valor de pixel horizontal pelo qual você deseja rolar.
//     y: É o valor de pixel vertical pelo qual você deseja rolar.
func (e Data) ScrollBy(x, y float64) {
	e.Object.Call("scrollBy", x, y)
}

// ScrollTo
//
// English:
//
// Scrolls to a particular set of coordinates in the document.
//
//   Input:
//     x: Is the pixel along the horizontal axis of the document that you want displayed in the upper left.
//     y: Is the pixel along the vertical axis of the document that you want displayed in the upper left.
//
// Português:
//
// Rola para um determinado conjunto de coordenadas no documento.
//
//   Entrada:
//     x: É o pixel ao longo do eixo horizontal do documento que você deseja exibir no canto superior esquerdo.
//     y: É o pixel ao longo do eixo vertical do documento que você deseja exibir no canto superior esquerdo.
func (e Data) ScrollTo(x, y float64) {
	e.Object.Call("scrollTo", x, y)
}
