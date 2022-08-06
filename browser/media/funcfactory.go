package media

func NewFactory() (ref *FactoryConfig) {
	ref = new(FactoryConfig)
	ref.Init()
	return
}

type FactoryConfig struct {
	config map[string]any
}

func (e *FactoryConfig) Init() {
	e.config = make(map[string]any)
}

func (e FactoryConfig) Get() (config map[string]any) {
	return e.config
}

func (e *FactoryConfig) DefaultAudio() (ref *FactoryConfig) {
	e.config = map[string]any{
		"audio": true,
	}

	return e
}

func (e *FactoryConfig) DefaultVideo() (ref *FactoryConfig) {
	e.config = map[string]any{
		"video": true,
	}

	return e
}

// AudioAutoGainControl
//
// English:
//
// Specifies whether automatic gain control is preferred and/or required.
//
//	Input:
//	  value: automatic gain control on/off.
//
// Português:
//
// Especifica se o controle automático de ganho é preferencial e ou obrigatório.
//
//	Entrada:
//	  value: controle automático de ganho ligado/desligado.
func (e *FactoryConfig) AudioAutoGainControl(value bool) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["autoGainControl"] = value
	return e
}

// AudioAutoGainControlExact
//
// English:
//
// Specifies whether automatic gain control is preferred and/or required.
//
//	Input:
//	  value: automatic gain control on/off, exact value.
//
// Português:
//
// Especifica se o controle automático de ganho é preferencial e ou obrigatório.
//
//	Entrada:
//	  value: controle automático de ganho ligado/desligado, valor exato.
func (e *FactoryConfig) AudioAutoGainControlExact(value bool) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["autoGainControl"] = map[string]any{"exact": value}
	return e
}

// AudioAutoGainControlIdeal
//
// English:
//
// Specifies whether automatic gain control is preferred and/or required.
//
//	Input:
//	  value: automatic gain control on/off, ideal value.
//
// Português:
//
// Especifica se o controle automático de ganho é preferencial e ou obrigatório.
//
//	Entrada:
//	  value: controle automático de ganho ligado/desligado, valor ideal.
func (e *FactoryConfig) AudioAutoGainControlIdeal(value bool) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["autoGainControl"] = map[string]any{"ideal": value}
	return e
}

// AudioChannelCount
//
// English:
//
// Specify the channel count or range of channel counts which are acceptable and/or required.
//
//	Input:
//	  value: number of channel count acceptable and/or required.
//
// Português:
//
// Especifica a contagem de canais ou intervalo de contagens de canais que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  value: total de canais aceitáveis ou requeridos.
func (e *FactoryConfig) AudioChannelCount(value int) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["channelCount"] = value
	return e
}

// AudioChannelCountExact
//
// English:
//
// Specify the channel count or range of channel counts which are acceptable and/or required.
//
//	Input:
//	  value: number of channel count acceptable and/or required, exact value.
//
// Português:
//
// Especifica a contagem de canais ou intervalo de contagens de canais que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  value: total de canais aceitáveis ou requeridos, valor exato.
func (e *FactoryConfig) AudioChannelCountExact(value int) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["channelCount"] = map[string]any{"exact": value}
	return e
}

// AudioChannelCountOptions
//
// English:
//
// Specify the channel count or range of channel counts which are acceptable and/or required.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Especifica a contagem de canais ou intervalo de contagens de canais que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) AudioChannelCountOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["audio"].(map[string]any)["channelCount"] = options
	return e
}

// AudioEchoCancellation
//
// English:
//
// Specify whether or not echo cancellation is preferred and/or required.
//
//	Input:
//	  value: echo cancellation on/off
//
// Português:
//
// Especifica se o cancelamento de eco é preferencial ou obrigatório ou não.
//
//	Entrada:
//	  value: cancelamento de eco ligado/desligado.
func (e *FactoryConfig) AudioEchoCancellation(value bool) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["echoCancellation"] = value
	return e
}

// AudioEchoCancellationExact
//
// English:
//
// Specify whether or not echo cancellation is preferred and/or required.
//
//	Input:
//	  value: echo cancellation on/off, exact value.
//
// Português:
//
// Especifica se o cancelamento de eco é preferencial ou obrigatório ou não.
//
//	Entrada:
//	  value: cancelamento de eco ligado/desligado, valor exato.
func (e *FactoryConfig) AudioEchoCancellationExact(value bool) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["echoCancellation"] = map[string]any{"exact": value}
	return e
}

