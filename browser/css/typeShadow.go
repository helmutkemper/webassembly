package css

import (
	"strings"
)

type Shadow []BoxShadow

func (s *Shadow) Add(shadow BoxShadow) {
	*s = append(*s, shadow)
}

func (s Shadow) String() string {
	var shadows = make([]string, 0)

	for _, shadow := range s {
		shadows = append(shadows, shadow.String())
	}
	return strings.Join(shadows, ",\n") + ";"
}
