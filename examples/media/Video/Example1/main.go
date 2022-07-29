// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
//
//  <video controls width="250">
//    <source src="https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm" type="video/webm">
//    <source src="https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.mp4" type="video/mp4">
//    Sorry, your browser doesn't support embedded videos.
//  </video>

//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"log"
)

func main() {

	stage := factoryBrowser.NewStage()

	videoEvent := make(chan event.Data)

	tagVideo := &html.TagVideo{}

	s1 := factoryBrowser.NewTagVideo().Reference(&tagVideo).AddListenerEnded(&videoEvent).Controls(true).Width(250).Append(
		factoryBrowser.NewTagSource().Src("https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm").Type("video/webm"),
		factoryBrowser.NewTagSource().Src("https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.mp4").Type("video/mp4"),
	)

	stage.Append(s1)

	go func() {
		for {
			select {
			case converted := <-videoEvent:
				log.Printf("%+v", converted.EventName)
				tagVideo.RemoveListenerEnded()
			}
		}
	}()

	done := make(chan struct{}, 0)
	<-done
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
