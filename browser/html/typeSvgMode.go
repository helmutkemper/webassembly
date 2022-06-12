package html

// SvgMode
//
// English:
//
// For each pixel among the layers to which it is applied, a blend mode takes the colors of the foreground and the
// background, performs a calculation on them, and returns a new color value.
//
// Changes between blend modes are not interpolated. Any change occurs immediately.
//
// Português:
//
// Para cada pixel entre as camadas às quais é aplicado, um modo de mesclagem pega as cores do primeiro plano e do
// plano de fundo, realiza um cálculo sobre elas e retorna um novo valor de cor.
//
// As alterações entre os modos de mesclagem não são interpoladas. Qualquer mudança ocorre imediatamente.
type SvgMode string

func (e SvgMode) String() string {
	return string(e)
}

const (
	// KSvgModeNormal
	//
	// English:
	//
	// The final color is the top color, regardless of what the bottom color is. The effect is like two opaque pieces of
	// paper overlapping.
	//
	// Português:
	//
	// A cor final é a cor de cima, independentemente da cor de baixo. O efeito é como dois pedaços opacos de papel
	// sobrepostos.
	KSvgModeNormal SvgMode = "normal"

	// KSvgModeMultiply
	//
	// English:
	//
	// The final color is the result of multiplying the top and bottom colors. A black layer leads to a black final layer,
	// and a white layer leads to no change. The effect is like two images printed on transparent film overlapping.
	//
	// Português:
	//
	// A cor final é o resultado da multiplicação das cores superior e inferior. Uma camada preta leva a uma camada final
	// preta e uma camada branca não leva a nenhuma alteração. O efeito é como duas imagens impressas em filme
	// transparente sobrepostas.
	KSvgModeMultiply SvgMode = "multiply"

	// KSvgModeScreen
	//
	// English:
	//
	// The final color is the result of inverting the colors, multiplying them, and inverting that value. A black layer
	// leads to no change, and a white layer leads to a white final layer. The effect is like two images shone onto a
	// projection screen.
	//
	// Português:
	//
	// A cor final é o resultado de inverter as cores, multiplicá-las e inverter esse valor. Uma camada preta não leva a
	// nenhuma alteração e uma camada branca leva a uma camada final branca. O efeito é como duas imagens brilhando em uma
	// tela de projeção.
	KSvgModeScreen SvgMode = "screen"

	// KSvgModeOverlay
	//
	// English:
	//
	// The final color is the result of multiply if the bottom color is darker, or screen if the bottom color is lighter.
	// This blend mode is equivalent to hard-light but with the layers swapped.
	//
	// Português:
	//
	// A cor final é o resultado da multiplicação se a cor de baixo for mais escura, ou da tela se a cor de baixo for mais
	// clara. Este modo de mesclagem é equivalente à luz forte, mas com as camadas trocadas.
	KSvgModeOverlay SvgMode = "overlay"

	// KSvgModeDarken
	//
	// English:
	//
	// The final color is composed of the darkest values of each color channel.
	//
	// Português:
	//
	// A cor final é composta pelos valores mais escuros de cada canal de cor.
	KSvgModeDarken SvgMode = "darken"

	// KSvgModeLighten
	//
	// English:
	//
	// The final color is composed of the lightest values of each color channel.
	//
	// Português:
	//
	// A cor final é composta pelos valores mais claros de cada canal de cor.
	KSvgModeLighten SvgMode = "lighten"

	// KSvgModeColorDodge
	//
	// English:
	//
	// The final color is the result of dividing the bottom color by the inverse of the top color.
	// A black foreground leads to no change. A foreground with the inverse color of the backdrop leads to a fully lit
	// color. This blend mode is similar to screen, but the foreground need only be as light as the inverse of the
	// backdrop to create a fully lit color.
	//
	// Português:
	//
	// A cor final é o resultado da divisão da cor de baixo pelo inverso da cor de cima.
	// Um primeiro plano preto não leva a nenhuma mudança. Um primeiro plano com a cor inversa do pano de fundo leva a uma
	// cor totalmente iluminada. Esse modo de mesclagem é semelhante à tela, mas o primeiro plano precisa ser tão claro
	// quanto o inverso do plano de fundo para criar uma cor totalmente iluminada.
	KSvgModeColorDodge SvgMode = "color-dodge"

	// KSvgModeColorBurn
	//
	// English:
	//
	// The final color is the result of inverting the bottom color, dividing the value by the top color, and inverting
	// that value.
	// A white foreground leads to no change. A foreground with the inverse color of the backdrop leads to a black final
	// image. This blend mode is similar to multiply, but the foreground need only be as dark as the inverse of the
	// backdrop to make the final image black.
	//
	// Português:
	//
	// A cor final é o resultado da inversão da cor inferior, dividindo o valor pela cor superior e invertendo esse valor.
	// Um primeiro plano branco não leva a nenhuma mudança. Um primeiro plano com a cor inversa do pano de fundo leva a
	// uma imagem final preta. Esse modo de mesclagem é semelhante à multiplicação, mas o primeiro plano precisa ser tão
	// escuro quanto o inverso do pano de fundo para tornar a imagem final preta.
	KSvgModeColorBurn SvgMode = "color-burn"

	// KSvgModeHardLight
	//
	// English:
	//
	// The final color is the result of multiply if the top color is darker, or screen if the top color is lighter.
	// This blend mode is equivalent to overlay but with the layers swapped. The effect is similar to shining a harsh
	// spotlight on the backdrop.
	//
	// Português:
	//
	// A cor final é o resultado da multiplicação se a cor de cima for mais escura, ou da tela se a cor de cima for mais
	// clara.
	// Este modo de mesclagem é equivalente à sobreposição, mas com as camadas trocadas. O efeito é semelhante ao de um
	// holofote forte sobre o pano de fundo.
	KSvgModeHardLight SvgMode = "hard-light"

	// KSvgModeSoftLight
	//
	// English:
	//
	// The final color is similar to hard-light, but softer. This blend mode behaves similar to hard-light.
	// The effect is similar to shining a diffused spotlight on the backdrop*.*
	//
	// Português:
	//
	// A cor final é semelhante à luz dura, mas mais suave. Esse modo de mesclagem se comporta de maneira semelhante à
	// luz forte.
	// O efeito é semelhante ao de um holofote difuso no fundo.
	KSvgModeSoftLight SvgMode = "soft-light"

	// KSvgModeDifference
	//
	// English:
	//
	// The final color is the result of subtracting the darker of the two colors from the lighter one.
	// A black layer has no effect, while a white layer inverts the other layer's color.
	//
	// Português:
	//
	// A cor final é o resultado da subtração da mais escura das duas cores da mais clara.
	// Uma camada preta não tem efeito, enquanto uma camada branca inverte a cor da outra camada.
	KSvgModeDifference SvgMode = "difference"

	// KSvgModeExclusion
	//
	// English:
	//
	// The final color is similar to difference, but with less contrast.
	// As with difference, a black layer has no effect, while a white layer inverts the other layer's color.
	//
	// Português:
	//
	// A cor final é semelhante à diferença, mas com menos contraste.
	// Tal como acontece com a diferença, uma camada preta não tem efeito, enquanto uma camada branca inverte a cor da
	// outra camada.
	KSvgModeExclusion SvgMode = "exclusion"

	// KSvgModeHue
	//
	// English:
	//
	// The final color has the hue of the top color, while using the saturation and luminosity of the bottom color.
	//
	// Português:
	//
	// A cor final tem a tonalidade da cor de cima, enquanto usa a saturação e luminosidade da cor de baixo.
	KSvgModeHue SvgMode = "hue"

	// KSvgModeSaturation
	//
	// English:
	//
	// The final color has the saturation of the top color, while using the hue and luminosity of the bottom color.
	// A pure gray backdrop, having no saturation, will have no effect.
	//
	// Português:
	//
	// A cor final tem a saturação da cor de cima, enquanto usa a tonalidade e a luminosidade da cor de baixo.
	// Um pano de fundo cinza puro, sem saturação, não terá efeito.
	KSvgModeSaturation SvgMode = "saturation"

	// KSvgModeColor
	//
	// English:
	//
	// The final color has the hue and saturation of the top color, while using the luminosity of the bottom color.
	// The effect preserves gray levels and can be used to colorize the foreground.
	//
	// Português:
	//
	// A cor final tem o matiz e a saturação da cor de cima, enquanto usa a luminosidade da cor de baixo.
	// O efeito preserva os níveis de cinza e pode ser usado para colorir o primeiro plano.
	KSvgModeColor SvgMode = "color"

	// KSvgModeLuminosity
	//
	// English:
	//
	// The final color has the luminosity of the top color, while using the hue and saturation of the bottom color.
	// This blend mode is equivalent to color, but with the layers swapped.
	//
	// Português:
	//
	// A cor final tem a luminosidade da cor de cima, enquanto usa o matiz e a saturação da cor de baixo.
	// Este modo de mesclagem é equivalente à cor, mas com as camadas trocadas.
	KSvgModeLuminosity SvgMode = "luminosity"
)
