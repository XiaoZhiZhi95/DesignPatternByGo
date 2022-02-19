package main

import ."DesignPattern/Creation/simplefactory/factory"

func main()  {
	gf := GunFactory{}

	//生成对应的产品
	ak := gf.GetGun(1, "myAkåç")

	//实现声音的个性化，只需要调用这一个sound的函数
	ak.Sound()
}
