package stage

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/document"
	keyboard "github.com/helmutkemper/iotmaker.webassembly/browser/event/keyBoard"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/mouse"
	"github.com/helmutkemper/iotmaker.webassembly/platform/engine"
	"log"
	"sync"
	"syscall/js"
)

// todo: https://developer.mozilla.org/en-US/docs/Web/API/Window/showOpenFilePicker
// todo: https://developer.mozilla.org/en-US/docs/Web/API/Window/showSaveFilePicker
// https://developer.mozilla.org/en-US/docs/Web/API/Window

// Compatible
//
// English:
//
// Defines the functions necessary for the graphic elements to work on the stage.
//
// Português:
//
// Define as funções necessárias ao funcionamento dos elementos gráficos no palco.
type Compatible interface {
	Get() js.Value
}

// Stage
//
// English:
//
// Defines the structure necessary for the wasm application stage to work.
//
// Português:
//
// Define a estrutura necessária para o funcionamento do palco da aplicação wasm.
type Stage struct {
	// selfDocument
	//
	// English:
	//
	// Defines the browser or graphic 'document'.
	//
	// Português:
	//
	// Define o 'documento' do navegador ou do elemnto gráfico.
	selfDocument js.Value

	// selfWindow
	//
	// English:
	//
	// Defines the browser window.
	//
	// Português:
	//
	// Define a janela do navegador.
	selfWindow js.Value

	// engine
	//
	// English:
	//
	// Defines the engine responsible for the calculations used in the application.
	//
	// The engine is responsible for making animations run at 120 fps.
	//
	// Português:
	//
	// Define a engine responssável pelos cálculos usados na aplicação.
	//
	// A engine é responssável por fazer a animações rodar a 120 fps.
	engine engine.IEngine

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

	// fnResize
	//
	// English:
	//
	// Gets the pointer of the stage's onResize event function.
	//
	// Português:
	//
	// Recebe o ponteiro da função de evento onResize do stage.
	fnResize *js.Func

	// fnLoad
	//
	// English:
	//
	// Gets the pointer of the stage's onLoad event function.
	//
	// Português:
	//
	// Recebe o ponteiro da função de evento onLoad do stage.
	fnLoad *js.Func

	// fnUnload
	//
	// English:
	//
	// Gets the pointer of the stage's onUnLoad event function.
	//
	// Português:
	//
	// Recebe o ponteiro da função de evento onUnLoad do stage.
	fnUnload *js.Func

	// fnKeyUp
	//
	// English:
	//
	// Gets the pointer of the stage's keyUp event function.
	//
	// Português:
	//
	// Recebe o ponteiro da função de evento keyUp do stage.
	fnKeyUp *js.Func

	// fnKeyDown
	//
	// English:
	//
	// Gets the pointer of the stage's keyDown event function.
	//
	// Português:
	//
	// Recebe o ponteiro da função de evento keyDown do stage.
	fnKeyDown *js.Func
}

// Engine
//
// English:
//
// Defines the engine responsible for the calculations used in the application.
//
//	Input:
//	  engine: object pointer compatible with engine interface
//
// The engine is responsible for making animations run at 120 fps.
//
//	Notes:
//	  * The engine object must have been properly initialized before being pointed.
//
// Português:
//
// Define a engine responssável pelos cálculos usados na aplicação.
//
//	Entrada:
//	  engine: ponteiro de objeto compatível com a interface de engine
//
// A engine é responssável por fazer a animações rodar a 120 fps.
//
//	Notas:
//	  * O objeto engine deve ter sido devidamente inicializado antes de ser apontado.
func (e *Stage) Engine(engine engine.IEngine) {
	e.engine = engine
}

