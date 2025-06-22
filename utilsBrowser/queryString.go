package utilsBrowser

import (
	"strings"
	"syscall/js"
)

// GetQueryStringParam
//
// English:
//
//	Returns the value of a query string parameter from the current URL.
//	If the key is not found, returns an empty string.
//
// Português:
//
//	Retorna o valor de um parâmetro da query string da URL atual.
//	Se a chave não for encontrada, retorna uma string vazia.
func GetQueryStringParam(key string) string {
	// window.location
	location := js.Global().Get("window").Get("location")

	// location.search => "?key=value&other=123"
	query := location.Get("search").String()

	// Remove the leading '?'
	query = strings.TrimPrefix(query, "?")

	// Split by '&'
	pairs := strings.Split(query, "&")

	// Loop over the parameters
	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 && parts[0] == key {
			return parts[1]
		}
	}

	return ""
}
