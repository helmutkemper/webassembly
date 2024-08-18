package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"net/http"
	"sync"
	"syscall/js"
	"time"
)

/*
Level ID: 0
Resolution: 156543.03392800014
Level ID: 1
Resolution: 78271.51696399994
Level ID: 2
Resolution: 39135.75848200009
Level ID: 3
Resolution: 19567.87924099992
Level ID: 4
Resolution: 9783.93962049996
Level ID: 5
Resolution: 4891.96981024998
Level ID: 6
Resolution: 2445.98490512499
Level ID: 7
Resolution: 1222.992452562495
Level ID: 8
Resolution: 611.4962262813797
Level ID: 9
Resolution: 305.74811314055756
Level ID: 10
Resolution: 152.87405657041106
Level ID: 11
Resolution: 76.43702828507324
Level ID: 12
Resolution: 38.21851414253662
Level ID: 13
Resolution: 19.10925707126831
Level ID: 14
Resolution: 9.554628535634155
Level ID: 15
Resolution: 4.77731426794937
Level ID: 16
Resolution: 2.388657133974685
Level ID: 17
Resolution: 1.1943285668550503
Level ID: 18
Resolution: 0.5971642835598172
Level ID: 19
Resolution: 0.29858214164761665
Level ID: 20
Resolution: 0.14929107082380833
Level ID: 21
Resolution: 0.07464553541190416
Level ID: 22
Resolution: 0.03732276770595208
Level ID: 23
Resolution: 0.01866138385297604
*/

const (
	tileSize = 256
)

type CacheType js.Value

type OsmCache struct {
	testX int
	testY int

	chZoom     chan int
	z          int
	isDrawing  bool
	isCentered bool
	isDragging bool

	cacheData     map[int]map[int]map[int]CacheType
	cacheUrl      map[string]struct{}
	cacheUrlMutex sync.RWMutex
	url           string

	gridLine   int
	gridColor  color.Color
	gridEnable bool

	longitude float64
	latitude  float64
	zoom      int

	usefulWindowWidth  int
	usefulWindowHeight int

	tileXIntegral   float64 // tileXIntegral tile x usado no cálculo da url com valor integral. A formula calcula um valor inteiro usado para formar a url e baixar a imagem, este é o valor integral, antes da conversão.
	tileYIntegral   float64 // tileYIntegral tile y usado no cálculo da url com valor integral. A formula calcula um valor inteiro usado para formar a url e baixar a imagem, este é o valor integral, antes da conversão.
	tileXOsmUrl     int     // tileXOsmUrl tile x usado no cálculo da url com valor convertido. A formula calcula um valor inteiro usado para formar a url e baixar a imagem, este é o valor convertido para url.
	tileYOsmUrl     int     // tileYOsmUrl tile y usado no cálculo da url com valor convertido. A formula calcula um valor inteiro usado para formar a url e baixar a imagem, este é o valor convertido para url.
	tileXPercentual float64 // tileXPercentual tile x usado no cálculo da url com valor percentual. A formula calcula um valor inteiro usado para formar a url e baixar a imagem, este é o valor percentual, ou seja x percentual * tile size é o valor x dentro do tile para a coordenada.
	tileYPercentual float64 // tileYPercentual tile y usado no cálculo da url com valor percentual. A formula calcula um valor inteiro usado para formar a url e baixar a imagem, este é o valor percentual, ou seja y percentual * tile size é o valor y dentro do tile para a coordenada.
	tileHCentral    int     // tileHCentral tile horizontal do centro do mapa. Ex.: Se o mapa tem 7 tiles, tileHCentral = 4
	tileVCentral    int     // tileVCentral tile vertical do centro do mapa. Ex.: Se o mapa tem 7 tiles, tileVCentral = 4

	mapWidth                int // mapWidth internal size width of the map div
	mapHeight               int // mapHeight internal size height of the map div
	horizontalQuantityTile  int // horizontalQuantityTile quantidade de tiles horizontais usados na imagem
	verticalQuantityTile    int // verticalQuantityTile quantidade de tiles verticais usados na imagem
	startXMouseMove         int // startXMouseMove captura o x inicial da div maps para o dragging do mouse
	startYMouseMove         int // startYMouseMove captura o y inicial da div maps para o dragging do mouse
	totalXMovementMouseMove int
	totalYMovementMouseMove int

	canvasTag *html.TagCanvas
	mapTag    *html.TagDiv
	maskTag   *html.TagDiv
}