// WindowName
//
// English:
//
// Sets the name of the window's browsing context.
//
// The name of the window is used primarily for setting targets for hyperlinks and forms. Browsing contexts do not need
// to have names.
//
// Modern browsers will reset Window.name to an empty string if a tab loads a page from a different domain, and restore
// the name if the original page is reloaded (e.g. by selecting the "back" button). This prevents an untrusted page from
// accessing any information that the previous page might have stored in the property (potentially the new page might
// also modify such data, which might then be read by the original page if it was reloaded).
//
// Window.name has also been used in some frameworks for providing cross-domain messaging (e.g. Dojo's
// dojox.io.windowName) as a more secure alternative to JSONP. Modern web applications hosting sensitive data should,
// however, not rely on window.name for cross-domain messaging — that is not its intended purpose and there are
// safer/better ways of sharing information between windows. Window.postMessage() is the recommended mechanism.
//
//	Notes:
//	  * window.name converts all stored values to their string representations using the toString method.
//
// Português:
//
// Define o nome do contexto de navegação da janela.
//
// O nome da janela é usado principalmente para definir destinos para hiperlinks e formulários. Os contextos de
// navegação não precisam ter nomes.
//
// Os navegadores modernos redefinirão Window.name para uma string vazia se uma guia carregar uma página de um domínio
// diferente e restaurará o nome se a página original for recarregada (por exemplo, selecionando o botão "voltar").
// Isso impede que uma página não confiável acesse qualquer informação que a página anterior possa ter armazenado na
// propriedade (potencialmente, a nova página também pode modificar esses dados, que podem ser lidos pela página
// original se ela for recarregada).
//
// Window.name também foi usado em algumas estruturas para fornecer mensagens entre domínios (por exemplo,
// dojox.io.windowName do Dojo) como uma alternativa mais segura ao JSONP. Os aplicativos da Web modernos que hospedam
// dados confidenciais não devem, no entanto, depender do window.name para mensagens entre domínios — essa não é a
// finalidade pretendida e existem maneiras melhores e mais seguras de compartilhar informações entre janelas.
// Window.postMessage() é o mecanismo recomendado.
//
//	Notas:
//	  * window.name converte todos os valores armazenados em suas representações de string usando o método toString.
func (e *Stage) WindowName(name string) {
	js.Global().Get("window").Set("name", name)
}

// NewWindow
//
// English:
//
// Opens a new simplified browser window.
//
//	Input:
//	  url: window content address.
//
//	Output:
//	  newWindow: Pointer to new Stage object with new window controls.
//
// Português:
//
// Abre uma nova janela simplificada no navegador.
//
//	Entrada:
//	  url: endereço do conteúdo da janela.
//
//	Saída:
//	  newWindow: Ponteiro para novo objeto Stage com os controles da nova janela.
func (e *Stage) NewWindow(url string) (newWindow *Stage) {
	// WASM has a bug, use: "_blank", "width=100,height=100"
	window := js.Global().Get("window").Call("open", url, "_blank", "width=100,height=100")

	return e.NewStage(window)
}

// NewStage
//
// English:
//
// Transforms a javascript object from the window.open() function into a working Stage pointer.
//
//	Notes:
//	 * This function is used internally, please use the NewWindow() function
//
// Português:
//
// Transforma um objeto javascript proveniente da função window.open() em um ponteiro de Stage funcional.
//
//	Notas:
//	 * Esta função é usada internamente, por favor, use a função NewWindow()
func (e *Stage) NewStage(window js.Value) (newWindow *Stage) {
	newWindow = new(Stage)
	newWindow.selfDocument = window.Get("document")
	newWindow.selfWindow = window.Get("window")
	newWindow.listener = new(sync.Map)
	return
}

// Init
//
// English:
//
//	Initializes the document with the browser's main document.
//
// Português:
//
//	Inicializa o documento com o documento principal do navegador.
func (e *Stage) Init() {
	e.selfDocument = js.Global().Get("document")
	e.selfWindow = js.Global().Get("window")
	e.listener = new(sync.Map)
}

// Append
//
// English:
//
// Adds new graphical elements to the application stage.
//
//	Input:
//	  elements: list of graphical objects compatible with the Compatible interface.
//
// Português:
//
// Adiciona novos elementos gráficos ao palco da aplicação.
//
//	Entrada:
//	  elements: lista de objetos gráficos compatíveis com a interface Compatible.
func (e *Stage) Append(elements ...Compatible) (ref *Stage) {
	fragment := js.Global().Get("document").Call("createDocumentFragment")
	for _, element := range elements {
		fragment.Call("appendChild", element.Get())
	}

	e.selfDocument.Get("body").Call("appendChild", fragment)
	return e
}

// Get
//
// English:
//
//	Returns the javascript object.
//
// Português:
//
//	Retorna o objeto javascript.
func (e *Stage) Get() js.Value {
	return e.selfDocument
}