// AudioEchoCancellationIdeal
//
// English:
//
// Specify whether or not echo cancellation is preferred and/or required.
//
//	Input:
//	  value: echo cancellation on/off, ideal value.
//
// Português:
//
// Especifica se o cancelamento de eco é preferencial ou obrigatório ou não.
//
//	Entrada:
//	  value: cancelamento de eco ligado/desligado, valor ideal.
func (e *FactoryConfig) AudioEchoCancellationIdeal(value bool) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["echoCancellation"] = map[string]any{"ideal": value}
	return e
}

// AudioLatency
//
// English:
//
// Specify the latency or range of latencies which are acceptable and/or required.
//
//	Input:
//	  value: latency which are acceptable and/or required.
//
// Português:
//
// Especifica a latência ou intervalo de latências que são aceitáveis e/ou necessárias.
//
//	Entrada:
//	  value: latência que são aceitáveis e ou necessárias.
func (e *FactoryConfig) AudioLatency(value float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["latency"] = value
	return e
}

// AudioLatencyExact
//
// English:
//
// Specify the latency or range of latencies which are acceptable and/or required.
//
//	Input:
//	  value: latency which are acceptable and/or required, exact value.
//
// Português:
//
// Especifica a latência ou intervalo de latências que são aceitáveis e/ou necessárias.
//
//	Entrada:
//	  value: latência que são aceitáveis e ou necessárias, valor exato.
func (e *FactoryConfig) AudioLatencyExact(value float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["latency"] = map[string]any{"exact": value}
	return e
}

// AudioLatencyOptions
//
// English:
//
// Specify the latency or range of latencies which are acceptable and/or required.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Especifica a latência ou intervalo de latências que são aceitáveis e/ou necessárias.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) AudioLatencyOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["audio"].(map[string]any)["latency"] = options
	return e
}

// AudioNoiseSuppression
//
// English:
//
// Specifies whether noise suppression is preferred and/or required.
//
//	Input:
//	  value: noise suppression on/off.
//
// Português:
//
// Especifica se a supressão de ruído é preferida e ou necessária.
//
//	Entrada:
//	  value: supressão de ruído ligada/desligada.
func (e *FactoryConfig) AudioNoiseSuppression(value bool) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["noiseSuppression"] = value
	return e
}

// AudioNoiseSuppressionExact
//
// English:
//
// Specifies whether noise suppression is preferred and/or required.
//
//	Input:
//	  value: noise suppression on/off, exact value.
//
// Português:
//
// Especifica se a supressão de ruído é preferida e ou necessária.
//
//	Entrada:
//	  value: supressão de ruído ligada/desligada, valor exato.
func (e *FactoryConfig) AudioNoiseSuppressionExact(value bool) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["noiseSuppression"] = map[string]any{"exact": value}
	return e
}

// AudioNoiseSuppressionIdeal
//
// English:
//
// Specifies whether noise suppression is preferred and/or required.
//
//	Input:
//	  value: noise suppression on/off, ideal value.
//
// Português:
//
// Especifica se a supressão de ruído é preferida e ou necessária.
//
//	Entrada:
//	  value: supressão de ruído ligada/desligada, valor ideal.
func (e *FactoryConfig) AudioNoiseSuppressionIdeal(value bool) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["noiseSuppression"] = map[string]any{"ideal": value}
	return e
}

// AudioSampleRate
//
// English:
//
// Specify the sample rate or range of sample rates which are acceptable and/or required.
//
//	Input:
//	  value: the sample rate value.
//
// Português:
//
// Especifica a taxa de amostragem ou intervalo de taxas de amostragem que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  value: valor da taxa de amostragem.
func (e *FactoryConfig) AudioSampleRate(value float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["sampleRate"] = value
	return e
}

// AudioSampleRateExact
//
// English:
//
// Specify the sample rate or range of sample rates which are acceptable and/or required.
//
//	Input:
//	  value: the sample rate, exact value.
//
// Português:
//
// Especifica a taxa de amostragem ou intervalo de taxas de amostragem que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  value: valor exato da taxa de amostragem.
func (e *FactoryConfig) AudioSampleRateExact(value float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["sampleRate"] = map[string]any{"exact": value}
	return e
}

// AudioSampleRateOptions
//
// English:
//
// Specify the sample rate or range of sample rates which are acceptable and/or required.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Especifica a taxa de amostragem ou intervalo de taxas de amostragem que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) AudioSampleRateOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["audio"].(map[string]any)["sampleRate"] = options
	return e
}

