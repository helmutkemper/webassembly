package engine

import (
	"math"
	"math/rand"
	"time"
)

const kUIdSize = 10

type FuncList struct {
	id string
	f  func()
}

type Engine struct {
	sleepFrame    int
	fps           int
	fpsLowLatency int
	kUIdCharList  []string

	// en: Warning! stopTicker should be a channel, however, conflict with webassembly <-done channer
	// pt_br: Cuidado! stopTicker deveria ser um channel, porém, deu conflito com o webassembly <-done channer
	stopTicker bool

	ticker           *time.Ticker
	tickerLowLatency *time.Ticker
	tickerVerifyFps  *time.Ticker

	funcListToHighLatency []FuncList
	funcListToSystem      []FuncList
	funcListToAfterSystem []FuncList
	funcListToMath        []FuncList
	funcListToDraw        []FuncList

	funcCursorDraw FuncList

	// pt_br: impede que o laço ocorra em intervalos muitos próximos e trave o
	// processamento do browser para outras tarefas
	slipFrame          int
	slipFrameTimeAlarm time.Duration

	// Contador de engine sobrecarregada. Não consegue gerar fps configurado.
	conterOverflow int
}

func (el *Engine) Init() {
	// fixme: must be a interval of time
	el.sleepFrame = 2
	el.fps = 120
	el.fpsLowLatency = 1

	el.kUIdCharList = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "_", "!", "@",
		"#", "$", "%", "&", "*", "(", ")", "-", "_", "+", "=", "[", "{", "}", "]", "/", "?", ";", ":", ".", ",", "<", ">",
		"|"}
	el.funcListToSystem = make([]FuncList, 0)
	el.funcListToAfterSystem = make([]FuncList, 0)
	el.funcListToMath = make([]FuncList, 0)
	el.funcListToDraw = make([]FuncList, 0)
	el.tickerStart()
}

func (el *Engine) SetSleepFrame(value int) {
	el.sleepFrame = value
}

func (el *Engine) GetSleepFrame() int {
	return el.sleepFrame
}

func (el *Engine) SetFPS(value int) {
	el.fps = value
	el.stopTicker = true
}

func (el *Engine) GetFPS() int {
	return el.fps
}

func (el *Engine) CursorAddDrawFunction(runnerFunc func()) string {
	UId := el.getUId()
	el.funcCursorDraw = FuncList{id: UId, f: runnerFunc}

	return UId
}

func (el *Engine) CursorRemoveDrawFunction(id string) {
	el.funcCursorDraw = FuncList{}
}

func (el *Engine) HighLatencyAddToFunctions(runnerFunc func()) (string, int) {
	UId := el.getUId()
	index := len(el.funcListToHighLatency)
	el.funcListToHighLatency = append(el.funcListToHighLatency, FuncList{id: UId, f: runnerFunc})

	return UId, index
}

func (el *Engine) HighLatencyDeleteFromFunctions(UId string) {
	for k, runner := range el.funcListToHighLatency {
		if runner.id == UId {
			el.funcListToHighLatency = append(el.funcListToHighLatency[:k], el.funcListToHighLatency[k+1:]...)
			break
		}
	}
}

func (el *Engine) HighLatencySetZIndex(UId string, index int) int {
	var function FuncList
	var pass = false
	var length = len(el.funcListToHighLatency)
	var listCopy = make([]FuncList, len(el.funcListToHighLatency))

	if index < 0 || index > length-1 {
		return math.MaxInt32
	}

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	if index == 0 {

		listCopy = append([]FuncList{function}, listCopy...)

	} else if index == length-1 {

		listCopy = append(listCopy, function)

	} else {

		firstPart := make([]FuncList, len(listCopy[:index]))
		lastPart := make([]FuncList, len(listCopy[index:]))

		copy(firstPart, listCopy[:index])
		copy(lastPart, listCopy[index:])

		firstPart = append(firstPart, function)

		listCopy = make([]FuncList, 0)
		listCopy = append(firstPart, lastPart...)

	}

	el.funcListToHighLatency = listCopy
	return index
}

func (el *Engine) HighLatencyGetZIndex(UId string) int {
	for k, runner := range el.funcListToHighLatency {
		if runner.id == UId {
			return k
		}
	}

	return math.MaxInt32
}

