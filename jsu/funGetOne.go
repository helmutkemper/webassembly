package jsu

import "syscall/js"

// OR
//
// English:
//
// This function is like js functionality: preview.captureStream = preview.captureStream || preview.mozCaptureStream;
//
// Português:
//
// Esta função é como a funcionalidade js: preview.captureStream = preview.captureStream || preview.mozCaptureStream;
func OR(list ...js.Value) (notNull js.Value) {
	for _, v := range list {
		if !(v.IsNull() == true || v.IsUndefined() == true) {
			return v
		}
	}

	return
}
