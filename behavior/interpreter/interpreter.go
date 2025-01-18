package interpreter

import (
	"fmt"
	"time"
)

/**
解释器模式
*/

// Expression 表达式接口
type Expression interface {
	Interpret() // 解释方法
}

type Move struct {
	// 鼠标位置坐标
	x int
	y int
}

func NewMove(x, y int) *Move {
	return &Move{
		x: x,
		y: y,
	}
}

func (m *Move) Interpret() {
	fmt.Printf("移动鼠标：[%d,%d]\n", m.x, m.y)
}

type LeftDown struct {
}

func (l *LeftDown) Interpret() {
	println("按下鼠标：左键")
}

type LeftUp struct {
}

func (l *LeftUp) Interpret() {
	println("按下鼠标：上键")
}

type RightDown struct {
}

func (l *RightDown) Interpret() {
	println("按下鼠标：右键")
}

type Delay struct {
	seconds int // 延时秒数
}

func NewDelay(seconds int) *Delay {
	return &Delay{seconds: seconds}
}

func (d *Delay) GetSeconds() int {
	return d.seconds
}

func (d *Delay) Interpret() {
	fmt.Printf("系统延迟：%d 秒钟\n", d.seconds)
	time.Sleep(time.Second * time.Duration(d.seconds))
}

type LeftClick struct {
	LeftDown *LeftDown
	LeftUp   *LeftUp
}

func NewLeftClick() *LeftClick {
	return &LeftClick{
		LeftDown: new(LeftDown),
		LeftUp:   new(LeftUp),
	}
}

func (l *LeftClick) Interpret() {
	//单击=先按下再松开
	l.LeftDown.Interpret()
	l.LeftDown.Interpret()
}

type Repetition struct {
	loopCount  int        // 循环次数
	expression Expression // 循环体表达式
}

func NewRepetition(expression Expression, loopCount int) *Repetition {
	return &Repetition{
		loopCount:  loopCount,
		expression: expression,
	}
}

func (r *Repetition) Interpret() {
	for r.loopCount > 0 {
		r.expression.Interpret()
		r.loopCount--
	}
}

type Sequence struct {
	expressions []Expression // 指令集序列
}

func NewSequence(expressions []Expression) *Sequence {
	return &Sequence{expressions}
}

func (s *Sequence) Interpret() {
	// 循环挨个解析每条指令
	for _, exp := range s.expressions {
		exp.Interpret()
	}
}
