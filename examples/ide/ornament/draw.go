package ornament

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"syscall/js"
)

// Draw Draw the visual elements of the device
type Draw interface {

	// Init Initializes the instance
	Init() (err error)

	// Update Draw the element design
	Update(x, y, width, height rulesDensity.Density) (err error)

	// GetSvg Returns the SVG tag with the element design
	GetSvg() (svg *html.TagSvg)

	// SetWarning sets the visibility of the warning mark
	SetWarning(warning bool)

	SetSelected(selected bool)

	// ToPngResized
	//
	// English:
	//
	//	 Transform the SVG in `PNG Data` to be used with `img.src = pngData`.
	//
	//		Input:
	//		  width and height: Final size or image scale.
	//
	//	   Rules:
	//	     * If the values are lower or equal to 5.0 the image will have its size multiplied by the past value.
	//	         In this case, both values should be less than or equal to 5.0;
	//	     * If the values are greater than 5.0, this will be the size of the image;
	//	     * If the values are equal to zero, the image will have the original size.
	//
	//	   Example:
	//	     js.Global().Get("document").Call("getElementById", "test").Set("src", url)
	//
	// Português:
	//
	//	 Transforma um SVG em `PNG Data` para ser usado com `img.src = pngData`.
	//
	//		Input:
	//		  width and height: Tamanho final ou escala da imagem.
	//
	//	   Regras:
	//	     * Caso os valores sejam menores ou iguais a 5.0 a imagem terá seu tamanho multiplicado pelo valor passado.
	//	         Nesse caso, os dois valores devem ser menores ou iguais a 5.0;
	//	     * Caso os valores sejam maiores do que 5.0, este será o tamanho da imagem;
	//	     * Caso os valores sejam iguais a zero, a imagem terá o tamanho original.
	//
	//	   Exemplo:
	//	     js.Global().Get("document").Call("getElementById", "test").Set("src", url)
	ToPngResized(width, height float64) (pngData js.Value)
}
