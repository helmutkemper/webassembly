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

	x              float64
	y              float64
	deltaY         float64
	validKeyVDelta int

	yFactor float64

	velocityX float64
	velocityY float64

	// runningStart
	//
	// English:
	//
	// Receives the starting time of the race to be used in the acceleration formula, both in the character's sprint and
	// braking.
	//
	// Português:
	//
	// Recebe o tempo inicial da corrida para ser usado na fórmula de aceleração, tanto na arrancada quanto no freio do
	// personagem.
	runningStart time.Time

	// runningStopSlip
	//
	// English:
	//
	// When set to false, the character decelerates with the running movie clip, when true, decelerates with the stopped
	// movie clip.
	//
	// Português:
	//
	// Quando definido como false, o personagem desacelera com o movie clip de corrida, quando true, desacelera com o
	// movie clip de parado.
	runningStopSlip bool

	gravityDelta   float64
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

	runningLeftStartFunction  func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)
	runningLeftStopFunction   func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)
	runningRightStartFunction func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)
	runningRightStopFunction  func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)
	gravityFunction           func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)

	//stateMachine []func()
}

// StopSlip
//
// English:
//
// When set to false, the character decelerates with the running movie clip, when true, decelerates with the stopped
// movie clip.
//
// Português:
//
// Quando definido como false, o personagem desacelera com o movie clip de corrida, quando true, desacelera com o
// movie clip de parado.
func (e *SpritePlayer) StopSlip(slip bool) (ref *SpritePlayer) {
	e.runningStopSlip = slip

	return e
}

// GravityFunc
//
// English:
//
// Defines the formula used to calculate gravity.
//
// Português:
//
// Define a fórmula usada para cálcular a gravidade.
//
// Default / Padrão:
//
//	func(x, y float64, t int64) (dx, dy float64) {
//	  dy = 0.0000015*float64(t*t) + 0.5
//	  if dy > 20 {
//	    dy = 20
//	  }
//	  return
//	}
func (e *SpritePlayer) GravityFunc(f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) (ref *SpritePlayer) {
	e.gravityFunction = f

	return e
}

func (e *SpritePlayer) RunningLeftStartFunc(f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) (ref *SpritePlayer) {
	e.runningLeftStartFunction = f

	return e
}

func (e *SpritePlayer) RunningLeftStopFunc(f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) (ref *SpritePlayer) {
	e.runningLeftStopFunction = f

	return e
}

func (e *SpritePlayer) RunningRightStartFunc(f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) (ref *SpritePlayer) {
	e.runningRightStartFunction = f

	return e
}

func (e *SpritePlayer) RunningRightStopFunc(f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) (ref *SpritePlayer) {
	e.runningRightStopFunction = f

	return e
}

func (e *SpritePlayer) DefineFloorVerySlippery() (ref *SpritePlayer) {
	e.StopSlip(true)
	e.RunningRightStartFunc(e.GetFormulaFloorVerySlipperyStart())
	e.RunningLeftStartFunc(e.GetFormulaFloorVerySlipperyStart())
	e.RunningRightStopFunc(e.GetFormulaFloorVerySlipperyStop())
	e.RunningLeftStopFunc(e.GetFormulaFloorVerySlipperyStop())

	return e
}

func (e *SpritePlayer) DefineFloorLittleSlippery() (ref *SpritePlayer) {
	e.StopSlip(true)
	e.RunningRightStartFunc(e.GetFormulaFloorLittleSlipperyStart())
	e.RunningLeftStartFunc(e.GetFormulaFloorLittleSlipperyStart())
	e.RunningRightStopFunc(e.GetFormulaFloorLittleSlipperyStop())
	e.RunningLeftStopFunc(e.GetFormulaFloorLittleSlipperyStop())

	return e
}

