//go:build js

package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"log"
	"reflect"
	"strings"
	"time"
)

const (
	kTagName = "admin"
)

func uuidStr() (uuidStr string) {
	uId, err := uuid.NewUUID()
	if err != nil {
		err = fmt.Errorf("controlCell.NewUUID().error: %v", err)
		return
	}
	uuidStr = uId.String()
	return
}

func (e *Admin) basicRangeCell(label, class, id string, min, max, value any, element reflect.Value) (div *html.TagDiv) {
	slicerTag := factoryBrowser.NewTagInputRange(id + "Range").
		Class("slider").
		Min(min).
		Max(max).
		Value(value).
		AddListenerInput(e.sliderEvent)

	numberTag := factoryBrowser.NewTagInputNumber(id + "Number").
		Class("number").
		Min(min).
		Max(max).
		Value(value).
		AddListenerInput(e.numberEvent)

	div = factoryBrowser.NewTagDiv().
		Id(id+"Filter").
		Class(class).
		Append(
			factoryBrowser.NewTagDiv().
				Class("text").
				Lang(html.KLanguageEnglishUnitedStates).
				Text(label),
			slicerTag,
			numberTag,
		)

	//e.eventListSlider = append(
	//	e.eventListSlider, func() {
	//		numberTag.Value(slicerTag.GetValue())
	//		if element.CanAddr() {
	//			element.SetInt(int64(slicerTag.GetValue()))
	//		}
	//	},
	//)
	//e.eventListNumber = append(
	//	e.eventListNumber, func() {
	//		slicerTag.Value(numberTag.GetValue())
	//		if element.CanAddr() {
	//			element.SetInt(int64(numberTag.GetValue()))
	//		}
	//	},
	//)
	return
}

func (e *Admin) basicDivCell(label, class, id string) (div *html.TagDiv) {
	div = factoryBrowser.NewTagDiv().
		Id(id + "Filter").
		Class(class).
		Append(
			factoryBrowser.NewTagDiv().
				Class("text").
				Lang(html.KLanguageEnglishUnitedStates).
				Text(label),
		)

	return
}

func main() {

	t := Test{
		Color: Color{
			Name:   "Color distance",
			Global: 128,
			Red:    10,
			Green:  20,
			Blue:   40,
		},
	}
	d := t.Init()

	go func() {
		for {
			log.Printf("Global: %v", t.Color.Global)
			time.Sleep(time.Second)
		}
	}()

	divPanel := factoryBrowser.NewTagDiv().Id("panel").Append(d)

	stage := factoryBrowser.NewStage()
	stage.Append(divPanel)

	done := make(chan struct{})
	<-done
}

type Color struct {
	Name   string `admin:"type:label"`
	Global int    `admin:"type:slice;min:0;max:255;value:128;label:Global"`
	Red    int    `admin:"type:slice;min:0;max:255;value:0;label:Red"`
	Green  int    `admin:"type:slice;min:0;max:255;value:0;label:Green"`
	Blue   int    `admin:"type:slice;min:0;max:255;value:0;label:Blue"`
}

type Test struct {
	Admin

	Color Color `admin:"type:controlCell"`
}

func (e *Test) Init() (div *html.TagDiv) {
	return e.init(reflect.ValueOf(e).Elem())
}

type Admin struct {
	sliderEvent chan generic.Data
	numberEvent chan generic.Data

	eventListSlider []func()
	eventListNumber []func()
}

func (e *Admin) makeControlCellLabel(element reflect.Value, data []string) (div *html.TagDiv) {
	for k := range data {
		values := strings.Split(data[k], ":")
		if len(values) == 2 {
			key := values[0]
			value := values[1]

			if key == "type" && value == "label" {
				log.Printf("element name: %v", element.Elem().String())
				return e.basicDivCell(element.Elem().String(), "internalCell", "0")
			}
		}
	}

	//todo: error de config!
	return
}

