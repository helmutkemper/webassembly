package html

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"syscall/js"
)

type commonEvents struct {
	selfElement *js.Value

	fnAbort                    *js.Func
	fnAuxclick                 *js.Func
	fnBeforeinput              *js.Func
	fnBeforeMatch              *js.Func
	fnBeforetoggle             *js.Func
	fnCancel                   *js.Func
	fnCanplay                  *js.Func
	fnCanplaythrough           *js.Func
	fnChange                   *js.Func
	fnClick                    *js.Func
	fnClose                    *js.Func
	fnContextlost              *js.Func
	fnContextmenu              *js.Func
	fnContextrestored          *js.Func
	fnCopy                     *js.Func
	fnCuechange                *js.Func
	fnCut                      *js.Func
	fnDblclick                 *js.Func
	fnDrag                     *js.Func
	fnDragend                  *js.Func
	fnDragenter                *js.Func
	fnDragleave                *js.Func
	fnDragover                 *js.Func
	fnDragstart                *js.Func
	fnDrop                     *js.Func
	fnDurationchange           *js.Func
	fnEmptied                  *js.Func
	fnEnded                    *js.Func
	fnFormdata                 *js.Func
	fnInput                    *js.Func
	fnInvalid                  *js.Func
	fnKeydown                  *js.Func
	fnKeypress                 *js.Func
	fnKeyup                    *js.Func
	fnLoadeddata               *js.Func
	fnLoadedmetadata           *js.Func
	fnLoadstart                *js.Func
	fnMousedown                *js.Func
	fnMouseenter               *js.Func
	fnMouseleave               *js.Func
	fnMousemove                *js.Func
	fnMouseout                 *js.Func
	fnMouseover                *js.Func
	fnMouseup                  *js.Func
	fnPaste                    *js.Func
	fnPause                    *js.Func
	fnPlay                     *js.Func
	fnPlaying                  *js.Func
	fnProgress                 *js.Func
	fnRatechange               *js.Func
	fnReset                    *js.Func
	fnScrollend                *js.Func
	fnSecuritypolicyviolation  *js.Func
	fnSeeked                   *js.Func
	fnSeeking                  *js.Func
	fnSelect                   *js.Func
	fnSlotchange               *js.Func
	fnStalled                  *js.Func
	fnSubmit                   *js.Func
	fnSuspend                  *js.Func
	fnTimeupdate               *js.Func
	fnToggle                   *js.Func
	fnVolumechange             *js.Func
	fnWaiting                  *js.Func
	fnWebkitanimationend       *js.Func
	fnWebkitanimationiteration *js.Func
	fnWebkitanimationstart     *js.Func
	fnWebkittransitionend      *js.Func
	fnWheel                    *js.Func
	fnBlur                     *js.Func
	fnError                    *js.Func
	fnFocus                    *js.Func
	fnLoad                     *js.Func
	fnResize                   *js.Func
	fnScroll                   *js.Func
	fnAfterprint               *js.Func
	fnBeforeprint              *js.Func
	fnBeforeunload             *js.Func
	fnHashchange               *js.Func
	fnLanguagechange           *js.Func
	fnMessage                  *js.Func
	fnMessageerror             *js.Func
	fnOffline                  *js.Func
	fnOnline                   *js.Func
	fnPageswap                 *js.Func
	fnPagehide                 *js.Func
	fnPagereveal               *js.Func
	fnPageshow                 *js.Func
	fnPopstate                 *js.Func
	fnRejectionhandled         *js.Func
	fnStorage                  *js.Func
	fnUnhandledrejection       *js.Func
	fnUnload                   *js.Func
	fnReadystatechange         *js.Func
	fnVisibilitychange         *js.Func
}

func (e *commonEvents) GetUuidStr() (uuidStr string) {
	uId, err := uuid.NewUUID()
	if err != nil {
		err = fmt.Errorf("commonEvents.NewUUID().error: %v", err)
		return
	}
	uuidStr = uId.String()
	return
}

func (e *commonEvents) AddListenerAbort(genericEvent chan generic.Data) {
	if e.fnAbort != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventAbort, this, args)
		return nil
	})
	e.fnAbort = &fn

	e.selfElement.Call(
		"addEventListener",
		"abort",
		*e.fnAbort,
	)
}

func (e *commonEvents) RemoveListenerAbort() {
	if e.fnAbort == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"abort",
		*e.fnAbort,
	)
	e.fnAbort = nil

}

func (e *commonEvents) AddListenerAuxclick(genericEvent chan generic.Data) {
	if e.fnAuxclick != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventAuxclick, this, args)
		return nil
	})
	e.fnAuxclick = &fn

	e.selfElement.Call(
		"addEventListener",
		"auxclick",
		*e.fnAuxclick,
	)
}

func (e *commonEvents) RemoveListenerAuxclick() {
	if e.fnAuxclick == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"auxclick",
		*e.fnAuxclick,
	)
	e.fnAuxclick = nil
}

func (e *commonEvents) AddListenerBeforeinput(genericEvent chan generic.Data) {
	if e.fnBeforeinput != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventBeforeinput, this, args)
		return nil
	})
	e.fnBeforeinput = &fn

	e.selfElement.Call(
		"addEventListener",
		"beforeinput",
		*e.fnBeforeinput,
	)
}

func (e *commonEvents) RemoveListenerBeforeinput() {
	if e.fnBeforeinput == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"beforeinput",
		*e.fnBeforeinput,
	)
	e.fnBeforeinput = nil
}

func (e *commonEvents) AddListenerBeforematch(genericEvent chan generic.Data) {
	if e.fnBeforeMatch != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventBeforematch, this, args)
		return nil
	})
	e.fnBeforeMatch = &fn

	e.selfElement.Call(
		"addEventListener",
		"beforematch",
		*e.fnBeforeMatch,
	)
}

func (e *commonEvents) RemoveListenerBeforematch() {
	if e.fnBeforeMatch == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"beforematch",
		*e.fnBeforeMatch,
	)
	e.fnBeforeMatch = nil
}

func (e *commonEvents) AddListenerBeforetoggle(genericEvent chan generic.Data) {
	if e.fnBeforetoggle != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventBeforetoggle, this, args)
		return nil
	})
	e.fnBeforetoggle = &fn

	e.selfElement.Call(
		"addEventListener",
		"beforetoggle",
		*e.fnBeforetoggle,
	)
}

