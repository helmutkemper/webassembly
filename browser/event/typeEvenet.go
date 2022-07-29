// Package event
//
// English:
//
// This package returns the variables available for generic events.
//
// Português:
//
// Este pacote retorna as variáveis disponíveis para eventos genéricos.
//
// Sorce: / Fonte:
// https://developer.mozilla.org/en-US/docs/Web/API/Event
package event

import (
	"syscall/js"
)

type EventName string

func (e EventName) String() string {
	return string(e)
}

const (
	// KEventCanPlay
	//
	// English:
	//
	// The browser can play the media, but estimates that not enough data has been loaded to play the media up to its end
	// without having to stop for further buffering of content.
	//
	// Português:
	//
	// O navegador pode reproduzir a mídia, mas estima que não foram carregados dados suficientes para reproduzir a mídia
	// até o fim sem ter que parar para mais armazenamento em buffer do conteúdo.
	KEventCanPlay EventName = "canplay"

	// KEventCanPlayThrough
	//
	// English:
	//
	// The browser estimates it can play the media up to its end without stopping for content buffering.
	//
	// Português:
	//
	// O navegador estima que pode reproduzir a mídia até o fim sem parar para o buffer de conteúdo.
	KEventCanPlayThrough EventName = "canplaythrough"

	// KEventComplete
	//
	// English:
	//
	// The rendering of an OfflineAudioContext is terminated.
	//
	// Português:
	//
	// A renderização de um OfflineAudioContext é encerrada.
	KEventComplete EventName = "complete"

	// KEventDurationChange
	//
	// English:
	//
	// The duration attribute has been updated.
	//
	// Português:
	//
	// O atributo de duração foi atualizado.
	KEventDurationChange EventName = "durationchange"

	// KEventEmptied
	//
	// English:
	//
	// The media has become empty; for example, this event is sent if the media has already been loaded (or partially
	// loaded), and the load() method is called to reload it.
	//
	// Português:
	//
	// A mídia ficou vazia; por exemplo, este evento é enviado se a mídia já foi carregada (ou parcialmente carregada) e
	// o método load() é chamado para recarregá-la.
	KEventEmptied EventName = "emptied"

	// KEventEnded
	//
	// English:
	//
	// Playback has stopped because the end of the media was reached.
	//
	// Português:
	//
	// A reprodução foi interrompida porque o fim da mídia foi atingido.
	KEventEnded EventName = "ended"

	// KEventLoadedData
	//
	// English:
	//
	// The first frame of the media has finished loading.
	//
	// Português:
	//
	// O primeiro quadro da mídia terminou de ser carregado.
	KEventLoadedData EventName = "loadeddata"

	// KEventLoadedMetadata
	//
	// English:
	//
	// The metadata has been loaded.
	//
	// Português:
	//
	// Os metadados foram carregados.
	KEventLoadedMetadata EventName = "loadedmetadata"

	// KEventPause
	//
	// English:
	//
	// Playback has been paused.
	//
	// Português:
	//
	// A reprodução foi pausada.
	KEventPause EventName = "pause"

	// KEventPlay
	//
	// English:
	//
	// Playback has begun.
	//
	// Português:
	//
	// A reprodução começou.
	KEventPlay EventName = "play"

	// KEventPlaying
	//
	// English:
	//
	// Playback is ready to start after having been paused or delayed due to lack of data.
	//
	// Português:
	//
	// A reprodução está pronta para iniciar após ter sido pausada ou atrasada devido à falta de dados.
	KEventPlaying EventName = "playing"

	// KEventProgress
	//
	// English:
	//
	// Fired periodically as the browser loads a resource.
	//
	// Português:
	//
	// Acionado periodicamente conforme o navegador carrega um recurso.
	KEventProgress EventName = "progress"

	// KEventRateChange
	//
	// English:
	//
	// The playback rate has changed.
	//
	// Português:
	//
	// A taxa de reprodução mudou.
	KEventRateChange EventName = "ratechange"

	// KEventSeeked
	//
	// English:
	//
	// A seek operation completed.
	//
	// Português:
	//
	// Uma operação de busca concluída.
	KEventSeeked EventName = "seeked"

	// KEventSeeking
	//
	// English:
	//
	// A seek operation began.
	//
	// Português:
	//
	// Uma operação de busca começou.
	KEventSeeking EventName = "seeking"

	// KEventStalled
	//
	// English:
	//
	// The user agent is trying to fetch media data, but data is unexpectedly not forthcoming.
	//
	// Português:
	//
	// O agente do usuário está tentando buscar dados de mídia, mas os dados inesperadamente não estão disponíveis.
	KEventStalled EventName = "stalled"

	// KEventSuspend
	//
	// English:
	//
	// Media data loading has been suspended.
	//
	// Português:
	//
	// O carregamento de dados de mídia foi suspenso.
	KEventSuspend EventName = "suspend"

	// KEventTimeUpdate
	//
	// English:
	//
	// The time indicated by the currentTime attribute has been updated.
	//
	// Português:
	//
	// A hora indicada pelo atributo currentTime foi atualizada.
	KEventTimeUpdate EventName = "timeupdate"

	// KEventVolumeChange
	//
	// English:
	//
	// The volume has changed.
	//
	// Português:
	//
	// O volume mudou.
	KEventVolumeChange EventName = "volumechange"

	// KEventWaiting
	//
	// English:
	//
	// Playback has stopped because of a temporary lack of data.
	//
	// Português:
	//
	// A reprodução parou devido a uma falta temporária de dados.
	KEventWaiting EventName = "waiting"
)

