package decorator

/**
装饰器模式
*/

type Showable interface {
	Show() // 定义展示行为
}

type Girl struct {
}

func (g *Girl) Show() {
	print("女孩的素颜")
}

// Decorator 化妆品粉饰器
type decorator struct {
	showable Showable //持有某个善于展示的家伙
}

// newDecorator 构造时注入这个家伙
func newDecorator(showable Showable) *decorator {
	return &decorator{showable: showable}
}

func (d *decorator) Show() {
	d.showable.Show() //直接调用不做加任何粉饰。
}

// FoundationMakeup 粉底彩妆
type FoundationMakeup struct {
	*decorator
}

func NewFoundationMakeup(showable Showable) Showable {
	return &FoundationMakeup{newDecorator(showable)} //调用化妆品父类注入
}

func (f *FoundationMakeup) Show() {
	print("打粉底(")
	f.showable.Show()
	print(")")
}

// Lipstick 口红
type Lipstick struct {
	*decorator
}

func NewLipstick(showable Showable) Showable {
	return &Lipstick{newDecorator(showable)}
}

func (l *Lipstick) Show() {
	print("涂口红(")
	l.showable.Show()
	print(")")
}
