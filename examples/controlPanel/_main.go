//go:build js

package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/components"
	"github.com/helmutkemper/iotmaker.webassembly/runTimeUtil"
	"log"
	"reflect"
	"strconv"
	"syscall/js"
)

type Color struct {
	Name   string `adminType:"label"`
	Min    int64
	Global components.RangeInt
	Red    int    `adminType:"rangeInt" adminMin:"0" adminMax:"255" adminValue:"0" adminLabel:"Red"`
	Green  int    `adminType:"rangeInt" adminMin:"0" adminMax:"255" adminValue:"0" adminLabel:"Green"`
	Blue   int    `adminType:"rangeFloat" adminMin:"0" adminMax:"255" adminValue:"0" adminStep:"0.01" adminLabel:"Blue"`
	Button string `adminType:"button" adminLabel:"Button"`
}

type Test struct {
	Components
	div html.TagDiv

	Color      Color `adminType:"rangeCell"`
	TestKemper int
}

func (e *Test) Init() (div *html.TagDiv, err error) {
	return e.init(e, reflect.ValueOf(e), reflect.TypeOf(e))
}

func main() {

	r := components.RangeInt{
		//Label: "Label",
		Text:  "text",
		Min:   -10,
		Max:   255,
		Value: 100,
		Step:  1,
		ChangeValue: func(value, minimum, maximum int64) {
			log.Printf("value: %v", value)
			log.Printf("minimum: %v", minimum)
			log.Printf("maximum: %v", maximum)
		},
	}
	t := Test{
		Color: Color{
			Name:   "Color distance",
			Global: r,
			Red:    10,
			Green:  20,
			Blue:   40,
		},
	}
	d, err := t.Init()
	if err != nil {
		panic(err)
	}

	divPanel := factoryBrowser.NewTagDiv().Id("panel").Append(d)

	stage := factoryBrowser.NewStage()
	stage.Append(divPanel)

	done := make(chan struct{})
	<-done
}

type Components struct {
	sliderEvent chan generic.Data
	numberEvent chan generic.Data
	buttonEvent chan generic.Data

	eventListSlider []func()
	eventListNumber []func()
	eventListButton []func()
}

func (e *Components) GetChanSlide() (event chan generic.Data) {
	return e.sliderEvent
}

func (e *Components) GetChanNumber() (event chan generic.Data) {
	return e.numberEvent
}

func (e *Components) GetChanButton() (event chan generic.Data) {
	return e.buttonEvent
}

func (e *Components) makeControlCellLabel(element reflect.Value) (div *html.TagDiv) {
	return e.basicDivCell(element.Elem().String(), "internalCell")
}

func (e *Components) makeControlCellButton(element reflect.Value, config *componentsButton) (div *html.TagDiv) {

	id := e.GetUId()

	div = factoryBrowser.NewTagDiv().
		Id(id+"Filter").
		Class("internalCell").
		Append(
			factoryBrowser.NewTagDiv().
				Class("text").
				Lang(html.Language(config.getLabel())).
				Text("Label"),
			factoryBrowser.NewTagInputButton(id+"Button").
				Class("button").
				Value(config.getLabel()).
				Lang(html.Language(config.getLanguage())).
				AddListenerClick(e.buttonEvent),
		)

	e.eventListButton = append(
		e.eventListButton, func() {
			//numberTag.Value(rangeTag.GetValue())
			if element.CanAddr() {
				element.SetString(e.GetUId())
			}
		},
	)

	return
}

func (e *Components) makeControlCellRange(element reflect.Value, config configGetInterface) (div *html.TagDiv) {

	id := e.GetUId()

	rangeTag := factoryBrowser.NewTagInputRange(id + "Range").
		Class("slider").
		Min(config.getMin()).
		Max(config.getMax()).
		Value(config.getValue()).
		Lang(html.Language(config.getLanguage())).
		AddListenerInput(e.sliderEvent)

	numberTag := factoryBrowser.NewTagInputNumber(id + "Number").
		Class("number").
		Min(config.getMin()).
		Max(config.getMax()).
		Value(config.getValue()).
		Lang(html.Language(config.getLabel())).
		AddListenerInput(e.numberEvent)

	div = factoryBrowser.NewTagDiv().
		Id(id+"Filter").
		Class("internalCell").
		Append(
			factoryBrowser.NewTagDiv().
				Class("text").
				Lang(html.Language(config.getLabel())).
				Text(config.getLabel()),
			rangeTag,
			numberTag,
		)

	e.eventListSlider = append(
		e.eventListSlider, func() {
			numberTag.Value(rangeTag.GetValue())
			if element.CanAddr() {
				element.SetInt(int64(rangeTag.GetValue()))
			}
		},
	)
	e.eventListNumber = append(
		e.eventListNumber, func() {
			rangeTag.Value(numberTag.GetValue())
			if element.CanAddr() {
				element.SetInt(int64(numberTag.GetValue()))
			}
		},
	)

	return
}

