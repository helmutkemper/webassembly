package _global

import "log"

// SetCapture
//
// English:
//
//  Introduced in the HTML Media Capture specification and valid for the file input type only, the
//  capture attribute defines which media—microphone, video, or camera—should be used to capture a
//  new file for upload with file upload control in supporting scenarios.
//
// Português:
//
//  Introduzido na especificação HTML Media Capture e válido apenas para o tipo de entrada de arquivo,
//  o atributo capture define qual mídia—microfone, vídeo ou câmera—deve ser usada para capturar um
//  novo arquivo para upload com controle de upload de arquivo em cenários de suporte.
func (e *GlobalAttributes) SetCapture(capture string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagInput:
	default:
		log.Printf("tag " + e.tag.String() + " does not support capture property")
	}

	e.selfElement.Set("capture", capture)
	return e
}
