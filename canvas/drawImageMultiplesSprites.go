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
//     This method is based on book Eloquent JavaScript from Marijn Haverbeke
//     chapter 17 - Canvas, https://eloquentjavascript.net/index.html
//     Thanks Marijn!
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
//     Este método é baseado no livro Eloquent JavaScript de Marijn Haverbeke
//     capítulo 17 - Canvas, https://eloquentjavascript.net/index.html
//     Obrigado Marijn!
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
func (el *Canvas) DrawImageMultiplesSprites(image interface{}, spriteWidth, spriteHeight, spriteFirstElementIndex, spriteLastElementIndex int, spriteChangeInterval time.Duration, x, y, width, height, lifeCycleLimit, lifeCycleRepeatLimit int, lifeCycleRepeatInterval time.Duration) {
	previousBackgroundImageData := el.SelfContext.Call("getImageData", x, y, width, height)
	go threadDrawImageMultiplesSprites(el, image, previousBackgroundImageData, spriteWidth, spriteHeight, spriteFirstElementIndex, spriteLastElementIndex, spriteChangeInterval, x, y, width, height, lifeCycleLimit, lifeCycleRepeatLimit, 1, lifeCycleRepeatInterval)
}

func threadDrawImageMultiplesSprites(el *Canvas, image, previousBackgroundImageData interface{}, spriteWidth, spriteHeight, spriteFirstElementIndex, spriteLastElementIndex int, spriteChangeInterval time.Duration, x, y, width, height, lifeCycleLimit, lifeCycleRepeatLimit, lifeCycleRepeatLimitCounter int, lifeCycleRepeatInterval time.Duration) {
	var cycle = spriteFirstElementIndex
	var lifeCycle = 0

	ticker := time.NewTicker(spriteChangeInterval)

	el.SelfContext.Call("clearRect", x, y, width, height)
	el.SelfContext.Call("putImageData", previousBackgroundImageData, x, y)
	el.SelfContext.Call("drawImage", image.(js.Value), cycle*spriteWidth, 0, spriteWidth, spriteHeight, x, y, width, height)

	for {
		select {
		case <-ticker.C:

			if cycle < spriteLastElementIndex {
				cycle += 1
			} else {
				cycle = spriteFirstElementIndex
				lifeCycle += 1
			}

			el.SelfContext.Call("clearRect", x, y, width, height)
			el.SelfContext.Call("putImageData", previousBackgroundImageData, x, y)
			el.SelfContext.Call("drawImage", image.(js.Value), cycle*spriteWidth, 0, spriteWidth, spriteHeight, x, y, width, height)

			if lifeCycleLimit != 0 && lifeCycleLimit == lifeCycle {
				if lifeCycleRepeatInterval != 0 && lifeCycleRepeatLimit == 0 || lifeCycleRepeatLimit != 0 && lifeCycleRepeatLimit != lifeCycleRepeatLimitCounter {

					go func() {
						time.Sleep(lifeCycleRepeatInterval)
						threadDrawImageMultiplesSprites(el, image, previousBackgroundImageData, spriteWidth, spriteHeight, spriteFirstElementIndex,
							spriteLastElementIndex, spriteChangeInterval, x, y, width, height, lifeCycleLimit, lifeCycleRepeatLimit,
							lifeCycleRepeatLimitCounter, lifeCycleRepeatInterval)
					}()

					lifeCycleRepeatLimitCounter += 1
				}
				return

			}
		}
	}
}
