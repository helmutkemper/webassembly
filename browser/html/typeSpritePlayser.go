package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
	"log"
	"sync"
	"time"
)

type SpriteStatus string

const (
	KSpriteStatusRightDown                SpriteStatus = "rightDown"
	KSpriteStatusRightUp                  SpriteStatus = "rightUp"
	KSpriteStatusRightUpFirst             SpriteStatus = "rightUpFirst"
	KSpriteStatusLeftDown                 SpriteStatus = "leftDown"
	KSpriteStatusLeftUp                   SpriteStatus = "leftUp"
	KSpriteStatusLeftUpFirst              SpriteStatus = "leftUpFirst"
	KSpriteStatusPlayerRunningHorizontal  SpriteStatus = "runningHorizonatl"
	KSpriteStatusPlayerStoppingHorizontal SpriteStatus = "stoppingHorizonatl"
	KSpriteStatusPlayerRunningVertical    SpriteStatus = "runningVertical"
	KSpriteStatusPlayerStoppingVertical   SpriteStatus = "stoppingVertical"

	KSpriteStatusGravityStart SpriteStatus = "gravityStart"
	KSpriteStatusGravityStop  SpriteStatus = "gravityStop"
)

type SpriteStatusList struct {
	sync sync.Mutex
	list []SpriteStatus
}

func (e *SpriteStatusList) Debug() {
	e.sync.Lock()
	defer e.sync.Unlock()

	log.Printf("%v", e.list)
}

func (e *SpriteStatusList) Init() {
	e.sync.Lock()
	defer e.sync.Unlock()

	e.list = make([]SpriteStatus, 0)
}

func (e *SpriteStatusList) Verify(statusList ...SpriteStatus) (ok bool) {
	e.sync.Lock()
	defer e.sync.Unlock()

	if len(e.list) == 0 || len(statusList) == 0 {
		return false
	}

	for _, status := range statusList {
		pass := false
		for k := range e.list {
			if e.list[k] == status {
				pass = true
				break
			}
		}
		if pass == true {
			continue
		}
		return false
	}

	return true
}

func (e *SpriteStatusList) Add(statusList ...SpriteStatus) {
	e.sync.Lock()
	defer e.sync.Unlock()

	for _, status := range statusList {
		for k := range e.list {
			if e.list[k] == status {
				return
			}
		}

		e.list = append(e.list, status)
	}
}

func (e *SpriteStatusList) Remove(statusList ...SpriteStatus) {
	e.sync.Lock()
	defer e.sync.Unlock()

	for _, status := range statusList {
		for k := range e.list {
			if e.list[k] == status {
				e.list = append(e.list[:k], e.list[k+1:]...)
				break
			}
		}
	}
}

type AccelerationFunction func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector DirectionHorizontal, yVector DirectionVertical) (dx, dy float64)

type DirectionHorizontal bool

const (
	KRight DirectionHorizontal = true
	KLeft  DirectionHorizontal = false
)

type DirectionVertical bool

const (
	KUp   DirectionVertical = true
	KDown DirectionVertical = false
)

