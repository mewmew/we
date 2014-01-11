package we

import (
	"fmt"
	"image"
)

// Button corresponds to a mouse button.
type Button int

// Mouse buttons.
const (
	Button1 Button = iota
	Button2
	Button3
	Button4
	Button5
	Button6
	Button7
	Button8

	// Aliases.
	ButtonLeft   = Button1
	ButtonRight  = Button2
	ButtonMiddle = Button3
)

// buttonNames maps from Button value to string description.
var buttonNames = map[Button]string{
	ButtonLeft:   "[left button]",
	ButtonRight:  "[right button]",
	ButtonMiddle: "[middle button]",
}

func (button Button) String() string {
	s, ok := buttonNames[button]
	if ok {
		return s
	}
	return fmt.Sprintf("[button %d]", int(button)+1)
}

// MousePress is triggered when a mouse button is pressed.
type MousePress struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Mouse button that was pressed.
	Button Button
	// Bitfield of modifier keys.
	Mod Mod
}

func (e MousePress) String() string {
	// Override the embedded String method of image.Point.
	return fmt.Sprintf("{%v %v %v}", e.Point, e.Button, e.Mod)
}

// MouseRelease is triggered when a mouse button is released.
type MouseRelease struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Mouse button that was released.
	Button Button
	// Bitfield of modifier keys.
	Mod Mod
}

func (e MouseRelease) String() string {
	// Override the embedded String method of image.Point.
	return fmt.Sprintf("{%v %v %v}", e.Point, e.Button, e.Mod)
}

// MouseMove is triggered when the mouse is moved from one location to another.
type MouseMove struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Coordinates of the mouse cursor at the beginning of the move event.
	From image.Point
}

func (e MouseMove) String() string {
	// Override the embedded String method of image.Point.
	return fmt.Sprintf("{%v %v}", e.Point, e.From)
}

// MouseDrag is triggered when the mouse is moved from one location to another
// while a mouse button is held down.
type MouseDrag struct {
	// Coordinates of the mouse cursor.
	image.Point
	// Coordinates of the mouse cursor at the beginning of the drag event.
	From image.Point
	// Mouse button that was held down.
	Button Button
	// Bitfield of modifier keys.
	Mod Mod
}

// MouseEnter is triggered when the mouse enters or leaves the window. On mouse
// enter the value is true, otherwise it is false.
type MouseEnter bool

func (enter MouseEnter) String() string {
	if enter {
		return "enter"
	}
	return "leave"
}

// ScrollX is triggered when the mouse wheel is scrolled on the horizontal axis.
type ScrollX struct {
	// Horizontal scroll offset.
	Off int
	// Bitfield of modifier keys.
	Mod Mod
}

// ScrollY is triggered when the mouse wheel is scrolled on the vertical axis.
type ScrollY struct {
	// Vertical scroll offset.
	Off int
	// Bitfield of modifier keys.
	Mod Mod
}