// MouseAuto
//
// English:
//
//	Sets the mouse pointer to auto.
//
// Português:
//
//	Define o ponteiro do mouse como automático.
func (e *Stage) MouseAuto() (ref *Stage) {
	e.selfDocument.Get("body").Set("style", mouse.KCursorAuto.String())
	return e
}

// MouseHide #testar
//
// English:
//
//	Sets the mouse pointer to hide.
//
// Português:
//
//	Define o ponteiro do mouse como oculto.
func (e *Stage) MouseHide() (ref *Stage) {
	e.selfDocument.Get("body").Set("style", mouse.KCursorNone.String())
	return e
}

// SetMouse #testar
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
func (e *Stage) SetMouse(value mouse.CursorType) (ref *Stage) {
	e.selfDocument.Get("body").Set("style", value.String())
	return e
}

// Remove
//
// English:
//
//	Removes an html element from the document.
//
//	 Input:
//	   value: js.Value element containing an html document.
//
// Português:
//
//	Remove um elemento html do documento.
//
//	 Entrada:
//	   value: elemento js.Value contendo um documento html.
func (e *Stage) Remove(value interface{}) (ref *Stage) {
	e.selfDocument.Get("body").Call("removeChild", value)
	return e
}

// GetWidth #testar
//
// English:
//
//	Returns the width of the document in pixels.
//
//	 Output:
//	   width: document size in pixels.
//
// Português:
//
//	Retorna o comprimento do documento em pixels.
//
//	 Saída:
//	   width: tamanho do documento em pixels.
func (e Stage) GetWidth() (width int) {
	return js.Global().Get("window").Get("innerWidth").Int()
}

// GetHeight #testar
//
// English:
//
//	Returns the length of the document in pixels.
//
//	 Output:
//	   width: document size in pixels.
//
// Português:
//
//	Retorna a altura do documento em pixels.
//
//	 Saída:
//	   width: tamanho do documento em pixels.
func (e Stage) GetHeight() (height int) {
	return js.Global().Get("window").Get("innerHeight").Int()
}

// ResizeStageToScreen #deprecated
//
// English:
//
//	Resizes the document to the size of the main document.
//
// Português:
//
//	Redimensiona o documento para o tamanho do documento principal.
func (e *Stage) ResizeStageToScreen() (ref *Stage) {
	e.selfDocument.Get("body").Set("width", js.Global().Get("document").Get("body").Get("clientWidth").Int())
	e.selfDocument.Get("body").Set("height", js.Global().Get("document").Get("body").Get("clientHeight").Int())
	return e
}

// GetById
//
// Português:
//
//	Retorna a referência do elemento através do seu ID.
//
//	 Entrada:
//	   id: string que diferência maiúsculas e minúsculas representando o ID único do elemento sendo
//	       procurado.
//	 Nota:
//	   * Elemento é uma referência a um objeto Element, ou null se um elemento com o ID especificado
//	     não estiver contido neste documento.
//	   * Se não existe um elemento com o id fornecido, esta função retorna null. Note, o parâmetro ID
//	     diferência maiúsculas e minúsculas. Assim document.getElementById("Main") retornará null ao
//	     invés do elemento <div id="main">, devido a "M" e "m" diferirem para o objetivo deste método;
//	   * Elementos que não estão no documento não são procurados por getElementById. Quando criar um
//	     elemento e atribuir um ID ao mesmo, você deve inserir o elemento na árvore do documento com
//	     insertBefore ou método similar antes que você possa acessá-lo com getElementById:
//
//	       var elemento = document.createElement("div");
//	       elemento.id = 'testqq';
//	       var el = document.getElementById('testqq'); // el será null!
//
//	   * Documentos não-HTML, a implementação do DOM deve ter informações que diz quais atributos são
//	     do tipo ID.  Atributos com o nome "id" não são do tipo ID a menos que assim sejam definidos
//	     nos documentos DTD. O atributo id é definido para ser um tipo ID em casos comuns de  XHTML,
//	     XUL, e outros. Implementações que não reconhecem se os atributos são do tipo ID, ou não são
//	     esperados retornam null.
func (e Stage) GetById(id string) (element interface{}) {
	elementRet := js.Global().Get("document").Call("getElementById", id)
	if elementRet.IsUndefined() == true || elementRet.IsNull() {
		log.Printf("getElementById(%v).undefined", id)
		return nil
	}

	return elementRet
}

