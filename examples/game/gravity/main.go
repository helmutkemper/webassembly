package main

import (
	keyboard "github.com/helmutkemper/iotmaker.webassembly/browser/event/keyBoard"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"log"
	"time"
)

const (
	KGravity = 0.00000004
	KVMax    = 6
)

func main() {

	var mainWidth = 2000
	var mainHeight = 1500

	// English: channel triggered by a keyboard event.
	// Português: channel disparado por um evento do teclado.
	eventOnKeyData := make(chan keyboard.Data)

	// English: Event triggered by sprite image change.
	// Português: Evento disparado por mudança na imagem da sprite.
	//eventOnChange := make(chan struct{})

	stage := factoryBrowser.NewStage()

	// English: When working with canvas, it is a good practice to use layers, for example, one layer for the background
	// image and another for the character.
	// Português: Quando trabalhar com canvas, é uma boa prática, usar camadas, por exemplo, uma camada para a imagem de
	// fundo e outra para o personargem.
	backgroundCanvas := factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight()).
		StrokeStyle(factoryColor.NewRed()).
		StrokeRect(0, 0, mainWidth, mainHeight).
		SetXY(0, 0)

	playerCanvas := factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight()).
		SetXY(0, 0).
		BeginPath().
		LineWidth(2.0).
		MoveTo(0, 60).
		LineTo(mainWidth, 60).
		Stroke()

	stage.Append(backgroundCanvas, playerCanvas)

	img1 := factoryBrowser.NewTagImg().Src("./platformPack_tilesheet.png", true)
	csvGround := `-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,0,0,0,0,0,0,0,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,0,0,0
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,0,0,0,0,3,3,3
-1,-1,-1,-1,-1,-1,-1,0,0,0,0,3,3,3,3,3,3,3
-1,-1,-1,0,0,0,0,3,3,3,3,3,3,3,3,3,3,3`
	csvWather := `-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
4,4,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1`
	cvsPositive := `-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,52,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,66,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1`
	csvNegative := `-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,25,-1,-1,25,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1`

	sptTiles := html.SpriteTiles{}
	err := sptTiles.AddCsv("groundPositive", cvsPositive, img1, 64, 64)
	if err != nil {
		log.Printf("SpriteTiles.AddCsv().error: %v", err)
		panic(nil)
	}
	err = sptTiles.AddCsv("groundNegative", csvNegative, img1, 64, 64)
	if err != nil {
		log.Printf("SpriteTiles.AddCsv().error: %v", err)
		panic(nil)
	}
	err = sptTiles.AddCsv("ground", csvGround, img1, 64, 64)
	if err != nil {
		log.Printf("SpriteTiles.AddCsv().error: %v", err)
		panic(nil)
	}
	err = sptTiles.AddCsv("groundWather", csvWather, img1, 64, 64)
	if err != nil {
		log.Printf("SpriteTiles.AddCsv().error: %v", err)
		panic(nil)
	}

	err = sptTiles.Draw(backgroundCanvas, "groundPositive")
	if err != nil {
		log.Printf("SpriteTiles.Draw().error: %v", err)
		panic(nil)
	}
	err = sptTiles.Draw(backgroundCanvas, "groundNegative")
	if err != nil {
		log.Printf("SpriteTiles.Draw().error: %v", err)
		panic(nil)
	}
	err = sptTiles.Draw(backgroundCanvas, "ground")
	if err != nil {
		log.Printf("SpriteTiles.Draw().error: %v", err)
		panic(nil)
	}
	err = sptTiles.Draw(backgroundCanvas, "groundWather")
	if err != nil {
		log.Printf("SpriteTiles.Draw().error: %v", err)
		panic(nil)
	}

	// English: Sprite factory containing the canvas element, the image and the size of each frame.
	// Português: Fábrica da sprite contendo o elemento canvas, a imagem e o tamanho de cada quadro.
	spt := &html.SpritePlayer{}
	spt.Init(
		stage,
		playerCanvas,
		"./player_big.png",
		48,
		60,
	).
		Corners(0, mainWidth-48, 0, mainHeight-60)

	// English: Pointer to the event channel informing the sprite frame change.
	// Português: Ponteiro para o channel do evento informando a mudança de quadro da sprite.
	//OnChange(&eventOnChange)

	// English: Using scale(-1,1) and drawImage() commands requires a much higher computational cost than pre-processing
	// the image and then using cache memory, so images are the first thing done, before executing the code.
	// Português: Usar comandos scale(-1,1) e drawImage() requerem um custo computacional muito mais elevado do que
	// pré-processar a imagem e depois usar uma memória cache, por isto, as imagens são a primeira coisa feita, antes de executar o código.

	// English: Adds a scene named "stopped".
	// Português: Adiciona uma cena de nome "parado".
	spt.CreateStoppedRightSide().
		// English: Each frame is formed by the column and row where the frame is, the time interval for the frame to remain
		// on the screen and the information if the image should have the direction inverted.
		// Português: Cada quadro é formado pela coluna e linha onde o quadro se encontra, o intervalo de tempo para o
		// quadro permanecer na tela e a informação se a imagem deve ter o sentido invertido.
		Add(8, 0, 0, false, false)

	spt.CreateStoppedLeftSide().
		Add(8, 0, 0, true, false)

	spt.CreateWalkingRightSide().
		Add(0, 0, 100*time.Millisecond, false, false).
		Add(1, 0, 100*time.Millisecond, false, false).
		Add(2, 0, 100*time.Millisecond, false, false).
		Add(3, 0, 100*time.Millisecond, false, false)

	spt.CreateWalkingLeftSide().
		Add(0, 0, 100*time.Millisecond, true, false).
		Add(1, 0, 100*time.Millisecond, true, false).
		Add(2, 0, 100*time.Millisecond, true, false).
		Add(3, 0, 100*time.Millisecond, true, false)

	spt.CreateStoppedLeftSide().
		Add(8, 0, 0, false, false)

	spt.CreateStoppedLeftSide().
		Add(8, 0, 0, true, false)

	// English: Adds a function to clear the canvas before drawing.
	// Português: Adiciona uma função para limpar o canvas antes do desenho.
	stage.AddDrawFunctions(func() {
		playerCanvas.ClearRect(0, 0, mainWidth, mainHeight)
	})
	// English: Adds the sprite drawing function to the canvas drawing function.
	// Português: Adiciona a função de desenho do sprite a função de desenho do canvas.
	stage.AddDrawFunctions(spt.Draw)

	stage.AddMathFunctions(func() {
		if sptTiles.TestCollisionBox(spt, "groundPositive") == true {
			log.Printf("groundPositive")
		}
	})

	// English: Whenever any function uses channel to communicate with a thread, ensure that the thread is working before
	// the information is written to the channel, or the code may crash.
	// Português: Sempre que alguma função use channel para se comunicar com uma thread, garanta que a thread esteja
	// funcionando antes da informação ser escrita no channel, ou o código pode travar.
	go func() {

		for {
			select {

			case data := <-eventOnKeyData:
				if data.Repeat == true {
					continue
				}

				switch data.EventName {
				case "keyup":
					switch data.Code {
					case "ArrowUp":
						fallthrough
					case "ArrowDown":
						spt.KeyVerticalUp()

					case "ArrowLeft":
						fallthrough
					case "ArrowRight":
						spt.KeyHorizontalUp()
					default:
						continue
					}

					spt.ControlDelta()
					if spt.PlayerStop() == true {
						continue
					}

				case "keydown":
					//log.Printf("%v", data.Code)

					switch data.Code {
					case "ArrowUp":
						if spt.PlayerUp() == true {
							continue
						}

					case "ArrowDown":
						if spt.PlayerDown() == true {
							continue
						}

					case "ArrowLeft":
						if spt.PlayerLeft() == true {
							continue
						}

					case "ArrowRight":
						if spt.PlayerRight() == true {
							continue
						}
					case "Space":
						spt.Gravity()

					default:
						continue
					}
				}
			}

			spt.PlayerImageToUse()
		}
	}()

	spt.StartStoppedRightSide()

	// English: Add keyboard events in the form of channel.
	// Português: Adiciona eventos do teclado na forma de channel.
	stage.AddListenerKeyUp(&eventOnKeyData)
	stage.AddListenerKeyDown(&eventOnKeyData)

	done := make(chan struct{})
	done <- struct{}{}
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
//
//
//
//
