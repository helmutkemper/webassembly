package html

import (
	"image/color"
	"strconv"
	"time"
)

func TypeToString(value interface{}, separadorA, separadorB string) (ret interface{}) {
	separadorALen := len(separadorA)
	separadorBLen := len(separadorB)

	str := ""

	switch converted := value.(type) {
	case float64:

		ret = strconv.FormatFloat(converted, 'g', 3, 32)

	case float32:

		ret = strconv.FormatFloat(100.0*float64(converted), 'g', 3, 64) + "%"

	case time.Duration:

		ret = converted.String()

	case time.Time:

		ret = converted.String()

	case color.RGBA:

		ret = RGBAToJs(converted)

	case []float64:

		for _, v := range converted {
			p := strconv.FormatFloat(v, 'g', 3, 32) + separadorA
			str += p
		}
		ret = str[:len(str)-separadorALen]

	case []float32:

		for _, v := range converted {
			p := strconv.FormatFloat(100.0*float64(v), 'g', 3, 64) + "%" + separadorA
			str += p
		}
		ret = str[:len(str)-separadorALen]

	case []time.Duration:

		for _, v := range converted {
			str += v.String() + separadorA
		}
		ret = str[:len(str)-separadorALen]

	case []time.Time:

		for _, v := range converted {
			str += v.String() + separadorA
		}
		ret = str[:len(str)-separadorALen]

	case []color.RGBA:

		for _, v := range converted {
			str += RGBAToJs(v) + separadorA
		}
		ret = str[:len(str)-separadorALen]

	case [][]float64:

		for _, list := range converted {
			for _, v := range list {
				p := strconv.FormatFloat(v, 'g', 3, 32) + separadorA
				str += p
			}
			str = str[:len(str)-separadorALen] + separadorB
		}

		ret = str[:len(str)-separadorBLen]

	case [][]float32:

		for _, list := range converted {
			for _, v := range list {
				p := strconv.FormatFloat(100.0*float64(v), 'g', 3, 64) + "%" + separadorA
				str += p
			}
			str = str[:len(str)-separadorALen] + separadorB
		}

		ret = str[:len(str)-separadorBLen]

	case [][]time.Duration:

		for _, list := range converted {
			for _, v := range list {
				str += v.String() + separadorA
			}
			str = str[:len(str)-separadorALen] + separadorB
		}

		ret = str[:len(str)-separadorBLen]

	case [][]time.Time:

		for _, list := range converted {
			for _, v := range list {
				str += v.String() + separadorA
			}
			str = str[:len(str)-separadorALen] + separadorB
		}

		ret = str[:len(str)-separadorBLen]

	case [][]color.RGBA:

		for _, list := range converted {
			for _, v := range list {
				str += RGBAToJs(v) + separadorA
			}
			str = str[:len(str)-separadorALen] + separadorB
		}

		ret = str[:len(str)-separadorBLen]

	default:

		ret = value
	}

	return
}
