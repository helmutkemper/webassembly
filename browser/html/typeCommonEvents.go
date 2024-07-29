package html

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"syscall/js"
)

type commonEvents struct {
	selfElement *js.Value
	listener    map[string]js.Func
}

func (e *commonEvents) ListenerRemove(event string) {
	if e.listener == nil {
		e.listener = make(map[string]js.Func)
	}

	e.selfElement.Call(
		"removeEventListener",
		event,
		e.listener[event],
	)

	delete(e.listener, event)
}

func (e *commonEvents) callFunc(funcObj reflect.Value, params ...interface{}) (results []interface{}, err error) {

	if funcObj.Kind() != reflect.Func {
		return nil, fmt.Errorf("funcObj is not of type reflect.Func")
	}

	if len(params) != funcObj.Type().NumIn() {
		return nil, fmt.Errorf("incorrect number of parameters")
	}

	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}

	out := funcObj.Call(in)

	results = make([]interface{}, len(out))
	for i, result := range out {
		results[i] = result.Interface()
	}
	return results, nil
}

func (e *commonEvents) isBoolean(str string) (boolean bool) {
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

func (e *commonEvents) isNumeric(str string) (numeric bool) {
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

func (e *commonEvents) ListenerAddReflect(event string, params []interface{}, functions []reflect.Value, reference any) {

	if e.listener == nil {
		e.listener = make(map[string]js.Func)
	}

	e.listener[event] = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return nil
		}

		e.selfElement.Set("isTrusted", args[0].Get("isTrusted"))

		//melhor forma de varrer um objeto
		//if e.debugListener {
		//	var obj = args[0]
		//	js.Global().Get("Object").Call("keys", obj).Call("forEach", js.FuncOf(func(this js.Value, args []js.Value) any {
		//		log.Printf("key: [%v]: %v", args[0].String(), obj.Get(args[0].String()).String())
		//		return nil
		//	}))
		//}

		for kFunc := range functions {
			element := reflect.ValueOf(params[kFunc])

			for {
				if element.Kind() != reflect.Pointer {
					break
				}

				element = element.Elem()
			}

			if !element.IsValid() {
				log.Printf("error: the event cannot populate the event variable.")
				log.Printf("This usually occurs when:")
				log.Printf("  you do not pass a pointer to the memory address where the variable is located;")
				log.Printf("  the field is not public, the first letter of the field within the struct is lowercase.")
				return nil
			}

			for i := 0; i != element.NumField(); i += 1 {
				fieldVal := element.Field(i)

				fieldTyp := reflect.TypeOf(element.Interface()).Field(i)

				propertyToGet := fieldTyp.Tag.Get("wasmGet")

				if propertyToGet == "-" || propertyToGet == "" {
					continue
				}

				if !fieldVal.CanSet() {
					log.Printf("error: the event is unable to populate the '%v' field of the event variable.", propertyToGet)
					log.Printf("This usually occurs when:")
					log.Printf("  the field is not public, the first letter of the field within the struct is lowercase.")
					continue
				}

				propertyValue := e.selfElement.Get(propertyToGet)

				var receiverType reflect.Kind
				switch propertyValue.Type() {
				case js.TypeString:
					receiverType = reflect.String
				case js.TypeNumber:
					if propertyValue.Truthy() && float64(propertyValue.Int()) == propertyValue.Float() {
						receiverType = reflect.Int64
					} else {
						receiverType = reflect.Float64
					}
				case js.TypeBoolean:
					receiverType = reflect.Bool
				default:
					receiverType = reflect.Invalid
					continue
				}

				switch fieldVal.Kind() {
				case reflect.Bool:
					if receiverType != reflect.Bool {
						if e.isBoolean(propertyValue.String()) {
							v, err := strconv.ParseBool(propertyValue.String())
							if err != nil {
								log.Printf("error: %v retornou um tipo diferente de boolean, %v", propertyToGet, propertyValue.String())
								continue
							}
							fieldVal.SetBool(v)
							continue
						}
						log.Printf("error: %v retornou um tipo diferente de boolean, %v", propertyToGet, propertyValue.String())
						continue
					}
					fieldVal.SetBool(propertyValue.Bool())
				case reflect.Int64:
					value := propertyValue.String()
					if value == "" {
						fieldVal.SetInt(0)
						continue
					}

					if e.isNumeric(value) {
						v, err := strconv.ParseInt(value, 10, 64)
						if err != nil {
							log.Printf("error: %v retornou um tipo diferente de integer, %v", propertyToGet, propertyValue.String())
							continue
						}
						fieldVal.SetInt(v)
						continue
					}
					if receiverType != reflect.Int64 {
						log.Printf("error: %v retornou um tipo diferente de integer, %v", propertyToGet, propertyValue.String())
						continue
					}
					fieldVal.SetInt(int64(propertyValue.Int()))
				case reflect.Float64:
					if receiverType != reflect.Float64 {
						value := propertyValue.String()
						if value == "" {
							fieldVal.SetFloat(0)
							continue
						}

						if e.isNumeric(value) {
							v, err := strconv.ParseFloat(value, 64)
							if err != nil {
								log.Printf("error: %v (%v) retornou um tipo diferente de float64, %v", propertyToGet, propertyValue.Type(), propertyValue.String())
								continue
							}
							fieldVal.SetFloat(v)
							continue
						}
						log.Printf("error: %v (%v) retornou um tipo diferente de float64, %v", propertyToGet, propertyValue.Type(), propertyValue.String())
						continue
					}
					fieldVal.SetFloat(propertyValue.Float())
				case reflect.String:
					if receiverType != reflect.String {
						log.Printf("error: %v retornou um tipo diferente de string, %v", propertyToGet, propertyValue.String())
						continue
					}
					fieldVal.SetString(propertyValue.String())
				default:
					log.Printf("error: %v deve ser string, int64, float64 ou bool", propertyToGet)
					continue
				}
			}

			if !functions[kFunc].IsValid() {
				log.Printf("error: method passado para ListenerAdd é inválido")
				return nil
			}

			if reference == nil {
				_, err := e.callFunc(functions[kFunc], element.Interface())
				if err != nil {
					log.Printf("error: chamar a função %v, retornou um erro: %v", functions[kFunc].Type().Name(), err)
					return nil
				}
			} else {
				_, err := e.callFunc(functions[kFunc], element.Interface(), reference)
				if err != nil {
					log.Printf("error: chamar a função %v, retornou um erro: %v", functions[kFunc].Type().Name(), err)
					return nil
				}
			}
		}

		return nil
	})

	e.selfElement.Call(
		"addEventListener",
		event,
		e.listener[event],
	)
}
