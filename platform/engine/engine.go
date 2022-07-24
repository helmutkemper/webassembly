package engine

import (
	"math"
	"math/rand"
	"time"
)

// kUIdSize
//
// English:
//
// Defines the size of the UID linked to each added role.
//
// Português:
//
// Define o tamanho do UID vinculado a cada função adicionada.
const kUIdSize = 10

// kFpsMin
//
// English:
//
// Sets the minimum amount of fps allowed for automatic engine tuning
//
// Português:
//
// Define a quantidade mínima de fps permitida para o ajuste automático da engine
const kFpsMin = 10

// kFpsMax
//
// English:
//
// Sets the maximum amount of fps allowed for automatic engine tuning
//
//   Notes:
//    * Setting high fps crashes the browser in a test done on Mac M1 with 16GB of RAM
//
// Português:
//
// Define a quantidade máxima de fps permitida para o ajuste automático da engine
//
//   Notas:
//    * Definir fps elevados travam o navegador em um teste feito no Mac M1 com 16GB de RAM
const kFpsMax = 120

type FuncList struct {
	id string
	f  func()
}

type Engine struct {
	sleepFrame    int
	fps           int
	fpsMin        int
	fpsMax        int
	fpsLowLatency int
	kUIdCharList  []string

	// todo: fazer um channel
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
	if el.fpsMin == 0 {
		el.fpsMin = kFpsMin
	}

	if el.fpsMax == 0 {
		el.fpsMax = kFpsMax
	}

	// fixme: must be a interval of time
	el.sleepFrame = 0
	el.fpsLowLatency = 10

	if el.fps == 0 {
		el.fps = el.fpsMax
	}

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

// SetSleepFrame
//
// English:
//
// Português:
//
// Sleep Frame pode dá um tempo no processamento para impedir travamentos.
//
// Esta funcionalidade é experimental e vem do C++, onde o laço pode travar o sistema.
func (el *Engine) SetSleepFrame(value int) {
	el.sleepFrame = value
}

// GetSleepFrame
//
// English:
//
// Português:
//
// Sleep Frame pode dá um tempo no processamento para impedir travamentos.
//
// Esta funcionalidade é experimental e vem do C++, onde o laço pode travar o sistema.
func (el *Engine) GetSleepFrame() int {
	return el.sleepFrame
}

// SetFpsMax
//
// English:
//
// Português:
//
// Define a quantidade máxima de FPS.
//
//   Notas:
//     * A quantidade máxima de FPS pode fazer o navegador travar.
func (el *Engine) SetFpsMax(value int) {
	el.fpsMax = value
}

// SetFpsMin
//
// English:
//
// Português:
//
// Define a quantidade mínima de FPS.
//
//   Notas:
//     * O valor do FPS cai de forma automática quando o sistema está muito ocupado.
func (el *Engine) SetFpsMin(value int) {
	el.fpsMin = value
}

// SetFPS
//
// English:
//
// Português:
//
// Define a quantidade de FPS
func (el *Engine) SetFPS(value int) {
	el.fps = value
	el.stopTicker = true
}

// GetFPS
//
// English:
//
// Returns the amount of current FPS used in calculations.
//
// Português:
//
// Retorna a quantidade de FPS atual usado nos cálculos.
func (el *Engine) GetFPS() int {
	return el.fps
}

// CursorAddDrawFunction
//
// English:
//
// Português:
//
// Permite recriar a função que desenha o cursor.
func (el *Engine) CursorAddDrawFunction(runnerFunc func()) string {
	UId := el.getUId()
	el.funcCursorDraw = FuncList{id: UId, f: runnerFunc}

	return UId
}

// CursorRemoveDrawFunction
//
// English:
//
// Português:
//
// Remove a função que recria o cursor.
func (el *Engine) CursorRemoveDrawFunction(id string) {
	el.funcCursorDraw = FuncList{}
}

// HighLatencyAddToFunctions
//
// English:
//
// Português:
//
// Adiciona uma função de alta latencia, uma função de baixa prioridade de execussão.
//
//   Entrada:
//     runnerFunc: função a ser executada.
//
//   Saída:
//     UId da função, usado para identificar a função na hora de remover.
//     total: quantidade total de funções em execução.
//
//   Notas:
//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
func (el *Engine) HighLatencyAddToFunctions(runnerFunc func()) (string, int) {
	UId := el.getUId()
	index := len(el.funcListToHighLatency)
	el.funcListToHighLatency = append(el.funcListToHighLatency, FuncList{id: UId, f: runnerFunc})

	return UId, index
}

// HighLatencyDeleteFromFunctions
//
// English:
//
// Português:
//
// Remove uma função de alta latencia adicionada pela função HighLatencyAddToFunctions().
//
//   Entrada:
//     UId: ID retornado pela função HighLatencyAddToFunctions().
//
//   Notas:
//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
func (el *Engine) HighLatencyDeleteFromFunctions(UId string) {
	for k, runner := range el.funcListToHighLatency {
		if runner.id == UId {
			el.funcListToHighLatency = append(el.funcListToHighLatency[:k], el.funcListToHighLatency[k+1:]...)
			break
		}
	}
}

// HighLatencySetZIndex
//
// English:
//
// Português:
//
// Permite trocar a ordem de execução da função, na lista de execução.
//
//   Entrada:
//     UId: ID retornado pela função HighLatencyAddToFunctions().
//     index: 0 para a primeira função da lista
//
//   Notas:
//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
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

// HighLatencyGetZIndex
//
// English:
//
// Português:
//
// Retorna o índice de execução da função na lista, onde 0 é a primera função a ser executada.
//
//   Entrada:
//     UId: ID retornado pela função HighLatencyAddToFunctions().
//
//   Saída:
//     index: Ordem de execução da função.
//
//   Notas:
//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
func (el *Engine) HighLatencyGetZIndex(UId string) int {
	for k, runner := range el.funcListToHighLatency {
		if runner.id == UId {
			return k
		}
	}

	return math.MaxInt32
}

// HighLatencySetAsFistFunctionToRun
//
// English:
//
// Português:
//
// Faz a função ser a primeira da lista de execuções para funções de alta latencia.
//
//   Entrada:
//     UId: ID retornado pela função HighLatencyAddToFunctions().
//
//   Saída:
//     index: Ordem de execução da função.
//
//   Notas:
//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
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

// HighLatencySetAsLastFunctionToRun
//
// English:
//
// Português:
//
// Faz a função ser a útima da lista de execuções para funções de alta latencia.
//
//   Entrada:
//     UId: ID retornado pela função HighLatencyAddToFunctions().
//
//   Saída:
//     index: Ordem de execução da função.
//
//   Notas:
//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
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

// SystemAddToFunctions
//
// English:
//
// Português:
//
// Adiciona uma função a lista de execuções.
//
//   Entrada:
//     runnerFunc: função a ser executada.
//
//   Saída:
//     UId da função, usado para identificar a função na hora de remover.
//     total: quantidade total de funções em execução.
//
//   Notas:
//     * Funções de sistema são as primeiras funções da lista de execuções e devem ser as funções de uso do sistema.
func (el *Engine) SystemAddToFunctions(runnerFunc func()) (string, int) {
	UId := el.getUId()
	index := len(el.funcListToSystem)
	el.funcListToSystem = append(el.funcListToSystem, FuncList{id: UId, f: runnerFunc})

	return UId, index
}

// SystemDeleteFromFunctions
//
// English:
//
// Português:
//
// Remove uma função da lista de funções do sistema, adicionada pela função SystemAddToFunctions().
//
//   Entrada:
//     UId: ID retornado pela função SystemAddToFunctions().
//
//   Notas:
//     * Funções de sistema são as primeiras funções da lista de execuções e devem ser as funções de uso do sistema.
func (el *Engine) SystemDeleteFromFunctions(UId string) {
	for k, runner := range el.funcListToSystem {
		if runner.id == UId {
			el.funcListToSystem = append(el.funcListToSystem[:k], el.funcListToSystem[k+1:]...)
			break
		}
	}
}

// SystemSetZIndex
//
// English:
//
// Português:
//
// Permite trocar a ordem de execução da função, na lista de execução.
//
//   Entrada:
//     UId: ID retornado pela função SystemAddToFunctions().
//     index: 0 para a primeira função da lista
//
//   Notas:
//     * Funções de sistema são as primeiras funções da lista de execuções e devem ser as funções de uso do sistema.
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

// SystemGetZIndex
//
// English:
//
// Português:
//
// Retorna o índice de execução da função na lista, onde 0 é a primera função a ser executada.
//
//   Entrada:
//     UId: ID retornado pela função SystemAddToFunctions().
//
//   Saída:
//     index: Ordem de execução da função.
func (el *Engine) SystemGetZIndex(UId string) int {
	for k, runner := range el.funcListToSystem {
		if runner.id == UId {
			return k
		}
	}

	return math.MaxInt32
}

//
//
// English:
//
// Português:
//
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

//
//
// English:
//
// Português:
//
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

//
//
// English:
//
// Português:
//
func (el *Engine) AfterSystemAddToFunctions(runnerFunc func()) (string, int) {
	UId := el.getUId()
	index := len(el.funcListToAfterSystem)
	el.funcListToAfterSystem = append(el.funcListToAfterSystem, FuncList{id: UId, f: runnerFunc})

	return UId, index
}

//
//
// English:
//
// Português:
//
func (el *Engine) AfterSystemDeleteFromFunctions(UId string) {
	for k, runner := range el.funcListToAfterSystem {
		if runner.id == UId {
			el.funcListToAfterSystem = append(el.funcListToAfterSystem[:k], el.funcListToAfterSystem[k+1:]...)
			break
		}
	}
}

//
//
// English:
//
// Português:
//
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

//
//
// English:
//
// Português:
//
func (el *Engine) AfterSystemGetZIndex(UId string) int {
	for k, runner := range el.funcListToAfterSystem {
		if runner.id == UId {
			return k
		}
	}

	return math.MaxInt32
}

//
//
// English:
//
// Português:
//
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

//
//
// English:
//
// Português:
//
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

//
//
// English:
//
// Português:
//
func (el *Engine) MathAddToFunctions(runnerFunc func()) (string, int) {
	UId := el.getUId()
	index := len(el.funcListToMath)
	el.funcListToMath = append(el.funcListToMath, FuncList{id: UId, f: runnerFunc})

	return UId, index
}

//
//
// English:
//
// Português:
//
func (el *Engine) MathDeleteFromFunctions(UId string) {
	for k, runner := range el.funcListToMath {
		if runner.id == UId {
			el.funcListToMath = append(el.funcListToMath[:k], el.funcListToMath[k+1:]...)
			break
		}
	}
}

//
//
// English:
//
// Português:
//
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

//
//
// English:
//
// Português:
//
func (el *Engine) MathGetZIndex(UId string) int {
	for k, runner := range el.funcListToMath {
		if runner.id == UId {
			return k
		}
	}

	return math.MaxInt32
}

//
//
// English:
//
// Português:
//
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

//
//
// English:
//
// Português:
//
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

//
//
// English:
//
// Português:
//
func (el *Engine) DrawAddToFunctions(runnerFunc func()) (string, int) {
	UId := el.getUId()
	index := len(el.funcListToDraw)
	el.funcListToDraw = append(el.funcListToDraw, FuncList{id: UId, f: runnerFunc})

	return UId, index
}

//
//
// English:
//
// Português:
//
func (el *Engine) DrawDeleteFromFunctions(UId string) {
	for k, runner := range el.funcListToDraw {
		if runner.id == UId {
			el.funcListToDraw = append(el.funcListToDraw[:k], el.funcListToDraw[k+1:]...)
			break
		}
	}
}

//
//
// English:
//
// Português:
//
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

//
//
// English:
//
// Português:
//
func (el *Engine) DrawGetZIndex(UId string) int {
	for k, runner := range el.funcListToDraw {
		if runner.id == UId {
			return k
		}
	}

	return math.MaxInt32
}

//
//
// English:
//
// Português:
//
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

//
//
// English:
//
// Português:
//
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

//
//
// English:
//
// Português:
//
// todo: uID deveria ser algo isolado
func (el *Engine) getUId() string {
	var UId = ""
	for i := 0; i != kUIdSize; i += 1 {
		UId += el.kUIdCharList[rand.Intn(len(el.kUIdCharList)-1)]
	}

	return UId
}

//
//
// English:
//
// Português:
//
func (el *Engine) tickerStart() {
	el.ticker = time.NewTicker(time.Second / time.Duration(el.fps))
	el.tickerLowLatency = time.NewTicker(time.Second / time.Duration(el.fpsLowLatency))
	el.tickerVerifyFps = time.NewTicker(1 * time.Second) //todo: constante ou configurar
	el.slipFrameTimeAlarm = time.Second / time.Duration(el.fps)
	// fixme: remover a função desnecessária
	go func() { el.tickerRunner() }()
}

//
//
// English:
//
// Português:
//
func (el *Engine) tickerRunner() {
	defer el.tickerStart()
	for {
		select {
		case <-el.tickerVerifyFps.C:
			if el.conterOverflow >= 3 {
				el.fps -= el.conterOverflow
				if el.fps < el.fpsMin {
					el.fps = el.fpsMin
				}
			} else {
				el.fps += 2
				if el.fps > el.fpsMax {
					el.fps = el.fpsMax
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
