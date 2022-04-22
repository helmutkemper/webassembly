package html

import "log"

// SetCharset
//
// English:
//
//  Space-separated character encodings the server accepts. The browser uses them in the order in
//  which they are listed. The default value means the same encoding as the page.
//  (In previous versions of HTML, character encodings could also be delimited by commas.)
//
// Português:
//
//  Codificações de caracteres separados por espaço que o servidor aceita. O navegador os utiliza na
//  ordem em que estão listados. O valor padrão significa a mesma codificação da página.
//  (Nas versões anteriores do HTML, as codificações de caracteres também podiam ser delimitadas
//  por vírgulas.)
func (e *GlobalAttributes) SetCharset(value string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagForm:
	default:
		log.Printf("tag " + e.tag.String() + " does not support accept-charset property")
	}

	e.selfElement.Set("accept-charset", value)
	return e
}
