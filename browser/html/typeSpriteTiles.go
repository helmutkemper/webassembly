package html

import (
	"errors"
	"fmt"
	keyboard "github.com/helmutkemper/iotmaker.webassembly/browser/event/keyBoard"
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

type Event struct {
	Element        CollisionBox
	Event          func()
	EventUp        func()
	EventUpLeft    func()
	EventUpRight   func()
	EventDown      func()
	EventDownLeft  func()
	EventDownRight func()
	TimeOut        time.Duration
}

type SpriteKeyboard struct {
	Key      string
	Up       func()
	UpExec   bool
	Down     func()
	DownExec bool
}

type SpriteData struct {
	Csv       [][]int
	Draw      map[int]js.Value
	Collision map[int][][]bool
	Box       [][]Box
	Events    map[int]Event
	Keyboard  map[int]SpriteKeyboard
	Triggered map[int]time.Time
	Width     int
	Height    int
	Visible   bool
	Process   bool
}

func (e *SpriteData) Init(width, height int) {
	e.Draw = make(map[int]js.Value)
	e.Collision = make(map[int][][]bool)
	e.Width = width
	e.Height = height
}

type SpriteTiles struct {
	name              string
	frame             map[string]SpriteData
	error             error
	NoEvents          map[int]Event
	Trigger           map[int]time.Time
	uId               int
	KeyCode           map[string]bool
	lastCollisionTime time.Time

	eventOnKeyData chan keyboard.Data

	lx      int
	ly      int
	lw      int
	lh      int
	lcanvas *TagCanvas
}

func (e *SpriteTiles) getUId() int {
	defer func() {
		e.uId += 1
	}()

	return e.uId
}

func (e *SpriteTiles) Init(stage *stage.Stage) {
	e.frame = make(map[string]SpriteData)
	e.NoEvents = make(map[int]Event)
	e.Trigger = make(map[int]time.Time)
	e.KeyCode = make(map[string]bool)

	e.lastCollisionTime = time.Now()

	e.eventOnKeyData = make(chan keyboard.Data)

	go func() {

		for {
			select {

			case data := <-e.eventOnKeyData:
				if data.Repeat == true {
					continue
				}
				switch data.EventName {
				case "keyup":
					e.KeyCode[data.Code] = false
				case "keydown":
					e.KeyCode[data.Code] = true
				}
			}

		}
	}()

	stage.AddListenerKeyUp(&e.eventOnKeyData)
	stage.AddListenerKeyDown(&e.eventOnKeyData)
}

func (e *SpriteTiles) GetError() (err error) {
	return e.error
}

func (e *SpriteTiles) AddNoEvent(element CollisionBox, timeOut time.Duration, eventList ...func()) (ref *SpriteTiles) {
	for _, event := range eventList {
		e.NoEvents[e.getUId()] = Event{Element: element, Event: event, TimeOut: timeOut}
	}

	return e
}

func (e *SpriteTiles) AddCollision(element CollisionBox, event func()) (ref *SpriteTiles) {
	data := e.frame[e.name]
	if data.Events == nil {
		data.Events = make(map[int]Event)
		data.Triggered = make(map[int]time.Time)
	}

	data.Events[e.getUId()] = Event{Element: element, Event: event}
	e.frame[e.name] = data
	return e
}

func (e *SpriteTiles) AddCollisionUp(element CollisionBox, event func()) (ref *SpriteTiles) {
	data := e.frame[e.name]
	if data.Events == nil {
		data.Events = make(map[int]Event)
		data.Triggered = make(map[int]time.Time)
	}

	data.Events[e.getUId()] = Event{Element: element, EventUp: event}
	e.frame[e.name] = data
	return e
}

func (e *SpriteTiles) AddCollisionUpLeft(element CollisionBox, event func()) (ref *SpriteTiles) {
	data := e.frame[e.name]
	if data.Events == nil {
		data.Events = make(map[int]Event)
		data.Triggered = make(map[int]time.Time)
	}

	data.Events[e.getUId()] = Event{Element: element, EventUpLeft: event}
	e.frame[e.name] = data
	return e
}

func (e *SpriteTiles) AddCollisionUpRight(element CollisionBox, event func()) (ref *SpriteTiles) {
	data := e.frame[e.name]
	if data.Events == nil {
		data.Events = make(map[int]Event)
		data.Triggered = make(map[int]time.Time)
	}

	data.Events[e.getUId()] = Event{Element: element, EventUpRight: event}
	e.frame[e.name] = data
	return e
}

func (e *SpriteTiles) AddCollisionDown(element CollisionBox, event func()) (ref *SpriteTiles) {
	data := e.frame[e.name]
	if data.Events == nil {
		data.Events = make(map[int]Event)
		data.Triggered = make(map[int]time.Time)
	}

	data.Events[e.getUId()] = Event{Element: element, EventDown: event}
	e.frame[e.name] = data
	return e
}

func (e *SpriteTiles) AddCollisionDownLeft(element CollisionBox, event func()) (ref *SpriteTiles) {
	data := e.frame[e.name]
	if data.Events == nil {
		data.Events = make(map[int]Event)
		data.Triggered = make(map[int]time.Time)
	}

	data.Events[e.getUId()] = Event{Element: element, EventDownLeft: event}
	e.frame[e.name] = data
	return e
}

func (e *SpriteTiles) AddCollisionDownRight(element CollisionBox, event func()) (ref *SpriteTiles) {
	data := e.frame[e.name]
	if data.Events == nil {
		data.Events = make(map[int]Event)
		data.Triggered = make(map[int]time.Time)
	}

	data.Events[e.getUId()] = Event{Element: element, EventDownRight: event}
	e.frame[e.name] = data
	return e
}

func (e *SpriteTiles) AddKeyboard(key string, down, up func()) (ref *SpriteTiles) {
	data := e.frame[e.name]
	if data.Keyboard == nil {
		data.Keyboard = make(map[int]SpriteKeyboard)
	}

	data.Keyboard[e.getUId()] = SpriteKeyboard{Key: key, Up: up, Down: down}
	e.frame[e.name] = data
	return e
}

func (e *SpriteTiles) AddCsv(sceneName, csv string, img *TagImg, visible, process bool, width, height int) (ref *SpriteTiles) {
	var err error

	e.name = sceneName
	if e.frame == nil {
		e.frame = make(map[string]SpriteData)
	}

	data := SpriteData{}
	data.Init(width, height)
	data.Visible = visible
	data.Process = process

	var tile int64
	var imageData js.Value

	lines := strings.Split(csv, "\n")
	data.Csv = make([][]int, len(lines))
	data.Box = make([][]Box, len(lines))
	for row, line := range lines {

		// English: occurs when there is a line break at the end of the file.
		// Português: ocorre quando tem uma quebra de linha no fim do arquivo.
		if line == "" {
			continue
		}

		cols := strings.Split(line, ",")
		data.Csv[row] = make([]int, len(cols))
		data.Box[row] = make([]Box, len(cols))
		for col, value := range cols {

			// English: occurs when there is a line break at the end of the file.
			// Português: ocorre quando tem uma quebra de linha no fim do arquivo.
			if value == "" {
				continue
			}

			tile, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				e.error = err
				return e
			}

			colsLen := img.GetWidth() / width
			colTile := int(tile) % colsLen
			rowTile := int(tile) / colsLen

			data.Csv[row][col] = int(tile)

			if tile == -1 {
				continue
			}

			canvas := new(TagCanvas)
			canvas.Init(width, height)

			canvas.ClearRect(0, 0, width, height)
			canvas.DrawImageComplete(img, colTile*width, rowTile*height, width, height, 0, 0, width, height)
			imageData = canvas.GetImageData(0, 0, width, height, false, false)

			//todo: descomentar - início
			//cBox := canvas.GetCollisionBoxFromImageData(imageData, width, height)
			//cBox.XImg = col * width
			//cBox.YImg = row * height
			//cBox.WidthImg = width
			//cBox.HeightImg = height
			//data.Box[row][col] = cBox
			//todo: descomentar - fim

			_, found := data.Draw[int(tile)]
			if found == true {
				continue
			}

			data.Draw[int(tile)] = imageData
			data.Collision[int(tile)] = canvas.GetCollisionDataFromImageData(imageData, width, height)
		}
	}

	e.frame[sceneName] = data
	return e
}

