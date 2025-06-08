package main

import (
	"sync"
	"syscall/js"
)

type Draggable interface {
	GetRect() Rect
	CollisionTotalEvent(isColliding bool, b *Draggable)
	CollisionPartialEvent(isColliding bool, b *Draggable)
}

type Rect struct {
	Top, Left, Bottom, Right int
}

type ManagerCollision struct {
	draggables []Draggable
	mu         sync.Mutex
}

func (e *ManagerCollision) Init() {
	e.draggables = make([]Draggable, 0)

	js.Global().Get("document").Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e.checkAllCollisions()
		return nil
	}))

	js.Global().Get("document").Call("addEventListener", "mouseup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e.checkAllCollisions()
		return nil
	}))
}

func (e *ManagerCollision) Add(draggable Draggable) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.draggables = append(e.draggables, draggable)
	return nil
}

func (e *ManagerCollision) checkAllCollisions() {
	e.mu.Lock()
	defer e.mu.Unlock()

	collidingTotalMap := make(map[int]struct {
		total bool
		b     Draggable
	})
	collidingPartialMap := make(map[int]struct {
		total bool
		b     Draggable
	})

	for i, objA := range e.draggables {
		rectA := objA.GetRect()

		for j, objB := range e.draggables {
			if i == j {
				continue
			}

			rectB := objB.GetRect()

			isFullyInside := e.isFullyInside(rectA, rectB) || e.isFullyInside(rectB, rectA)
			isNotFullyInside := e.isColliding(rectA, rectB) || e.isColliding(rectB, rectA)
			isColliding := isFullyInside || isNotFullyInside

			if isColliding && isFullyInside {
				collidingTotalMap[i] = struct {
					total bool
					b     Draggable
				}{total: true, b: objB}
			}

			if isColliding && !isFullyInside {
				collidingPartialMap[i] = struct {
					total bool
					b     Draggable
				}{total: false, b: objB}
			}
		}
	}

	for i, objA := range e.draggables {
		if val, exists := collidingTotalMap[i]; exists {
			objA.CollisionTotalEvent(true, &val.b)
		} else {
			objA.CollisionTotalEvent(false, nil)
		}

		if val, exists := collidingPartialMap[i]; exists {
			objA.CollisionPartialEvent(true, &val.b)
		} else {
			objA.CollisionPartialEvent(false, nil)
		}
	}
}

func (e *ManagerCollision) isColliding(rectA, rectB Rect) bool {
	return !(rectA.Right < rectB.Left ||
		rectA.Left > rectB.Right ||
		rectA.Bottom < rectB.Top ||
		rectA.Top > rectB.Bottom)
}

func (e *ManagerCollision) isFullyInside(rectA, rectB Rect) bool {
	return rectA.Top <= rectB.Top &&
		rectA.Left <= rectB.Left &&
		rectA.Bottom >= rectB.Bottom &&
		rectA.Right >= rectB.Right
}
