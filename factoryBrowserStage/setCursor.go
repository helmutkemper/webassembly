package factoryBrowserStage

import (
	mouseChnn "github.com/helmutkemper/iotmaker.santa_isabel_theater.channels-go/mouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
)

var mouseChannelSendToThread chan mouseChnn.Channel

func SetCursor(name mouse.CursorType) {
	var mouseConf = mouseChnn.Channel{
		Name: name.String(),
	}
	mouseChannelSendToThread <- mouseConf
}

func init() {
	mouseChannelSendToThread = make(chan mouseChnn.Channel)

	go func() {
		for {
			select {
			case m := <-mouseChannelSendToThread:

				for _, cursor := range PreLoadMouseList {
					if cursor.id == m.Name {
						imageCursor.Img = cursor.Img.Get()
					}
				}

			}
		}
	}()
}
