package media

type WhiteBalanceMode string

func (e WhiteBalanceMode) String() string {
	return string(e)
}

const (
	KWhiteBalanceModeNone       WhiteBalanceMode = "none"
	KWhiteBalanceModeManual     WhiteBalanceMode = "manual"
	KWhiteBalanceModeSingleShot WhiteBalanceMode = "single-shot"
	KWhiteBalanceModeContinuous WhiteBalanceMode = "continuous"
)
