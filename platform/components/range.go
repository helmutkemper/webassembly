package components

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
)

type __rangeChane struct {
	Value float64 `wasmGet:"value"`
}

// Range
//
// English:
//
//		Component width label, input range and input number
//
//	 Public methods
//	   * Max(int64/float64)
//	   * Min(int64/float64)
//	   * Step(int64/float64)
//	   * Value(int64/float64)
//
//	 Methods for internal use, should not be used by you and should not be shadowed (rewritten)
//	   * OnChangeNumber
//	   * OnChangeRange
//
//		    Example:
//
//		    +- panel --------------------------------------------------------+
//		    |                                                                |
//		    |  +- panelCel -----------------------------------------------+  |
//		    |  |                                                          |  |
//		    |  | +- labelCel -------------------------------------------+ |  |
//		    |  | | Label                                              ˇ | |  |
//		    |  | +- compCel --------------------------------------------+ |  |
//		    |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |
//		    |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |
//		    |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |
//		    |  | +------------------------------------------------------+ |  |
//		    |  |                                                          |  |
//		    |  +----------------------------------------------------------+  |
//		    |                                                                |
//		    +----------------------------------------------------------------+
//
//	     type ColorRange struct {
//	       components.Range
//
//	       Color       float64              `wasmPanel:"type:value;min:0;max:50;step:1;default:0"` // Field with the value of the element
//	       TagRange    *html.TagInputRange  `wasmPanel:"type:inputTagRange"` // [optional] Reference to the range component
//	       TagNumber   *html.TagInputNumber `wasmPanel:"type:inputTagNumber"` // [optional] Reference to the number component
//	       ColorChange *Click               `wasmPanel:"type:listener;event:change;func:OnChange"` // [optional] Event pointer. Create one for each event, change, click, etc.
//	     }
//
//	     func (e *ColorRange) Init() { // [optional]
//	       anything you want
//
//	       you can change the default properties
//	       e.Step(1)
//	       e.Max(255)
//	       e.Min(0)
//	       e.Value(0)
//	     }
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

// OnChangeNumber Event called when the component value changes.
//
//	This function must not be called by the user, and it must not be shadowed.
//	It is public so that reflect can access it.
func (e *Range) OnChangeNumber(stt __rangeChane) {
	e.__rangeTag.Value(stt.Value)
}

// OnChangeRange Event called when the component value changes.
//
//	This function must not be called by the user, and it must not be shadowed.
//	It is public so that reflect can access it.
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

func (e *Range) value(value any) {
	e.__rangeTag.Value(value)
	e.__numberTag.Value(value)
}

// Max Sets the maximum value of the component.
//
//	When the component was created, a field within the configuration struct received the tag `wasmPanel:"type:value"`
//	and this was created as an int64 or float64 type, therefore the value passed to the function must respect the type
func (e *Range) Max(max any) {
	if e.__rangeTag == nil {
		e.__max = max
		return
	}

	e.max(max)
}

// Min Sets the minimum value of the component.
//
//	When the component was created, a field within the configuration struct received the tag `wasmPanel:"type:value"`
//	and this was created as an int64 or float64 type, therefore the value passed to the function must respect the type
func (e *Range) Min(min any) {
	if e.__rangeTag == nil {
		e.__min = min
		return
	}

	e.min(min)
}

// Step Sets the step value of the component.
//
//	When the component was created, a field within the configuration struct received the tag `wasmPanel:"type:value"`
//	and this was created as an int64 or float64 type, therefore the value passed to the function must respect the type
func (e *Range) Step(step any) {
	if e.__rangeTag == nil {
		e.__step = step
		return
	}

	e.step(step)
}

// Value Sets the value of the component.
//
//	When the component was created, a field within the configuration struct received the tag `wasmPanel:"type:value"`
//	and this was created as an int64 or float64 type, therefore the value passed to the function must respect the type
func (e *Range) Value(value any) {
	if e.__rangeTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
