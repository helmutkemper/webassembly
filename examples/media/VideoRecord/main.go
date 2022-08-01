//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/event"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/mouse"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"log"
	"syscall/js"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	options := js.Global().Get("Object")
	options.Set("type", "text/html")

	aFileParts := js.Global().Get("Array")
	aFileParts.SetIndex(0, "<a id=\"a\"><b id=\"b\">hey!</b></a>")

	js.Global().Get("window").Get("Blob").New([]interface{}{"<a id=\"a\"><b id=\"b\">hey!</b></a>"}, options)

	<-done

	click := make(chan mouse.Data)
	playing := make(chan event.Data)
	videoPreview := &html.TagVideo{}
	videoRecording := &html.TagVideo{}
	startButton := &html.TagDiv{}
	downloadButton := &html.TagA{}

	stage := factoryBrowser.NewStage()

	left := factoryBrowser.NewTagDiv().Class("left").Append(
		factoryBrowser.NewTagDiv().Reference(&startButton).AddListenerClick(&click).Id("startButton").Class("button").Text("Start Recording"),
		factoryBrowser.NewTagH2().Text("Preview"),
		factoryBrowser.NewTagVideo().Reference(&videoPreview).AddListenerPlaying(&playing).Id("preview").Width(160).Height(120).AutoPlay(true).Muted(true),
	)
	right := factoryBrowser.NewTagDiv().Class("right").Append(
		factoryBrowser.NewTagDiv().Id("stopButton").Class("button").Text("Stop Recording"),
		factoryBrowser.NewTagH2().Text("Recording"),
		factoryBrowser.NewTagVideo().Reference(&videoRecording).Id("recording").Width(160).Height(120).Controls(true),
		factoryBrowser.NewTagA().Reference(&downloadButton).Id("downloadButton").Class("button").Text("Download"),
	)

	stage.Append(left, right)
	log.Printf("%v", js.Global().Get("Promise"))

	var tOut = &time.Ticker{}
	go func() {

		captureStream := js.Value{}
		stream := js.Value{}
		//forEach := js.FuncOf(func(this js.Value, args []js.Value) any {
		//	log.Printf("%v", this)
		//	for _, v := range args {
		//		v.Call("stop")
		//	}
		//	return nil
		//})
		//recordedChunks := js.FuncOf(func(this js.Value, args []js.Value) any {
		//	recordedChunks := args[0]
		//	options := js.Global().Get("Object")
		//	options.Set("type", "video/webm")
		//	recordedBlob := js.Global().Get("Blob").New(recordedChunks, options)
		//	videoRecording.Src(js.Global().Call("createObjectURL", recordedBlob))
		//	return nil
		//})

		for {
			select {
			case <-tOut.C:
				log.Printf("%v", stream.Call("getTracks"))
				//stream.Get("getTracks").Call("forEach", forEach, js.Global().Get("Object"))
			case <-click:
				// timeout
				tOut = time.NewTicker(5 * time.Second)

				success := js.FuncOf(func(this js.Value, args []js.Value) any {
					// capture user media promise
					stream = args[0]
					videoPreview.Get().Set("srcObject", args[0])
					downloadButton.Get().Set("href", args[0])
					//preview.srcObject = stream;
					//downloadButton.href = stream;
					//preview.captureStream = preview.captureStream || preview.mozCaptureStream;

					if !(videoPreview.Get().Get("captureStream").IsNull() || videoPreview.Get().Get("captureStream").IsUndefined()) {
						captureStream = videoPreview.Get().Get("captureStream")
					} else {
						captureStream = videoPreview.Get().Get("mozCaptureStream")
					}
					videoPreview.Get().Set("captureStream", captureStream)

					//pw := js.Global().Call("getElementById", "preview")
					//pw.Get("srcObject")
					//js.Global().Get("MediaRecorder").New()
					return nil
				})

				// capture user media
				options := js.Global().Get("Object")
				options.Set("video", true)
				options.Set("audi", true)
				js.Global().Get("navigator").Get("mediaDevices").Call("getUserMedia", options).Call("then", success)

			case <-playing:
				// onPlaying promise
				log.Printf("videoPreview.Get().Call(\"captureStream\"): %v", videoPreview.Get().Call("captureStream"))
				//videoPreview.Get().Call("captureStream", recordedChunks)
				//startRecording()

				options := js.Global().Get("Object")
				options.Set("type", "video/webm")
				recordedBlob := js.Global().Call("Blob").New(videoPreview.Get().Call("captureStream"), options)
				videoRecording.Src(js.Global().Call("createObjectURL", recordedBlob))
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
