package components

import "strconv"

type Isolation struct {
	value string
	start int
	end   int
	key   string
}

func (e *Isolation) exchangeForKey(output string, list []Isolation) (result string) {
	bufOut := ""
	for i := len(list) - 1; i != -1; i -= 1 {
		bufKey := "'$#@!'"
		key := strconv.FormatInt(int64(i), 10)
		bufKey += key
		bufKey += "'!@#$'"
		list[i].key = bufKey

		bufOut = ""
		bufOut += output[:list[i].start]
		bufOut += bufKey
		bufOut += output[list[i].end:]

		output = bufOut
	}

	return output
}

func (e *Isolation) isolate(data string) (output string, list []Isolation) {
	list = make([]Isolation, 0)

	// counts how many quotes were removed as the total text size decreases
	interactions := 0
	// counts how many quotes were removed in the skip function, as the total size of the text decreases
	skips := 0

	start := 0
	end := 0
	value := ""
	skipNext := false
	match := false
	capture := false
	for index := 0; index != len(data); index += 1 {
		char := data[index]

		if char == byte('\\') {
			skipNext = true
			skips += 1
			continue
		}

		if char == byte('\'') {
			if skipNext {
				skipNext = false
			} else {
				match = true
				capture = !capture

				if capture {
					start = index - (interactions * 2) - skips
				} else {
					end = index - 1 - (interactions * 2) - skips
					interactions += 1
					list = append(list, Isolation{value: value, start: start, end: end})
					value = ""
				}
			}
		}

		if match {
			match = false
		} else if capture {
			output += string(char)
			value += string(char)
			//fmt.Printf("\033[31;1;4m%v\033[0m", string(char))
		} else {
			output += string(char)
			//fmt.Printf("%v", string(char))
		}
	}
	//fmt.Print("\n\n")
	return
}
