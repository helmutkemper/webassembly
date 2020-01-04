package canvas

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/mouse"

func (el *Canvas) SetMouseCursor(cursor mouse.CursorType) {
	el.SelfElement.Call("setAttribute", "style", cursor.String())
}
