package product

import "fmt"

// IGun
/*  @Description: 基础的抽象产品类
*/
type IGun interface {
	SetName(name string)
	GetName() string
	Sound()
}

// Gun
/*  @Description: 产品类通用的使用方法
*/
type Gun struct{
	name string
}

func (gun *Gun)SetName(name string){
	gun.name = name
}

func (gun *Gun)GetName() string {
	return gun.name
}

func (gun *Gun)Sound()  {
	fmt.Println("All Gun Sound like Gagaga")
}



