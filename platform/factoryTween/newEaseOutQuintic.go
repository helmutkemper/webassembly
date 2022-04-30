package factoryTween

import (
	"github.com/helmutkemper/iotmaker.webassembly/platform/tween"
	"time"
)

// NewEaseOutQuintic
//
// English: Ease tween ease out quintic
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//                Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: x; args[1]: y}
//
// Português: Facilitador de interpolação ease out quintic
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//                Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: x; args[1]: y}
func NewEaseOutQuintic(
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments ...interface{}),
	loop int,
	arguments ...interface{},
) *tween.Tween {

	t := &tween.Tween{}
	t.SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(tween.KEaseOutQuintic).
		Start()

	return t
}