func (e *commonEvents) RemoveListenerBeforetoggle() {
	if e.fnBeforetoggle == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"beforetoggle",
		*e.fnBeforetoggle,
	)
	e.fnBeforetoggle = nil
}

func (e *commonEvents) AddListenerCancel(genericEvent chan generic.Data) {
	if e.fnCancel != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventCancel, this, args)
		return nil
	})
	e.fnCancel = &fn

	e.selfElement.Call(
		"addEventListener",
		"cancel",
		*e.fnCancel,
	)

}

func (e *commonEvents) RemoveListenerCancel() {
	if e.fnCancel == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"cancel",
		*e.fnCancel,
	)
	e.fnCancel = nil

}

func (e *commonEvents) AddListenerCanplay(genericEvent chan generic.Data) {
	if e.fnCanplay != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventCanplay, this, args)
		return nil
	})
	e.fnCanplay = &fn

	e.selfElement.Call(
		"addEventListener",
		"canplay",
		*e.fnCanplay,
	)

}

func (e *commonEvents) RemoveListenerCanplay() {
	if e.fnCanplay == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"canplay",
		*e.fnCanplay,
	)
	e.fnCanplay = nil

}

func (e *commonEvents) AddListenerCanplaythrough(genericEvent chan generic.Data) {
	if e.fnCanplaythrough != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventCanplaythrough, this, args)
		return nil
	})
	e.fnCanplaythrough = &fn

	e.selfElement.Call(
		"addEventListener",
		"canplaythrough",
		*e.fnCanplaythrough,
	)

}

func (e *commonEvents) RemoveListenerCanplaythrough() {
	if e.fnCanplaythrough == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"canplaythrough",
		*e.fnCanplaythrough,
	)
	e.fnCanplaythrough = nil

}

func (e *commonEvents) AddListenerChange(genericEvent chan generic.Data) {
	if e.fnChange != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventChange, this, args)
		return nil
	})
	e.fnChange = &fn

	e.selfElement.Call(
		"addEventListener",
		"change",
		*e.fnChange,
	)

}

func (e *commonEvents) RemoveListenerChange() {
	if e.fnChange == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"change",
		*e.fnChange,
	)
	e.fnChange = nil

}

func (e *commonEvents) AddListenerClick(genericEvent chan generic.Data) {
	if e.fnClick != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventClick, this, args)
		return nil
	})
	e.fnClick = &fn

	e.selfElement.Call(
		"addEventListener",
		"click",
		*e.fnClick,
	)

}

func (e *commonEvents) RemoveListenerClick() {
	if e.fnClick == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"click",
		*e.fnClick,
	)
	e.fnClick = nil

}

func (e *commonEvents) AddListenerClose(genericEvent chan generic.Data) {
	if e.fnClose != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventClose, this, args)
		return nil
	})
	e.fnClose = &fn

	e.selfElement.Call(
		"addEventListener",
		"close",
		*e.fnClose,
	)

}

func (e *commonEvents) RemoveListenerClose() {
	if e.fnClose == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"close",
		*e.fnClose,
	)
	e.fnClose = nil

}

func (e *commonEvents) AddListenerContextlost(genericEvent chan generic.Data) {
	if e.fnContextlost != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventContextlost, this, args)
		return nil
	})
	e.fnContextlost = &fn

	e.selfElement.Call(
		"addEventListener",
		"contextlost",
		*e.fnContextlost,
	)

}

func (e *commonEvents) RemoveListenerContextlost() {
	if e.fnContextlost == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"contextlost",
		*e.fnContextlost,
	)
	e.fnContextlost = nil

}

func (e *commonEvents) AddListenerContextmenu(genericEvent chan generic.Data) {
	if e.fnContextmenu != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventContextmenu, this, args)
		return nil
	})
	e.fnContextmenu = &fn

	e.selfElement.Call(
		"addEventListener",
		"contextmenu",
		*e.fnContextmenu,
	)

}

func (e *commonEvents) RemoveListenerContextmenu() {
	if e.fnContextmenu == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"contextmenu",
		*e.fnContextmenu,
	)
	e.fnContextmenu = nil

}

func (e *commonEvents) AddListenerContextrestored(genericEvent chan generic.Data) {
	if e.fnContextrestored != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventContextrestored, this, args)
		return nil
	})
	e.fnContextrestored = &fn

	e.selfElement.Call(
		"addEventListener",
		"contextrestored",
		*e.fnContextrestored,
	)

}

func (e *commonEvents) RemoveListenerContextrestored() {
	if e.fnContextrestored == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"contextrestored",
		*e.fnContextrestored,
	)
	e.fnContextrestored = nil

}

func (e *commonEvents) AddListenerCopy(genericEvent chan generic.Data) {
	if e.fnCopy != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventCopy, this, args)
		return nil
	})
	e.fnCopy = &fn

	e.selfElement.Call(
		"addEventListener",
		"copy",
		*e.fnCopy,
	)

}

func (e *commonEvents) RemoveListenerCopy() {
	if e.fnCopy == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"copy",
		*e.fnCopy,
	)
	e.fnCopy = nil

}

func (e *commonEvents) AddListenerCuechange(genericEvent chan generic.Data) {
	if e.fnCuechange != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventCuechange, this, args)
		return nil
	})
	e.fnCuechange = &fn

	e.selfElement.Call(
		"addEventListener",
		"cuechange",
		*e.fnCuechange,
	)

}

func (e *commonEvents) RemoveListenerCuechange() {
	if e.fnCuechange == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"cuechange",
		*e.fnCuechange,
	)
	e.fnCuechange = nil

}

func (e *commonEvents) AddListenerCut(genericEvent chan generic.Data) {
	if e.fnCut != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventCut, this, args)
		return nil
	})
	e.fnCut = &fn

	e.selfElement.Call(
		"addEventListener",
		"cut",
		*e.fnCut,
	)

}

func (e *commonEvents) RemoveListenerCut() {
	if e.fnCut == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"cut",
		*e.fnCut,
	)
	e.fnCut = nil

}

func (e *commonEvents) AddListenerDblclick(genericEvent chan generic.Data) {
	if e.fnDblclick != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventDblclick, this, args)
		return nil
	})
	e.fnDblclick = &fn

	e.selfElement.Call(
		"addEventListener",
		"dblclick",
		*e.fnDblclick,
	)

}

func (e *commonEvents) RemoveListenerDblclick() {
	if e.fnDblclick == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"dblclick",
		*e.fnDblclick,
	)
	e.fnDblclick = nil

}