func (el *Engine) HighLatencySetAsFistFunctionToRun(UId string) int {
	var function FuncList
	var pass = false
	var listCopy = make([]FuncList, len(el.funcListToHighLatency))

	copy(listCopy, el.funcListToHighLatency)

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToHighLatency = append([]FuncList{function}, listCopy...)

	return len(el.funcListToHighLatency) - 1
}

func (el *Engine) HighLatencySetAsLastFunctionToRun(UId string) int {
	var function FuncList
	var pass = false
	var listCopy = make([]FuncList, len(el.funcListToHighLatency))

	copy(listCopy, el.funcListToHighLatency)

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToHighLatency = append(listCopy, function)

	return 0
}

func (el *Engine) SystemAddToFunctions(runnerFunc func()) (string, int) {
	UId := el.getUId()
	index := len(el.funcListToSystem)
	el.funcListToSystem = append(el.funcListToSystem, FuncList{id: UId, f: runnerFunc})

	return UId, index
}

func (el *Engine) SystemDeleteFromFunctions(UId string) {
	for k, runner := range el.funcListToSystem {
		if runner.id == UId {
			el.funcListToSystem = append(el.funcListToSystem[:k], el.funcListToSystem[k+1:]...)
			break
		}
	}
}

func (el *Engine) SystemSetZIndex(UId string, index int) int {
	var function FuncList
	var pass = false
	var length = len(el.funcListToSystem)
	var listCopy = make([]FuncList, len(el.funcListToSystem))

	if index < 0 || index > length-1 {
		return math.MaxInt32
	}

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	if index == 0 {

		listCopy = append([]FuncList{function}, listCopy...)

	} else if index == length-1 {

		listCopy = append(listCopy, function)

	} else {

		firstPart := make([]FuncList, len(listCopy[:index]))
		lastPart := make([]FuncList, len(listCopy[index:]))

		copy(firstPart, listCopy[:index])
		copy(lastPart, listCopy[index:])

		firstPart = append(firstPart, function)

		listCopy = make([]FuncList, 0)
		listCopy = append(firstPart, lastPart...)

	}

	el.funcListToSystem = listCopy
	return index
}

func (el *Engine) SystemGetZIndex(UId string) int {
	for k, runner := range el.funcListToSystem {
		if runner.id == UId {
			return k
		}
	}

	return math.MaxInt32
}

func (el *Engine) SystemSetAsFistFunctionToRun(UId string) int {
	var function FuncList
	var pass = false
	var listCopy = make([]FuncList, len(el.funcListToSystem))

	copy(listCopy, el.funcListToSystem)

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToSystem = append([]FuncList{function}, listCopy...)

	return len(el.funcListToSystem) - 1
}

func (el *Engine) SystemSetAsLastFunctionToRun(UId string) int {
	var function FuncList
	var pass = false
	var listCopy = make([]FuncList, len(el.funcListToSystem))

	copy(listCopy, el.funcListToSystem)

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToSystem = append(listCopy, function)

	return 0
}

func (el *Engine) AfterSystemAddToFunctions(runnerFunc func()) (string, int) {
	UId := el.getUId()
	index := len(el.funcListToAfterSystem)
	el.funcListToAfterSystem = append(el.funcListToAfterSystem, FuncList{id: UId, f: runnerFunc})

	return UId, index
}

func (el *Engine) AfterSystemDeleteFromFunctions(UId string) {
	for k, runner := range el.funcListToAfterSystem {
		if runner.id == UId {
			el.funcListToAfterSystem = append(el.funcListToAfterSystem[:k], el.funcListToAfterSystem[k+1:]...)
			break
		}
	}
}

func (el *Engine) AfterSystemSetZIndex(UId string, index int) int {
	var function FuncList
	var pass = false
	var length = len(el.funcListToAfterSystem)
	var listCopy = make([]FuncList, len(el.funcListToAfterSystem))

	if index < 0 || index > length-1 {
		return math.MaxInt32
	}

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	if index == 0 {

		listCopy = append([]FuncList{function}, listCopy...)

	} else if index == length-1 {

		listCopy = append(listCopy, function)

	} else {

		firstPart := make([]FuncList, len(listCopy[:index]))
		lastPart := make([]FuncList, len(listCopy[index:]))

		copy(firstPart, listCopy[:index])
		copy(lastPart, listCopy[index:])

		firstPart = append(firstPart, function)

		listCopy = make([]FuncList, 0)
		listCopy = append(firstPart, lastPart...)

	}

	el.funcListToAfterSystem = listCopy
	return index
}