// EventManager
//
// English:
//
// Capture event information and format to Golang
//
//	Output:
//	  data: list with all the information provided by the browser.
//
// Português:
//
// Captura as informações do evento e formata para o Golang
//
//	Saída:
//	  data: lista com todas as informações fornecidas pelo navegador.
func EventManager(name EventName, this js.Value, args []js.Value) (data Data) {
	var event = Event{}
	event.Object = args[0]

	data.EventName = name
	data.This = this
	data.Bubbles = event.GetBubbles()
	data.Cancelable = event.GetCancelable()
	data.Composed = event.GetComposed()
	data.CurrentTarget = event.GetCurrentTarget()
	data.DefaultPrevented = event.GetDefaultPrevented()
	data.EventPhase = event.GetEventPhase()
	data.IsTrusted = event.GetIsTrusted()
	data.Target = event.GetTarget()
	data.TimeStamp = event.GetTimeStamp()
	data.Type = event.GetType()

	return
}

type Event struct {
	Object js.Value
}

// GetBubbles
//
// English:
//
// A boolean value indicating whether or not the event bubbles up through the DOM.
//
// Português:
//
// Um valor booleano que indica se o evento bubbles up ou não pelo DOM.
func (e *Event) GetBubbles() (bubbles bool) {
	return e.Object.Get("bubbles").Bool()
}

// GetCancelable
//
// English:
//
// A boolean value indicating whether the event is cancelable.
//
// Português:
//
// Um valor booleano que indica se o evento pode ser cancelado.
func (e *Event) GetCancelable() (cancelable bool) {
	return e.Object.Get("cancelable").Bool()
}

// GetComposed
//
// English:
//
// A boolean indicating whether or not the event can bubble across the boundary between the shadow DOM and the
// regular DOM.
//
// Português:
//
// Um booleano que indica se o evento pode ou não atravessar o limite entre o Shadow DOM e o DOM regular.
func (e *Event) GetComposed() (composed bool) {
	return e.Object.Get("composed").Bool()
}

// GetCurrentTarget
//
// English:
//
// A reference to the currently registered target for the event.
//
// This is the object to which the event is currently slated to be sent. It's possible this has been changed along the
// way through retargeting.
//
// Português:
//
// Uma referência ao destino atualmente registrado para o evento.
//
// Este é o objeto para o qual o evento está programado para ser enviado. É possível que isso tenha sido alterado ao
// longo do caminho por meio do redirecionamento.
func (e *Event) GetCurrentTarget() (currentTarget js.Value) {
	return e.Object.Get("currentTarget")
}

// GetDefaultPrevented
//
// English:
//
// Indicates whether or not the call to event.preventDefault() canceled the event.
//
// Português:
//
// Indica se a chamada para event.preventDefault() cancelou ou não o evento.
func (e *Event) GetDefaultPrevented() (defaultPrevented bool) {
	return e.Object.Get("defaultPrevented").Bool()
}

// GetEventPhase
//
// English:
//
// The eventPhase read-only property of the Event interface indicates which phase of the event flow is currently
// being evaluated.
//
// Português:
//
// A propriedade somente leitura eventPhase da interface Event indica qual fase do fluxo de eventos está sendo avaliada
// no momento.
func (e *Event) GetEventPhase() (eventPhase Phase) {
	return Phase(e.Object.Get("eventPhase").Int())
}

// GetIsTrusted
//
// English:
//
// Indicates whether or not the event was initiated by the browser (after a user click, for instance) or by a script
// (using an event creation method, for example).
//
// Português:
//
// Indica se o evento foi ou não iniciado pelo navegador (após um clique do usuário, por exemplo) ou por um script
// (usando um método de criação de eventos, por exemplo).
func (e *Event) GetIsTrusted() (isTrusted bool) {
	return e.Object.Get("isTrusted").Bool()
}

// GetTarget
//
// English:
//
// A reference to the object to which the event was originally dispatched.
//
// Português:
//
// Uma referência ao objeto para o qual o evento foi originalmente despachado.
func (e *Event) GetTarget() (target js.Value) {
	return e.Object.Get("target")
}

