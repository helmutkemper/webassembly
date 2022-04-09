package canvas

import (
	"image/color"
	"syscall/js"
)

// AddColorStopPosition
//
// English:
//
//  Specifies the colors and stop positions in a gradient object
//     Input:
//       gradient:     A gradient object created by CreateLinearGradient() or CreateRadialGradient()
//                     methods;
//       stopPosition: A value between 0.0 and 1.0 that represents the position between start (0%) and
//                     end (100%) in a gradient;
//       color:        A color RGBA value to display at the stop position;
//
//     Note:
//
//       * You can call the addColorStopPosition() method multiple times to change a gradient. If you
//         omit this method for gradient objects, the gradient will not be visible. You need to create
//         at least one color stop to have a visible gradient.
//
// Português:
//
//  Especifica a cor e a posição final para a cor dentro do gradiente
//   Entrada:
//     gradient:     Objeto de gradiente criado pelos métodos CreateLinearGradient() ou
//                   CreateRadialGradient();
//     stopPosition: Um valor entre 0.0 e 1.0 que representa a posição entre o início (0%) e o fim
//                   (100%) dentro do gradiente;
//     color:        Uma cor no formato RGBA para ser mostrada na posição determinada.
//
//     Nota: Você pode chamar o método AddColorStopPosition() várias vezes para adicionar várias cores ao gradiente, porém, se
//     você omitir o método, o gradiente não será visível. Você tem a obrigação de chamar o método pelo menos uma vez com uma cor
//     para que o gradiente seja visível.
func (el *Canvas) AddColorStopPosition(gradient interface{}, stopPosition float64, color color.RGBA) {
	gradient.(js.Value).Call("addColorStop", stopPosition, RGBAToJs(color))
}
