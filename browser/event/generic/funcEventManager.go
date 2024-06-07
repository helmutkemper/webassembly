package generic

import "syscall/js"

type EventName string

func (e EventName) String() string {
	return string(e)
}

const (
	KEventAbort             EventName = "abort"
	KEventAutocomplete      EventName = "autocomplete"
	KEventAutocompleteError EventName = "autocompleteerror"
	KEventBlur              EventName = "blur"
	KEventCancel            EventName = "cancel"
	KEventCanplay           EventName = "canplay"
	KEventCanplaythrough    EventName = "canplaythrough"
	KEventChange            EventName = "change"
	KEventClick             EventName = "click"
	KEventClose             EventName = "close"
	KEventContextmenu       EventName = "contextmenu"
	KEventCuechange         EventName = "cuechange"
	KEventDblclick          EventName = "dblclick"
	KEventDrag              EventName = "drag"
	KEventDragend           EventName = "dragend"
	KEventDragenter         EventName = "dragenter"
	KEventDragleave         EventName = "dragleave"
	KEventDragover          EventName = "dragover"
	KEventDragstart         EventName = "dragstart"
	KEventDrop              EventName = "drop"
	KEventDurationchange    EventName = "durationchange"
	KEventEmptied           EventName = "emptied"
	KEventEnded             EventName = "ended"
	KEventError             EventName = "error"
	KEventFocus             EventName = "focus"
	KEventInput             EventName = "input"
	KEventInvalid           EventName = "invalid"
	KEventKeydown           EventName = "keydown"
	KEventKeypress          EventName = "keypress"
	KEventKeyup             EventName = "keyup"
	KEventLoad              EventName = "load"
	KEventLoadeddata        EventName = "loadeddata"
	KEventLoadedmetadata    EventName = "loadedmetadata"
	KEventLoadstart         EventName = "loadstart"
	KEventMousedown         EventName = "mousedown"
	KEventMouseenter        EventName = "mouseenter"
	KEventMouseleave        EventName = "mouseleave"
	KEventMousemove         EventName = "mousemove"
	KEventMouseout          EventName = "mouseout"
	KEventMouseover         EventName = "mouseover"
	KEventMouseup           EventName = "mouseup"
	KEventMousewheel        EventName = "mousewheel"
	KEventPause             EventName = "pause"
	KEventPlay              EventName = "play"
	KEventPlaying           EventName = "playing"
	KEventProgress          EventName = "progress"
	KEventRatechange        EventName = "ratechange"
	KEventReset             EventName = "reset"
	KEventResize            EventName = "resize"
	KEventScroll            EventName = "scroll"
	KEventSeeked            EventName = "seeked"
	KEventSeeking           EventName = "seeking"
	KEventSelect            EventName = "select"
	KEventShow              EventName = "show"
	KEventSort              EventName = "sort"
	KEventStalled           EventName = "stalled"
	KEventSubmit            EventName = "submit"
	KEventSuspend           EventName = "suspend"
	KEventTimeupdate        EventName = "timeupdate"
	KEventToggle            EventName = "toggle"
	KEventVolumechange      EventName = "volumechange"
	KEventWaiting           EventName = "waiting"
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
	data.Accesskey = event.GetAccesskey()
	data.Autocapitalize = event.GetAutocapitalize()
	data.Autofocus = event.GetAutofocus()
	data.Class = event.GetClass()
	data.Contenteditable = event.GetContenteditable()
	data.Dir = event.GetDir()
	data.Draggable = event.GetDraggable()
	data.Enterkeyhint = event.GetEnterkeyhint()
	data.Exportparts = event.GetExportparts()
	data.Hidden = event.GetHidden()
	data.Id = event.GetId()
	data.Inert = event.GetInert()
	data.InputMode = event.GetInputMode()
	data.Is = event.GetIs()
	data.ItemId = event.GetItemId()
	data.Itemprop = event.GetItemprop()
	data.ItemRef = event.GetItemRef()
	data.ItemScope = event.GetItemScope()
	data.ItemType = event.GetItemType()
	data.Lang = event.GetLang()
	data.Nonce = event.GetNonce()
	data.Popover = event.GetPopover()
	data.Role = event.GetRole()
	data.Spellcheck = event.GetSpellcheck()
	data.Style = event.GetStyle()
	data.TabIndex = event.GetTabIndex()
	data.Title = event.GetTitle()
	data.Translate = event.GetTranslate()
	data.This = this

	return
}