type SpritePlayer struct {
	spt   *Sprite
	stage stage.Functions

	status SpriteStatusList

	x                 float64
	y                 float64
	velocityX         float64
	velocityY         float64
	velocityXInertial float64
	velocityYInertial float64

	xVector        DirectionHorizontal
	yVector        DirectionVertical
	xVectorDesired DirectionHorizontal
	yVectorDesired DirectionVertical

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

	horizontalFunction AccelerationFunction
	verticalFunction   AccelerationFunction

	runningLeftStartFunction  AccelerationFunction
	runningLeftStopFunction   AccelerationFunction
	runningRightStartFunction AccelerationFunction
	runningRightStopFunction  AccelerationFunction
	gravityFunction           AccelerationFunction

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
func (e *SpritePlayer) GravityFunc(f AccelerationFunction) (ref *SpritePlayer) {
	e.gravityFunction = f

	return e
}

func (e *SpritePlayer) RunningLeftStartFunc(f AccelerationFunction) (ref *SpritePlayer) {
	e.runningLeftStartFunction = f

	return e
}

func (e *SpritePlayer) RunningLeftStopFunc(f AccelerationFunction) (ref *SpritePlayer) {
	e.runningLeftStopFunction = f

	return e
}

func (e *SpritePlayer) RunningRightStartFunc(f AccelerationFunction) (ref *SpritePlayer) {
	e.runningRightStartFunction = f

	return e
}

func (e *SpritePlayer) RunningRightStopFunc(f AccelerationFunction) (ref *SpritePlayer) {
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

func (e *SpritePlayer) GetFormulaFloorVerySlipperyStart() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector DirectionHorizontal, yVector DirectionVertical) (dx, dy float64) {
		dx = 0.000001*float64(tRunning*tRunning) - inertialSpeedX/2
		if dx > 2 {
			dx = 2
		}

		dy = 0.0000015*float64(tRunning*tRunning) + 0.5
		if dy > 20 {
			dy = 20
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorVerySlipperyStop() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector DirectionHorizontal, yVector DirectionVertical) (dx, dy float64) {
		dx = velocityX - (0.00000007 * float64(tRunning*tRunning))
		if dx < 0 {
			dx = 0
		}

		dy = 0.0000015*float64(tRunning*tRunning) + 0.5
		if dy > 20 {
			dy = 20
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorLittleSlipperyStart() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector DirectionHorizontal, yVector DirectionVertical) (dx, dy float64) {
		dx = 0.000002 * float64(tRunning*tRunning)
		if dx > 2 {
			dx = 2
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorLittleSlipperyStop() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector DirectionHorizontal, yVector DirectionVertical) (dx, dy float64) {
		dx = velocityX - (0.000005 * float64(tRunning*tRunning))
		if dx < 0 {
			dx = 0
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorDefaultStart() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector DirectionHorizontal, yVector DirectionVertical) (dx, dy float64) {
		dx = 0.00001*float64(tRunning*tRunning) + 0.5
		if dx > 2 {
			dx = 2
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorDefaultStop() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector DirectionHorizontal, yVector DirectionVertical) (dx, dy float64) {
		dx = velocityX - (0.00002 * float64(tRunning*tRunning))
		if dx < 0 {
			dx = 0
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaGravityDefault() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector DirectionHorizontal, yVector DirectionVertical) (dx, dy float64) {
		dx = inertialSpeedX - 0.003*float64(tGravity)
		if dx < 0 {
			dx = 0
		}
		dy = 0.0000015*float64(tRunning*tRunning) + 0.5
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

	e.status.Init()

	e.gravityStart = time.Now()
	e.runningStart = time.Now()

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

	// limita direita/esquerda
	e.limitXMin = -1.0
	e.limitXMax = -1.0

	e.limitYMin = -1.0
	e.limitYMax = -1.0

	e.stage.AddMathFunctions(e.statusVerify)

	return e
}

func (e *SpritePlayer) statusVerify() {

	// sanitização - início

	// Teclas UP removem o status de teclas DOWN e vice-versa.
	if e.status.Verify(KSpriteStatusRightUp) == true {
		e.status.Remove(KSpriteStatusRightDown)
	}
	if e.status.Verify(KSpriteStatusLeftUp) == true {
		e.status.Remove(KSpriteStatusLeftDown)
	}
	if e.status.Verify(KSpriteStatusRightDown) == true {
		e.status.Remove(KSpriteStatusRightUp)
	}
	if e.status.Verify(KSpriteStatusLeftDown) == true {
		e.status.Remove(KSpriteStatusLeftUp)
	}

	// Tecla de movimentação DOWN remove a tecla de movimentação UP oposta.
	if e.status.Verify(KSpriteStatusLeftDown, KSpriteStatusRightUp) == true {
		e.status.Remove(KSpriteStatusRightUp)
	}
	if e.status.Verify(KSpriteStatusRightDown, KSpriteStatusLeftUp) == true {
		e.status.Remove(KSpriteStatusLeftUp)
	}

	// sanitização - fim

	//e.status.Debug()

	// Teclas UP não são mais usadas depois deste ponto.
	if e.status.Verify(KSpriteStatusRightUp) == true {
		e.status.Remove(KSpriteStatusRightUp)
		e.onRunningRightStop()
	}
	if e.status.Verify(KSpriteStatusLeftUp) == true {
		e.status.Remove(KSpriteStatusLeftUp)
		e.onRunningLeftStop()
	}

	// Caso uma tecla horizontal esteja precionada, o player anda
	if e.status.Verify(KSpriteStatusRightDown) == true {
		e.status.Add(KSpriteStatusPlayerRunningHorizontal)
	}
	if e.status.Verify(KSpriteStatusLeftDown) == true {
		e.status.Add(KSpriteStatusPlayerRunningHorizontal)
	}

	// inicio da gravidade
	if e.status.Verify(KSpriteStatusGravityStart) == true {
		e.status.Remove(KSpriteStatusGravityStart)
		e.onGravityStart()
		e.status.Add(KSpriteStatusPlayerRunningVertical)
	}

	if e.status.Verify(KSpriteStatusGravityStop) == true {
		e.status.Remove(KSpriteStatusGravityStop, KSpriteStatusPlayerRunningVertical)
	}
	// fim da gravidade

	if e.status.Verify(KSpriteStatusPlayerRunningHorizontal) == true {
		if e.horizontalFunction != nil {
			tRunning := time.Since(e.runningStart).Milliseconds()
			tGravity := time.Since(e.gravityStart).Milliseconds()
			e.velocityX, e.velocityY = e.horizontalFunction(e.velocityXInertial, e.velocityYInertial, e.velocityX, e.velocityY, e.x, e.y, tRunning, tGravity, e.xVector, e.yVector)
			if e.xVectorDesired == KRight {
				e.DX(e.velocityX)
			} else {
				e.DX(-e.velocityX)
			}

			if e.status.Verify(KSpriteStatusPlayerStoppingHorizontal) == true {
				if e.velocityX == 0 {
					e.status.Remove(KSpriteStatusPlayerRunningHorizontal, KSpriteStatusPlayerStoppingHorizontal)

					if e.runningStopSlip == false {
						e.StartStoppedRightSide()
					}
				}
			}
		}
	}

	if e.status.Verify(KSpriteStatusPlayerRunningVertical) == true {
		if e.verticalFunction != nil {
			tRunning := time.Since(e.runningStart).Milliseconds()
			tGravity := time.Since(e.gravityStart).Milliseconds()
			e.velocityX, e.velocityY = e.verticalFunction(e.velocityXInertial, e.velocityYInertial, e.velocityX, e.velocityY, e.x, e.y, tRunning, tGravity, e.xVector, e.yVector)
			if e.yVectorDesired == KUp {
				e.DY(-e.velocityY)
			} else {
				e.DY(e.velocityY)
			}

			if e.status.Verify(KSpriteStatusPlayerStoppingVertical) == true {
				if e.velocityY == 0 {
					e.status.Remove(KSpriteStatusPlayerRunningVertical, KSpriteStatusPlayerStoppingVertical)
				}
			}
		}
	}
}

func (e *SpritePlayer) X(x int) (ref *SpritePlayer) {
	e.xVector = float64(x) >= e.x

	e.x = float64(x)

	if e.limitXMax != -1 && e.x > e.limitXMax {
		e.x = e.limitXMax
	}

	if e.limitXMin != -1 && e.x < e.limitXMin {
		e.x = e.limitXMin
	}

	e.spt.x = int(e.x)

	return e
}

func (e *SpritePlayer) DX(x float64) (ref *SpritePlayer) {
	e.xVector = x+e.x >= e.x

	e.x = x + e.x

	if e.limitXMax != -1 && e.x > e.limitXMax {
		e.x = e.limitXMax
	}

	if e.limitXMin != -1 && e.x < e.limitXMin {
		e.x = e.limitXMin
	}

	e.spt.x = int(e.x)

	return e
}

func (e *SpritePlayer) Y(y int) (ref *SpritePlayer) {
	e.yVector = float64(y) < e.y

	e.y = float64(y)

	if e.limitYMax != -1 && e.x > e.limitYMax {
		e.y = e.limitYMax
	}

	if e.limitYMin != -1 && e.y < e.limitYMin {
		e.y = e.limitYMin
	}

	e.spt.y = int(e.y)

	return e
}

func (e *SpritePlayer) DY(y float64) (ref *SpritePlayer) {
	e.yVector = y+e.y < e.y

	e.y = y + e.y

	if e.limitYMax != -1 && e.x > e.limitYMax {
		e.y = e.limitYMax
	}

	if e.limitYMin != -1 && e.y < e.limitYMin {
		e.y = e.limitYMin
	}

	e.spt.y = int(e.y)

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

func (e *SpritePlayer) onRunningLeftStop() {
	if e.runningStopSlip == true {
		e.StartStoppedLeftSide()
	}

	e.horizontalFunction = e.runningLeftStopFunction

	if e.xVector == KRight {
		e.velocityXInertial = -e.velocityX
	} else {
		e.velocityXInertial = e.velocityX
	}

	e.runningStart = time.Now()

}

func (e *SpritePlayer) RunningLeftStop() {
	e.status.Add(
		KSpriteStatusLeftUp,
		//KSpriteStatusLeftUpFirst,
		KSpriteStatusPlayerStoppingHorizontal,
	)
}

func (e *SpritePlayer) RunningLeftStart() {
	e.xVectorDesired = KLeft
	e.StartWalkingLeftSide()
	e.horizontalFunction = e.runningLeftStartFunction

	if e.xVector == KRight {
		e.velocityXInertial = e.velocityX
	} else {
		e.velocityXInertial = -e.velocityX
	}

	e.runningStart = time.Now()

	e.status.Add(
		KSpriteStatusLeftDown,
		//KSpriteStatusLeftDownFirst,
	)
}

func (e *SpritePlayer) RunningRightStop() {
	//e.onRunningRightStop()

	e.status.Add(
		KSpriteStatusRightUp,
		//KSpriteStatusRightUpFirst,
		KSpriteStatusPlayerStoppingHorizontal,
	)
}

func (e *SpritePlayer) onRunningRightStop() {
	if e.runningStopSlip == true {
		e.StartStoppedRightSide()
	}

	e.horizontalFunction = e.runningRightStopFunction

	if e.xVector == KRight {
		e.velocityXInertial = -e.velocityX
	} else {
		e.velocityXInertial = e.velocityX
	}

	e.runningStart = time.Now()
}

func (e *SpritePlayer) RunningRightStart() {
	e.xVectorDesired = KRight
	e.StartWalkingRightSide()
	e.horizontalFunction = e.runningRightStartFunction

	if e.xVector == KRight {
		e.velocityXInertial = -e.velocityX
	} else {
		e.velocityXInertial = e.velocityX
	}

	e.runningStart = time.Now()
	e.status.Add(
		KSpriteStatusRightDown,
		//KSpriteStatusRightDownFirst,
	)
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

func (e *SpritePlayer) onGravityStart() {
	e.yVectorDesired = KDown

	if e.yVector == KUp {
		e.velocityYInertial = -e.velocityY
	} else {
		e.velocityYInertial = e.velocityY
	}

	e.verticalFunction = e.gravityFunction
	e.gravityStart = time.Now()
}

func (e *SpritePlayer) Gravity() {
	if e.gravityStarted == true {
		return
	}
	e.gravityStarted = true

	e.status.Add(
		KSpriteStatusGravityStart,
	)
}

func (e *SpritePlayer) GravityStop() {
	if e.gravityStarted == false {
		return
	}
	e.gravityStarted = false

	e.status.Add(
		KSpriteStatusGravityStop,
	)
}

func (e *SpritePlayer) GetCollisionBox() (box Box) {
	return e.spt.GetCollisionBox()
}