func (e *commonEvents) AddListenerDrag(genericEvent chan generic.Data) {
	if e.fnDrag != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventDrag, this, args)
		return nil
	})
	e.fnDrag = &fn

	e.selfElement.Call(
		"addEventListener",
		"drag",
		*e.fnDrag,
	)

}

func (e *commonEvents) RemoveListenerDrag() {
	if e.fnDrag == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"drag",
		*e.fnDrag,
	)
	e.fnDrag = nil

}

func (e *commonEvents) AddListenerDragend(genericEvent chan generic.Data) {
	if e.fnDragend != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventDragend, this, args)
		return nil
	})
	e.fnDragend = &fn

	e.selfElement.Call(
		"addEventListener",
		"dragend",
		*e.fnDragend,
	)

}

func (e *commonEvents) RemoveListenerDragend() {
	if e.fnDragend == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"dragend",
		*e.fnDragend,
	)
	e.fnDragend = nil

}

func (e *commonEvents) AddListenerDragenter(genericEvent chan generic.Data) {
	if e.fnDragenter != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventDragenter, this, args)
		return nil
	})
	e.fnDragenter = &fn

	e.selfElement.Call(
		"addEventListener",
		"dragenter",
		*e.fnDragenter,
	)

}

func (e *commonEvents) RemoveListenerDragenter() {
	if e.fnDragenter == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"dragenter",
		*e.fnDragenter,
	)
	e.fnDragenter = nil

}

func (e *commonEvents) AddListenerDragleave(genericEvent chan generic.Data) {
	if e.fnDragleave != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventDragleave, this, args)
		return nil
	})
	e.fnDragleave = &fn

	e.selfElement.Call(
		"addEventListener",
		"dragleave",
		*e.fnDragleave,
	)

}

func (e *commonEvents) RemoveListenerDragleave() {
	if e.fnDragleave == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"dragleave",
		*e.fnDragleave,
	)
	e.fnDragleave = nil

}

func (e *commonEvents) AddListenerDragover(genericEvent chan generic.Data) {
	if e.fnDragover != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventDragover, this, args)
		return nil
	})
	e.fnDragover = &fn

	e.selfElement.Call(
		"addEventListener",
		"dragover",
		*e.fnDragover,
	)

}

func (e *commonEvents) RemoveListenerDragover() {
	if e.fnDragover == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"dragover",
		*e.fnDragover,
	)
	e.fnDragover = nil

}

func (e *commonEvents) AddListenerDragstart(genericEvent chan generic.Data) {
	if e.fnDragstart != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventDragstart, this, args)
		return nil
	})
	e.fnDragstart = &fn

	e.selfElement.Call(
		"addEventListener",
		"dragstart",
		*e.fnDragstart,
	)

}

func (e *commonEvents) RemoveListenerDragstart() {
	if e.fnDragstart == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"dragstart",
		*e.fnDragstart,
	)
	e.fnDragstart = nil

}

func (e *commonEvents) AddListenerDrop(genericEvent chan generic.Data) {
	if e.fnDrop != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventDrop, this, args)
		return nil
	})
	e.fnDrop = &fn

	e.selfElement.Call(
		"addEventListener",
		"drop",
		*e.fnDrop,
	)

}

func (e *commonEvents) RemoveListenerDrop() {
	if e.fnDrop == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"drop",
		*e.fnDrop,
	)
	e.fnDrop = nil

}

func (e *commonEvents) AddListenerDurationchange(genericEvent chan generic.Data) {
	if e.fnDurationchange != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventDurationchange, this, args)
		return nil
	})
	e.fnDurationchange = &fn

	e.selfElement.Call(
		"addEventListener",
		"durationchange",
		*e.fnDurationchange,
	)

}

func (e *commonEvents) RemoveListenerDurationchange() {
	if e.fnDurationchange == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"durationchange",
		*e.fnDurationchange,
	)
	e.fnDurationchange = nil

}

func (e *commonEvents) AddListenerEmptied(genericEvent chan generic.Data) {
	if e.fnEmptied != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventEmptied, this, args)
		return nil
	})
	e.fnEmptied = &fn

	e.selfElement.Call(
		"addEventListener",
		"emptied",
		*e.fnEmptied,
	)

}

func (e *commonEvents) RemoveListenerEmptied() {
	if e.fnEmptied == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"emptied",
		*e.fnEmptied,
	)
	e.fnEmptied = nil

}

func (e *commonEvents) AddListenerEnded(genericEvent chan generic.Data) {
	if e.fnEnded != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventEnded, this, args)
		return nil
	})
	e.fnEnded = &fn

	e.selfElement.Call(
		"addEventListener",
		"ended",
		*e.fnEnded,
	)

}

func (e *commonEvents) RemoveListenerEnded() {
	if e.fnEnded == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"ended",
		*e.fnEnded,
	)
	e.fnEnded = nil

}

func (e *commonEvents) AddListenerFormdata(genericEvent chan generic.Data) {
	if e.fnFormdata != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventFormdata, this, args)
		return nil
	})
	e.fnFormdata = &fn

	e.selfElement.Call(
		"addEventListener",
		"formdata",
		*e.fnFormdata,
	)

}

func (e *commonEvents) RemoveListenerFormdata() {
	if e.fnFormdata == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"formdata",
		*e.fnFormdata,
	)
	e.fnFormdata = nil

}

func (e *commonEvents) AddListenerInput(genericEvent chan generic.Data) {
	if e.fnInput != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventInput, this, args)
		return nil
	})
	e.fnInput = &fn

	e.selfElement.Call(
		"addEventListener",
		"input",
		*e.fnInput,
	)

}

func (e *commonEvents) RemoveListenerInput() {
	if e.fnInput == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"input",
		*e.fnInput,
	)
	e.fnInput = nil

}

func (e *commonEvents) AddListenerInvalid(genericEvent chan generic.Data) {
	if e.fnInvalid != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventInvalid, this, args)
		return nil
	})
	e.fnInvalid = &fn

	e.selfElement.Call(
		"addEventListener",
		"invalid",
		*e.fnInvalid,
	)

}

func (e *commonEvents) RemoveListenerInvalid() {
	if e.fnInvalid == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"invalid",
		*e.fnInvalid,
	)
	e.fnInvalid = nil

}

func (e *commonEvents) AddListenerKeydown(genericEvent chan generic.Data) {
	if e.fnKeydown != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventKeydown, this, args)
		return nil
	})
	e.fnKeydown = &fn

	e.selfElement.Call(
		"addEventListener",
		"keydown",
		*e.fnKeydown,
	)

}