func (e *Components) makeControlCell(element reflect.Value) (elements []html.Compatible, err error) {
	elements = make([]html.Compatible, 0)

	if element.Kind() == reflect.Pointer {
		element = element.Elem()
	}

	for i := 0; i < element.NumField(); i++ {

		field := element.Field(i)
		typeField := element.Type().Field(i)
		tag := typeField.Tag

		if field.CanInterface() { // Verifica se o campo é exportado
			tagName := tag.Get("adminType")
			if tagName == "-" {
				continue
			}

			switch tagName {
			case "label":

				elements = append(elements, e.makeControlCellLabel(field.Addr()))
			case "rangeInt":

				config := new(componentsRangeInt)
				if err = config.populate(tag); err != nil {
					file, line, funcName := runTimeUtil.Trace()
					err = errors.Join(fmt.Errorf("%v(line: %v).populate().error: %v", funcName, line, err))
					err = errors.Join(fmt.Errorf("file: %v", file), err)
					return
				}

				// If you pass field.Addr() it is no longer addressable
				elements = append(elements, e.makeControlCellRange(field, config))
			case "rangeFloat":

				config := new(adminRangeFloat)
				if err = config.populate(tag); err != nil {
					file, line, funcName := runTimeUtil.Trace()
					err = errors.Join(fmt.Errorf("%v(line: %v).populate().error: %v", funcName, line, err))
					err = errors.Join(fmt.Errorf("file: %v", file), err)
					return
				}

				// If you pass field.Addr() it is no longer addressable
				elements = append(elements, e.makeControlCellRange(field, config))

			case "button":

				config := new(componentsButton)
				if err = config.populate(tag); err != nil {
					file, line, funcName := runTimeUtil.Trace()
					err = errors.Join(fmt.Errorf("%v(line: %v).populate().error: %v", funcName, line, err))
					err = errors.Join(fmt.Errorf("file: %v", file), err)
					return
				}
				elements = append(elements, e.makeControlCellButton(field, config))
			}

		}
	}

	return
}

func (e *Components) process(valueOf reflect.Value, typeOf reflect.Type) (div *html.TagDiv) {

	id := e.GetUId()
	div = factoryBrowser.NewTagDiv().Id(id).Class("panelCel")

	if typeOf.Kind() == reflect.Pointer {
		typeOf = typeOf.Elem()
		valueOf = valueOf.Elem()
	}

	if typeOf.Kind() == reflect.Struct {
		// Itera sobre os campos da struct
		for i := 0; i < typeOf.NumField(); i++ {
			field := typeOf.Field(i)
			fieldValue := valueOf.Field(i)

			//fieldName := field.Name
			fieldType := field.Type

			if fieldType.String() == "components.RangeInt" {
				log.Printf(">>>>>>>--------------------------------")
				if fieldValue.CanInterface() {

					e.processRange(div, fieldValue, fieldValue.Type())

				}
			}

			// Imprime o nome do campo e o nome do tipo do campo
			//if fieldValue.CanInterface() {
			//	log.Printf("Campo: %s, Tipo: %s, Value: %v\n", fieldName, fieldType, fieldValue.Interface())
			//	log.Printf("%v", fieldValue)
			//}

			if fieldValue.CanInterface() && field.Type.Kind() == reflect.Struct {
				div.Append(
					e.process(fieldValue, field.Type),
				)
			}
		}

	}

	return
}

