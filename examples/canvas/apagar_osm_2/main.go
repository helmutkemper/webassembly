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
	testX     int
	testY     int
	testTileX int
	testTileY int

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

	log.Printf("%v,%v", e.testX, e.testY)
	return nil
}

func (e *OsmCache) dblClickWindowFunc(_ js.Value, _ []js.Value) any {

	var x = e.testX
	log.Printf("e.testX: %v", e.testX)
	e.testTileX = e.testX / tileSize
	log.Printf("e.testTileX: %v", e.testTileX)
	e.testX = e.testX % tileSize
	log.Printf("e.testX: %v", e.testX)

	go func() {
		e.calculate()
		e.Centralize()
		e.mapTag.AddStyle("left", fmt.Sprintf("%vpx", e.totalXMovementMouseMove+(x*-1)))
	}()

	//e.mapTag.AddStyle("left", fmt.Sprintf("%vpx", e.testX))
	//e.mapTag.AddStyle("top", fmt.Sprintf("%vpx", e.totalYMovementMouseMove))

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
	// https://wiki.openstreetmap.org/wiki/Slippy_map_tilenames#Resolution_and_Scale
	// (ler) https://cfis.savagexi.com/2006/05/03/google-maps-deconstructed/

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

	e.tileHCentral = int(math.Ceil(float64(e.horizontalQuantityTile)/2.0)) + (e.testTileX * -1)
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

	var lon = 33.07737274977919
	var lat = 68.96301643204939
	//max: 19 min: 5
	var zoom = 19

	lon = 33.07805939528697
	lat = 68.96301643067

	//resolution := 1.0 / 0.14929107082380833
	//lonMax := 360.0 / 2.0
	//tiles := math.Sqrt(274877906944.0)
	//lonPerTile := lonMax / tiles
	//lon += resolution

	point := Point{}
	point.SetDegrees(lon, lat)

	cache.SetGrid(true, 1, factoryColor.NewGrayHalfTransparent())
	cache.LoadOsm(lon, lat, zoom)
	cache.Centralize()

	d := HorizontalTileDistance(point, zoom)
	a := Angle{}
	a.SetDegrees(90)
	pointB := DestinationPoint(point, d, a)
	log.Printf("%v,%v", pointB.GetLongitudeAsDegrees(), pointB.GetLatitudeAsDegrees())
	select {}
}

func DestinationPoint(point Point, distance Distance, angle Angle) (destination Point) {
	earthRadius := EarthRadius(point)

	//latB = asin( sin( latA ) * cos( d / R ) +
	//  cos( latA ) * sin( d / R ) * cos( θ ) )
	latitude := math.Asin(math.Sin(point.Rad[LATITUDE])*math.Cos(distance.GetMeters()/earthRadius.GetMeters()) +
		math.Cos(point.Rad[LATITUDE])*math.Sin(distance.GetMeters()/earthRadius.GetMeters())*math.Cos(angle.GetAsRadians()))

	//lonB = lonA + atan2( sin( θ ) *
	//  sin( d / R ) * cos( latA ),
	//  cos( d / R ) − sin( latA ) * sin( latB ) )
	longitude := point.Rad[LONGITUDE] + math.Atan2(math.Sin(DegreesToRadians(angle.GetAsRadians()))*
		math.Sin(distance.GetMeters()/earthRadius.GetMeters())*math.Cos(point.Rad[LATITUDE]),
		math.Cos(distance.GetMeters()/earthRadius.GetMeters())-math.Sin(point.Rad[LATITUDE])*math.Sin(latitude))

	destination.SetRadians(longitude, latitude)
	return
}

func HorizontalTileDistance(point Point, zoom int) (distance Distance) {
	// https://wiki.openstreetmap.org/wiki/Zoom_levels
	earthCircumference := EarthCircumference(point)
	distance.SetMeters(earthCircumference.GetMeters() * math.Cos(point.Rad[LATITUDE]) / math.Pow(2.0, float64(zoom)))
	return
}