func (e *commonEvents) RemoveListenerKeydown() {
	if e.fnKeydown == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"keydown",
		*e.fnKeydown,
	)
	e.fnKeydown = nil

}

func (e *commonEvents) AddListenerKeypress(genericEvent chan generic.Data) {
	if e.fnKeypress != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventKeypress, this, args)
		return nil
	})
	e.fnKeypress = &fn

	e.selfElement.Call(
		"addEventListener",
		"keypress",
		*e.fnKeypress,
	)

}

func (e *commonEvents) RemoveListenerKeypress() {
	if e.fnKeypress == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"keypress",
		*e.fnKeypress,
	)
	e.fnKeypress = nil

}

func (e *commonEvents) AddListenerKeyup(genericEvent chan generic.Data) {
	if e.fnKeyup != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventKeyup, this, args)
		return nil
	})
	e.fnKeyup = &fn

	e.selfElement.Call(
		"addEventListener",
		"keyup",
		*e.fnKeyup,
	)

}

func (e *commonEvents) RemoveListenerKeyup() {
	if e.fnKeyup == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"keyup",
		*e.fnKeyup,
	)
	e.fnKeyup = nil

}

func (e *commonEvents) AddListenerLoadeddata(genericEvent chan generic.Data) {
	if e.fnLoadeddata != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventLoadeddata, this, args)
		return nil
	})
	e.fnLoadeddata = &fn

	e.selfElement.Call(
		"addEventListener",
		"loadeddata",
		*e.fnLoadeddata,
	)

}

func (e *commonEvents) RemoveListenerLoadeddata() {
	if e.fnLoadeddata == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"loadeddata",
		*e.fnLoadeddata,
	)
	e.fnLoadeddata = nil

}

func (e *commonEvents) AddListenerLoadedmetadata(genericEvent chan generic.Data) {
	if e.fnLoadedmetadata != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventLoadedmetadata, this, args)
		return nil
	})
	e.fnLoadedmetadata = &fn

	e.selfElement.Call(
		"addEventListener",
		"loadedmetadata",
		*e.fnLoadedmetadata,
	)

}

func (e *commonEvents) RemoveListenerLoadedmetadata() {
	if e.fnLoadedmetadata == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"loadedmetadata",
		*e.fnLoadedmetadata,
	)
	e.fnLoadedmetadata = nil

}

func (e *commonEvents) AddListenerLoadstart(genericEvent chan generic.Data) {
	if e.fnLoadstart != nil {

	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventLoadstart, this, args)
		return nil
	})
	e.fnLoadstart = &fn

	e.selfElement.Call(
		"addEventListener",
		"loadstart",
		*e.fnLoadstart,
	)
	return
}

func (e *commonEvents) RemoveListenerLoadstart() {
	if e.fnLoadstart == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"loadstart",
		*e.fnLoadstart,
	)
	e.fnLoadstart = nil

}

func (e *commonEvents) AddListenerMousedown(genericEvent chan generic.Data) {
	if e.fnMousedown != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventMousedown, this, args)
		return nil
	})
	e.fnMousedown = &fn

	e.selfElement.Call(
		"addEventListener",
		"mousedown",
		*e.fnMousedown,
	)

}

func (e *commonEvents) RemoveListenerMousedown() {
	if e.fnMousedown == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"mousedown",
		*e.fnMousedown,
	)
	e.fnMousedown = nil

}

func (e *commonEvents) AddListenerMouseenter(genericEvent chan generic.Data) {
	if e.fnMouseenter != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventMouseenter, this, args)
		return nil
	})
	e.fnMouseenter = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseenter",
		*e.fnMouseenter,
	)

}

func (e *commonEvents) RemoveListenerMouseenter() {
	if e.fnMouseenter == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseenter",
		*e.fnMouseenter,
	)
	e.fnMouseenter = nil

}

func (e *commonEvents) AddListenerMouseleave(genericEvent chan generic.Data) {
	if e.fnMouseleave != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventMouseleave, this, args)
		return nil
	})
	e.fnMouseleave = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseleave",
		*e.fnMouseleave,
	)

}

func (e *commonEvents) RemoveListenerMouseleave() {
	if e.fnMouseleave == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseleave",
		*e.fnMouseleave,
	)
	e.fnMouseleave = nil

}

func (e *commonEvents) AddListenerMousemove(genericEvent chan generic.Data) {
	if e.fnMousemove != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventMousemove, this, args)
		return nil
	})
	e.fnMousemove = &fn

	e.selfElement.Call(
		"addEventListener",
		"mousemove",
		*e.fnMousemove,
	)

}

func (e *commonEvents) RemoveListenerMousemove() {
	if e.fnMousemove == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"mousemove",
		*e.fnMousemove,
	)
	e.fnMousemove = nil

}

func (e *commonEvents) AddListenerMouseout(genericEvent chan generic.Data) {
	if e.fnMouseout != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventMouseout, this, args)
		return nil
	})
	e.fnMouseout = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseout",
		*e.fnMouseout,
	)

}

func (e *commonEvents) RemoveListenerMouseout() {
	if e.fnMouseout == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseout",
		*e.fnMouseout,
	)
	e.fnMouseout = nil

}

func (e *commonEvents) AddListenerMouseover(genericEvent chan generic.Data) {
	if e.fnMouseover != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventMouseover, this, args)
		return nil
	})
	e.fnMouseover = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseover",
		*e.fnMouseover,
	)

}

func (e *commonEvents) RemoveListenerMouseover() {
	if e.fnMouseover == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseover",
		*e.fnMouseover,
	)
	e.fnMouseover = nil

}

func (e *commonEvents) AddListenerMouseup(genericEvent chan generic.Data) {
	if e.fnMouseup != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventMouseup, this, args)
		return nil
	})
	e.fnMouseup = &fn

	e.selfElement.Call(
		"addEventListener",
		"mouseup",
		*e.fnMouseup,
	)

}

func (e *commonEvents) RemoveListenerMouseup() {
	if e.fnMouseup == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"mouseup",
		*e.fnMouseup,
	)
	e.fnMouseup = nil

}

func (e *commonEvents) AddListenerPaste(genericEvent chan generic.Data) {
	if e.fnPaste != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventPaste, this, args)
		return nil
	})
	e.fnPaste = &fn

	e.selfElement.Call(
		"addEventListener",
		"paste",
		*e.fnPaste,
	)

}

