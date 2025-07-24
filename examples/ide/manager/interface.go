package manager

import "syscall/js"

type Icon interface {
	GetIconName() (name string)
	GetIconCategory() (name string)
	SetStatus(status int)
	GetStatus() (staus int)
	GetIcon() (icon js.Value)
}
