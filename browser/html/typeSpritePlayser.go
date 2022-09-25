package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
	"github.com/helmutkemper/iotmaker.webassembly/platform/eventQueue"
	"log"
	"time"
)

const (
	moveRight                                       = "moveRight"
	moveRightConfig                                 = "moveRightConfig"
	moveRightStop                                   = "moveRightStop"
	moveLeft                                        = "moveLeft"
	moveLeftConfig                                  = "moveLeftConfig"
	KSpriteStatusLeftDownConfigured                 = "KSpriteStatusLeftDownConfigured"
	moveLeftStop                                    = "moveLeftStop"
	playerRunningHorizontal                         = "playerRunningHorizontal"
	playerStoppingHorizontal                        = "playerStoppingHorizontal"
	KSpriteStatusPlayerStoppingHorizontal           = "KSpriteStatusPlayerStoppingHorizontal"
	KSpriteStatusPlayerStoppingHorizontalConfigured = "KSpriteStatusPlayerStoppingHorizontalConfigured"
	KSpriteStatusPlayerRunningVertical              = "KSpriteStatusPlayerRunningVertical"
	KSpriteStatusPlayerStoppingVertical             = "KSpriteStatusPlayerStoppingVertical"
	SpriteStatusFreeFallStart                       = "SpriteStatusFreeFallStart"
	SpriteStatusFreeFall                            = "SpriteStatusFreeFall"
	KSpriteStatusFreeFallImpact                     = "KSpriteStatusFreeFallImpact"

	KSpriteStatusJumpingStart      = "KSpriteStatusJumpingStart"
	KSpriteStatusJumpingStop       = "KSpriteStatusJumpingStop"
	KSpriteStatusJumpingInProgress = "KSpriteStatusJumpingInProgress"
	KSpriteStatusJumpingEnable     = "KSpriteStatusJumpingEnable"
	KSpriteStatusJumpingDisable    = "KSpriteStatusJumpingDisable"

	// KSpriteStatusMovieClipStop
	//
	// Indica quando o movie clip do player parado deve ser usado pela configuração determinada pelo desenvolvedor.
	// e.runningStopSlip == true: Define o uso do movie clip de parado mesmo quando a velocidade de x difere de zero
	// e.runningStopSlip == false: Define o uso do movie clip de parado apenas quando a velocidade de x for zero
	KSpriteStatusMovieClipStop = "KSpriteStatusMovieClipStop"
)

type AccelerationFunction func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector, xVectorDesired DirectionHorizontal, yVector, yVectorDesired DirectionVertical) (dx, dy float64)

type DirectionHorizontal bool

func (e DirectionHorizontal) String() string {
	if e == KRight {
		return "right"
	}

	return "left"
}

const (
	KRight DirectionHorizontal = true
	KLeft  DirectionHorizontal = false
)

type DirectionVertical bool

func (e DirectionVertical) String() string {
	if e == KUp {
		return "up"
	}

	return "down"
}

const (
	KUp   DirectionVertical = true
	KDown DirectionVertical = false
)

type SpritePlayer struct {
	spt   *Sprite
	stage stage.Functions

	status     eventQueue.EventQueue
	statusPrev map[string]bool

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

	freeFallRegistered bool
	freeFallStart      time.Time

	jumpingEnabled bool
	jumpingStart   time.Time
	jumpingStop    time.Time

	limitXMin float64
	limitXMax float64
	limitYMin float64
	limitYMax float64

	horizontalTmpFunction AccelerationFunction
	horizontalFunction    AccelerationFunction
	verticalFunction      AccelerationFunction

	runningLeftStartFunction   AccelerationFunction
	runningLeftStopFunction    AccelerationFunction
	runningRightStartFunction  AccelerationFunction
	runningRightStopFunction   AccelerationFunction
	freeFallFunction           AccelerationFunction
	gravityFunctionAirFriction AccelerationFunction

	onRunningRightStopConfigure              bool
	onRunningLeftStopConfigure               bool
	onRunningRightConfigBeforeStartConfigure bool
	onRunningLeftConfigBeforeStartConfigure  bool
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
// Define a fórmula usada para calcular a gravidade.
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
	e.freeFallFunction = f

	return e
}

