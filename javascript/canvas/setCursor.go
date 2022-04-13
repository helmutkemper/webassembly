package canvas

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"

func (el *Canvas) SetMouseCursor(cursor browserMouse.CursorType) {
	el.SelfElement.Call("setAttribute", "style", cursor.String())
}
