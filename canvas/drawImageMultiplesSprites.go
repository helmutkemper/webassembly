package canvas

import (
	"syscall/js"
	"time"
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
func (el *Canvas) DrawImageMultiplesSprites(image interface{}, spriteW, spriteH int) {

	ticker := time.NewTicker(80 * time.Millisecond)

	go func(el *Canvas, image interface{}, spriteW, spriteH int) {
		var cycle = 0

		dataInterface := el.SelfContext.Call("getImageData", 0, 0, spriteW, spriteH)

		el.SelfContext.Call("drawImage", image.(js.Value), cycle*spriteW, 0, spriteW, spriteH, 0, 0, spriteW, spriteH)

		for {
			select {
			case <-ticker.C:
				cycle = (cycle + 1) % 8
				el.SelfContext.Call("clearRect", 0, 0, spriteW, spriteH)
				el.SelfContext.Call("putImageData", dataInterface, 0, 0)

				el.SelfContext.Call("drawImage", image.(js.Value), cycle*spriteW, 0, spriteW, spriteH, 0, 0, spriteW, spriteH)
			}
		}
	}(el, image, spriteW, spriteH)

}