func (el *Engine) AfterSystemGetZIndex(UId string) int {
	for k, runner := range el.funcListToAfterSystem {
		if runner.id == UId {
			return k
		}
	}

	return math.MaxInt32
}

func (el *Engine) AfterSystemSetAsFistFunctionToRun(UId string) int {
	var function FuncList
	var pass = false
	var listCopy = make([]FuncList, len(el.funcListToAfterSystem))

	copy(listCopy, el.funcListToAfterSystem)

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToAfterSystem = append([]FuncList{function}, listCopy...)

	return len(el.funcListToAfterSystem) - 1
}

func (el *Engine) AfterSystemSetAsLastFunctionToRun(UId string) int {
	var function FuncList
	var pass = false
	var listCopy = make([]FuncList, len(el.funcListToAfterSystem))

	copy(listCopy, el.funcListToAfterSystem)

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToAfterSystem = append(listCopy, function)

	return 0
}

func (el *Engine) MathAddToFunctions(runnerFunc func()) (string, int) {
	UId := el.getUId()
	index := len(el.funcListToMath)
	el.funcListToMath = append(el.funcListToMath, FuncList{id: UId, f: runnerFunc})

	return UId, index
}

func (el *Engine) MathDeleteFromFunctions(UId string) {
	//todo: ela ia ser mais rápido se eu usase o int de MathAddToFunctions()
	// exemplo em tween.tickerRunnerPrepare()

	for k, runner := range el.funcListToMath {
		if runner.id == UId {
			el.funcListToMath = append(el.funcListToMath[:k], el.funcListToMath[k+1:]...)
			break
		}
	}
}

func (el *Engine) MathSetZIndex(UId string, index int) int {
	var function FuncList
	var pass = false
	var length = len(el.funcListToMath)
	var listCopy = make([]FuncList, len(el.funcListToMath))

	if index < 0 || index > length-1 {
		return math.MaxInt32
	}

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	if index == 0 {

		listCopy = append([]FuncList{function}, listCopy...)

	} else if index == length-1 {

		listCopy = append(listCopy, function)

	} else {

		firstPart := make([]FuncList, len(listCopy[:index]))
		lastPart := make([]FuncList, len(listCopy[index:]))

		copy(firstPart, listCopy[:index])
		copy(lastPart, listCopy[index:])

		firstPart = append(firstPart, function)

		listCopy = make([]FuncList, 0)
		listCopy = append(firstPart, lastPart...)

	}

	el.funcListToMath = listCopy
	return index
}

func (el *Engine) MathGetZIndex(UId string) int {
	for k, runner := range el.funcListToMath {
		if runner.id == UId {
			return k
		}
	}

	return math.MaxInt32
}

func (el *Engine) MathSetAsFistFunctionToRun(UId string) int {
	var function FuncList
	var pass = false
	var listCopy = make([]FuncList, len(el.funcListToMath))

	copy(listCopy, el.funcListToMath)

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToMath = append([]FuncList{function}, listCopy...)

	return len(el.funcListToMath) - 1
}

func (el *Engine) MathSetAsLastFunctionToRun(UId string) int {
	var function FuncList
	var pass = false
	var listCopy = make([]FuncList, len(el.funcListToMath))

	copy(listCopy, el.funcListToMath)

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToMath = append(listCopy, function)

	return 0
}

func (el *Engine) DrawAddToFunctions(runnerFunc func()) (string, int) {
	UId := el.getUId()
	index := len(el.funcListToDraw)
	el.funcListToDraw = append(el.funcListToDraw, FuncList{id: UId, f: runnerFunc})

	return UId, index
}

func (el *Engine) DrawDeleteFromFunctions(UId string) {
	for k, runner := range el.funcListToDraw {
		if runner.id == UId {
			el.funcListToDraw = append(el.funcListToDraw[:k], el.funcListToDraw[k+1:]...)
			break
		}
	}
}

