package easingTween

import (
	"github.com/helmutkemper/iotmaker.webassembly/interfaces"
	"github.com/helmutkemper/iotmaker.webassembly/platform/engine"
	"github.com/helmutkemper/iotmaker.webassembly/platform/globalEngine"
	"time"
)

type Tween struct {
	engine             engine.IEngine
	engineHasFunction  bool
	startValue         float64
	endValue           float64
	arguments          interface{}
	startTime          time.Time
	duration           time.Duration
	tweenFunc          func(currentTime, duration, currentPercentage, startValue, endValue, changeInValue float64) float64
	interaction        func(value, percentToComplete float64, arguments interface{})
	onCycleStart       func(value float64, arguments interface{})
	onCycleEnd         func(value float64, arguments interface{})
	onStart            func(value float64, arguments interface{})
	onEnd              func(value float64, arguments interface{})
	onInvert           func(value float64, arguments interface{})
	doNotReverseMotion bool
	invert             bool
	repeat             int
	fpsUId             string
	loopStartValue     float64
	loopEndValue       float64
}

// SetEngine
//
// English:
//
//  Defines a new engine for time control.
//
//   Input:
//     value: object compatible with the engine.IEEngine interface
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Define uma nova engine para controle de tempo.
//
//   Entrada:
//     value: objeto compatível com ã interface engine.IEngine
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetEngine(value engine.IEngine) (object interfaces.TweenInterface) {
	el.engine = value
	return el
}

// SetTweenFunc
//
// English:
//
//  Defines the tween math function to control the loop of interactions
//
//   Input:
//     value: tween math function.
//       currentTime:   current time, int64(time.Duration);
//       duration:      total time, int64(time.Duration);
//       startValue:    initial value;
//       endValue:      final value;
//       changeInValue: startValue - endValue
//
//   Output:
//     object: reference to the current Tween object.
//
//   Note:
//     * To create a new function, base it on the linear function, where:
//         return changeInValue * currentTime / duration + startValue
//
// Português:
//
//  Define a função matemática tween para controle do ciclo de interações
//
//   Entrada:
//     value: função matemática tween.
//       currentTime:   tempo atual, int64(time.Duration);
//       duration:      tempo total, int64(time.Duration);
//       startValue:    valor inicial;
//       endValue:      valor final;
//       changeInValue: startValue - endValue
//
//   Saída:
//     object: referência para o objeto Tween corrente.
//
//   Nota:
//     * Para criar uma nova função, tenha como base a função linear, onde:
//         return changeInValue * currentTime / duration + startValue
func (el *Tween) SetTweenFunc(value func(currentTime, duration, currentPercentage, startValue, endValue, changeInValue float64) (percent float64)) (object interfaces.TweenInterface) {
	el.tweenFunc = value
	return el
}

// SetValues
//
// English:
//
//  Defines the initial and final values of the interactions cycle.
//
//   Input:
//     start: initial value for the beginning of the cycle of interactions;
//     end:   final value for the end of the iteration cycle.
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Defines os valores inicial e final do ciclo de interações.
//
//   Entrada:
//     start: valor inicial para o início do ciclo de interações;
//     end:   valor final para o fim do ciclo de interações.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetValues(start, end float64) (object interfaces.TweenInterface) {
	el.startValue = start
	el.endValue = end
	return el
}

// SetDuration
//
// English:
//
//  Defines the total cycle time of interactions.
//
//   Input:
//     value: time.Duration contendo o tempo do ciclo de interações.
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Define o tempo total do ciclo de interações.
//
//   Entrada:
//     value: time.Duration contendo o tempo do ciclo de interações.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetDuration(value time.Duration) (object interfaces.TweenInterface) {
	el.duration = value
	return el
}

// SetDoNotReverseMotion
//
// English:
//
//  Defines the option of reversing values at the end of each cycle.
//
//   Input:
//     value: true to not revert the values at the end of each cycle.
//
//   Output:
//     object: reference to the current Tween object.
//
//   Notas:
//     * In case of loop, the order of event functions are: SetOnStartFunc(), SetOnCycleStartFunc(),
//       SetOnCycleEndFunc(), SetOnInvertFunc(), SetOnCycleStartFunc(), SetOnCycleEndFunc(),
//       SetOnInvertFunc() ...
//     * SetOnEndFunc() will only be called at the end of all interactions;
//     * This function prevents inversion of values, but the SetOnInvertFunc() event function
//       continues to be called.
//
// Português:
//
//  Define a opção de reversão de valores ao final de cada ciclo.
//
//   Entrada:
//     value: true para não reverter os valores ao final de cada ciclo.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
//
//   Notas:
//     * Em caso de laço, a ordem das funções de eventos são: SetOnStartFunc(), SetOnCycleStartFunc(),
//       SetOnCycleEndFunc(), SetOnInvertFunc(), SetOnCycleStartFunc(), SetOnCycleEndFunc(),
//       SetOnInvertFunc() ...
//     * SetOnEndFunc() só será chamada ao final de todas as interações.
//     * Esta função impede a inversão de valores, mas, a função de evento SetOnInvertFunc() continua
//       sendo chamada.
func (el *Tween) SetDoNotReverseMotion(value bool) (object interfaces.TweenInterface) {
	el.doNotReverseMotion = value
	return el
}

