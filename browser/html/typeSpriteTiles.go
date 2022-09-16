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
	Element CollisionBox
	Event   func()
}

type SpriteKeyboard struct {
	Key      string
	Up       func()
	UpExec   bool
	Down     func()
	DownExec bool
}

type SpriteData struct {
	Csv      [][]int
	Draw     map[int]js.Value
	Colision map[int][][]bool
	Box      [][]Box
	Events   map[int]Event
	Keyboard map[int]SpriteKeyboard
	Trigged  map[int]time.Time
	Width    int
	Height   int
	Visible  bool
	Proccess bool
}

func (e *SpriteData) Init(width, height int) {
	e.Draw = make(map[int]js.Value)
	e.Colision = make(map[int][][]bool)
	e.Width = width
	e.Height = height
}

type SpriteTiles struct {
	name     string
	frame    map[string]SpriteData
	error    error
	NoEvents map[int]Event
	Trigged  map[int]time.Time
	uId      int

	KeyCode map[string]bool

	eventOnKeyData chan keyboard.Data
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
	e.Trigged = make(map[int]time.Time)
	e.KeyCode = make(map[string]bool)

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

func (e *SpriteTiles) AddNoEvent(element CollisionBox, eventList ...func()) (ref *SpriteTiles) {
	for _, event := range eventList {
		e.NoEvents[e.getUId()] = Event{Element: element, Event: event}
	}

	return e
}

func (e *SpriteTiles) AddColision(element CollisionBox, event func()) (ref *SpriteTiles) {
	data := e.frame[e.name]
	if data.Events == nil {
		data.Events = make(map[int]Event)
		data.Trigged = make(map[int]time.Time)
	}

	data.Events[e.getUId()] = Event{Element: element, Event: event}
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

func (e *SpriteTiles) AddCsv(sceneName, csv string, img *TagImg, visible, proccess bool, width, height int) (ref *SpriteTiles) {
	var err error

	e.name = sceneName
	if e.frame == nil {
		e.frame = make(map[string]SpriteData)
	}

	data := SpriteData{}
	data.Init(width, height)
	data.Visible = visible
	data.Proccess = proccess

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

			cBox := canvas.GetCollisionBox()
			cBox.X += col * width
			cBox.Y += row * height
			data.Box[row][col] = cBox

			_, found := data.Draw[int(tile)]
			if found == true {
				continue
			}

			data.Draw[int(tile)] = imageData
			data.Colision[int(tile)] = canvas.GetCollisionData()
		}
	}

	e.frame[sceneName] = data
	return e
}

func (e *SpriteTiles) Draw(canvas *TagCanvas, sceneName string) (err error) {
	data, ok := e.frame[sceneName]

	if data.Visible == false {
		return
	}

	if ok == false {
		err = errors.New(fmt.Sprintf("scene %v not found", sceneName))
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

func (e *SpriteTiles) Proccess() {

	var thisCBox Box
	var pass = false
	for name, data := range e.frame {
		_ = name //fixme
		if data.Proccess == false {
			continue
		}

		rowMaxAllowed := len(data.Box) - 1
		colMaxAllowed := len(data.Box[0]) - 1
		for eventID, v := range data.Events {
			var elementCBox = v.Element.GetCollisionBox()

			rowMin := elementCBox.Y / data.Height
			if rowMin < 0 {
				rowMin = 0
			}

			rowMax := (elementCBox.Y + elementCBox.Height) / data.Height
			if rowMax > rowMaxAllowed {
				rowMax = rowMaxAllowed
			}

			colMin := elementCBox.X / data.Width
			if colMin < 0 {
				colMin = 0
			}

			colMax := (elementCBox.X + elementCBox.Width) / data.Width
			if colMax > colMaxAllowed {
				colMax = colMaxAllowed
			}

			for row := rowMin; row <= rowMax; row += 1 {
				for col := colMin; col <= colMax; col += 1 {
					thisCBox = data.Box[row][col]
					if thisCBox.X < elementCBox.X+elementCBox.Width &&
						thisCBox.X+thisCBox.Width > elementCBox.X &&
						thisCBox.Y < elementCBox.Y+elementCBox.Height &&
						thisCBox.Y+thisCBox.Height > elementCBox.Y {

						pass = true
						//collision

						if v.Event != nil {
							v.Event()
							data.Trigged[eventID] = time.Now()
						}

						break
					}
				}
			}

			if pass == true {
				for k := range data.Keyboard {
					keyOn, keyFound := e.KeyCode[data.Keyboard[k].Key]
					if keyFound == true && keyOn == true && data.Keyboard[k].DownExec == false {
						tmp := data.Keyboard[k]
						tmp.DownExec = true
						tmp.UpExec = false
						data.Keyboard[k] = tmp

						if data.Keyboard[k].Down != nil {
							data.Keyboard[k].Down()
						}
					} else if keyFound == true && keyOn == false && data.Keyboard[k].UpExec == false {
						tmp := data.Keyboard[k]
						tmp.UpExec = true
						tmp.DownExec = false
						data.Keyboard[k] = tmp

						if data.Keyboard[k].Up != nil {
							data.Keyboard[k].Up()
						}
					}
				}
			}

			if pass == false {
				for eventID, v := range e.NoEvents {
					if v.Event != nil {
						v.Event()
						e.Trigged[eventID] = time.Now()
					}
				}
			}
		}

		e.frame[name] = data
	}
}

func (e *SpriteTiles) TestCollisionBox(element CollisionBox, sceneName string) (collision bool) {
	if e.frame == nil {
		return
	}

	data, ok := e.frame[sceneName]
	if ok == false {
		return
	}

	elementCBox := element.GetCollisionBox()

	var thisCBox Box
	for row, rowData := range data.Csv {
		for col := range rowData {
			thisCBox = data.Box[row][col]
			if thisCBox.X < elementCBox.X+elementCBox.Width &&
				thisCBox.X+thisCBox.Width > elementCBox.X &&
				thisCBox.Y < elementCBox.Y+elementCBox.Height &&
				thisCBox.Y+thisCBox.Height > elementCBox.Y {
				return true
			}
		}
	}

	return false
}