func (e *Components) processRange(panelCel *html.TagDiv, valueOf reflect.Value, typeOf reflect.Type) {

	/*                                                                     */
	/* +- panel --------------------------------------------------------+  */
	/* |                                                                |  */
	/* |  +- panelCel -----------------------------------------------+  |  */
	/* |  |                                                          |  |  */
	/* |  | +- labelCel -------------------------------------------+ |  |  */
	/* |  | | Label                                              ˇ | |  |  */
	/* |  | +- compCel --------------------------------------------+ |  |  */
	/* |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |  */
	/* |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |  */
	/* |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |  */
	/* |  | +------------------------------------------------------+ |  |  */
	/* |  |                                                          |  |  */
	/* |  +----------------------------------------------------------+  |  */
	/* |                                                                |  */
	/* |  +- panelCel -----------------------------------------------+  |  */
	/* |  |                                                          |  |  */
	/* |  | +- labelCel -------------------------------------------+ |  |  */
	/* |  | | Label                                              ˇ | |  |  */
	/* |  | +- compCel --------------------------------------------+ |  |  */
	/* |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |  */
	/* |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |  */
	/* |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |  */
	/* |  | +------------------------------------------------------+ |  |  */
	/* |  |                                                          |  |  */
	/* |  +----------------------------------------------------------+  |  */
	/* |                                                                |  */
	/* +----------------------------------------------------------------+  */
	/*                                                                     */

	id := e.GetUId()

	//panelCel = factoryBrowser.NewTagDiv().Id(id + "-panel-cel").Class("panelCel")
	//labelCel := factoryBrowser.NewTagDiv().Id(id + "-label-cel").Class("labelCel")
	compCel := factoryBrowser.NewTagDiv().Id(id + "-comp-cel").Class("compCel")

	//labelTag := factoryBrowser.NewTagDiv().Id(id + "-label-tag").Class("text")
	textTag := factoryBrowser.NewTagDiv().Id(id + "-text-tag").Class("text")
	rangeTag := factoryBrowser.NewTagInputRange(id + "-input-range").Class("slider")
	numberTag := factoryBrowser.NewTagInputNumber(id + "-input-number").Class("number")

	//labelCel.Append(labelTag)
	compCel.Append(textTag, rangeTag, numberTag)
	panelCel.Append(compCel)

	var eChangeValue components.RangeIntFunc

	//elementType := reflect.TypeOf(el)
	//elementValue := reflect.ValueOf(el)

	if typeOf.Kind() == reflect.Pointer {
		typeOf = typeOf.Elem()
		valueOf = valueOf.Elem()
	}

	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		fieldValue := valueOf.Field(i)

		fieldName := field.Name
		//fieldType := field.Type

		//log.Printf("name: %v", fieldName)
		//log.Printf("type: %v", fieldType)

		switch fieldName {
		//case "Label":
		//labelTag.Text(fieldValue.Interface())

		case "Text":
			textTag.Text(fieldValue.Interface())

		case "Min":
			rangeTag.Min(fieldValue.Interface())
			numberTag.Min(fieldValue.Interface())

		case "Max":
			rangeTag.Max(fieldValue.Interface())
			numberTag.Max(fieldValue.Interface())

		case "Value":
			rangeTag.Value(fieldValue.Interface())
			numberTag.Value(fieldValue.Interface())

		case "Step":
			rangeTag.Step(fieldValue.Interface())
			numberTag.Step(fieldValue.Interface())

		case "ChangeValue":
			eChangeValue = fieldValue.Interface().(components.RangeIntFunc)

		}
	}

	elJs := rangeTag.Get()
	var fn js.Func
	fn = js.FuncOf(func(this js.Value, _ []js.Value) interface{} {

		valFl := int64(0)
		maxFl := int64(0)
		minFl := int64(0)

		valObj := this.Get("value")
		if !valObj.IsNull() && !valObj.IsUndefined() && !valObj.IsNaN() {
			valFl, _ = strconv.ParseInt(valObj.String(), 10, 64)
		}

		maxObj := this.Get("max")
		if !maxObj.IsNull() && !maxObj.IsUndefined() && !maxObj.IsNaN() {
			maxFl, _ = strconv.ParseInt(maxObj.String(), 10, 64)
		}

		minObj := this.Get("min")
		if !minObj.IsNull() && !minObj.IsUndefined() && !minObj.IsNaN() {
			minFl, _ = strconv.ParseInt(minObj.String(), 10, 64)
		}

		if eChangeValue != nil {
			eChangeValue(valFl, minFl, maxFl)
		}
		return nil
	})
	elJs.Call(
		"addEventListener",
		"change",
		fn,
	)

	//Lang(html.Language(config.getLabel())).
	//Text(config.getLabel()),

	//e.eventListSlider = append(
	//	e.eventListSlider, func() {
	//		numberTag.Value(rangeTag.GetValue())
	//		if element.CanAddr() {
	//			element.SetInt(int64(rangeTag.GetValue()))
	//		}
	//	},
	//)
	//e.eventListNumber = append(
	//	e.eventListNumber, func() {
	//		rangeTag.Value(numberTag.GetValue())
	//		if element.CanAddr() {
	//			element.SetInt(int64(numberTag.GetValue()))
	//		}
	//	},
	//)

	//div.Append(
	//	factoryBrowser.NewTagDiv().Id(id+"Label").Class("internalCell").Append(label),
	//	factoryBrowser.NewTagDiv().Id(id+"Internal").Class("internalCell").Append(
	//		text,
	//		rangeTag,
	//		numberTag,
	//	),
	//)
	return
}