func EarthCircumference(point Point) (distance Distance) {
	earthRadius := EarthRadius(point)
	distance.SetMeters(2.0 * math.Pi * earthRadius.GetMeters())
	return
}

func EarthRadius(point Point) (distance Distance) {
	distance.SetMeters(
		math.Sqrt(
			(math.Pow(math.Pow(GeoidalMajor, 2.0)*math.Cos(point.Rad[LATITUDE]), 2.0) +
				math.Pow(math.Pow(GeoidalMinor, 2.0)*math.Sin(point.Rad[LATITUDE]), 2.0)) /
				(math.Pow(GeoidalMajor*math.Cos(point.Rad[LATITUDE]), 2.0) +
					math.Pow(GeoidalMinor*math.Sin(point.Rad[LATITUDE]), 2.0))),
	)

	return
}

func DegreesToRadians(degreesAFlt float64) float64 {
	return math.Pi * degreesAFlt / 180.0
}

func RadiansToDegrees(degreesAFlt float64) float64 {
	return 180.0 * degreesAFlt / math.Pi
}

const (
	GeoidalMajor = Wgs84A
	GeoidalMinor = Wgs84B

	// Wgs84A
	// Geoide WGS84, major semi axis in meters
	// Warning, changing this value affect an innumerable amount of testing
	Wgs84A float64 = 6378137.0

	// Wgs84B
	// Geoide WGS84, minor semi axis in meters
	// Warning, changing this value affect an innumerable amount of testing
	Wgs84B float64 = 6356752.314245

	// DEGREES
	// Degrees symbol for human notation
	// Warning, changing this value affect an innumerable amount of testing
	DEGREES string = "°"

	// MINUTES
	// Minutes from degrees symbol for human notation
	// Warning, changing this value affect an innumerable amount of testing
	MINUTES string = "´"

	// SECONDS
	// Seconds from degrees symbol for human notation
	// Warning, changing this value affect an innumerable amount of testing
	SECONDS string = "´´"

	// RADIANS
	// Rads symbol from human notation
	// Warning, changing this value affect an innumerable amount of testing
	RADIANS string = "rad"

	LONGITUDE = 0
	LATITUDE  = 1
)

type Distance struct {
	Meters       float64 // distance
	Kilometers   float64 // distance
	unit         string  // distance unit
	preserveUnit string  // original unit
}

type DistanceList struct {
	List []Distance
}

// GetMeters Get distance value
func (d *Distance) GetMeters() float64 {
	return d.Meters
}

// GetKilometers Get distance value
func (d *Distance) GetKilometers() float64 {
	return d.Kilometers
}

// GetUnit Get distance unit
func (d *Distance) GetUnit() string {
	return d.unit
}

func (d *Distance) GetOriginalUnit() string {
	return d.preserveUnit
}

// AddMeters Set distance as meters
func (d *Distance) AddMeters(m float64) (ref *Distance) {

	d.Meters += m
	d.Kilometers += m / 1000
	return d
}

// SetMeters Set distance as meters
func (d *Distance) SetMeters(m float64) (ref *Distance) {
	d.Meters = m
	d.Kilometers = m / 1000
	d.unit = "m"
	d.preserveUnit = "m"
	return d
}

func (d *Distance) SetMetersIfGreaterThan(m float64) (ref *Distance) {
	test := math.Max(d.Meters, m)

	d.Meters = test
	d.Kilometers = test / 1000
	d.unit = "m"
	d.preserveUnit = "m"
	return d
}

func (d *Distance) SetKilometersIfGreaterThan(km float64) (ref *Distance) {
	test := math.Max(d.Kilometers, km)

	d.Meters = test * 1000
	d.Kilometers = test
	d.unit = "km"
	d.preserveUnit = "km"
	return d
}

func (d *Distance) SetMetersIfLessThan(m float64) (ref *Distance) {
	test := math.Min(d.Meters, m)

	d.Meters = test
	d.Kilometers = test / 1000
	d.unit = "m"
	d.preserveUnit = "m"
	return d
}

