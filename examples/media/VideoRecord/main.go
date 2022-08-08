//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/mouse"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/browser/media"
	"log"
	"syscall/js"
)

func main() {

	done := make(chan struct{}, 0)

	var obj = js.Global().Get("navigator").Get("mediaDevices").Call("getSupportedConstraints")
	js.Global().Get("Object").Call("keys", obj).Call("forEach", js.FuncOf(func(this js.Value, args []js.Value) any {
		log.Printf("key: [%v]: %v", args[0].String(), obj.Get(args[0].String()).String())
		return nil
	}))

	clickStart := make(chan mouse.Data, 0)
	clickStop := make(chan mouse.Data, 0)

	videoPreview := &html.TagVideo{}
	videoRecording := &html.TagVideo{}
	startButton := &html.TagDiv{}
	stopButton := &html.TagDiv{}
	downloadButton := &html.TagA{}

	stage := factoryBrowser.NewStage()

	left := factoryBrowser.NewTagDiv().Class("left").Append(
		factoryBrowser.NewTagDiv().Reference(&startButton).AddListenerClick(&clickStart).Id("startButton").Class("button").Text("Start Recording"),
		factoryBrowser.NewTagH2().Text("Preview"),
		factoryBrowser.NewTagVideo().Reference(&videoPreview).Id("preview").Width(160).Height(120).AutoPlay(true).Muted(true),
	)
	right := factoryBrowser.NewTagDiv().Class("right").Append(
		factoryBrowser.NewTagDiv().Reference(&stopButton).AddListenerClick(&clickStop).Id("stopButton").Class("button").Text("Stop Recording"),
		factoryBrowser.NewTagH2().Text("Recording"),
		factoryBrowser.NewTagVideo().Reference(&videoRecording).Id("recording").Width(160).Height(120).Controls(true),
		factoryBrowser.NewTagA().Reference(&downloadButton).Id("downloadButton").Class("button").Text("Download"),
	)

	stage.Append(left, right)

	go func() {

		for {
			select {
			case <-clickStart:
				config := media.NewFactory().
					VideoWidth(800).
					VideoHeight(600).
					VideoIso(200).
					AudioAutoGainControl(true).
					AudioChannelCount(2).
					AudioVolumeOptions(-1, -1, 100).
					VideoPointsOfInterest(
						[]any{
							map[string]any{
								"x": 20,
								"y": 20,
							},
							map[string]any{
								"x": 40,
								"y": 40,
							},
						},
					)
				videoPreview.RecordingUserMedia(config)
			case <-clickStop:
				recording := videoPreview.RecordingUserMediaStop()
				if recording == false {
					continue
				}

				src, err := videoPreview.GetRecordingSrcData()
				if err != nil {
					log.Printf("err: %v", err.Error())
					continue
				}

				videoRecording.Src(src)

				downloadButton.HRef(src)
				downloadButton.Download("RecordedVideo.webm")
			}
		}
	}()

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
