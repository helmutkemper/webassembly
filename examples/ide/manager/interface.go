package manager

import (
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"syscall/js"
)

type Icon interface {
	SetDelay(delay float64)
	SetDuration(duration float64)
	//SetX(x rulesDensity.Density)
	GetX() (x int)
	//SetY(y rulesDensity.Density)
	GetY() (y int)
	SetWidth(width rulesDensity.Density)
	GetWidth() (width int)
	SetHeight(height rulesDensity.Density)
	GetHeight() (height int)
	SetSize(size float64)
	SetOpening(statusOpening int)
	Init()

	GetIconName() (name string)
	GetIconCategory() (name string)
	SetStatus(status int)
	GetStatus() (staus int)
	GetIcon() (icon js.Value)
}
