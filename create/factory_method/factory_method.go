package factory_method

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
工厂方法模式
*/

// Enemy 敌军抽象
type Enemy interface {
	// Show 在地图上绘制敌军
	Show()
}

type enemy struct {
	x int
	y int
}

//初始化坐标
func newEnemy(x, y int) *enemy {
	return &enemy{
		x: x,
		y: y,
	}
}

// Airplane 飞机
type Airplane struct {
	*enemy
}

func NewAirplane(x, y int) *Airplane {
	//调用父类构造子初始化坐标
	return &Airplane{newEnemy(x, y)}
}

func (a *Airplane) Show() {
	fmt.Printf("飞机出现坐标: (%d,%d)\n", a.x, a.y)
	println("飞机向玩家发起攻击...")
}

// Tank 坦克
type Tank struct {
	*enemy
}

func NewTank(x, y int) *Tank {
	return &Tank{newEnemy(x, y)}
}

func (t *Tank) Show() {
	fmt.Printf("坦克出现坐标: (%d,%d)\n", t.x, t.y)
	println("坦克向玩家发起攻击...")
}

// Factory 生产敌军的工厂
type Factory interface {
	Create(screenWidth int) Enemy
}

// RandomFactory 随机工厂
type RandomFactory struct {
	once sync.Once
}

func NewRandomFactory() *RandomFactory {
	f := &RandomFactory{}
	f.once.Do(func() {
		// 设置随机数种子. 只需要设置一次，如果多次设置相同种子，则产生的随机数相同
		rand.Seed(time.Now().Unix())
	})
	return f
}

func (r *RandomFactory) Create(screenWidth int) Enemy {
	if rand.Intn(100)%2 == 0 {
		return NewAirplane(rand.Intn(screenWidth), 0)
	} else {
		return NewTank(rand.Intn(screenWidth), 0)
	}
}

// Boss 敌军
type Boss struct {
	*enemy
}

func NewBoss(x, y int) *Boss {
	return &Boss{newEnemy(x, y)}
}

func (b *Boss) Show() {
	fmt.Printf("Boss 出现坐标: (%d,%d)\n", b.x, b.y)
	println("Boss 向玩家发起攻击...")
}

// BossFactory 工厂
type BossFactory struct {
}

func (bf *BossFactory) Create(screenWidth int) Enemy {
	// boss应该出现在屏幕中央
	return NewBoss(screenWidth/2, 0)
}