// AudioSampleSize
//
// English:
//
// Specify the sample size or range of sample sizes which are acceptable and/or required.
//
//	Input:
//	  value: the value of sample size
//
// Português:
//
// Especifica o tamanho da amostra ou intervalo de tamanhos de amostra que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  value: o valor do tamanho da amostra.
func (e *FactoryConfig) AudioSampleSize(value float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["sampleSize"] = value
	return e
}

// AudioSampleSizeExact
//
// English:
//
// Specify the sample size or range of sample sizes which are acceptable and/or required.
//
//	Input:
//	  value: the exact value of sample size.
//
// Português:
//
// Especifica o tamanho da amostra ou intervalo de tamanhos de amostra que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  value: o valor exato do tamanho da amostra.
func (e *FactoryConfig) AudioSampleSizeExact(value float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["sampleSize"] = map[string]any{"exact": value}
	return e
}

// AudioSampleSizeOptions
//
// English:
//
// Specify the sample size or range of sample sizes which are acceptable and/or required.
//
//	Input:
//	  value: the value of sample size
//
// Português:
//
// Especifica o tamanho da amostra ou intervalo de tamanhos de amostra que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  value: o valor do tamanho da amostra.
func (e *FactoryConfig) AudioSampleSizeOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["audio"].(map[string]any)["sampleSize"] = options
	return e
}

// AudioVolume
//
// English:
//
// Specify the volume or range of volumes which are acceptable and/or required.
//
//	Input:
//	  value: the volume value.
//
// Português:
//
// Especifica o volume ou intervalo de volumes que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  value: o valor do volume.
func (e *FactoryConfig) AudioVolume(value float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["volume"] = value
	return e
}

// AudioVolumeExact
//
// English:
//
// Specify the volume or range of volumes which are acceptable and/or required.
//
//	Input:
//	  value: the exact volume value.
//
// Português:
//
// Especifica o volume ou intervalo de volumes que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  value: o valor exato do volume.
func (e *FactoryConfig) AudioVolumeExact(value float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	e.config["audio"].(map[string]any)["volume"] = map[string]any{"exact": value}
	return e
}

// AudioVolumeOptions
//
// English:
//
// Specify the volume or range of volumes which are acceptable and/or required.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Especifica o volume ou intervalo de volumes que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) AudioVolumeOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["audio"] == nil {
		e.config["audio"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["audio"].(map[string]any)["volume"] = options
	return e
}

// ImageWhiteBalanceMode
//
// English:
//
// A const Specify one of KWhiteBalanceModeNone, KWhiteBalanceModeManual, KWhiteBalanceModeSingleShot, or
// KWhiteBalanceModeContinuous.
//
// Português:
//
// Uma constante especificando KWhiteBalanceModeNone, KWhiteBalanceModeManual, KWhiteBalanceModeSingleShot ou
// KWhiteBalanceModeContinuous.
func (e *FactoryConfig) ImageWhiteBalanceMode(value WhiteBalanceMode) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["whiteBalanceMode"] = value.String()

	return e
}

// ImageExposureMode
//
// English:
//
// A const Specify one of KExposureModeNone, KExposureModeManual, KExposureModeSingleShot, or
// KExposureModeContinuous.
//
// Português:
//
// Uma constante especificando um de KExposureModeNone, KExposureModeManual, KExposureModeSingleShot ou
// KExposureModeContinuous.
func (e *FactoryConfig) ImageExposureMode(value ExposureMode) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["exposureMode"] = value.String()

	return e
}

// ImageFocusMode
//
// English:
//
// A const Specify one of KFocusModeNone, KFocusModeManual, KFocusModeSingleShot, or KFocusModeContinuous.
//
// Português:
//
// Uma constante especificando um de KFocusModeNone, KFocusModeManual, KFocusModeSingleShot ou KFocusModeContinuous.
func (e *FactoryConfig) ImageFocusMode(value FocusMode) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["focusMode"] = value.String()

	return e
}

