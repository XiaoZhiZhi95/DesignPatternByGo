package main

import (
	."DesignPattern/structural/Decorator/baseAndDecorator"

)

func main() {
	var sh ShapeBase

	sh = NewCircle()
	sh.Draw()
	colorDe := NewColorDecorator("Green", sh)
	colorDe.Draw()
}
