package iotmaker_platform_IDraw

type IFilterGradientInterface interface {
	SetP0(point interface{})
	SetP1(point interface{})
	PrepareFilter(platform ICanvasGradient)

	// Fill
	// en: The fill() method fills the current drawing (path). The default color is
	//     black.
	//     Tip: Use the fillStyle property to fill with another color/gradient.
	//     Note: If the path is not closed, the fill() method will add a line from the
	//     last point to the start point of the path to close the path (like
	//     closePath()), and then fill the path.
	//
	// pt_br: O método fill() preenche e pinta do desenho (caminho). A cor padrão é
	//     preto.
	//     Dica: Use a propriedade fillStyle para adicionar uma cor ou gradient.
	//     Nota: Se o caminho não estiver fechado, o método fill() irá adicioná uma
	//     linha do último ao primeiro ponto do caminho para fechar o caminho
	//     (semelhante ao método closePath()) e só então irá pintar
	Fill(gradient interface{})

	// Stroke
	// en: The stroke() method actually draws the path you have defined with all those
	//     moveTo() and lineTo() methods. The default color is black.
	//     Tip: Use the strokeStyle property to draw with another color/gradient.
	//
	// pt_br: O método stroke() desenha o caminho definido com os métodos moveTo() e
	//     lineTo() usando a cor padrão, preta.
	//     Dica: Use a propriedade strokeStyle para desenhar com outra cor ou usar um
	//     gradiente
	Stroke(gradient interface{})
}
