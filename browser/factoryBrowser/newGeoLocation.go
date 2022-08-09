package factoryBrowser

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/geolocation"
)

// NewGeoLocation
//
// English:
//
// The Geolocation interface represents an object able to obtain the position of the device programmatically. It gives
// Web content access to the location of the device. This allows a website or app to offer customized results based on
// the user's location.
//
// An object with this interface is obtained using the navigator.geolocation property implemented by the Navigator
// object.
//
//	Notes:
//	 * For security reasons, when a web page tries to access location information, the user is notified and asked to
//	   grant permission. Be aware that each browser has its own policies and methods for requesting this permission.
//	 * This feature is available only in secure contexts (HTTPS), in some or all supporting browsers.
//
// Português:
//
// A interface Geolocation representa um objeto capaz de obter a posição do dispositivo programaticamente.
// Dá acesso ao conteúdo da Web à localização do dispositivo. Isso permite que um site ou aplicativo ofereça resultados
// personalizados com base na localização do usuário.
//
// An object with this interface is obtained using the navigator.geolocation property implemented by the Navigator
// object.
//
//	Notes:
//	 * Por motivos de segurança, quando uma página da Web tenta acessar informações de localização, o usuário é
//	   notificado e solicitado a conceder permissão. Esteja ciente de que cada navegador tem suas próprias políticas e
//	   métodos para solicitar essa permissão.
//	 * Este recurso está disponível apenas em contextos seguros (HTTPS), em alguns ou em todos os navegadores compatíveis.
func NewGeoLocation() (ref *geolocation.GeoLocation) {
	return new(geolocation.GeoLocation)
}
