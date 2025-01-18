package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	// 设屏幕宽度为 100
	screenWidth := 100
	println("游戏开始")
	var factory Factory
	factory = new(RandomFactory)
	for i := 0; i < 10; i++ {
		factory.Create(screenWidth).Show()
	}
	println("抵达关底")
	factory = new(BossFactory)
	factory.Create(screenWidth).Show()
}
