package components

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/runTimeUtil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

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
	//log.Printf("tag data: %v", data)
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

type Components struct {
	panelFather *html.TagDiv
	panelBody   *html.TagDiv
}

type __rangeChane struct {
	Value float64 `wasmGet:"value"`
}

type Range struct {
	__max   any
	__min   any
	__step  any
	__value any

	// __rangeChane is the pointer sent when the `change` event happens
	__change *__rangeChane

	// reference of component elements
	__rangeTag  *html.TagInputRange  `wasmPanel:"type:inputTagRange"`
	__numberTag *html.TagInputNumber `wasmPanel:"type:inputTagNumber"`
}

func (e *Range) OnChangeNumber(stt __rangeChane) {
	e.__rangeTag.Value(stt.Value)
}

func (e *Range) OnChangeRange(stt __rangeChane) {
	e.__numberTag.Value(stt.Value)
}

func (e *Range) init() {
	if e.__max != nil {
		e.max(e.__max)
		e.__max = nil
	}

	if e.__min != nil {
		e.min(e.__min)
		e.__min = nil
	}

	if e.__step != nil {
		e.step(e.__step)
		e.__step = nil
	}

	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Range) max(max any) {
	e.__rangeTag.Max(max)
	e.__numberTag.Max(max)
}

func (e *Range) min(min any) {
	e.__rangeTag.Min(min)
	e.__numberTag.Min(min)
}

func (e *Range) step(step any) {
	e.__rangeTag.Step(step)
	e.__numberTag.Step(step)
}

func (e *Range) value(step any) {
	e.__rangeTag.Value(step)
	e.__numberTag.Value(step)
}

func (e *Range) Max(max any) {
	if e.__rangeTag == nil {
		e.__max = max
		return
	}

	e.max(max)
}

func (e *Range) Min(min any) {
	if e.__rangeTag == nil {
		e.__min = min
		return
	}

	e.min(min)
}

func (e *Range) Step(step any) {
	if e.__rangeTag == nil {
		e.__step = step
		return
	}

	e.step(step)
}

func (e *Range) GetValue() (value any) {
	return e.__rangeTag.GetValue()
}

func (e *Range) Value(value any) {
	if e.__rangeTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}

func (e *Components) Init(el any) (err error) {
	element := reflect.ValueOf(el)
	typeof := reflect.TypeOf(el)
	e.createDivsFather()
	err = e.process(element, typeof)
	if err != nil {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).process().error: %v", funcName, line, err))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
		return
	}

	stage := factoryBrowser.NewStage()
	stage.Append(e.panelFather)

	return
}

func (e *Components) GetUId() (uuidStr string, err error) {
	uId, err := uuid.NewUUID()
	if err != nil {
		file, line, funcName := runTimeUtil.Trace()
		err = errors.Join(fmt.Errorf("%v(line: %v).NewUUID().error: %v", funcName, line, err))
		err = errors.Join(fmt.Errorf("file: %v", file), err)
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
						file, line, funcName := runTimeUtil.Trace()
						err = errors.Join(fmt.Errorf("%v(line: %v).process().error: %v", funcName, line, err))
						err = errors.Join(fmt.Errorf("file: %v", file), err)
						return
					}
				case "compCel":
					// ignore
				case "component":
					divCompCel := factoryBrowser.NewTagDiv().Class("compCel")

					err = e.processComponent(fieldVal, fieldTyp.Type, tagData, divCompCel)
					if err != nil {
						file, line, funcName := runTimeUtil.Trace()
						err = errors.Join(fmt.Errorf("%v(line: %v).processComponent().error: %v", funcName, line, err))
						err = errors.Join(fmt.Errorf("file: %v", file), err)
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
						file, line, funcName := runTimeUtil.Trace()
						err = errors.Join(fmt.Errorf("%v(line: %v).process().error: %v", funcName, line, err))
						err = errors.Join(fmt.Errorf("file: %v", file), err)
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

					if fieldVal.CanSet() && fieldVal.IsNil() {
						newInstance := reflect.New(fieldVal.Type().Elem())
						fieldVal.Set(newInstance)
					}

					err = e.processComponentRange(fieldVal, tagData, divComponent)
					if err != nil {
						file, line, funcName := runTimeUtil.Trace()
						err = errors.Join(fmt.Errorf("%v(line: %v).processComponentRange().error: %v", funcName, line, err))
						err = errors.Join(fmt.Errorf("file: %v", file), err)
						return
					}
				case "button":
					e.processComponentButton(fieldVal, tagData, divComponent)
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

	if element.Kind() != reflect.Pointer {
		err = fmt.Errorf("error: component `%v` must be a pointer", 9 /*element.Type().Name()*/)
		return
	}

	if !element.CanInterface() {
		err = fmt.Errorf("error: component `%v` cannot be interfacied", element.Elem().Type().Name())
		return
	}

	if !element.CanSet() {
		err = fmt.Errorf("error: component `%v` cannot be set", element.Elem().Type().Name())
		return
	}

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
		err = errors.Join(err, fmt.Errorf("       type MyComponent struct {"))
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

		// __rangeChane is the pointer sent when the `change` event happens
		rangeComponent.__change = new(__rangeChane)
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
			// __rangeChane is the type pointer contained in components.Range and collects value
			new(__rangeChane),
			// __rangeChane is the type pointer contained in components.Range and collects value
			new(__rangeChane),
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
				fieldVal.Interface().(*html.TagInputRange).Min(-5).Max(10).Value(-5)

			// Checks whether the reference to the input number tag was requested by the user
			case "inputTagNumber":
				fieldVal.Set(reflect.ValueOf(inputNumber))
				fieldVal.Interface().(*html.TagInputNumber).Min(-5).Max(10).Value(-5)

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
						// __rangeChane is the type pointer contained in components.Range and collects value
						new(__rangeChane),
						// __rangeChane is the type pointer contained in components.Range and collects value
						new(__rangeChane),
					}

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
		}
	}

	method := elementOriginal.MethodByName("Init")
	if method.IsValid() {
		method.Call(nil)
	}

	return
}