// ImagePointsOfInterest
//
// English:
//
// The pixel coordinates on the sensor of one or more points of interest.
//
//	Input:
//	  value: The pixel coordinates on the sensor
//	    map[string]any{"x": valueOfX, "y": valueOfY}
//	    []any{map[string]any{"x": valueOfX, "y": valueOfY}, map[string]any{"x": valueOfX, "y": valueOfY}}
//
// This is either an object in the form { x:value, y:value } or an array of such objects, where value is a
// double-precision integer.
//
// Português:
//
// O pixel coordena no sensor de um ou mais pontos de interesse.
//
//	Entrada:
//	  value: As coordenadas do sensor em pexels
//	    map[string]any{"x": valorDeX, "y": valorDeY}
//	    []any{map[string]any{"x": valorDeX, "y": valorDeY}, map[string]any{"x": valorDeX, "y": valorDeY}}
//
// Este é um objeto no formato { x:value, y:value } ou uma matriz de tais objetos, onde value é um inteiro de precisão
// dupla.
func (e *FactoryConfig) ImagePointsOfInterest(value any) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["pointsOfInterest"] = value

	return e
}

// ImageExposureCompensation
//
// English:
//
// Specify f-stop adjustment by up to ±3.
//
//	Input:
//	  value: f-stop adjustment by up to ±3.
//
// Português:
//
// Especifica o ajuste f-stop em até ±3.
//
//	Entrada:
//	  value: ajuste f-stop em até ±3.
func (e *FactoryConfig) ImageExposureCompensation(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["exposureCompensation"] = value

	return e
}

// ImageExposureCompensationExact
//
// English:
//
// Specify f-stop adjustment by up to ±3.
//
//	Input:
//	  value: f-stop adjustment by up to ±3, exact value.
//
// Português:
//
// Especifica o ajuste f-stop em até ±3.
//
//	Entrada:
//	  value: ajuste f-stop em até ±3, valor exato.
func (e *FactoryConfig) ImageExposureCompensationExact(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["exposureCompensation"] = map[string]any{"exact": value}

	return e
}

// ImageExposureCompensationOptions
//
// English:
//
// Specify f-stop adjustment by up to ±3.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Especifica o ajuste f-stop em até ±3.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) ImageExposureCompensationOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["exposureCompensation"] = options

	return e
}

// ImageColorTemperature
//
// English:
//
// Specify a desired color temperature in degrees kelvin.
//
//	Input:
//	  value: color temperature in degrees kelvin.
//
// Português:
//
// Especifica uma temperatura de cor desejada em graus kelvin.
//
//	Entrada:
//	  value: temperatura de cor desejada em graus kelvin.
func (e *FactoryConfig) ImageColorTemperature(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["colorTemperature"] = value

	return e
}

// ImageColorTemperatureExact
//
// English:
//
// Specify a desired color temperature in degrees kelvin.
//
//	Input:
//	  value: Exact value of color temperature in degrees kelvin.
//
// Português:
//
// Especifica uma temperatura de cor desejada em graus kelvin.
//
//	Entrada:
//	  value: Valor exato da temperatura de cor desejada em graus kelvin.
func (e *FactoryConfig) ImageColorTemperatureExact(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["colorTemperature"] = map[string]any{"exact": value}

	return e
}

// ImageColorTemperatureOptions
//
// English:
//
// Specify a desired color temperature in degrees kelvin.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Especifica uma temperatura de cor desejada em graus kelvin.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) ImageColorTemperatureOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["colorTemperature"] = options

	return e
}

// ImageIso
//
// English:
//
// Specify a desired iso setting.
//
//	Input:
//	  value: a desired iso setting.
//
// Português:
//
// Especifica uma configuração iso desejada.
//
//	Entrada:
//	  value: uma configuração iso desejada.
func (e *FactoryConfig) ImageIso(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["iso"] = value

	return e
}

// ImageIsoExact
//
// English:
//
// Specify a desired iso setting.
//
//	Input:
//	  value: an exact value of desired iso setting.
//
// Português:
//
// Especifica uma configuração iso desejada.
//
//	Entrada:
//	  value: um valor exato da configuração iso desejada.
func (e *FactoryConfig) ImageIsoExact(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["iso"] = map[string]any{"exact": value}

	return e
}

// ImageIsoOptions
//
// English:
//
// Specify a desired iso setting.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Especifica uma configuração iso desejada.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) ImageIsoOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["iso"] = options

	return e
}

// ImageBrightness
//
// English:
//
// Specify a desired brightness setting.
//
//	Input:
//	  value: a desired brightness setting.
//
// Português:
//
// Especifica uma configuração de brilho desejada.
//
//	Entrada:
//	  value: a configuração de brilho desejada.
func (e *FactoryConfig) ImageBrightness(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["brightness"] = value

	return e
}

