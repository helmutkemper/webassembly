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
	deltaLeft      float64
	deltaRight     float64
	deltaY         float64
	lastLeftSide   bool
	validKeyHDelta int
	validKeyVDelta int

	xFactor float64
	yFactor float64

	gravityStart   time.Time
	gravityMathId  string
	gravityStarted bool

	limitXMin float64
	limitXMax float64
	limitYMin float64
	limitYMax float64

	stagePlayerRightStartSupport string
	stagePlayerRightStopSupport  string
	stagePlayerLeftStartSupport  string
	stagePlayerLeftStopSupport   string

	gravityFunction func(x, y float64, t int64) (dx, dy float64)
}

func (e *SpritePlayer) GravityFunc(f func(x, y float64, t int64) (dx, dy float64)) (ref *SpritePlayer) {
	e.gravityFunction = f

	return e
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

func (e *SpritePlayer) X(x int) (ref *SpritePlayer) {
	e.x = float64(x)
	e.spt.x = x

	return e
}

func (e *SpritePlayer) Y(y int) (ref *SpritePlayer) {
	e.y = float64(y)
	e.spt.y = y

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

func (e *SpritePlayer) CreateScene(name string) (ref *spriteConfig) {
	return e.spt.Scene(name)
}

func (e *SpritePlayer) CreateStoppedRightSide() (ref *spriteConfig) {
	return e.spt.Scene("stoppedRightSide")
}

func (e *SpritePlayer) StartStoppedRightSide() (ref *SpritePlayer) {
	err := e.spt.Start("stoppedRightSide")
	if err != nil {
		log.Printf("bug: StartStoppedRightSide()")
	}

	return e
}

func (e *SpritePlayer) CreateStoppedLeftSide() (ref *spriteConfig) {
	return e.spt.Scene("stoppedLeftSide")
}

func (e *SpritePlayer) StartStoppedLeftSide() (ref *SpritePlayer) {
	err := e.spt.Start("stoppedLeftSide")
	if err != nil {
		log.Printf("bug: StartStoppedLeftSide()")
	}

	return e
}

func (e *SpritePlayer) CreateWalkingRightSide() (ref *spriteConfig) {
	return e.spt.Scene("walkingRightSide")
}

func (e *SpritePlayer) StartWalkingRightSide() (ref *SpritePlayer) {
	err := e.spt.Start("walkingRightSide")
	if err != nil {
		log.Printf("bug: StartWalkingRightSide()")
	}

	return e
}

func (e *SpritePlayer) CreateWalkingLeftSide() (ref *spriteConfig) {
	return e.spt.Scene("walkingLeftSide")
}

func (e *SpritePlayer) StartWalkingLeftSide() (ref *SpritePlayer) {
	err := e.spt.Start("walkingLeftSide")
	if err != nil {
		log.Printf("bug: StartWalkingLeftSide()")
	}

	return e
}

func (e *SpritePlayer) CreateFallRightSide() (ref *spriteConfig) {
	return e.spt.Scene("fallRightSide")
}

func (e *SpritePlayer) StartFallRightSide() (ref *SpritePlayer) {
	err := e.spt.Start("fallRightSide")
	if err != nil {
		log.Printf("bug: StartFallRightSide()")
	}

	return e
}

func (e *SpritePlayer) CreateFallLeftSide() (ref *spriteConfig) {
	return e.spt.Scene("fallLeftSide")
}

func (e *SpritePlayer) StartFallLeftSide() (ref *SpritePlayer) {
	err := e.spt.Start("fallLeftSide")
	if err != nil {
		log.Printf("bug: StartFallLeftSide()")
	}

	return e
}

func (e *SpritePlayer) KeyVerticalUp() {
	e.validKeyVDelta -= 1
}

func (e *SpritePlayer) KeyHorizontalUp() {
	e.validKeyHDelta -= 1
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
	if e.validKeyVDelta <= 0 {
		e.deltaY = 0
		e.validKeyVDelta = 0
	}
}

func (e *SpritePlayer) PlayerStop() /*(cont bool)*/ {
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

		//return true
	}

	//return false
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

func (e *SpritePlayer) PlayerLeftStop() {
	if e.gravityStarted == true {
		return
	}

	e.StartStoppedLeftSide()
	e.stage.DeleteMathFunctions(e.stagePlayerLeftStartSupport)
	e.stagePlayerLeftStopSupport, _ = e.stage.AddMathFunctions(e.playerLeftStopSupport)
}

func (e *SpritePlayer) playerLeftStopSupport() {
	if e.deltaLeft == 0 {
		e.stage.DeleteMathFunctions(e.stagePlayerLeftStopSupport)
		return
	}

	e.x = e.x - e.deltaLeft

	if e.deltaLeft > 0 {
		e.deltaLeft -= 0.025
	} else if e.deltaLeft < 0 {
		e.deltaLeft = 0
	}

	e.spt.X(mathUtil.FloatToInt(e.x))
}

func (e *SpritePlayer) PlayerLeftStart() /*(cont bool)*/ {
	e.lastLeftSide = true
	e.StartWalkingLeftSide()
	e.stagePlayerLeftStartSupport, _ = e.stage.AddMathFunctions(e.playerLeftStartSupport)
}

func (e *SpritePlayer) playerLeftStartSupport() /*(cont bool)*/ {
	e.x = e.x - e.deltaLeft

	if e.deltaLeft < 1.5 {
		e.deltaLeft += 0.025
	}

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

func (e *SpritePlayer) PlayerRightStop() {
	if e.gravityStarted == true {
		return
	}

	e.StartStoppedRightSide()
	e.stage.DeleteMathFunctions(e.stagePlayerRightStartSupport)
	e.stagePlayerRightStopSupport, _ = e.stage.AddMathFunctions(e.playerRightStopSupport)
}

func (e *SpritePlayer) playerRightStopSupport() {
	if e.deltaRight == 0 {
		e.stage.DeleteMathFunctions(e.stagePlayerRightStopSupport)
		return
	}

	e.x = e.x + e.deltaRight

	if e.deltaRight > 0 {
		e.deltaRight -= 0.025
	} else if e.deltaRight < 0 {
		e.deltaRight = 0
	}

	e.spt.X(mathUtil.FloatToInt(e.x))
}

func (e *SpritePlayer) PlayerRightStart() {
	e.lastLeftSide = false
	e.StartWalkingRightSide()
	e.stagePlayerRightStartSupport, _ = e.stage.AddMathFunctions(e.playerRightStartSupport)
}

func (e *SpritePlayer) playerRightStartSupport() {
	e.x = e.x + e.deltaRight

	if e.deltaRight < 1.5 {
		e.deltaRight += 0.025
	}

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

func (e *SpritePlayer) WalkingLeft() {
	var err error
	err = e.spt.Start("walkingLeftSide")
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func (e *SpritePlayer) WalkingRight() {
	var err error
	err = e.spt.Start("walkingRightSide")
	if err != nil {
		log.Printf("error: %v", err)
	}
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

func (e *SpritePlayer) Fall() {
	if e.lastLeftSide == false {
		e.StartFallRightSide()
	} else {
		e.StartFallLeftSide()
	}
}

func (e *SpritePlayer) Gravity() {
	if e.gravityStarted == true {
		return
	}
	e.gravityStarted = true

	log.Printf("gravity")
	if e.gravityMathId != "" {
		return
	}

	e.deltaGravity = 0

	e.gravityStart = time.Now()
	e.gravityMathId, _ = e.stage.AddMathFunctions(e.gravitySupport)
}

func (e *SpritePlayer) gravitySupport() {
	if e.limitYMax != -1 && e.y >= e.limitYMax {
		return
	}
	e.y += e.deltaGravity

	// todo: fazer e eliminar constantes
	t := time.Since(e.gravityStart).Milliseconds()
	if e.gravityFunction != nil {
		_, e.deltaGravity = e.gravityFunction(e.x, e.y, t)
	}
	//e.deltaGravity = 0.00001*float64(t*t) + 1.0
	//if e.deltaGravity > 20 {
	//	e.deltaGravity = 20
	//}

	e.ControlY()

	if e.limitYMax != -1 && e.y >= e.limitYMax && e.gravityMathId != "" {
		e.GravityStop()
	}
}

func (e *SpritePlayer) GravityStop() {
	if e.gravityStarted == false {
		return
	}
	e.gravityStarted = false

	log.Printf("gravity stop")
	e.stage.DeleteMathFunctions(e.gravityMathId)
	e.gravityMathId = ""
}

func (e *SpritePlayer) GetCollisionBox() (box Box) {
	return e.spt.GetCollisionBox()
}