func (e *Components) processComponentButton(element reflect.Value, tagData *tag, father *html.TagDiv) {

	var buttonClick func()
	var buttonPress func()
	var buttonRelease func()
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
		factoryBrowser.NewTagSpan().Text(tagData.Label),
		factoryBrowser.NewTagButton().Class("inputButton").Value(tagData.Value),
	)
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

	log.Printf("tag data: %+v", tagData)

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
	father.Append(factoryBrowser.NewTagDiv().Class("panelHeader").Append(
		factoryBrowser.NewTagDiv().Class("headerText").Text(element.Interface()),
		factoryBrowser.NewTagDiv().Class("dragIcon"),
		factoryBrowser.NewTagDiv().Class("closeIconPanel").Text("ˇ"),
	),
	)

}

type ComponentInterface interface {
	canComponentize() (err error)
}

type component struct {
}

func (e *component) GetUId() (uuidStr string) {
	uId, err := uuid.NewUUID()
	if err != nil {
		err = fmt.Errorf("controlCell.NewUUID().error: %v", err)
		log.Printf("%v", err) //todo: melhorar
		return
	}
	uuidStr = uId.String()
	return
}

// <div class="labelCel">
//
//	<div class="labelText">Label</div>
//	<div class="closeIcon">ˇ</div>
//
// </div>
func (e *component) processLabelCel(cel LabelCel, div *html.TagDiv) {
	id := e.GetUId()
	div.Append(
		factoryBrowser.NewTagDiv().Id(id+"-labelCel").Class("labelCel").Append(
			factoryBrowser.NewTagDiv().Id(id+"-labelText").Class("labelText").Text(cel.Label),
			factoryBrowser.NewTagDiv().Id(id+"-closeIcon").Class("closeIcon").Text("ˇ"),
		),
	)
	return
}

// <div class="component">
//
//	<span>Text inside component</span>
//	<input type="range" class="inputRange">
//	<input type="number" class="inputNumber">
//
// </div>
func (e *component) processInputRange(cel RangeInt, div *html.TagDiv) {
	id := e.GetUId()
	div.Append(
		factoryBrowser.NewTagDiv().Id(id+"-component").Class("component").Append(
			factoryBrowser.NewTagDiv().Text(cel.Text),
			factoryBrowser.NewTagInputRange().Class("inputRange").Min(cel.Min).Max(cel.Max).Step(cel.Step).Value(cel.Value),
			factoryBrowser.NewTagInputNumber().Class("inputNumber").Min(cel.Min).Max(cel.Max).Step(cel.Step).Value(cel.Value),
		),
	)
}

func (e *component) canComponentize() (err error) {
	return nil
}

func (e *component) isBoolean(str string) (boolean bool) {
	str = strings.ToLower(str)
	if str == "true" {
		return true
	}
	if str == "false" {
		return true
	}
	if str == "t" {
		return true
	}
	if str == "f" {
		return true
	}
	if str == "1" {
		return true
	}
	if str == "0" {
		return true
	}

	return false
}

