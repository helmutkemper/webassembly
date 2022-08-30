package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
	"github.com/helmutkemper/iotmaker.webassembly/mathUtil"
	"log"
	"time"
)

type SpritePlayer struct {
	spt   *Sprite
	stage stage.Functions

	deltaGravity   float64
	x              float64
	y              float64
	deltaX         float64
	deltaY         float64
	lastLeftSide   bool
	validKeyHDelta int
	validKeyVDelta int

	xFactor float64
	yFactor float64

	gravityStart  time.Time
	gravityMathId string

	limitXMin float64
	limitXMax float64
	limitYMin float64
	limitYMax float64
}

func (e *SpritePlayer) Init(stage stage.Functions, canvas *TagCanvas, imgPath string, width, height int) (ref *SpritePlayer) {
	e.stage = stage

	e.spt = new(Sprite)
	e.spt.Canvas(canvas)
	e.spt.Image(imgPath)
	e.spt.SpriteWidth(width)
	e.spt.SpriteHeight(height)
	e.spt.Init()

	//e.stage.AddDrawFunctions(e.spt.Draw)

	e.stage.AddMathFunctions(func() {
		e.ControlX()
		e.ControlY()
	})

	// andar direita/esquerda
	e.xFactor = 1.0
	// subir/descer
	e.yFactor = 1.0

	// limita direita/esquerda
	e.limitXMin = -1.0
	e.limitXMax = -1.0

	e.limitYMin = -1.0
	e.limitYMax = -1.0

	return e
}

func (e *SpritePlayer) Draw() {
	e.spt.Draw()
}

func (e *SpritePlayer) Corners(xMin, xMax, yMin, yMax int) (ref *SpritePlayer) {
	e.limitXMin = float64(xMin)
	e.limitXMax = float64(xMax)
	e.limitYMin = float64(yMin)
	e.limitYMax = float64(yMax)

	return e
}

func (e *SpritePlayer) CreateStoppedRightSide() (ref *spriteConfig) {
	return e.spt.Scene("stoppedRightSide")
}

func (e *SpritePlayer) StartStoppedRightSide() {
	err := e.spt.Start("stoppedRightSide")
	if err != nil {
		log.Printf("bug: StartStoppedRightSide()")
	}
}

func (e *SpritePlayer) CreateStoppedLeftSide() (ref *spriteConfig) {
	return e.spt.Scene("stoppedLeftSide")
}

func (e *SpritePlayer) StartStoppedLeftSide() {
	err := e.spt.Start("stoppedLeftSide")
	if err != nil {
		log.Printf("bug: StartStoppedLeftSide()")
	}
}

func (e *SpritePlayer) CreateWalkingRightSide() (ref *spriteConfig) {
	return e.spt.Scene("walkingRightSide")
}

func (e *SpritePlayer) StartWalkingRightSide() {
	err := e.spt.Start("walkingRightSide")
	if err != nil {
		log.Printf("bug: StartWalkingRightSide()")
	}
}

func (e *SpritePlayer) CreateWalkingLeftSide() (ref *spriteConfig) {
	return e.spt.Scene("walkingLeftSide")
}

func (e *SpritePlayer) StartWalkingLeftSide() {
	err := e.spt.Scene("walkingLeftSide")
	if err != nil {
		log.Printf("bug: StartWalkingLeftSide()")
	}
}

func (e *SpritePlayer) CreateFallRightSide() (ref *spriteConfig) {
	return e.spt.Scene("fallRightSide")
}

func (e *SpritePlayer) StartFallRightSide() {
	err := e.spt.Start("fallRightSide")
	if err != nil {
		log.Printf("bug: StartFallRightSide()")
	}
}

func (e *SpritePlayer) CreateFallLeftSide() (ref *spriteConfig) {
	return e.spt.Scene("fallLeftSide")
}

func (e *SpritePlayer) StartFallLeftSide() {
	err := e.spt.Start("fallLeftSide")
	if err != nil {
		log.Printf("bug: StartFallLeftSide()")
	}
}

func (e *SpritePlayer) KeyVerticalUp() {
	e.validKeyVDelta -= 1
}

func (e *SpritePlayer) KeyHorizontalUp() {
	e.validKeyHDelta -= 1
}