func (e *OsmCache) GetJsImage(x, y, z int) (found bool, img js.Value) {
	var data CacheType
	data, found = e.cacheData[z][x][y]

	if !found {
		e.AddTile(x, y, z)
		data, found = e.cacheData[z][x][y]
	}

	return found, js.Value(data)
}

func (e *OsmCache) ClearX(x, z int) {
	delete(e.cacheData[z], x)
}

func (e *OsmCache) ClearY(y, z int) {
	for x := range e.cacheData {
		delete(e.cacheData[z][x], y)
	}
}

func (e *OsmCache) ClearZ(z int) {
	delete(e.cacheData, z)
}

func (e *OsmCache) Add(x, y, z int, img CacheType) {
	if e.cacheData[z] == nil {
		e.cacheData[z] = make(map[int]map[int]CacheType)
	}

	if e.cacheData[z][x] == nil {
		e.cacheData[z][x] = make(map[int]CacheType)
	}

	e.cacheData[z][x][y] = img
}

func (e *OsmCache) addCacheUrl(url string) (found bool) {
	e.cacheUrlMutex.Lock()
	defer e.cacheUrlMutex.Unlock()

	if _, found = e.cacheUrl[url]; found {
		return
	}
	e.cacheUrl[url] = struct{}{}

	return
}

func (e *OsmCache) AddTile(x, y, z int) {
	if _, found := e.cacheData[z][x][y]; found {
		// todo: flag force update
		return
	}

	var err error
	var req *http.Request
	var resp *http.Response
	var img image.Image

	if e.url == "" {
		e.url = "https://tile.openstreetmap.org/%v/%v/%v.png"
	}

	url := fmt.Sprintf(e.url, z, x, y)
	if found := e.addCacheUrl(url); found {
		return
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return
	}
	//req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("Access-Control-Allow-Methods", "GET")
	req.Header.Set("Access-Control-Allow-Headers", "*")

	client := http.Client{
		Timeout: 15 * time.Second, // todo: permitir setup
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return
	}

	//resp, err = http.Get(fmt.Sprintf(e.url, z, x, y))
	//if err != nil {
	//	log.Fatalf("Erro ao carregar a imagem: %v", err)
	//}
	defer func() {
		_ = resp.Body.Close()
	}()

	img, err = png.Decode(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao decodificar a imagem: %v", err)
	}

	jsImg := js.Global().Get("Image").New()
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		log.Printf("Erro ao bufferizar a imagem: %v", err)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	jsImg.Set("src", "data:image/png;base64,"+base64.StdEncoding.EncodeToString(buf.Bytes()))
	jsImg.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		wg.Done()
		return nil
	}))
	wg.Wait()
	e.Add(x, y, z, CacheType(jsImg))

	return
}

func (e *OsmCache) mouseDownFunc(_ js.Value, args []js.Value) any {
	event := args[0]
	e.isDragging = true

	// Pega o viewPoint da tela de onde o click aconteceu.
	e.startXMouseMove = event.Get("clientX").Int()
	e.startYMouseMove = event.Get("clientY").Int()

	// Altera o cursor
	e.mapTag.AddStyle("cursor", "grabbing")
	return nil
}

