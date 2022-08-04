package media

type ExposureMode string

func (e ExposureMode) String() string {
	return string(e)
}

const (
	KExposureModeNone       ExposureMode = "none"
	KExposureModeManual     ExposureMode = "manual"
	KExposureModeSingleShot ExposureMode = "single-shot"
	KExposureModeContinuous ExposureMode = "continuous"
)
