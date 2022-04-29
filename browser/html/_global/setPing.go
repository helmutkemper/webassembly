package _global

import (
	"log"
	"strings"
)

// SetPing
//
// English:
//
//  A space-separated list of URLs. When the link is followed, the browser will send POST requests
//  with the body PING to the URLs. Typically for tracking.
//
// Português:
//
//  Uma lista de URLs separados por espaços. Quando o link for seguido, o navegador enviará
//  solicitações POST com o corpo PING para as URLs. Normalmente para rastreamento.
func (e *GlobalAttributes) SetPing(ping ...string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagA:
	default:
		log.Printf("tag " + e.tag.String() + " does not support ping property")
	}

	e.selfElement.Set("hreflang", strings.Join(ping, " "))
	return e
}
