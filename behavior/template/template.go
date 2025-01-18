package template

import (
	"math/rand"
	"time"
)

/**
模板方法模式
*/

// CigarLighterInterface 点烟器接口
type CigarLighterInterface interface {
	ElectrifyDC16V() //供电方法，16V直流电
}

// GPS 导航的实现
type GPS struct {
}

func (g *GPS) ElectrifyDC16V() {
	println("连接卫星")
	println("定位。。。")
}

// CigarLighter 点烟器的实现
type CigarLighter struct {
}

func (c *CigarLighter) ElectrifyDC16V() {
	for t := 1000; t > 0; t-- {
		println("加热电炉丝")
	}
	println("点烟器弹出")
}

type Mammal interface {
	FeedMilk()
}

// 哺乳动物
type mammal struct {
	female bool
	Move   func() // Move 哺乳动物当然可以移动，但具体怎么移动还不知道。
}

func newMammal(move func()) *mammal {
	return &mammal{
		female: true,
		Move:   move,
	}
}

// FeedMilk 既然是哺乳动物当然会喂奶了，但这里约束为只能母的喂奶
func (m *mammal) FeedMilk() {
	if m.female {
		println("喂奶")
	} else {
		println("公的不会")
	}
}

type Human struct {
	*mammal
}

func NewHuman() *Human {
	return &Human{newMammal(func() {
		println("两条腿走路……")
	})}
}

// Whale 鲸鱼
type Whale struct {
	*mammal
}

func NewWhale() *Whale {
	return &Whale{newMammal(func() {
		println("游泳……")
	})}
}

// Bat 蝙蝠
type Bat struct {
	*mammal
}

func NewBat() *Bat {
	return &Bat{newMammal(func() {
		println("用翅膀飞……")
	})}
}

// -------
// 模仿抽象类定义。已实现方法设计为接口，由抽象类实现；未实现的方法设计为函数属性 交由具体类去实现

type PM interface {
	Kickoff()
}
type pm struct {
	Analyze func()       //需求分析
	Design  func()       //设计
	Develop func(string) //开发
	Test    func() bool  //测试
	Release func()       //发布
}

func (p *pm) Kickoff() {
	p.Analyze()
	p.Design()
	p.Develop("开始开发")
	//如果测试失败，则继续开发改Bug。
	for !p.Test() {
		p.Develop("开发改 bug")
	}
	p.Release()
}

type AutoTestPM struct {
	*pm
}

func NewAutoTestPM() PM {
	rand.Seed(time.Now().Unix())
	return &AutoTestPM{&pm{
		Analyze: func() {
			println("需求分析")
		},
		Design: func() {
			println("业务设计")
		},
		Develop: func(msg string) {
			println(msg)
		},
		Test: func() bool {
			if rand.Intn(10)%2 == 0 {
				println("测试失败")
				return false
			}
			println("测试通过")
			return true
		},
		Release: func() {
			println("进行发版")
		},
	}}
}
