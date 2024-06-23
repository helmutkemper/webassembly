package components

import "strings"

type tag struct {
	Event    string
	Type     string
	Min      string
	Max      string
	Value    string
	Step     string
	Key      string
	Label    string
	Listener string
	Func     string
	Default  string
}

func (e *tag) getTagKeyValue(data string) (key, value string) {
	list := strings.Split(data, ":")
	key = list[0]
	value = list[1]
	return
}

func (e *tag) init(tagRaw string) {

	list := strings.Split(tagRaw, ";")
	for k := range list {
		key, value := e.getTagKeyValue(list[k])
		switch key {
		case "event":
			e.Event = value
		case "label":
			e.Label = value
		case "type":
			e.Type = value
		case "min":
			e.Min = value
		case "max":
			e.Max = value
		case "step":
			e.Step = value
		case "key":
			e.Key = value
		case "value":
			e.Value = value
		case "listener":
			e.Listener = value
		case "func":
			e.Func = value
		case "default":
			e.Default = value
		}
	}

	return
}
