package media

type Audio struct {

	// GainControl
	//
	// English:
	//
	// A ConstrainBoolean object which specifies whether automatic gain control is preferred and/or required.
	//
	// Português:
	//
	// Um objeto ConstrainBoolean que especifica se o controle automático de ganho é preferencial e ou obrigatório.
	AutoGainControl ConstrainBoolean //`js:"autoGainControl"`

	// ChannelCount
	//
	// English:
	//
	// A ConstrainULong specifying the channel count or range of channel counts which are acceptable and/or required.
	//
	// Português:
	//
	// A ConstrainULong specifying the channel count or range of channel counts which are acceptable and/or required.
	ChannelCount ConstrainULong `js:"channelCount"`

	// EchoCancellation
	//
	// English:
	//
	// A ConstrainBoolean object specifying whether or not echo cancellation is preferred and/or required.
	//
	// Português:
	//
	// Um objeto ConstrainBoolean especificando se o cancelamento de eco é preferencial ou obrigatório ou não.
	EchoCancellation ConstrainBoolean `js:"echoCancellation"`

	// Latency
	//
	// English:
	//
	// A ConstrainDouble specifying the latency or range of latencies which are acceptable and/or required.
	//
	// Português:
	//
	// Um ConstrainDouble especificando a latência ou intervalo de latências que são aceitáveis e/ou necessárias.
	Latency ConstrainDouble `js:"latency"`

	// NoiseSuppression
	//
	// English:
	//
	// A ConstrainBoolean which specifies whether noise suppression is preferred and/or required.
	//
	// Português:
	//
	// Um ConstrainBoolean que especifica se a supressão de ruído é preferida e ou necessária.
	NoiseSuppression ConstrainBoolean `js:"noiseSuppression"`

	// SampleRate
	//
	// English:
	//
	// A ConstrainULong specifying the sample rate or range of sample rates which are acceptable and/or required.
	//
	// Português:
	//
	// Um ConstrainULong especificando a taxa de amostragem ou intervalo de taxas de amostragem que são aceitáveis e/ou
	// necessários.
	SampleRate ConstrainULong `js:"sampleRate"`

	// SampleSize
	//
	// English:
	//
	// A ConstrainULong specifying the sample size or range of sample sizes which are acceptable and/or required.
	//
	// Português:
	//
	// Um ConstrainULong especificando o tamanho da amostra ou intervalo de tamanhos de amostra que são aceitáveis e/ou
	// necessários.
	SampleSize ConstrainULong `js:"sampleSize"`

	// Volume
	//
	// English:
	//
	// A ConstrainDouble specifying the volume or range of volumes which are acceptable and/or required.
	//
	// Português:
	//
	// Um ConstrainDouble especificando o volume ou intervalo de volumes que são aceitáveis e/ou necessários.
	Volume ConstrainDouble `js:"volume"`
}

func (e *Audio) mountAutoGainControl(audio *map[string]interface{}) {
	if e.AutoGainControl.Exact.IsSet() {
		(*audio)["autoGainControl"] = map[string]interface{}{"exact": e.AutoGainControl.Exact.Bool()}
		return
	}

	if e.AutoGainControl.Ideal.IsSet() {
		(*audio)["autoGainControl"] = map[string]interface{}{"ideal": e.AutoGainControl.Ideal.Bool()}
		return
	}

	if e.AutoGainControl.Value.IsSet() {
		(*audio)["autoGainControl"] = e.AutoGainControl.Value.Bool()
		return
	}
}

func (e *Audio) mountChannelCount(audio *map[string]interface{}) {
	if e.ChannelCount.Exact != nil {
		(*audio)["channelCount"] = map[string]interface{}{"exact": e.ChannelCount.Exact}
		return
	}

	if e.ChannelCount.Max != nil || e.ChannelCount.Min != nil || e.ChannelCount.Ideal != nil {
		(*audio)["channelCount"] = make(map[string]interface{})

		if e.ChannelCount.Max != nil {
			(*audio)["channelCount"].(map[string]interface{})["max"] = e.ChannelCount.Max
		}

		if e.ChannelCount.Min != nil {
			(*audio)["channelCount"].(map[string]interface{})["min"] = e.ChannelCount.Min
		}

		if e.ChannelCount.Ideal != nil {
			(*audio)["channelCount"].(map[string]interface{})["ideal"] = e.ChannelCount.Ideal
		}

		return
	}

	if e.ChannelCount.Value != nil {
		(*audio)["channelCount"] = e.ChannelCount.Value
		return
	}
}

func (e *Audio) mountEchoCancellation(audio *map[string]interface{}) {
	if e.EchoCancellation.Exact.IsSet() {
		(*audio)["echoCancellation"] = map[string]interface{}{"exact": e.EchoCancellation.Exact.Bool()}
		return
	}

	if e.EchoCancellation.Ideal.IsSet() {
		(*audio)["echoCancellation"] = map[string]interface{}{"ideal": e.EchoCancellation.Ideal.Bool()}
		return
	}

	if e.EchoCancellation.Value.IsSet() {
		(*audio)["echoCancellation"] = e.EchoCancellation.Value.Bool()
		return
	}
}