// AddListener
//
// English:
//
//  Associates a function with a mouse event.
//
//   Example:
//     stage.AddListener(browserMouse.KEventMouseOver, onMouseEvent)
//     timer := time.NewTimer(10 * time.Second)
//     go func() {
//       select {
//         case <-timer.C:
//         stage.RemoveListener(mouse.KEventMouseOver)
//       }
//     }()
//
// Português:
//
//  Associa uma função a um evento do mouse.
//
//   Exemplo:
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
//func (e *Stage) AddListener(eventType interface{}, manager mouse.SimpleManager) (ref *Stage) {
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
//		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventAnimation.EventAnimation:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventClipBoard.EventClipBoard:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventDrag.EventDrag:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventFocus.EventFocus:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventHashChange.EventHashChange:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventInput.EventInput:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventKeyboard.EventKeyboard:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	//case mouse.Event:
//	//	e.listener.Store(converted.String(), mouseMoveEvt)
//	//	e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventPageTransition.EventPageTransition:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventUi.EventUi:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	case eventWheel.EventWheel:
//		e.listener.Store(converted.String(), mouseMoveEvt)
//		e.selfDocument.Call("addEventListener", converted.String(), mouseMoveEvt)
//
//	default:
//		log.Fatalf("event must be a event type")
//	}
//
//	return e
//}

//func (e *Stage) RemoveListener(eventType interface{}) (ref *Stage) {
//
//	switch converted := eventType.(type) {
//	case event.Event:
//		f, _ := e.listener.Load(converted.String())
//		e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	case eventAnimation.EventAnimation:
//		f, _ := e.listener.Load(converted.String())
//		e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	case eventClipBoard.EventClipBoard:
//		f, _ := e.listener.Load(converted.String())
//		e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	case eventDrag.EventDrag:
//		f, _ := e.listener.Load(converted.String())
//		e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	case eventFocus.EventFocus:
//		f, _ := e.listener.Load(converted.String())
//		e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	case eventHashChange.EventHashChange:
//		f, _ := e.listener.Load(converted.String())
//		e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	case eventInput.EventInput:
//		f, _ := e.listener.Load(converted.String())
//		e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	case eventKeyboard.EventKeyboard:
//		f, _ := e.listener.Load(converted.String())
//		e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	//case mouse.Event:
//	//	f, _ := e.listener.Load(converted.String())
//	//	e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	case eventPageTransition.EventPageTransition:
//		f, _ := e.listener.Load(converted.String())
//		e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	case eventUi.EventUi:
//		f, _ := e.listener.Load(converted.String())
//		e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	case eventWheel.EventWheel:
//		f, _ := e.listener.Load(converted.String())
//		e.selfDocument.Call("removeEventListener", converted.String(), f)
//
//	default:
//		log.Fatalf("event must be a event type")
//	}
//
//	return e
//}

// GetFPS
//
// English:
//
// Returns the amount of current FPS used in calculations.
//
// Português:
//
// Retorna a quantidade de FPS atual usado nos cálculos.
func (e *Stage) GetFPS() (fps int) {
	return e.engine.GetFPS()
}

// SetFPS
//
// English:
//
// Sets the total FPS used in calculations and moves.
//
// Português:
//
// Define o total de FPS usado nos cálculos e movimentos.
func (e *Stage) SetFPS(value int) {
	e.engine.SetFPS(value)
}

// AddCursorDrawFunction
//
// English:
//
// Allows you to recreate the function that draws the cursor.
//
// Português:
//
// Permite recriar a função que desenha o cursor.
func (e *Stage) AddCursorDrawFunction(runnerFunc func()) {
	e.engine.CursorAddDrawFunction(runnerFunc)
}

// RemoveCursorDrawFunction
//
// English:
//
// Removes the role responsible for recreating the cursor.
//
// Português:
//
// Remove a função responssável por recria o cursor.
func (e *Stage) RemoveCursorDrawFunction() {
	e.engine.CursorRemoveDrawFunction()
}

