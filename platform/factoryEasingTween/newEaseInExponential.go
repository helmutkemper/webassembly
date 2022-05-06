package factoryEasingTween

import (
	"github.com/helmutkemper/iotmaker.webassembly/platform/easingTween"
	"time"
)

// NewEaseInExponential
//
// English: Ease tween in exponential
//     duration: animation duration
//     startValue: initial value
//     endValue: final value
//     onStepFunc: on step function
//     loop: number of loops or -1 for infinite loops
//     arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//                Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: x; args[1]: y}
//
// Português: Facilitador de interpolação in exponential
//     duration: duração da animação
//     startValue: valor inicial
//     endValue: valor final
//     onStepFunc: função para o evento passo
//     loop: número de interações ou -1 para um número infinito de interações
//     arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//                Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: x; args[1]: y}
func NewEaseInExponential(
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc func(value, percentToComplete float64, arguments interface{}),
	loop int,
	arguments ...interface{},
) *easingTween.Tween {

	t := &easingTween.Tween{}
	t.SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(onStepFunc).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KEaseInExponential).
		Start()

	return t
}