func (e *Components) init(el any, valueOf reflect.Value, typeOf reflect.Type) (div *html.TagDiv, err error) {

	e.sliderEvent = make(chan generic.Data)
	e.numberEvent = make(chan generic.Data)
	e.buttonEvent = make(chan generic.Data)

	go func() {
		for {
			select {
			case <-e.sliderEvent:
				for k := range e.eventListSlider {
					e.eventListSlider[k]()
				}

			case <-e.numberEvent:
				for k := range e.eventListNumber {
					e.eventListNumber[k]()
				}

			case <-e.buttonEvent:
				for k := range e.eventListButton {
					e.eventListButton[k]()
				}
			}
		}
	}()

	div = factoryBrowser.NewTagDiv().Class("panel") //.Class("cell")
	div.Append(e.process(valueOf, typeOf))
	return

	element := reflect.ValueOf(el)

	if element.Kind() == reflect.Pointer {
		element = element.Elem()
	}

	for i := 0; i < element.NumField(); i++ {
		field := element.Field(i)
		typeField := element.Type().Field(i)
		tag := typeField.Tag

		if field.CanInterface() { // Verifica se o campo é exportado
			tagName := tag.Get("components") //todo: constante
			if tagName == "-" {
				continue
			}

			switch field.Type().Name() {
			case "components.RangeInt":
				log.Printf(">>>>>>>>>> reconheceu")
			}

			switch tagName {
			case "rangeCell":
				// Se o campo for um struct e for endereçável, chame init recursivamente
				if field.Kind() == reflect.Struct && field.CanAddr() {

					var elements []html.Compatible
					if elements, err = e.makeControlCell(field.Addr()); err != nil {
						file, line, funcName := runTimeUtil.Trace()
						err = errors.Join(fmt.Errorf("%v(line: %v).makeControlCell().error: %v", funcName, line, err))
						err = errors.Join(fmt.Errorf("file: %v", file), err)
						return
					}
					div.Append(elements...)

				}
			case "buttonCell":
				if field.Kind() == reflect.Struct && field.CanAddr() {
					var elements []html.Compatible
					if elements, err = e.makeControlCell(field.Addr()); err != nil {
						file, line, funcName := runTimeUtil.Trace()
						err = errors.Join(fmt.Errorf("%v(line: %v).makeControlCell().error: %v", funcName, line, err))
						err = errors.Join(fmt.Errorf("file: %v", file), err)
						return
					}
					div.Append(elements...)
				}
			}
		}
	}

	return
}

