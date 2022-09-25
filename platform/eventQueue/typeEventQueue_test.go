package eventQueue

import (
	"testing"
	"time"
)

func TestKeyboardEventQueue_GetActivesEvent(t *testing.T) {
	var queue EventQueue
	queue.Init()

	queue.AddEvent("ArrowRight", true)
	queue.AddEvent("ArrowLeft", true)
	queue.AddEvent("ArrowLeft", false)

	empty, list := queue.getActivesEvent()
	if empty != false || len(list) != 1 {
		t.Fatal("active list inconsistency, test 1")
	}

	if list[0].label != "ArrowRight" || list[0].active != true {
		t.Fatal("active list inconsistency, test 2")
	}

	queue.Clear()

	queue.AddEvent("ArrowRight", true)
	queue.AddEvent("ArrowLeft", true)
	queue.AddEvent("ArrowLeft", false)
	queue.AddEvent("ArrowRight", false)
	queue.AddEvent("ArrowLeft", true)
	queue.AddEvent("ArrowRight", true)
	empty, list = queue.getActivesEvent()
	if empty != false || len(list) != 2 {
		t.Fatal("active list inconsistency, test 3")
	}

	if list[0].label != "ArrowRight" {
		t.Fatal("active list key[0] inconsistency, test 4")
	}
}

