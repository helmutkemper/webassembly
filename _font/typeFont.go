package font

import (
	"strconv"
)

type Font struct {
	// en: due to the problem of javascript not being able to measure the font height
	// correctly, all fonts are determined in pixels
	//
	// pt_br: devido ao problema do javascript não conseguir medir a altura da fonte
	// de forma correta, todas as fontes são determinadas em pixels
	Size int

	// en: use a color.RGBA{} struct or a color name factory, ex.:
	// factoryColorNames.NewBlack()
	//
	// pt_br: use um struct color.RGBA{} ou uma fabrica com as cores de fontes, ex.:
	// factoryColorNames.NewBlack()
	//Color color.RGBA

	// en: use a string like 'verdana' ou a font family factory, ex.:
	// factoryFontFamily.NewArialBlack()
	//
	// pt_br: use uma string, tipo 'verdana' ou uma fábrica de famílias de fonte, ex.:
	// factoryFontFamily.NewArialBlack()
	Family string

	// en: use a string, like 'bold' or a font style factory, ex.:
	// factoryFontStyle.NewItalic()
	//
	// pt_br: use uma string, tipo 'bold' ou uma fábrica de estilo de fonte, ex.:
	// factoryFontStyle.NewItalic()
	Style string

	// en: use a string 'Small-Caps' or factoryFontVariant.NewSmallCaps()
	//
	// pt_br: use a string 'Small-Caps' ou factoryFontVariant.NewSmallCaps()
	Variant string

	// en: use a string like 'bold' or a font weight factory, ex.:
	// factoryFontWeight.NewBold()
	//
	// pt_br: use uma string, tipo 'bold' ou uma fábrica de peso de fonte, ex.:
	// factoryFontWeight.NewBold()
	Weight string
}

// en: Format the browser canvas font string as w3school rules
//     Note: browser canvas don't use color
//
// pt_br: Formata a string de fonte para elemento canvas no formato do w3school
//     Nota: o elemento canvas do navegador não usa cor
func (el *Font) String() string {
	return el.Style + " " + el.Variant + el.Weight + strconv.FormatInt(int64(el.Size), 10) + "px " + el.Family
}