func (e *SpritePlayer) GetFormulaFloorVerySlipperyStart() (f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) {
	return func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64) {
		dx = 0.000001 * float64(t*t)
		if dx > 2 {
			dx = 2
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorVerySlipperyStop() (f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) {
	return func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64) {
		dx = velocityX - (0.00000025 * float64(t*t))
		if dx < 0 {
			dx = 0
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorLittleSlipperyStart() (f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) {
	return func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64) {
		dx = 0.000005 * float64(t*t)
		if dx > 2 {
			dx = 2
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorLittleSlipperyStop() (f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) {
	return func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64) {
		dx = velocityX - (0.000005 * float64(t*t))
		if dx < 0 {
			dx = 0
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorDefaultStart() (f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) {
	return func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64) {
		dx = 0.00001*float64(t*t) + 0.5
		if dx > 2 {
			dx = 2
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorDefaultStop() (f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) {
	return func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64) {
		dx = velocityX - (0.00002 * float64(t*t))
		if dx < 0 {
			dx = 0
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaGravityDefault() (f func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64)) {
	return func(velocityX, velocityY, x, y float64, t int64) (dx, dy float64) {
		dy = 0.0000015*float64(t*t) + 0.5
		if dy > 20 {
			dy = 20
		}
		return
	}
}

func (e *SpritePlayer) Init(stage stage.Functions, canvas *TagCanvas, imgPath string, width, height int) (ref *SpritePlayer) {
	e.stage = stage

	//e.stateMachine = make([]func(), 0)
	//e.stage.AddMathFunctions(e.stateMachineExecute)

	e.spt = new(Sprite)
	e.spt.Canvas(canvas)
	e.spt.Image(imgPath)
	e.spt.SpriteWidth(width)
	e.spt.SpriteHeight(height)
	e.spt.Init()

	e.gravityFunction = e.GetFormulaGravityDefault()
	e.runningRightStartFunction = e.GetFormulaFloorDefaultStart()
	e.runningLeftStartFunction = e.GetFormulaFloorDefaultStart()
	e.runningRightStopFunction = e.GetFormulaFloorDefaultStop()
	e.runningLeftStopFunction = e.GetFormulaFloorDefaultStop()

	//e.stage.AddDrawFunctions(e.spt.Draw)

	// andar direita/esquerda
	//e.xFactor = 1.0
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

//func (e *SpritePlayer) KeyHorizontalUp() {
//	e.validKeyHDelta -= 1
//}

func (e *SpritePlayer) ControlX() {
	//e.x = e.x + e.deltaX

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

//func (e *SpritePlayer) PlayerStop() /*(cont bool)*/ {
//	var err error
//	if e.validKeyHDelta <= 0 && e.validKeyVDelta <= 0 {
//		if e.lastLeftSide == false {
//			err = e.spt.Start("stoppedRightSide")
//		} else {
//			err = e.spt.Start("stoppedLeftSide")
//		}
//		if err != nil {
//			log.Printf("error: %v", err)
//		}
//
//		//return true
//	}
//
//	//return false
//}

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

func (e *SpritePlayer) RunningLeftStop() {
	if e.stagePlayerLeftStartSupport == "" || e.stagePlayerLeftStopSupport != "" {
		return
	}

	if e.stagePlayerRightStartSupport != "" {
		e.stage.DeleteMathFunctions(e.stagePlayerRightStartSupport)
		e.stagePlayerRightStartSupport = ""
	}
	if e.stagePlayerRightStopSupport != "" {
		e.stage.DeleteMathFunctions(e.stagePlayerRightStopSupport)
		e.stagePlayerRightStopSupport = ""
	}

	if e.runningStopSlip == true {
		e.StartStoppedLeftSide()
	}

	e.stage.DeleteMathFunctions(e.stagePlayerLeftStartSupport)
	e.stagePlayerLeftStartSupport = ""

	e.runningStart = time.Now()
	e.stagePlayerLeftStopSupport, _ = e.stage.AddMathFunctions(e.runningLeftStopSupport)
}

func (e *SpritePlayer) runningLeftStopSupport() {
	// English: space width - sprite width
	// Português: comprimento do espaço - comprimento da sprite
	if e.limitXMax != -1 && e.x > e.limitXMax {
		e.x = e.limitXMax
	}

	if e.limitXMin != -1 && e.x < e.limitXMin {
		e.x = e.limitXMin
	}

	t := time.Since(e.runningStart).Milliseconds()
	if e.runningLeftStopFunction != nil {
		e.velocityX, e.velocityY = e.runningLeftStopFunction(e.velocityX, e.velocityY, e.x, e.y, t)
		e.x -= e.velocityX
		//e.y -= e.velocityY todo: testar

		if e.velocityX == 0 {
			if e.runningStopSlip == false {
				e.StartStoppedLeftSide()
			}

			e.stage.DeleteMathFunctions(e.stagePlayerLeftStopSupport)
			e.stagePlayerLeftStopSupport = ""
		}
	}

	e.spt.X(mathUtil.FloatToInt(e.x))
}

func (e *SpritePlayer) RunningLeftStart() {
	if e.stagePlayerLeftStartSupport != "" {
		return
	}

	//stop left
	if e.stagePlayerLeftStopSupport != "" {
		e.stage.DeleteMathFunctions(e.stagePlayerLeftStopSupport)
		e.stagePlayerLeftStopSupport = ""
	}
	//start right
	if e.stagePlayerRightStartSupport != "" {
		e.stage.DeleteMathFunctions(e.stagePlayerRightStartSupport)
		e.stagePlayerRightStartSupport = ""
	}
	//stop right
	if e.stagePlayerRightStopSupport != "" {
		e.stage.DeleteMathFunctions(e.stagePlayerRightStopSupport)
		e.stagePlayerRightStopSupport = ""
	}

	//e.lastLeftSide = true
	e.runningStart = time.Now()
	e.StartWalkingLeftSide()
	e.stagePlayerLeftStartSupport, _ = e.stage.AddMathFunctions(e.runningLeftStartSupport)
}

func (e *SpritePlayer) runningLeftStartSupport() {
	// English: space width - sprite width
	// Português: comprimento do espaço - comprimento da sprite
	if e.limitXMax != -1 && e.x > e.limitXMax {
		e.x = e.limitXMax
	}

	if e.limitXMin != -1 && e.x < e.limitXMin {
		e.x = e.limitXMin
	}

	t := time.Since(e.runningStart).Milliseconds()
	if e.runningLeftStartFunction != nil {
		e.velocityX, e.velocityY = e.runningLeftStartFunction(e.velocityX, e.velocityY, e.x, e.y, t)
		e.x -= e.velocityX
		//e.y -= e.velocityY //todo testar
	}

	e.spt.X(mathUtil.FloatToInt(e.x))
}

func (e *SpritePlayer) RunningRightStop() {
	if e.stagePlayerRightStartSupport == "" || e.stagePlayerRightStopSupport != "" {
		return
	}

	if e.stagePlayerLeftStartSupport != "" {
		e.stage.DeleteMathFunctions(e.stagePlayerLeftStartSupport)
		e.stagePlayerLeftStartSupport = ""
	}
	if e.stagePlayerLeftStopSupport != "" {
		e.stage.DeleteMathFunctions(e.stagePlayerLeftStopSupport)
		e.stagePlayerLeftStopSupport = ""
	}

	if e.runningStopSlip == true {
		e.StartStoppedRightSide()
	}

	e.stage.DeleteMathFunctions(e.stagePlayerRightStartSupport)
	e.stagePlayerRightStartSupport = ""

	e.runningStart = time.Now()
	e.stagePlayerRightStopSupport, _ = e.stage.AddMathFunctions(e.runningRightStopSupport)
}

func (e *SpritePlayer) runningRightStopSupport() {
	// English: space width - sprite width
	// Português: comprimento do espaço - comprimento da sprite
	if e.limitXMax != -1 && e.x > e.limitXMax {
		e.x = e.limitXMax
	}

	if e.limitXMin != -1 && e.x < e.limitXMin {
		e.x = e.limitXMin
	}

	t := time.Since(e.runningStart).Milliseconds()
	if e.runningRightStopFunction != nil {
		e.velocityX, e.velocityY = e.runningRightStopFunction(e.velocityX, e.velocityY, e.x, e.y, t)
		e.x += e.velocityX
		//e.y += e.velocityY //todo testar

		if e.velocityX == 0 {
			if e.runningStopSlip == false {
				e.StartStoppedRightSide()
			}

			e.stage.DeleteMathFunctions(e.stagePlayerRightStopSupport)
			e.stagePlayerRightStopSupport = ""
		}
	}

	e.spt.X(mathUtil.FloatToInt(e.x))
}

func (e *SpritePlayer) RunningRightStart() {
	if e.stagePlayerRightStartSupport != "" {
		return
	}

	//stop right
	if e.stagePlayerRightStopSupport != "" {
		e.stage.DeleteMathFunctions(e.stagePlayerRightStopSupport)
		e.stagePlayerRightStopSupport = ""
	}
	//start left
	if e.stagePlayerLeftStartSupport != "" {
		e.stage.DeleteMathFunctions(e.stagePlayerLeftStartSupport)
		e.stagePlayerLeftStartSupport = ""
	}
	//stop left
	if e.stagePlayerLeftStopSupport != "" {
		e.stage.DeleteMathFunctions(e.stagePlayerLeftStopSupport)
		e.stagePlayerLeftStopSupport = ""
	}

	//e.lastLeftSide = false
	e.runningStart = time.Now()
	e.StartWalkingRightSide()
	e.stagePlayerRightStartSupport, _ = e.stage.AddMathFunctions(e.runningRightStartSupport)
}

func (e *SpritePlayer) runningRightStartSupport() {
	// English: space width - sprite width
	// Português: comprimento do espaço - comprimento da sprite
	if e.limitXMax != -1 && e.x > e.limitXMax {
		e.x = e.limitXMax
	}

	if e.limitXMin != -1 && e.x < e.limitXMin {
		e.x = e.limitXMin
	}

	t := time.Since(e.runningStart).Milliseconds()
	if e.runningRightStartFunction != nil {
		e.velocityX, e.velocityY = e.runningRightStartFunction(e.velocityX, e.velocityY, e.x, e.y, t)
		e.x += e.velocityX
		//e.y += e.velocityY
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

//func (e *SpritePlayer) PlayerImageToUse() {
//	var err error
//	if e.lastLeftSide == true {
//		err = e.spt.Start("walkingLeftSide")
//		if err != nil {
//			log.Printf("error: %v", err)
//		}
//	} else {
//		err = e.spt.Start("walkingRightSide")
//		if err != nil {
//			log.Printf("error: %v", err)
//		}
//	}
//}

//func (e *SpritePlayer) Fall() {
//	if e.lastLeftSide == false {
//		e.StartFallRightSide()
//	} else {
//		e.StartFallLeftSide()
//	}
//}

func (e *SpritePlayer) Gravity() {
	if e.gravityStarted == true {
		return
	}
	e.gravityStarted = true

	log.Printf("gravity")
	if e.gravityMathId != "" {
		return
	}

	e.gravityDelta = 0

	e.gravityStart = time.Now()
	e.gravityMathId, _ = e.stage.AddMathFunctions(e.gravitySupport)
}

func (e *SpritePlayer) gravitySupport() {
	if e.limitYMax != -1 && e.y >= e.limitYMax {
		return
	}
	e.y += e.gravityDelta

	t := time.Since(e.gravityStart).Milliseconds()
	if e.gravityFunction != nil {
		_, e.gravityDelta = e.gravityFunction(0, 0, e.x, e.y, t)
	}

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
