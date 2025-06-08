package main

import (
	"errors"
)

type Device interface {
	GetID() string
	GetZIndex() int
	SetZIndex(int)
}

type ManagerZIndex struct {
	Map map[string]Device
}

func NewManagerZIndex() *ManagerZIndex {
	return &ManagerZIndex{Map: make(map[string]Device)}
}

func (e *ManagerZIndex) Add(device Device) error {
	if device == nil {
		return errors.New("device is nil")
	}

	e.Map[device.GetID()] = device
	return nil
}
