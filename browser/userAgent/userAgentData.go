package userAgent

import (
	"syscall/js"
)

type Hints string

func (e Hints) String() string {
	return string(e)
}

const (
	KHintsArchitecture    Hints = "architecture"
	KHintsBitness         Hints = "bitness"
	KHintsModel           Hints = "model"
	KHintsPlatformVersion Hints = "platformVersion"
	KHintsFullVersionList Hints = "fullVersionList"
)

type Brand struct {
	Brand   string
	Version string
}

type Data struct {

	// Brands
	//
	// English:
	//
	// Returns an array of objects containing brand and version specifying the browser brand and its version (the same
	// information as provided by NavigatorUAData.brands).
	//
	// Note that this information can be sent to a server in the Sec-CH-UA header (a low-entropy client hint).
	//
	// Português:
	//
	// Retorna uma matriz de objetos contendo marca e versão especificando a marca do navegador e sua versão (as mesmas
	// informações fornecidas por NavigatorUAData.brands).
	//
	// Observe que essas informações podem ser enviadas a um servidor no cabeçalho Sec-CH-UA (uma dica de cliente de
	// baixa entropia).
	Brands []Brand

	// Mobile
	//
	// English:
	//
	// Returns true if the user agent is running on a mobile device (the same information as provided by
	// NavigatorUAData.mobile).
	//
	// Note that this information can be sent to a server in the Sec-CH-UA-Mobile header (a low-entropy client hint).
	//
	// Português:
	//
	// Retorna true se o agente do usuário estiver sendo executado em um dispositivo móvel (as mesmas informações
	// fornecidas pelo NavigatorUAData.mobile).
	//
	// Observe que essas informações podem ser enviadas a um servidor no cabeçalho Sec-CH-UA-Mobile (uma dica de cliente
	// de baixa entropia).
	Mobile bool

	// Platform
	//
	// English:
	//
	// Returns a string describing the platform the user agent is running on, like "Windows" (the same information as
	// provided by NavigatorUAData.platform).
	//
	// Note that this information can be sent to a server in the Sec-CH-UA-Platform header (a low-entropy client hint).
	//
	// Português:
	//
	// Retorna uma string descrevendo a plataforma na qual o agente do usuário está sendo executado, como "Windows" (as
	// mesmas informações fornecidas pelo NavigatorUAData.platform).
	//
	// Observe que essas informações podem ser enviadas para um servidor no cabeçalho Sec-CH-UA-Platform (uma dica de
	// cliente de baixa entropia).
	Platform string

	// Architecture
	//
	// English:
	//
	// A string containing the platform architecture.
	//
	// For example, "x86". Note that this information can be sent to a server in the Sec-CH-UA-Arch header after the
	// server explicitly requests it in the Accept-CH header.
	//
	// Português:
	//
	// Uma string contendo a arquitetura da plataforma.
	//
	// Por exemplo, "x86". Observe que essas informações podem ser enviadas para um servidor no cabeçalho Sec-CH-UA-Arch
	// após o servidor solicitar explicitamente no cabeçalho Accept-CH.
	Architecture string

	// Bitness
	//
	// English:
	//
	// A string containing the architecture bitness. For example, "32" or "64".
	//
	// Note that this information can be sent to a server in the Sec-CH-UA-Bitness header if the server explicitly
	// requests it in the Accept-CH header.
	//
	// Português:
	//
	// Uma string contendo o número de bits da arquitetura.
	//
	// Por exemplo, "32" ou "64". Observe que essas informações podem ser enviadas para um servidor no cabeçalho
	// Sec-CH-UA-Bitness se o servidor solicitar explicitamente no cabeçalho Accept-CH.
	Bitness string

	// Model
	//
	// English:
	//
	// A string containing the model of mobile device. For example, "Pixel 2XL".
	//
	// If device is not a mobile device or if device model is not known, model will be "". Note that this information can
	// be sent to a server in the Sec-CH-UA-Model header if the server explicitly requests it in the Accept-CH header.
	//
	// Português:
	//
	// Uma string contendo o modelo do dispositivo móvel.
	//
	// Por exemplo, "Pixel 2XL". Se o dispositivo não for um dispositivo móvel ou se o modelo do dispositivo não for
	// conhecido, o modelo será "". Observe que essas informações podem ser enviadas para um servidor no cabeçalho
	// Sec-CH-UA-Model se o servidor solicitar explicitamente no cabeçalho Accept-CH.
	Model string

	// PlatformVersion
	//
	// English:
	//
	// A string containing the platform version.
	//
	// Platform name itself is always available as low-entropy hint platform. For example, "10.0". Note that this
	// information can be sent to a server in the Sec-CH-UA-Platform-Version header if the server explicitly requests it
	// in the Accept-CH header.
	//
	// Português:
	//
	// Uma string contendo a versão da plataforma.
	//
	// O próprio nome da plataforma está sempre disponível como plataforma de dica de baixa entropia. Por exemplo, "10,0".
	// Observe que essas informações podem ser enviadas para um servidor no cabeçalho Sec-CH-UA-Platform-Version se o
	// servidor solicitar explicitamente no cabeçalho Accept-CH.
	PlatformVersion string

	// FullVersionList
	//
	// English:
	//
	// An array of objects with properties "brand" and "version" representing the browser name and full version
	// respectively.
	//
	// For example, {"brand": "Google Chrome", "version": "103.0.5060.134"}, {"brand": "Chromium", "version":
	// "103.0.5060.134"}.
	//
	// Please note that one object may intentionally contain invalid information to prevent sites from relying on a fixed
	// list of browsers. Note that this information can be sent to a server in the Sec-CH-UA-Full-Version-List header if
	// the server explicitly requests it in the Accept-CH header.
	//
	// Português:
	//
	// Uma matriz de objetos com propriedades "marca" e "versão" representando o nome do navegador e a versão completa,
	// respectivamente.
	//
	// Por exemplo, {"brand": "Google Chrome", "version": "103.0.5060.134"}, {"brand": "Chromium", "version":
	// "103.0.5060.134"}.
	//
	// Observe que um objeto pode conter intencionalmente informações inválidas para evitar que os sites dependam de uma
	// lista fixa de navegadores. Observe que essas informações podem ser enviadas para um servidor no cabeçalho
	// Sec-CH-UA-Full-Version-List se o servidor solicitar explicitamente no cabeçalho Accept-CH.
	FullVersionList []Brand

	object js.Value
}