func (e *commonEvents) RemoveListenerPaste() {
	if e.fnPaste == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"paste",
		*e.fnPaste,
	)
	e.fnPaste = nil

}

func (e *commonEvents) AddListenerPause(genericEvent chan generic.Data) {
	if e.fnPause != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventPause, this, args)
		return nil
	})
	e.fnPause = &fn

	e.selfElement.Call(
		"addEventListener",
		"pause",
		*e.fnPause,
	)

}

func (e *commonEvents) RemoveListenerPause() {
	if e.fnPause == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"pause",
		*e.fnPause,
	)
	e.fnPause = nil

}

func (e *commonEvents) AddListenerPlay(genericEvent chan generic.Data) {
	if e.fnPlay != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventPlay, this, args)
		return nil
	})
	e.fnPlay = &fn

	e.selfElement.Call(
		"addEventListener",
		"play",
		*e.fnPlay,
	)

}

func (e *commonEvents) RemoveListenerPlay() {
	if e.fnPlay == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"play",
		*e.fnPlay,
	)
	e.fnPlay = nil

}

func (e *commonEvents) AddListenerPlaying(genericEvent chan generic.Data) {
	if e.fnPlaying != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventPlaying, this, args)
		return nil
	})
	e.fnPlaying = &fn

	e.selfElement.Call(
		"addEventListener",
		"playing",
		*e.fnPlaying,
	)

}

func (e *commonEvents) RemoveListenerPlaying() {
	if e.fnPlaying == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"playing",
		*e.fnPlaying,
	)
	e.fnPlaying = nil

}

func (e *commonEvents) AddListenerProgress(genericEvent chan generic.Data) {
	if e.fnProgress != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventProgress, this, args)
		return nil
	})
	e.fnProgress = &fn

	e.selfElement.Call(
		"addEventListener",
		"progress",
		*e.fnProgress,
	)

}

func (e *commonEvents) RemoveListenerProgress() {
	if e.fnProgress == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"progress",
		*e.fnProgress,
	)
	e.fnProgress = nil

}

func (e *commonEvents) AddListenerRatechange(genericEvent chan generic.Data) {
	if e.fnRatechange != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventRatechange, this, args)
		return nil
	})
	e.fnRatechange = &fn

	e.selfElement.Call(
		"addEventListener",
		"ratechange",
		*e.fnRatechange,
	)

}

func (e *commonEvents) RemoveListenerRatechange() {
	if e.fnRatechange == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"ratechange",
		*e.fnRatechange,
	)
	e.fnRatechange = nil

}

func (e *commonEvents) AddListenerReset(genericEvent chan generic.Data) {
	if e.fnReset != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventReset, this, args)
		return nil
	})
	e.fnReset = &fn

	e.selfElement.Call(
		"addEventListener",
		"reset",
		*e.fnReset,
	)

}

func (e *commonEvents) RemoveListenerReset() {
	if e.fnReset == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"reset",
		*e.fnReset,
	)
	e.fnReset = nil

}

func (e *commonEvents) AddListenerScrollend(genericEvent chan generic.Data) {
	if e.fnScrollend != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventScrollend, this, args)
		return nil
	})
	e.fnScrollend = &fn

	e.selfElement.Call(
		"addEventListener",
		"scrollend",
		*e.fnScrollend,
	)

}

func (e *commonEvents) RemoveListenerScrollend() {
	if e.fnScrollend == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"scrollend",
		*e.fnScrollend,
	)
	e.fnScrollend = nil

}

func (e *commonEvents) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) {
	if e.fnSecuritypolicyviolation != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventSecuritypolicyviolation, this, args)
		return nil
	})
	e.fnSecuritypolicyviolation = &fn

	e.selfElement.Call(
		"addEventListener",
		"securitypolicyviolation",
		*e.fnSecuritypolicyviolation,
	)

}

func (e *commonEvents) RemoveListenerSecuritypolicyviolation() {
	if e.fnSecuritypolicyviolation == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"securitypolicyviolation",
		*e.fnSecuritypolicyviolation,
	)
	e.fnSecuritypolicyviolation = nil

}

func (e *commonEvents) AddListenerSeeked(genericEvent chan generic.Data) {
	if e.fnSeeked != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventSeeked, this, args)
		return nil
	})
	e.fnSeeked = &fn

	e.selfElement.Call(
		"addEventListener",
		"seeked",
		*e.fnSeeked,
	)

}

func (e *commonEvents) RemoveListenerSeeked() {
	if e.fnSeeked == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"seeked",
		*e.fnSeeked,
	)
	e.fnSeeked = nil

}

func (e *commonEvents) AddListenerSeeking(genericEvent chan generic.Data) {
	if e.fnSeeking != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventSeeking, this, args)
		return nil
	})
	e.fnSeeking = &fn

	e.selfElement.Call(
		"addEventListener",
		"seeking",
		*e.fnSeeking,
	)

}

func (e *commonEvents) RemoveListenerSeeking() {
	if e.fnSeeking == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"seeking",
		*e.fnSeeking,
	)
	e.fnSeeking = nil

}

func (e *commonEvents) AddListenerSelect(genericEvent chan generic.Data) {
	if e.fnSelect != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventSelect, this, args)
		return nil
	})
	e.fnSelect = &fn

	e.selfElement.Call(
		"addEventListener",
		"select",
		*e.fnSelect,
	)

}

func (e *commonEvents) RemoveListenerSelect() {
	if e.fnSelect == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"select",
		*e.fnSelect,
	)
	e.fnSelect = nil

}

func (e *commonEvents) AddListenerSlotchange(genericEvent chan generic.Data) {
	if e.fnSlotchange != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventSlotchange, this, args)
		return nil
	})
	e.fnSlotchange = &fn

	e.selfElement.Call(
		"addEventListener",
		"slotchange",
		*e.fnSlotchange,
	)

}

func (e *commonEvents) RemoveListenerSlotchange() {
	if e.fnSlotchange == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"slotchange",
		*e.fnSlotchange,
	)
	e.fnSlotchange = nil

}

func (e *commonEvents) AddListenerStalled(genericEvent chan generic.Data) {
	if e.fnStalled != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventStalled, this, args)
		return nil
	})
	e.fnStalled = &fn

	e.selfElement.Call(
		"addEventListener",
		"stalled",
		*e.fnStalled,
	)

}

