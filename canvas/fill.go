package canvas

// en: The fill() method fills the current drawing (path). The default color is black.
//     Tip: Use the fillStyle property to fill with another color/gradient.
//     Note: If the path is not closed, the fill() method will add a line from the last point to the start point of the
//     path to close the path (like closePath()), and then fill the path.
// pt_br: O método fill() preenche e pinta do desenho (caminho). A cor padrão é preto.
//     Dica: Use a propriedade fillStyle para adicionar uma cor ou gradient.
//     Nota: Se o caminho não estiver fechado, o método fill() irá adicioná uma linha do último ao primeiro ponto do
//     caminho para fechar o caminho (semelhante ao método closePath()) e só então irá pintar
func (el *Canvas) Fill() {
	el.SelfContext.Call("fill")
}
