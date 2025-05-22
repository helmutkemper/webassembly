package components

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// tag
//
// English:
//
//	Process the tag `wasmPanel`
//
// Português:
//
//	Processa a tag `wasmPanel`
type tag struct {
	Attach        string
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
	Latitude      string
	Longitude     string
	Zoom          string
	Icon          string
	IconRight     string
	IconLeft      string
	Menu          string
	SubMenu       string
	Action        string
	Left          string
	Top           string
	Drag          string
	Minimize      string
	Close         string
	Options       string
	Columns       string
}

// getTagKeyValue
//
// English:
//
//	Separates tag elements using point and comma as a separator. Eg.: `wasmPanel:"type:headerText;label:Control panel"`
//
// Português:
//
//	Separa os elementos da tag, usando ponto e vírgula como separador. Ex.: `wasmPanel:"type:headerText;label:Control panel"`
func (e *tag) getTagKeyValue(data string, isolationData []Isolation) (err error, key, value string) {
	pairKeyValue := strings.Split(data, ":")
	if len(pairKeyValue) != 2 {
		err = fmt.Errorf("error: the data `%v` has an invalid tag key/value pair", data)
		return
	}

	key = pairKeyValue[0]
	value = pairKeyValue[1]

	for k := range isolationData {
		value = strings.Replace(value, isolationData[k].key, isolationData[k].value, -1)
	}

	return
}

// Search
//
// English:
//
//	Search for a key / value combination within the tag
//
// Português:
//
//	Procura por uma combinação chave / valor dentro da tag
func (e *tag) Search(key, value string) (found bool) {
	element := reflect.ValueOf(*e)
	for i := 0; i < element.NumField(); i++ {

		a := element.Type().Field(i).Name == key
		b := element.Field(i).String() == value

		if a && b {
			return true
		}
	}
	return
}

// init
//
// English:
//
//	Receives the data RAW and returns the processed data
//
// Português:
//
//	Recebe o dado raw e devolve o dado processado
func (e *tag) init(tagRaw string) (err error) {

	isolate := Isolation{}
	output, isolationData := isolate.isolate(tagRaw)
	result := isolate.exchangeForKey(output, isolationData)

	result = strings.TrimRight(result, ";")
	list := strings.Split(result, ";")
	for k := range list {
		var key, value string
		if err, key, value = e.getTagKeyValue(list[k], isolationData); err != nil {
			err = errors.Join(fmt.Errorf("the raw `%v` data tag has processed width error", tagRaw), err)
			return
		}

		switch key {
		case "attach":
			e.Attach = value
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
		case "icon":
			e.Icon = value
		case "iconRight":
			e.IconRight = value
		case "iconLeft":
			e.IconLeft = value
		case "menu":
			e.Menu = value
		case "subMenu":
			e.SubMenu = value
		case "action":
			e.Action = value
		case "left":
			e.Left = value
		case "top":
			e.Top = value
		case "drag":
			e.Drag = value
		case "minimize":
			e.Minimize = value
		case "close":
			e.Close = value
		case "options":
			e.Options = value
		case "columns":
			e.Columns = value
		default:
			err = fmt.Errorf("a tag was not processed correctly for the key: %v, full value: `%v`", key, tagRaw)
			return
		}
	}

	return
}