func (e *SpritePlayer) ControlX() {
	e.x = e.x + e.deltaX

	// English: space width - sprite width
	// Português: comprimento do espaço - comprimento da sprite
	if e.limitXMax != -1 && e.x > e.limitXMax {
		e.x = e.limitXMax
	}

	if e.limitXMin != -1 && e.x < e.limitXMin {
		e.x = e.limitXMin
	}

	e.spt.X(mathUtil.FloatToInt(e.x))
}

func (e *SpritePlayer) ControlY() {
	e.y = e.y + e.deltaY

	// English: space heght - sprite height
	// Português: altura do espaço - altura da sprite
	if e.limitYMax != -1 && e.y > e.limitYMax {
		e.y = e.limitYMax
	}

	if e.limitYMin != -1 && e.y < e.limitYMin {
		e.y = e.limitYMin
	}

	e.spt.Y(mathUtil.FloatToInt(e.y))
}

func (e *SpritePlayer) ControlDelta() {
	if e.validKeyHDelta <= 0 {
		e.deltaX = 0
		e.validKeyHDelta = 0
	}

	if e.validKeyVDelta <= 0 {
		e.deltaY = 0
		e.validKeyVDelta = 0
	}
}

func (e *SpritePlayer) PlayerStop() (cont bool) {
	var err error
	if e.validKeyHDelta <= 0 && e.validKeyVDelta <= 0 {
		if e.lastLeftSide == false {
			err = e.spt.Start("stoppedRightSide")
		} else {
			err = e.spt.Start("stoppedLeftSide")
		}
		if err != nil {
			log.Printf("error: %v", err)
		}

		return true
	}

	return false
}

func (e *SpritePlayer) PlayerUp() (cont bool) {
	if e.validKeyVDelta > 0 {
		return true
	}

	e.validKeyVDelta += 1
	e.deltaY = -e.yFactor
	return false
}

func (e *SpritePlayer) PlayerDown() (cont bool) {
	if e.validKeyVDelta > 0 {
		return true
	}

	e.validKeyVDelta += 1
	e.deltaY = e.yFactor
	return false
}

func (e *SpritePlayer) PlayerLeft() (cont bool) {
	if e.validKeyHDelta > 0 {
		return true
	}

	e.validKeyHDelta += 1
	e.deltaX = -e.xFactor
	e.lastLeftSide = true
	return false
}

func (e *SpritePlayer) PlayerRight() (cont bool) {
	if e.validKeyHDelta > 0 {
		return true
	}

	e.validKeyHDelta += 1
	e.deltaX = e.xFactor
	e.lastLeftSide = false
	return false
}

func (e *SpritePlayer) PlayerImageToUse() {
	var err error
	if e.lastLeftSide == true {
		err = e.spt.Start("walkingLeftSide")
		if err != nil {
			log.Printf("error: %v", err)
		}
	} else {
		err = e.spt.Start("walkingRightSide")
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
}

func (e *SpritePlayer) Gravity() {
	if e.gravityMathId != "" {
		return
	}

	if e.lastLeftSide == false {
		e.StartFallRightSide()
	} else {
		e.StartFallRightSide()
	}

	e.gravityStart = time.Now()
	e.gravityMathId, _ = e.stage.AddMathFunctions(e.gravitySupport)
}

func (e *SpritePlayer) gravitySupport() {
	if e.y >= e.limitYMax {
		return
	}
	e.y += e.deltaGravity

	// todo: fazer e eliminar constantes
	t := time.Since(e.gravityStart).Milliseconds()
	e.deltaGravity = 0.00001*float64(t*t) + 1.0
	if e.deltaGravity > 20 {
		e.deltaGravity = 20
	}

	e.ControlY()
	e.spt.Y(mathUtil.FloatToInt(e.y))

	if e.limitYMax != -1 && e.y >= e.limitYMax && e.gravityMathId != "" {
		e.GravityStop()
	}
}

func (e *SpritePlayer) GravityStop() {
	e.stage.DeleteMathFunctions(e.gravityMathId)
	e.gravityMathId = ""
}

func (e *SpritePlayer) GetCollisionBox() (box Box) {
	return e.spt.GetCollisionBox()
}
