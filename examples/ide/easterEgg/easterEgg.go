package easterEgg

import (
	"fmt"
	"github.com/helmutkemper/webassembly/browser/html"
	"slices"
	"strings"
	"time"
)

// MorseCode Makes an SVG blink in code morse
type MorseCode struct {
	flashInterval int
	ticker        *time.Ticker
	flashStop     chan struct{}

	dit      []bool
	dah      []bool
	space    []bool
	splitter []bool

	period           []bool
	comma            []bool
	questionMark     []bool
	colon            []bool
	semiColon        []bool
	dash             []bool
	apostrophe       []bool
	slash            []bool
	quotationMark    []bool
	exclamationPoint []bool

	a []bool
	b []bool
	c []bool
	d []bool
	e []bool
	f []bool
	g []bool
	h []bool
	i []bool
	j []bool
	k []bool
	l []bool
	m []bool
	n []bool
	o []bool
	p []bool
	q []bool
	r []bool
	s []bool
	t []bool
	u []bool
	v []bool
	w []bool
	x []bool
	y []bool
	z []bool

	n0 []bool
	n1 []bool
	n2 []bool
	n3 []bool
	n4 []bool
	n5 []bool
	n6 []bool
	n7 []bool
	n8 []bool
	n9 []bool
}

func (e *MorseCode) FlashEnd() {
	if e.ticker != nil {
		e.ticker.Stop()
		e.ticker = nil

		e.flashStop <- struct{}{}
	}
}

// FlashMarkErrorMsg Faz um SVG piscar enviando um SOS
func (e *MorseCode) FlashMarkErrorMsg(svg *html.TagSvg) {

	viewingStandard, _ := e.TextToMorse("please, correct the error before running this code.  ")

	if e.ticker != nil {
		e.ticker.Stop()
		e.flashStop <- struct{}{}
	}

	e.ticker = time.NewTicker(time.Duration(e.flashInterval) * time.Millisecond)

	go func() {
		localViewingStandard := make([]bool, len(viewingStandard))
		copy(localViewingStandard, viewingStandard)

		svg.AddStyle("visibility", "visible")
		time.Sleep(1 * time.Minute)

		for {
			select {
			case <-e.flashStop:
				return

			case <-e.ticker.C:
				view := localViewingStandard[0]
				localViewingStandard = localViewingStandard[1:]

				if len(localViewingStandard) == 0 {
					localViewingStandard = make([]bool, len(viewingStandard))
					copy(localViewingStandard, viewingStandard)

					svg.AddStyle("visibility", "visible")
					time.Sleep(1 * time.Minute)
					continue
				}

				if view {
					svg.AddStyle("visibility", "visible")
				} else {
					svg.AddStyle("visibility", "hidden")
				}
			}
		}
	}()

}