// populateBrand
//
// English:
//
// Returns an array of objects containing brand and version specifying the browser brand and its version (the same
// information as provided by NavigatorUAData.brands).
//
// Note that this information can be sent to a server in the Sec-CH-UA header (a low-entropy client hint).
//
// Português:
//
// Retorna uma matriz de objetos contendo marca e versão especificando a marca do navegador e sua versão (as mesmas
// informações fornecidas por NavigatorUAData.brands).
//
// Observe que essas informações podem ser enviadas a um servidor no cabeçalho Sec-CH-UA (uma dica de cliente de
// baixa entropia).
func (e *Data) populateBrand() {
	if e.object.Get("brands").IsNull() || e.object.Get("brands").IsUndefined() {
		return
	}

	l := e.object.Get("brands").Get("length").Int()
	i := 0
	brands := make([]Brand, l)
	forEach := js.FuncOf(func(this js.Value, args []js.Value) any {
		brands[i].Brand = args[0].Get("brand").String()
		brands[i].Version = args[0].Get("version").String()
		i += 1
		return nil
	})

	js.Global().Get("Object").Call("values", e.object.Get("brands")).Call("forEach", forEach)
	e.Brands = brands
}

// populateMobile
//
// English:
//
// Returns true if the user agent is running on a mobile device (the same information as provided by
// NavigatorUAData.mobile).
//
// Note that this information can be sent to a server in the Sec-CH-UA-Mobile header (a low-entropy client hint).
//
// Português:
//
// Retorna true se o agente do usuário estiver sendo executado em um dispositivo móvel (as mesmas informações
// fornecidas pelo NavigatorUAData.mobile).
//
// Observe que essas informações podem ser enviadas a um servidor no cabeçalho Sec-CH-UA-Mobile (uma dica de cliente
// de baixa entropia).
func (e *Data) populateMobile() {
	if e.object.Get("mobile").IsNull() || e.object.Get("mobile").IsUndefined() {
		return
	}

	e.Mobile = e.object.Get("mobile").Bool()
}