// SetLoops
//
// English:
//
//  Defines the number of loops before the end of the function.
//
//   Notes:
//     * At each new iteration of the loop, a movement inversion will occur, unless the
//       SetDoNotReverseMotion(true) function is used;
//     * For infinite loops, set the value to -1;
//     * In case of loop, the order of event functions are: SetOnStartFunc(), SetOnCycleStartFunc(),
//       SetOnCycleEndFunc(), SetOnInvertFunc(), SetOnCycleStartFunc(), SetOnCycleEndFunc(),
//       SetOnInvertFunc() ...
//     * SetOnEndFunc() will only be called at the end of all interactions.
//
// Português:
//
//  Define a quantidade de laços antes do fim da função.
//
//   Notas:
//     * A cada nova interação do laço ocorrerá uma inversão de movimento, a não ser que seja usada a
//       função SetDoNotReverseMotion(true);
//     * Para laços infinitos, defina o valor como sendo -1;
//     * Em caso de laço, a ordem das funções de eventos são: SetOnStartFunc(), SetOnCycleStartFunc(),
//       SetOnCycleEndFunc(), SetOnInvertFunc(), SetOnCycleStartFunc(), SetOnCycleEndFunc(),
//       SetOnInvertFunc() ...
//     * SetOnEndFunc() só será chamada ao final de todas as interações.
func (el *Tween) SetLoops(value int) (object interfaces.TweenInterface) {
	el.repeat = value
	return el
}

// SetOnStartFunc
//
// English:
//
//  Add the function to be called when the animation starts.
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: initial value defined in startValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada quando a animação inicia.
//
//   Entrada:
//     function: func(value float64, arguments ...interface{})
//       value: valor inicial definido em startValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnStartFunc(function func(value float64, arguments interface{})) (object interfaces.TweenInterface) {
	el.onStart = function
	return el
}

// SetOnEndFunc
//
// English:
//
//  Add the function to be called when the animation ends.
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: final value defined in endValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada quando a animação inicia.
//
//   Entrada:
//     function: func(value float64, arguments ...interface{})
//       value: valor final definido em endValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnEndFunc(function func(value float64, arguments interface{})) (object interfaces.TweenInterface) {
	el.onEnd = function
	return el
}

// SetOnCycleStartFunc
//
// English:
//
//  Adds the function to be called at the beginning of the interpolation cycle
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: initial value defined in startValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada no início do ciclo de interpolação
//
//   Entrada:
//     function: func(value float64, arguments ...interface{})
//       value: valor inicial definido em startValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnCycleStartFunc(function func(value float64, arguments interface{})) (object interfaces.TweenInterface) {
	el.onCycleStart = function
	return el
}

// SetOnCycleEndFunc
//
// English:
//
//  Adds the function to be called at the ending of the interpolation cycle
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: final value defined in endValue
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada no fim do ciclo de interpolação
//
//   Entrada:
//     function: func(value float64, arguments ...interface{})
//       value: valor final definido em endValue
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnCycleEndFunc(function func(value float64, arguments interface{})) (object interfaces.TweenInterface) {
	el.onCycleEnd = function
	return el
}

// SetOnStepFunc
//
// English:
//
//  Adds the function to be called for each iteration.
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: current value
//       percentToComplete: value between 0.0 and 1.0 indicating the percentage of the process
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada a cada interação
//
//   Entrada:
//     function: func(value float64, arguments ...interface{})
//       value: valor corrente
//       percentToComplete: valor entre 0.0 e 1.0 indicando o percentual do processo
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnStepFunc(function func(value, percentToComplete float64, arguments interface{})) (object interfaces.TweenInterface) {
	el.interaction = function
	return el
}

// SetOnInvertFunc
//
// English:
//
//  Adds the function to be called on inversion of the interpolation cycle
//
//   Input:
//     function: func(value float64, arguments ...interface{})
//       value: current value
//       arguments: list of values passed to event functions, defined in SetArguments()
//
// Português:
//
//  Adiciona a função a ser chamada a cada interação
//
//   Entrada:
//     function: func(value, percentToComplete float64, arguments ...interface{})
//       value: valor corrente
//       arguments: lista de valores passados para as funções de evento, definidos em SetArguments()
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) SetOnInvertFunc(function func(value float64, arguments interface{})) (object interfaces.TweenInterface) {
	el.onInvert = function
	return el
}

