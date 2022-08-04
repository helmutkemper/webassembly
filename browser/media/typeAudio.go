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
	AutoGainControl ConstrainBoolean `js:"autoGainControl"`

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
