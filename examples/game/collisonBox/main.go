//go:build js

package main

import (
	"fmt"
	keyboard "github.com/helmutkemper/webassembly/browser/event/keyBoard"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"image/color"
	"log"
)

func main() {

	text := "please, use arrow keys and space bar"

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

	groundBox := &html.CollisionBox{}
	groundBox.Init(groundData, w, h)
	groundBox.X(200)
	groundBox.Y(200)

	canvasA.PutImageData(groundBox.GetData(), groundBox.GetX(), groundBox.GetY())

	go func() {
		positionX := 200
		positionY := 140
		imageBoxColor := color.RGBA{A: 255}
		collisionBoxColor := color.RGBA{R: 255, A: 0}
		useOptimized := false

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
					case "Digit0":
						imageBoxColor = color.RGBA{A: 255}
						collisionBoxColor = color.RGBA{R: 255, A: 0}
						useOptimized = false
					case "Digit1":
						imageBoxColor = color.RGBA{A: 0}
						collisionBoxColor = color.RGBA{R: 255, A: 255}
						useOptimized = true
					default:
						log.Printf("data.Code: %v", data.Code)
					}
				}

				canvasB.DrawImageComplete(img2, i*w, 0*h, w, h, 0, 0, w, h)
				playerData := canvasB.GetImageData(0, 0, w, h, false, false)
				canvasB.ClearRect(0*w, 0*h, w, h)

				playerBox := new(html.CollisionBox)
				playerBox.Debug(&imageBoxColor, &collisionBoxColor)
				playerBox.Init(playerData, w, h)
				playerBox.UseOptimized(useOptimized)
				playerBox.X(positionX)
				playerBox.Y(positionY)

				canvasB.ClearRect(0, 0, 800, 600)
				canvasB.PutImageData(playerBox.GetData(), playerBox.GetX(), playerBox.GetY())

				text = "please, use arrow keys to move the player, 0 or 1 to change collision box type and space bar to change the image<br><br>Player collision:<br>"

				upPx, rightPx, downPx, leftPx := playerBox.DistanceCorrection(groundBox)
				upLeft, upRight, downLeft, downRight := playerBox.Quadrant(groundBox)

				textUp := "false"
				if upLeft || upRight {
					textUp = "<span style=\"color:red\">true</span>"
				}

				textLeft := "false"
				if upLeft || downLeft {
					textLeft = "<span style=\"color:red\">true</span>"
				}

				textRight := "false"
				if upRight || downRight {
					textRight = "<span style=\"color:red\">true</span>"
				}

				textDown := "false"
				if downLeft || downRight {
					textDown = "<span style=\"color:red\">true</span>"
				}

				textUpLeft := "false"
				if upLeft {
					textUpLeft = "<span style=\"color:red\">true</span>"
				}

				textUpRight := "false"
				if upRight {
					textUpRight = "<span style=\"color:red\">true</span>"
				}

				textDownLeft := "false"
				if downLeft {
					textDownLeft = "<span style=\"color:red\">true</span>"
				}

				textDownRight := "false"
				if downRight {
					textDownRight = "<span style=\"color:red\">true</span>"
				}

				text += fmt.Sprintf("up [%vpx]: <b>%v</b><br>", upPx, textUp)
				text += fmt.Sprintf("left [%vpx]: <b>%v</b><br>", leftPx, textLeft)
				text += fmt.Sprintf("right [%vpx]: <b>%v</b><br>", rightPx, textRight)
				text += fmt.Sprintf("down [%vpx]: <b>%v</b><br>", downPx, textDown)
				text += fmt.Sprintf("upLeft: <b>%v</b><br>", textUpLeft)
				text += fmt.Sprintf("upRight: <b>%v</b><br>", textUpRight)
				text += fmt.Sprintf("downLeft: <b>%v</b><br>", textDownLeft)
				text += fmt.Sprintf("downRight: <b>%v</b><br>", textDownRight)

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
