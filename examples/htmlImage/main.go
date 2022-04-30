//go:build js
// +build js

package main

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.webassembly/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.webassembly/html"
	"time"
)

var imgPlayer html.Image

func main() {

	done := make(chan struct{})
	stage := global.Global.Stage

	imgPlayer = factoryBrowserImage.NewImage(
		480,
		60,
		map[string]interface{}{
			"id":  "player_big",
			"src": "./player_big.png",
		},
		true,
		true,
	)

	p := factoryImage.NewMultipleSpritesImage(
		"id_multiple_sprite_player_bug",
		global.Global.Stage,
		global.Global.Canvas,
		global.Global.ScratchPad,
		nil,
		imgPlayer.Get(),
		480,
		60,
		0,
		9,
		1000*time.Millisecond,
		0,
		0,
		48,
		60,
		global.Global.Density,
		global.Global.DensityManager,
	)
	stage.AddToDraw(p)

	<-done
}