func (e *commonEvents) RemoveListenerStalled() {
	if e.fnStalled == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"stalled",
		*e.fnStalled,
	)
	e.fnStalled = nil

}

func (e *commonEvents) AddListenerSubmit(genericEvent chan generic.Data) {
	if e.fnSubmit != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventSubmit, this, args)
		return nil
	})
	e.fnSubmit = &fn

	e.selfElement.Call(
		"addEventListener",
		"submit",
		*e.fnSubmit,
	)

}

func (e *commonEvents) RemoveListenerSubmit() {
	if e.fnSubmit == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"submit",
		*e.fnSubmit,
	)
	e.fnSubmit = nil

}

func (e *commonEvents) AddListenerSuspend(genericEvent chan generic.Data) {
	if e.fnSuspend != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventSuspend, this, args)
		return nil
	})
	e.fnSuspend = &fn

	e.selfElement.Call(
		"addEventListener",
		"suspend",
		*e.fnSuspend,
	)

}

func (e *commonEvents) RemoveListenerSuspend() {
	if e.fnSuspend == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"suspend",
		*e.fnSuspend,
	)
	e.fnSuspend = nil

}

func (e *commonEvents) AddListenerTimeupdate(genericEvent chan generic.Data) {
	if e.fnTimeupdate != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventTimeupdate, this, args)
		return nil
	})
	e.fnTimeupdate = &fn

	e.selfElement.Call(
		"addEventListener",
		"timeupdate",
		*e.fnTimeupdate,
	)

}

func (e *commonEvents) RemoveListenerTimeupdate() {
	if e.fnTimeupdate == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"timeupdate",
		*e.fnTimeupdate,
	)
	e.fnTimeupdate = nil

}

func (e *commonEvents) AddListenerToggle(genericEvent chan generic.Data) {
	if e.fnToggle != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventToggle, this, args)
		return nil
	})
	e.fnToggle = &fn

	e.selfElement.Call(
		"addEventListener",
		"toggle",
		*e.fnToggle,
	)

}

func (e *commonEvents) RemoveListenerToggle() {
	if e.fnToggle == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"toggle",
		*e.fnToggle,
	)
	e.fnToggle = nil

}

func (e *commonEvents) AddListenerVolumechange(genericEvent chan generic.Data) {
	if e.fnVolumechange != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventVolumechange, this, args)
		return nil
	})
	e.fnVolumechange = &fn

	e.selfElement.Call(
		"addEventListener",
		"volumechange",
		*e.fnVolumechange,
	)

}

func (e *commonEvents) RemoveListenerVolumechange() {
	if e.fnVolumechange == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"volumechange",
		*e.fnVolumechange,
	)
	e.fnVolumechange = nil

}

func (e *commonEvents) AddListenerWaiting(genericEvent chan generic.Data) {
	if e.fnWaiting != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventWaiting, this, args)
		return nil
	})
	e.fnWaiting = &fn

	e.selfElement.Call(
		"addEventListener",
		"waiting",
		*e.fnWaiting,
	)

}

func (e *commonEvents) RemoveListenerWaiting() {
	if e.fnWaiting == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"waiting",
		*e.fnWaiting,
	)
	e.fnWaiting = nil

}

func (e *commonEvents) AddListenerWebkitanimationend(genericEvent chan generic.Data) {
	if e.fnWebkitanimationend != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventWebkitanimationend, this, args)
		return nil
	})
	e.fnWebkitanimationend = &fn

	e.selfElement.Call(
		"addEventListener",
		"webkitanimationend",
		*e.fnWebkitanimationend,
	)

}

func (e *commonEvents) RemoveListenerWebkitanimationend() {
	if e.fnWebkitanimationend == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"webkitanimationend",
		*e.fnWebkitanimationend,
	)
	e.fnWebkitanimationend = nil

}

func (e *commonEvents) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) {
	if e.fnWebkitanimationiteration != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventWebkitanimationiteration, this, args)
		return nil
	})
	e.fnWebkitanimationiteration = &fn

	e.selfElement.Call(
		"addEventListener",
		"webkitanimationiteration",
		*e.fnWebkitanimationiteration,
	)

}

func (e *commonEvents) RemoveListenerWebkitanimationiteration() {
	if e.fnWebkitanimationiteration == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"webkitanimationiteration",
		*e.fnWebkitanimationiteration,
	)
	e.fnWebkitanimationiteration = nil

}

func (e *commonEvents) AddListenerWebkitanimationstart(genericEvent chan generic.Data) {
	if e.fnWebkitanimationstart != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventWebkitanimationstart, this, args)
		return nil
	})
	e.fnWebkitanimationstart = &fn

	e.selfElement.Call(
		"addEventListener",
		"webkitanimationstart",
		*e.fnWebkitanimationstart,
	)

}

func (e *commonEvents) RemoveListenerWebkitanimationstart() {
	if e.fnWebkitanimationstart == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"webkitanimationstart",
		*e.fnWebkitanimationstart,
	)
	e.fnWebkitanimationstart = nil

}

func (e *commonEvents) AddListenerWebkittransitionend(genericEvent chan generic.Data) {
	if e.fnWebkittransitionend != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventWebkittransitionend, this, args)
		return nil
	})
	e.fnWebkittransitionend = &fn

	e.selfElement.Call(
		"addEventListener",
		"webkittransitionend",
		*e.fnWebkittransitionend,
	)

}

func (e *commonEvents) RemoveListenerWebkittransitionend() {
	if e.fnWebkittransitionend == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"webkittransitionend",
		*e.fnWebkittransitionend,
	)
	e.fnWebkittransitionend = nil

}

func (e *commonEvents) AddListenerWheel(genericEvent chan generic.Data) {
	if e.fnWheel != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventWheel, this, args)
		return nil
	})
	e.fnWheel = &fn

	e.selfElement.Call(
		"addEventListener",
		"wheel",
		*e.fnWheel,
	)

}

func (e *commonEvents) RemoveListenerWheel() {
	if e.fnWheel == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"wheel",
		*e.fnWheel,
	)
	e.fnWheel = nil

}

func (e *commonEvents) AddListenerBlur(genericEvent chan generic.Data) {
	if e.fnBlur != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventBlur, this, args)
		return nil
	})
	e.fnBlur = &fn

	e.selfElement.Call(
		"addEventListener",
		"blur",
		*e.fnBlur,
	)

}