// ImageBrightnessExact
//
// English:
//
// Specify a desired brightness setting.
//
//	Input:
//	  value: an exact value of desired brightness setting.
//
// Português:
//
// Especifica uma configuração de brilho desejada.
//
//	Entrada:
//	  value: o valor exato da configuração de brilho desejada.
func (e *FactoryConfig) ImageBrightnessExact(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["brightness"] = map[string]any{"exact": value}

	return e
}

// ImageBrightnessOptions
//
// English:
//
// A ConstrainDouble specifying a desired brightness setting.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Um ConstrainDouble especificando uma configuração de brilho desejada.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) ImageBrightnessOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["brightness"] = options

	return e
}

// ImageContrast
//
// English:
//
// Specify the degree of difference between light and dark.
//
//	Input:
//	  value: the degree of difference between light and dark.
//
// Português:
//
// Especifica o grau de diferença entre claro e escuro.
//
//	Entrada:
//	  value: o grau de diferença entre claro e escuro.
func (e *FactoryConfig) ImageContrast(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["contrast"] = value

	return e
}

// ImageContrastExact
//
// English:
//
// Specify the degree of difference between light and dark.
//
//	Input:
//	  value: the degree of difference between light and dark.
//
// Português:
//
// Especifica o grau de diferença entre claro e escuro.
//
//	Entrada:
//	  value: o grau de diferença entre claro e escuro.
func (e *FactoryConfig) ImageContrastExact(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["contrast"] = map[string]any{"exact": value}

	return e
}

// ImageContrastOptions
//
// English:
//
// Specify the degree of difference between light and dark.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Especifica o grau de diferença entre claro e escuro.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) ImageContrastOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["contrast"] = options

	return e
}

// ImageSaturation
//
// English:
//
// Specify the degree of color intensity.
//
//	Input:
//	  value: the degree of color intensity.
//
// Português:
//
// Especifica o grau de intensidade da cor.
//
//	Entrada:
//	  value: o grau de intensidade da cor.
func (e *FactoryConfig) ImageSaturation(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["saturation"] = value

	return e
}

// ImageSaturationExact
//
// English:
//
// Specify the degree of color intensity.
//
//	Input:
//	  value: the exact degree of color intensity.
//
// Português:
//
// Especifica o grau de intensidade da cor.
//
//	Entrada:
//	  value: o grau exato de intensidade da cor.
func (e *FactoryConfig) ImageSaturationExact(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["saturation"] = map[string]any{"exact": value}

	return e
}

// ImageSaturationOptions
//
// English:
//
// Specify the degree of color intensity.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Especifica o grau de intensidade da cor.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) ImageSaturationOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["saturation"] = options

	return e
}

// ImageSharpness
//
// English:
//
// Specify the intensity of edges.
//
//	Input:
//	  value: the intensity of edges.
//
// Português:
//
// Especifica à intensidade das arestas.
//
//	Entrada:
//	  value: à intensidade das arestas.
func (e *FactoryConfig) ImageSharpness(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["sharpness"] = value

	return e
}

// ImageSharpnessExact
//
// English:
//
// Specify the intensity of edges.
//
//	Input:
//	  value: the intensity of edges.
//
// Português:
//
// Especifica à intensidade das arestas.
//
//	Entrada:
//	  value: à intensidade das arestas.
func (e *FactoryConfig) ImageSharpnessExact(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["sharpness"] = map[string]any{"exact": value}

	return e
}

// ImageSharpnessOptions
//
// English:
//
// Specify the intensity of edges.
//
//	Input:
//	  value: the intensity of edges.
//
// Português:
//
// Especifica à intensidade das arestas.
//
//	Entrada:
//	  value: à intensidade das arestas.
func (e *FactoryConfig) ImageSharpnessOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["sharpness"] = options

	return e
}

// ImageFocusDistance
//
// English:
//
// Specify distance to a focused object.
//
//	Input:
//	  value: distance to a focused object.
//
// Português:
//
// Especifica a distância para um objeto focado.
//
//	Entrada:
//	  value: distância para um objeto focado.
func (e *FactoryConfig) ImageFocusDistance(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["focusDistance"] = value

	return e
}

// ImageFocusDistanceExact
//
// English:
//
// Specify distance to a focused object.
//
//	Input:
//	  value: distance to a focused object.
//
// Português:
//
// Especifica a distância para um objeto focado.
//
//	Entrada:
//	  value: distância para um objeto focado.
func (e *FactoryConfig) ImageFocusDistanceExact(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["focusDistance"] = map[string]any{"exact": value}

	return e
}