// populatePlatform
//
// English:
//
// Returns a string describing the platform the user agent is running on, like "Windows" (the same information as
// provided by NavigatorUAData.platform).
//
// Note that this information can be sent to a server in the Sec-CH-UA-Platform header (a low-entropy client hint).
//
// Português:
//
// Retorna uma string descrevendo a plataforma na qual o agente do usuário está sendo executado, como "Windows" (as
// mesmas informações fornecidas pelo NavigatorUAData.platform).
//
// Observe que essas informações podem ser enviadas para um servidor no cabeçalho Sec-CH-UA-Platform (uma dica de
// cliente de baixa entropia).
func (e *Data) populatePlatform() {
	if e.object.Get("platform").IsNull() || e.object.Get("platform").IsUndefined() {
		return
	}

	e.Platform = e.object.Get("platform").String()
}

// populateArchitecture
//
// English:
//
// A string containing the platform architecture.
//
// For example, "x86". Note that this information can be sent to a server in the Sec-CH-UA-Arch header after the
// server explicitly requests it in the Accept-CH header.
//
// Português:
//
// Uma string contendo a arquitetura da plataforma.
//
// Por exemplo, "x86". Observe que essas informações podem ser enviadas para um servidor no cabeçalho Sec-CH-UA-Arch
// após o servidor solicitar explicitamente no cabeçalho Accept-CH.
func (e *Data) populateArchitecture() {
	if e.object.Get("architecture").IsNull() || e.object.Get("architecture").IsUndefined() {
		return
	}

	e.Architecture = e.object.Get("architecture").String()
}

// populateBitness
//
// English:
//
// A string containing the architecture bitness. For example, "32" or "64".
//
// Note that this information can be sent to a server in the Sec-CH-UA-Bitness header if the server explicitly
// requests it in the Accept-CH header.
//
// Português:
//
// Uma string contendo o número de bits da arquitetura.
//
// Por exemplo, "32" ou "64". Observe que essas informações podem ser enviadas para um servidor no cabeçalho
// Sec-CH-UA-Bitness se o servidor solicitar explicitamente no cabeçalho Accept-CH.
func (e *Data) populateBitness() {
	if e.object.Get("bitness").IsNull() || e.object.Get("bitness").IsUndefined() {
		return
	}

	e.Bitness = e.object.Get("bitness").String()
}

// populateModel
//
// English:
//
// A string containing the model of mobile device. For example, "Pixel 2XL".
//
// If device is not a mobile device or if device model is not known, model will be "". Note that this information can
// be sent to a server in the Sec-CH-UA-Model header if the server explicitly requests it in the Accept-CH header.
//
// Português:
//
// Uma string contendo o modelo do dispositivo móvel.
//
// Por exemplo, "Pixel 2XL". Se o dispositivo não for um dispositivo móvel ou se o modelo do dispositivo não for
// conhecido, o modelo será "". Observe que essas informações podem ser enviadas para um servidor no cabeçalho
// Sec-CH-UA-Model se o servidor solicitar explicitamente no cabeçalho Accept-CH.
func (e *Data) populateModel() {
	if e.object.Get("model").IsNull() || e.object.Get("model").IsUndefined() {
		return
	}

	e.Model = e.object.Get("model").String()
}

// populatePlatformVersion
//
// English:
//
// A string containing the platform version.
//
// Platform name itself is always available as low-entropy hint platform. For example, "10.0". Note that this
// information can be sent to a server in the Sec-CH-UA-Platform-Version header if the server explicitly requests it
// in the Accept-CH header.
//
// Português:
//
// Uma string contendo a versão da plataforma.
//
// O próprio nome da plataforma está sempre disponível como plataforma de dica de baixa entropia. Por exemplo, "10,0".
// Observe que essas informações podem ser enviadas para um servidor no cabeçalho Sec-CH-UA-Platform-Version se o
// servidor solicitar explicitamente no cabeçalho Accept-CH.
func (e *Data) populatePlatformVersion() {
	if e.object.Get("platformVersion").IsNull() || e.object.Get("platformVersion").IsUndefined() {
		return
	}

	e.PlatformVersion = e.object.Get("platformVersion").String()
}

