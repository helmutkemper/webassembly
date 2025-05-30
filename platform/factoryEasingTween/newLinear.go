package factoryEasingTween

import (
	"github.com/helmutkemper/webassembly/platform/easingTween"
	"time"
)

// NewLinear
//
// English:
//
//	Ease tween linear.
//
//	 Input:
//
//	   duration: animation duration
//	   startValue: initial value
//	   endValue: final value
//	   onStepFunc: on step function
//	   loop: number of loops or -1 for infinite loops
//	   arguments: array of arguments passed for functions onStart, onEnd, onInvert and onStep.
//	      Example: ..., [arguments] x, y) will be onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
//
// Português:
//
//	Facilitador de interpolação linear.
//
//	 Entrada:
//
//	   duration: duração da animação
//	   startValue: valor inicial
//	   endValue: valor final
//	   onStepFunc: função para o evento passo
//	   loop: número de interações ou -1 para um número infinito de interações
//	   arguments: array de argumentos passados para as funções onStart, onEnd, onInvert e onStep.
//	      Exemplo: ..., [argumentos] x, y) será onStartFunc(value, args...) { args[0]: this; args[1]: x; args[2]: y}
func NewLinear(
	duration time.Duration,
	startValue,
	endValue float64,
	onStepFunc interface{},
	loop int,
	arguments ...interface{},
) *easingTween.Tween {

	var function func(value, percentToComplete float64, arguments interface{})
	switch converted := onStepFunc.(type) {
	case func(value, percentToComplete float64, arguments interface{}):
		function = converted
	case func() (f func(value, percentToComplete float64, arguments interface{})):
		function = converted()
	}

	t := &easingTween.Tween{}
	t.SetDuration(duration).
		SetValues(startValue, endValue).
		SetOnStepFunc(function).
		SetLoops(loop).
		SetArgumentsFunc(arguments).
		SetTweenFunc(easingTween.KLinear).
		Start()

	return t
}
