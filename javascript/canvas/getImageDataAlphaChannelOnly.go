package canvas

// GetImageDataAlphaChannelOnly
// en: Returns an ImageData map[x][y]uint8 that copies the pixel alpha channel for
// the specified rectangle on a canvas
//     x: The x coordinate (in pixels) of the upper-left corner to start copy from
//     y: The y coordinate (in pixels) of the upper-left corner to start copy from
//     width: The width of the rectangular area you will copy
//     height: The height of the rectangular area you will copy
//     return: map[x(int)][y(int)]uint8
//             Note: return x and y are NOT relative to the coordinate (0,0) on the
//             image, are relative to the coordinate (0,0) on the canvas
//
//     Note: The ImageData object is not a picture, it specifies a part (rectangle)
//     on the canvas, and holds information only for alpha channel of every pixel
//     inside that rectangle.
//
//     For every pixel in the map[x][y] there are one piece of information, the
//     alpha channel uint8 value (from 0-255; 0 is transparent and 255 is fully
//     visible)
//
//     Tip: After you have manipulated the color/alpha information in the map[x][y],
//     you can copy the image data back onto the canvas with the putImageData()
//     method.
//
// pr_br: Retorna um mapa map[x][y]uint8 com parte dos dados da imagem contida
// no retângulo especificado.
//     x: Coordenada x (em pixels) do canto superior esquerdo de onde os dados vão
//     ser copiados
//     y: Coordenada y (em pixels) do canto superior esquerdo de onde os dados vão
//     ser copiados
//     width: comprimento do retângulo a ser copiado
//     height: altura do retângulo a ser copiado
//     return: map[x(int)][y(int)]uint8
//             Nota: x e y do retorno não são relativos a coordenada (0,0) da imagem,
//             são relativos a coordenada (0,0) do canvas
//
//     Nota: Os dados da imagem não são uma figura, eles representam uma parte
//     retangular do canvas e guardam informações apenas do canal alpha de cada pixel
//     contido nessa área
//
//     Para cada pixel contido no mapa há apenas uma peça da informação do canal
//     alpha com valores no formato uint8, com valoes de 0-255; onde, 0 é
//     transparente e 255 é totalmente visível
//
//     Dica: Depois de manipular as informações de cor/alpha contidas no map[x][y],
//     elas podem ser colocadas de volta no canvas com o método putImageData().
func (el *Canvas) GetImageDataAlphaChannelOnly(x, y, width, height int) map[int]map[int]uint8 {

	dataInterface := el.SelfContext.Call("getImageData", x, y, width, height)
	dataJs := dataInterface.Get("data")

	ret := make(map[int]map[int]uint8)

	var rgbaLength int = 4

	var tmp uint8
	var i int = 0
	var xp int
	var yp int
	for yp = 0; yp != height; yp += 1 {
		for xp = 0; xp != width; xp += 1 {

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			tmp = uint8(dataJs.Index(i + 3).Int())

			i += rgbaLength

			if tmp == 0 {
				continue
			}

			if len(ret[x+xp]) == 0 {
				ret[x+xp] = make(map[int]uint8)
			}

			ret[x+xp][y+yp] = tmp
		}
	}

	return ret
}
