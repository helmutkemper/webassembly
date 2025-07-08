package devices

import "syscall/js"

type Icon interface {
	GetIconName() (name string)
	GetIconCategory() (name string)
	GetIcon() (icon []js.Value)
}
