package main

import (
	"fmt"
	"math"
	"syscall/js"
)

var startX, startY int
var startLon, startLat float64
var zoom int
var tileUrlTemplate string
var lon, lat float64

func lonLatToTile(lon, lat float64, zoom int) (int, int) {
	scale := 1 << zoom
	x := int((lon + 180.0) / 360.0 * float64(scale))
	y := int((1.0 - (math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0)) / math.Pi)) / 2.0 * float64(scale))
	return x, y
}

func lonLatToPixels(lon, lat float64, zoom int) (float64, float64) {
	scale := 1 << zoom
	x := (lon + 180.0) / 360.0 * float64(scale) * 256
	y := (1.0 - (math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0)) / math.Pi)) / 2.0 * float64(scale) * 256
	return x, y
}

func pixelsToLonLat(pixelX, pixelY float64, zoom int) (float64, float64) {
	scale := 1 << zoom
	lon := (pixelX/256.0/float64(scale))*360.0 - 180.0
	n := math.Pi - 2.0*math.Pi*(pixelY/256.0)/float64(scale)
	lat := 180.0 / math.Pi * math.Atan(math.Sinh(n))
	return lon, lat
}

func drawTile(ctx js.Value, tileX, tileY, zoom int, offsetX, offsetY int, tileUrlTemplate string) {
	img := js.Global().Get("Image").New()
	url := fmt.Sprintf(tileUrlTemplate, zoom, tileX, tileY)
	img.Set("src", url)
	img.Call("addEventListener", "load", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ctx.Call("drawImage", img, offsetX, offsetY)
		return nil
	}))
}

func drawMap(this js.Value, args []js.Value) interface{} {
	canvas := js.Global().Get("document").Call("getElementById", "mapCanvas")
	ctx := canvas.Call("getContext", "2d")

	lon = -0.1276 // Longitude de Londres
	lat = 51.5074 // Latitude de Londres
	zoom = 17
	tileUrlTemplate = "https://tile.openstreetmap.org/%d/%d/%d.png" // URL template para os tiles

	tileSize := 256
	centerX := canvas.Get("width").Int()/2 - tileSize/2
	centerY := canvas.Get("height").Int()/2 - tileSize/2

	drawTiles(ctx, lon, lat, zoom, centerX, centerY, tileUrlTemplate)

	canvas.Call("addEventListener", "mousedown", js.FuncOf(onMouseDown))

	return nil
}

func onMouseDown(this js.Value, args []js.Value) interface{} {
	event := args[0]
	startX = event.Get("clientX").Int()
	startY = event.Get("clientY").Int()
	startLon, startLat = currentLonLat()

	canvas := js.Global().Get("document").Call("getElementById", "mapCanvas")
	canvas.Call("addEventListener", "mousemove", js.FuncOf(onMouseMove))
	canvas.Call("addEventListener", "mouseup", js.FuncOf(onMouseUp))
	canvas.Call("addEventListener", "mouseout", js.FuncOf(onMouseUp))

	return nil
}

func onMouseMove(this js.Value, args []js.Value) interface{} {
	event := args[0]
	dx := event.Get("clientX").Int() - startX
	//dy := event.Get("clientY").Int() - startY
	dy := 0

	dlat, dlon := calculateDelta(lat, lon, zoom, dx, dy)
	lon = dlon + lon
	lat = dlat + lat

	canvas := js.Global().Get("document").Call("getElementById", "mapCanvas")
	ctx := canvas.Call("getContext", "2d")
	ctx.Call("clearRect", 0, 0, canvas.Get("width").Int(), canvas.Get("height").Int())

	centerX := canvas.Get("width").Int()/2 - 256/2
	centerY := canvas.Get("height").Int()/2 - 256/2

	drawTiles(ctx, lon, lat, zoom, centerX, centerY, tileUrlTemplate)

	//scale := 1 << zoom
	//lonPerPixel := 360.0 / float64(scale*256)
	//latPerPixel := 180.0 / math.Pi * math.Atan(math.Sinh(math.Pi*(2.0*(256.0*float64(scale)-float64(dy))/256.0/float64(scale)-1.0)))
	//
	//newLon := startLon - float64(dx)*lonPerPixel
	//newLat := startLat - float64(dy)*latPerPixel
	//
	//canvas := js.Global().Get("document").Call("getElementById", "mapCanvas")
	//ctx := canvas.Call("getContext", "2d")
	////ctx.Call("clearRect", 0, 0, canvas.Get("width").Int(), canvas.Get("height").Int())
	//
	//centerX := canvas.Get("width").Int()/2 - 256/2
	//centerY := canvas.Get("height").Int()/2 - 256/2
	//
	//drawTiles(ctx, newLon, newLat, zoom, centerX, centerY, tileUrlTemplate)

	return nil
}

func onMouseUp(this js.Value, args []js.Value) interface{} {
	canvas := js.Global().Get("document").Call("getElementById", "mapCanvas")
	canvas.Call("removeEventListener", "mousemove", js.FuncOf(onMouseMove))
	canvas.Call("removeEventListener", "mouseup", js.FuncOf(onMouseUp))
	canvas.Call("removeEventListener", "mouseout", js.FuncOf(onMouseUp))

	return nil
}

func drawTiles(ctx js.Value, lon, lat float64, zoom, centerX, centerY int, tileUrlTemplate string) {
	tileX, tileY := lonLatToTile(lon, lat, zoom)

	drawTile(ctx, tileX, tileY, zoom, centerX, centerY, tileUrlTemplate)
	drawTile(ctx, tileX+1, tileY, zoom, centerX+256, centerY, tileUrlTemplate)
	drawTile(ctx, tileX-1, tileY, zoom, centerX-256, centerY, tileUrlTemplate)
	drawTile(ctx, tileX, tileY+1, zoom, centerX, centerY+256, tileUrlTemplate)
	drawTile(ctx, tileX, tileY-1, zoom, centerX, centerY-256, tileUrlTemplate)
	drawTile(ctx, tileX+1, tileY+1, zoom, centerX+256, centerY+256, tileUrlTemplate)
	drawTile(ctx, tileX-1, tileY-1, zoom, centerX-256, centerY-256, tileUrlTemplate)
	drawTile(ctx, tileX+1, tileY-1, zoom, centerX+256, centerY-256, tileUrlTemplate)
	drawTile(ctx, tileX-1, tileY+1, zoom, centerX-256, centerY+256, tileUrlTemplate)
}

func currentLonLat() (float64, float64) {
	canvas := js.Global().Get("document").Call("getElementById", "mapCanvas")
	width := canvas.Get("width").Int()
	height := canvas.Get("height").Int()
	centerPixelX := float64(width / 2)
	centerPixelY := float64(height / 2)

	return pixelsToLonLat(centerPixelX, centerPixelY, zoom)
}

func calculateDelta(lat, lon float64, zoom int, deltaX, deltaY int) (float64, float64) {
	// Convert latitude to radians
	latRad := lat * math.Pi / 180

	// Calculate the size of one pixel in degrees
	longitudePerPixel := 360 / (math.Pow(2, float64(zoom)+8))
	latitudePerPixel := longitudePerPixel * math.Cos(latRad)

	// Calculate the change in longitude and latitude
	deltaLon := float64(deltaX) * longitudePerPixel
	deltaLat := float64(deltaY) * latitudePerPixel

	// Return the new longitude and latitude
	return deltaLat, deltaLon
}

func main() {
	zoom = 17
	js.Global().Set("drawMap", js.FuncOf(drawMap))
	select {}
}