// AddHighLatencyFunctions
//
// English:
//
// Adds a high latency function, a low execution priority function.
//
//	Input:
//	  runnerFunc: function to be performed.
//
//	Output:
//	  UId: used to identify the function when removing.
//	  total: total number of functions running.
//
//	Notes:
//	  * High latency functions are secondary functions designed to run at a lower FPS rate.
//
// Português:
//
// Adiciona uma função de alta latencia, uma função de baixa prioridade de execussão.
//
//	Entrada:
//	  runnerFunc: função a ser executada.
//
//	Saída:
//	  UId da função, usado para identificar a função na hora de remover.
//	  total: quantidade total de funções em execução.
//
//	Notas:
//	  * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
func (e *Stage) AddHighLatencyFunctions(runnerFunc func()) (UId string, total int) {
	UId, total = e.engine.HighLatencyAddToFunctions(runnerFunc)
	return
}

// DeleteHighLatencyFunctions
//
// English:
//
// Removes a high latency function added by the HighLatencyAddToFunctions() function.
//
//	Input:
//	  UId: ID returned by the HighLatencyAddToFunctions() function.
//
//	Notes:
//	  * High latency functions are secondary functions designed to run at a lower FPS rate.
//
// Português:
//
// Remove uma função de alta latencia adicionada pela função HighLatencyAddToFunctions().
//
//	Entrada:
//	  UId: ID retornado pela função HighLatencyAddToFunctions().
//
//	Notas:
//	  * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
func (e *Stage) DeleteHighLatencyFunctions(UId string) {
	e.engine.HighLatencyDeleteFromFunctions(UId)
}

// SetHighLatencyZIndex
//
// English:
//
// Allows you to change the order of execution of the function, in the execution list.
//
//	Input:
//	  UId: ID returned by the HighLatencyAddToFunctions() function.
//	  index: 0 for the first function in the list
//
//	Notes:
//	  * High latency functions are secondary functions designed to run at a lower FPS rate.
//
// Português:
//
// Permite trocar a ordem de execução da função, na lista de execução.
//
//	Entrada:
//	  UId: ID retornado pela função HighLatencyAddToFunctions().
//	  index: 0 para a primeira função da lista
//
//	Notas:
//	  * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
func (e *Stage) SetHighLatencyZIndex(UId string, index int) int {
	return e.engine.HighLatencySetZIndex(UId, index)
}

// AddMathFunctions
//
// English:
//
// Português:
func (e *Stage) AddMathFunctions(runnerFunc func()) (UId string, total int) {
	UId, total = e.engine.MathAddToFunctions(runnerFunc)
	return
}

// DeleteMathFunctions
//
// English:
//
// Português:
func (e *Stage) DeleteMathFunctions(UId string) {
	e.engine.MathDeleteFromFunctions(UId)
}

// SetMathZIndex
//
// English:
//
// Português:
func (e *Stage) SetMathZIndex(UId string, index int) int {
	return e.engine.MathSetZIndex(UId, index)
}

// AddDrawFunctions
//
// English:
//
// Português:
func (e *Stage) AddDrawFunctions(runnerFunc func()) (UId string, total int) {
	UId, total = e.engine.DrawAddToFunctions(runnerFunc)
	return
}

// DeleteDrawFunctions
//
// English:
//
// Português:
func (e *Stage) DeleteDrawFunctions(UId string) {
	e.engine.DrawDeleteFromFunctions(UId)
}

// SetDrawZIndex
//
// English:
//
// Português:
func (e *Stage) SetDrawZIndex(UId string, index int) int {
	return e.engine.DrawSetZIndex(UId, index)
}

func (e *Stage) AddListenerKeyUp(keyBoardCh chan keyboard.Data) (ref *Stage) {
	if e.fnKeyUp != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		keyBoardCh <- keyboard.EventManager(keyboard.KEventKeyUp, this, args)
		return nil
	})
	e.fnKeyUp = &fn

	e.selfDocument.Call(
		"addEventListener",
		"keyup",
		*e.fnKeyUp,
	)
	return e
}

func (e *Stage) AddListenerKeyDown(keyBoardCh chan keyboard.Data) (ref *Stage) {
	if e.fnKeyDown != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		keyBoardCh <- keyboard.EventManager(keyboard.KEventKeyDown, this, args)
		return nil
	})
	e.fnKeyDown = &fn

	e.selfDocument.Call(
		"addEventListener",
		"keydown",
		*e.fnKeyDown,
	)
	return e
}