// SetArgumentsFunc
//
// English:
//
//  Determines the arguments passed to event functions.
//
//   Input:
//     arguments: list of interfaces{} passed to all event functions when they are invoked.
//
//   Output:
//     object: reference to the current Tween object.
//
//   Note:
//     * If you need complex functions, remember to use pointers to data in the arguments.
//
// Português:
//
//  Determina os argumentos passados para as funções de eventos.
//
//   Entrada:
//     arguments: lista de interfaces{} passadas para todas as funções de eventos quando elas são invocadas.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
//
//   Nota:
//     * Caso necessite de funções complexas, lembre-se de usar ponteiros para dados nos argumentos.
func (el *Tween) SetArgumentsFunc(arguments interface{}) (object interfaces.TweenInterface) {
	el.arguments = arguments
	return el
}

// Start
//
// English:
//
//  Starts the interaction according to the chosen tween function.
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Inicia a interação conforme a função tween escolhida.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) Start() (object interfaces.TweenInterface) {

	if el.engine == nil {
		el.engine = globalEngine.Engine
	}

	if el.tweenFunc == nil {
		el.tweenFunc = KLinear
	}

	el.startTime = time.Now()
	el.invert = true

	if el.onStart != nil {
		el.onStart(el.startValue, el.arguments)
	}

	el.tickerRunnerPrepare(el.startValue, el.endValue)

	return el
}

func (el *Tween) tickerRunnerPrepare(startValue, endValue float64) {
	if el.onCycleStart != nil {
		el.onCycleStart(el.startValue, el.arguments)
	}

	el.loopStartValue = startValue
	el.loopEndValue = endValue

	if el.engineHasFunction == false {
		el.engineHasFunction = true
		el.fpsUId, _ = el.engine.MathAddToFunctions(el.tickerRunnerRun)
	}
}

func (el *Tween) tickerRunnerRun() {
	elapsed := time.Since(el.startTime)
	percent := elapsed.Seconds() / el.duration.Seconds()
	value := el.tweenFunc(elapsed.Seconds(), el.duration.Seconds(), percent, el.loopStartValue, el.loopEndValue, el.loopEndValue-el.loopStartValue)

	if el.interaction != nil {
		el.interaction(value, percent, el.arguments)
	}

	if elapsed >= el.duration {

		el.Stop()

		if el.repeat == 0 && el.onEnd != nil {
			el.onEnd(value, el.arguments)
		}

		if el.repeat != 0 {
			el.startTime = time.Now()

			if el.onInvert != nil {
				el.onInvert(value, el.arguments)
			}

			if el.doNotReverseMotion == true {
				el.tickerRunnerPrepare(el.startValue, el.endValue)
			} else {
				if el.invert == true {
					el.tickerRunnerPrepare(el.endValue, el.startValue)
				} else {
					el.tickerRunnerPrepare(el.startValue, el.endValue)
				}
				el.invert = !el.invert
			}

			el.repeat -= 1
		}
	}
}

// End
//
// English:
//
//  Terminates all interactions of the chosen Tween function, without invoking the onCycleEnd and
//  onEnd functions.
//
//   Saída:
//     object: reference to the current Tween object.
//
// Português:
//
// Termina todas as interações da função Tween escolhida, sem invocar as funções onCycleEnd e onEnd.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) End() (object interfaces.TweenInterface) {
	el.engineHasFunction = false
	el.engine.MathDeleteFromFunctions(el.fpsUId)

	return el
}

// Stop
//
// English:
//
//  Ends all interactions of the chosen Tween function, interacting with the onCycleEnd and onEnd
//  functions, respectively, in that order, if they have been defined.
//
//   Output:
//     object: reference to the current Tween object.
//
// Português:
//
//  Termina todas as interações da função Tween escolhida, interagindo com as funções onCycleEnd e
//  onEnd, respectivamente nessa ordem, se elas tiverem sido definidas.
//
//   Saída:
//     object: referência para o objeto Tween corrente.
func (el *Tween) Stop() (object interfaces.TweenInterface) {
	el.engineHasFunction = false
	el.engine.MathDeleteFromFunctions(el.fpsUId)

	if el.onCycleEnd != nil {
		el.onCycleEnd(el.endValue, el.arguments)
	}

	if el.onEnd != nil && el.repeat == 0 {
		el.onEnd(el.endValue, el.arguments)
	}

	return el
}
