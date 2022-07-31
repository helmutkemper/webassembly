//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

func main() {

	done := make(chan struct{}, 0)

	videoPreview := &html.TagVideo{}

	stage := factoryBrowser.NewStage()

	left := factoryBrowser.NewTagDiv().Class("left").Append(
		factoryBrowser.NewTagDiv().Id("startButton").Class("button").Text("Start Recording"),
		factoryBrowser.NewTagH2().Text("Preview"),
		factoryBrowser.NewTagVideo().Reference(&videoPreview).Id("preview").Width(160).Height(120).AutoPlay(true).Muted(true),
	)
	right := factoryBrowser.NewTagDiv().Class("right").Append(
		factoryBrowser.NewTagDiv().Id("stopButton").Class("button").Text("Stop Recording"),
		factoryBrowser.NewTagH2().Text("Recording"),
		factoryBrowser.NewTagVideo().Id("recording").Width(160).Height(120).Controls(true),
		factoryBrowser.NewTagA().Id("downloadButton").Class("button").Text("Download"),
	)

	stage.Append(left, right)

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
