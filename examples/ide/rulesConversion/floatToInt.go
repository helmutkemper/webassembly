package rulesConversion

import "github.com/helmutkemper/webassembly/mathUtil"

func FloatToInt(f float64) (i int) {
	//return int(f)
	return mathUtil.FloatToInt(f)
}
