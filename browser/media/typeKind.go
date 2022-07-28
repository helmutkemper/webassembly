package media

type Kind string

func (e Kind) String() string {
	return string(e)
}

const (
	KKindVideoInput  Kind = "videoinput"
	KKindAudioInput  Kind = "audioinput"
	KKindAudioOutput Kind = "audiooutput"
)
