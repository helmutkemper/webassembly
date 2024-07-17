package components

import (
	"errors"
	"fmt"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Components struct {
	ref         interface{}
	panelFather *html.TagDiv
	panelBody   *html.TagDiv
}

func (e *Components) Init(el any) (panel *html.TagDiv, err error) {
	element := reflect.ValueOf(el)
	typeof := reflect.TypeOf(el)
	e.createDivsFather()

	err = e.process(element, typeof)
	if err != nil {
		//file, line, funcName := runTimeUtil.Trace()
		//err = errors.Join(fmt.Errorf("%v(line: %v).process().error: %v", funcName, line, err))
		//err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}

	e.panelFather.Class("panel")

	return e.panelFather, err
}

//func (e *Components) GetUId() (uuidStr string, err error) {
//	uId, err := uuid.NewUUID()
//	if err != nil {
//		//file, line, funcName := runTimeUtil.Trace()
//		//err = errors.Join(fmt.Errorf("%v(line: %v).NewUUID().error: %v", funcName, line, err))
//		//err = errors.Join(fmt.Errorf("file: %v", file), err)
//		return
//	}
//	uuidStr = uId.String()
//	return
//}

func (e *Components) createDivsFather() {
	e.panelFather = factoryBrowser.NewTagDiv().Class("panel")
	e.panelBody = factoryBrowser.NewTagDiv().Class("panelBody")
}

func (e *Components) process(element reflect.Value, typeof reflect.Type) (err error) {

	if element.Kind() == reflect.Pointer {
		if element.CanSet() && element.IsNil() {
			newInstance := reflect.New(element.Type().Elem())
			element.Set(newInstance)
		}

		element = element.Elem()
		typeof = typeof.Elem()
	}

	if element.CanInterface() {
		for i := 0; i != element.NumField(); i += 1 {
			fieldVal := element.Field(i)
			fieldTyp := typeof.Field(i)

			//if fieldVal.Kind() == reflect.Struct {
			//	e.process(fieldVal, fieldTyp.Type)
			//}

			tagRaw := fieldTyp.Tag.Get("wasmPanel")
			if tagRaw != "" {
				tagData := new(tag)
				tagData.init(tagRaw)

				switch tagData.Type {
				case "headerText":
					e.processHeaderText(fieldVal, tagData.Label, e.panelFather)
					// Espera criar panelHeader para que panelBody fique abaixo
					e.panelFather.Append(e.panelBody)
				case "panelBody":

					// initialize the panelBody pointer
					if fieldVal.Kind() == reflect.Pointer {
						if fieldVal.CanSet() && fieldVal.IsNil() {
							newInstance := reflect.New(fieldVal.Type().Elem())
							fieldVal.Set(newInstance)
						}

						e.ref = fieldVal.Interface()
					}

					err = e.process(fieldVal, fieldTyp.Type)
					if err != nil {
						//file, line, funcName := runTimeUtil.Trace()
						//err = errors.Join(fmt.Errorf("%v(line: %v).process().error: %v", funcName, line, err))
						//err = errors.Join(fmt.Errorf("file: %v", file), err)
						return
					}
				case "compCel":
					// ignore
				case "component":
					divCompCel := factoryBrowser.NewTagDiv().Class("compCel")

					err = e.processComponent(fieldVal, fieldTyp.Type, tagData, divCompCel)
					if err != nil {
						//file, line, funcName := runTimeUtil.Trace()
						//err = errors.Join(fmt.Errorf("%v(line: %v).processComponent().error: %v", funcName, line, err))
						//err = errors.Join(fmt.Errorf("file: %v", file), err)
						return
					}

					panelCel := factoryBrowser.NewTagDiv().Class("panelCel")

					e.processLabelCel(tagData.Label, panelCel)

					panelCel.Append(
						divCompCel,
					)

					e.panelBody.Append(panelCel)
				case "celLabel":
					//log.Printf("%v", tagData.Type)
				case "celText":
					//log.Printf("%v", tagData.Type)
				case "range":
					//log.Printf("%v", tagData.Type)
				case "onchange":
					//log.Printf("%v", tagData.Type)
				default:
					//log.Printf(">>>>>>>> %v", tagData.Type)
					err = e.process(fieldVal, fieldTyp.Type)
					if err != nil {
						//file, line, funcName := runTimeUtil.Trace()
						//err = errors.Join(fmt.Errorf("%v(line: %v).process().error: %v", funcName, line, err))
						//err = errors.Join(fmt.Errorf("file: %v", file), err)
						return
					}
				}
			}
		}
	}

	return
}

func (e *Components) processComponent(element reflect.Value, typeof reflect.Type, tagData *tag, father *html.TagDiv) (err error) {

	if element.CanInterface() {

		for {
			if element.Kind() != reflect.Pointer {
				break
			}

			if element.CanSet() && element.IsNil() {
				newInstance := reflect.New(element.Type().Elem())
				element.Set(newInstance)
			}

			element = element.Elem()

		}

		for i := 0; i != element.NumField(); i += 1 {
			divComponent := factoryBrowser.NewTagDiv().Class("component")

			var fieldTyp reflect.StructField
			fieldVal := element.Field(i)
			if typeof.Kind() == reflect.Pointer {
				fieldTyp = typeof.Elem().Field(i)
			} else {
				fieldTyp = typeof.Field(i)
			}

			tagRaw := fieldTyp.Tag.Get("wasmPanel")
			if tagRaw != "" {
				tagData := new(tag)
				tagData.init(tagRaw)

				switch tagData.Type {
				case "range":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Range (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:range;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Range (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentRange(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "button":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Button (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:button;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Button (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentButton(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "select":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Select (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:select;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Select (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentSelect(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "radio":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Radio (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:radio;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Radio (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentRadio(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "checkbox":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Checkbox (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:checkbox;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Checkbox (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentCheckbox(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "color":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Color (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:color;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Color (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentColor(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "date":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Date (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:date;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Date (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentDate(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "week":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Week (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:week;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Week (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentWeek(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "text":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Text (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:text;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Text (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentText(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "time":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Time (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:time;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Time (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentTime(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "month":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Month (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:month;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Month (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentMonth(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "password":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Password (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:password;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Password (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentPassword(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "mail":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.Mail (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:mail;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.Mail (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentMail(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				case "textArea":

					if fieldVal.Kind() != reflect.Pointer {
						err = fmt.Errorf("component.TextArea (%v) requires a pointer to the component, example", fieldVal.Type().Name())
						err = errors.Join(err, fmt.Errorf("type %v struct {", element.Type().Name()))
						err = errors.Join(err, fmt.Errorf("  %v *%v `wasmPanel:\"type:text;label:...\"`", fieldTyp.Name, fieldVal.Type().Name()))
						err = errors.Join(err, fmt.Errorf("}"))
						return
					}

					if !fieldVal.CanSet() || !fieldVal.CanInterface() {
						err = fmt.Errorf("component.TextArea (%v) requires a public field, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentTextArea(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}

				default:
					//log.Printf(">>>>>>>> %v", tagData.Type)
				}
			}

			father.Append(divComponent)
		}
	}

	return
}

func (e *Components) processLabelCel(label string, father *html.TagDiv) {
	// <div class="labelCel">
	//   <div class="labelText">Label</div>
	//   <div class="closeIcon">ˇ</div>
	// </div>
	father.Append(
		factoryBrowser.NewTagDiv().Class("labelCel").Append(
			factoryBrowser.NewTagDiv().Class("labelText").Text(label),
			factoryBrowser.NewTagDiv().Class("closeIcon").Text("ˇ"),
		),
	)
}

//func (e *Components) callFunc(funcObj reflect.Value, params ...interface{}) (results []interface{}, err error) {
//
//	if funcObj.Kind() != reflect.Func {
//		return nil, fmt.Errorf("funcObj is not of type reflect.Func")
//	}
//
//	if len(params) != funcObj.Type().NumIn() {
//		return nil, fmt.Errorf("incorrect number of parameters")
//	}
//
//	in := make([]reflect.Value, len(params))
//	for i, param := range params {
//		in[i] = reflect.ValueOf(param)
//	}
//
//	out := funcObj.Call(in)
//
//	results = make([]interface{}, len(out))
//	for i, result := range out {
//		results[i] = result.Interface()
//	}
//	return results, nil
//}

// searchFieldByTagType Procura um campo pelo tipo da tag
func (e *Components) searchFieldByTagType(typeName, eventName string, element reflect.Value) (fieldVal reflect.Value, fieldTyp reflect.StructField, found bool) {
	for i := 0; i != element.NumField(); i += 1 {
		fieldVal = element.Field(i)
		fieldTyp = reflect.TypeOf(element.Interface()).Field(i)

		tagDataInternal := new(tag)
		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal.init(tagRaw)
		}

		if tagDataInternal.Type == typeName && eventName == "" {
			found = true
			return
		}

		if tagDataInternal.Type == typeName && tagDataInternal.Event == eventName {
			found = true
			return
		}
	}

	return
}

func (e *Components) verifyTypeFromElement(fieldVal reflect.Value, valueType reflect.Kind) (dataType reflect.Kind, value any, ok bool) {
	switch valueType {
	case reflect.Int64:
		dataType = reflect.Int64
		value = fieldVal.Int()
	case reflect.Float64:
		dataType = reflect.Float64
		value = fieldVal.Float()
	case reflect.String:
		dataType = reflect.String
		value = fieldVal.String()
	case reflect.Bool:
		dataType = reflect.Bool
		value = fieldVal.Bool()
	default:
		dataType = reflect.Invalid
		return
	}

	ok = true
	return
}

//func (e *Components) verifyTypeNumericFromElement(fieldVal reflect.Value, valueType reflect.Kind) (dataType reflect.Kind, value any, ok bool) {
//	switch valueType {
//	case reflect.Int64:
//		dataType = reflect.Int64
//		value = fieldVal.Int()
//	case reflect.Float64:
//		dataType = reflect.Float64
//		value = fieldVal.Float()
//	default:
//		dataType = reflect.Invalid
//		return
//	}
//
//	ok = true
//	return
//}

func (e *Components) processComponentRange(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	rangeComponent := Range{}

	inputRange := factoryBrowser.NewTagInputRange().Class("inputRange")
	inputNumber := factoryBrowser.NewTagInputNumber().Class("inputNumber") //.Min(eMin).Max(eMax).Step(eStep).Value(rangeVal).ListenerAdd(generic.KEventChange, captureData, listenerFunc).ListenerDebug(true)

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Range` was done
	if fieldRange := element.FieldByName("Range"); !fieldRange.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Range` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Range"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value int64 `wasmPanel:\"type:value;min:0;max:50;step:1;default:0\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Range
		newInstance := reflect.New(fieldRange.Type())
		fieldRange.Set(newInstance.Elem())

		// Initializes the two input tags within Range
		rangeComponent.__rangeTag = inputRange
		rangeComponent.__numberTag = inputNumber

		// __rangeOnInputEvent is the pointer sent when the `change` event happens
		rangeComponent.__change = new(__rangeOnInputEvent)

		// populates the component.Range within the user component
		componentRange := element.FieldByName("Range")
		componentRange.Set(reflect.ValueOf(rangeComponent))
	}

	// Search for the listener input tag and if it does not exist, set up the controller control function
	if _, _, found := e.searchFieldByTagType("listener", "input", element); !found {
		var methods []reflect.Value
		var params []interface{}

		// Passes the functions to be executed in the listener
		methods = []reflect.Value{
			// rangeComponent is the struct components.Range and OnChangeNumber is a function belonging to the struct components.Range
			reflect.ValueOf(&rangeComponent).MethodByName("OnChangeNumber"),
			// rangeComponent is the struct components.Range and OnChangeRange is a function belonging to the struct components.Range
			reflect.ValueOf(&rangeComponent).MethodByName("OnChangeRange"),
		}

		// Pass variable pointers
		params = []interface{}{
			// __rangeOnInputEvent is the type pointer contained in components.Range and collects value
			new(__rangeOnInputEvent),
			// __rangeOnInputEvent is the type pointer contained in components.Range and collects value
			new(__rangeOnInputEvent),
		}

		// explanation
		//   inputNumber.ListenerAdd() accepts two arrays, one for the function to be invoked, and the other with the data to be passed
		inputRange.ListenerAddReflect("input", params, methods, element.Interface())
		inputNumber.ListenerAddReflect("input", params, methods, element.Interface())
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {

			// Checks whether the reference to the input range tag was requested by the user
			case "inputTagRange":
				fieldVal.Set(reflect.ValueOf(inputRange))
				//fieldVal.Interface().(*html.TagInputRange).Min(-5).Max(10).Value(-5)

			// Checks whether the reference to the input number tag was requested by the user
			case "inputTagNumber":
				fieldVal.Set(reflect.ValueOf(inputNumber))
				//fieldVal.Interface().(*html.TagInputNumber).Min(-5).Max(10).Value(-5)

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.Float64:
					if value.(float64) != 0 {
						passValue = true
					}

				case reflect.Int64:
					if value.(int64) != 0 {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type int64 or float64", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Fill in the numeric fields
				inputRange.Value(value)
				inputNumber.Value(value)

				// If the value is zero, and the user has determined a value other than zero,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputRange.Value(tagDataInternal.Default)
					inputNumber.Value(tagDataInternal.Default)
				}

				// todo: isto não tem sentido aqui - inicio
				switch dataType {
				case reflect.Float64:
					if tagDataInternal.Min != "" {
						eMin, _ := strconv.ParseFloat(tagDataInternal.Min, 64)
						inputRange.Min(eMin)
						inputNumber.Min(eMin)
					}
					if tagDataInternal.Max != "" {
						eMax, _ := strconv.ParseFloat(tagDataInternal.Max, 64)
						inputRange.Max(eMax)
						inputNumber.Max(eMax)
					}
					if tagDataInternal.Step != "" {
						eStep, _ := strconv.ParseFloat(tagDataInternal.Step, 64)
						inputRange.Step(eStep)
						inputNumber.Step(eStep)
					}
				case reflect.Int64:
					if tagDataInternal.Min != "" {
						eMin, _ := strconv.ParseInt(tagDataInternal.Min, 10, 64)
						inputRange.Min(eMin)
						inputNumber.Min(eMin)
					}
					if tagDataInternal.Max != "" {
						eMax, _ := strconv.ParseInt(tagDataInternal.Max, 10, 64)
						inputRange.Max(eMax)
						inputNumber.Max(eMax)
					}
					if tagDataInternal.Step != "" {
						eStep, _ := strconv.ParseInt(tagDataInternal.Step, 10, 64)
						inputRange.Step(eStep)
						inputNumber.Step(eStep)
					}
				default:

				}
				// todo: isto não tem sentido aqui - fim

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				// explanation
				//   inputNumber.ListenerAdd() accepts two arrays, one for the function to be invoked, and the other with the data to be passed
				//   The first element of the array is the user function
				//   From the second element onwards, they are internal functions and must be called after the user function in case the user has changed any value.
				//inputRange.ListenerAddReflect(tagDataInternal.Event, params, methods, element.Interface())
				//inputNumber.ListenerAddReflect(tagDataInternal.Event, params, methods, element.Interface())

				inputRange.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
				inputNumber.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)

			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagDataFather.Label),
		inputRange,
		inputNumber,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Range{}) {
			r := fieldVal.Interface().(Range)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentColor(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	colorComponent := Color{}

	inputColor := factoryBrowser.NewTagInputColor().Class("component .component-color")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Color` was done
	if fieldColor := element.FieldByName("Color"); !fieldColor.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Color` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Color"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed color\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Color
		newInstance := reflect.New(fieldColor.Type())
		fieldColor.Set(newInstance.Elem())

		// Initializes the two input tags within Color
		colorComponent.__colorTag = inputColor

		// __colorOnInputEvent is the pointer sent when the `change` event happens
		colorComponent.__change = new(__colorOnInputEvent)

		// populates the component.Color within the user component
		componentColor := element.FieldByName("Color")
		componentColor.Set(reflect.ValueOf(colorComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {

			// Checks whether the reference to the input color tag was requested by the user
			case "inputTagColor":
				fieldVal.Set(reflect.ValueOf(inputColor))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputColor.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputColor.Value(tagDataInternal.Default)
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputColor.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagDataFather.Label),
		inputColor,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Color{}) {
			r := fieldVal.Interface().(Color)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentDate(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	dateComponent := Date{}

	inputDate := factoryBrowser.NewTagInputDate().Class("component .component-date")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Date` was done
	if fieldDate := element.FieldByName("Date"); !fieldDate.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Date` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Date"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed date;placeHolder:Place holder date\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Date
		newInstance := reflect.New(fieldDate.Type())
		fieldDate.Set(newInstance.Elem())

		// Initializes the two input tags within Date
		dateComponent.__dateTag = inputDate

		// __dateOnInputEvent is the pointer sent when the `change` event happens
		dateComponent.__change = new(__dateOnInputEvent)

		// populates the component.Date within the user component
		componentDate := element.FieldByName("Date")
		componentDate.Set(reflect.ValueOf(dateComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {

			// Checks whether the reference to the input date tag was requested by the user
			case "inputTagDate":
				fieldVal.Set(reflect.ValueOf(inputDate))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputDate.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputDate.Value(tagDataInternal.Default)
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputDate.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagDataFather.Label),
		inputDate,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Date{}) {
			r := fieldVal.Interface().(Date)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentText(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	textComponent := Text{}

	inputText := factoryBrowser.NewTagInputText().Class("component .component-text")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Text` was done
	if fieldText := element.FieldByName("Text"); !fieldText.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Text` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Text"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed text;placeHolder:Place holder text\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Text
		newInstance := reflect.New(fieldText.Type())
		fieldText.Set(newInstance.Elem())

		// Initializes the two input tags within Text
		textComponent.__textTag = inputText

		// __textOnInputEvent is the pointer sent when the `change` event happens
		textComponent.__change = new(__textOnInputEvent)

		// populates the component.Text within the user component
		componentText := element.FieldByName("Text")
		componentText.Set(reflect.ValueOf(textComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {

			// Checks whether the reference to the input text tag was requested by the user
			case "inputTagText":
				fieldVal.Set(reflect.ValueOf(inputText))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputText.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputText.Value(tagDataInternal.Default)
				}

				inputText.Placeholder(tagDataInternal.PlaceHolder)

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputText.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagDataFather.Label),
		inputText,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Text{}) {
			r := fieldVal.Interface().(Text)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentTime(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	timeComponent := Time{}

	inputTime := factoryBrowser.NewTagInputTime().Class("component .component-time")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Time` was done
	if fieldTime := element.FieldByName("Time"); !fieldTime.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Time` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Time"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed time;placeHolder:Place holder time\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Time
		newInstance := reflect.New(fieldTime.Type())
		fieldTime.Set(newInstance.Elem())

		// Initializes the two input tags within Time
		timeComponent.__timeTag = inputTime

		// __timeOnInputEvent is the pointer sent when the `change` event happens
		timeComponent.__change = new(__timeOnInputEvent)

		// populates the component.Time within the user component
		componentTime := element.FieldByName("Time")
		componentTime.Set(reflect.ValueOf(timeComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {

			// Checks whether the reference to the input time tag was requested by the user
			case "inputTagTime":
				fieldVal.Set(reflect.ValueOf(inputTime))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputTime.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					d, _ := time.ParseDuration(tagDataInternal.Default)
					inputTime.Value(Timespan(d).Format(time.TimeOnly))
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputTime.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagDataFather.Label),
		inputTime,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Time{}) {
			r := fieldVal.Interface().(Time)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentMonth(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	monthComponent := Month{}

	inputMonth := factoryBrowser.NewTagInputMonth().Class("component .component-month")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Month` was done
	if fieldMonth := element.FieldByName("Month"); !fieldMonth.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Month` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Month"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed month;placeHolder:Place holder month\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Month
		newInstance := reflect.New(fieldMonth.Type())
		fieldMonth.Set(newInstance.Elem())

		// Initializes the two input tags within Month
		monthComponent.__monthTag = inputMonth

		// __monthOnInputEvent is the pointer sent when the `change` event happens
		monthComponent.__change = new(__monthOnInputEvent)

		// populates the component.Month within the user component
		componentMonth := element.FieldByName("Month")
		componentMonth.Set(reflect.ValueOf(monthComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {

			// Checks whether the reference to the input month tag was requested by the user
			case "inputTagMonth":
				fieldVal.Set(reflect.ValueOf(inputMonth))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputMonth.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputMonth.Value(tagDataInternal.Default)
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputMonth.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagDataFather.Label),
		inputMonth,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Month{}) {
			r := fieldVal.Interface().(Month)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentWeek(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	weekComponent := Week{}

	inputWeek := factoryBrowser.NewTagInputWeek().Class("component .component-week")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Week` was done
	if fieldWeek := element.FieldByName("Week"); !fieldWeek.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Week` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Week"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed week;placeHolder:Place holder week\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Week
		newInstance := reflect.New(fieldWeek.Type())
		fieldWeek.Set(newInstance.Elem())

		// Initializes the two input tags within Week
		weekComponent.__weekTag = inputWeek

		// __weekOnInputEvent is the pointer sent when the `change` event happens
		weekComponent.__change = new(__weekOnInputEvent)

		// populates the component.Week within the user component
		componentWeek := element.FieldByName("Week")
		componentWeek.Set(reflect.ValueOf(weekComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {

			// Checks whether the reference to the input week tag was requested by the user
			case "inputTagWeek":
				fieldVal.Set(reflect.ValueOf(inputWeek))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputWeek.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputWeek.Value(tagDataInternal.Default)
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputWeek.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagDataFather.Label),
		inputWeek,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Week{}) {
			r := fieldVal.Interface().(Week)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentPassword(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	passwordComponent := Password{}

	inputPassword := factoryBrowser.NewTagInputPassword().Class("component .component-password")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Password` was done
	if fieldPassword := element.FieldByName("Password"); !fieldPassword.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Password` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Password"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed password;placeHolder:Place holder text\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Password
		newInstance := reflect.New(fieldPassword.Type())
		fieldPassword.Set(newInstance.Elem())

		// Initializes the two input tags within Password
		passwordComponent.__passwordTag = inputPassword

		// __passwordOnInputEvent is the pointer sent when the `change` event happens
		passwordComponent.__change = new(__passwordOnInputEvent)

		// populates the component.Password within the user component
		componentPassword := element.FieldByName("Password")
		componentPassword.Set(reflect.ValueOf(passwordComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {

			// Checks whether the reference to the input password tag was requested by the user
			case "inputTagPassword":
				fieldVal.Set(reflect.ValueOf(inputPassword))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputPassword.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputPassword.Value(tagDataInternal.Default)
				}

				inputPassword.Placeholder(tagDataInternal.PlaceHolder)

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputPassword.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagDataFather.Label),
		inputPassword,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Password{}) {
			r := fieldVal.Interface().(Password)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentMail(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	mailComponent := Mail{}

	inputMail := factoryBrowser.NewTagInputEMail().Class("component .component-mail")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Mail` was done
	if fieldMail := element.FieldByName("Mail"); !fieldMail.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Mail` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Mail"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed Mail;placeHolder:Place holder text\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Mail
		newInstance := reflect.New(fieldMail.Type())
		fieldMail.Set(newInstance.Elem())

		// Initializes the two input tags within Mail
		mailComponent.__mailTag = inputMail

		// __mailOnInputEvent is the pointer sent when the `change` event happens
		mailComponent.__change = new(__mailOnInputEvent)

		// populates the component.Mail within the user component
		componentMail := element.FieldByName("Mail")
		componentMail.Set(reflect.ValueOf(mailComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {

			// Checks whether the reference to the input mail tag was requested by the user
			case "inputTagMail":
				fieldVal.Set(reflect.ValueOf(inputMail))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputMail.Value(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputMail.Value(tagDataInternal.Default)
				}

				inputMail.Placeholder(tagDataInternal.PlaceHolder)

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputMail.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagDataFather.Label),
		inputMail,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Mail{}) {
			r := fieldVal.Interface().(Mail)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentTextArea(element reflect.Value, tagDataFather *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	textAreaComponent := TextArea{}

	inputTextArea := factoryBrowser.NewTagTextArea().Class("component .component-textarea")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.TextArea` was done
	if fieldTextArea := element.FieldByName("TextArea"); !fieldTextArea.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.TextArea` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.TextArea"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmPanel:\"type:value;default:Predefined fixed text;placeHolder:Place holder text\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize TextArea
		newInstance := reflect.New(fieldTextArea.Type())
		fieldTextArea.Set(newInstance.Elem())

		// Initializes the textArea tag within TextArea
		textAreaComponent.__textAreaTag = inputTextArea

		// __textAreaOnInputEvent is the pointer sent when the `change` event happens
		textAreaComponent.__change = new(__textAreaOnInputEvent)

		// populates the component.TextArea within the user component
		componentTextArea := element.FieldByName("TextArea")
		componentTextArea.Set(reflect.ValueOf(textAreaComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {

			// Checks whether the reference to the textArea tag was requested by the user
			case "inputTagTextArea":
				fieldVal.Set(reflect.ValueOf(inputTextArea))

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				inputTextArea.Text(value)

				// If the value is zero, and the user has determined a value other than blank,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Default != "" {
					inputTextArea.Text(tagDataInternal.Default)
				}

				inputTextArea.Placeholder(tagDataInternal.PlaceHolder)

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					fieldVal.Interface(),
				}

				inputTextArea.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagDataFather.Label),
		inputTextArea,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(TextArea{}) {
			r := fieldVal.Interface().(TextArea)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentButton(element reflect.Value, tagData *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	elementOriginal := element
	buttonComponent := Button{}

	inputButton := factoryBrowser.NewTagInputButton().Class("inputButton") // .Value("-- ok --")

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move element to pointer struct
	element = element.Elem()

	// Checks if the import of `components.Range` was done
	if fieldRange := element.FieldByName("Button"); !fieldRange.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Button` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type OnClickEvent struct {"))
		err = errors.Join(err, fmt.Errorf("         IsTrusted bool `wasmGet:\"isTrusted\"`"))
		err = errors.Join(err, fmt.Errorf("         Value     string `wasmGet:\"value\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *OnClickEvent, ref %v) OnClick(event OnClickEvent) {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         log.Printf(\"Trusted: %%v\", event.IsTrusted)"))
		err = errors.Join(err, fmt.Errorf("         log.Printf(\"Value:   %%v\", event.Value)"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Button"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         Label    string        `wasmPanel:\"type:value;default:Ok\"`"))
		err = errors.Join(err, fmt.Errorf("         RunEvent *OnClickEvent `wasmPanel:\"type:listener;event:click;func:OnClick\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Range
		newInstance := reflect.New(fieldRange.Type())
		fieldRange.Set(newInstance.Elem())

		// Initializes the two input tags within Range
		buttonComponent.__buttonTag = inputButton

		// populates the component.Range within the user component
		componentButton := element.FieldByName("Button")
		componentButton.Set(reflect.ValueOf(buttonComponent))
	}

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {

			// Checks whether the reference to the input range tag was requested by the user
			case "inputTagButton":
				fieldVal.Set(reflect.ValueOf(inputButton))
				//fieldVal.Interface().(*html.TagInputButton)

			// Checks if the value tag was created
			case "value":

				// Captures the value of the component defined by the value tag
				dataType, value, ok = e.verifyTypeFromElement(fieldVal, fieldVal.Kind())
				if !ok {
					err = fmt.Errorf("%v.%v type '%v', must be a type int64, float64, bool or string", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Checks if the field is non-zero, i.e. defined by the user
				// Limits the types accepted by numeric fields
				// The limitation on int64, float64, string and bool types is determined by the golang webassembly
				passValue := false
				switch dataType {
				case reflect.String:
					if value.(string) != "" {
						passValue = true
					}

				default:
					err = fmt.Errorf("%v.%v type '%v', must be a type int64 or float64", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
					return
				}

				// Fill in the numeric fields
				inputButton.Value(value)

				// If the value is zero, and the user has determined a value other than zero,
				// fill in the field with the default value
				if !passValue && tagDataInternal.Label != "" {
					inputButton.Value(tagDataInternal.Label)
				}

			// listener defines the field received by the event function
			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					// tagDataInternal.Func is the user function
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					// fieldVal.Interface() is the struct pointer that collects user data
					fieldVal.Interface(),
				}

				// explanation
				//   inputNumber.ListenerAdd() accepts two arrays, one for the function to be invoked, and the other with the data to be passed
				//   The first element of the array is the user function
				//   From the second element onwards, they are internal functions and must be called after the user function in case the user has changed any value.
				//inputButton.ListenerAddReflect(tagDataInternal.Event, params, methods, element.Interface())
				inputButton.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)

			}

		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagData.Label),
		inputButton,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Button{}) {
			r := fieldVal.Interface().(Button)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

// verifyTypesComponentSelect checks the types of the components.Select configuration
func (e *Components) verifyTypesComponentSelect(element reflect.Value) (err error) {

	elemTpl := element.Type()
	elemTplOriginal := element.Type()
	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := elemTpl.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {
			case "value":

				if fieldVal.Kind() != reflect.Pointer {
					err = fmt.Errorf("the field %v, inside %v, must be a pointer of slice struct", fieldTyp.Name, elemTplOriginal.Name())
					err = errors.Join(err, fmt.Errorf("       Example:"))
					err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
					err = errors.Join(err, fmt.Errorf("         components.Select"))
					err = errors.Join(err, fmt.Errorf("         "))
					err = errors.Join(err, fmt.Errorf("         %v *[]SelectData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name))
					err = errors.Join(err, fmt.Errorf("       }"))
					err = errors.Join(err, fmt.Errorf("       "))
					err = errors.Join(err, fmt.Errorf("       type SelectData struct {"))
					err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
					err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
					err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
					err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
					err = errors.Join(err, fmt.Errorf("       }"))

					return
				}

				if fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.Elem().Kind() != reflect.Slice {
					err = fmt.Errorf("the field %v, inside %v, must be a pointer of slice struct", fieldTyp.Name, elemTplOriginal.Name())
					err = errors.Join(err, fmt.Errorf("       Example:"))
					err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
					err = errors.Join(err, fmt.Errorf("         components.Select"))
					err = errors.Join(err, fmt.Errorf("         "))
					err = errors.Join(err, fmt.Errorf("         %v *[]SelectData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name))
					err = errors.Join(err, fmt.Errorf("       }"))
					err = errors.Join(err, fmt.Errorf("       "))
					err = errors.Join(err, fmt.Errorf("       type SelectData struct {"))
					err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
					err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
					err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
					err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
					err = errors.Join(err, fmt.Errorf("       }"))

					return
				}

				fieldVal = fieldVal.Elem()
				fieldTpl := fieldVal.Type().Elem()

				if fieldTpl.Kind() != reflect.Struct {
					err = fmt.Errorf("the field %v, inside %v, must be a pointer of slice struct", fieldTyp.Name, elemTplOriginal.Name())
					err = errors.Join(err, fmt.Errorf("       Example:"))
					err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
					err = errors.Join(err, fmt.Errorf("         components.Select"))
					err = errors.Join(err, fmt.Errorf("         "))
					err = errors.Join(err, fmt.Errorf("         %v *[]SelectData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name))
					err = errors.Join(err, fmt.Errorf("       }"))
					err = errors.Join(err, fmt.Errorf("       "))
					err = errors.Join(err, fmt.Errorf("       type SelectData struct {"))
					err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
					err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
					err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
					err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
					err = errors.Join(err, fmt.Errorf("       }"))

					return
				}

				for k := 0; k != fieldTpl.NumField(); k += 1 {
					fieldTyp := fieldTpl.Field(k)

					tagRaw := fieldTyp.Tag.Get("wasmPanel")
					if tagRaw != "" {
						tagDataInternal := new(tag)
						tagDataInternal.init(tagRaw)

						switch tagDataInternal.Type {
						case "label":
							if fieldTyp.Type.Kind() != reflect.String {
								err = fmt.Errorf("the tag type:%v, inside %v, must be a string", tagDataInternal.Type, fieldTpl.Name())
								err = errors.Join(err, fmt.Errorf("       Example:"))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
								err = errors.Join(err, fmt.Errorf("         components.Select"))
								err = errors.Join(err, fmt.Errorf("         "))
								err = errors.Join(err, fmt.Errorf("         %v *[]%v `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name, fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("       }"))
								err = errors.Join(err, fmt.Errorf("       "))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
								err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
								err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
								err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
								err = errors.Join(err, fmt.Errorf("       }"))

								return
							}
						case "value":
							if fieldTyp.Type.Kind() != reflect.String {
								err = fmt.Errorf("the tag type:%v, inside %v, must be a string", tagDataInternal.Type, fieldTpl.Name())
								err = errors.Join(err, fmt.Errorf("       Example:"))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
								err = errors.Join(err, fmt.Errorf("         components.Select"))
								err = errors.Join(err, fmt.Errorf("         "))
								err = errors.Join(err, fmt.Errorf("         %v *[]%v `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name, fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("       }"))
								err = errors.Join(err, fmt.Errorf("       "))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
								err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
								err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
								err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
								err = errors.Join(err, fmt.Errorf("       }"))

								return
							}
						case "disabled":
							if fieldTyp.Type.Kind() != reflect.Bool {
								err = fmt.Errorf("the tag type:%v, inside %v, must be a boolean", tagDataInternal.Type, fieldTpl.Name())
								err = errors.Join(err, fmt.Errorf("       Example:"))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
								err = errors.Join(err, fmt.Errorf("         components.Select"))
								err = errors.Join(err, fmt.Errorf("         "))
								err = errors.Join(err, fmt.Errorf("         %v *[]%v `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name, fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("       }"))
								err = errors.Join(err, fmt.Errorf("       "))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
								err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
								err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
								err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
								err = errors.Join(err, fmt.Errorf("       }"))

								return
							}
						case "selected":
							if fieldTyp.Type.Kind() != reflect.Bool {
								err = fmt.Errorf("the tag type:%v, inside %v, must be a boolean", tagDataInternal.Type, fieldTpl.Name())
								err = errors.Join(err, fmt.Errorf("       Example:"))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", elemTplOriginal.Name()))
								err = errors.Join(err, fmt.Errorf("         components.Select"))
								err = errors.Join(err, fmt.Errorf("         "))
								err = errors.Join(err, fmt.Errorf("         %v *[]%v `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label N,value N\"`", fieldTyp.Name, fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("       }"))
								err = errors.Join(err, fmt.Errorf("       "))
								err = errors.Join(err, fmt.Errorf("       type %v struct {", fieldTpl.Name()))
								err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
								err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
								err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"`"))
								err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"`"))
								err = errors.Join(err, fmt.Errorf("       }"))

								return
							}
						}

					}
				}
				return
			}

		}

	}

	return
}

func (e *Components) processComponentSelect(element reflect.Value, tagData *tag, father *html.TagDiv) (err error) {

	inputSelect := factoryBrowser.NewTagSelect().Class("inputSelect")

	elementOriginal := element
	selectComponent := Select{}

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move the element from pointer to struct
	element = element.Elem()

	// Checks if the import of `components.Select` was done
	if fieldSelect := element.FieldByName("Select"); !fieldSelect.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Select` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Select"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         List *[]SelectData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label 3,value 3\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       type SelectData struct {"))
		err = errors.Join(err, fmt.Errorf("         Label    string `wasmPanel:\"type:label\"`"))
		err = errors.Join(err, fmt.Errorf("         Value    string `wasmPanel:\"type:value\"`"))
		err = errors.Join(err, fmt.Errorf("         Disabled bool   `wasmPanel:\"type:disabled\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Selected bool   `wasmPanel:\"type:selected\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       // Note: Use `>` to set value as selected. ie. >label,value"))
		return
	} else {
		// Initialize Select
		newInstance := reflect.New(fieldSelect.Type())
		fieldSelect.Set(newInstance.Elem())

		// Initializes the input tags within Select
		selectComponent.__selectTag = inputSelect

		// __selectOnInputEvent is the pointer sent when the `change` event happens
		selectComponent.__change = new(__selectOnInputEvent)

		// populates the component.Select within the user component
		componentRange := element.FieldByName("Select")
		componentRange.Set(reflect.ValueOf(selectComponent))
	}

	err = e.verifyTypesComponentSelect(element)
	if err != nil {
		return
	}

	fieldTyp := element.Type()
	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := fieldTyp.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {
			case "inputTagSelect":
				fieldVal.Set(reflect.ValueOf(inputSelect))

			case "value":

				// pointer is not nil
				// Move the element from pointer to struct
				fieldVal = fieldVal.Elem()

				if fieldVal.Len() == 0 {

					if tagDataInternal.Default != "" {
						optionList := strings.Split(tagDataInternal.Default, ",")
						if len(optionList)%2 != 0 {
							err = fmt.Errorf("%v.%v: the correct format from tag value is: `wasmPanel:\"type:value;default:label1,value1,label2,value2,labelN,valueN\"`, where value and label, must be a pair", element.Type().Name(), fieldTyp.Name)
							return
						}

						for k := 0; k != len(optionList); k += 2 {
							// if label start with `>` the option is selected
							selected := false
							if strings.HasPrefix(optionList[k], ">") {
								optionList[k] = optionList[k][1:]
								selected = true
							}
							inputSelect.NewOption(optionList[k], optionList[k+1], false, selected)
						}
					}

				} else {

					// run inside slice data
					for iField := 0; iField != fieldVal.Len(); iField += 1 {
						keyVal := fieldVal.Index(iField)

						// get label, value, disabled and selected
						var label, value string
						var disabled, selected bool
						for ik := 0; ik != keyVal.NumField(); ik += 1 {
							optionVal := keyVal.Field(ik)
							optionTyp := reflect.TypeOf(keyVal.Interface()).Field(ik)

							optionTagRaw := optionTyp.Tag.Get("wasmPanel")
							if optionTagRaw != "" {
								optionTag := new(tag)
								optionTag.init(optionTagRaw)

								switch optionTag.Type {
								case "label":
									label = optionVal.Interface().(string)
								case "value":
									value = optionVal.Interface().(string)
								case "disabled":
									disabled = optionVal.Interface().(bool)
								case "selected":
									selected = optionVal.Interface().(bool)
								}
							}
						}

						inputSelect.NewOption(label, value, disabled, selected)
					}
				}

			//

			case "listener":

				// The field must be a pointer, or it cannot be populated
				if fieldVal.Kind() != reflect.Pointer {
					log.Printf("error: %v deve ser um ponteiro", fieldVal.Type().Name())
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: %v não pode ser definido automaticamente.", fieldVal.Type().Name())
					log.Printf("         isto geralmente acontece quando %v não é público.", fieldVal.Type().Name())
					continue
				}

				// Checks if the field is nil and initializes the pointer
				// The less work for the user, the greater the chance they will like the system
				if fieldVal.CanSet() && fieldVal.IsNil() {
					newInstance := reflect.New(fieldVal.Type().Elem())
					fieldVal.Set(newInstance)
				}

				if fieldVal.IsNil() {
					err = fmt.Errorf("o campo %v não foi inicializado de forma correta. ele deve ser público", fieldVal.Type().Name())
					return
				}

				var methods []reflect.Value
				var params []interface{}

				// Passes the functions to be executed in the listener
				methods = []reflect.Value{
					// tagDataInternal.Func is the user function
					fieldVal.MethodByName(tagDataInternal.Func),
				}

				// Pass variable pointers
				params = []interface{}{
					// fieldVal.Interface() is the struct pointer that collects user data
					fieldVal.Interface(),
				}

				// explanation
				//   inputSelect.ListenerAdd() accepts two arrays, one for the function to be invoked, and the other with the data to be passed
				//   The first element of the array is the user function
				//   From the second element onwards, they are internal functions and must be called after the user function in case the user has changed any value.
				inputSelect.ListenerAddReflect(tagDataInternal.Event, params, methods, e.ref)
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagData.Label),
		inputSelect,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Select{}) {
			r := fieldVal.Interface().(Select)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentRadio(element reflect.Value, tagData *tag, father *html.TagDiv) (err error) {

	inputDivRadio := factoryBrowser.NewTagDiv().Class("inputRadio")

	elementOriginal := element
	radioComponent := Radio{}

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move the element from pointer to struct
	element = element.Elem()

	// Checks if the import of `components.Radio` was done
	if fieldRadio := element.FieldByName("Radio"); !fieldRadio.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Radio` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Radio"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         List *[]RadioData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label 3,value 3\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       type RadioData struct {"))
		err = errors.Join(err, fmt.Errorf("         TagRadio *html.TagInputRadio `wasmPanel:\"type:inputTagRadio\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         TagLabel *html.TagLabel      `wasmPanel:\"type:inputTagLabel\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Label    string              `wasmPanel:\"type:label\"`"))
		err = errors.Join(err, fmt.Errorf("         Value    string              `wasmPanel:\"type:value\"`"))
		err = errors.Join(err, fmt.Errorf("         Disabled bool                `wasmPanel:\"type:disabled\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Selected bool                `wasmPanel:\"type:selected\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Change   *RadioChange        `wasmPanel:\"type:listener;event:change;func:OnChangeEvent\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       // Note: Use `>` to set value as selected. ie. >label,value"))
		err = errors.Join(err, fmt.Errorf("       type RadioChange struct {"))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmGet:\"value\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *RadioChange) OnChangeEvent(event RadioChange, reference *Body) {"))
		err = errors.Join(err, fmt.Errorf("         log.Printf(\"value: %%v\", event.Value)"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Radio
		newInstance := reflect.New(fieldRadio.Type())
		fieldRadio.Set(newInstance.Elem())

		// Initializes the input tags within Radio
		//radioComponent.__radioTag = inputRadio // todo: fazer

		// __radioOnInputEvent is the pointer sent when the `change` event happens
		radioComponent.__change = new(__radioOnInputEvent)

		// populates the component.Radio within the user component
		componentRange := element.FieldByName("Radio")
		componentRange.Set(reflect.ValueOf(radioComponent))
	}

	err = e.verifyTypesComponentSelect(element) // todo: mudar este nome
	if err != nil {
		return
	}

	fieldNameInputTagLabel := ""
	fieldNameInputTagRadio := ""
	fieldNameLabel := ""
	fieldNameValue := ""
	fieldNameDisabled := ""
	fieldNameSelected := ""
	fieldNameListener := ""
	tagListener := new(tag)
	typeListener := reflect.StructField{}

	fieldTyp := element.Type()
	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := fieldTyp.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {
			case "value":

				// fieldVal.Interface() é *[]struct{...}, por isto .Elem().Elem(), ou *[] -> struct{...}
				fieldTyp := reflect.TypeOf(fieldVal.Interface()).Elem().Elem()
				for k := 0; k != fieldTyp.NumField(); k += 1 {
					fieldTyp := fieldTyp.Field(k)
					tagRaw := fieldTyp.Tag.Get("wasmPanel")
					if tagRaw != "" {
						tagDataInternal := new(tag)
						tagDataInternal.init(tagRaw)

						switch tagDataInternal.Type {
						case "inputTagLabel":
							fieldNameInputTagLabel = fieldTyp.Name
						case "inputTagRadio":
							fieldNameInputTagRadio = fieldTyp.Name
						case "label":
							fieldNameLabel = fieldTyp.Name
						case "value":
							fieldNameValue = fieldTyp.Name
						case "disabled":
							fieldNameDisabled = fieldTyp.Name
						case "selected":
							fieldNameSelected = fieldTyp.Name
						case "listener":
							fieldNameListener = fieldTyp.Name
							tagListener = tagDataInternal
							typeListener = fieldTyp
						}
					}
				}
			}
		}
	}

	//fieldTyp := element.Type()

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := fieldTyp.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {
			//case "inputTagRadio":
			//	fieldVal.Set(reflect.ValueOf(inputRadio))

			case "value":

				fieldValPointer := fieldVal

				// pointer is not nil
				// Move the element from pointer to struct
				fieldVal = fieldVal.Elem()

				var inputLabel *html.TagLabel
				var inputRadio *html.TagInputRadio

				if fieldVal.IsZero() {

					var sliceValue reflect.Value
					var sliceType reflect.Type

					// fieldVal.Interface() é *[]struct{...}, por isto .Elem(), ou * -> []struct{...}
					sliceValue = reflect.ValueOf(fieldValPointer.Interface()).Elem()
					sliceType = reflect.TypeOf(sliceValue.Interface())
					newSlice := reflect.MakeSlice(sliceType, 0, 0)
					sliceValue.Set(newSlice)

					elemType := sliceType.Elem()
					newElem := reflect.New(elemType).Elem()

					if tagDataInternal.Default != "" {
						optionList := strings.Split(tagDataInternal.Default, ",")
						if len(optionList)%2 != 0 {
							err = fmt.Errorf("%v.%v: the correct format from tag value is: `wasmPanel:\"type:value;default:label1,value1,label2,value2,labelN,valueN\"`, where value and label, must be a pair", element.Type().Name(), fieldTyp.Name)
							return
						}

						for k := 0; k != len(optionList); k += 2 {

							inputRadio = factoryBrowser.NewTagInputRadio()
							inputLabel = factoryBrowser.NewTagLabel()

							// if label start with `>` the option is selected
							selected := false
							if strings.HasPrefix(optionList[k], ">") {
								optionList[k] = optionList[k][1:]
								selected = true
							}

							inputRadio.Value(optionList[k+1]).Disabled(false).Checked(selected).Class("inputRadio").Name(tagDataInternal.Name)
							inputLabel.Text(optionList[k]).Append(inputRadio)

							if fieldNameInputTagLabel != "" {
								newElem.FieldByName(fieldNameInputTagLabel).Set(reflect.ValueOf(inputLabel))
							}

							if fieldNameInputTagRadio != "" {
								newElem.FieldByName(fieldNameInputTagRadio).Set(reflect.ValueOf(inputRadio))
							}

							if fieldNameLabel != "" {
								newElem.FieldByName(fieldNameLabel).SetString(optionList[k+1])
							}

							if fieldNameValue != "" {
								newElem.FieldByName(fieldNameValue).SetString(optionList[k])
							}

							if fieldNameDisabled != "" {
								newElem.FieldByName(fieldNameDisabled).SetBool(false)
							}

							if fieldNameSelected != "" {
								newElem.FieldByName(fieldNameDisabled).SetBool(selected)
							}

							if fieldNameListener != "" {
								// The field must be a pointer, or it cannot be populated
								if typeListener.Type.Kind() != reflect.Pointer {
									log.Printf("error: %v.%v deve ser um ponteiro", newElem.Type().Name(), typeListener.Type.Name())
									return
								}

								if !typeListener.IsExported() {
									log.Printf("error: %v.%v não pode ser definido automaticamente.", newElem.Type().Name(), fieldNameListener)
									log.Printf("         isto geralmente acontece quando %v.%v não é público.", newElem.Type().Name(), fieldNameListener)
									return
								}

								newInstance := reflect.New(typeListener.Type.Elem())
								newElem.FieldByName(fieldNameListener).Set(newInstance)

								var methods []reflect.Value
								var params []interface{}

								// Passes the functions to be executed in the listener
								methods = []reflect.Value{
									// tagDataInternal.Func is the user function
									newElem.FieldByName(fieldNameListener).MethodByName(tagListener.Func),
								}

								// Pass variable pointers
								params = []interface{}{
									// fieldVal.Interface() is the struct pointer that collects user data
									newElem.FieldByName(fieldNameListener).Interface(),
								}

								inputRadio.ListenerAddReflect(tagListener.Event, params, methods, e.ref)
							}

							sliceValue.Set(reflect.Append(sliceValue, newElem))

							inputDivRadio.Append(
								factoryBrowser.NewTagSpan().Append(inputLabel),
							)
						}
					}

				} else {

					// run inside slice data
					for iField := 0; iField != fieldVal.Len(); iField += 1 {
						keyVal := fieldVal.Index(iField)

						inputRadio = factoryBrowser.NewTagInputRadio()
						inputLabel = factoryBrowser.NewTagLabel()

						// get label, value, disabled and selected
						var label, value string
						var disabled, selected bool
						for ik := 0; ik != keyVal.NumField(); ik += 1 {

							optionVal := keyVal.Field(ik)
							optionTyp := reflect.TypeOf(keyVal.Interface()).Field(ik)

							optionTagRaw := optionTyp.Tag.Get("wasmPanel")
							if optionTagRaw != "" {
								optionTag := new(tag)
								optionTag.init(optionTagRaw)

								switch optionTag.Type {
								case "inputTagLabel":
									optionVal.Set(reflect.ValueOf(inputLabel))
								case "inputTagRadio":
									optionVal.Set(reflect.ValueOf(inputRadio))
								case "label":
									label = optionVal.Interface().(string)
								case "value":
									value = optionVal.Interface().(string)
								case "disabled":
									disabled = optionVal.Interface().(bool)
								case "selected":
									selected = optionVal.Interface().(bool)
								case "listener":
									// The field must be a pointer, or it cannot be populated
									if typeListener.Type.Kind() != reflect.Pointer {
										log.Printf("error: %v.%v deve ser um ponteiro", optionVal.Type().Name(), typeListener.Type.Name())
										return
									}

									if !typeListener.IsExported() {
										log.Printf("error: %v.%v não pode ser definido automaticamente.", optionVal.Type().Name(), fieldNameListener)
										log.Printf("         isto geralmente acontece quando %v.%v não é público.", optionVal.Type().Name(), fieldNameListener)
										return
									}

									newInstance := reflect.New(typeListener.Type.Elem())
									keyVal.FieldByName(fieldNameListener).Set(newInstance)

									var methods []reflect.Value
									var params []interface{}

									// Passes the functions to be executed in the listener
									methods = []reflect.Value{
										// tagDataInternal.Func is the user function
										keyVal.FieldByName(fieldNameListener).MethodByName(tagListener.Func),
									}

									// Pass variable pointers
									params = []interface{}{
										// fieldVal.Interface() is the struct pointer that collects user data
										keyVal.FieldByName(fieldNameListener).Interface(),
									}

									inputRadio.ListenerAddReflect(tagListener.Event, params, methods, e.ref)
								}

							}
						}

						inputRadio.Value(value).Disabled(disabled).Checked(selected).Class("inputRadio").Name(tagDataInternal.Name)
						inputLabel.Text(label).Append(inputRadio)

						inputDivRadio.Append(
							factoryBrowser.NewTagSpan().Append(inputLabel),
						)

						//inputSelect.NewOption(label, value, disabled, selected)
					}
				}

			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagData.Label),
		inputDivRadio,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Radio{}) {
			r := fieldVal.Interface().(Radio)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentCheckbox(element reflect.Value, tagData *tag, father *html.TagDiv) (err error) {

	inputDivCheckbox := factoryBrowser.NewTagDiv().Class("inputCheckbox")

	elementOriginal := element
	checkboxComponent := Checkbox{}

	// Initializes the pointer if it is nil
	if element.IsNil() {
		newInstance := reflect.New(element.Type().Elem())
		element.Set(newInstance)
	}

	// Move the element from pointer to struct
	element = element.Elem()

	// Checks if the import of `components.Checkbox` was done
	if fieldCheckbox := element.FieldByName("Checkbox"); !fieldCheckbox.IsValid() {
		err = fmt.Errorf("error: component %v needs to embed `components.Checkbox` directly", element.Type().Name())
		err = errors.Join(err, fmt.Errorf("       Example:"))
		err = errors.Join(err, fmt.Errorf("       type %v struct {", element.Type().Name()))
		err = errors.Join(err, fmt.Errorf("         components.Checkbox"))
		err = errors.Join(err, fmt.Errorf("         "))
		err = errors.Join(err, fmt.Errorf("         List *[]CheckboxData `wasmPanel:\"type:value;default:label 1,value 1,>label 2,value 2,label 3,value 3\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       type CheckboxData struct {"))
		err = errors.Join(err, fmt.Errorf("         TagCheckbox *html.TagInputCheckbox `wasmPanel:\"type:inputTagCheckbox\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         TagLabel *html.TagLabel      `wasmPanel:\"type:inputTagLabel\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Label    string              `wasmPanel:\"type:label\"`"))
		err = errors.Join(err, fmt.Errorf("         Value    string              `wasmPanel:\"type:value\"`"))
		err = errors.Join(err, fmt.Errorf("         Disabled bool                `wasmPanel:\"type:disabled\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Selected bool                `wasmPanel:\"type:selected\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("         Change   *CheckboxChange        `wasmPanel:\"type:listener;event:change;func:OnChangeEvent\"` // [optional]"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       // Note: Use `>` to set value as selected. ie. >label,value"))
		err = errors.Join(err, fmt.Errorf("       type CheckboxChange struct {"))
		err = errors.Join(err, fmt.Errorf("         Value string `wasmGet:\"value\"`"))
		err = errors.Join(err, fmt.Errorf("       }"))
		err = errors.Join(err, fmt.Errorf("       func (e *CheckboxChange) OnChangeEvent(event CheckboxChange, reference *Body) {"))
		err = errors.Join(err, fmt.Errorf("         log.Printf(\"value: %%v\", event.Value)"))
		err = errors.Join(err, fmt.Errorf("       }"))
		return
	} else {
		// Initialize Checkbox
		newInstance := reflect.New(fieldCheckbox.Type())
		fieldCheckbox.Set(newInstance.Elem())

		// Initializes the input tags within Checkbox
		//checkboxComponent.__checkboxTag = inputCheckbox // todo: fazer

		// __checkboxOnInputEvent is the pointer sent when the `change` event happens
		checkboxComponent.__change = new(__checkboxOnInputEvent)

		// populates the component.Checkbox within the user component
		componentRange := element.FieldByName("Checkbox")
		componentRange.Set(reflect.ValueOf(checkboxComponent))
	}

	err = e.verifyTypesComponentSelect(element) // todo: mudar este nome
	if err != nil {
		return
	}

	fieldNameInputTagLabel := ""
	fieldNameInputTagCheckbox := ""
	fieldNameLabel := ""
	fieldNameValue := ""
	fieldNameDisabled := ""
	fieldNameSelected := ""
	fieldNameListener := ""
	tagListener := new(tag)
	typeListener := reflect.StructField{}

	fieldTyp := element.Type()
	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := fieldTyp.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {
			case "value":

				// fieldVal.Interface() é *[]struct{...}, por isto .Elem().Elem(), ou *[] -> struct{...}
				fieldTyp := reflect.TypeOf(fieldVal.Interface()).Elem().Elem()
				for k := 0; k != fieldTyp.NumField(); k += 1 {
					fieldTyp := fieldTyp.Field(k)
					tagRaw := fieldTyp.Tag.Get("wasmPanel")
					if tagRaw != "" {
						tagDataInternal := new(tag)
						tagDataInternal.init(tagRaw)

						switch tagDataInternal.Type {
						case "inputTagLabel":
							fieldNameInputTagLabel = fieldTyp.Name
						case "inputTagCheckbox":
							fieldNameInputTagCheckbox = fieldTyp.Name
						case "label":
							fieldNameLabel = fieldTyp.Name
						case "value":
							fieldNameValue = fieldTyp.Name
						case "disabled":
							fieldNameDisabled = fieldTyp.Name
						case "selected":
							fieldNameSelected = fieldTyp.Name
						case "listener":
							fieldNameListener = fieldTyp.Name
							tagListener = tagDataInternal
							typeListener = fieldTyp
						}
					}
				}
			}
		}
	}

	//fieldTyp := element.Type()

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		fieldTyp := fieldTyp.Field(i)

		tagRaw := fieldTyp.Tag.Get("wasmPanel")
		if tagRaw != "" {
			tagDataInternal := new(tag)
			tagDataInternal.init(tagRaw)

			switch tagDataInternal.Type {
			//case "inputTagCheckbox":
			//	fieldVal.Set(reflect.ValueOf(inputCheckbox))

			case "value":

				fieldValPointer := fieldVal

				// pointer is not nil
				// Move the element from pointer to struct
				fieldVal = fieldVal.Elem()

				var inputLabel *html.TagLabel
				var inputCheckbox *html.TagInputCheckBox

				if fieldVal.IsZero() {

					var sliceValue reflect.Value
					var sliceType reflect.Type

					// fieldVal.Interface() é *[]struct{...}, por isto .Elem(), ou * -> []struct{...}
					sliceValue = reflect.ValueOf(fieldValPointer.Interface()).Elem()
					sliceType = reflect.TypeOf(sliceValue.Interface())
					newSlice := reflect.MakeSlice(sliceType, 0, 0)
					sliceValue.Set(newSlice)

					elemType := sliceType.Elem()
					newElem := reflect.New(elemType).Elem()

					if tagDataInternal.Default != "" {
						optionList := strings.Split(tagDataInternal.Default, ",")
						if len(optionList)%2 != 0 {
							err = fmt.Errorf("%v.%v: the correct format from tag value is: `wasmPanel:\"type:value;default:label1,value1,label2,value2,labelN,valueN\"`, where value and label, must be a pair", element.Type().Name(), fieldTyp.Name)
							return
						}

						for k := 0; k != len(optionList); k += 2 {

							inputCheckbox = factoryBrowser.NewTagInputCheckBox()
							inputLabel = factoryBrowser.NewTagLabel()

							// if label start with `>` the option is selected
							selected := false
							if strings.HasPrefix(optionList[k], ">") {
								optionList[k] = optionList[k][1:]
								selected = true
							}

							inputCheckbox.Value(optionList[k+1]).Disabled(false).Checked(selected).Class("inputCheckbox").Name(tagDataInternal.Name)
							inputLabel.Text(optionList[k]).Append(inputCheckbox)

							if fieldNameInputTagLabel != "" {
								newElem.FieldByName(fieldNameInputTagLabel).Set(reflect.ValueOf(inputLabel))
							}

							if fieldNameInputTagCheckbox != "" {
								newElem.FieldByName(fieldNameInputTagCheckbox).Set(reflect.ValueOf(inputCheckbox))
							}

							if fieldNameLabel != "" {
								newElem.FieldByName(fieldNameLabel).SetString(optionList[k+1])
							}

							if fieldNameValue != "" {
								newElem.FieldByName(fieldNameValue).SetString(optionList[k])
							}

							if fieldNameDisabled != "" {
								newElem.FieldByName(fieldNameDisabled).SetBool(false)
							}

							if fieldNameSelected != "" {
								newElem.FieldByName(fieldNameDisabled).SetBool(selected)
							}

							if fieldNameListener != "" {
								// The field must be a pointer, or it cannot be populated
								if typeListener.Type.Kind() != reflect.Pointer {
									log.Printf("error: %v.%v deve ser um ponteiro", newElem.Type().Name(), typeListener.Type.Name())
									return
								}

								if !typeListener.IsExported() {
									log.Printf("error: %v.%v não pode ser definido automaticamente.", newElem.Type().Name(), fieldNameListener)
									log.Printf("         isto geralmente acontece quando %v.%v não é público.", newElem.Type().Name(), fieldNameListener)
									return
								}

								newInstance := reflect.New(typeListener.Type.Elem())
								newElem.FieldByName(fieldNameListener).Set(newInstance)

								var methods []reflect.Value
								var params []interface{}

								// Passes the functions to be executed in the listener
								methods = []reflect.Value{
									// tagDataInternal.Func is the user function
									newElem.FieldByName(fieldNameListener).MethodByName(tagListener.Func),
								}

								// Pass variable pointers
								params = []interface{}{
									// fieldVal.Interface() is the struct pointer that collects user data
									newElem.FieldByName(fieldNameListener).Interface(),
								}

								inputCheckbox.ListenerAddReflect(tagListener.Event, params, methods, e.ref)
							}

							sliceValue.Set(reflect.Append(sliceValue, newElem))

							inputDivCheckbox.Append(
								factoryBrowser.NewTagSpan().Append(inputLabel),
							)
						}
					}

				} else {

					// run inside slice data
					for iField := 0; iField != fieldVal.Len(); iField += 1 {
						keyVal := fieldVal.Index(iField)

						inputCheckbox = factoryBrowser.NewTagInputCheckBox()
						inputLabel = factoryBrowser.NewTagLabel()

						// get label, value, disabled and selected
						var label, value string
						var disabled, selected bool
						for ik := 0; ik != keyVal.NumField(); ik += 1 {

							optionVal := keyVal.Field(ik)
							optionTyp := reflect.TypeOf(keyVal.Interface()).Field(ik)

							optionTagRaw := optionTyp.Tag.Get("wasmPanel")
							if optionTagRaw != "" {
								optionTag := new(tag)
								optionTag.init(optionTagRaw)

								switch optionTag.Type {
								case "inputTagLabel":
									optionVal.Set(reflect.ValueOf(inputLabel))
								case "inputTagCheckbox":
									optionVal.Set(reflect.ValueOf(inputCheckbox))
								case "label":
									label = optionVal.Interface().(string)
								case "value":
									value = optionVal.Interface().(string)
								case "disabled":
									disabled = optionVal.Interface().(bool)
								case "selected":
									selected = optionVal.Interface().(bool)
								case "listener":
									// The field must be a pointer, or it cannot be populated
									if typeListener.Type.Kind() != reflect.Pointer {
										log.Printf("error: %v.%v deve ser um ponteiro", optionVal.Type().Name(), typeListener.Type.Name())
										return
									}

									if !typeListener.IsExported() {
										log.Printf("error: %v.%v não pode ser definido automaticamente.", optionVal.Type().Name(), fieldNameListener)
										log.Printf("         isto geralmente acontece quando %v.%v não é público.", optionVal.Type().Name(), fieldNameListener)
										return
									}

									newInstance := reflect.New(typeListener.Type.Elem())
									keyVal.FieldByName(fieldNameListener).Set(newInstance)

									var methods []reflect.Value
									var params []interface{}

									// Passes the functions to be executed in the listener
									methods = []reflect.Value{
										// tagDataInternal.Func is the user function
										keyVal.FieldByName(fieldNameListener).MethodByName(tagListener.Func),
									}

									// Pass variable pointers
									params = []interface{}{
										// fieldVal.Interface() is the struct pointer that collects user data
										keyVal.FieldByName(fieldNameListener).Interface(),
									}

									inputCheckbox.ListenerAddReflect(tagListener.Event, params, methods, e.ref)
								}

							}
						}

						inputCheckbox.Value(value).Disabled(disabled).Checked(selected).Class("inputCheckbox").Name(tagDataInternal.Name)
						inputLabel.Text(label).Append(inputCheckbox)

						inputDivCheckbox.Append(
							factoryBrowser.NewTagSpan().Append(inputLabel),
						)

						//inputSelect.NewOption(label, value, disabled, selected)
					}
				}

			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagData.Label),
		inputDivCheckbox,
	)

	for i := 0; i != element.NumField(); i += 1 {
		fieldVal := element.Field(i)
		if fieldVal.Type() == reflect.TypeOf(Checkbox{}) {
			r := fieldVal.Interface().(Checkbox)
			r.init()
			break
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processHeaderText(element reflect.Value, defaultText string, father *html.TagDiv) {

	var ok bool
	var text string

	text, ok = element.Interface().(string)
	if ok && text != "" {

		defaultText = text
	}

	// <div class="panelHeader">
	//   <div class="headerText">Panel</div>
	//   <div class="dragIcon"></div>
	//   <div class="closeIconPanel">ˇ</div>
	// </div>
	father.Append(
		factoryBrowser.NewTagDiv().Class("panelHeader").Append(
			factoryBrowser.NewTagDiv().Class("headerText").Text(defaultText),
			factoryBrowser.NewTagDiv().Class("dragIcon"),
			factoryBrowser.NewTagDiv().Class("closeIconPanel").Text("ˇ"),
		),
	)

}