func (e *SpritePlayer) GravityAirFrictionFunc(f AccelerationFunction) (ref *SpritePlayer) {
	e.gravityFunctionAirFriction = f

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
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector, xVectorDesired DirectionHorizontal, yVector, yVectorDesired DirectionVertical) (dx, dy float64) {
		// English: if the player is moving in the opposite direction, it moves less. more friction with the floor.
		// This rule is here to facilitate the exchange of formulas.
		// Português: Caso o player esteja mudando para direção oposta, ele se desloca menos. mais atrito com o chão.
		// Esta regra fica aqui para facilitar a troca de fórmulas.
		if xVector != xVectorDesired {
			inertialSpeedX /= 2
		}

		dx = inertialSpeedX + 0.000001*float64(tRunning*tRunning)

		if xVectorDesired == KLeft {
			dx *= -1.0
		}
		if dx > 2.0 {
			dx = 2.0
		} else if dx < -2.0 {
			dx = -2.0
		}

		return
	}
}

func (e *SpritePlayer) GetFormulaFloorVerySlipperyStop() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector, xVectorDesired DirectionHorizontal, yVector, yVectorDesired DirectionVertical) (dx, dy float64) {
		dx = inertialSpeedX - 0.0000007*float64(tRunning*tRunning)
		if xVectorDesired == KLeft {
			dx *= -1.0
			if dx > 0 {
				dx = 0
			}
		} else if dx < 0 {
			dx = 0
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorLittleSlipperyStart() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector, xVectorDesired DirectionHorizontal, yVector, yVectorDesired DirectionVertical) (dx, dy float64) {
		dx = 0.000002 * float64(tRunning*tRunning)

		// English: if the player is moving in the opposite direction, it moves less. more friction with the floor.
		// This rule is here to facilitate the exchange of formulas.
		// Português: Caso o player esteja mudando para direção oposta, ele se desloca menos. mais atrito com o chão.
		// Esta regra fica aqui para facilitar a troca de fórmulas.
		if xVector != xVectorDesired {
			inertialSpeedX /= 2
		}

		dx = inertialSpeedX + 0.000002*float64(tRunning*tRunning)

		if xVectorDesired == KLeft {
			dx *= -1.0
		}
		if dx > 2.0 {
			dx = 2.0
		} else if dx < -2.0 {
			dx = -2.0
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorLittleSlipperyStop() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector, xVectorDesired DirectionHorizontal, yVector, yVectorDesired DirectionVertical) (dx, dy float64) {
		dx = inertialSpeedX - 0.000005*float64(tRunning*tRunning)
		if xVectorDesired == KLeft {
			dx *= -1.0
			if dx > 0 {
				dx = 0
			}
		} else if dx < 0 {
			dx = 0
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorDefaultStart() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector, xVectorDesired DirectionHorizontal, yVector, yVectorDesired DirectionVertical) (dx, dy float64) {

		// English: When the player is moving in the opposite direction, not using inertialSpeedX makes the movement more pleasant
		// Português: Quando o player está se movendo na direção oposta, não usar inertialSpeedX deixa o movimento mais agradável
		dx = 0.000008*float64(tRunning*tRunning) + 0.5

		if xVectorDesired == KLeft {
			dx *= -1.0
		}
		if dx > 2.0 {
			dx = 2.0
		} else if dx < -2.0 {
			dx = -2.0
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaFloorDefaultStop() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector, xVectorDesired DirectionHorizontal, yVector, yVectorDesired DirectionVertical) (dx, dy float64) {

		dx = (inertialSpeedX / 8) - 0.00002*float64(tRunning*tRunning)
		dx += 0.5
		if xVectorDesired == KLeft {
			dx *= -1.0
			if dx > 0 {
				dx = 0
			}
		} else if dx < 0 {
			dx = 0
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaGravityDefault() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector, xVectorDesired DirectionHorizontal, yVector, yVectorDesired DirectionVertical) (dx, dy float64) {
		dy = 0.0000005*float64(tRunning*tRunning) + 0.5
		if dy > 20 {
			dy = 20
		}
		return
	}
}

func (e *SpritePlayer) GetFormulaAirFriction() (f AccelerationFunction) {
	return func(inertialSpeedX, inertialSpeedY, velocityX, velocityY, x, y float64, tRunning, tGravity int64, xVector, xVectorDesired DirectionHorizontal, yVector, yVectorDesired DirectionVertical) (dx, dy float64) {
		dx = inertialSpeedX - (0.0000002*float64(tRunning*tRunning) + 0.0)
		if dx < 0 {
			dx = 0
		}
		return
	}
}

func (e *SpritePlayer) Init(stage stage.Functions, canvas *TagCanvas, imgPath string, width, height int) (ref *SpritePlayer) {
	e.stage = stage

	e.status.Init()
	e.status.AddOppositeEvent(moveRight, moveLeft)
	e.status.AddOppositeEvent(moveRightStop, moveLeftStop)

	e.statusPrev = make(map[string]bool)

	e.freeFallStart = time.Now()
	e.runningStart = time.Now()

	e.spt = new(Sprite)
	e.spt.Canvas(canvas)
	e.spt.Image(imgPath)
	e.spt.SpriteWidth(width)
	e.spt.SpriteHeight(height)
	e.spt.Init()

	e.freeFallFunction = e.GetFormulaGravityDefault()
	e.gravityFunctionAirFriction = e.GetFormulaAirFriction()
	e.runningRightStartFunction = e.GetFormulaFloorDefaultStart()
	e.runningLeftStartFunction = e.GetFormulaFloorDefaultStart()
	e.runningRightStopFunction = e.GetFormulaFloorDefaultStop()
	e.runningLeftStopFunction = e.GetFormulaFloorDefaultStop()

	// limita direita/esquerda
	e.limitXMin = -1.0
	e.limitXMax = -1.0

	e.limitYMin = -1.0
	e.limitYMax = -1.0

	e.xVector = KRight
	e.xVectorDesired = KRight

	e.MovieClipStopped()

	e.stage.AddMathFunctions(e.statusVerify)

	return e
}

func (e *SpritePlayer) statusVerify() {

	_, empty, list := e.status.GetStatus()

	if empty == true {
		return
	}

	var activeList = make(map[string]bool)
	for _, v := range list {
		activeList[v.GetLabel()] = true
	}

	//e.status.Debug()

	if activeList[SpriteStatusFreeFallStart] {
		e.status.AddEvent(SpriteStatusFreeFallStart, false)
		e.status.AddEvent(SpriteStatusFreeFall, true)
		e.onFreeFall()
	}

	if activeList[SpriteStatusFreeFall] {
		if e.verticalFunction != nil {
			tRunning := time.Since(e.runningStart).Milliseconds()
			tGravity := time.Since(e.freeFallStart).Milliseconds()
			e.velocityX, e.velocityY = e.verticalFunction(e.velocityXInertial, e.velocityYInertial, e.velocityX, e.velocityY, e.x, e.y, tRunning, tGravity, e.xVector, e.xVectorDesired, e.yVector, e.yVectorDesired)

			if e.yVectorDesired == KUp {
				e.DY(-e.velocityY)
			} else {
				e.DY(e.velocityY)
			}
		} else {
			log.Print("bug: vertical function is nil")
		}
	}

	if activeList[moveRight] && !activeList[moveRightConfig] {
		e.status.AddEvent(moveRightConfig, true)
		e.status.AddEvent(playerRunningHorizontal, true)

		// English: This block configures when right movement starts
		// Português: Este bloco configura quando o movimento para a direita começa
		e.velocityXInertial = e.velocityX
		e.xVectorDesired = KRight
		e.horizontalFunction = e.runningRightStartFunction
		e.runningStart = time.Now()
		e.MovieClipRunning()
	}

	if activeList[moveRightConfig] && !activeList[moveRight] {
		delete(activeList, moveRightConfig)
		e.status.AddEvent(moveRightConfig, false)

		// English: This block removes the configuration flag
		// Português: Este bloco remove o flag indicativo de configuração
	}

	if activeList[moveLeft] && !activeList[moveLeftConfig] {
		e.status.AddEvent(moveLeftConfig, true)
		e.status.AddEvent(playerRunningHorizontal, true)

		// English: This block configures when left movement starts
		// Português: Este bloco configura quando o movimento para a esquerda começa
		e.velocityXInertial = -e.velocityX
		e.xVectorDesired = KLeft
		e.horizontalFunction = e.runningLeftStartFunction
		e.runningStart = time.Now()
		e.MovieClipRunning()
	}

	if activeList[moveLeftConfig] && !activeList[moveLeft] {
		delete(activeList, moveLeftConfig)
		e.status.AddEvent(moveLeftConfig, false)

		// English: This block removes the configuration flag
		// Português: Este bloco remove o flag indicativo de configuração
	}

	if e.statusPrev[moveLeft] && activeList[moveRight] {
		// English: This block is triggered when the movement changes from left to right
		// Português: Este bloco é acionado quando o movimento muda da esquerda para a direita

	} else if e.statusPrev[moveRight] && activeList[moveLeft] {
		// English: This block is triggered when the movement changes from right to left
		// Português: Este bloco é acionado quando o movimento muda da direita para a esquerda

	} else if e.statusPrev[moveLeft] && activeList[moveLeftStop] {
		// English: This block is triggered when the movement changes from walking left to stopping left
		// Português: Este bloco é acionado quando o movimento muda de andar para a esquerda para parar a esquerda

		e.horizontalFunction = e.runningRightStopFunction
		e.velocityXInertial = -e.velocityX
		e.runningStart = time.Now()
	} else if e.statusPrev[moveRight] && activeList[moveRightStop] {
		// English: This block is triggered when the movement changes from walking right to stopping right
		// Português: Este bloco é acionado quando o movimento muda de andar para a direita para parar a direita

		e.horizontalFunction = e.runningRightStopFunction
		e.velocityXInertial = e.velocityX
		e.runningStart = time.Now()
	} else if e.statusPrev[moveLeftStop] && activeList[moveLeft] {
		// English: This block is triggered when movement changes from stopping left to walking left
		// Português: Este bloco é acionado quando o movimento muda de parando em direção a esquerda para andando em direção a esquerda

	} else if e.statusPrev[playerStoppingHorizontal] && activeList[moveRight] && !activeList[moveRightConfig] {
		// English: This block is triggered when movement changes from stopping right to walking right
		// Português: Este bloco é acionado quando o movimento muda de parando em direção a direita para andando em direção a direita

	} else if e.statusPrev[playerStoppingHorizontal] && activeList[moveRight] && !activeList[moveRightConfig] {
		// English: This block is triggered when movement changes from stopping left to walking right
		// Português: Este bloco é acionado quando o movimento muda de parando em direção a esquerda para andando em direção a direita

	} else if e.statusPrev[moveRightStop] && activeList[moveLeft] {
		// English: This block is triggered when movement changes from stopping right to walking left
		// Português: Este bloco é acionado quando o movimento muda de parando em direção a esquerda para andando em direção a esquerda

	}

	if (activeList[moveRight] || activeList[moveLeft]) && activeList[KSpriteStatusMovieClipStop] {
		delete(activeList, KSpriteStatusMovieClipStop)

		// English: This block is triggered when the player is moving with the movie clip stopped and user triggers movement again (slippery floor)
		// Português: Este bloco é acionado quando o player está se deslocando com o movie clip de parado e usuário aciona movimento novamente (chão escorregadio)
		e.status.AddEvent(KSpriteStatusMovieClipStop, false)
	}

	if !activeList[moveRight] && !activeList[moveLeft] &&
		activeList[playerStoppingHorizontal] {

		// English: This block is triggered when right/left movement is not triggered
		// Português: Este bloco é acionado quando o movimento de direita/esquerda não é acionado

		// English: The stopped movie clip makes the player slip on the floor
		// Português: O movie clip de parado faz o player escorregar no chão
		if e.runningStopSlip == true && !activeList[KSpriteStatusMovieClipStop] {
			e.status.AddEvent(KSpriteStatusMovieClipStop, true)
		}
	}

	if activeList[moveRightStop] {
		delete(activeList, moveRight)

		// English: This block is triggered by the command to stop the movement to the right
		// Português: Este bloco é acionado pelo comando de parar o movimento para à direita
		e.status.AddEvent(playerStoppingHorizontal, true)
		e.status.AddEvent(moveRight, false)
		e.status.AddEvent(moveRightStop, false)
	}

	if activeList[moveLeftStop] {
		delete(activeList, moveLeft)

		// English: This block is triggered by the command to stop the movement to the left
		// Português: Este bloco é acionado pelo comando de parar o movimento para à esquerda
		e.status.AddEvent(playerStoppingHorizontal, true)
		e.status.AddEvent(moveLeft, false)
		e.status.AddEvent(moveLeftStop, false)
	}

	if activeList[playerStoppingHorizontal] {
		// English: This block is triggered by the stop horizontal movement command
		// Português: Este bloco é acionado pelo comando de parar o movimento horizontal

		if e.velocityX == 0 {

			e.status.AddEvent(playerRunningHorizontal, false)
			e.status.AddEvent(playerStoppingHorizontal, false)

			if e.runningStopSlip == false {
				e.status.AddEvent(KSpriteStatusMovieClipStop, true)
			}
		}
	}

	if activeList[playerRunningHorizontal] {
		// English: This block is triggered by the horizontal movement command and triggers the acceleration formula, both for running and stopping.
		// Português: Este bloco é acionado pelo comando de movimento horizontal e aciona a fórmula de aceleração, tanto para correr quanto parar

		if e.horizontalFunction != nil {
			tRunning := time.Since(e.runningStart).Milliseconds()
			tGravity := time.Since(e.freeFallStart).Milliseconds()
			e.velocityX, e.velocityY = e.horizontalFunction(e.velocityXInertial, e.velocityYInertial, e.velocityX, e.velocityY, e.x, e.y, tRunning, tGravity, e.xVector, e.xVectorDesired, e.yVector, e.yVectorDesired)
			e.DX(e.velocityX)
		} else {
			log.Print("bug: horizontal function is nil")
		}
	}

	if activeList[KSpriteStatusMovieClipStop] {

		// English: This block triggers the stopped movie clip
		// Português: Este bloco aciona o movie clip parado
		e.status.AddEvent(KSpriteStatusMovieClipStop, false)
		e.MovieClipStopped()
	}

	e.statusPrev = make(map[string]bool)
	for k, v := range activeList {
		e.statusPrev[k] = v
	}
}

func (e *SpritePlayer) MovieClipStopped() (ref *SpritePlayer) {
	if e.xVectorDesired == KRight {
		e.MovieClipStoppedRightSide()
	} else {
		e.MovieClipStoppedLeftSide()
	}

	return e
}

func (e *SpritePlayer) MovieClipRunning() (ref *SpritePlayer) {
	if e.xVectorDesired == KRight {
		e.MovieClipWalkingRightSide()
	} else {
		e.MovieClipWalkingLeftSide()
	}

	return e
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

func (e *SpritePlayer) MovieClipStoppedRightSide() (ref *SpritePlayer) {
	err := e.spt.Start("stoppedRightSide")
	if err != nil {
		log.Printf("bug: StartStoppedRightSide()")
	}

	return e
}

func (e *SpritePlayer) CreateStoppedLeftSide() (ref *spriteConfig) {
	return e.spt.Scene("stoppedLeftSide")
}

func (e *SpritePlayer) MovieClipStoppedLeftSide() (ref *SpritePlayer) {
	err := e.spt.Start("stoppedLeftSide")
	if err != nil {
		log.Printf("bug: StartStoppedLeftSide()")
	}

	return e
}

func (e *SpritePlayer) CreateWalkingRightSide() (ref *spriteConfig) {
	return e.spt.Scene("walkingRightSide")
}

func (e *SpritePlayer) MovieClipWalkingRightSide() (ref *SpritePlayer) {
	err := e.spt.Start("walkingRightSide")
	if err != nil {
		log.Printf("bug: StartWalkingRightSide()")
	}

	return e
}

func (e *SpritePlayer) CreateWalkingLeftSide() (ref *spriteConfig) {
	return e.spt.Scene("walkingLeftSide")
}

func (e *SpritePlayer) MovieClipWalkingLeftSide() (ref *SpritePlayer) {
	err := e.spt.Start("walkingLeftSide")
	if err != nil {
		log.Printf("bug: StartWalkingLeftSide()")
	}

	return e
}

func (e *SpritePlayer) CreateFallRightSide() (ref *spriteConfig) {
	return e.spt.Scene("fallRightSide")
}

func (e *SpritePlayer) MovieClipFall() (ref *SpritePlayer) {
	if e.xVectorDesired == KRight {
		e.MovieClipFallRightSide()
	} else {
		e.MovieClipFallLeftSide()
	}

	return e
}

func (e *SpritePlayer) MovieClipFallRightSide() (ref *SpritePlayer) {
	err := e.spt.Start("fallRightSide")
	if err != nil {
		log.Printf("bug: StartFallRightSide()")
	}

	return e
}

func (e *SpritePlayer) CreateFallLeftSide() (ref *spriteConfig) {
	return e.spt.Scene("fallLeftSide")
}

func (e *SpritePlayer) MovieClipFallLeftSide() (ref *SpritePlayer) {
	err := e.spt.Start("fallLeftSide")
	if err != nil {
		log.Printf("bug: StartFallLeftSide()")
	}

	return e
}

func (e *SpritePlayer) RunningRightStop() {
	e.status.AddEvent(
		moveRightStop,
		true,
	)
}

func (e *SpritePlayer) RunningLeftStop() {
	e.status.AddEvent(
		moveLeftStop,
		true,
	)
}

func (e *SpritePlayer) RunningRightStart() {
	e.status.AddEvent(
		moveRight,
		true,
	)
}

func (e *SpritePlayer) RunningLeftStart() {
	e.status.AddEvent(
		moveLeft,
		true,
	)
}

func (e *SpritePlayer) JumpingStart() {
	e.status.AddEvent(
		KSpriteStatusJumpingStart,
		true,
	)
}

func (e *SpritePlayer) JumpingStop() {
	e.status.AddEvent(
		KSpriteStatusJumpingStop,
		true,
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

func (e *SpritePlayer) onFreeFall() {
	e.yVectorDesired = KDown

	if e.yVector == KUp {
		e.velocityYInertial = -e.velocityY
	} else {
		e.velocityYInertial = e.velocityY
	}

	//if e.velocityX < 0 {
	//	e.velocityXInertial = -e.velocityX
	//} else {
	//	e.velocityXInertial = e.velocityX
	//}

	e.verticalFunction = e.freeFallFunction
	e.freeFallStart = time.Now()
}

func (e *SpritePlayer) FreeFallEnable() {
	if e.freeFallRegistered == true {
		return
	}
	e.freeFallRegistered = true

	e.status.AddEvent(
		SpriteStatusFreeFallStart,
		true,
	)
}

func (e *SpritePlayer) FreeFallDisable() {
	if e.freeFallRegistered == false {
		return
	}
	e.freeFallRegistered = false

	e.status.AddEvent(
		SpriteStatusFreeFall,
		false,
	)
}

func (e *SpritePlayer) JumpingEnable() {
	if e.jumpingEnabled == true {
		return
	}
	e.jumpingEnabled = true

	// fixme: descomentar
	//e.status.AddEvent(
	//	KSpriteStatusJumpingEnable,
	//	true,
	//)
}

func (e *SpritePlayer) JumpingDisable() {
	if e.jumpingEnabled == false {
		return
	}
	e.jumpingEnabled = false

	e.status.AddEvent(
		KSpriteStatusJumpingDisable,
		true,
	)
}

func (e *SpritePlayer) GetCollisionBox() (box Box) {
	return e.spt.GetCollisionBox()
}
