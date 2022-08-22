package main

import (
	keyboard "github.com/helmutkemper/iotmaker.webassembly/browser/event/keyBoard"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"log"
	"time"
)

func main() {
	var err error

	// English: channel triggered by a keyboard event.
	// Português: channel disparado por um evento do teclado.
	eventOnKeyData := make(chan keyboard.Data)

	// English: Event triggered by sprite image change.
	// Português: Evento disparado por mudança na imagem da sprite.
	eventOnChange := make(chan struct{})

	stage := factoryBrowser.NewStage()

	backgroundCanvas := factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight()).FillStyle(factoryColor.NewRed()).FillRect(0, 0, 800, 600).SetXY(0, 0)
	playerCanvas := factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight()).SetXY(0, 0)
	stage.Append(backgroundCanvas, playerCanvas)

	spt := factoryBrowser.NewSprite(
		playerCanvas,
		"./player_big.png",
		48,
		60,
	).
		OnChange(&eventOnChange)

	// Português: Adds a scene named "stopped".
	// Português: Adiciona uma cena de nome "parado".
	spt.Scene("stopped").
		Add(8, 0, 0, false, false)

	spt.Scene("normalWalk").
		Add(0, 0, 100*time.Millisecond, false, false).
		Add(1, 0, 100*time.Millisecond, false, false).
		Add(2, 0, 100*time.Millisecond, false, false).
		Add(3, 0, 100*time.Millisecond, false, false)

	spt.Scene("invertedFloor").
		Add(0, 0, 100*time.Millisecond, true, false).
		Add(1, 0, 100*time.Millisecond, true, false).
		Add(2, 0, 100*time.Millisecond, true, false).
		Add(3, 0, 100*time.Millisecond, true, false)

	// English: Adds a function to clear the canvas before drawing.
	// Português: Adiciona uma função para limpar o canvas antes do desenho.
	stage.AddDrawFunctions(func() {
		playerCanvas.ClearRect(0, 0, 800, 600)
	})
	// English: Adds the sprite drawing function to the canvas drawing function.
	// Português: Adiciona a função de desenho do sprite a função de desenho do canvas.
	stage.AddDrawFunctions(spt.Draw)

	go func() {
		x := 0
		delta := 0
		for {
			select {

			case <-eventOnChange:
				x = x + delta
				spt.X(x)

			case data := <-eventOnKeyData:
				if data.Repeat == true {
					continue
				}

				switch data.EventName {
				case "keyup":
					err = spt.Start("stopped")
					if err != nil {
						log.Printf("error: %v", err)
					}
					delta = 0
				case "keydown":
					switch data.Code {
					case "ArrowLeft":
						if spt.GetScene() != "invertedFloor" {
							spt.Scene("stopped").
								Add(8, 0, 0, true, false)

							err = spt.Start("invertedFloor")
							if err != nil {
								log.Printf("error: %v", err)
							}
						}

						delta = -9

					case "ArrowRight":
						if spt.GetScene() != "normalWalk" {
							spt.Scene("stopped").
								Add(8, 0, 0, false, false)

							err = spt.Start("normalWalk")
							if err != nil {
								log.Printf("error: %v", err)
							}
						}

						delta = 9
					}
				}
			}
		}
	}()

	spt.Y(100)
	err = spt.Start("stopped")
	if err != nil {
		log.Printf("error: %v", err)
	}

	stage.AddListenerKeyUp(&eventOnKeyData)
	stage.AddListenerKeyDown(&eventOnKeyData)

	done := make(chan struct{})
	done <- struct{}{}
}