// ImageFocusDistanceOptions
//
// English:
//
// Specify distance to a focused object.
//
//	Input:
//	  value: distance to a focused object.
//
// Português:
//
// Especifica a distância para um objeto focado.
//
//	Entrada:
//	  value: distância para um objeto focado.
func (e *FactoryConfig) ImageFocusDistanceOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["focusDistance"] = options

	return e
}

// ImageZoom
//
// English:
//
// Specify the desired focal length.
//
//	Input:
//	  value: the desired focal length.
//
// Português:
//
// Especifica a distância focal desejada.
//
//	Entrada:
//	  value: distância focal desejada.
func (e *FactoryConfig) ImageZoom(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["zoom"] = value

	return e
}

// ImageZoomExact
//
// English:
//
// Specify the desired focal length.
//
//	Input:
//	  value: the exact value of desired focal length.
//
// Português:
//
// Especifica a distância focal desejada.
//
//	Entrada:
//	  value: o valor exato da distância focal desejada.
func (e *FactoryConfig) ImageZoomExact(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["zoom"] = map[string]any{"exact": value}

	return e
}

// ImageZoomOptions
//
// English:
//
// Specify the desired focal length.
//
//	Input:
//	  value: the desired focal length.
//
// Português:
//
// Especifica a distância focal desejada.
//
//	Entrada:
//	  value: distância focal desejada.
func (e *FactoryConfig) ImageZoomOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["zoom"] = options

	return e
}

// ImageTorch
//
// English:
//
// Define whether the fill light is continuously connected, meaning it stays on as long as the track is active.
//
// Português:
//
// Define se a luz de preenchimento está continuamente conectada, o que significa que permanece acesa enquanto a trilha
// estiver ativa.
func (e *FactoryConfig) ImageTorch(value bool) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["torch"] = value
	return e
}

// ImageTorchExact
//
// English:
//
// Define whether the fill light is continuously connected, meaning it stays on as long as the track is active.
//
// Português:
//
// Define se a luz de preenchimento está continuamente conectada, o que significa que permanece acesa enquanto a trilha
// estiver ativa.
func (e *FactoryConfig) ImageTorchExact(value bool) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["torch"] = map[string]any{"exact": value}
	return e
}

// ImageTorchIdeal
//
// English:
//
// Define whether the fill light is continuously connected, meaning it stays on as long as the track is active.
//
// Português:
//
// Define se a luz de preenchimento está continuamente conectada, o que significa que permanece acesa enquanto a trilha
// estiver ativa.
func (e *FactoryConfig) ImageTorchIdeal(value bool) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["torch"] = map[string]any{"ideal": value}
	return e
}

// ImageAspectRatio
//
// English:
//
// Specify the video aspect ratio or range of aspect ratios which are acceptable and/or required.
//
//	Input:
//	  value: the video aspect ratio
//
// Português:
//
// Especifica a proporção do vídeo ou o intervalo de proporções do vídeo que são aceitáveis e/ou obrigatórios.
//
//	Entrada:
//	  value: a proporção do vídeo
func (e *FactoryConfig) ImageAspectRatio(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["aspectRatio"] = value

	return e
}

// ImageAspectRatioExact
//
// English:
//
// Specify the video aspect ratio or range of aspect ratios which are acceptable and/or required.
//
//	Input:
//	  value: the video aspect ratio
//
// Português:
//
// Especifica a proporção do vídeo ou o intervalo de proporções do vídeo que são aceitáveis e/ou obrigatórios.
//
//	Entrada:
//	  value: a proporção do vídeo
func (e *FactoryConfig) ImageAspectRatioExact(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["aspectRatio"] = map[string]any{"exact": value}

	return e
}

// ImageAspectRatioOptions
//
// English:
//
// Specify the video aspect ratio or range of aspect ratios which are acceptable and/or required.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Especifica a proporção do vídeo ou o intervalo de proporções do vídeo que são aceitáveis e/ou obrigatórios.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) ImageAspectRatioOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["aspectRatio"] = options

	return e
}

