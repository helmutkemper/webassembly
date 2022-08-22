package html

import (
	"errors"
	"syscall/js"
	"time"
)

// SpriteScene
//
// English:
//
// # Archive each sprite individually
//
// Português:
//
// Arquiva cada sprite individualmente
type spriteScene struct {
	sprite    *Sprite
	intervel  time.Duration
	imageData js.Value
}

// Add
//
// English:
//
// Adds a sprite to the scene.
//
//	Input::
//	   col: column where the desired sprite is located. Starts at zero.
//	   row: line where the desired sprite is. Starts at zero.
//	   interval: time for the next sprite change.
//	   flipHorizontal: flips the image horizontally.
//	   flipVertical: flips the image vertically.
//
//	Notes:
//	  * Use interval equal to zero to stop a move.
//
// Português:
//
// Adiciona uma sprite a cena.
//
//	Entrada:
//	   col: coluna onde a sprite desejada se encontra. Começa em zero.
//	   row: linha onde a sprite desejada se encontra. Começa em zero.
//	   interval: tempo para a próxima troca de sprite.
//	   flipHorizontal: inverte a imagem no sentido horizontal.
//	   flipVertical: inverte a imagem no sentido vertical.
//
//	Notas:
//	  * Use interval igual a zero para parar um movimento.
func (e *spriteScene) Add(col, row int, interval time.Duration, flipHorizontal, flipVertical bool) (ref *spriteScene) {
	canvas := new(TagCanvas)
	canvas.Init(e.sprite.spriteWidth, e.sprite.spriteHeight)

	var imageData js.Value
	var imageDataFlipped js.Value

	canvas.context.Call("clearRect", 0, 0, e.sprite.spriteWidth, e.sprite.spriteHeight)
	canvas.context.Call("drawImage", e.sprite.img.Get(), col*e.sprite.spriteWidth, row*e.sprite.spriteHeight, e.sprite.spriteWidth, e.sprite.spriteHeight, 0, 0, e.sprite.spriteWidth, e.sprite.spriteHeight)
	imageData = canvas.context.Call("getImageData", 0, 0, e.sprite.spriteWidth, e.sprite.spriteHeight)

	if flipHorizontal == false && flipVertical == false {
		e.sprite.scene[e.sprite.sceneName] = append(e.sprite.scene[e.sprite.sceneName], spriteScene{imageData: imageData, intervel: interval})
		return e
	}

	data1 := imageData.Get("data")

	imageDataFlipped = canvas.context.Call("getImageData", 0, 0, e.sprite.spriteWidth, e.sprite.spriteHeight)
	data2 := imageDataFlipped.Get("data")

	if flipHorizontal == true && flipVertical == true {
		e.sprite.flipData(data1, data2)
	} else if flipHorizontal == true && flipVertical == false {
		e.sprite.flipDataHorizontal(data1, data2)
	} else if flipHorizontal == false && flipVertical == true {
		e.sprite.flipDataVertival(data1, data2)
	}

	e.sprite.scene[e.sprite.sceneName] = append(e.sprite.scene[e.sprite.sceneName], spriteScene{imageData: imageDataFlipped, intervel: interval})

	return e
}

// Sprite
//
// English:
//
// Sprite lets you add a sequence of images and gives a sense of movement.
//
// Português:
//
// Sprite permite adicionar uma sequência de imagens e dá sensação de movimento.
type Sprite struct {
	canvas               *TagCanvas
	img                  Compatible
	spriteWidth          int
	spriteHeight         int
	x                    int
	y                    int
	clearRectDeltaX      int
	clearRectDeltaY      int
	clearRectDeltaWidth  int
	clearRectDeltaHeight int
	init                 bool
	onChange             *chan struct{}
	onEnd                *chan struct{}

	scene           map[string][]spriteScene
	sceneName       string
	imageDataToDraw js.Value
	running         bool
	stopCh          chan struct{}
}

// Y
//
// English:
//
// Y coordinate of where the drawing will be inserted on the canvas.
//
// Português:
//
// Cordenada Y de onde o desenho será inserido no canvas.
func (e *Sprite) Y(y int) (ref *Sprite) {
	e.y = y
	return e
}