func (d *Distance) SetKilometersIfLessThan(km float64) (ref *Distance) {
	test := math.Min(d.Kilometers, km)

	d.Meters = test * 1000
	d.Kilometers = test
	d.unit = "km"
	d.preserveUnit = "km"
	return d
}

// AddKilometers Set distance as kilometers
func (d *Distance) AddKilometers(km float64) (ref *Distance) {
	d.Meters += km * 1000
	d.Kilometers += km
	d.unit = "Km"
	d.preserveUnit = "Km"
	return d
}

// SetKilometers Set distance as kilometers
func (d *Distance) SetKilometers(km float64) (ref *Distance) {
	d.Meters = km * 1000
	d.Kilometers = km
	d.unit = "Km"
	d.preserveUnit = "Km"
	return d
}

// ToMetersString Get distance as string
func (d *Distance) ToMetersString() string {
	return fmt.Sprintf("%1.2fm", d.Meters)
}

func (d *Distance) ToKilometersString() string {
	return fmt.Sprintf("%1.2fKm", d.Kilometers)
}

type Angle struct {
	Degrees      float64 //Angle
	Radians      float64 //Angle
	Unit         string  //Unit
	PreserveUnit string  //original Unit
}

type AngleList struct {
	List []Angle
}

// SetDecimalDegrees Set Angle value as decimal degrees
func (a *Angle) SetDecimalDegrees(degrees, primes, seconds float64) (ref *Angle) {
	a.Degrees = degrees + primes/60.0 + seconds/3600.0
	a.Radians = DegreesToRadians(degrees + primes/60.0 + seconds/3600.0)
	a.Unit = DEGREES
	a.PreserveUnit = DEGREES
	return a
}

// SetDegrees Set Angle value as degrees
func (a *Angle) SetDegrees(angle float64) (ref *Angle) {
	a.Degrees = angle
	a.Radians = DegreesToRadians(angle)
	a.Unit = DEGREES
	a.PreserveUnit = DEGREES
	return a
}

// SetRadians Set Angle value as radians
func (a *Angle) SetRadians(angle float64) (ref *Angle) {
	a.Radians = angle
	a.Degrees = RadiansToDegrees(angle)
	a.Unit = RADIANS
	a.PreserveUnit = RADIANS
	return a
}

// AddDegrees Set Angle value as degrees
func (a *Angle) AddDegrees(angle float64) (ref *Angle) {
	a.Degrees = a.Degrees + angle
	a.Radians = DegreesToRadians(a.Degrees)
	return a
}

// GetAsRadians Get Angle
func (a *Angle) GetAsRadians() float64 {
	return a.Radians
}

// GetAsDegrees Get Angle
func (a *Angle) GetAsDegrees() float64 {
	return a.Degrees
}

// GetUnit Get Unit
func (a *Angle) GetUnit() string {
	return a.Unit
}

// GetOriginalUnit Get original Unit before conversion
func (a *Angle) GetOriginalUnit() string {
	return a.PreserveUnit
}

// ToDegreesString Get Angle as string
func (a *Angle) ToDegreesString() string {
	return fmt.Sprintf("%1.3f%v", a.Degrees, DEGREES)
}

// ToRadiansString Get Angle as string
func (a *Angle) ToRadiansString() string {
	return fmt.Sprintf("%1.3f%v", a.Radians, RADIANS)
}

