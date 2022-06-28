package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeConvolveMatrix
//
// English:
//
// The <feConvolveMatrix> SVG filter primitive applies a matrix convolution filter effect. A convolution combines pixels
// in the input image with neighboring pixels to produce a resulting image. A wide variety of imaging operations can be
// achieved through convolutions, including blurring, edge detection, sharpening, embossing and beveling.
//
// A matrix convolution is based on an n-by-m matrix (the convolution kernel) which describes how a given pixel value
// in the input image is combined with its neighboring pixel values to produce a resulting pixel value. Each result
// pixel is determined by applying the kernel matrix to the corresponding source pixel and its neighboring pixels.
//
// The basic convolution formula which is applied to each color value for a given pixel is:
//
//   COLORX,Y = ( SUM I=0 to [orderY-1] { SUM J=0 to [orderX-1] { SOURCE X-targetX+J, Y-targetY+I *
//                kernelMatrixorderX-J-1, orderY-I-1 } } ) / divisor + bias * ALPHAX,Y
//
// where "orderX" and "orderY" represent the X and Y values for the 'order' attribute, "targetX" represents the value of
// the 'targetX' attribute, "targetY" represents the value of the 'targetY' attribute, "kernelMatrix" represents the
// value of the 'kernelMatrix' attribute, "divisor" represents the value of the 'divisor' attribute, and "bias"
// represents the value of the 'bias' attribute.
//
// Note in the above formulas that the values in the kernel matrix are applied such that the kernel matrix is rotated
// 180 degrees relative to the source and destination images in order to match convolution theory as described in many
// computer graphics textbooks.
//
// To illustrate, suppose you have a input image which is 5 pixels by 5 pixels, whose color values for one of the color
// channels are as follows:
//
//   0    20  40 235 235
//   100 120 140 235 235
//   200 220 240 235 235
//   225 225 255 255 255
//   225 225 255 255 255
//
// and you define a 3-by-3 convolution kernel as follows:
//
//   1 2 3
//   4 5 6
//   7 8 9
//
// Let's focus on the color value at the second row and second column of the image (source pixel value is 120). Assuming
// the simplest case (where the input image's pixel grid aligns perfectly with the kernel's pixel grid) and assuming
// default values for attributes 'divisor', 'targetX' and 'targetY', then resulting color value will be:
//
//   (9*  0 + 8* 20 + 7* 40 +
//    6*100 + 5*120 + 4*140 +
//    3*200 + 2*220 + 1*240) /
//   (9+8+7+6+5+4+3+2+1)
//
// Português:
//
// A primitiva de filtro SVG <feConvolveMatrix> aplica um efeito de filtro de convolução de matriz. Uma convolução
// combina pixels na imagem de entrada com pixels vizinhos para produzir uma imagem resultante. Uma ampla variedade de
// operações de imagem pode ser alcançada por meio de convoluções, incluindo desfoque, detecção de bordas, nitidez,
// relevo e chanfro.
//
// Uma convolução de matriz é baseada em uma matriz n por m (o kernel de convolução) que descreve como um determinado
// valor de pixel na imagem de entrada é combinado com seus valores de pixel vizinhos para produzir um valor de pixel
// resultante. Cada pixel resultante é determinado pela aplicação da matriz kernel ao pixel de origem correspondente e
// seus pixels vizinhos. A fórmula de convolução básica que é aplicada a cada valor de cor para um determinado pixel é:
//
//   COLORX,Y = ( SUM I=0 to [orderY-1] { SUM J=0 to [orderX-1] { SOURCE X-targetX+J, Y-targetY+I *
//                kernelMatrixorderX-J-1, orderY-I-1 } } ) / divisor + bias * ALPHAX,Y
//
// onde "orderX" e "orderY" representam os valores X e Y para o atributo 'order', "targetX" representa o valor do
// atributo 'targetX', "targetY" representa o valor do atributo 'targetY', "kernelMatrix" representa o valor do atributo
// 'kernelMatrix', "divisor" representa o valor do atributo 'divisor' e "bias" representa o valor do atributo 'bias'.
//
// Observe nas fórmulas acima que os valores na matriz do kernel são aplicados de tal forma que a matriz do kernel é
// girada 180 graus em relação às imagens de origem e destino para corresponder à teoria de convolução conforme descrito
// em muitos livros de computação gráfica.
//
// Para ilustrar, suponha que você tenha uma imagem de entrada com 5 pixels por 5 pixels, cujos valores de cor para um
// dos canais de cores sejam os seguintes:
//
//   0    20  40 235 235
//   100 120 140 235 235
//   200 220 240 235 235
//   225 225 255 255 255
//   225 225 255 255 255
//
// e você define um kernel de convolução 3 por 3 da seguinte forma:
//
//   1 2 3
//   4 5 6
//   7 8 9
//
// Vamos nos concentrar no valor da cor na segunda linha e na segunda coluna da imagem (o valor do pixel de origem é
// 120). Assumindo o caso mais simples (onde a grade de pixels da imagem de entrada se alinha perfeitamente com a grade
// de pixels do kernel) e assumindo valores padrão para os atributos 'divisor', 'targetX' e 'targetY', o valor da cor
// resultante será:
//
//   (9*  0 + 8* 20 + 7* 40 +
//    6*100 + 5*120 + 4*140 +
//    3*200 + 2*220 + 1*240) /
//   (9+8+7+6+5+4+3+2+1)
func NewTagSvgFeConvolveMatrix() (ref *html.TagSvgFeConvolveMatrix) {
	ref = &html.TagSvgFeConvolveMatrix{}
	ref.Init()

	return ref
}
