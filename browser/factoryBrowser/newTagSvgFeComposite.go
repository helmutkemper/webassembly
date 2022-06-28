package factoryBrowser

import "github.com/helmutkemper/iotmaker.webassembly/browser/html"

// NewTagSvgFeComposite
//
// English:
//
// The <feComposite> SVG filter primitive performs the combination of two input images pixel-wise in image space using
// one of the Porter-Duff compositing operations: over, in, atop, out, xor, lighter, or arithmetic.
//
// The table below shows each of these operations using an image of the MDN logo composited with a red circle:
//   * Over: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_over.png
//     The source graphic defined by the in attribute (the MDN logo) is placed over the destination graphic defined by
//     the in2 attribute (the circle).
//     This is the default operation, which will be used if no operation or an unsupported operation is specified.
//   * In: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_in.png
//     The parts of the source graphic defined by the in attribute that overlap the destination graphic defined in the
//     in2 attribute, replace the destination graphic.
//   * Out: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_out.png
//     The parts of the source graphic defined by the in attribute that fall outside the destination graphic defined in
//     the in2 attribute, are displayed.
//   * Atop: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_atop.png
//     The parts of the source graphic defined in the in attribute, which overlap the destination graphic defined in the
//     in2 attribute, replace the destination graphic. The parts of the destination graphic that do not overlap with the
//     source graphic stay untouched.
//   * Xor: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_xor.png
//     The non-overlapping regions of the source graphic defined in the in attribute and the destination graphic defined
//     in the in2 attribute are combined.
//   * Lighter: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_lighter.png
//     The sum of the source graphic defined in the in attribute and the destination graphic defined in the in2
//     attribute is displayed.
//   * Arithmetic: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_arithmetic.png
//     The arithmetic operation is useful for combining the output from the <feDiffuseLighting> and <feSpecularLighting>
//     filters with texture data. If the arithmetic operation is chosen, each result pixel is computed using the
//     following formula:
//       result = k1*i1*i2 + k2*i1 + k3*i2 + k4
//       where:
//         * i1 and i2 indicate the corresponding pixel channel values of the input image, which map to in and in2
//           respectively
//         * k1, k2, k3, and k4 indicate the values of the attributes with the same name.
//
// Português:
//
// A primitiva de filtro SVG <feComposite> executa a combinação de duas imagens de entrada pixel a pixel no espaço da
// imagem usando uma das operações de composição de Porter-Duff: sobre, dentro, em cima, fora, xor, mais leve ou
// aritmética.
//
// A tabela abaixo mostra cada uma dessas operações usando uma imagem do logotipo do MDN composta por um círculo
// vermelho:
//   * Over: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_over.png
//     O gráfico de origem definido pelo atributo in (o logotipo do MDN) é colocado sobre o gráfico de destino definido
//     pelo atributo in2 (o círculo).
//     Esta é a operação padrão, que será usada se nenhuma operação ou uma operação não suportada for especificada.
//   * In: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_in.png
//     As partes do gráfico de origem definidas pelo atributo in que se sobrepõem ao gráfico de destino definido no
//     atributo in2 substituem o gráfico de destino.
//   * Out: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_out.png
//     As partes do gráfico de origem definido pelo atributo in que estão fora do gráfico de destino definido no
//     atributo in2 são exibidas.
//   * Atop: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_atop.png
//     As partes do gráfico de origem definidas no atributo in, que se sobrepõem ao gráfico de destino definido no
//     atributo in2, substituem o gráfico de destino.
//     As partes do gráfico de destino que não se sobrepõem ao gráfico de origem permanecem intactas.
//   * Xor: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_xor.png
//     As regiões não sobrepostas do gráfico de origem definido no atributo in e o gráfico de destino definido no
//     atributo in2 são combinados.
//   * Lighter: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_lighter.png
//     A soma do gráfico de origem definido no atributo in e o gráfico de destino definido no atributo in2 é exibido.
//   * Arithmetic: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite/operation_arithmetic.png
//     A operação aritmética é útil para combinar a saída dos filtros <feDiffeLighting> e <feSpecularLighting> com dados
//     de textura. Se a operação aritmética for escolhida, cada pixel resultante é calculado usando a seguinte fórmula:
//       resultado = k1*i1*i2 + k2*i1 + k3*i2 + k4
//       onde:
//         * i1 e i2 indicam os valores de canal de pixel correspondentes da imagem de entrada, que mapeiam para in e
//           in2 respectivamente
//         * k1, k2, k3 e k4 indicam os valores dos atributos com o mesmo nome.
func NewTagSvgFeComposite() (ref *html.TagSvgFeComposite) {
	ref = &html.TagSvgFeComposite{}
	ref.Init()

	return ref
}