type Point struct {
	Id int64 `bson:"id"`
	// Array de localização geográfica.
	// [0:x:longitude,1:y:latitude]
	// Este campo deve obrigatoriamente ser um array devido a indexação do MongoDB
	Loc [2]float64 `bson:"loc"`
	Rad [2]float64 `bson:"rad"`
	// Unidade original do ponto. Serve para manter a resposta no formato original.
	// Versão dentro do Open Street Maps
	Version int64 `bson:"version"`

	Visible bool `bson:"visible"`

	// TimeStamp dentro do Open Street Maps
	TimeStamp time.Time `bson:"timeStamp"`
	// ChangeSet dentro do Open Street Maps
	ChangeSet int64 `bson:"changeSet"`
	// User Id dentro do Open Street Maps
	UId int64 `bson:"userId"`
	// User Name dentro do Open Street Maps
	User string `bson:"-"`
	// Tags do Open Street Maps
	// As Tags contêm _todo tipo de informação, desde como elas foram importadas, ao nome de um estabelecimento comercial,
	// por exemplo.
	Tag map[string]string `bson:"tag"`

	// Dados do usuário
	// Como o GO é fortemente tipado, eu obtive problemas em estender o struct de forma satisfatória e permitir ao usuário
	// do sistema gravar seus próprios dados, por isto, este campo foi criado. Use-o a vontade.
	Data map[string]string `bson:"data"`
	Role string            `bson:"role"`
	// Node usado apenas para o parser do arquivo
	GeoJSon        string `bson:"geoJSon,omitempty"`
	GeoJSonFeature string `bson:"geoJSonFeature"`

	HasKeyValue bool `bson:"hasKeyValue" json:"-"`
}

// SetDegrees Set latitude and longitude as degrees
func (el *Point) SetDegrees(longitude, latitude float64) {
	el.Loc = [2]float64{longitude, latitude}
	el.Rad = [2]float64{DegreesToRadians(longitude), DegreesToRadians(latitude)}
}

// SetDecimalDrees Set latitude and longitude as decimal degrees
func (el *Point) SetDecimalDrees(longitudeDegrees, longitudePrimes, longitudeSeconds, latitudeDegrees, latitudePrimes, latitudeSeconds int64) {
	el.Loc = [2]float64{float64(latitudeDegrees) + float64(latitudePrimes)/60.0 + float64(latitudeSeconds)/3600.0, float64(longitudeDegrees) + float64(longitudePrimes)/60.0 + float64(longitudeSeconds)/3600.0}
	el.Rad = [2]float64{DegreesToRadians(float64(latitudeDegrees) + float64(latitudePrimes)/60.0 + float64(latitudeSeconds)/3600.0), DegreesToRadians(float64(longitudeDegrees) + float64(longitudePrimes)/60.0 + float64(longitudeSeconds)/3600.0)}
}

// SetRadians Set latitude and longitude as radians
func (el *Point) SetRadians(longitude, latitude float64) {
	el.Loc = [2]float64{RadiansToDegrees(longitude), RadiansToDegrees(latitude)}
	el.Rad = [2]float64{longitude, latitude}
}

// ToDecimalDegreesString Get latitude and longitude
func (el *Point) ToDecimalDegreesString() string {

	dec := math.Abs(el.Loc[LONGITUDE])
	degLng := math.Floor(dec)
	minLng := math.Floor((dec - degLng) * 60.0)
	secLng := (dec - degLng - (minLng / 60.0)) * 3600.0
	if el.Loc[LONGITUDE] < 0 {
		degLng *= -1
	}

	dec = math.Abs(el.Loc[LATITUDE])
	degLat := math.Floor(dec)
	minLat := math.Floor((dec - degLat) * 60.0)
	secLat := (dec - degLat - (minLat / 60.0)) * 3600.0

	if el.Loc[LATITUDE] < 0 {
		degLat *= -1
	}

	return fmt.Sprintf("(%v%v%v%v%2.2f%v,%v%v%v%v%2.2f%v)", degLat, DEGREES, minLat, MINUTES, secLat, SECONDS, degLng, DEGREES, minLng, MINUTES, secLng, SECONDS)
}

func (el *Point) GetLatitudeAsDegrees() float64 {
	return el.Loc[LATITUDE]
}

func (el *Point) GetLongitudeAsDegrees() float64 {
	return el.Loc[LONGITUDE]
}

func (el *Point) GetLatitudeAsRadians() float64 {
	return el.Rad[LATITUDE]
}

func (el *Point) GetLongitudeAsRadians() float64 {
	return el.Rad[LONGITUDE]
}

func (el *Point) ToGoogleMapString() string {
	return fmt.Sprintf("%1.5f, %1.5f [ Please, copy and past this value on google maps search ]", el.Loc[LATITUDE], el.Loc[LONGITUDE])
}

