package html

import (
	"errors"
	"syscall/js"
	"time"
)

type spriteConfig struct {
	data   []spriteScene
	sprite *Sprite
}

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
	intervel      time.Duration
	imageData     js.Value
	collisionData [][]bool
	collisionBox  Box
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
func (e *spriteConfig) Add(col, row int, interval time.Duration, flipHorizontal, flipVertical bool) (ref *spriteConfig) {
	canvas := new(TagCanvas)
	canvas.Init(e.sprite.spriteWidth, e.sprite.spriteHeight)

	var imageData js.Value

	canvas.ClearRect(0, 0, e.sprite.spriteWidth, e.sprite.spriteHeight)
	canvas.DrawImageComplete(e.sprite.img, col*e.sprite.spriteWidth, row*e.sprite.spriteHeight, e.sprite.spriteWidth, e.sprite.spriteHeight, 0, 0, e.sprite.spriteWidth, e.sprite.spriteHeight)
	imageData = canvas.GetImageData(0, 0, e.sprite.spriteWidth, e.sprite.spriteHeight, flipVertical, flipHorizontal)

	config := e.sprite.scene[e.sprite.sceneName]
	config.data = append(
		config.data,
		spriteScene{
			imageData:     imageData,
			intervel:      interval,
			collisionData: canvas.GetCollisionData(),
			collisionBox:  canvas.GetCollisionBox(),
		},
	)
	e.sprite.scene[e.sprite.sceneName] = config
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

	scene     map[string]spriteConfig
	sceneName string

	running   bool
	stopCh    chan struct{}
	stoppedCh chan struct{}

	// draw and colision - start
	imageData     js.Value
	collisionData [][]bool
	collisionBox  Box
	// draw and colision - end
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
func (e *Sprite) Scene(name string) (ref *spriteConfig) {

	if e.scene == nil {
		e.scene = make(map[string]spriteConfig)
	}

	e.sceneName = name
	data := make([]spriteScene, 0)
	config := spriteConfig{
		sprite: e,
		data:   data,
	}
	e.scene[e.sceneName] = config

	ref = new(spriteConfig)
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
	if e.sceneName == name {
		return
	}

	e.sceneName = name

	if e.running == true {
		e.stopCh <- struct{}{}
		<-e.stoppedCh
	}

	config, ok := e.scene[name]
	if ok == false {
		err = errors.New("scene name not found")
		return
	}

	e.imageData = config.data[0].imageData
	e.collisionData = config.data[0].collisionData
	e.collisionBox = config.data[0].collisionBox

	if e.onChange != nil && len(*e.onChange) > 1 {
		*e.onChange <- struct{}{}
	}

	if len(config.data) == 1 {
		return
	}

	e.running = true

	go func(data []spriteScene) {
		defer func() {
			e.stoppedCh <- struct{}{}
		}()

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

				e.imageData = data[i].imageData
				e.collisionData = data[i].collisionData
				e.collisionBox = data[i].collisionBox
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
	}(config.data)

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
	e.stoppedCh = make(chan struct{})
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
	if e.imageData.IsUndefined() == true {
		return
	}

	e.canvas.context.Call("putImageData", e.imageData, e.x, e.y)

	cBox := e.GetCollisionBox()
	e.canvas.StrokeRect(cBox.X, cBox.Y, cBox.Width, cBox.Height).Stroke()
}

func (e *Sprite) GetCollisionBox() (box Box) {
	cBox := e.collisionBox
	cBox.X += e.x
	cBox.Y += e.y

	return cBox
}

func (e *Sprite) TestCollisionBox(element CollisionBox) (collision bool) {
	thisCBox := e.collisionBox
	elementCBox := element.GetCollisionBox()
	if thisCBox.X < elementCBox.X+elementCBox.Width &&
		thisCBox.X+thisCBox.Width > elementCBox.X &&
		thisCBox.Y < elementCBox.Y+elementCBox.Height &&
		thisCBox.Y+thisCBox.Height > elementCBox.Y {
		return true
	}
	return false
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