func (e *Stage) AddListenerResize(windowEvet chan document.Data) (ref *Stage) {
	if e.fnResize != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		windowEvet <- document.EventManager(document.KEventResize, this, args)
		return nil
	})
	e.fnResize = &fn

	e.selfWindow.Get("window").Call(
		"addEventListener",
		"resize",
		*e.fnResize,
	)
	return e
}

func (e *Stage) RemoveListenerResize() (ref *Stage) {
	if e.fnResize == nil {
		return e
	}

	e.selfWindow.Get("window").Call(
		"removeEventListener",
		"resize",
		*e.fnResize,
	)
	e.fnResize = nil
	return e
}

func (e *Stage) AddListenerLoad(windowEvet chan document.Data) (ref *Stage) {
	if e.fnLoad != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		windowEvet <- document.EventManager(document.KEventLoad, this, args)
		return nil
	})
	e.fnLoad = &fn

	e.selfWindow.Get("window").Call(
		"addEventListener",
		"load",
		*e.fnLoad,
	)
	return e
}

func (e *Stage) AddListenerOnUnload(windowEvet chan document.Data) (ref *Stage) {
	if e.fnUnload != nil {
		return e
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		windowEvet <- document.EventManager(document.KEventUnLoad, this, args)
		return nil
	})
	e.fnUnload = &fn

	e.selfWindow.Get("window").Call(
		"addEventListener",
		"unload",
		*e.fnUnload,
	)
	return e
}

func (e *Stage) RemoveListenerOnLoad() (ref *Stage) {
	if e.fnLoad == nil {
		return e
	}

	e.selfWindow.Get("window").Call(
		"removeEventListener",
		"resize",
		*e.fnLoad,
	)
	e.fnLoad = nil
	return e
}

// ------------------------------------------------------------------------------------