func (el *Point) ToLeafletMapString() string {
	return fmt.Sprintf("[%1.5f, %1.5f],", el.Loc[LATITUDE], el.Loc[LONGITUDE])
}

//func (el Point) GetBoundingBox(distanceAStt Distance) Box {
//	return BoundingBox(el, distanceAStt)
//}

func (el *Point) GetDestinationPoint(distance Distance, angle Angle) Point {
	return DestinationPoint(*el, distance, angle)
}

//func (el *Point) GetDirectionBetweenTwoPoints(point Point) Angle {
//	return DirectionBetweenTwoPoints(el, point)
//}

//func (el *Point) GetDistanceBetweenTwoPoints(point Point) Distance {
//	return DistanceBetweenTwoPoints(el, point)
//}

func (el *Point) add(point Point) Point {
	var ret = Point{}
	ret.SetDegrees(el.Loc[LONGITUDE]+point.Loc[LONGITUDE], el.Loc[LATITUDE]+point.Loc[LATITUDE])
	return ret
}

func (el *Point) sub(point Point) Point {
	var ret = Point{}
	ret.SetDegrees(el.Loc[LONGITUDE]-point.Loc[LONGITUDE], el.Loc[LATITUDE]-point.Loc[LATITUDE])
	return ret
}

func (el *Point) plus(value float64) Point {
	var ret = Point{}
	ret.SetDegrees(el.Loc[LONGITUDE]*value, el.Loc[LATITUDE]*value)
	return ret
}

func (el *Point) div(value float64) Point {
	var ret = Point{}
	ret.SetDegrees(el.Loc[LONGITUDE]/value, el.Loc[LATITUDE]/value)
	return ret
}

func (el *Point) equality(point Point) bool {
	return el.Loc[LONGITUDE] == point.Loc[LONGITUDE] && el.Loc[LATITUDE] == point.Loc[LATITUDE]
}

func (el *Point) dotProduct(point Point) float64 {
	return el.Loc[LONGITUDE]*point.Loc[LONGITUDE] + el.Loc[LATITUDE]*point.Loc[LATITUDE]
}

func (el *Point) distanceSquared(point Point) float64 {
	return (point.Loc[LONGITUDE]-el.Loc[LONGITUDE])*(point.Loc[LONGITUDE]-el.Loc[LONGITUDE]) + (point.Loc[LATITUDE]-el.Loc[LATITUDE])*(point.Loc[LATITUDE]-el.Loc[LATITUDE])
}

func (el *Point) pythagoras(point Point) float64 {
	return math.Sqrt(el.distanceSquared(point))
}

func (el *Point) Distance(pointA, pointB Point) float64 {
	var l2 = pointA.distanceSquared(pointB)
	if l2 == 0.0 {
		return el.pythagoras(pointA) // v == w case
	}

	// Consider the line extending the segment, parameterized as v + t (w - v)
	// We find projection of point p onto the line.
	// It falls where t = [(p-v) . (w-v)] / |w-v|^2
	var pA = el.sub(pointA)
	var pB = pointB.sub(pointA)
	var t = pA.dotProduct(pB) / l2
	if t < 0.0 {
		return el.pythagoras(pointA)
	} else if t > 1.0 {
		return el.pythagoras(pointB)
	}
	var pC = pointB.sub(pointA)
	pC = pC.plus(t)
	pC = pointA.add(pC)

	return el.pythagoras(pC)
}

func (el *Point) decisionDistance(pointsA []Point) float64 {
	var i int
	var curDistance float64
	var dst = el.pythagoras(pointsA[LONGITUDE])
	for i = 1; i < len(pointsA); i += 1 {
		curDistance = el.pythagoras(pointsA[i])
		if curDistance < dst {
			dst = curDistance
		}
	}

	return dst
}

func (el *Point) IsContainedInTheList(pointsAAStt []Point) bool {
	for _, point := range pointsAAStt {
		if el.equality(point) {
			return true
		}
	}

	return false
}
