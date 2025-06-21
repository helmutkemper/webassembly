package utilsMath

import (
	"log"
	"testing"
)

func TestFloatToInt(t *testing.T) {
	for i := 0.0; i < 1.0; i += 0.1 {
		log.Printf("%v = FloatToInt(%v)", FloatToInt(i), i)
		if i < 0.5 && FloatToInt(i) != 0.0 {
			t.FailNow()
		} else if i >= 0.5 && FloatToInt(i) != 1.0 {
			t.FailNow()
		}
	}
}