// Blur
//
// English:
//
// Shifts focus away from the window.
//
// Português:
//
// Desvia o foco da janela.
func (e *Stage) Blur() {
	e.selfWindow.Call("blur")
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
//	Notes
//	 * Close() has no effect when called on Window objects returned by HTMLIFrameElement.contentWindow.
//
// Português:
//
// O método Close() fecha a janela atual ou a janela na qual foi chamado.
//
// Este método só pode ser chamado em janelas que foram abertas por um script usando o método Window.open(). Se a janela
// não foi aberta por um script, um erro semelhante a este aparece no console: Scripts não podem fechar janelas que não
// foram abertas por script.
//
//	Notas
//	 * Close() não tem efeito quando chamado em objetos Window retornados por HTMLIFrameElement.contentWindow.
func (e *Stage) Close() {
	e.selfWindow.Call("close")
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
func (e *Stage) Focus() {
	e.selfWindow.Call("focus")
}

// MoveTo
//
// English:
//
// Moves the window to the specified coordinates.
//
//	Input:
//	  x: is the horizontal coordinate to be moved to.
//	  y: is the vertical coordinate to be moved to.
//
//	Notes:
//	 * This function moves the window to an absolute location. In contrast, window.moveBy() moves the window relative to its current location.
//
// Português:
//
// Move a janela para as coordenadas especificadas.
//
//	Entrada:
//	  x: é a coordenada horizontal a ser movida.
//	  y: é a coordenada vertical a ser movida.
//
//	Notes:
//	 * This function moves the window to an absolute location. In contrast, window.moveBy() moves the window relative to its current location.
func (e *Stage) MoveTo(x, y float64) {
	e.selfWindow.Call("moveTo", x, y)
}

// MoveBy
//
// English:
//
// The moveBy() method of the Window interface moves the current window by a specified amount.
//
//	Input:
//	  deltaX: is the amount of pixels to move the window horizontally. Positive values are to the right, while negative
//	    values are to the left.
//	  deltaY: is the amount of pixels to move the window vertically. Positive values are down, while negative values
//	    are up.
//
//	Notes:
//	 * This function moves the window relative to its current location. In contrast, window.moveTo() moves the window
//	   to an absolute location.
//
// Português:
//
// O método moveBy() da interface Window move a janela atual por um valor especificado.
//
//	Entrada:
//	  deltaX: é a quantidade de pixels para mover a janela horizontalmente. Os valores positivos estão à direita,
//	    enquanto os valores negativos estão à esquerda.
//	  deltaY: é a quantidade de pixels para mover a janela verticalmente. Os valores positivos estão em baixa, enquanto
//	    os valores negativos estão em alta.
//
//	Notas:
//	 * Esta função move a janela em relação à sua localização atual. Em contraste, window.moveTo() move a janela para
//	   um local absoluto.
func (e *Stage) MoveBy(deltaX, deltaY float64) {
	e.selfWindow.Call("moveBy", deltaX, deltaY)
}

// ResizeBy
//
// English:
//
// The Window.resizeBy() method resizes the current window by a specified amount.
//
//	Input:
//	  xDelta: Number of pixels to grow the window horizontally.
//	  yDelta: Number of pixels to grow the window vertically.
//
// Português:
//
// O método Window.resizeBy() redimensiona a janela atual em um valor especificado.
//
//	Entrada:
//	  xDelta: Número de pixels para aumentar a janela horizontalmente.
//	  yDelta: Número de pixels para aumentar a janela verticalmente.
func (e *Stage) ResizeBy(deltaX, deltaY float64) {
	e.selfWindow.Call("resizeBy", deltaX, deltaY)
}

// ResizeTo
//
// English:
//
// The Window.resizeTo() method dynamically resizes the window.
//
//	Input:
//	  width: An integer representing the new outerWidth in pixels (including scroll bars, title bars, etc).
//	  height: An integer value representing the new outerHeight in pixels (including scroll bars, title bars, etc).
//
// Português:
//
// O método Window.resizeTo() redimensiona dinamicamente a janela.
//
//	Entrada:
//	  width: Um inteiro que representa o novo outerWidth em pixels (incluindo barras de rolagem, barras de título etc.)
//	  height: Um valor inteiro que representa o novo outerHeight em pixels (incluindo barras de rolagem, barras de
//	    título etc.)
func (e *Stage) ResizeTo(width, height float64) {
	e.selfWindow.Call("resizeTo", width, height)
}

// Scroll
//
// English:
//
// Scrolls the window to a particular place in the document.
//
//	Input:
//	  x: Is the pixel along the horizontal axis of the document that you want displayed in the upper left.
//	  y: is the pixel along the vertical axis of the document that you want displayed in the upper left.
//
// Português:
//
// Rola a janela para um local específico no documento.
//
//	Entrada:
//	  x: É o pixel ao longo do eixo horizontal do documento que você deseja exibir no canto superior esquerdo.
//	  y: é o pixel ao longo do eixo vertical do documento que você deseja exibir no canto superior esquerdo.
func (e *Stage) Scroll(x, y float64) {
	e.selfWindow.Call("scroll", x, y)
}

// ScrollBy
//
// English:
//
// Scrolls the document in the window by the given amount.
//
//	Input:
//	  x: Is the horizontal pixel value that you want to scroll by.
//	  y: Is the vertical pixel value that you want to scroll by.
//
// Português:
//
// Rola o documento na janela pela quantidade especificada.
//
//	Entrada:
//	  x: É o valor de pixel horizontal pelo qual você deseja rolar.
//	  y: É o valor de pixel vertical pelo qual você deseja rolar.
func (e *Stage) ScrollBy(x, y float64) {
	e.selfWindow.Call("scrollBy", x, y)
}

// ScrollTo
//
// English:
//
// Scrolls to a particular set of coordinates in the document.
//
//	Input:
//	  x: Is the pixel along the horizontal axis of the document that you want displayed in the upper left.
//	  y: Is the pixel along the vertical axis of the document that you want displayed in the upper left.
//
// Português:
//
// Rola para um determinado conjunto de coordenadas no documento.
//
//	Entrada:
//	  x: É o pixel ao longo do eixo horizontal do documento que você deseja exibir no canto superior esquerdo.
//	  y: É o pixel ao longo do eixo vertical do documento que você deseja exibir no canto superior esquerdo.
func (e *Stage) ScrollTo(x, y float64) {
	e.selfWindow.Call("scrollTo", x, y)
}

//------------------------------------------------------------------------------------------

// GetName
//
// English:
//
// Gets the name of the window's browsing context.
//
// Português:
//
// Obtém o nome do contexto de navegação da janela.
func (e *Stage) GetName() (name string) {
	return e.selfWindow.Get("name").String()
}

// GetInnerWidth
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
func (e *Stage) GetInnerWidth() (width float64) {
	return e.selfWindow.Get("innerWidth").Float()
}

// GetInnerHeight
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
func (e *Stage) GetInnerHeight() (height float64) {
	return e.selfWindow.Get("innerHeight").Float()
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
func (e *Stage) GetFrameLength() (length int) {
	return e.selfWindow.Get("length").Int()
}

// GetIsClosed
//
// English:
//
// # The Window.closed read-only property indicates whether the referenced window is closed or not
//
// Português:
//
// A propriedade somente leitura Window.closed indica se a janela referenciada está fechada ou não
func (e *Stage) GetIsClosed() (closed bool) {
	return e.selfWindow.Get("closed").Bool()
}

// GetOuterHeight
//
// English:
//
// The Window.outerHeight read-only property returns the height in pixels of the whole browser window, including any
// sidebar, window chrome, and window-resizing borders/handles.
//
//	Notes:
//	  * To change the size of a window, see window.resizeBy() and window.resizeTo().
//
// Português:
//
// A propriedade somente leitura Window.outerHeight retorna a altura em pixels de toda a janela do navegador, incluindo
// qualquer barra lateral, cromo de janela e alças de borda de redimensionamento de janela.
//
//	Notas:
//	  * Para alterar o tamanho de uma janela, consulte window.resizeBy() e window.resizeTo().
func (e *Stage) GetOuterHeight() (outerHeight float64) {
	return e.selfWindow.Get("outerHeight").Float()
}

// GetOuterWidth
//
// English:
//
// Window.outerWidth read-only property returns the width of the outside of the browser window. It represents the width
// of the whole browser window including sidebar (if expanded), window chrome and window resizing borders/handles.
//
//	Notes:
//	 * To change the size of a window, see window.resizeBy() and window.resizeTo().
//
// Português:
//
// A propriedade somente leitura Window.outerWidth retorna a largura da parte externa da janela do navegador.
// Ele representa a largura de toda a janela do navegador, incluindo barra lateral (se expandida), cromo da janela e
// alças de bordas de redimensionamento de janela.
//
//	Notas:
//	 * Para alterar o tamanho de uma janela, consulte window.resizeBy() e window.resizeTo().
func (e *Stage) GetOuterWidth() (outerWidth float64) {
	return e.selfWindow.Get("outerWidth").Float()
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
func (e *Stage) GetScrollX() (scrollX float64) {
	return e.selfWindow.Get("scrollX").Float()
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
func (e *Stage) GetScrollY() (scrollY float64) {
	return e.selfWindow.Get("scrollY").Float()
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
func (e *Stage) GetScreenX() (screenX float64) {
	return e.selfWindow.Get("screenX").Float()
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
func (e *Stage) GetScreenY() (screenY float64) {
	return e.selfWindow.Get("screenY").Float()
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
//   - The opener can be omitted by specifying rel=noopener on a link, or passing noopener in the windowFeatures
//     parameter.
//   - Windows opened because of links with a target of _blank don't get an opener, unless explicitly requested with
//     rel=opener.
//   - Having a Cross-Origin-Opener-Policy header with a value of same-origin prevents setting opener.
//     Since the new window is loaded in a different browsing context, it won't have a reference to the opening window.
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
//   - O opener pode ser omitido especificando rel=noopener em um link ou passando noopener no parâmetro
//     windowFeatures.
//   - O Windows aberto devido a links com um destino de _blank não obtém um abridor, a menos que solicitado
//     explicitamente com rel=opener.
//   - Ter um cabeçalho Cross-Opener-Policy com um valor de mesma origem impede a configuração do abridor.
//     Como a nova janela é carregada em um contexto de navegação diferente, ela não terá uma referência à janela de
//     abertura.
func (e *Stage) GetOpener() (opener js.Value) {
	return e.selfWindow.Get("opener")
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
func (e *Stage) GetParent() (parent js.Value) {
	return e.selfWindow.Get("parent")
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
func (e *Stage) GetScreen() (parent js.Value) {
	return e.selfWindow.Get("screen")
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
func (e *Stage) GetScrollBars() (scrollbars js.Value) {
	return e.selfWindow.Get("scrollbars")
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
func (e *Stage) GetStatusBar() (statusbar js.Value) {
	return e.selfWindow.Get("statusbar")
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
func (e *Stage) GetTop() (top js.Value) {
	return e.selfWindow.Get("top")
}
