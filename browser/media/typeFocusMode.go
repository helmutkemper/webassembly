package media

type FocusMode string

func (e FocusMode) String() string {
	return string(e)
}

const (
	KFocusModeNone       FocusMode = "none"
	KFocusModeManual     FocusMode = "manual"
	KFocusModeSingleShot FocusMode = "single-shot"
	KFocusModeContinuous FocusMode = "continuous"
)