// GetTimeStamp
//
// English:
//
// This value is the number of milliseconds elapsed from the beginning of the time origin until the event was created.
// If the global object is Window, the time origin is the moment the user clicked on the link, or the script that
// initiated the loading of the document. In a worker, the time origin is the moment of creation of the worker.
//
// Português:
//
// Esse valor é o número de milissegundos decorridos desde o início da origem do tempo até a criação do evento.
// Se o objeto global for Window, a origem do tempo é o momento em que o usuário clicou no link, ou o script que
// iniciou o carregamento do documento. Em um trabalhador, a origem do tempo é o momento de criação do trabalhador.
func (e *Event) GetTimeStamp() (timeStamp float64) {
	return e.Object.Get("timeStamp").Float()
}

// GetType
//
// English:
//
// Returns a string containing the event's type. It is set when the event is constructed and is the name commonly used
// to refer to the specific event, such as click, load, or error.
//
// Português:
//
// Retorna uma string contendo o tipo do evento. Ele é definido quando o evento é construído e é o nome comumente usado
// para se referir ao evento específico, como clique, carregamento ou erro.
func (e *Event) GetType() (typeString string) {
	return e.Object.Get("type").String()
}

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

	// Bubbles
	//
	// English:
	//
	// A boolean value indicating whether or not the event bubbles up through the DOM.
	//
	// Português:
	//
	// Um valor booleano que indica se o evento bubbles up ou não pelo DOM.
	Bubbles bool

	// GetCancelable
	//
	// English:
	//
	// A boolean value indicating whether the event is cancelable.
	//
	// Português:
	//
	// Um valor booleano que indica se o evento pode ser cancelado.
	Cancelable bool

	// Composed
	//
	// English:
	//
	// A boolean indicating whether or not the event can bubble across the boundary between the shadow DOM and the
	// regular DOM.
	//
	// Português:
	//
	// Um booleano que indica se o evento pode ou não atravessar o limite entre o Shadow DOM e o DOM regular.
	Composed bool

	// CurrentTarget
	//
	// English:
	//
	// A reference to the currently registered target for the event.
	//
	// This is the object to which the event is currently slated to be sent. It's possible this has been changed along the
	// way through retargeting.
	//
	// Português:
	//
	// Uma referência ao destino atualmente registrado para o evento.
	//
	// Este é o objeto para o qual o evento está programado para ser enviado. É possível que isso tenha sido alterado ao
	// longo do caminho por meio do redirecionamento.
	CurrentTarget js.Value

	// DefaultPrevented
	//
	// English:
	//
	// Indicates whether or not the call to event.preventDefault() canceled the event.
	//
	// Português:
	//
	// Indica se a chamada para event.preventDefault() cancelou ou não o evento.
	DefaultPrevented bool

	// EventPhase
	//
	// English:
	//
	// The eventPhase read-only property of the Event interface indicates which phase of the event flow is currently
	// being evaluated.
	//
	// Português:
	//
	// A propriedade somente leitura eventPhase da interface Event indica qual fase do fluxo de eventos está sendo avaliada
	// no momento.
	EventPhase Phase

	// IsTrusted
	//
	// English:
	//
	// Indicates whether or not the event was initiated by the browser (after a user click, for instance) or by a script
	// (using an event creation method, for example).
	//
	// Português:
	//
	// Indica se o evento foi ou não iniciado pelo navegador (após um clique do usuário, por exemplo) ou por um script
	// (usando um método de criação de eventos, por exemplo).
	IsTrusted bool

	// Target
	//
	// English:
	//
	// A reference to the object to which the event was originally dispatched.
	//
	// Português:
	//
	// Uma referência ao objeto para o qual o evento foi originalmente despachado.
	Target js.Value

	// TimeStamp
	//
	// English:
	//
	// This value is the number of milliseconds elapsed from the beginning of the time origin until the event was created.
	// If the global object is Window, the time origin is the moment the user clicked on the link, or the script that
	// initiated the loading of the document. In a worker, the time origin is the moment of creation of the worker.
	//
	// Português:
	//
	// Esse valor é o número de milissegundos decorridos desde o início da origem do tempo até a criação do evento.
	// Se o objeto global for Window, a origem do tempo é o momento em que o usuário clicou no link, ou o script que
	// iniciou o carregamento do documento. Em um trabalhador, a origem do tempo é o momento de criação do trabalhador.
	TimeStamp float64

	// Type
	//
	// English:
	//
	// Returns a string containing the event's type. It is set when the event is constructed and is the name commonly used
	// to refer to the specific event, such as click, load, or error.
	//
	// Português:
	//
	// Retorna uma string contendo o tipo do evento. Ele é definido quando o evento é construído e é o nome comumente usado
	// para se referir ao evento específico, como clique, carregamento ou erro.
	Type string
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
