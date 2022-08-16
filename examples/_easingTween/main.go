//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
	"time"
)

//  .animate {
//    width: 29px;
//    height: 50px;
//    position: absolute;
//    background-image: url("./small.png");
//  }

func main() {

	done := make(chan struct{}, 0)

	var stage = stage.Stage{}
	stage.Init()

	factoryBrowser.NewTagDiv("div_o").
		Class("animate").
		NewEasingTweenLinear("x", 3*time.Second, 50, 300, onUpdateX, -1).
		NewEasingTweenLinear("y", 3*time.Second, 50, 300, onUpdateY, -1).
		AppendToStage()

	<-done
}

func onUpdateX(x, _ float64, args interface{}) {
	this := args.([]interface{})[0].(*html.TagDiv)
	this.SetX(int(x))
}

func onUpdateY(y, p float64, args interface{}) {
	this := args.([]interface{})[0].(*html.TagDiv)
	this.SetY(int(y))
}
