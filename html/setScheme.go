package html

import "log"

// SetScheme
//
// English:
//
//  A string containing the permitted scheme for the protocol that the site wishes to handle.
//  For example, you can register to handle SMS text message links by passing the "sms" scheme.
//
// For security reasons, registerProtocolHandler() restricts which schemes can be registered.
//
//   Note:
//     * A custom scheme may be registered as long as:
//       * The custom scheme's name begins with web+
//       * The custom scheme's name includes at least 1 letter after the web+ prefix
//       * The custom scheme has only lowercase ASCII letters in its name.
//
// Português:
//
//  Uma string contendo o esquema permitido para o protocolo que o site deseja manipular. Por exemplo,
//  você pode se registrar para lidar com links de mensagens de texto SMS passando o esquema "sms".
//
// Por motivos de segurança, registerProtocolHandler() restringe quais esquemas podem ser registrados.
//
//   Nota:
//     * A custom scheme may be registered as long as:
//       * O nome dos esquemas personalizados começa com web+
//       * O nome do esquema personalizado inclui pelo menos 1 letra após o prefixo web+
//       * O esquema personalizado tem apenas letras ASCII minúsculas em seu nome.
func (e *GlobalAttributes) SetScheme(scheme Scheme) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagA:
	default:
		log.Printf("tag " + e.tag.String() + " does not support scheme property")
	}

	e.selfElement.Set("scheme", scheme)
	return e
}
