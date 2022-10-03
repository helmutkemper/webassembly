package html

import (
	"image/color"
	"log"
	"math"
	"syscall/js"
)

type Box struct {
	xImg      int
	yImg      int
	widthImg  int
	heightImg int
	x         int
	y         int
	width     int
	height    int

	collisionDataFunction func(red, green, blue, alpha uint8) bool

	boxImageColor     *color.RGBA
	boxCollisionColor *color.RGBA

	data js.Value
}

func (e *Box) X(x int) {
	e.xImg = x
}

func (e *Box) Y(y int) {
	e.yImg = y
}

func (e *Box) GetX() (x int) {
	return e.xImg
}

func (e *Box) GetY() (y int) {
	return e.yImg
}

func (e *Box) GetXBox() (x int) {
	return e.xImg + e.x
}

func (e *Box) GetYBox() (y int) {
	return e.yImg + e.y
}

func (e *Box) GetWidthBox() (width int) {
	return e.width
}

func (e *Box) GetHeightBox() (height int) {
	return e.height
}

func (e *Box) GetWidth() (width int) {
	return e.widthImg
}

func (e *Box) GetHeight() (height int) {
	return e.heightImg
}

func (e *Box) GetData() (data js.Value) {
	return e.data
}

func (e *Box) Debug(boxImageColor, boxCollisionColor *color.RGBA) {
	e.boxImageColor = boxImageColor
	e.boxCollisionColor = boxCollisionColor
}

func (e *Box) Populate(data js.Value, width, height int) {
	e.data = data
	e.widthImg = width
	e.heightImg = height

	if e.collisionDataFunction == nil {
		e.collisionDataFunction = func(r, g, b, a uint8) bool {
			if a != 0 {
				return true
			}

			return false
		}
	}

	e.getCollisionBoxFromImageData()

	if e.boxCollisionColor != nil && e.boxImageColor != nil {
		e.debug(*e.boxCollisionColor, *e.boxImageColor)
	}
}

func (e *Box) DataFunction(f func(red, green, blue, alpha uint8) bool) {
	e.collisionDataFunction = f
}

func (e *Box) getCollisionDataFromImageData() (data [][]bool) {
	dataJs := e.data.Get("data")

	var rgbaLength = 4

	var i = 0
	var x int
	var y int

	var red uint8
	var green uint8
	var blue uint8
	var alpha uint8

	data = make([][]bool, e.heightImg)
	for y = 0; y != e.heightImg; y += 1 {
		data[y] = make([]bool, e.widthImg)
		for x = 0; x != e.widthImg; x += 1 {

			red = uint8(dataJs.Index(i + 0).Int())
			green = uint8(dataJs.Index(i + 1).Int())
			blue = uint8(dataJs.Index(i + 2).Int())
			alpha = uint8(dataJs.Index(i + 3).Int())

			data[y][x] = e.collisionDataFunction(red, green, blue, alpha)

			i += rgbaLength
		}
	}

	return
}

func (e *Box) debug(collisionColor, imageColor color.RGBA) {
	var rgbaLength = 4

	var i = 0
	var col int
	var row int

	for row = 0; row != e.heightImg; row += 1 {
		for col = 0; col != e.widthImg; col += 1 {

			//box de colisão - início
			//A := true //col == x || col == x+width
			//B := true //row == y || row == y+height
			//C := col >= x && col <= x+width
			//D := row >= y && row <= y+height
			//if (A && D) && (B && C) {}
			//box de colisão - fim

			AImage := col == 0 || col == e.widthImg-1
			BImage := row == 0 || row == e.heightImg-1
			CImage := col >= 0 && col <= e.widthImg-1
			DImage := row >= 0 && row <= e.heightImg-1

			if (AImage && DImage) || (BImage && CImage) {
				e.data.Get("data").SetIndex(i+0, int(imageColor.R))
				e.data.Get("data").SetIndex(i+1, int(imageColor.G))
				e.data.Get("data").SetIndex(i+2, int(imageColor.B))
				e.data.Get("data").SetIndex(i+3, int(imageColor.A))
			}

			ACollision := col == e.x || col == e.x+e.width
			BCollision := row == e.y || row == e.y+e.height
			CCollision := col >= e.x && col <= e.x+e.width
			DCollision := row >= e.y && row <= e.y+e.height

			if (ACollision && DCollision) || (BCollision && CCollision) {
				e.data.Get("data").SetIndex(i+0, int(collisionColor.R))
				e.data.Get("data").SetIndex(i+1, int(collisionColor.G))
				e.data.Get("data").SetIndex(i+2, int(collisionColor.B))
				e.data.Get("data").SetIndex(i+3, int(collisionColor.A))
			}

			i += rgbaLength
		}
	}
}

