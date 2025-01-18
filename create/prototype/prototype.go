package prototype

import (
	"math/rand"
	"time"
)

/**
原型模式
*/

// EnemyPlane 敌机
type EnemyPlane struct {
	x int //敌机横坐标
	//敌机一开始是从顶部出来所以纵坐标y必然是0, 所以实例化的时候就不给设置 y 值
	y      int //敌机纵坐标
	bullet *Bullet
}

// Bullet 子弹
type Bullet struct {
	count int // 子弹数量
}

func NewBullet() *Bullet {
	// 初始化 5000 发子弹
	return &Bullet{count: 5000}
}

// SetCount 填充子弹
func (b *Bullet) SetCount(count int) {
	b.count = count
}

// NewEnemyPlane 构造方法
func NewEnemyPlane(x int) *EnemyPlane {
	// 设置随机种子
	rand.Seed(time.Now().Unix())
	return &EnemyPlane{
		x:      x,
		y:      0,
		bullet: NewBullet(),
	}
}

// SetX 此处开放setX，为了让克隆后的实例重新修改x坐标
func (e *EnemyPlane) SetX(x int) {
	e.x = x
}

func (e *EnemyPlane) GetX() int {
	return e.x
}

func (e *EnemyPlane) GetY() int {
	return e.y
}

// Fly 让敌机飞
func (e *EnemyPlane) Fly() {
	e.y++ //每调用一次，敌机飞行时纵坐标＋1
}

// Shoot 射击
func (e *EnemyPlane) Shoot() {
	e.bullet.count--
}

func (b *Bullet) Clone() *Bullet {
	bt := *b
	return &bt
}

// Clone 克隆方法
func (e *EnemyPlane) Clone() *EnemyPlane {
	// 值传递，创建出了一个新的对象
	ep := *e
	ep.bullet = e.bullet.Clone()
	// 返回新对象的地址
	return &ep
}

// EnemyPlaneFactory 敌机工厂
type EnemyPlaneFactory struct {
	protoType *EnemyPlane
}

func NewEnemyPlaneFactory() *EnemyPlaneFactory {
	// 创造一个敌机原型
	return &EnemyPlaneFactory{protoType: NewEnemyPlane(200)}
}

func (epf *EnemyPlaneFactory) GetInstance(x int) *EnemyPlane {
	clone := epf.protoType.Clone()
	clone.SetX(x) //重新设置克隆机的x坐标
	return clone
}
