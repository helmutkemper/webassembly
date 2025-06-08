package utils

import (
	"errors"
	"fmt"
	"sync"
)

type SequentialInterface interface {
	GetId(base string) (id string, err error)
}

// SequentialId generates sequential IDs from a base string
type SequentialId struct {
	mu      sync.Mutex
	mapData map[string]int
}

// GetId creates a sequential ID based on the base string
// @param base - The base string to create the ID from
// @returns - The generated ID, base + "_" + count
func (e *SequentialId) GetId(base string) (id string, err error) {
	if base == "" {
		err = errors.New("base is required")
		return
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	if e.mapData == nil {
		e.mapData = make(map[string]int)
	}

	if _, exists := e.mapData[base]; !exists {
		e.mapData[base] = 0
	}
	count := e.mapData[base]
	e.mapData[base]++
	id = fmt.Sprintf("%s_%d", base, count)
	return
}
