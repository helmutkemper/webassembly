//go:build js

package main

import (
	"fmt"
	keyboard "github.com/helmutkemper/iotmaker.webassembly/browser/event/keyBoard"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"image/color"
	"log"
)

func main() {

	text := "please, use arrow keys"

	imageBoxColor := color.RGBA{A: 0}
	collisionBoxColor := color.RGBA{R: 255, A: 255}

	// English: channel triggered by a keyboard event.
	// Português: channel disparado por um evento do teclado.
	eventOnKeyData := make(chan keyboard.Data)

	canvasA := factoryBrowser.NewTagCanvas(800, 600)
	canvasA.SetXY(0, 0)

	canvasB := factoryBrowser.NewTagCanvas(800, 600)
	canvasB.SetXY(0, 0)

	divA := factoryBrowser.NewTagDiv().
		SetXY(10, 10).
		Html(text)

	var stage = factoryBrowser.NewStage()
	stage.Append(canvasA, canvasB, divA)

	img1 := factoryBrowser.NewTagImg().
		Src("./platformPack_tilesheet.png", true)

	img2 := factoryBrowser.NewTagImg().
		Src("./player_big.png", true)

	w := 48
	h := 60

	i := 0
	canvasA.DrawImageComplete(img1, i*w, 0*h, w, h, 0, 0, w, h)
	groundData := canvasA.GetImageData(0, 0, w, h, false, false)
	canvasA.ClearRect(0*w, 0*h, w, h)

	groundBox := html.Box{}
	groundBox.Debug(&imageBoxColor, &collisionBoxColor)
	groundBox.Populate(groundData, w, h)
	groundBox.X(200)
	groundBox.Y(200)

	canvasA.PutImageData(groundBox.GetData(), groundBox.GetX(), groundBox.GetY())

	go func() {
		positionX := 200
		positionY := 140

		for {
			select {
			case data := <-eventOnKeyData:
				switch data.EventName {
				case "keydown":
					switch data.Code {
					case "ArrowUp":
						positionY -= 1
					case "ArrowDown":
						positionY += 1
					case "ArrowLeft":
						positionX -= 1
					case "ArrowRight":
						positionX += 1
					case "Space":
						i += 1
						if i > 9 {
							i = 0
						}
						log.Printf("i: %v", i)
					}
				}

				canvasB.DrawImageComplete(img2, i*w, 0*h, w, h, 0, 0, w, h)
				playerData := canvasB.GetImageData(0, 0, w, h, false, false)
				canvasB.ClearRect(0*w, 0*h, w, h)

				playerBox := html.Box{}
				playerBox.Debug(&imageBoxColor, &collisionBoxColor)
				playerBox.Populate(playerData, w, h)
				playerBox.X(positionX)
				playerBox.Y(positionY)

				canvasB.ClearRect(0, 0, 800, 600)
				canvasB.PutImageData(playerBox.GetData(), playerBox.GetX(), playerBox.GetY())

				text = "please, use arrow keys<br><br>Player collision:<br>"

				upLeft, upRight, downLeft, downRight := playerBox.Quadrant(groundBox)
				text += fmt.Sprintf("up: <b>%v</b><br>", upLeft || upRight)
				text += fmt.Sprintf("left: <b>%v</b><br>", upLeft || downLeft)
				text += fmt.Sprintf("right: <b>%v</b><br>", upRight || downRight)
				text += fmt.Sprintf("down: <b>%v</b><br>", downLeft || downRight)
				text += fmt.Sprintf("upLeft: <b>%v</b><br>", upLeft)
				text += fmt.Sprintf("upRight: <b>%v</b><br>", upRight)
				text += fmt.Sprintf("downLeft: <b>%v</b><br>", downLeft)
				text += fmt.Sprintf("downRight: <b>%v</b><br>", downRight)

				divA.Html(text)
			}
		}
	}()

	// English: Add keyboard events in the form of channel.
	// Português: Adiciona eventos do teclado na forma de channel.
	stage.AddListenerKeyUp(&eventOnKeyData)
	stage.AddListenerKeyDown(&eventOnKeyData)

	done := make(chan struct{})
	done <- struct{}{}
}