func TestKeyboardEventQueue_AddSpecialEvent(t *testing.T) {
	var queue EventQueue
	queue.Init()

	var specialKeyFire1 = make([]Event, 4)
	specialKeyFire1[0].Set("ArrowRight", true, 500*time.Millisecond)
	specialKeyFire1[1].Set("ArrowDown", true, 500*time.Millisecond)
	specialKeyFire1[2].Set("ArrowLeft", true, 500*time.Millisecond)
	specialKeyFire1[3].Set("ArrowUp", true, 500*time.Millisecond)

	var specialKeyFire2 = make([]Event, 2)
	specialKeyFire2[0].Set("ArrowRight", true, 500*time.Millisecond)
	specialKeyFire2[1].Set("ArrowLeft", true, 500*time.Millisecond)

	var specialKeyFire3 = make([]Event, 8)
	specialKeyFire3[0].Set("ArrowRight", true, 500*time.Millisecond)
	specialKeyFire3[1].Set("ArrowRight", false, 500*time.Millisecond)
	specialKeyFire3[2].Set("ArrowDown", true, 500*time.Millisecond)
	specialKeyFire3[3].Set("ArrowDown", false, 500*time.Millisecond)
	specialKeyFire3[4].Set("ArrowLeft", true, 500*time.Millisecond)
	specialKeyFire3[5].Set("ArrowLeft", false, 500*time.Millisecond)
	specialKeyFire3[6].Set("ArrowUp", true, 500*time.Millisecond)
	specialKeyFire3[7].Set("ArrowUp", false, 500*time.Millisecond)

	var specialKeyFire4 = make([]Event, 6)
	specialKeyFire4[0].Set("ArrowDown", true, 500*time.Millisecond)
	specialKeyFire4[1].Set("ArrowDown", false, 500*time.Millisecond)
	specialKeyFire4[2].Set("ArrowLeft", true, 500*time.Millisecond)
	specialKeyFire4[3].Set("ArrowLeft", false, 500*time.Millisecond)
	specialKeyFire4[4].Set("ArrowUp", true, 500*time.Millisecond)
	specialKeyFire4[5].Set("ArrowUp", false, 500*time.Millisecond)

	var specialKeyFire5 = make([]Event, 4)
	specialKeyFire5[0].Set("ArrowDown", true, 500*time.Millisecond)
	specialKeyFire5[1].Set("ArrowDown", false, 500*time.Millisecond)
	specialKeyFire5[2].Set("ArrowRight", true, 500*time.Millisecond)
	specialKeyFire5[3].Set("ArrowRight", false, 500*time.Millisecond)

	var specialKeyFire6 = make([]Event, 2)
	specialKeyFire6[0].Set("ArrowLeft", true, 500*time.Millisecond)
	specialKeyFire6[1].Set("ArrowUp", true, 500*time.Millisecond)

	queue.AddSpecialEvent("specialKeyFire1", specialKeyFire1)
	queue.AddSpecialEvent("specialKeyFire2", specialKeyFire2)
	queue.AddSpecialEvent("specialKeyFire3", specialKeyFire3)
	queue.AddSpecialEvent("specialKeyFire4", specialKeyFire4)
	queue.AddSpecialEvent("specialKeyFire5", specialKeyFire5)
	queue.AddSpecialEvent("specialKeyFire6", specialKeyFire6)

	namesToTest := []string{"specialKeyFire3", "specialKeyFire4", "specialKeyFire1", "specialKeyFire5", "specialKeyFire2", "specialKeyFire6"}
	eventsToTest := [][]Event{
		{{label: "ArrowRight", active: true}, {label: "ArrowRight", active: false}, {label: "ArrowDown", active: true}, {label: "ArrowDown", active: false}, {label: "ArrowLeft", active: true}, {label: "ArrowLeft", active: false}, {label: "ArrowUp", active: true}, {label: "ArrowUp", active: false}},
		{{label: "ArrowDown", active: true}, {label: "ArrowDown", active: false}, {label: "ArrowLeft", active: true}, {label: "ArrowLeft", active: false}, {label: "ArrowUp", active: true}, {label: "ArrowUp", active: false}},
		{{label: "ArrowRight", active: true}, {label: "ArrowDown", active: true}, {label: "ArrowLeft", active: true}, {label: "ArrowUp", active: true}},
		{{label: "ArrowDown", active: true}, {label: "ArrowDown", active: false}, {label: "ArrowRight", active: true}, {label: "ArrowRight", active: false}},
		{{label: "ArrowRight", active: true}, {label: "ArrowLeft", active: true}},
		{{label: "ArrowLeft", active: true}, {label: "ArrowUp", active: true}}}

	for k := range queue.listName {
		if queue.listName[k].label != namesToTest[k] {
			t.Fatal("names of events orders error")
		}
	}

	for ke := range queue.list {
		for k := range queue.list[ke] {
			if queue.list[ke][k].label != eventsToTest[ke][k].label || queue.list[ke][k].active != eventsToTest[ke][k].active {
				t.Fatal("order of events list error")
			}
		}
	}

	queue.AddOppositeEvent("ArrowLeft", "ArrowRight")

	var empty bool
	var list []Event

	empty, list = queue.AddEvent("ArrowRight", true)
	if empty != false || len(list) != 1 || list[0].label != "ArrowRight" {
		t.Fatal("list inconsistency, test 1")
	}

	empty, list = queue.AddEvent("ArrowLeft", true)
	if empty != false || len(list) != 1 || list[0].label != "specialKeyFire2" {
		t.Fatal("list inconsistency, test 2")
	}

	empty, list = queue.AddEvent("ArrowRight", false)
	if empty != false || len(list) != 1 || list[0].label != "ArrowLeft" {
		t.Fatal("list inconsistency, test 3")
	}

	empty, list = queue.AddEvent("ArrowLeft", false)
	if empty != true || len(list) != 0 {
		t.Fatal("list inconsistency, test 4")
	}

	empty, list = queue.AddEvent("ArrowLeft", true)
	if empty != false || len(list) != 1 || list[0].label != "ArrowLeft" {
		t.Fatal("list inconsistency, test 5")
	}

	empty, list = queue.AddEvent("ArrowRight", true)
	if empty != false || len(list) != 1 || list[0].label != "ArrowRight" {
		t.Fatal("list inconsistency, test 6")
	}

	empty, list = queue.AddEvent("ArrowRight", false)
	if empty != false || len(list) != 1 || list[0].label != "ArrowLeft" {
		t.Fatal("list inconsistency, test 7")
	}

	empty, list = queue.AddEvent("ArrowLeft", false)
	if empty != true || len(list) != 0 {
		t.Fatal("list inconsistency, test 8")
	}

	empty, list = queue.AddEvent("ArrowDown", true)
	empty, list = queue.AddEvent("ArrowDown", false)
	empty, list = queue.AddEvent("ArrowLeft", true)
	empty, list = queue.AddEvent("ArrowLeft", false)
	empty, list = queue.AddEvent("ArrowUp", true)
	empty, list = queue.AddEvent("ArrowUp", false)
	if empty != false || len(list) != 1 || list[0].label != "specialKeyFire4" {
		t.Fatal("list inconsistency, test 9")
	}

	empty, list = queue.AddEvent("ArrowDown", true)
	time.Sleep(300 * time.Millisecond)
	empty, list = queue.AddEvent("ArrowDown", false)
	time.Sleep(300 * time.Millisecond)
	empty, list = queue.AddEvent("ArrowLeft", true)
	time.Sleep(300 * time.Millisecond)
	empty, list = queue.AddEvent("ArrowLeft", false)
	time.Sleep(300 * time.Millisecond)
	empty, list = queue.AddEvent("ArrowUp", true)
	time.Sleep(300 * time.Millisecond)
	empty, list = queue.AddEvent("ArrowUp", false)
	if empty != false || len(list) != 1 || list[0].label != "specialKeyFire4" {
		t.Fatal("list inconsistency, test 9")
	}

	empty, list = queue.AddEvent("ArrowDown", true)
	time.Sleep(300 * time.Millisecond)
	empty, list = queue.AddEvent("ArrowDown", false)
	time.Sleep(300 * time.Millisecond)
	empty, list = queue.AddEvent("ArrowLeft", true)
	time.Sleep(600 * time.Millisecond) //over timeout
	empty, list = queue.AddEvent("ArrowLeft", false)
	time.Sleep(300 * time.Millisecond)
	empty, list = queue.AddEvent("ArrowUp", true)
	time.Sleep(300 * time.Millisecond)
	empty, list = queue.AddEvent("ArrowUp", false)
	if empty != true || len(list) != 0 {
		t.Fatal("list inconsistency, test 10")
	}
}
