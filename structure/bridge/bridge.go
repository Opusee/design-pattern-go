package bridge

/**
桥接模式
*/

// Ruler 尺子
type Ruler interface {
	Regularize() // 规则化笔触走向
}

type SquareRuler struct {
}

func (s *SquareRuler) Regularize() {
	//尺子模板画出正方形
	println("□")
}

type TriangleRuler struct {
}

func (t *TriangleRuler) Regularize() {
	//尺子模板画出三角形
	println("△")
}

type CircleRuler struct {
}

func (c *CircleRuler) Regularize() {
	//尺子模板画出圆形
	println("○")
}

type penFiled struct {
	//尺子的引用
	ruler Ruler
}

func newPenFiled(ruler Ruler) *penFiled {
	return &penFiled{ruler: ruler}
}

// Pen 画笔
type Pen interface {
	Draw()
}

type BlackPen struct {
	*penFiled
}

func NewBlackPen(ruler Ruler) *BlackPen {
	return &BlackPen{newPenFiled(ruler)}
}

func (v *BlackPen) Draw() {
	print("黑")
	v.ruler.Regularize()
}

type WhitePen struct {
	*penFiled
}

func NewWhitePen(ruler Ruler) *WhitePen {
	return &WhitePen{newPenFiled(ruler)}
}

func (r *WhitePen) Draw() {
	print("白")
	r.ruler.Regularize()
}
