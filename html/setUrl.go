package html

import "log"

// SetUrl
//
// English:
//
//  A string containing the URL of the handler.
//
// This URL must include %s, as a placeholder that will be replaced with the escaped URL to be
// handled.
//
//   Note:
//     * The handler URL must use the https scheme. Older browsers also supported http.
//
// Português:
//
//  Uma string contendo o URL do manipulador.
//
// Este URL deve incluir %s, como um marcador de posição que será substituído pelo URL de escape a
// ser tratado.
//
//   Nota:
//     * A URL do manipulador deve usar o esquema https. Os navegadores mais antigos também
//       suportavam http.
func (e *GlobalAttributes) SetUrl(url string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagA:
	default:
		log.Printf("tag " + e.tag.String() + " does not support url property")
	}

	e.selfElement.Set("url", url)
	return e
}