func (e *component) isNumeric(str string) (numeric bool) {
	for k := range str {
		switch str[k] {
		case '.':
			continue
		case '-':
			continue
		case '0':
			continue
		case '1':
			continue
		case '2':
			continue
		case '3':
			continue
		case '4':
			continue
		case '5':
			continue
		case '6':
			continue
		case '7':
			continue
		case '8':
			continue
		case '9':
			continue
		default:
			return false
		}
	}

	return true
}

type RangeIntFunc func(value, minimum, maximum int64)

type PanelCel struct {
	LabelCel     LabelCel
	ComponentCel ComponentCel
}

type Label struct {
	component

	Label string
}

type LabelCel struct {
	component

	Label string
}

type Content struct {
	Label   string
	Content []any
}

type ComponentCel struct {
	component

	Content []Content
}

type RangeInt struct {
	component

	Text        string
	Min         int64
	Max         int64
	Value       int64
	Step        int64
	ChangeValue RangeIntFunc
}

func (e *RangeInt) canComponentize() (err error) {
	return nil
}

func (e *RangeInt) init() (err error) {
	return nil
}

type Panel struct {
	component

	Content []any
}

func (e *Panel) Init() (panel *html.TagDiv, err error) {
	// <div class="panel">
	//    <div class="panelHeader">
	//        <div class="headerText">Panel</div>
	//        <div class="dragIcon"></div>
	//        <div class="closeIconPanel">ˇ</div>
	//    </div>
	//    <div class="panelBody"></div>
	// </div>

	panelHeader := factoryBrowser.NewTagDiv().Class("panelHeader").Append(
		// <div class="headerText">Panel</div>
		factoryBrowser.NewTagDiv().Class("headerText").Text("Panel"), //todo: texto
		// <div class="dragIcon"></div>
		factoryBrowser.NewTagDiv().Class("dragIcon"),
		// <div class="closeIconPanel">ˇ</div>
		factoryBrowser.NewTagDiv().Class("closeIconPanel").Text("ˇ"),
	)

	// <div class="panelBody"></div>
	panelBody := factoryBrowser.NewTagDiv().Class("panelBody")

	// <div class="panel">
	panel = factoryBrowser.NewTagDiv().Class("panel").Append(panelHeader, panelBody)

	//for k := range e.Content {
	//	switch el := e.Content[k].(type) {
	//	case PanelCel:
	//		e.processLabelCel(el.LabelCel, panel)
	//		for k := range el.ComponentCel.Content {
	//			var eLabel any
	//			switch el := el.ComponentCel.Content[k].(type) {
	//			case LabelCel:
	//				eLabel = el.Label
	//				_ = eLabel
	//				//e.processLabel(el, panel)
	//			case RangeInt:
	//				e.processInputRange(el, panelBody)
	//			//case PanelCel:
	//			//	e.processLabelCel(el.LabelCel, panel)
	//			default:
	//				log.Printf(">> >> %+v", reflect.ValueOf(el).Kind())
	//			}
	//		}
	//
	//	default:
	//		log.Printf(">> %+v", reflect.TypeOf(el).Name())
	//	}

	//reflectElm := reflect.ValueOf(e.Content[k])
	//reflectTyp := reflect.TypeOf(e.Content[k])
	//
	//if reflectElm.Kind() == reflect.Pointer {
	//	reflectElm = reflectElm.Elem()
	//	reflectTyp = reflectTyp.Elem()
	//}
	//
	//for i := 0; i < reflectTyp.NumField(); i++ {
	//	field := reflectTyp.Field(i)
	//	fieldValue := reflectElm.Field(i)
	//
	//	fieldName := field.Name
	//	fieldType := field.Type
	//
	//	if fieldValue.CanInterface() {
	//		switch fieldType.String() {
	//		case "components.LabelCel":
	//			e.processLabelCel(fieldValue.Interface().(LabelCel), panelBody)
	//		case "components.ComponentCel":
	//			for i := 0; i < fieldValue.Len(); i++ {
	//				fieldValue = fieldValue.Index(i)
	//				log.Printf("%+v", fieldValue.Interface())
	//
	//			}
	//
	//		default:
	//			log.Printf("%v, %v", fieldType.String(), fieldName)
	//		}
	//	}
	//}

	//}
	//
	//stage := factoryBrowser.NewStage()
	//stage.Append(panel)

	return
}
