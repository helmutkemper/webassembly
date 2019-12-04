package canvas

// en: The stroke() method actually draws the path you have defined with all those moveTo() and lineTo() methods. The default color is black.
//     Tip: Use the strokeStyle property to draw with another color/gradient.
//
// pt_br: O método stroke() desenha o caminho definido com os métodos moveTo() e lineTo() usando a cor padrão, preta.
//     Dica: Use a propriedade strokeStyle para desenhar com outra cor ou usar um gradiente
func (el *Canvas) Stroke() {
	el.SelfContext.Call("stroke")
}
