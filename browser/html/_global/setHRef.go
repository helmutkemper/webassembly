package _global

import "log"

// SetHRef
//
// English:
//
//  The URL that the hyperlink points to. Links are not restricted to HTTP-based URLs — they can use
//  any URL scheme supported by browsers:
//    * Sections of a page with fragment URLs;
//    * Pieces of media files with media fragments;
//    * Telephone numbers with tel: URLs;
//    * Email addresses with mailto: URLs;
//    * While web browsers may not support other URL schemes, web sites can with
//      registerProtocolHandler().
//
// Português:
//
//  A URL para a qual o hiperlink aponta. Os links não são restritos a URLs baseados em HTTP — eles
//  podem usar qualquer esquema de URL suportado pelos navegadores:
//    * Seções de uma página com URLs de fragmento;
//    * Pedaços de arquivos de mídia com fragmentos de mídia;
//    * Números de telefone com tel: URLs;
//    * Endereços de e-mail com mailto: URLs;
//    * Embora os navegadores da Web possam não suportar outros esquemas de URL, os sites da Web podem
//      com registerProtocolHandler().
func (e *GlobalAttributes) SetHRef(href string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagA:
	default:
		log.Printf("tag " + e.tag.String() + " does not support href property")
	}

	e.selfElement.Set("href", href)
	return e
}