// populateFullVersionList
//
// English:
//
// An array of objects with properties "brand" and "version" representing the browser name and full version
// respectively.
//
// For example, {"brand": "Google Chrome", "version": "103.0.5060.134"}, {"brand": "Chromium", "version":
// "103.0.5060.134"}.
//
// Please note that one object may intentionally contain invalid information to prevent sites from relying on a fixed
// list of browsers. Note that this information can be sent to a server in the Sec-CH-UA-Full-Version-List header if
// the server explicitly requests it in the Accept-CH header.
//
// Português:
//
// Uma matriz de objetos com propriedades "marca" e "versão" representando o nome do navegador e a versão completa,
// respectivamente.
//
// Por exemplo, {"brand": "Google Chrome", "version": "103.0.5060.134"}, {"brand": "Chromium", "version":
// "103.0.5060.134"}.
//
// Observe que um objeto pode conter intencionalmente informações inválidas para evitar que os sites dependam de uma
// lista fixa de navegadores. Observe que essas informações podem ser enviadas para um servidor no cabeçalho
// Sec-CH-UA-Full-Version-List se o servidor solicitar explicitamente no cabeçalho Accept-CH.
func (e *Data) populateFullVersionList() {
	if e.object.Get("fullVersionList").IsNull() || e.object.Get("fullVersionList").IsUndefined() {
		return
	}

	l := e.object.Get("brands").Get("length").Int()
	i := 0
	fvs := make([]Brand, l)
	forEach := js.FuncOf(func(this js.Value, args []js.Value) any {
		fvs[i].Brand = args[0].Get("brand").String()
		fvs[i].Version = args[0].Get("version").String()
		i += 1
		return nil
	})

	js.Global().Get("Object").Call("values", e.object.Get("fullVersionList")).Call("forEach", forEach)
	e.FullVersionList = fvs
}

// GetHighEntropyValues
//
// English:
//
// The GetHighEntropyValues() function with a dictionary object containing the high entropy values the user-agent returns.
//
//	Notes:
//	  * The terms high entropy and low entropy refer to the amount of information these values reveal about the browser. The values returned as properties are deemed low entropy, and unlikely to identify a user. The values returned by NavigatorUAData.getHighEntropyValues() could potentially reveal more information. These values are therefore retrieved via a Promise, allowing time for the browser to request user permission, or make other checks.
//
// Português:
//
// The getHighEntropyValues() method of the NavigatorUAData interface is a Promise that resolves with a dictionary object containing the high entropy values the user-agent returns.
//
//	Notes:
//	  * The terms high entropy and low entropy refer to the amount of information these values reveal about the browser. The values returned as properties are deemed low entropy, and unlikely to identify a user. The values returned by NavigatorUAData.getHighEntropyValues() could potentially reveal more information. These values are therefore retrieved via a Promise, allowing time for the browser to request user permission, or make other checks.
func GetHighEntropyValues(hints ...Hints) (data Data) {
	wait := make(chan struct{})
	success := js.FuncOf(func(this js.Value, args []js.Value) any {
		data = Data{object: args[0]}
		data.populateBrand()
		data.populateMobile()
		data.populatePlatform()
		data.populateArchitecture()
		data.populateBitness()
		data.populateModel()
		data.populatePlatformVersion()
		data.populateFullVersionList()

		wait <- struct{}{}
		return nil
	})

	hintsArr := make([]interface{}, len(hints))
	for k, hint := range hints {
		hintsArr[k] = hint.String()
	}

	js.Global().Get("navigator").Get("userAgentData").Call("getHighEntropyValues", hintsArr).Call("then", success)
	<-wait

	return
}
