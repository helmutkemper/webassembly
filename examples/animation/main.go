//go:build js
// +build js

//
package main

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	doc "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mathUtil"
	"log"
	"strconv"
	"time"
)

func main() {

	done := make(chan struct{}, 0)
	document := global.Global.Document

	// Carrega a imagem
	factoryBrowserImage.NewImage(
		29,
		50,
		map[string]interface{}{
			"id":  "spacecraft",
			"src": "./small.png",
		},
		true,
		false,
	)

	var d html.Div
	d.NewDiv("example").
		Css("vivo").
		CssAddList("red", "user", "red").
		CssAddList("yellow", "user", "yellow").
		CssAddList("normal", "user").
		CssToggleTime(time.Second, "red", "yellow").
		CssToggleLoop(10).
		CssOnLoopEnd("normal").
		CssToggleStart().
		AppendById("palco")

	var err error
	//document.GetElementById(document, "palco")
	for a := 0; a != 1; a += 1 {

		id := "div_" + strconv.FormatInt(int64(a), 10)
		_, err = document.CreateElementAndAppend(
			"palco",
			"div",
			[]string{"animate"},
			doc.P{P: "id", V: id},
		)
		if err != nil {
			log.Printf("document.CreateElement().error: %v", err.Error())
		}
		var e = document.GetElementById(document, id)
		var border = 200
		factoryTween.NewSelectRandom(
			time.Duration(mathUtil.Int(1000, 3000))*time.Millisecond,
			mathUtil.Float64FomInt(border, global.Global.Document.GetDocumentWidth()-29-border),
			mathUtil.Float64FomInt(border, global.Global.Document.GetDocumentWidth()-29-border),
			func(x, p float64, ars ...interface{}) {
				px := strconv.FormatFloat(x, 'E', 10, 32) + "px"
				document.SetElementStyle(e, "left", px)
			},
			-1,
		)

		factoryTween.NewSelectRandom(
			time.Duration(mathUtil.Int(1000, 3000))*time.Millisecond,
			mathUtil.Float64FomInt(border, global.Global.Document.GetDocumentHeight()-50-border),
			mathUtil.Float64FomInt(border, global.Global.Document.GetDocumentHeight()-50-border),
			func(y, p float64, ars ...interface{}) {
				py := strconv.FormatFloat(y, 'E', 10, 32) + "px"
				document.SetElementStyle(e, "top", py)
			},
			-1,
		)
	}

	<-done
}
