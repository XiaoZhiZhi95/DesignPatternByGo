package product

import "fmt"

// ak47
/*  @Description: 一个具体实现产品的子类
*/
type ak47 struct {
	Gun		//以此实现对Gun的继承
}

func (ak *ak47)Sound()  {
	fmt.Println(ak.GetName(), "Ak47 Sound Like BangBang")
}

func NewAk47(n string) IGun {
	return &ak47{
		Gun{
			n,
		},
	}
}
