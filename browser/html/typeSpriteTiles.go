package html

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

type SpriteData struct {
	Csv      [][]int
	Draw     map[int]js.Value
	Colision map[int][][]bool
	Box      [][]Box
	Width    int
	Height   int
}

func (e *SpriteData) Init(width, height int) {
	e.Draw = make(map[int]js.Value)
	e.Colision = make(map[int][][]bool)
	e.Width = width
	e.Height = height
}

type SpriteTiles struct {
	frame map[string]SpriteData
}

func (e *SpriteTiles) AddCsv(sceneName, csv string, img *TagImg, width, height int) (err error) {
	if e.frame == nil {
		e.frame = make(map[string]SpriteData)
	}

	data := SpriteData{}
	data.Init(width, height)

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
				return
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

			canvas.context.Call("clearRect", 0, 0, width, height)
			canvas.context.Call("drawImage", img.Get(), colTile*width, rowTile*height, width, height, 0, 0, width, height)
			imageData = canvas.context.Call("getImageData", 0, 0, width, height)

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
	return
}

func (e *SpriteTiles) Draw(canvas *TagCanvas, sceneName string) (err error) {
	data, ok := e.frame[sceneName]

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

func (e *SpriteTiles) TestCollisionBox(element CollisionBox, sceneName string) (collision bool) {
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