func (e *commonEvents) RemoveListenerBlur() {
	if e.fnBlur == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"blur",
		*e.fnBlur,
	)
	e.fnBlur = nil

}

func (e *commonEvents) AddListenerError(genericEvent chan generic.Data) {
	if e.fnError != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventError, this, args)
		return nil
	})
	e.fnError = &fn

	e.selfElement.Call(
		"addEventListener",
		"error",
		*e.fnError,
	)

}

func (e *commonEvents) RemoveListenerError() {
	if e.fnError == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"error",
		*e.fnError,
	)
	e.fnError = nil

}

func (e *commonEvents) AddListenerFocus(genericEvent chan generic.Data) {
	if e.fnFocus != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventFocus, this, args)
		return nil
	})
	e.fnFocus = &fn

	e.selfElement.Call(
		"addEventListener",
		"focus",
		*e.fnFocus,
	)

}

func (e *commonEvents) RemoveListenerFocus() {
	if e.fnFocus == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"focus",
		*e.fnFocus,
	)
	e.fnFocus = nil

}

func (e *commonEvents) AddListenerLoad(genericEvent chan generic.Data) {
	if e.fnLoad != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventLoad, this, args)
		return nil
	})
	e.fnLoad = &fn

	e.selfElement.Call(
		"addEventListener",
		"load",
		*e.fnLoad,
	)

}

func (e *commonEvents) RemoveListenerLoad() {
	if e.fnLoad == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"load",
		*e.fnLoad,
	)
	e.fnLoad = nil

}

func (e *commonEvents) AddListenerResize(genericEvent chan generic.Data) {
	if e.fnResize != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventResize, this, args)
		return nil
	})
	e.fnResize = &fn

	e.selfElement.Call(
		"addEventListener",
		"resize",
		*e.fnResize,
	)

}

func (e *commonEvents) RemoveListenerResize() {
	if e.fnResize == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"resize",
		*e.fnResize,
	)
	e.fnResize = nil

}

func (e *commonEvents) AddListenerScroll(genericEvent chan generic.Data) {
	if e.fnScroll != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventScroll, this, args)
		return nil
	})
	e.fnScroll = &fn

	e.selfElement.Call(
		"addEventListener",
		"scroll",
		*e.fnScroll,
	)

}

func (e *commonEvents) RemoveListenerScroll() {
	if e.fnScroll == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"scroll",
		*e.fnScroll,
	)
	e.fnScroll = nil

}

func (e *commonEvents) AddListenerAfterprint(genericEvent chan generic.Data) {
	if e.fnAfterprint != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventAfterprint, this, args)
		return nil
	})
	e.fnAfterprint = &fn

	e.selfElement.Call(
		"addEventListener",
		"afterprint",
		*e.fnAfterprint,
	)

}

func (e *commonEvents) RemoveListenerAfterprint() {
	if e.fnAfterprint == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"afterprint",
		*e.fnAfterprint,
	)
	e.fnAfterprint = nil

}

func (e *commonEvents) AddListenerBeforeprint(genericEvent chan generic.Data) {
	if e.fnBeforeprint != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventBeforeprint, this, args)
		return nil
	})
	e.fnBeforeprint = &fn

	e.selfElement.Call(
		"addEventListener",
		"beforeprint",
		*e.fnBeforeprint,
	)

}

func (e *commonEvents) RemoveListenerBeforeprint() {
	if e.fnBeforeprint == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"beforeprint",
		*e.fnBeforeprint,
	)
	e.fnBeforeprint = nil

}

func (e *commonEvents) AddListenerBeforeunload(genericEvent chan generic.Data) {
	if e.fnBeforeunload != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventBeforeunload, this, args)
		return nil
	})
	e.fnBeforeunload = &fn

	e.selfElement.Call(
		"addEventListener",
		"beforeunload",
		*e.fnBeforeunload,
	)

}

func (e *commonEvents) RemoveListenerBeforeunload() {
	if e.fnBeforeunload == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"beforeunload",
		*e.fnBeforeunload,
	)
	e.fnBeforeunload = nil

}

func (e *commonEvents) AddListenerHashchange(genericEvent chan generic.Data) {
	if e.fnHashchange != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventHashchange, this, args)
		return nil
	})
	e.fnHashchange = &fn

	e.selfElement.Call(
		"addEventListener",
		"hashchange",
		*e.fnHashchange,
	)

}

func (e *commonEvents) RemoveListenerHashchange() {
	if e.fnHashchange == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"hashchange",
		*e.fnHashchange,
	)
	e.fnHashchange = nil

}

func (e *commonEvents) AddListenerLanguagechange(genericEvent chan generic.Data) {
	if e.fnLanguagechange != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventLanguagechange, this, args)
		return nil
	})
	e.fnLanguagechange = &fn

	e.selfElement.Call(
		"addEventListener",
		"languagechange",
		*e.fnLanguagechange,
	)

}

func (e *commonEvents) RemoveListenerLanguagechange() {
	if e.fnLanguagechange == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"languagechange",
		*e.fnLanguagechange,
	)
	e.fnLanguagechange = nil

}

func (e *commonEvents) AddListenerMessage(genericEvent chan generic.Data) {
	if e.fnMessage != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventMessage, this, args)
		return nil
	})
	e.fnMessage = &fn

	e.selfElement.Call(
		"addEventListener",
		"message",
		*e.fnMessage,
	)

}

func (e *commonEvents) RemoveListenerMessage() {
	if e.fnMessage == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"message",
		*e.fnMessage,
	)
	e.fnMessage = nil

}

func (e *commonEvents) AddListenerMessageerror(genericEvent chan generic.Data) {
	if e.fnMessageerror != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventMessageerror, this, args)
		return nil
	})
	e.fnMessageerror = &fn

	e.selfElement.Call(
		"addEventListener",
		"messageerror",
		*e.fnMessageerror,
	)

}

func (e *commonEvents) RemoveListenerMessageerror() {
	if e.fnMessageerror == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"messageerror",
		*e.fnMessageerror,
	)
	e.fnMessageerror = nil

}

func (e *commonEvents) AddListenerOffline(genericEvent chan generic.Data) {
	if e.fnOffline != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventOffline, this, args)
		return nil
	})
	e.fnOffline = &fn

	e.selfElement.Call(
		"addEventListener",
		"offline",
		*e.fnOffline,
	)

}

func (e *commonEvents) RemoveListenerOffline() {
	if e.fnOffline == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"offline",
		*e.fnOffline,
	)
	e.fnOffline = nil

}