func (el *Engine) DrawSetZIndex(UId string, index int) int {
	var function FuncList
	var pass = false
	var length = len(el.funcListToDraw)
	var listCopy = make([]FuncList, len(el.funcListToDraw))

	if index < 0 || index > length-1 {
		return math.MaxInt32
	}

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	if index == 0 {

		listCopy = append([]FuncList{function}, listCopy...)

	} else if index == length-1 {

		listCopy = append(listCopy, function)

	} else {

		firstPart := make([]FuncList, len(listCopy[:index]))
		lastPart := make([]FuncList, len(listCopy[index:]))

		copy(firstPart, listCopy[:index])
		copy(lastPart, listCopy[index:])

		firstPart = append(firstPart, function)

		listCopy = make([]FuncList, 0)
		listCopy = append(firstPart, lastPart...)

	}

	el.funcListToDraw = listCopy
	return index
}

func (el *Engine) DrawGetZIndex(UId string) int {
	for k, runner := range el.funcListToDraw {
		if runner.id == UId {
			return k
		}
	}

	return math.MaxInt32
}

func (el *Engine) DrawSetAsFistFunctionToRun(UId string) int {
	var function FuncList
	var pass = false
	var listCopy = make([]FuncList, len(el.funcListToDraw))

	copy(listCopy, el.funcListToDraw)

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToDraw = append([]FuncList{function}, listCopy...)

	return len(el.funcListToDraw) - 1
}

func (el *Engine) DrawSetAsLastFunctionToRun(UId string) int {
	var function FuncList
	var pass = false
	var listCopy = make([]FuncList, len(el.funcListToDraw))

	copy(listCopy, el.funcListToDraw)

	for k, runner := range listCopy {
		if runner.id == UId {
			pass = true
			function = runner
			listCopy = append(listCopy[:k], listCopy[k+1:]...)
			break
		}
	}

	if pass == false {
		return math.MaxInt32
	}

	el.funcListToDraw = append(listCopy, function)

	return 0
}

// todo: uID deveria ser algo isolado
func (el *Engine) getUId() string {
	var UId = ""
	for i := 0; i != kUIdSize; i += 1 {
		UId += el.kUIdCharList[rand.Intn(len(el.kUIdCharList)-1)]
	}

	return UId
}

func (el *Engine) tickerStart() {
	el.ticker = time.NewTicker(time.Second / time.Duration(el.fps))
	el.tickerLowLatency = time.NewTicker(time.Second / time.Duration(el.fpsLowLatency))
	el.tickerVerifyFps = time.NewTicker(1 * time.Second) //todo: constante ou configurar
	el.slipFrameTimeAlarm = time.Second / time.Duration(el.fps)
	// fixme: remover a função desnecessária
	go func() { el.tickerRunner() }()
}

func (el *Engine) tickerRunner() {
	defer el.tickerStart()
	for {
		select {
		case <-el.tickerVerifyFps.C:
			if el.conterOverflow >= 3 {
				el.fps -= el.conterOverflow
				if el.fps < 10 { //todo: constante ou configurar
					el.fps = 10
				}
			} else {
				el.fps += 2
				if el.fps > 120 { //todo: constante ou configurar
					el.fps = 120
				}
			}
			el.conterOverflow = 0
			//log.Printf("fps: %v", el.fps)

		case <-el.tickerLowLatency.C:

			for _, runnerFunc := range el.funcListToHighLatency {
				if runnerFunc.f != nil {
					runnerFunc.f()
				}
			}

		case <-el.ticker.C:

			el.ticker.Reset(time.Second / time.Duration(el.fps))

			if el.slipFrame != 0 {
				el.slipFrame -= 1
				continue
			}

			if el.stopTicker == true {
				el.stopTicker = false
				return
			}

			start := time.Now()

			for _, runnerFunc := range el.funcListToSystem {
				if runnerFunc.f != nil {
					runnerFunc.f()
				}
			}

			for _, runnerFunc := range el.funcListToAfterSystem {
				if runnerFunc.f != nil {
					runnerFunc.f()
				}
			}

			for _, runnerFunc := range el.funcListToMath {
				if runnerFunc.f != nil {
					runnerFunc.f()
				}
			}

			for _, runnerFunc := range el.funcListToDraw {
				if runnerFunc.f != nil {
					runnerFunc.f()
				}
			}

			if el.funcCursorDraw.f != nil {
				el.funcCursorDraw.f()
			}

			elapsed := time.Since(start)
			if elapsed > el.slipFrameTimeAlarm {
				//fmt.Printf("Esta dando timeout kemper!!\n")
				el.conterOverflow += 1
				el.slipFrame = el.sleepFrame
			}
		}
	}
}