func (e *OsmCache) mouseMoveFunc(_ js.Value, args []js.Value) any {
	if !e.isDragging {
		return nil
	}

	event := args[0]

	// Pega o viewPoint da tela de onde o movimento aconteceu.
	var actualX = event.Get("clientX").Int()
	var actualY = event.Get("clientY").Int()

	// Calcula o delta de deslocamento em relação ao ponto inicial
	var deltaX = e.startXMouseMove - actualX
	var deltaY = e.startYMouseMove - actualY

	e.testX += deltaX
	e.testY += deltaY

	e.mapTag.AddStyle("left", fmt.Sprintf("%vpx", e.mapTag.GetStyleInt("left")-deltaX))
	e.mapTag.AddStyle("top", fmt.Sprintf("%vpx", e.mapTag.GetStyleInt("top")-deltaY))

	// Atualiza o viewPoint da tela de onde o click aconteceu para onde o mouse está
	e.startXMouseMove = actualX
	e.startYMouseMove = actualY

	e.isCentered = false

	return nil
}

func (e *OsmCache) mouseUpFunc(_ js.Value, _ []js.Value) any {
	e.isDragging = false
	e.mapTag.AddStyle("cursor", "grab")
	return nil
}

func (e *OsmCache) dblClickWindowFunc(_ js.Value, _ []js.Value) any {

	e.testX = int(e.tileXIntegral + float64(e.testX))
	e.testY = int(e.tileYIntegral + float64(e.testY))

	dLongitude := (360.0 / math.Pow(2, float64(e.zoom))) * (float64(e.testX))

	log.Printf("(%v,%v)", e.longitude, e.latitude)
	log.Printf("(%v,%v)", e.longitude+dLongitude, e.latitude)

	e.chZoom <- e.zoom
	////log.Printf("(%v,%v)", e.longitude, e.latitude)
	//e.isCentered = false
	//e.calculate()
	//e.Centralize()

	e.testX = 0
	e.testY = 0
	return nil
}

func (e *OsmCache) clickWindowFunc(_ js.Value, args []js.Value) any {
	events := args[0]

	if e.canvasTag.GetId() != events.Get("target").Get("id").String() {
		e.z = 0
		return nil
	}

	e.z = 1

	return nil
}

func (e *OsmCache) keyUpWindowFunc(_ js.Value, args []js.Value) any {
	event := args[0]

	if e.isDrawing {
		return nil
	}

	if e.z == 0 {
		return nil
	}

	switch event.Get("key").String() {
	case "-":
		e.z = -1
	case "=":
		e.z = 1
	case "+":
		e.z = 1
	case "c":
		e.Centralize()
		return nil
	case "C":
		e.Centralize()
		return nil
	default:
		return nil
	}

	e.chZoom <- e.zoom + e.z

	return nil
}

func (e *OsmCache) resizeWindowFunc(_ js.Value, _ []js.Value) any {
	e.resizeScreen()
	e.calculate()

	return nil
}

func (e *OsmCache) SetGrid(enable bool, line int, color color.Color) {
	e.gridEnable = enable
	e.gridLine = line
	e.gridColor = color
}

func (e *OsmCache) Centralize() {

	e.prepareCenterMap()

	e.mapTag.AddStyle("left", fmt.Sprintf("%vpx", e.totalXMovementMouseMove))
	e.mapTag.AddStyle("top", fmt.Sprintf("%vpx", e.totalYMovementMouseMove))

	e.isCentered = true

	e.canvasTag.BeginPath()
	e.canvasTag.LineWidth(5)
	e.canvasTag.StrokeStyle(color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff})
	e.canvasTag.Arc((e.tileHCentral-1)*tileSize+int(e.tileXPercentual*float64(tileSize)), (e.tileVCentral-1)*tileSize+int(e.tileYPercentual*float64(tileSize)), 10, 0, 2.0*math.Pi, false)
	e.canvasTag.Stroke()
}

func (e *OsmCache) prepareMapTagSize() {
	// https://developer.mozilla.org/en-US/docs/Web/API/Element/clientWidth
	e.mapWidth = e.maskTag.Get().Get("clientWidth").Int()
	e.mapHeight = e.maskTag.Get().Get("clientHeight").Int()
}

