package canvas

import (
	"syscall/js"
)

// todo: SelfContextType deve ser um enum

// en: The Canvas API provides a means for drawing graphics via JavaScript and the HTML <canvas> element. Among other
// things, it can be used for animation, game graphics, data visualization, photo manipulation, and real-time video
// processing.
//
// The Canvas API largely focuses on 2D graphics. The WebGL API, which also uses the <canvas> element, draws
// hardware-accelerated 2D and 3D graphics.
type Canvas struct {
	SelfContext     js.Value
	SelfContextType int
	SelfElement     js.Value
}
