package baseAndDecorator

import "fmt"

// ShapeBase
/*  @Description: 基础的形状类
*/
type ShapeBase interface {
	Draw()
}

// Circle circle类
type Circle struct {}
func NewCircle() *Circle {
	return &Circle{}
}
func (c *Circle) Draw()  {
	fmt.Println("Circle Draw Method")
}