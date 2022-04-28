package _global

import (
	"log"
	"strings"
	"syscall/js"
)

// SetCustomScheme
//
// English:
//
//  A string containing the permitted scheme for the protocol that the site wishes to handle.
//  For example, you can register to handle SMS text message links by passing the "sms" scheme.
//
//   Input:
//     scheme: A string containing the permitted scheme for the protocol that the site wishes to
//             handle. For example, you can register to handle SMS text message links by passing the
//             "sms" scheme;
//     url:    A string containing the URL of the handler. This URL must include %s, as a placeholder
//             that will be replaced with the escaped URL to be handled.
//
//   Secure context:
//     * This feature is available only in secure contexts (HTTPS), in some or all supporting
//       browsers.
//
// The Navigator method registerProtocolHandler() lets websites register their ability to open or
// handle particular URL schemes (aka protocols).
//
// For example, this API lets webmail sites open mailto: URLs, or VoIP sites open tel: URLs.
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
//   Entrada:
//     scheme: Uma string contendo o esquema permitido para o protocolo que o site deseja manipular.
//             Por exemplo, você pode se registrar para lidar com links de mensagens de texto SMS
//             passando o esquema "sms";
//     url:    Uma string contendo o URL do manipulador. Este URL deve incluir %s, como um marcador
//             de posição que será substituído pelo URL de escape a ser tratado.
//
//   Contexto seguro:
//     * Esse recurso está disponível apenas em contextos seguros (HTTPS), em alguns ou em todos os
//       navegadores compatíveis.
//
// O método do Navigator registerProtocolHandler() permite que os sites registrem sua capacidade de
// abrir ou manipular determinados esquemas de URL (também conhecidos como protocolos).
//
// Por exemplo, esta API permite que sites de webmail abram mailto: URLs ou sites VoIP abram tel:
// URLs.
//
// Por motivos de segurança, registerProtocolHandler() restringe quais esquemas podem ser registrados.
//
//   Nota:
//     * A custom scheme may be registered as long as:
//       * O nome dos esquemas personalizados começa com web+
//       * O nome do esquema personalizado inclui pelo menos 1 letra após o prefixo web+
//       * O esquema personalizado tem apenas letras ASCII minúsculas em seu nome.
func (e *GlobalAttributes) SetCustomScheme(scheme, url string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagA:
	default:
		log.Printf("tag " + e.tag.String() + " does not support scheme property")
	}

	if strings.HasPrefix(scheme, "web+") == false {
		log.Printf(e.tag.String() + ".SetCustomScheme().error: scheme must start with 'web+' prefix")
		return
	}

	e.selfElement.Set("scheme", scheme)
	js.Global().Get("navigator").Call("registerProtocolHandler", scheme, url)
	return e
}
