package baseAndDecorator

import (
	"fmt"
)

// ColorDecorator
/*  @Description: 一个颜色装饰器，需要将形状的类进行结合
*/
type ColorDecorator struct {
	Color   string
	DeShape ShapeBase
}

// NewColorDecorator
/* @Description: 获取一个装饰器
*  @param c
*  @param s
*  @return *ColorDecorator
*/
func NewColorDecorator(c string, s ShapeBase) *ColorDecorator {
	return &ColorDecorator{
		Color: c,
		DeShape: s,
	}
}
// Draw
/* @Description: 使用颜色装饰器重新复写Draw函数
*  @receiver cd
*/
func (cd *ColorDecorator) Draw() {
	fmt.Println("ColorDecorator Add New Color = ", cd.Color)
	cd.DeShape.Draw()
}
