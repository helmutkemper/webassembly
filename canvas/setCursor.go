package canvas

import "github.com/helmutkemper/iotmaker.platform.webbrowser/mouse"

func (el *Canvas) SetMouseCursor(cursor mouse.CursorType) {
	el.SelfElement.Call("setAttribute", "style", cursor.String())
}
