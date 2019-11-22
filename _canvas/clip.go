package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Clips a region of any shape and size from the original canvas
//     The clip() method clips a region of any shape and size from the original canvas.
//     Tip: Once a region is clipped, all future drawing will be limited to the clipped region (no access to other
//     regions on the canvas). You can however save the current canvas region using the save() method before using the
//     clip() method, and restore it (with the restore() method) any time in the future.
func (el *Canvas) Clip(x, y iotmaker_types.Coordinate) {
	el.SelfContext.Call("clip", x, y)
}
