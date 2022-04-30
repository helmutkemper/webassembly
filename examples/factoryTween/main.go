//go:build js
// +build js

// GOARCH=wasm;GOOS=js
// -o main.wasm
package main

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mathUtil"
	"github.com/helmutkemper/iotmaker.webassembly/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.webassembly/html"
	"strconv"
	"time"
)

var imgSpace html.Image

// enviroment var: GOARCH=wasm;GOOS=js
// go argument:    -o main.wasm
func main() {

	done := make(chan struct{}, 0)
	stage := global.Global.Stage

	// a imagem deve ser pre-carregada no navegador
	imgSpace = factoryBrowserImage.NewImage(
		29,
		50,
		map[string]interface{}{
			"width": 29,
			"heght": 50,
			"id":    "spacecraft",
			"src":   "./small.png",
		},
		true,
		false,
	)

	for a := 0; a != 100; a += 1 {
		i := factoryImage.NewImage(
			"id_"+strconv.FormatInt(int64(a), 10),
			global.Global.Stage,
			global.Global.Canvas,
			global.Global.ScratchPad,
			nil,
			imgSpace.Get(),
			-100,
			-100,
			29,
			50,
			global.Global.Density,
			global.Global.DensityManager,
		)
		i.SetDraggableToDesktop()
		stage.AddToDraw(i)

		factoryTween.NewSelectRandom(
			time.Duration(mathUtil.Int(500, 3000))*time.Millisecond,
			mathUtil.Float64FomInt(0, 1000),
			mathUtil.Float64FomInt(0, 1000),
			false,
			nil,
			nil,
			nil,
			nil,
			nil,
			func(x, p float64, ars ...interface{}) {
				i.MoveX(int(x))
			},
			-1,
		)

		factoryTween.NewSelectRandom(
			time.Duration(mathUtil.Int(500, 3000))*time.Millisecond,
			mathUtil.Float64FomInt(0, 800),
			mathUtil.Float64FomInt(0, 900),
			false,
			nil,
			nil,
			nil,
			nil,
			nil,
			func(y, p float64, ars ...interface{}) {
				i.MoveY(int(y))
			},
			-1,
		)

	}

	<-done
}
