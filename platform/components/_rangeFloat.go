package components

import "reflect"

type RangeFloat struct {
	element     reflect.Value
	label       string
	min         float64
	max         float64
	value       float64
	step        float64
	changeValue func(value, minimum, maximum float64)
	changeMin   func(value, minimum, maximum float64)
	changeMax   func(value, minimum, maximum float64)
	changeStep  func(value, minimum, maximum float64)
}

func (e *RangeFloat) SetLabel(label string) {
	e.label = label
}

func (e *RangeFloat) GetLabel() (label string) {
	return e.label
}

func (e *RangeFloat) SetMin(min float64) {
	e.min = min
}

func (e *RangeFloat) GetMin() (min float64) {
	return e.min
}

func (e *RangeFloat) SetMax(max float64) {
	e.max = max
}

func (e *RangeFloat) GetMax() (max float64) {
	return e.max
}

func (e *RangeFloat) SetValue(value float64) {
	e.value = value
}

func (e *RangeFloat) GetValue() (value float64) {
	return e.value
}

func (e *RangeFloat) SetStep(step float64) {
	e.step = step
}

func (e *RangeFloat) GetStep() (step float64) {
	return e.step
}

func (e *RangeFloat) SetChangeValue(function func(value, minimum, maximum float64)) {
	e.changeValue = function
}

func (e *RangeFloat) SetChangeMin(function func(value, minimum, maximum float64)) {
	e.changeMin = function
}

func (e *RangeFloat) SetChangeMax(function func(value, minimum, maximum float64)) {
	e.changeMax = function
}

func (e *RangeFloat) SetChangeStep(function func(value, minimum, maximum float64)) {
	e.changeStep = function
}
