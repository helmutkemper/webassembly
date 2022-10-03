package html

import (
	"image/color"
	"math"
	"syscall/js"
)

// CollisionBox
//
// English:
//
// Automates collision detection using an auto-size box.
//
// Português:
//
// Automatiza a detecção de colisão usando uma caixa de tamanho automático.
type CollisionBox struct {
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

// X
//
// English:
//
// Sets the X of the collision box (X of the image)
//
// Português:
//
// Define o X da caixa de colisão (X da imagem)
func (e *CollisionBox) X(x int) {
	e.xImg = x
}

// Y
//
// English:
//
// Sets the Y of the collision box (Y of the image)
//
// Português:
//
// Define o Y da caixa de colisão (Y da imagem)
func (e *CollisionBox) Y(y int) {
	e.yImg = y
}

// GetX
//
// English:
//
// Returns the X of the collision box (X of the image)
//
// Português:
//
// Retorna o valor de X da caixa de colisão (X da imagem)
func (e *CollisionBox) GetX() (x int) {
	return e.xImg
}

// GetY
//
// English:
//
// Returns the Y of the collision box (Y of the image)
//
// Português:
//
// Retorna o valor de Y da caixa de colisão (Y da imagem)
func (e *CollisionBox) GetY() (y int) {
	return e.yImg
}

// GetXBox
//
// English:
//
// Returns the value of X from the optimized collision box within the image.
//
// Português:
//
// Retorna o valor de X da caixa de colisão otimizada, dentro da imagem.
func (e *CollisionBox) GetXBox() (x int) {
	return e.xImg + e.x
}

// GetYBox
//
// English:
//
// Returns the value of Y from the optimized collision box within the image.
//
// Português:
//
// Retorna o valor de Y da caixa de colisão otimizada, dentro da imagem.
func (e *CollisionBox) GetYBox() (y int) {
	return e.yImg + e.y
}

// GetWidthBox
//
// English:
//
// Returns the width of the optimized collision box within the image.
//
// Português:
//
// Retorna o comprimento da caixa de colisão otimizada, dentro da imagem.
func (e *CollisionBox) GetWidthBox() (width int) {
	return e.width
}

// GetHeightBox
//
// English:
//
// Returns the height of the optimized collision box within the image.
//
// Português:
//
// Retorna a altura da caixa de colisão otimizada, dentro da imagem.
func (e *CollisionBox) GetHeightBox() (height int) {
	return e.height
}

// GetWidth
//
// English:
//
// Returns the width of the image.
//
// Português:
//
// Retorna o comprimento da imagem.
func (e *CollisionBox) GetWidth() (width int) {
	return e.widthImg
}

// GetHeight
//
// English:
//
// Returns the height of the image.
//
// Português:
//
// Retorna o altura da imagem.
func (e *CollisionBox) GetHeight() (height int) {
	return e.heightImg
}

// GetData
//
// English:
//
// Returns the data used to form the image on the canvas.
//
// Português:
//
// Retorna o dado usado para formar a imagem no canvas.
func (e *CollisionBox) GetData() (data js.Value) {
	return e.data
}

// Debug
//
// English:
//
// Paint two colored lines, one in the image box and one in the optimized collision box.
//
// Português:
//
// Pinta duas linhas coloridas, uma na caixa da imagem e outra na caixa de colisão otimizada.
func (e *CollisionBox) Debug(boxImageColor, boxCollisionColor *color.RGBA) {
	e.boxImageColor = boxImageColor
	e.boxCollisionColor = boxCollisionColor
}

// Init
//
// English:
//
// Inicializa o objeto
//
//	Input:
//	  data: Data used to create the image on the canvas. Use the canvas.GetImageData() function;
//	  width: image width;
//	  height: image height.
//
// Português:
//
// Inicializa o objeto
//
//	Entrada:
//	  data: Dado usado para criar a imagem no canvas. Use a função canvas.GetImageData();
//	  width: comprimento da imagem;
//	  height: altura da imagem.
func (e *CollisionBox) Init(data js.Value, width, height int) {
	e.data = data
	e.widthImg = width
	e.heightImg = height

	if e.collisionDataFunction == nil {
		e.collisionDataFunction = func(r, g, b, a uint8) bool {
			return a != 0
		}
	}

	e.getCollisionBoxFromImageData()

	if e.boxCollisionColor != nil && e.boxImageColor != nil {
		e.debug(*e.boxCollisionColor, *e.boxImageColor)
	}
}

// DataFunction
//
// English:
//
// Default function used to detect the optimized collision box.
//
//	 Default:
//	   func(r, g, b, a uint8) bool {
//				return a != 0
//			}
//
// Português:
//
// Função padrão usada para detectar a caixa de colisão otimizada.
//
//	 Padrão:
//	   func(r, g, b, a uint8) bool {
//				return a != 0
//			}
func (e *CollisionBox) DataFunction(f func(red, green, blue, alpha uint8) bool) {
	e.collisionDataFunction = f
}

// getCollisionDataFromImageData
//
// English:
//
// Returns a Boolean slice formed by [y][x]color.RGBA.A != 0
//
// Português:
//
// Retorna um slice booleano formado por [y][x]color.RGBA.A != 0
func (e *CollisionBox) getCollisionDataFromImageData() (data [][]bool) {
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

// debug
//
// English:
//
// Draw two colored squares around the image and collision-optimized boxes.
//
// Português:
//
// Desenha dois quadrados coloridos em torno das caixas da imagem e de colisão otimizada.
func (e *CollisionBox) debug(collisionColor, imageColor color.RGBA) {
	var rgbaLength = 4

	var i = 0
	var col int
	var row int

	for row = 0; row != e.heightImg; row += 1 {
		for col = 0; col != e.widthImg; col += 1 {

			//CollisionBox de colisão - início
			//A := true //col == x || col == x+width
			//B := true //row == y || row == y+height
			//C := col >= x && col <= x+width
			//D := row >= y && row <= y+height
			//if (A && D) && (B && C) {}
			//CollisionBox de colisão - fim

			// Image
			A := col == 0 || col == e.widthImg-1
			B := row == 0 || row == e.heightImg-1
			C := col >= 0 && col <= e.widthImg-1
			D := row >= 0 && row <= e.heightImg-1

			if (A && D) || (B && C) {
				e.data.Get("data").SetIndex(i+0, int(imageColor.R))
				e.data.Get("data").SetIndex(i+1, int(imageColor.G))
				e.data.Get("data").SetIndex(i+2, int(imageColor.B))
				e.data.Get("data").SetIndex(i+3, int(imageColor.A))
			}

			// Collision
			A = col == e.x || col == e.x+e.width
			B = row == e.y || row == e.y+e.height
			C = col >= e.x && col <= e.x+e.width
			D = row >= e.y && row <= e.y+e.height

			if (A && D) || (B && C) {
				e.data.Get("data").SetIndex(i+0, int(collisionColor.R))
				e.data.Get("data").SetIndex(i+1, int(collisionColor.G))
				e.data.Get("data").SetIndex(i+2, int(collisionColor.B))
				e.data.Get("data").SetIndex(i+3, int(collisionColor.A))
			}

			i += rgbaLength
		}
	}
}

// getCollisionBoxFromImageData
//
// English:
//
// Calculates the optimized collision box.
//
// Português:
//
// Calcula a caixa de colisão otimizada.
func (e *CollisionBox) getCollisionBoxFromImageData() {

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

func (e *CollisionBox) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (e *CollisionBox) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// collision
//
// English:
//
// Detect collision. This function is not public, as it fails by 1px with the quadrant function, due to integer division.
//
// Português:
//
// Detecta colisão. Esta função não é pública, pois, pois, ela falha em 1px com a função de quadrante, devido a divisão de inteiro.
func (e *CollisionBox) collision(element CollisionBox) (collision bool) {
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
// Returns the collision quadrant where the tested element hit the CollisionBox.
//
// Português:
//
// Retorna o quadrante da colisão onde o elemento testado bateu na caixa.
func (e *CollisionBox) Quadrant(element CollisionBox) (upLeft, upRight, downLeft, downRight bool) {
	q := CollisionBox{}
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

	return
}

// DistanceCorrection
//
// English:
//
// Returns the distance in pixels when one object overlaps another.
//
// Português:
//
// Retorna a distância em pixels quando um objeto se sobrepõe a outro.
func (e *CollisionBox) DistanceCorrection(element CollisionBox) (up, right, down, left int) {
	up = element.GetYBox() + element.GetHeightBox() - e.GetYBox()
	right = e.GetXBox() + e.GetWidthBox() - element.GetXBox()
	down = e.GetYBox() + e.GetHeightBox() - element.GetYBox()
	left = element.GetXBox() + element.GetWidthBox() - e.GetXBox()
	return
}

// collisionUpLeft
//
// English:
//
// Returns if left and up collision is heard, used for testing only.
//
// Português:
//
// Retorna se ouve colisão a esquerda e em cima, usado apenas para teste.
func (e *CollisionBox) collisionUpLeft(element CollisionBox) (collision bool) {
	a := CollisionBox{}
	a.x = e.x
	a.y = e.y
	a.width = e.width / 2
	a.height = e.height / 2
	return a.collision(element)
}

// collisionUpRight
//
// English:
//
// Returns if right and up collision is heard, used for testing only.
//
// Português:
//
// Retorna se ouve colisão a direita e em cima, usado apenas para teste.
func (e *CollisionBox) collisionUpRight(element CollisionBox) (collision bool) {
	b := CollisionBox{}
	b.x = e.x + e.width/2
	b.y = e.y
	b.width = e.width / 2
	b.height = e.height / 2
	return b.collision(element)
}

// collisionDownLeft
//
// English:
//
// Returns if left and down collision is heard, used for testing only.
//
// Português:
//
// Retorna se ouve colisão a esquerda e em baixo, usado apenas para teste.
func (e *CollisionBox) collisionDownLeft(element CollisionBox) (collision bool) {
	c := CollisionBox{}
	c.x = e.x
	c.y = e.y + e.height/2
	c.width = e.width / 2
	c.height = e.height / 2
	return c.collision(element)
}

// collisionDownRight
//
// English:
//
// Returns if right and down collision is heard, used for testing only.
//
// Português:
//
// Retorna se ouve colisão a direita e em baixo, usado apenas para teste.
func (e *CollisionBox) collisionDownRight(element CollisionBox) (collision bool) {
	d := CollisionBox{}
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