// X
//
// English:
//
// X coordinate of where the drawing will be inserted on the canvas.
//
// Português:
//
// Cordenada X de onde o desenho será inserido no canvas.
func (e *Sprite) X(x int) (ref *Sprite) {
	e.x = x
	return e
}

// OnEnd
//
// English:
//
// Channel fired when a sena is completed.
//
// Português:
//
// Channel disparado quando uma sena é terminada.
func (e *Sprite) OnEnd(onEnd *chan struct{}) (ref *Sprite) {
	*onEnd = make(chan struct{}, 1)
	e.onEnd = onEnd
	return e
}

// OnChange
//
// English:
//
// Channel fired when a sprite is switched.
//
// Português:
//
// Channel disparado quando uma sprite é trocada.
func (e *Sprite) OnChange(onChange *chan struct{}) (ref *Sprite) {
	*onChange = make(chan struct{}, 1)
	e.onChange = onChange
	return e
}

// GetScene
//
// English:
//
// Returns the name of the current scene.
//
// Português:
//
// Retorna o nome da sena atual.
func (e *Sprite) GetScene() (name string) {
	return e.sceneName
}

// Scene
//
// English:
//
// Adds a new scene.
//
// Português:
//
// Adiciona uma nova sena.
func (e *Sprite) Scene(name string) (ref *spriteScene) {

	if e.scene == nil {
		e.scene = make(map[string][]spriteScene)
	}

	e.sceneName = name
	e.scene[e.sceneName] = make([]spriteScene, 0)

	ref = new(spriteScene)
	ref.sprite = e

	return
}

// Start
//
// English:
//
// Start a scene.
//
// Português:
//
// Inicia uma cena.
func (e *Sprite) Start(name string) (err error) {

	e.sceneName = name

	if e.running == true {
		e.stopCh <- struct{}{}
	}

	data, ok := e.scene[name]
	if ok == false {
		err = errors.New("scene name not found")
		return
	}

	e.imageDataToDraw = data[0].imageData
	if e.onChange != nil && len(*e.onChange) > 1 {
		*e.onChange <- struct{}{}
	}

	if len(data) == 1 {
		e.stopCh = make(chan struct{})
		return
	}

	e.running = true

	go func() {
		var i = 0
		var l = len(data) - 1

		var timer = time.NewTimer(data[i].intervel)
		for {
			select {
			case <-e.stopCh:
				e.running = false
				return
			case <-timer.C:
				if i < l {
					i += 1
				} else {
					if e.onEnd != nil && len(*e.onEnd) == 0 {
						*e.onEnd <- struct{}{}
					}

					i = 0
				}

				e.imageDataToDraw = data[i].imageData
				if e.onChange != nil && len(*e.onChange) == 0 {
					*e.onChange <- struct{}{}
				}

				if data[i].intervel == 0 {
					e.running = false
					return
				}

				timer = time.NewTimer(data[i].intervel)
			}
		}
	}()

	return
}

// SpriteHeight
//
// English:
//
// Sprite height in pixels.
//
// Português:
//
// Altura da sprite em pixels.
func (e *Sprite) SpriteHeight(spriteHeight int) (ref *Sprite) {
	e.spriteHeight = spriteHeight
	return e
}

// SpriteWidth
//
// English:
//
// Sprite length in pixels.
//
// Português:
//
// Comprimento da sprite em pixels.
func (e *Sprite) SpriteWidth(spriteWidth int) (ref *Sprite) {
	e.spriteWidth = spriteWidth
	return e
}

// Canvas
//
// English:
//
// Reference to the canvas element where the animation will take place.
//
// Português:
//
// Referência ao elemento canvas onde a animação ocorrerá.
func (e *Sprite) Canvas(canvas *TagCanvas) (ref *Sprite) {
	e.canvas = canvas
	return e
}

// Image
//
// English:
//
// Image path with the sprite to be animated.
//
// Português:
//
// Caminho da imagem com a sprite a ser animada.
func (e *Sprite) Image(src string) (ref *Sprite) {
	img := &TagImg{}
	img.Init()
	img.Src(src, true)

	e.img = img
	return e
}

