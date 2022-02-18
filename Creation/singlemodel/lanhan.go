package singlemodel

import "sync"

type SingleLhan struct {}

var instance *SingleLhan	//懒汉模式的单例
var lock sync.Mutex	//锁
var once sync.Once	//原子操作

// GetEhanInstance
/* @Description: 直接获取，线程不安全，会导致多个创建
*  @return *SingleEhan
 */
func GetEhanInstance() *SingleLhan {
	if instance == nil {
		instance = new(SingleLhan)
	}

	return instance
}

// GetInstanceByMutex
/* @Description: 先加锁，解决并发安全，但是效率降低
*  @return *SingleEhan
 */
func GetInstanceByMutex() *SingleLhan {
	lock.Lock()			//进来先加锁
	defer lock.Unlock()

	if instance == nil {
		instance = new(SingleLhan)
	}

	return instance
}

// GetInstanceByInner
/* @Description: 加锁了，但是也非并发安全
*  @return *SingleEhan
 */
func GetInstanceByInner() *SingleLhan  {
	if instance == nil {
		lock.Lock()
		instance = new(SingleLhan)
		lock.Unlock()
	}

	return instance
}

// GetInstanceByDouble
/* @Description: 双重检测，保证安全
*  @return *SingleEhan
 */
func GetInstanceByDouble() *SingleLhan {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = new(SingleLhan)
		}
		lock.Unlock()
	}

	return instance
}

// GetInstanceByOnce
/* @Description: go语言的原子操作
*  @return *SingleLhan
 */
func GetInstanceByOnce() *SingleLhan {
	once.Do(func() {
		instance = new(SingleLhan)
	})

	return instance
}
