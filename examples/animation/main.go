//go:build js
// +build js

//
package main

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/css"
	document2 "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
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

	var err error
	document.GetElementById(document, "palco")
	for a := 0; a != 1; a += 1 {

		id := "div_" + strconv.FormatInt(int64(a), 10)
		var cssClass = css.Class{}
		cssClass.SetList("current", "animate")
		err = document.CreateElement(document, "palco", "div", document2.Property{Property: "id", Value: id}, cssClass)
		if err != nil {
			log.Printf("document.CreateElement().error: %v", err.Error())
		}
		var e = document.GetElementById(document, id)

		factoryTween.NewEaseInBack(
			//time.Duration(mathUtil.Int(2000, 5000))*time.Millisecond,
			2*time.Second,
			//mathUtil.Float64FomInt(0, global.Global.Document.GetDocumentWidth()-29),
			//mathUtil.Float64FomInt(0, global.Global.Document.GetDocumentWidth()-29),
			0,
			float64(global.Global.Document.GetDocumentWidth()-29),
			func(x, p float64, ars ...interface{}) {
				px := strconv.FormatFloat(x, 'E', 10, 32) + "px"
				document.SetElementStyle(e, "left", px)
			},
			-1,
		)

		factoryTween.NewEaseInBack(
			//time.Duration(mathUtil.Int(2000, 5000))*time.Millisecond,
			2*time.Second,
			//mathUtil.Float64FomInt(0, global.Global.Document.GetDocumentHeight()-50),
			0,
			//mathUtil.Float64FomInt(0, global.Global.Document.GetDocumentHeight()-50),
			float64(global.Global.Document.GetDocumentHeight()-50),
			func(y, p float64, ars ...interface{}) {
				py := strconv.FormatFloat(y, 'E', 10, 32) + "px"
				document.SetElementStyle(e, "top", py)
			},
			-1,
		)

	}

	<-done
}
