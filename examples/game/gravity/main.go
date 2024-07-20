package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"log"
	"strconv"
	"time"
)

func main() {

	var mainWidth = 2000
	var mainHeight = 1500

	// English: Event triggered by sprite image change.
	// Português: Evento disparado por mudança na imagem da sprite.
	//eventOnChange := make(chan struct{})

	stage := factoryBrowser.NewStage()

	div1 := factoryBrowser.NewTagDiv().SetX(0).SetY(0)
	stage.Append(div1)

	go func() {
		for {
			div1.Html("fps: " + strconv.FormatInt(int64(stage.GetFPS()), 10))
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// English: When working with canvas, it is a good practice to use layers, for example, one layer for the background
	// image and another for the character.
	// Português: Quando trabalhar com canvas, é uma boa prática, usar camadas, por exemplo, uma camada para a imagem de
	// fundo e outra para o personargem.
	backgroundCanvas := factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight())

	playerCanvas := factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight()).
		SetXY(0, 0).
		BeginPath().
		LineWidth(2.0).
		MoveTo(0, 60).
		LineTo(mainWidth, 60).
		Stroke()

	stage.Append(backgroundCanvas, playerCanvas)

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
		X(300).
		DefineFloorVerySlippery().
		Corners(-1, mainWidth-48, -1, mainHeight-60)

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

	spt.CreateFallRightSide().
		Add(8, 0, 0, false, false)

	spt.CreateStoppedLeftSide().
		Add(8, 0, 0, true, false)

	spt.CreateFallLeftSide().
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

	spt.MovieClipStopped()

	// English: Adds a function to clear the canvas before drawing.
	// Português: Adiciona uma função para limpar o canvas antes do desenho.
	stage.AddDrawFunctions(func() {
		playerCanvas.ClearRect(0, 0, mainWidth, mainHeight)
	})
	// English: Adds the sprite drawing function to the canvas drawing function.
	// Português: Adiciona a função de desenho do sprite a função de desenho do canvas.
	stage.AddDrawFunctions(spt.Draw)

	img1 := factoryBrowser.NewTagImg().Src("./platformPack_tilesheet.png", true)
	//img1W := 12 * 64
	//img1H := 7 * 64
	//cv := factoryBrowser.NewTagCanvas(img1.GetWidth(), img1.GetHeight()).
	//	DrawImageTile(img1, 0*64, 0*64, 0, 0, img1W, img1H)
	//data := cv.GetImageDataMatrix(0, 0, img1W, img1H)
	//
	//dataMatrix := make([][]any, len(data))
	//for y := 0; y != len(data); y += 1 {
	//	dataMatrix[y] = make([]any, len(data[y]))
	//	for x := 0; x != len(data[y]); x += 1 {
	//		dataMatrix[y][x] = color.RGBA{
	//			R: data[y][x][0],
	//			G: data[y][x][1],
	//			B: data[y][x][2],
	//			A: data[y][x][3],
	//		}
	//	}
	//}
	//
	//ctr := contour.Contour{}
	//ctr.VerifyFunction(func(pMatrix *[][]any, x, y int) bool {
	//	return (*pMatrix)[y][x].(color.RGBA).A > 100
	//})
	//ctr.PopulateFunction(func(pMatrix *[][]any, x, y int) {
	//	(*pMatrix)[y][x] = color.RGBA{R: 128, G: 0, B: 0, A: 255}
	//})
	//_ = ctr.Init(&dataMatrix)
	//
	//ctr.Verify()
	//
	//dataMatrix = ctr.GetData()
	//for y := 0; y != len(dataMatrix); y += 1 {
	//	for x := 0; x != len(dataMatrix[y]); x += 1 {
	//		if dataMatrix[y][x] == nil {
	//			continue
	//		}
	//
	//		data[y][x][0] = dataMatrix[y][x].(color.RGBA).R
	//		data[y][x][1] = dataMatrix[y][x].(color.RGBA).G
	//		data[y][x][2] = dataMatrix[y][x].(color.RGBA).B
	//		data[y][x][3] = dataMatrix[y][x].(color.RGBA).A
	//	}
	//}
	//
	//cv.ImageData(data, img1W, img1H)
	//stage.Append(cv)

	csvGround := `-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,0,0,0,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0`

	csvWather := `-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1
-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1`

	sptTiles := html.SpriteTiles{}
	sptTiles.Init(stage)
	sptTiles.AddNoEvent(
		spt, 0, spt.FreeFallEnable,
	)
	sptTiles.AddNoEvent(
		spt, 500*time.Millisecond, spt.JumppingDisable,
	)
	//sptTiles.AddNoEvent(spt, spt.Fall)
	stage.AddMathFunctions(sptTiles.Process)

	err := sptTiles.AddCsv("ground", csvGround, img1, true, true, 64, 64).
		AddCollision(spt, spt.FreeFallDisable).
		AddCollision(spt, spt.JumppingEnable).
		AddKeyboard("ArrowRight", spt.RunningRightStart, spt.RunningRightStop).
		AddKeyboard("ArrowLeft", spt.RunningLeftStart, spt.RunningLeftStop).
		AddKeyboard("ArrowUp", spt.JumppingStart, spt.JumppingStop).
		GetError()
	if err != nil {
		log.Printf("SpriteTiles.AddCsv().error: %v", err)
		panic(nil)
	}
	err = sptTiles.AddCsv("groundWather", csvWather, img1, true, true, 64, 64).
		GetError()
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

	// English: Add keyboard events in the form of channel.
	// Português: Adiciona eventos do teclado na forma de channel.
	//stage.AddListenerKeyUp(&eventOnKeyData)
	//stage.AddListenerKeyDown(&eventOnKeyData)

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
