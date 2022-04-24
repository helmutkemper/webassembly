package factoryBrowser

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"

func NewTagImage(id, src string, width, height int, waitLoad bool) (ref *html.TagImage) {
	ref = &html.TagImage{}

	ref.CreateElement(html.KTagImg, src, width, height, waitLoad)
	ref.Id(id)

	rc1 := &html.TagCanvas{}
	rc1.CreateElement(html.KTagCanvas, width, height)
	rc1.DrawImage(ref.GetJs())
	rc1.AppendById("stage")
	data := rc1.GetImageData(0, 0, width, height)

	rc2 := &html.TagCanvas{}
	rc2.CreateElement(html.KTagCanvas, width, height)
	rc2.AppendById("stage")
	rc2.PutImageData(data, width, height)

	return ref
}
