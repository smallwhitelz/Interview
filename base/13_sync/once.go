package _3_sync

import "sync"

type Service interface {
	DoSomething()
}

// 单例模式：一个结构体只有一个实例
// singleton 推荐用小写，用大写的话，别人会绕开GetInstance方法自己创建实例
type singleton struct {
}

func (s *singleton) DoSomething() {
}

var (
	instance *singleton
	initOnce = &sync.Once{}
)

// 这种是懒汉单例
// 直接返回singleton会报将一个私有结构体返回给公共，不建议这么做，所以单例一般配合一个接口
func GetInstance() Service {
	initOnce.Do(func() {
		instance = &singleton{}
	})
	return instance
}

// 饥饿模式
var instance1 = &singleton{}

func GetInstance1() Service {
	return instance1
}

// 还有一种饥饿模式，初始化复杂结构体，借用init
var instance2 *singleton

func init() {
	// 一大堆代码
	instance2 = &singleton{}
}

func GetInstance2() Service {
	return instance2
}
