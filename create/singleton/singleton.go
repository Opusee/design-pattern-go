package singleton

import "sync"

/**
单例模式
*/

var (
	god  *God
	once sync.Once
)

// God 单例实例
type God struct {
}

// GetInstance 构造方法，用于获取单例对象
func GetInstance() *God {
	once.Do(func() {
		println("实例化 God")
		god = new(God)
	})
	return god
}
