package singlemodel

type SingleEhan struct{}

var instance2 *SingleEhan

func init() {
	instance2 = new(SingleEhan)
}

// GetEHanInstance
/* @Description: 直接获取
*  @return *SingleEhan
*/
func GetEHanInstance() *SingleEhan {
	return instance2
}