// Init
//
// English:
//
// Initializes the Sprite object.
//
// Português:
//
// Inicializa o objeto Sprite.
func (e *Sprite) Init() {

	e.stopCh = make(chan struct{})
	e.init = true
}

// Draw
//
// English:
//
// Function to be added to the engine's Draw() function.
//
// Português:
//
// Função a ser adicionada a funçãp Draw() da engine.
func (e *Sprite) Draw() {
	if e.imageDataToDraw.IsUndefined() == true {
		return
	}

	e.canvas.context.Call("putImageData", e.imageDataToDraw, e.x, e.y)
}

// flipData
//
// English:
//
// Reverses the direction of the image, both vertically and horizontally.
//
// Português:
//
// Reverte o sentido da imagem, tanto na vertical quanto na horizontal.
func (e *Sprite) flipData(dataJs, flipped js.Value) {
	var x, y int

	var rgbaLength = 4

	var i = 0
	x = 0
	y = 0

	m := e.spriteHeight * e.spriteWidth * 4
	for y = 0; y != e.spriteHeight; y += 1 {
		for x = 0; x != e.spriteWidth; x += 1 {

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			flipped.SetIndex(m-i+0, uint8(dataJs.Index(i+0).Int()))
			flipped.SetIndex(m-i+1, uint8(dataJs.Index(i+1).Int()))
			flipped.SetIndex(m-i+2, uint8(dataJs.Index(i+2).Int()))
			flipped.SetIndex(m-i+3, uint8(dataJs.Index(i+3).Int()))

			i += rgbaLength
		}
	}
}

// flipDataHorizontal
//
// English:
//
// Reverses the direction of the image horizontally.
//
// Português:
//
// Inverte o sentido da imagem na horizontal.
func (e *Sprite) flipDataHorizontal(dataJs, flipped js.Value) {
	var x, y, m, i int

	var rgbaLength = 4

	for y = 0; y != e.spriteHeight; y += 1 {
		for x = 0; x != e.spriteWidth; x += 1 {
			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			flipped.SetIndex(m+0+(e.spriteWidth-x-1)*4, uint8(dataJs.Index(i+0).Int()))
			flipped.SetIndex(m+1+(e.spriteWidth-x-1)*4, uint8(dataJs.Index(i+1).Int()))
			flipped.SetIndex(m+2+(e.spriteWidth-x-1)*4, uint8(dataJs.Index(i+2).Int()))
			flipped.SetIndex(m+3+(e.spriteWidth-x-1)*4, uint8(dataJs.Index(i+3).Int()))

			i += rgbaLength
		}
		m = m + e.spriteWidth*4
	}
}

// flipDataVertival
//
// English:
//
// Reverses the direction of the image vertically.
//
// Português:
//
// Inverte o sentido da imagem na vertical.
func (e *Sprite) flipDataVertival(dataJs, flipped js.Value) {
	var x, y, m, i int

	var rgbaLength = 4

	for y = 0; y != e.spriteHeight; y += 1 {

		// English: Shifts the data position in the matrix to an ordinary two-dimensional matrix.
		// Português: Desloca a posição do dado na matriz para uma matriz bidimencional comum.
		m = (e.spriteHeight*e.spriteWidth - e.spriteWidth*(y+1)) + 1

		for x = 0; x != e.spriteWidth; x += 1 {
			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			// English:   +x shifts the data in the common matrix
			//            *4 shifts the data in the js image array
			// Português: +x desloca o dado na matriz comum
			//            *4 desloca o dado na matriz de imagem js
			flipped.SetIndex((m+x)*4+0, uint8(dataJs.Index(i+0).Int()))
			flipped.SetIndex((m+x)*4+1, uint8(dataJs.Index(i+1).Int()))
			flipped.SetIndex((m+x)*4+2, uint8(dataJs.Index(i+2).Int()))
			flipped.SetIndex((m+x)*4+3, uint8(dataJs.Index(i+3).Int()))

			i += rgbaLength
		}

	}
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