// basicDivCell Draw a div with text
func (e *Components) basicDivCell(label, class string) (div *html.TagDiv) {

	id := e.GetUId()

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

func (e *Components) GetUId() (uuidStr string) {
	uId, err := uuid.NewUUID()
	if err != nil {
		err = fmt.Errorf("controlCell.NewUUID().error: %v", err)
		return
	}
	uuidStr = uId.String()
	return
}

type componentsButton struct {
	label    string
	language string
}

func (e *componentsButton) getLabel() (label string) {
	return e.label
}

func (e *componentsButton) getLanguage() (language string) {
	return e.language
}

func (e *componentsButton) populate(tag reflect.StructTag) (err error) {
	label := tag.Get("adminLabel")
	language := tag.Get("adminLanguage")

	if label == "" {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(value).error: %v", funcName, line, fmt.Errorf("the label value must be informed")))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	} else {
		e.label = label
	}

	if language == "" {
		e.language = "en-us"
	} else {
		e.language = language
	}

	return
}

type componentsRangeInt struct {
	min      int64
	max      int64
	value    int64
	step     int64
	label    string
	language string
}

func (e *componentsRangeInt) getMin() (min any) {
	return e.min
}

func (e *componentsRangeInt) getMax() (max any) {
	return e.max
}

func (e *componentsRangeInt) getValue() (value any) {
	return e.value
}

func (e *componentsRangeInt) getStep() (step any) {
	return e.step
}

func (e *componentsRangeInt) getLabel() (label string) {
	return e.label
}

func (e *componentsRangeInt) getLanguage() (language string) {
	return e.language
}

func (e *componentsRangeInt) populate(tag reflect.StructTag) (err error) {
	minStr := tag.Get("adminMin")
	maxStr := tag.Get("adminMax")
	value := tag.Get("adminValue")
	step := tag.Get("adminStep")
	label := tag.Get("adminLabel")
	language := tag.Get("adminLanguage")

	if minStr == "" {
		e.min = 0
	} else if e.min, err = strconv.ParseInt(minStr, 10, 64); err != nil {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(min).error: %v", funcName, line, err))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}

	if maxStr == "" {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(max).error: %v", funcName, line, fmt.Errorf("the maximum value must be informed")))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	} else if e.max, err = strconv.ParseInt(maxStr, 10, 64); err != nil {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(max).error: %v", funcName, line, err))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}

	if value == "" {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(value).error: %v", funcName, line, fmt.Errorf("the default value must be informed")))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	} else if e.value, err = strconv.ParseInt(value, 10, 64); err != nil {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(value).error: %v", funcName, line, err))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}

	if step == "" {
		e.step = 1
	} else if e.step, err = strconv.ParseInt(step, 10, 64); err != nil {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(step).error: %v", funcName, line, err))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}

	if label == "" {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(value).error: %v", funcName, line, fmt.Errorf("the label value must be informed")))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	} else {
		e.label = label
	}

	if language == "" {
		e.language = "en-us"
	} else {
		e.language = language
	}

	return
}

type configGetInterface interface {
	getMin() (min any)
	getMax() (max any)
	getValue() (value any)
	getStep() (step any)
	getLabel() (label string)
	getLanguage() (language string)
}

type adminRangeFloat struct {
	min      float64
	max      float64
	value    float64
	step     float64
	label    string
	language string
}

func (e *adminRangeFloat) getMin() (min any) {
	return e.min
}

func (e *adminRangeFloat) getMax() (max any) {
	return e.max
}

func (e *adminRangeFloat) getValue() (value any) {
	return e.value
}

func (e *adminRangeFloat) getStep() (step any) {
	return e.step
}

func (e *adminRangeFloat) getLabel() (label string) {
	return e.label
}

func (e *adminRangeFloat) getLanguage() (language string) {
	return e.language
}

func (e *adminRangeFloat) populate(tag reflect.StructTag) (err error) {
	minStr := tag.Get("adminMin")
	maxStr := tag.Get("adminMax")
	value := tag.Get("adminValue")
	step := tag.Get("adminStep")
	label := tag.Get("adminLabel")
	language := tag.Get("adminLanguage")

	if minStr == "" {
		e.min = 0
	} else if e.min, err = strconv.ParseFloat(minStr, 64); err != nil {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(min).error: %v", funcName, line, err))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}

	if maxStr == "" {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(max).error: %v", funcName, line, fmt.Errorf("the maximum value must be informed")))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	} else if e.max, err = strconv.ParseFloat(maxStr, 64); err != nil {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(max).error: %v", funcName, line, err))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}

	if value == "" {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(value).error: %v", funcName, line, fmt.Errorf("the default value must be informed")))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	} else if e.value, err = strconv.ParseFloat(value, 64); err != nil {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(value).error: %v", funcName, line, err))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}

	if step == "" {
		e.step = 1
	} else if e.step, err = strconv.ParseFloat(step, 64); err != nil {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(step).error: %v", funcName, line, err))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}

	if label == "" {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).ParseInt(value).error: %v", funcName, line, fmt.Errorf("the label value must be informed")))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	} else {
		e.label = label
	}

	if language == "" {
		e.language = "en-us"
	} else {
		e.language = language
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
