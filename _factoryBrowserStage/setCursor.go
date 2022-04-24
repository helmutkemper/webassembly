package factoryBrowserStage

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
)

//var mouseChannelSendToThread chan mouseChnn.Cursor

// fixme:

func SetCursor(name mouse.CursorType) {

	for _, cursor := range PreLoadMouseList {
		if cursor.id == name.String() {
			imageCursor.Img = cursor.Img.Get()
		}
	}

}
