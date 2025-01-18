package abstract_factory

import "fmt"

/**
抽象工厂模式
*/

// Unit 兵种
type Unit interface {
	Show()
	Attack()
}

type unit struct {
	attack  int // 攻击力
	defence int // 防御力
	health  int // 血量
	x       int // 横坐标
	y       int // 纵坐标
}

func newUnit(attack, defence, health, x, y int) *unit {
	return &unit{
		attack:  attack,
		defence: defence,
		health:  health,
		x:       x,
		y:       y,
	}
}

// Marine 海军陆战队士兵
type Marine struct {
	*unit
}

// NewMarine 构造方法
func NewMarine(x, y int) *Marine {
	return &Marine{newUnit(6, 5, 40, x, y)}
}

func (m *Marine) Show() {
	fmt.Printf("士兵出现在坐标：(%d,%d)\n", m.x, m.y)
}

func (m *Marine) Attack() {
	println("士兵用机关枪射击，攻击力：", m.attack)
}

// Tank 坦克
type Tank struct {
	*unit
}

func NewTank(x, y int) *Tank {
	return &Tank{newUnit(25, 100, 150, x, y)}
}

func (t *Tank) Show() {
	fmt.Printf("坦克出现在坐标：(%d,%d)\n", t.x, t.y)
}

func (t *Tank) Attack() {
	println("坦克用炮轰击，攻击力：", t.attack)
}

// Battleship 巨型战舰
type Battleship struct {
	*unit
}

func NewBattleship(x, y int) *Battleship {
	return &Battleship{newUnit(25, 200, 500, x, y)}
}

func (b *Battleship) Show() {
	fmt.Printf("战舰出现在坐标：(%d,%d)\n", b.x, b.y)
}

func (b *Battleship) Attack() {
	println("战舰用激光炮打击，攻击力：", b.attack)
}

// Roach 外星蟑螂兵
type Roach struct {
	*unit
}

func NewRoach(x, y int) *Roach {
	return &Roach{newUnit(5, 2, 35, x, y)}
}

func (r *Roach) Show() {
	fmt.Printf("蟑螂兵出现在坐标：(%d,%d)\n", r.x, r.y)
}

func (r *Roach) Attack() {
	println("蟑螂兵用爪子挠，攻击力：", r.attack)
}

// Spitter 外星毒液口水兵
type Spitter struct {
	*unit
}

func NewSpitter(x, y int) *Spitter {
	return &Spitter{newUnit(10, 8, 80, x, y)}
}

func (s *Spitter) Show() {
	fmt.Printf("口水兵出现在坐标：(%d,%d)\n", s.x, s.y)
}

func (s *Spitter) Attack() {
	println("口水兵用毒液喷射，攻击力：", s.attack)
}

// Mammoth 外星猛犸巨兽
type Mammoth struct {
	*unit
}

func NewMammoth(x, y int) *Mammoth {
	return &Mammoth{newUnit(20, 100, 400, x, y)}
}

func (m *Mammoth) Show() {
	fmt.Printf("猛犸巨兽兵出现在坐标：(%d,%d)\n", m.x, m.y)
}

func (m *Mammoth) Attack() {
	println("猛犸巨兽用獠牙顶，攻击力：", m.attack)
}

type AbstractFactory interface {
	CreateLowClass() Unit  // 工厂方法：制造低级兵种
	CreateMidClass() Unit  // 工厂方法：制造中级兵种
	CreateHighClass() Unit // 工厂方法：制造高级兵种
}

// HumanFactory 人族工厂
type HumanFactory struct {
	x int
	y int
}

func NewHumanFactory(x, y int) *HumanFactory {
	return &HumanFactory{
		x: x,
		y: y,
	}
}

func (h *HumanFactory) CreateLowClass() Unit {
	marine := NewMarine(h.x, h.y)
	println("制造海军陆战队员成功。")
	return marine
}

func (h *HumanFactory) CreateMidClass() Unit {
	tank := NewTank(h.x, h.y)
	println("制造变形坦克成功。")
	return tank
}

func (h *HumanFactory) CreateHighClass() Unit {
	battleship := NewBattleship(h.x, h.y)
	println("制造巨型战舰成功。")
	return battleship
}

type AlienFactory struct {
	x int
	y int
}

// NewAlienFactory 外星虫族工厂
func NewAlienFactory(x, y int) *AlienFactory {
	return &AlienFactory{
		x: x,
		y: y,
	}
}

func (a *AlienFactory) CreateLowClass() Unit {
	roach := NewRoach(a.x, a.y)
	println("制造蟑螂兵成功。")
	return roach
}

func (a *AlienFactory) CreateMidClass() Unit {
	spitter := NewSpitter(a.x, a.y)
	println("制造毒液兵成功。")
	return spitter
}

func (a *AlienFactory) CreateHighClass() Unit {
	mammoth := NewMammoth(a.x, a.y)
	println("制造猛犸巨兽成功。")
	return mammoth
}
