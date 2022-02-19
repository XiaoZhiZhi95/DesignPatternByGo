package factory

import ."DesignPattern/Creation/simplefactory/product"

// GunFactory 工厂类，实现获取不同的子类
type GunFactory struct {}

// GetGun
/* @Description: 实现获取不同的子类产品
*  @receiver gf
*  @param t
*  @param n
*  @return IGun
*/
func (gf *GunFactory)GetGun(t int, n string) IGun {
	switch t {
	case 1:
		return NewAk47(n)
	case 2:
		return NewM4(n)
	default:
		return NewAk47(n)
	}
}


