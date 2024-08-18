package components

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"log"
	"math"
	"syscall/js"
)

// __osmOnInputEvent faz a captura de dados do event input
type __osmOnInputEvent struct {
	ClientX int64 `wasmGetArg:"clientX"`
	ClientY int64 `wasmGetArg:"clientY"`
	StartX  int64
	StartY  int64
}

func (e *__osmOnInputEvent) OnMousedown(event html.TagCanvas, ref any) {
	log.Printf("event type: %T", event)
	log.Printf("> OnMousedown.clientX: %v", event.MouseClientX)
	log.Printf("> OnMousedown.clientY: %v", event.MouseClientY)

}

type Osm struct {
	__longitude any
	__latitude  any
	__zoom      any
	__change    *__osmOnInputEvent
	__width     int
	__height    int
	__url       string

	__osmTag    *html.TagDiv    `wasmPanel:"type:TagDiv"`
	__canvasTag *html.TagCanvas `wasmPanel:"type:TagCanvas"`

	__startX   int
	__startY   int
	__startLon float64
	__startLat float64
}

func (e *Osm) onMouseDown(this js.Value, args []js.Value) interface{} {
	e.__longitude = e.__canvasTag.GetOsmLongitude()
	e.__latitude = e.__canvasTag.GetOsmLatitude()
	e.__zoom = e.__canvasTag.GetOsmZoom()

	event := args[0]
	e.__startX = event.Get("clientX").Int()
	e.__startY = event.Get("clientY").Int()
	e.__startLon = e.__longitude.(float64)
	e.__startLat = e.__latitude.(float64)

	e.__canvasTag.Get().Call("addEventListener", "mousemove", js.FuncOf(e.onMouseMove))
	e.__canvasTag.Get().Call("addEventListener", "mouseup", js.FuncOf(e.onMouseUp))
	e.__canvasTag.Get().Call("addEventListener", "mouseout", js.FuncOf(e.onMouseUp))

	return nil
}

func (e *Osm) onMouseMove(this js.Value, args []js.Value) interface{} {
	event := args[0]
	dx := event.Get("clientX").Int() - e.__startX
	dy := event.Get("clientY").Int() - e.__startY

	e.__canvasTag.ShiftInPixels(dx, dy, e.__zoom.(int), e.__longitude.(float64), e.__latitude.(float64))
	return nil

}

func (e *Osm) onMouseUp(this js.Value, args []js.Value) interface{} {
	e.__canvasTag.Get().Call("removeEventListener", "mousemove", js.FuncOf(e.onMouseMove))
	e.__canvasTag.Get().Call("removeEventListener", "mouseup", js.FuncOf(e.onMouseUp))
	e.__canvasTag.Get().Call("removeEventListener", "mouseout", js.FuncOf(e.onMouseUp))

	return nil
}

func (e *Osm) pixelsToLonLat(pixelX, pixelY float64, zoom int) (float64, float64) {
	scale := 1 << uint(zoom)
	lon := (pixelX/256.0)/float64(scale)*360.0 - 180.0
	n := math.Pi - 2.0*math.Pi*(pixelY/256.0)/float64(scale)
	lat := 180.0 / math.Pi * math.Atan(math.Sinh(n))
	return lon, lat
}

func (e *Osm) SetValue(longitude, latitude float64, zoom int) {
	e.__longitude = longitude
	e.__latitude = latitude
	e.__zoom = zoom

	if e.__canvasTag == nil {
		return
	}

	e.setValue(longitude, latitude, zoom)
}

func (e *Osm) SetOsmURL(url string) {
	e.__url = url

	if e.__canvasTag == nil {
		return
	}
}

func (e *Osm) setOsmURL(url string) {
	e.__canvasTag.SetOsmURL(url)
}

func (e *Osm) setValue(longitude, latitude float64, zoom int) {
	e.__canvasTag.SetOsm(longitude, latitude, zoom, 0, 0)
}

func (e *Osm) init() {
	if e.__url != "" {
		e.SetOsmURL(e.__url)
	}

	if e.__longitude != nil && e.__latitude != nil && e.__zoom != nil {
		e.setValue(e.__longitude.(float64), e.__latitude.(float64), e.__zoom.(int))
	}
}