// ImageFacingModeExact
//
// English:
//
// Specify the video aspect ratio or range of aspect ratios which are acceptable and/or required.
//
//	Input:
//	  value: the video aspect ratio
//
// Português:
//
// Especifica a proporção do vídeo ou o intervalo de proporções do vídeo que são aceitáveis e/ou obrigatórios.
//
//	Entrada:
//	  value: a proporção do vídeo
func (e *FactoryConfig) ImageFacingModeExact(value FacingMode) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["facingMode"] = map[string]any{"exact": value.String()}

	return e
}

// ImageFacingModeIdeal
//
// English:
//
// Specify the video aspect ratio or range of aspect ratios which are acceptable and/or required.
//
//	Input:
//	  value: the video aspect ratio
//
// Português:
//
// Especifica a proporção do vídeo ou o intervalo de proporções do vídeo que são aceitáveis e/ou obrigatórios.
//
//	Entrada:
//	  value: a proporção do vídeo
func (e *FactoryConfig) ImageFacingModeIdeal(value FacingMode) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["facingMode"] = map[string]any{"ideal": value.String()}

	return e
}

// ImageFrameRate
//
// English:
//
// Specify the frame rate or range of frame rates which are acceptable and/or required.
//
//	Input:
//	  value: the frame rate
//
// Português:
//
// Especifica a taxa de quadros ou intervalo de taxas de quadros que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  value: a taxa de quadros
func (e *FactoryConfig) ImageFrameRate(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["frameRate"] = value

	return e
}

// ImageFrameRateExact
//
// English:
//
// Specify the frame rate or range of frame rates which are acceptable and/or required.
//
//	Input:
//	  value: the exact value of frame rate
//
// Português:
//
// Especifica a taxa de quadros ou intervalo de taxas de quadros que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  value: o valor exato da taxa de quadros
func (e *FactoryConfig) ImageFrameRateExact(value float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["frameRate"] = map[string]any{"exact": value}

	return e
}

// ImageFrameRateOptions
//
// English:
//
// Specify the frame rate or range of frame rates which are acceptable and/or required.
//
//	Input:
//	  min: The smallest permissible value. If the value cannot remain equal to or greater than this value, matching
//	    will fail.
//	  max: The largest permissible value.  If the value cannot remain equal to or less than this value, matching
//	    will fail.
//	  ideal: An ideal value. If possible, this value will be used, but if it's not possible, the user agent will use
//	    the closest possible match.
//
//	Notes:
//	  * Use -1 to ignore the property.
//
// Português:
//
// Especifica a taxa de quadros ou intervalo de taxas de quadros que são aceitáveis e/ou necessários.
//
//	Entrada:
//	  min: O menor valor permitido. Se o valor não puder permanecer igual ou maior que esse valor, a correspondência
//	    falhará.
//	  max: O maior valor permitido. Se o valor não puder permanecer igual ou menor que esse valor, a correspondência
//	    falhará.
//	  ideal: Um valor ideal. Se possível, esse valor será usado, mas se não for possível, o agente do usuário usará a
//	    correspondência mais próxima possível.
//
//	Notas:
//	  * Use o valor -1 para ignorar a propriedade.
func (e *FactoryConfig) ImageFrameRateOptions(min, max, ideal float64) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["frameRate"] = options

	return e
}

// ImageHeight
//
// English:
//
// Specify the video height or range of heights which are acceptable and/or required.
//
//	Input:
//	  value: video height
//
// Português:
//
// Especifica à altura do vídeo ou o intervalo de alturas que são aceitáveis e/ou obrigatórios.
//
//	Entrada:
//	  value: altura do vídeo
func (e *FactoryConfig) ImageHeight(value int) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["height"] = value

	return e
}

// ImageHeightExact
//
// English:
//
// Specify the video height or range of heights which are acceptable and/or required.
//
//	Input:
//	  value: video height
//
// Português:
//
// Especifica à altura do vídeo ou o intervalo de alturas que são aceitáveis e/ou obrigatórios.
//
//	Entrada:
//	  value: altura do vídeo
func (e *FactoryConfig) ImageHeightExact(value int) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["height"] = map[string]any{"exact": value}

	return e
}

// ImageHeightOptions
//
// English:
//
// Specify the video height or range of heights which are acceptable and/or required.
//
//	Input:
//	  value: video height
//
// Português:
//
// Especifica à altura do vídeo ou o intervalo de alturas que são aceitáveis e/ou obrigatórios.
//
//	Entrada:
//	  value: altura do vídeo
func (e *FactoryConfig) ImageHeightOptions(min, max, ideal int) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["height"] = options

	return e
}

