package components

import (
	"strings"
)

type tag struct {
	Event         string
	Type          string
	Min           string
	Max           string
	Value         string
	Step          string
	Key           string
	Label         string
	Listener      string
	Func          string
	Default       string
	PlaceHolder   string
	Name          string
	Width         string
	Height        string
	Size          string
	Level         string
	Color         string
	Background    string
	DisableBorder string
}

func (e *tag) getTagKeyValue(data string, isolationData []Isolation) (key, value string) {
	pairKeyValue := strings.Split(data, ":")
	key = pairKeyValue[0]
	value = pairKeyValue[1]

	for k := range isolationData {
		value = strings.Replace(value, string(isolationData[k].key), string(isolationData[k].value), -1)
	}

	return
}

func (e *tag) init(tagRaw string) {

	isolate := Isolation{}
	output, isolationData := isolate.isolate(tagRaw)
	result := isolate.exchangeForKey(output, isolationData)

	list := strings.Split(string(result), ";")
	for k := range list {
		key, value := e.getTagKeyValue(list[k], isolationData)
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
		case "placeHolder":
			e.PlaceHolder = value
		case "name":
			e.Name = value
		case "height":
			e.Height = value
		case "width":
			e.Width = value
		case "size":
			e.Size = value
		case "level":
			e.Level = value
		case "color":
			e.Color = value
		case "background":
			e.Background = value
		case "disableBorder":
			e.DisableBorder = value
		}
	}

	return
}
