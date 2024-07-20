package factoryBrowser

import (
	"github.com/helmutkemper/webassembly/browser/html"
)

func NewSprite(
	canvas *html.TagCanvas,
	imgPath string,
	width, height int,

) (ref *html.Sprite) {

	ref = new(html.Sprite)
	ref.Canvas(canvas)
	ref.Image(imgPath)
	ref.SpriteWidth(width)
	ref.SpriteHeight(height)
	ref.Init()
	return ref
}