func (e *SpriteTiles) Draw(canvas *TagCanvas, sceneName string) (err error) {
	data, ok := e.frame[sceneName]

	e.lcanvas = canvas

	if ok == false {
		err = errors.New(fmt.Sprintf("scene %v not found", sceneName))
		return
	}

	if data.Visible == false {
		return
	}

	for row, rowData := range data.Csv {
		for col := range rowData {
			spriteId := data.Csv[row][col]
			if spriteId == -1 {
				continue
			}

			canvas.context.Call("putImageData", data.Draw[spriteId], col*data.Width, row*data.Height)
		}
	}

	return
}

//func (e *SpriteTiles) Process() {
//
//	var sceneElementCBox Box
//	var pass = false
//	for frameName, frameData := range e.frame {
//		if frameData.Process == false {
//			continue
//		}
//
//		rowMaxAllowed := len(frameData.Box) - 1
//		colMaxAllowed := len(frameData.Box[0]) - 1
//		for frameEventID, frameDataEvent := range frameData.Events {
//			var elementCBox = frameDataEvent.Element.GetCollisionBox()
//
//			rowMin := elementCBox.YImg / frameData.Height
//			if rowMin < 0 {
//				rowMin = 0
//			}
//
//			rowMax := (elementCBox.YImg + elementCBox.HeightImg) / frameData.Height
//			if rowMax > rowMaxAllowed {
//				rowMax = rowMaxAllowed
//			}
//
//			colMin := elementCBox.XImg / frameData.Width
//			if colMin < 0 {
//				colMin = 0
//			}
//
//			colMax := (elementCBox.XImg + elementCBox.WidthImg) / frameData.Width
//			if colMax > colMaxAllowed {
//				colMax = colMaxAllowed
//			}
//
//			for row := rowMin; row <= rowMax; row += 1 {
//				for col := colMin; col <= colMax; col += 1 {
//					sceneElementCBox = frameData.Box[row][col]
//
//					e.lx = elementCBox.X + elementCBox.XImg
//					e.ly = elementCBox.Y + elementCBox.YImg
//					e.lw = elementCBox.Width
//					e.lh = elementCBox.Height
//
//					e.lcanvas.StrokeStyle(factoryColor.NewRed()).
//						ClearRect(e.lx-5, e.ly-5, e.lw+10, e.lh+10).
//						StrokeRect(e.lx, e.ly, e.lw, e.lh).
//						Stroke()
//
//					if elementCBox.Collision(sceneElementCBox) {
//
//						pass = true
//						//collision
//
//						e.lastCollisionTime = time.Now()
//						frameData.Triggered[frameEventID] = time.Now()
//
//						upLeft, upRight, downLeft, downRight := elementCBox.Quadrant(sceneElementCBox)
//						log.Printf("quadrant: %v, %v, %v, %v", upLeft, upRight, downLeft, downRight)
//						// todo: isto deve fazer o ajuste de penetração da colisão
//						//diffY := elementCBox.Y + elementCBox.Height - sceneElementCBox.Y
//						//if diffY > 5 {
//						//	elementCBox.Y -= diffY
//						//}
//
//						if (upLeft || upRight) && frameDataEvent.EventUp != nil {
//							frameDataEvent.EventUp()
//							log.Printf("up: %v", (elementCBox.YImg+elementCBox.Y)-(sceneElementCBox.YImg+sceneElementCBox.Y+sceneElementCBox.Height))
//							//frameDataEvent.Element.AdjustBox(0, (elementCBox.YImg+elementCBox.Y)-(sceneElementCBox.YImg+sceneElementCBox.Y+sceneElementCBox.Height))
//							break
//						}
//
//						if upLeft && frameDataEvent.EventUpLeft != nil {
//							frameDataEvent.EventUpLeft()
//							break
//						}
//
//						if upRight && frameDataEvent.EventUpRight != nil {
//							frameDataEvent.EventUpRight()
//							break
//						}
//
//						if (downLeft || downRight) && frameDataEvent.EventDown != nil {
//							frameDataEvent.EventDown()
//							log.Printf("down: %v", (sceneElementCBox.YImg+sceneElementCBox.Y)-(elementCBox.YImg+elementCBox.Y+elementCBox.Height))
//							frameDataEvent.Element.AdjustBox(0, (sceneElementCBox.YImg+sceneElementCBox.Y)-(elementCBox.YImg+elementCBox.Y+elementCBox.Height)-2)
//							break
//						}
//
//						if downLeft && frameDataEvent.EventDownLeft != nil {
//							frameDataEvent.EventDownLeft()
//							break
//						}
//
//						if downRight && frameDataEvent.EventDownRight != nil {
//							frameDataEvent.EventDownRight()
//							break
//						}
//
//						if frameDataEvent.Event != nil {
//							frameDataEvent.Event()
//							break
//						}
//					}
//
//					if pass == true {
//						break
//					}
//				}
//			}
//
//			if pass == true { //fixme: eu posso querer usar as teclas em queda livre - colocar timeout como em jumpEnable
//				for k := range frameData.Keyboard {
//					keyOn, keyFound := e.KeyCode[frameData.Keyboard[k].Key]
//					if keyFound == true && keyOn == true && frameData.Keyboard[k].DownExec == false {
//						tmp := frameData.Keyboard[k]
//						tmp.DownExec = true
//						tmp.UpExec = false
//						frameData.Keyboard[k] = tmp
//
//						if frameData.Keyboard[k].Down != nil {
//							frameData.Keyboard[k].Down()
//						}
//					} else if keyFound == true && keyOn == false && frameData.Keyboard[k].UpExec == false {
//						tmp := frameData.Keyboard[k]
//						tmp.UpExec = true
//						tmp.DownExec = false
//						frameData.Keyboard[k] = tmp
//
//						if frameData.Keyboard[k].Up != nil {
//							frameData.Keyboard[k].Up()
//						}
//					}
//				}
//			}
//
//			// collision
//			if pass == false {
//
//				for noEventsEventID, noEventsEvent := range e.NoEvents {
//					if noEventsEvent.TimeOut != 0 && time.Now().Sub(e.lastCollisionTime) < noEventsEvent.TimeOut {
//						continue
//					}
//
//					if noEventsEvent.Event != nil {
//						noEventsEvent.Event()
//						e.Trigger[noEventsEventID] = time.Now()
//					}
//				}
//			}
//		}
//
//		e.frame[frameName] = frameData
//	}
//}