func (e *OsmCache) prepareWindowSize() {
	e.usefulWindowWidth = js.Global().Get("document").Get("documentElement").Get("clientWidth").Int()
	e.usefulWindowHeight = js.Global().Get("document").Get("documentElement").Get("clientHeight").Int()
}

func (e *OsmCache) prepareOsmTile() {
	e.tileXIntegral = (e.longitude + 180.0) / 360.0 * math.Pow(2, float64(e.zoom))
	e.tileYIntegral = (1.0 - (math.Log(math.Tan(e.latitude*math.Pi/180.0)+1.0/math.Cos(e.latitude*math.Pi/180.0)) / math.Pi)) / 2.0 * math.Pow(2, float64(e.zoom))

	e.tileXOsmUrl = int(math.Floor(e.tileXIntegral))
	e.tileYOsmUrl = int(math.Floor(e.tileYIntegral))

	e.tileXPercentual = e.tileXIntegral - float64(e.tileXOsmUrl)
	e.tileYPercentual = e.tileYIntegral - float64(e.tileYOsmUrl)
}

func (e *OsmCache) prepareCenterMap() {
	e.totalXMovementMouseMove = e.mapWidth/2 - ((e.tileHCentral-1)*tileSize + int(e.tileXPercentual*float64(tileSize)))
	e.totalYMovementMouseMove = e.mapHeight/2 - ((e.tileVCentral-1)*tileSize + int(e.tileYPercentual*float64(tileSize)))
}

func (e *OsmCache) calculateTotalTilesImage() {
	e.horizontalQuantityTile = int(math.Ceil(float64(e.mapWidth) / float64(tileSize)))
	if e.mapWidth <= tileSize {
		e.horizontalQuantityTile += 1
	}

	e.verticalQuantityTile = int(math.Ceil(float64(e.mapHeight) / float64(tileSize)))
	if e.mapHeight <= tileSize {
		e.verticalQuantityTile += 1
	}

	e.horizontalQuantityTile *= 2 // todo: configurar
	e.verticalQuantityTile *= 2   // todo: configurar

	if e.horizontalQuantityTile%2 == 0 {
		e.horizontalQuantityTile += 1
	}

	if e.verticalQuantityTile%2 == 0 {
		e.verticalQuantityTile += 1
	}

	e.tileHCentral = int(math.Ceil(float64(e.horizontalQuantityTile) / 2.0))
	e.tileVCentral = int(math.Ceil(float64(e.verticalQuantityTile) / 2.0))
}

func (e *OsmCache) calculate() {

	e.prepareOsmTile()
	e.prepareMapTagSize()
	e.calculateTotalTilesImage()

	e.canvasTag.SetSize(e.horizontalQuantityTile*tileSize, e.verticalQuantityTile*tileSize)

	e.prepareCenterMap()

	if e.isCentered {
		e.Centralize()
	}

	e.isDrawing = true
	defer func() {
		e.isDrawing = false
	}()

	for v := 0; v != e.verticalQuantityTile; v += 1 {
		for h := 0; h != e.horizontalQuantityTile; h += 1 {
			x := e.tileXOsmUrl + h - e.tileHCentral + 1
			y := e.tileYOsmUrl + v - e.tileVCentral + 1
			go func(x, y, z int) {
				e.AddTile(x, y, z)
			}(x, y, e.zoom)
		}
	}

	for {
		var pass = true
		for v := 0; v != e.verticalQuantityTile; v += 1 {
			for h := 0; h != e.horizontalQuantityTile; h += 1 {
				x := e.tileXOsmUrl + h - e.tileHCentral + 1
				y := e.tileYOsmUrl + v - e.tileVCentral + 1
				found, jsImg := e.GetJsImage(x, y, e.zoom)
				if !found {
					pass = false
					continue
				}

				e.canvasTag.DrawImage(jsImg, h*tileSize, v*tileSize, tileSize, tileSize)

				if e.gridEnable {
					e.canvasTag.StrokeStyle(e.gridColor).
						LineWidth(e.gridLine).
						StrokeRect(h*tileSize, v*tileSize, tileSize, tileSize)
				}
			}
		}

		if pass {
			break
		}
		time.Sleep(time.Nanosecond)
	}

	e.canvasTag.BeginPath()
	e.canvasTag.LineWidth(5)
	e.canvasTag.StrokeStyle(color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff})
	e.canvasTag.Arc((e.tileHCentral-1)*tileSize+int(e.tileXPercentual*float64(tileSize)), (e.tileVCentral-1)*tileSize+int(e.tileYPercentual*float64(tileSize)), 10, 0, 2.0*math.Pi, false)
	e.canvasTag.Stroke()
}