func (e *Admin) makeControlCellSlice(element reflect.Value, data []string) (div *html.TagDiv) {

	var eLabel, eMax, eMin, eValue any

	switch element.Elem().Kind() {
	case reflect.Int:
		eValue = element.Elem().Int()
	case reflect.Float64:
		eValue = element.Elem().Float()
	}

	for k := range data {
		values := strings.Split(data[k], ":")
		if len(values) == 2 {
			key := values[0]
			value := values[1]

			switch key {
			case "min":
				eMin = value
			case "max":
				eMax = value
			case "label":
				eLabel = value
			}
		}
	}

	return e.basicRangeCell(eLabel.(string), "internalCell", "1", eMin, eMax, eValue, element.Elem())
}

func (e *Admin) makeControlCell(element reflect.Value) (elements []html.Compatible) {
	elements = make([]html.Compatible, 0)

	if element.Kind() == reflect.Pointer {
		element = element.Elem()
	}

	for i := 0; i < element.NumField(); i++ {
		field := element.Field(i)
		typeField := element.Type().Field(i)
		tag := typeField.Tag

		if field.CanInterface() { // Verifica se o campo é exportado
			tagName := tag.Get(kTagName)
			if tagName == "-" {
				continue
			}

			configList := strings.Split(tagName, ";")
			for k := range configList {
				values := strings.Split(configList[k], ":")
				if len(values) == 2 {
					key := values[0]
					value := values[1]

					switch key {
					case "type":

						switch value {
						case "label":
							elements = append(elements, e.makeControlCellLabel(field.Addr(), configList))
						case "slice":
							elements = append(elements, e.makeControlCellSlice(field.Addr(), configList))
						}

					}

				}
			}
		}
	}

	return
}

func (e *Admin) channelLoopB() {
	defer func() {
		log.Printf("não deveria ter entrado aqui")
	}()
	for {
		select {
		//case <-e.sliderEvent:
		//	for k := range e.eventListSlider {
		//		e.eventListSlider[k]()
		//	}

		case <-e.numberEvent:
			for k := range e.eventListNumber {
				e.eventListNumber[k]()
			}
		}
	}
}

func (e *Admin) channelLoopA() {
	defer func() {
		log.Printf("não deveria ter entrado aqui")
	}()
	for {
		select {
		case <-e.sliderEvent:
			for k := range e.eventListSlider {
				e.eventListSlider[k]()
			}

			//case <-e.numberEvent:
			//	for k := range e.eventListNumber {
			//		e.eventListNumber[k]()
			//	}
		}
	}
}

func (e *Admin) init(element reflect.Value) (div *html.TagDiv) {
	//e.sliderEvent = make(chan generic.Data)
	//e.numberEvent = make(chan generic.Data)

	go e.channelLoopA()
	go e.channelLoopB()

	div = factoryBrowser.NewTagDiv().Class("cell")

	if element.Kind() == reflect.Pointer {
		element = element.Elem()
	}

	for i := 0; i < element.NumField(); i++ {
		field := element.Field(i)
		typeField := element.Type().Field(i)
		tag := typeField.Tag

		if field.CanInterface() { // Verifica se o campo é exportado
			tagName := tag.Get(kTagName)
			if tagName == "-" {
				continue
			}

			configList := strings.Split(tagName, ";")
			for k := range configList {
				values := strings.Split(configList[k], ":")
				if len(values) == 2 {
					key := values[0]
					value := values[1]

					switch key {
					case "type":

						switch value {
						case "controlCell":
							if field.Kind() == reflect.Struct && field.CanAddr() {
								div.Append(e.makeControlCell(field.Addr())...)
							}
						}
					}
				}
			}

			// Se o campo for um struct e for endereçável, chame init recursivamente
			if field.Kind() == reflect.Struct && field.CanAddr() {
				e.init(field.Addr())
			}
		}
	}

	return
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
