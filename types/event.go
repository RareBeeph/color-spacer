package types

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"sync"
)

type Event struct {
	sync.Mutex

	StopPropagating bool

	EventType  EType
	MousePos   pixel.Vec
	MouseVel   pixel.Vec
	InitialPos pixel.Vec
	Buttons    []pixelgl.Button
}

type EType int

const (
	Click EType = iota
	Drag
)

func (e *Event) Contains(b pixelgl.Button) bool {
	for _, t := range e.Buttons {
		if b == t {
			return true
		}
	}
	return false
}
