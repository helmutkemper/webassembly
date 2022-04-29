package html

import (
	"strconv"
	"strings"
)

type Font struct {
	//Style
	//
	// English:
	//
	//  Specifies the font style.
	//
	// Português:
	//
	//  Especifica o estilo da fonte.
	Style FontStyleRule

	// Variant
	//
	// English:
	//
	//  Specifies the font variant.
	//
	// Português:
	//
	//  Especifica a variante da fonte.
	Variant FontVariantRule

	// Weight
	//
	// English:
	//
	//  Specifies the font weight.
	//
	// Português:
	//
	//  Especifica o peso da fonte.
	Weight FontWeightRule

	// Size
	//
	// English:
	//
	//  Specifies the font size and the line-height, in pixels.
	//
	// Português:
	//
	//  Especifica o tamanho da fonte e a altura da linha, em pixels.
	Size int

	// Family
	//
	// English:
	//
	//  Specifies the font family.
	//
	// Português:
	//
	//  Especifica a família de fontes.
	Family string

	// Caption
	//
	// English:
	//
	//  Use the font captioned controls (like buttons, drop-downs, etc.)
	//
	// Português:
	//
	//  Use os controles legendados de fonte (como botões, menus suspensos etc.)
	Caption string

	// Icon
	//
	// English:
	//
	//  Use the font used to label icons.
	//
	// Português:
	//
	//  Use a fonte usada para rotular os ícones.
	Icon string

	// Menu
	//
	// English:
	//
	//  Use the font used in menus (drop-down menus and menu lists).
	//
	// Português:
	//
	//  Use a fonte usada nos menus (menus suspensos e listas de menus).
	Menu string

	// MessageBox
	//
	// English:
	//
	//  Use the font used in dialog boxes.
	//
	// Português:
	//
	//  Use a fonte usada nas caixas de diálogo.
	MessageBox string

	// SmallCaption
	//
	// English:
	//
	//  Use the font used for labeling small controls.
	//
	// Português:
	//
	//  Use a fonte usada para rotular pequenos controles.
	SmallCaption string

	// StatusBar
	//
	// English:
	//
	//  Use the fonts used in window status bar.
	//
	// Português:
	//
	//  Use as fontes usadas na barra de status da janela.
	StatusBar string
}

func (e *Font) String() string {
	var ret string
	if e.Style != "" {
		ret += e.Style.String()
		ret += " "
	}

	if e.Variant != "" {
		ret += e.Variant.String()
		ret += " "
	}

	if e.Weight != "" {
		ret += e.Weight.String()
		ret += " "
	}

	if e.Size != 0 {
		ret += strconv.FormatInt(int64(e.Size), 10)
		ret += "px "
	}

	if e.Family != "" {
		ret += e.Family
		ret += " "
	}

	if e.Caption != "" {
		ret += e.Caption
		ret += " "
	}

	if e.Icon != "" {
		ret += e.Icon
		ret += " "
	}

	if e.Menu != "" {
		ret += e.Menu
		ret += " "
	}

	if e.MessageBox != "" {
		ret += e.MessageBox
		ret += " "
	}

	if e.SmallCaption != "" {
		ret += e.SmallCaption
		ret += " "
	}

	if e.StatusBar != "" {
		ret += e.StatusBar
		ret += " "
	}

	return strings.TrimSpace(ret)
}