// ImageWidth
//
// English:
//
// Specify the video width or range of widths which are acceptable and/or required.
//
//	Input:
//	  value: the video width
//
// Português:
//
// Especifica a largura do vídeo ou o intervalo de larguras que são aceitáveis e/ou obrigatórios.
//
//	Entrada:
//	  value: largura do vídeo
func (e *FactoryConfig) ImageWidth(value int) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["width"] = value

	return e
}

// ImageWidthExact
//
// English:
//
// Specify the video width or range of widths which are acceptable and/or required.
//
//	Input:
//	  value: the exact value of video width
//
// Português:
//
// Especifica a largura do vídeo ou o intervalo de larguras que são aceitáveis e/ou obrigatórios.
//
//	Entrada:
//	  value: o valor exato da largura do vídeo
func (e *FactoryConfig) ImageWidthExact(value int) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	e.config["image"].(map[string]any)["width"] = map[string]any{"exact": value}

	return e
}

// ImageWidthOptions
//
// English:
//
// Specify the video width or range of widths which are acceptable and/or required.
//
//	Input:
//	  value: the exact value of video width
//
// Português:
//
// Especifica a largura do vídeo ou o intervalo de larguras que são aceitáveis e/ou obrigatórios.
//
//	Entrada:
//	  value: o valor exato da largura do vídeo
func (e *FactoryConfig) ImageWidthOptions(min, max, ideal int) (ref *FactoryConfig) {
	if e.config["image"] == nil {
		e.config["image"] = make(map[string]any)
	}

	options := make(map[string]any)
	if min > -1 {
		options["min"] = min
	}

	if max > -1 {
		options["max"] = max
	}

	if ideal > -1 {
		options["ideal"] = ideal
	}

	e.config["image"].(map[string]any)["width"] = options

	return e
}

// DeviceIdExact
//
// English:
//
// Specify a device ID or an array of device IDs which are acceptable and/or required.
//
// Português:
//
// Especifica um ID de dispositivo ou uma matriz de IDs de dispositivo que são aceitáveis e/ou obrigatórios.
// https://developer.mozilla.org/en-US/docs/Web/API/MediaTrackConstraints/deviceId
func (e *FactoryConfig) DeviceIdExact(value any) (ref *FactoryConfig) {
	if e.config["media"] == nil {
		e.config["media"] = make(map[string]any)
	}

	e.config["media"].(map[string]any)["deviceId"] = map[string]any{"exact": value}

	return e
}

// DeviceIdIdeal
//
// English:
//
// Specify a device ID or an array of device IDs which are acceptable and/or required.
//
// Português:
//
// Especifica um ID de dispositivo ou uma matriz de IDs de dispositivo que são aceitáveis e/ou obrigatórios.
// https://developer.mozilla.org/en-US/docs/Web/API/MediaTrackConstraints/deviceId
func (e *FactoryConfig) DeviceIdIdeal(value any) (ref *FactoryConfig) {
	if e.config["media"] == nil {
		e.config["media"] = make(map[string]any)
	}

	e.config["media"].(map[string]any)["deviceId"] = map[string]any{"ideal": value}

	return e
}

// GroupIdExact
//
// English:
//
// Specify a group ID or an array of group IDs which are acceptable and/or required.
//
// Português:
//
// Especifica um ID de grupo ou uma matriz de IDs de grupo que são aceitáveis e/ou obrigatórios.
// https://developer.mozilla.org/en-US/docs/Web/API/MediaTrackConstraints/groupId
func (e *FactoryConfig) GroupIdExact(value any) (ref *FactoryConfig) {
	if e.config["media"] == nil {
		e.config["media"] = make(map[string]any)
	}

	e.config["media"].(map[string]any)["groupId"] = map[string]any{"exact": value}

	return e
}

// GroupIdIdeal
//
// English:
//
// Specify a group ID or an array of group IDs which are acceptable and/or required.
//
// Português:
//
// Especifica um ID de grupo ou uma matriz de IDs de grupo que são aceitáveis e/ou obrigatórios.
// https://developer.mozilla.org/en-US/docs/Web/API/MediaTrackConstraints/groupId
func (e *FactoryConfig) GroupIdIdeal(value any) (ref *FactoryConfig) {
	if e.config["media"] == nil {
		e.config["media"] = make(map[string]any)
	}

	e.config["media"].(map[string]any)["groupId"] = map[string]any{"ideal": value}

	return e
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
