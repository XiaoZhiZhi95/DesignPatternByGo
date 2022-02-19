package product

import "fmt"

// m4
/*  @Description: 一个具体实现产品的子类
*/
type m4 struct {
	Gun	//以此实现对Gun的继承
}

func (m *m4)Sound()  {
	fmt.Println(m.GetName(), "M4 Sound Like BiuBiu")
}

func NewM4(n string) IGun {
	return &ak47{
		Gun{
			n,
		},
	}
}