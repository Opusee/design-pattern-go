package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	println("游戏开始。。。")
	println("双方挖矿攒钱。。。")

	println()
	//第一位玩家选择了地球人族
	println("工人建造人族工厂。。。")
	var factory AbstractFactory
	factory = NewHumanFactory(10, 10)

	marine := factory.CreateLowClass()
	marine.Show()

	tank := factory.CreateMidClass()
	tank.Show()

	ship := factory.CreateHighClass()
	ship.Show()

	println()
	//另一位玩家选择了外星族
	println("工蜂建造外星虫族工厂。。。")
	factory = NewAlienFactory(200, 200)

	roach := factory.CreateLowClass()
	roach.Show()

	spitter := factory.CreateMidClass()
	spitter.Show()

	mammoth := factory.CreateHighClass()
	mammoth.Show()

	println()
	println("两族开始大混战。。。")
	marine.Attack()
	roach.Attack()
	spitter.Attack()
	tank.Attack()
	mammoth.Attack()
	ship.Attack()
}
