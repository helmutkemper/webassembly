package main

import (
	keyboard "github.com/helmutkemper/iotmaker.webassembly/browser/event/keyBoard"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"log"
	"time"
)

var (
	spt *html.Sprite

	x              = 0
	y              = 0
	deltaX         = 0
	deltaY         = 0
	lastRightSide  = true
	validKeyHDelta = 0
	validKeyVDelta = 0
)

func controlX() {
	x = x + deltaX

	// English: space width - sprite width
	// Português: comprimento do espaço - comprimento da sprite
	if x > 800-48 {
		x = 800 - 48
	}
	if x < 0 {
		x = 0
	}
}

func controlY() {
	y = y + deltaY

	// English: space heght - sprite height
	// Português: altura do espaço - altura da sprite
	if y > 600-60 {
		y = 600 - 60
	}
	if y < 0 {
		y = 0
	}
}

func controlDelta() {
	if validKeyHDelta <= 0 {
		deltaX = 0
		validKeyHDelta = 0
	}

	if validKeyVDelta <= 0 {
		deltaY = 0
		validKeyVDelta = 0
	}
}

func playerStop() (cont bool) {
	var err error
	if validKeyHDelta <= 0 && validKeyVDelta <= 0 {
		if lastRightSide == true {
			err = spt.Start("normalStopped")
		} else {
			err = spt.Start("invertedStopped")
		}
		if err != nil {
			log.Printf("error: %v", err)
		}

		return true
	}

	return false
}

func playerUp() (cont bool) {
	if validKeyVDelta > 0 {
		return true
	}

	validKeyVDelta += 1
	deltaY = -9
	return false
}

func playerDown() (cont bool) {
	if validKeyVDelta > 0 {
		return true
	}

	validKeyVDelta += 1
	deltaY = 9
	return false
}

func playerLeft() (cont bool) {
	if validKeyHDelta > 0 {
		return true
	}

	validKeyHDelta += 1
	deltaX = -9
	lastRightSide = false
	return false
}

func playerRight() (cont bool) {
	if validKeyHDelta > 0 {
		return true
	}

	validKeyHDelta += 1
	deltaX = 9
	lastRightSide = true
	return false
}

func playerImageToUse() {
	var err error
	if lastRightSide == false {
		err = spt.Start("invertedWalking")
		if err != nil {
			log.Printf("error: %v", err)
		}
	} else {
		err = spt.Start("normalWalking")
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
}

func main() {
	var err error

	// English: channel triggered by a keyboard event.
	// Português: channel disparado por um evento do teclado.
	eventOnKeyData := make(chan keyboard.Data)

	// English: Event triggered by sprite image change.
	// Português: Evento disparado por mudança na imagem da sprite.
	eventOnChange := make(chan struct{})

	stage := factoryBrowser.NewStage()

	// English: When working with canvas, it is a good practice to use layers, for example, one layer for the background
	// image and another for the character.
	// Português: Quando trabalhar com canvas, é uma boa prática, usar camadas, por exemplo, uma camada para a imagem de
	// fundo e outra para o personargem.
	backgroundCanvas := factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight()).
		StrokeStyle(factoryColor.NewRed()).
		StrokeRect(0, 0, 800, 600).
		SetXY(0, 0)

	playerCanvas := factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight()).
		SetXY(0, 0)

	stage.Append(backgroundCanvas, playerCanvas)

	// English: Sprite factory containing the canvas element, the image and the size of each frame.
	// Português: Fábrica da sprite contendo o elemento canvas, a imagem e o tamanho de cada quadro.
	spt = factoryBrowser.NewSprite(
		playerCanvas,
		"./player_big.png",
		48,
		60,
	).
		// English: Pointer to the event channel informing the sprite frame change.
		// Português: Ponteiro para o channel do evento informando a mudança de quadro da sprite.
		OnChange(&eventOnChange)

	// English: Using scale(-1,1) and drawImage() commands requires a much higher computational cost than pre-processing
	// the image and then using cache memory, so images are the first thing done, before executing the code.
	// Português: Usar comandos scale(-1,1) e drawImage() requerem um custo computacional muito mais elevado do que
	// pré-processar a imagem e depois usar uma memória cache, por isto, as imagens são a primeira coisa feita, antes de executar o código.

	// English: Adds a scene named "stopped".
	// Português: Adiciona uma cena de nome "parado".
	spt.Scene("normalStopped").
		// English: Each frame is formed by the column and row where the frame is, the time interval for the frame to remain
		// on the screen and the information if the image should have the direction inverted.
		// Português: Cada quadro é formado pela coluna e linha onde o quadro se encontra, o intervalo de tempo para o
		// quadro permanecer na tela e a informação se a imagem deve ter o sentido invertido.
		Add(8, 0, 0, false, false)

	spt.Scene("invertedStopped").
		Add(8, 0, 0, true, false)

	spt.Scene("normalWalking").
		Add(0, 0, 100*time.Millisecond, false, false).
		Add(1, 0, 100*time.Millisecond, false, false).
		Add(2, 0, 100*time.Millisecond, false, false).
		Add(3, 0, 100*time.Millisecond, false, false)

	spt.Scene("invertedWalking").
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

	// English: Whenever any function uses channel to communicate with a thread, ensure that the thread is working before
	// the information is written to the channel, or the code may crash.
	// Português: Sempre que alguma função use channel para se comunicar com uma thread, garanta que a thread esteja
	// funcionando antes da informação ser escrita no channel, ou o código pode travar.
	go func() {

		for {
			select {

			case <-eventOnChange:
				controlX()
				spt.X(x)

				controlY()
				spt.Y(y)

			case data := <-eventOnKeyData:
				if data.Repeat == true {
					continue
				}

				switch data.EventName {
				case "keyup":
					switch data.Code {
					case "ArrowUp":
						validKeyVDelta -= 1
					case "ArrowDown":
						validKeyVDelta -= 1
					case "ArrowLeft":
						validKeyHDelta -= 1
					case "ArrowRight":
						validKeyHDelta -= 1
					}

					controlDelta()
					if playerStop() == true {
						continue
					}

				case "keydown":
					switch data.Code {
					case "ArrowUp":
						if playerUp() == true {
							continue
						}

					case "ArrowDown":
						if playerDown() == true {
							continue
						}

					case "ArrowLeft":
						if playerLeft() == true {
							continue
						}

					case "ArrowRight":
						if playerRight() == true {
							continue
						}

					default:
						continue
					}
				}
			}

			playerImageToUse()
		}
	}()

	err = spt.Start("normalStopped")
	if err != nil {
		log.Printf("error: %v", err)
	}

	// English: Add keyboard events in the form of channel.
	// Português: Adiciona eventos do teclado na forma de channel.
	stage.AddListenerKeyUp(&eventOnKeyData)
	stage.AddListenerKeyDown(&eventOnKeyData)

	done := make(chan struct{})
	done <- struct{}{}
}