func (e *Box) getCollisionBoxFromImageData() {

	var xMin = math.MaxInt
	var yMin = math.MaxInt
	var xMax = math.MinInt
	var yMax = math.MinInt

	collisionData := e.getCollisionDataFromImageData()
	for y, dataY := range collisionData {
		for x, value := range dataY {
			if value == true {
				xMin = e.min(xMin, x)
				yMin = e.min(yMin, y)
				xMax = e.max(xMax, x)
				yMax = e.max(yMax, y)
			}
		}
	}

	e.x = xMin
	e.y = yMin
	e.width = xMax - xMin
	e.height = yMax - yMin
}

func (e *Box) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (e *Box) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (e *Box) collision(element Box) (collision bool) {
	if e.GetXBox() <= element.GetXBox()+element.GetWidthBox() &&
		e.GetXBox()+e.GetWidthBox() >= element.GetXBox() &&
		e.GetYBox() < element.GetYBox()+element.GetHeightBox() &&
		e.GetYBox()+e.GetHeightBox() >= element.GetYBox() {
		return true
	}

	return false
}

// Quadrant
//
// English:
//
// Returns the collision quadrant where the tested element hit the box.
//
// Português:
//
// Retorna o quadrante da colisão onde o elemento testado bateu na caixa.
func (e *Box) Quadrant(element Box) (upLeft, upRight, downLeft, downRight bool) {
	q := Box{}
	q.xImg = e.xImg
	q.yImg = e.yImg
	q.widthImg = e.widthImg
	q.heightImg = e.heightImg
	q.x = e.x
	q.y = e.y
	q.width = e.width
	q.height = e.height

	q.width = e.width / 2
	q.height = e.height / 2
	upLeft = q.collision(element)

	q.x = e.x + q.width
	q.y = e.y
	upRight = q.collision(element)

	q.x = e.x
	q.y = e.y + q.height
	downLeft = q.collision(element)

	q.x = e.x + q.width
	q.y = e.y + q.height
	downRight = q.collision(element)

	if downLeft || downRight {
		log.Printf("down: %v", e.GetYBox()+e.GetHeightBox()-element.GetYBox())
	}
	if upLeft || upRight {
		log.Printf("up: %v", element.GetYBox()+element.GetHeightBox()-e.GetYBox())
	}
	if upLeft || downLeft {
		log.Printf("left: %v", element.GetXBox()+element.GetWidthBox()-e.GetXBox())
	}
	if upRight || downRight {
		log.Printf("right: %v", e.GetXBox()+e.GetWidthBox()-element.GetXBox())
	}
	return
}

func (e *Box) collisionUpLeft(element Box) (collision bool) {
	a := Box{}
	a.x = e.x
	a.y = e.y
	a.width = e.width / 2
	a.height = e.height / 2
	return a.collision(element)
}

func (e *Box) collisionUpRight(element Box) (collision bool) {
	b := Box{}
	b.x = e.x + e.width/2
	b.y = e.y
	b.width = e.width / 2
	b.height = e.height / 2
	return b.collision(element)
}

func (e *Box) collisionDownLeft(element Box) (collision bool) {
	c := Box{}
	c.x = e.x
	c.y = e.y + e.height/2
	c.width = e.width / 2
	c.height = e.height / 2
	return c.collision(element)
}

func (e *Box) collisionDownRight(element Box) (collision bool) {
	d := Box{}
	d.x = e.x + e.width/2
	d.y = e.y + e.height/2
	d.width = e.width / 2
	d.height = e.height / 2
	return d.collision(element)
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
