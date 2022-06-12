package html

type Overflow string

func (e Overflow) String() string {
	return string(e)
}

const (
	KOverflowVisible Overflow = "visible"
	KOverflowHidden  Overflow = "hidden"
	KOverflowScroll  Overflow = "scroll"
	KOverflowAuto    Overflow = "auto"
)
