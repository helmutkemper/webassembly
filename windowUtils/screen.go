package windowUtils

import "syscall/js"

// GetScreenSize
//
// English:
//
//	Retrieves the current screen dimensions and returns the width and height in pixels.
//
// Português:
//
//	Recupera as dimensões da tela atual e retorna a largura e a altura em pixels.
func GetScreenSize() (width, height int) {
	window := js.Global()
	width = window.Get("innerWidth").Int()
	height = window.Get("innerHeight").Int()
	return
}
