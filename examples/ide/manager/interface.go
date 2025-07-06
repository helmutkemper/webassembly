package manager

import "syscall/js"

type Icon interface {
	GetIconName() (name string)
	GetIconCategory() (name string)
	GetIcon(disabled bool) (icon js.Value)
}