// Init Monta o alfabeto Morse
func (e *MorseCode) Init() {

	// Waiting time for each boolean used on code morse
	e.flashInterval = 60

	e.dit = []bool{true, true, false, false}
	// The duration of a dah is three times the duration of a dit
	e.dah = []bool{true, true, true, true, true, true, false, false, false}
	// space, 3x dit
	e.space = []bool{false, false, false, false, false, false}
	// splitter, 7x dit
	e.splitter = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false}

	e.period = slices.Concat(e.dit, e.dah, e.dit, e.dah, e.dit, e.dah)           //.
	e.comma = slices.Concat(e.dah, e.dah, e.dit, e.dit, e.dah, e.dah)            //,
	e.questionMark = slices.Concat(e.dit, e.dit, e.dah, e.dah, e.dit, e.dit)     //?
	e.colon = slices.Concat(e.dah, e.dah, e.dah, e.dit, e.dit, e.dit)            //:
	e.semiColon = slices.Concat(e.dah, e.dit, e.dah, e.dit, e.dah, e.dit)        //;
	e.dash = slices.Concat(e.dah, e.dit, e.dit, e.dit, e.dit, e.dah)             //-
	e.apostrophe = slices.Concat(e.dit, e.dah, e.dah, e.dah, e.dah, e.dit)       //'
	e.slash = slices.Concat(e.dah, e.dit, e.dit, e.dah, e.dit)                   // /
	e.quotationMark = slices.Concat(e.dit, e.dah, e.dit, e.dit, e.dah, e.dit)    //"
	e.exclamationPoint = slices.Concat(e.dah, e.dit, e.dah, e.dit, e.dah, e.dah) //!

	e.a = slices.Concat(e.dit, e.dah)
	e.b = slices.Concat(e.dah, e.dit, e.dit, e.dit)
	e.c = slices.Concat(e.dah, e.dit, e.dah, e.dit)
	e.d = slices.Concat(e.dah, e.dit, e.dit)
	e.e = slices.Concat(e.dit)
	e.f = slices.Concat(e.dit, e.dit, e.dah, e.dit)
	e.g = slices.Concat(e.dah, e.dah, e.dit)
	e.h = slices.Concat(e.dit, e.dit, e.dit, e.dit)
	e.i = slices.Concat(e.dit, e.dit)
	e.j = slices.Concat(e.dit, e.dah, e.dah, e.dah)
	e.k = slices.Concat(e.dah, e.dit, e.dah)
	e.l = slices.Concat(e.dit, e.dah, e.dit, e.dit)
	e.m = slices.Concat(e.dah, e.dah)
	e.n = slices.Concat(e.dah, e.dit)
	e.o = slices.Concat(e.dah, e.dah, e.dah)
	e.p = slices.Concat(e.dit, e.dah, e.dah, e.dit)
	e.q = slices.Concat(e.dah, e.dah, e.dit, e.dah)
	e.r = slices.Concat(e.dit, e.dah, e.dit)
	e.s = slices.Concat(e.dit, e.dit, e.dit)
	e.t = slices.Concat(e.dah)
	e.u = slices.Concat(e.dit, e.dit, e.dah)
	e.v = slices.Concat(e.dit, e.dit, e.dit, e.dah)
	e.w = slices.Concat(e.dit, e.dah, e.dah)
	e.x = slices.Concat(e.dah, e.dit, e.dit, e.dah)
	e.y = slices.Concat(e.dah, e.dit, e.dah, e.dah)
	e.z = slices.Concat(e.dah, e.dah, e.dit, e.dit)

	e.n1 = slices.Concat(e.dit, e.dah, e.dah, e.dah, e.dah)
	e.n2 = slices.Concat(e.dit, e.dit, e.dah, e.dah, e.dah)
	e.n3 = slices.Concat(e.dit, e.dit, e.dit, e.dah, e.dah)
	e.n4 = slices.Concat(e.dit, e.dit, e.dit, e.dit, e.dah)
	e.n5 = slices.Concat(e.dit, e.dit, e.dit, e.dit, e.dit)
	e.n6 = slices.Concat(e.dah, e.dit, e.dit, e.dit, e.dit)
	e.n7 = slices.Concat(e.dah, e.dah, e.dit, e.dit, e.dit)
	e.n8 = slices.Concat(e.dah, e.dah, e.dah, e.dit, e.dit)
	e.n9 = slices.Concat(e.dah, e.dah, e.dah, e.dah, e.dit)
	e.n0 = slices.Concat(e.dah, e.dah, e.dah, e.dah, e.dah)
}

// TextToMorse Converts text into Boolean Slice Morse
func (e *MorseCode) TextToMorse(text string) (code []bool, err error) {
	text = strings.ToLower(text)

	code = make([]bool, 0)
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '.':
			code = append(code, e.period...)
		case ',':
			code = append(code, e.comma...)
		case '?':
			code = append(code, e.questionMark...)
		case ':':
			code = append(code, e.colon...)
		case ';':
			code = append(code, e.semiColon...)
		case '-':
			code = append(code, e.dash...)
		case '\'':
			code = append(code, e.apostrophe...)
		case '/':
			code = append(code, e.slash...)
		case '"':
			code = append(code, e.quotationMark...)
		case '!':
			code = append(code, e.exclamationPoint...)
		case 'a':
			code = append(code, e.a...)
		case 'b':
			code = append(code, e.b...)
		case 'c':
			code = append(code, e.c...)
		case 'd':
			code = append(code, e.d...)
		case 'e':
			code = append(code, e.e...)
		case 'f':
			code = append(code, e.f...)
		case 'g':
			code = append(code, e.g...)
		case 'h':
			code = append(code, e.h...)
		case 'i':
			code = append(code, e.i...)
		case 'j':
			code = append(code, e.j...)
		case 'k':
			code = append(code, e.k...)
		case 'l':
			code = append(code, e.l...)
		case 'm':
			code = append(code, e.m...)
		case 'n':
			code = append(code, e.n...)
		case 'o':
			code = append(code, e.o...)
		case 'p':
			code = append(code, e.p...)
		case 'q':
			code = append(code, e.q...)
		case 'r':
			code = append(code, e.r...)
		case 's':
			code = append(code, e.s...)
		case 't':
			code = append(code, e.t...)
		case 'u':
			code = append(code, e.u...)
		case 'v':
			code = append(code, e.v...)
		case 'w':
			code = append(code, e.w...)
		case 'x':
			code = append(code, e.x...)
		case 'y':
			code = append(code, e.y...)
		case 'z':
			code = append(code, e.z...)
		case '1':
			code = append(code, e.n1...)
		case '2':
			code = append(code, e.n2...)
		case '3':
			code = append(code, e.n3...)
		case '4':
			code = append(code, e.n4...)
		case '5':
			code = append(code, e.n5...)
		case '6':
			code = append(code, e.n6...)
		case '7':
			code = append(code, e.n7...)
		case '8':
			code = append(code, e.n8...)
		case '9':
			code = append(code, e.n9...)
		case '0':
			code = append(code, e.n0...)
		case ' ':
			code = append(code, e.splitter...)
		default:
			err = fmt.Errorf("character not supported: %v", text[i])
			return
		}

		code = append(code, e.space...)
	}

	return
}
