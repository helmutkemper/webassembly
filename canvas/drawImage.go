package canvas

import (
	"log"
	"syscall/js"
)

// en: Draws an image, canvas, or video onto the canvas
//     image: Specifies the image, canvas, or video element to use
//     sx: [optional] The x coordinate where to start clipping
//     sy: [optional] The y coordinate where to start clipping
//     sWidth: [optional] The width of the clipped image
//     sHeight: [optional] The height of the clipped image
//     x: The x coordinate where to place the image on the canvas
//     y: The y coordinate where to place the image on the canvas
//     width: [optional] The width of the image to use (stretch or reduce the image)
//     height: [optional] The height of the image to use (stretch or reduce the image)
//
//     Position the image on the canvas:
//     Golang Syntax: platform.DrawImage(img, x, y)
//
//     Position the image on the canvas, and specify width and height of the image:
//     Golang Syntax: platform.DrawImage(img, x, y, width, height)
//
//     Clip the image and position the clipped part on the canvas:
//     Golang Syntax: platform.drawImage(img, sx, sy, sWidth, sHeight, x, y, width,
//                    height)
//
// pt_br: Desenha uma imagem, canvas ou vídeo no elemento canvas
//     image: Especifica a imagem, canvas ou vídeo a ser usado
//     sx: [opcional] Coordenada x de onde o corte vai começar
//     sy: [opcional] Coordenada y de onde o corte vai começar
//     sWidth: [opcional] largura do corte
//     sHeight: [opcional] altura do corte
//     x: Coordenada x do canvas de onde o corte vai ser colocado
//     y: Coordenada y do canvas de onde o corte vai ser colocado
//     width: [opcional] Novo comprimento da imagem
//     height: [opcional] Nova largura da imagem
//
//     Posiciona a imagem no canvas
//     Golang Sintaxe: platform.DrawImage(img, x, y)
//
//     Posiciona a imagem no canvas e determina um novo tamanho da imagem final
//     Golang Sintaxe: platform.DrawImage(img, x, y, width, height)
//
//     Corta um pedaço da imagem e determina uma nova posição e tamanho para a imagem
//     final
//     Golang Sintaxe: platform.drawImage(img, sx, sy, sWidth, sHeight, x, y, width,
//                     height)
func (el *Canvas) DrawImage(image interface{}, value ...float64) {

	if len(value) == 8 {
		sx := value[0]
		sy := value[1]
		sWidth := value[2]
		sHeight := value[3]
		x := value[4]
		y := value[5]
		width := value[6]
		height := value[7]

		el.SelfContext.Call("drawImage", image.(js.Value), sx, sy, sWidth, sHeight, x, y, width, height)
	} else if len(value) == 4 {
		x := value[0]
		y := value[1]
		width := value[2]
		height := value[3]

		el.SelfContext.Call("drawImage", image.(js.Value), x, y, width, height)
	} else if len(value) == 2 {
		x := value[0]
		y := value[1]

		el.SelfContext.Call("drawImage", image.(js.Value), x, y)
	} else {
		log.Fatalf("canvas.drawImage must be canvas.drawImage(image, sx, sy, sWidth, sHeight, x, y, width, height), canvas.drawImage(image, x, y, width, height) or canvas.drawImage(image, x, y)")
	}
}
