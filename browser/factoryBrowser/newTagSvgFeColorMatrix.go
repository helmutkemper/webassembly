package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeColorMatrix
//
// English:
//
// The <feColorMatrix> SVG filter element changes colors based on a transformation matrix.
// Every pixel's color value [R,G,B,A] is matrix multiplied by a 5 by 5 color matrix to create new color [R',G',B',A'].
//
//   Notes:
//     * The prime symbol ' is used in mathematics indicate the result of a transformation.
//
//  | R' |     | r1 r2 r3 r4 r5 |   | R |
//  | G' |     | g1 g2 g3 g4 g5 |   | G |
//  | B' |  =  | b1 b2 b3 b4 b5 | * | B |
//  | A' |     | a1 a2 a3 a4 a5 |   | A |
//  | 1  |     |  0  0  0  0  1 |   | 1 |
//
// In simplified terms, below is how each color channel in the new pixel is calculated. The last row is ignored because
// its values are constant.
//
//  R' = r1*R + r2*G + r3*B + r4*A + r5
//  G' = g1*R + g2*G + g3*B + g4*A + g5
//  B' = b1*R + b2*G + b3*B + b4*A + b5
//  A' = a1*R + a2*G + a3*B + a4*A + a5
//
// Take the amount of red in the new pixel, or R':
//
// It is the sum of:
//   * r1 times the old pixel's red R,
//   * r2 times the old pixel's green G,
//   * r3 times of the old pixel's blue B,
//   * r4 times the old pixel's alpha A,
//   * plus a shift r5.
//
// These specified amounts can be any real number, though the final R' will be clamped between 0 and 1. The same goes
// for G', B', and A'.
//
//  R'      =      r1 * R      +        r2 * G      +       r3 * B      +       r4 * A       +       r5
//  New red = [ r1 * old red ] + [ r2 * old green ] + [ r3 * old Blue ] + [ r4 * old Alpha ] + [ shift of r5 ]
//
// If, say, we want to make a completely black image redder, we can make the r5 a positive real number x, boosting the
// redness on every pixel of the new image by x.
//
// An identity matrix looks like this:
//
//       R G B A W
//  R' | 1 0 0 0 0 |
//  G' | 0 1 0 0 0 |
//  B' | 0 0 1 0 0 |
//  A' | 0 0 0 1 0 |
//
// In it, every new value is exactly 1 times its old value, with nothing else added. It is recommended to start
// manipulating the matrix from here.
//
// Português:
//
// O elemento de filtro SVG <feColorMatrix> altera as cores com base em uma matriz de transformação.
// O valor de cor de cada pixel [R,G,B,A] é uma matriz multiplicada por uma matriz de cores de 5 por 5 para criar uma
// nova cor [R',G',B',A'].
//
//   Notas:
//     * O símbolo primo ' é usado em matemática para indicar o resultado de uma transformação.
//
//  | R' |     | r1 r2 r3 r4 r5 |   | R |
//  | G' |     | g1 g2 g3 g4 g5 |   | G |
//  | B' |  =  | b1 b2 b3 b4 b5 | * | B |
//  | A' |     | a1 a2 a3 a4 a5 |   | A |
//  | 1  |     |  0  0  0  0  1 |   | 1 |
//
// Em termos simplificados, abaixo está como cada canal de cor no novo pixel é calculado. A última linha é ignorada
// porque seus valores são constantes.
//
//  R' = r1*R + r2*G + r3*B + r4*A + r5
//  G' = g1*R + g2*G + g3*B + g4*A + g5
//  B' = b1*R + b2*G + b3*B + b4*A + b5
//  A' = a1*R + a2*G + a3*B + a4*A + a5
//
// Pegue a quantidade de vermelho no novo pixel, ou R':
//
// É a soma de:
//   * r1 vezes o antigo pixel vermelho, R,
//   * r2 vezes o antigo pixel verde, G,
//   * r3 vezes o antigo pixel azul, B,
//   * r4 vezes o antigo pixel alpha, A,
//   * mais um turno r5.
//
// Esses valores especificados podem ser qualquer número real, embora o R' final seja fixado entre 0 e 1. O mesmo vale
// para G', B' e A'.
//
//  R'      =      r1 * R      +        r2 * G      +       r3 * B      +       r4 * A       +       r5
//  New red = [ r1 * old red ] + [ r2 * old green ] + [ r3 * old Blue ] + [ r4 * old Alpha ] + [ shift of r5 ]
//
// Se, digamos, queremos tornar uma imagem completamente preta mais vermelha, podemos tornar o r5 um número real
// positivo x, aumentando a vermelhidão em cada pixel da nova imagem em x.
//
// Uma matriz identidade fica assim:
//
//       R G B A W
//  R' | 1 0 0 0 0 |
//  G' | 0 1 0 0 0 |
//  B' | 0 0 1 0 0 |
//  A' | 0 0 0 1 0 |
//
// Nele, cada novo valor é exatamente 1 vez seu valor antigo, sem mais nada adicionado. Recomenda-se começar a
// manipular a matriz a partir daqui.
func NewTagSvgFeColorMatrix() (ref *html.TagSvgFeColorMatrix) {
	ref = &html.TagSvgFeColorMatrix{}
	ref.Init()

	return ref
}
