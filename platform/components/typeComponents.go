package components

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type Components struct {
	panelFather *html.TagDiv
	panelBody   *html.TagDiv
}

func (e *Components) Init(el any) (err error) {
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

	stage := factoryBrowser.NewStage()
	stage.Append(e.panelFather)

	return
}

func (e *Components) GetUId() (uuidStr string, err error) {
	uId, err := uuid.NewUUID()
	if err != nil {
		//file, line, funcName := runTimeUtil.Trace()
		//err = errors.Join(fmt.Errorf("%v(line: %v).NewUUID().error: %v", funcName, line, err))
		//err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}
	uuidStr = uId.String()
	return
}

func (e *Components) createDivsFather() {
	e.panelFather = factoryBrowser.NewTagDiv().Class("panel")
	e.panelBody = factoryBrowser.NewTagDiv().Class("panelBody")
}

func (e *Components) process(element reflect.Value, typeof reflect.Type) (err error) {
	if element.Kind() == reflect.Pointer {
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
					e.processHeaderText(fieldVal, e.panelFather)
					// Espera criar panelHeader para que panelBody fique abaixo
					e.panelFather.Append(e.panelBody)
				case "panelBody":
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
						err = fmt.Errorf("component.Range (%v) requires a public fiel, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
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
						err = fmt.Errorf("component.Button (%v) requires a public fiel, '%v' with the first letter capitalized", fieldTyp.Name, strings.ToUpper(fieldTyp.Name[:1])+fieldTyp.Name[1:])
						return
					}

					err = e.processComponentButton(fieldVal, tagData, divComponent)
					if err != nil {
						return
					}
				case "select":
					e.processComponentSelect(fieldVal, tagData, divComponent)
				case "radio":
					divComponent.Class("component component-radio")
					e.processComponentRadio(fieldVal, tagData, divComponent)
				case "checkbox":
					divComponent.Class("component component-checkbox")
					e.processComponentCheckbox(fieldVal, tagData, divComponent)
				case "color":
					e.processComponentColor(fieldVal, tagData, divComponent)
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

func (e *Components) verifyTypeNumericFromElement(fieldVal reflect.Value, valueType reflect.Kind) (dataType reflect.Kind, value any, ok bool) {
	switch valueType {
	case reflect.Int64:
		dataType = reflect.Int64
		value = fieldVal.Int()
	case reflect.Float64:
		dataType = reflect.Float64
		value = fieldVal.Float()
	default:
		dataType = reflect.Invalid
		return
	}

	ok = true
	return
}

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
		inputRange.ListenerAddReflect("input", params, methods)
		inputNumber.ListenerAddReflect("input", params, methods)
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

				switch tagDataInternal.Event {
				// If the user wants to use the `input` event, the code assembles the user event and the panel event
				case "input":

					// Passes the functions to be executed in the listener
					methods = []reflect.Value{
						// tagDataInternal.Func is the user function
						fieldVal.MethodByName(tagDataInternal.Func),
						// rangeComponent is the struct components.Range and OnChangeNumber is a function belonging to the struct components.Range
						reflect.ValueOf(&rangeComponent).MethodByName("OnChangeNumber"),
						// rangeComponent is the struct components.Range and OnChangeRange is a function belonging to the struct components.Range
						reflect.ValueOf(&rangeComponent).MethodByName("OnChangeRange"),
					}

					// Pass variable pointers
					params = []interface{}{
						// fieldVal.Interface() is the struct pointer that collects user data
						fieldVal.Interface(),
						// __rangeOnInputEvent is the type pointer contained in components.Range and collects value
						new(__rangeOnInputEvent),
						// __rangeOnInputEvent is the type pointer contained in components.Range and collects value
						new(__rangeOnInputEvent),
					}

				// If the user uses another event, different from `input`, it just mounts the user event
				default:

					// Passes the functions to be executed in the listener
					methods = []reflect.Value{
						fieldVal.MethodByName(tagDataInternal.Func),
					}

					// Pass variable pointers
					params = []interface{}{
						fieldVal.Interface(),
					}

				}

				// explanation
				//   inputNumber.ListenerAdd() accepts two arrays, one for the function to be invoked, and the other with the data to be passed
				//   The first element of the array is the user function
				//   From the second element onwards, they are internal functions and must be called after the user function in case the user has changed any value.
				inputRange.ListenerAddReflect(tagDataInternal.Event, params, methods)
				inputNumber.ListenerAddReflect(tagDataInternal.Event, params, methods)

				//case "func":
				//	if !fieldVal.CanInterface() {
				//		err = fmt.Errorf("%v.%v must be a public field", element.Type().Name(), fieldTyp.Name)
				//		return
				//	}
				//
				//	if fieldVal.Kind() != reflect.Func {
				//		err = fmt.Errorf("%v.%v must be a function, but a %v was received", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
				//		return
				//	}
				//
				//	if _, ok = fieldVal.Interface().(func(any)); !ok {
				//		err = fmt.Errorf("%v.%v must be a function type func pointer 'func(ars any)', but a %v was received", element.Type().Name(), fieldTyp.Name, fieldVal.Kind())
				//		return
				//	}
				//
				//	if _, _, found := e.searchFieldByTagType("data", tagDataInternal.Event, element); !found {
				//		err = fmt.Errorf("in the struct named %v, there is a tag named %v 'type:func' for the event '%v', but there is no field with the tag 'type:data;event:%v' containing the pointer to the struct with the data information to be collected for the `%v` event", element.Type().Name(), fieldTyp.Name, tagDataInternal.Event, tagDataInternal.Event, tagDataInternal.Event)
				//		return err
				//	}
				//
				//	if dataVal, dataType, found := e.searchFieldByTagType("data", tagDataInternal.Event, element); found {
				//		if !fieldVal.IsNil() && dataVal.IsNil() {
				//			err = fmt.Errorf("%v.%v is nil, that is, it did not receive a pointer to a variable and cannot be set", element.Type().Name(), dataType.Name)
				//			return
				//		}
				//	}
				//
				//	if dataVal, dataType, found := e.searchFieldByTagType("data", tagDataInternal.Event, element); found {
				//		if !fieldVal.IsNil() && !dataVal.CanSet() {
				//			err = fmt.Errorf("%v.%v can't be set", element.Type().Name(), dataType.Name)
				//			return
				//		}
				//	}

				//case "data":
				//	if !fieldVal.CanInterface() {
				//		err = fmt.Errorf("%v.%v must be a public field", element.Type().Name(), fieldTyp.Name)
				//		return
				//	}
				//
				//	if fieldVal.Kind() != reflect.Pointer {
				//		err = fmt.Errorf("%v.%v must be a struct pointer, but a %v was received: %+v", element.Type().Name(), fieldTyp.Name, fieldVal.Kind(), fieldVal.Interface())
				//		return
				//	}
				//
				//	if _, _, found := e.searchFieldByTagType("func", tagDataInternal.Event, element); !found {
				//		err = fmt.Errorf("in the struct named %v, there is a tag named %v 'type:data' for the event '%v', but there is no field with the tag 'type:func;event:%v' containing the function pointer 'func(arg any)' of event `%v`", element.Type().Name(), fieldTyp.Name, tagDataInternal.Event, tagDataInternal.Event, tagDataInternal.Event)
				//		return err
				//	}

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

func (e *Components) processComponentButton(element reflect.Value, tagData *tag, father *html.TagDiv) (err error) {

	var dataType reflect.Kind
	var value any
	var ok bool

	buttonComponent := Button{}

	inputButton := factoryBrowser.NewTagInputButton().Class("inputButton").Value("-- ok --")

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
		err = errors.Join(err, fmt.Errorf("       func (e *OnClickEvent) OnClick(event OnClickEvent) {"))
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
				if !passValue && tagDataInternal.Default != "" {
					inputButton.Value(tagDataInternal.Default)
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
				inputButton.ListenerAddReflect(tagDataInternal.Event, params, methods)

			}

		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagData.Label),
		inputButton,
	)

	return
}

func (e *Components) processComponentColor(element reflect.Value, tagData *tag, father *html.TagDiv) {

	var buttonClick func()
	var buttonPress func()
	var buttonRelease func()
	var label string
	var value string
	var disabled bool
	_ = buttonClick   //todo: fazer
	_ = buttonPress   //todo: fazer
	_ = buttonRelease //todo: fazer
	var ok bool

	if element.CanInterface() {
		for i := 0; i != element.NumField(); i += 1 {
			fieldVal := element.Field(i)
			fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

			tagDataInternal := new(tag)
			tagRaw := fieldTyp.Tag.Get("wasmPanel")
			if tagRaw != "" {
				tagDataInternal.init(tagRaw)

				switch tagDataInternal.Type {
				case "disabled":
					disabled, ok = fieldVal.Interface().(bool)
					if !ok {
						log.Printf("error: Disabled deve ser do tipo booleano")
					}
				case "label":
					label, ok = fieldVal.Interface().(string)
					if !ok {
						log.Printf("error: Label deve ser do tipo string")
					}
				case "value":
					value, ok = fieldVal.Interface().(string)
					if !ok {
						log.Printf("error: Value deve ser do tipo string")
					}
				case "onclick":
					buttonClick, ok = fieldVal.Interface().(func())
					if !ok {
						log.Printf("error: A função onclick deve ser func()")
					}
				case "onpress":
					buttonPress, ok = fieldVal.Interface().(func())
					if !ok {
						log.Printf("error: A função onpress deve ser func()")
					}
				case "onrelease":
					buttonRelease, ok = fieldVal.Interface().(func())
					if !ok {
						log.Printf("error: A função onrelease deve ser func()")
					}
				}
			}
		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(label),
		factoryBrowser.NewTagInputColor().Class("inputButton").Value(value).Disabled(disabled),
	)
}

func (e *Components) processComponentSelect(element reflect.Value, tagData *tag, father *html.TagDiv) {

	selectTag := factoryBrowser.NewTagSelect().Class("inputSelect")

	if element.CanInterface() {
		for i := 0; i != element.Len(); i += 1 {
			subElement := element.Index(i)

			var ok bool
			var disabled bool
			var selected bool
			var label string
			var value string

			for i := 0; i != subElement.NumField(); i += 1 {

				fieldVal := subElement.Field(i)
				fieldTyp := reflect.TypeOf(subElement.Interface()).Field(i)

				tagDataInternal := new(tag)
				tagRaw := fieldTyp.Tag.Get("wasmPanel")
				if tagRaw != "" {
					tagDataInternal.init(tagRaw)

					switch tagDataInternal.Type {
					case "disabled":
						disabled, ok = fieldVal.Interface().(bool)
						if !ok {
							log.Printf("error: disabled deve ser do tipo booleano")
						}
					case "selected":
						selected, ok = fieldVal.Interface().(bool)
						if !ok {
							log.Printf("error: selected deve ser do tipo booleano")
						}
					case "label":
						label, ok = fieldVal.Interface().(string)
						if !ok {
							log.Printf("error: label deve ser do tipo string")
						}
					case "value":
						value, ok = fieldVal.Interface().(string)
						if !ok {
							log.Printf("error: value deve ser do tipo string")
						}
					}
				}
			}

			selectTag.NewOption(label, value, disabled, selected)

		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagData.Label),
		selectTag,
	)
}

func (e *Components) processComponentRadio(element reflect.Value, tagData *tag, father *html.TagDiv) {

	radioOptionsTag := factoryBrowser.NewTagDiv().Class("radioOptions")

	if element.CanInterface() {
		for i := 0; i != element.Len(); i += 1 {
			subElement := element.Index(i)

			var ok bool
			var name string
			var disabled bool
			var selected bool
			var label string
			var value string

			for i := 0; i != subElement.NumField(); i += 1 {

				fieldVal := subElement.Field(i)
				fieldTyp := reflect.TypeOf(subElement.Interface()).Field(i)

				tagDataInternal := new(tag)
				tagRaw := fieldTyp.Tag.Get("wasmPanel")
				if tagRaw != "" {
					tagDataInternal.init(tagRaw)

					switch tagDataInternal.Type {
					case "name":
						name, ok = fieldVal.Interface().(string)
						if !ok {
							log.Printf("error: name deve ser do tipo string")
						}
					case "disabled":
						disabled, ok = fieldVal.Interface().(bool)
						if !ok {
							log.Printf("error: disabled deve ser do tipo booleano")
						}
					case "selected":
						selected, ok = fieldVal.Interface().(bool)
						if !ok {
							log.Printf("error: selected deve ser do tipo booleano")
						}
					case "label":
						label, ok = fieldVal.Interface().(string)
						if !ok {
							log.Printf("error: label deve ser do tipo string")
						}
					case "value":
						value, ok = fieldVal.Interface().(string)
						if !ok {
							log.Printf("error: value deve ser do tipo string")
						}
					}
				}
			}

			radioOptionsTag.Append(
				factoryBrowser.NewTagLabel().Text(label).Append(
					factoryBrowser.NewTagInputRadio().Class("inputRadio").Name(name).Value(value).Disabled(disabled).Checked(selected),
				),
			)

		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagData.Label),
		radioOptionsTag,
	)
}

func (e *Components) processComponentCheckbox(element reflect.Value, tagData *tag, father *html.TagDiv) {

	radioOptionsTag := factoryBrowser.NewTagDiv().Class("checkboxOptions")

	if element.CanInterface() {
		for i := 0; i != element.Len(); i += 1 {
			subElement := element.Index(i)

			var ok bool
			var name string
			var disabled bool
			var selected bool
			var label string
			var value string

			for i := 0; i != subElement.NumField(); i += 1 {

				fieldVal := subElement.Field(i)
				fieldTyp := reflect.TypeOf(subElement.Interface()).Field(i)

				tagDataInternal := new(tag)
				tagRaw := fieldTyp.Tag.Get("wasmPanel")
				if tagRaw != "" {
					tagDataInternal.init(tagRaw)

					switch tagDataInternal.Type {
					case "name":
						name, ok = fieldVal.Interface().(string)
						if !ok {
							log.Printf("error: name deve ser do tipo string")
						}
					case "disabled":
						disabled, ok = fieldVal.Interface().(bool)
						if !ok {
							log.Printf("error: disabled deve ser do tipo booleano")
						}
					case "selected":
						selected, ok = fieldVal.Interface().(bool)
						if !ok {
							log.Printf("error: selected deve ser do tipo booleano")
						}
					case "label":
						label, ok = fieldVal.Interface().(string)
						if !ok {
							log.Printf("error: label deve ser do tipo string")
						}
					case "value":
						value, ok = fieldVal.Interface().(string)
						if !ok {
							log.Printf("error: value deve ser do tipo string")
						}
					}
				}
			}

			radioOptionsTag.Append(
				factoryBrowser.NewTagLabel().Text(label).Append(
					factoryBrowser.NewTagInputCheckBox().Class("inputRadio").Name(name).Value(value).Disabled(disabled).Checked(selected),
				),
			)

		}
	}

	father.Append(
		factoryBrowser.NewTagSpan().Text(tagData.Label),
		radioOptionsTag,
	)
}

func (e *Components) processHeaderText(element reflect.Value, father *html.TagDiv) {

	// <div class="panelHeader">
	//   <div class="headerText">Panel</div>
	//   <div class="dragIcon"></div>
	//   <div class="closeIconPanel">ˇ</div>
	// </div>
	father.Append(
		factoryBrowser.NewTagDiv().Class("panelHeader").Append(
			factoryBrowser.NewTagDiv().Class("headerText").Text(element.Interface()),
			factoryBrowser.NewTagDiv().Class("dragIcon"),
			factoryBrowser.NewTagDiv().Class("closeIconPanel").Text("ˇ"),
		),
	)

}