func (e *OsmCache) LoadOsm(longitude, latitude float64, zoom int) {

	e.longitude = longitude
	e.latitude = latitude
	e.zoom = zoom

	e.calculate()

	go func() {
		for {
			select {
			case value := <-e.chZoom:
				//if value == e.zoom {
				//	continue
				//}

				if value <= 5 {
					value = 5
				}

				if value >= 19 {
					value = 19
				}

				e.zoom = value
				e.calculate()
			}
		}
	}()
}

func (e *OsmCache) resizeScreen() {
	e.prepareWindowSize()

	e.maskTag.AddStyle("width", fmt.Sprintf("%vpx", e.usefulWindowWidth))
	e.maskTag.AddStyle("height", fmt.Sprintf("%vpx", e.usefulWindowHeight))
}

func (e *OsmCache) Init(screen bool) (div *html.TagDiv) {
	e.chZoom = make(chan int)
	e.cacheUrl = make(map[string]struct{})
	e.cacheData = make(map[int]map[int]map[int]CacheType)
	e.z = -1
	e.isCentered = true

	e.canvasTag = factoryBrowser.NewTagCanvas(256, 256).Id("canvas")
	e.mapTag = factoryBrowser.NewTagDiv().Id("map").Class("draggable-content").Append(e.canvasTag)
	e.maskTag = factoryBrowser.NewTagDiv().Id("mask").Class("fixed-div").Append(e.mapTag)

	if screen {
		e.resizeScreen()
	}

	e.mapTag.Get().Call("addEventListener", "mousedown", js.FuncOf(e.mouseDownFunc))
	e.mapTag.Get().Call("addEventListener", "mousemove", js.FuncOf(e.mouseMoveFunc))
	e.mapTag.Get().Call("addEventListener", "mouseup", js.FuncOf(e.mouseUpFunc))
	e.mapTag.Get().Call("addEventListener", "mouseout", js.FuncOf(e.mouseUpFunc))
	js.Global().Get("window").Call("addEventListener", "click", js.FuncOf(e.clickWindowFunc))
	js.Global().Get("window").Call("addEventListener", "dblclick", js.FuncOf(e.dblClickWindowFunc))
	js.Global().Get("window").Call("addEventListener", "keyup", js.FuncOf(e.keyUpWindowFunc))
	js.Global().Get("window").Call("addEventListener", "resize", js.FuncOf(e.resizeWindowFunc))
	//e.mapTag.Get().Call("addEventListener", "click", js.FuncOf(e.mouseScrollFunc))

	return e.maskTag
}

func main() {

	var cache = new(OsmCache)
	osm := cache.Init(true)

	stage := factoryBrowser.NewStage()
	stage.Append(osm)

	var lat = 68.96301643204939
	var lon = 33.07737274977919
	//max: 19 min: 5
	var zoom = 19

	cache.SetGrid(true, 1, factoryColor.NewGrayHalfTransparent())
	cache.LoadOsm(lon, lat, zoom)
	cache.Centralize()

	select {}
}