func (e *Audio) mountLatency(audio *map[string]interface{}) {
	if e.Latency.Exact != nil {
		(*audio)["latency"] = map[string]interface{}{"exact": e.Latency.Exact}
		return
	}

	if e.Latency.Max != nil || e.Latency.Min != nil || e.Latency.Ideal != nil {
		(*audio)["latency"] = make(map[string]interface{})

		if e.Latency.Max != nil {
			(*audio)["latency"].(map[string]interface{})["max"] = e.Latency.Max
		}

		if e.Latency.Min != nil {
			(*audio)["latency"].(map[string]interface{})["min"] = e.Latency.Min
		}

		if e.Latency.Ideal != nil {
			(*audio)["latency"].(map[string]interface{})["ideal"] = e.Latency.Ideal
		}

		return
	}

	if e.Latency.Value != nil {
		(*audio)["latency"] = e.Latency.Value
		return
	}
}

func (e *Audio) mountNoiseSuppression(audio *map[string]interface{}) {
	if e.NoiseSuppression.Exact.IsSet() {
		(*audio)["noiseSuppression"] = map[string]interface{}{"exact": e.NoiseSuppression.Exact.Bool()}
		return
	}

	if e.NoiseSuppression.Ideal.IsSet() {
		(*audio)["noiseSuppression"] = map[string]interface{}{"ideal": e.NoiseSuppression.Ideal.Bool()}
		return
	}

	if e.NoiseSuppression.Value.IsSet() {
		(*audio)["noiseSuppression"] = e.NoiseSuppression.Value.Bool()
		return
	}
}

func (e *Audio) mountSampleRate(audio *map[string]interface{}) {
	if e.SampleRate.Exact != nil {
		(*audio)["sampleRate"] = map[string]interface{}{"exact": e.SampleRate.Exact}
		return
	}

	if e.SampleRate.Max != nil || e.SampleRate.Min != nil || e.SampleRate.Ideal != nil {
		(*audio)["sampleRate"] = make(map[string]interface{})

		if e.SampleRate.Max != nil {
			(*audio)["sampleRate"].(map[string]interface{})["max"] = e.SampleRate.Max
		}

		if e.SampleRate.Min != nil {
			(*audio)["sampleRate"].(map[string]interface{})["min"] = e.SampleRate.Min
		}

		if e.SampleRate.Ideal != nil {
			(*audio)["sampleRate"].(map[string]interface{})["ideal"] = e.SampleRate.Ideal
		}

		return
	}

	if e.SampleRate.Value != nil {
		(*audio)["sampleRate"] = e.SampleRate.Value
		return
	}
}

func (e *Audio) mountSampleSize(audio *map[string]interface{}) {
	if e.SampleSize.Exact != nil {
		(*audio)["sampleSize"] = map[string]interface{}{"exact": e.SampleSize.Exact}
		return
	}

	if e.SampleSize.Max != nil || e.SampleSize.Min != nil || e.SampleSize.Ideal != nil {
		(*audio)["sampleSize"] = make(map[string]interface{})

		if e.SampleSize.Max != nil {
			(*audio)["sampleSize"].(map[string]interface{})["max"] = e.SampleSize.Max
		}

		if e.SampleSize.Min != nil {
			(*audio)["sampleSize"].(map[string]interface{})["min"] = e.SampleSize.Min
		}

		if e.SampleSize.Ideal != nil {
			(*audio)["sampleSize"].(map[string]interface{})["ideal"] = e.SampleSize.Ideal
		}

		return
	}

	if e.SampleSize.Value != nil {
		(*audio)["sampleSize"] = e.SampleSize.Value
		return
	}
}

func (e *Audio) mountVolume(audio *map[string]interface{}) {
	if e.Volume.Exact != nil {
		(*audio)["volume"] = map[string]interface{}{"exact": e.Volume.Exact}
		return
	}

	if e.Volume.Max != nil || e.Volume.Min != nil || e.Volume.Ideal != nil {
		(*audio)["volume"] = make(map[string]interface{})

		if e.Volume.Max != nil {
			(*audio)["volume"].(map[string]interface{})["max"] = e.Volume.Max
		}

		if e.Volume.Min != nil {
			(*audio)["volume"].(map[string]interface{})["min"] = e.Volume.Min
		}

		if e.Volume.Ideal != nil {
			(*audio)["volume"].(map[string]interface{})["ideal"] = e.Volume.Ideal
		}

		return
	}

	if e.Volume.Value != nil {
		(*audio)["volume"] = e.Volume.Value
		return
	}
}

func (e *Audio) mount(audio *map[string]interface{}) {
	if *audio == nil {
		*audio = make(map[string]interface{})
	}

	e.mountAutoGainControl(audio)
	e.mountChannelCount(audio)
	e.mountEchoCancellation(audio)
	e.mountLatency(audio)
	e.mountNoiseSuppression(audio)
	e.mountSampleRate(audio)
	e.mountSampleSize(audio)
	e.mountVolume(audio)
}