func (e *commonEvents) AddListenerOnline(genericEvent chan generic.Data) {
	if e.fnOnline != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventOnline, this, args)
		return nil
	})
	e.fnOnline = &fn

	e.selfElement.Call(
		"addEventListener",
		"online",
		*e.fnOnline,
	)

}

func (e *commonEvents) RemoveListenerOnline() {
	if e.fnOnline == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"online",
		*e.fnOnline,
	)
	e.fnOnline = nil

}

func (e *commonEvents) AddListenerPageswap(genericEvent chan generic.Data) {
	if e.fnPageswap != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventPageswap, this, args)
		return nil
	})
	e.fnPageswap = &fn

	e.selfElement.Call(
		"addEventListener",
		"pageswap",
		*e.fnPageswap,
	)

}

func (e *commonEvents) RemoveListenerPageswap() {
	if e.fnPageswap == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"pageswap",
		*e.fnPageswap,
	)
	e.fnPageswap = nil

}

func (e *commonEvents) AddListenerPagehide(genericEvent chan generic.Data) {
	if e.fnPagehide != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventPagehide, this, args)
		return nil
	})
	e.fnPagehide = &fn

	e.selfElement.Call(
		"addEventListener",
		"pagehide",
		*e.fnPagehide,
	)

}

func (e *commonEvents) RemoveListenerPagehide() {
	if e.fnPagehide == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"pagehide",
		*e.fnPagehide,
	)
	e.fnPagehide = nil

}

func (e *commonEvents) AddListenerPagereveal(genericEvent chan generic.Data) {
	if e.fnPagereveal != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventPagereveal, this, args)
		return nil
	})
	e.fnPagereveal = &fn

	e.selfElement.Call(
		"addEventListener",
		"pagereveal",
		*e.fnPagereveal,
	)

}

func (e *commonEvents) RemoveListenerPagereveal() {
	if e.fnPagereveal == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"pagereveal",
		*e.fnPagereveal,
	)
	e.fnPagereveal = nil

}

func (e *commonEvents) AddListenerPageshow(genericEvent chan generic.Data) {
	if e.fnPageshow != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventPageshow, this, args)
		return nil
	})
	e.fnPageshow = &fn

	e.selfElement.Call(
		"addEventListener",
		"pageshow",
		*e.fnPageshow,
	)

}

func (e *commonEvents) RemoveListenerPageshow() {
	if e.fnPageshow == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"pageshow",
		*e.fnPageshow,
	)
	e.fnPageshow = nil

}

func (e *commonEvents) AddListenerPopstate(genericEvent chan generic.Data) {
	if e.fnPopstate != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventPopstate, this, args)
		return nil
	})
	e.fnPopstate = &fn

	e.selfElement.Call(
		"addEventListener",
		"popstate",
		*e.fnPopstate,
	)

}

func (e *commonEvents) RemoveListenerPopstate() {
	if e.fnPopstate == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"popstate",
		*e.fnPopstate,
	)
	e.fnPopstate = nil

}

func (e *commonEvents) AddListenerRejectionhandled(genericEvent chan generic.Data) {
	if e.fnRejectionhandled != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventRejectionhandled, this, args)
		return nil
	})
	e.fnRejectionhandled = &fn

	e.selfElement.Call(
		"addEventListener",
		"rejectionhandled",
		*e.fnRejectionhandled,
	)

}

func (e *commonEvents) RemoveListenerRejectionhandled() {
	if e.fnRejectionhandled == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"rejectionhandled",
		*e.fnRejectionhandled,
	)
	e.fnRejectionhandled = nil

}

func (e *commonEvents) AddListenerStorage(genericEvent chan generic.Data) {
	if e.fnStorage != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventStorage, this, args)
		return nil
	})
	e.fnStorage = &fn

	e.selfElement.Call(
		"addEventListener",
		"storage",
		*e.fnStorage,
	)

}

func (e *commonEvents) RemoveListenerStorage() {
	if e.fnStorage == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"storage",
		*e.fnStorage,
	)
	e.fnStorage = nil

}

func (e *commonEvents) AddListenerUnhandledrejection(genericEvent chan generic.Data) {
	if e.fnUnhandledrejection != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventUnhandledrejection, this, args)
		return nil
	})
	e.fnUnhandledrejection = &fn

	e.selfElement.Call(
		"addEventListener",
		"unhandledrejection",
		*e.fnUnhandledrejection,
	)

}

func (e *commonEvents) RemoveListenerUnhandledrejection() {
	if e.fnUnhandledrejection == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"unhandledrejection",
		*e.fnUnhandledrejection,
	)
	e.fnUnhandledrejection = nil

}

func (e *commonEvents) AddListenerUnload(genericEvent chan generic.Data) {
	if e.fnUnload != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventUnload, this, args)
		return nil
	})
	e.fnUnload = &fn

	e.selfElement.Call(
		"addEventListener",
		"unload",
		*e.fnUnload,
	)

}

func (e *commonEvents) RemoveListenerUnload() {
	if e.fnUnload == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"unload",
		*e.fnUnload,
	)
	e.fnUnload = nil

}

func (e *commonEvents) AddListenerReadystatechange(genericEvent chan generic.Data) {
	if e.fnReadystatechange != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventReadystatechange, this, args)
		return nil
	})
	e.fnReadystatechange = &fn

	e.selfElement.Call(
		"addEventListener",
		"readystatechange",
		*e.fnReadystatechange,
	)

}

func (e *commonEvents) RemoveListenerReadystatechange() {
	if e.fnReadystatechange == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"readystatechange",
		*e.fnReadystatechange,
	)
	e.fnReadystatechange = nil

}

func (e *commonEvents) AddListenerVisibilitychange(genericEvent chan generic.Data) {
	if e.fnVisibilitychange != nil {
		return
	}

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}
		genericEvent <- generic.EventManager(generic.KEventVisibilitychange, this, args)
		return nil
	})
	e.fnVisibilitychange = &fn

	e.selfElement.Call(
		"addEventListener",
		"visibilitychange",
		*e.fnVisibilitychange,
	)

}

func (e *commonEvents) RemoveListenerVisibilitychange() {
	if e.fnVisibilitychange == nil {
		return
	}

	e.selfElement.Call(
		"removeEventListener",
		"visibilitychange",
		*e.fnVisibilitychange,
	)
	e.fnVisibilitychange = nil
}
