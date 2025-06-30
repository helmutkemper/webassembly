package rulesSequentialId

import (
	"fmt"
	"regexp"
	"sync"
)

var seqId sequentialId
var re *regexp.Regexp

func init() {
	// Regular expression is slow, so it should be initialized once, whenever possible
	re = regexp.MustCompile("^(.*?)(_[0-9]+)$")
}

// GetIdFromBase Returns a sequential ID according to the base.
//
//	As a rule, the ID will be generated using the following rule: base + _ + sequential number
//
//	This makes all devices of the same type have IDs initialized with the device type
func GetIdFromBase(base string) (id string) {
	return seqId.getId(base)
}

// sequentialId generates sequential IDs from a base string
type sequentialId struct {
	mu      sync.Mutex
	mapData map[string]int
}

// GetId creates a sequential ID based on the base string
// @param base - The base string to create the ID from
// @returns - The generated ID, base + "_" + count
func (e *sequentialId) getId(base string) (id string) {
	if base == "" {
		base = "empty_base"
	}

	base = re.ReplaceAllString(base, "$1")

	e.mu.Lock()
	defer e.mu.Unlock()

	if e.mapData == nil {
		e.mapData = make(map[string]int)
	}

	if _, exists := e.mapData[base]; !exists {
		e.mapData[base] = 0
	}
	count := e.mapData[base]
	id = fmt.Sprintf("%s_%d", base, count)
	e.mapData[base]++
	return
}
