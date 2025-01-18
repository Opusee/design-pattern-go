package flyweight

import "fmt"

/**
享元模式
*/

type Tile struct {
	image string //地砖所用的图片材质
	//地砖所在坐标
	x int
	y int
}

func NewTile(image string, x, y int) *Tile {
	t := new(Tile)
	t.image = image
	fmt.Printf("从磁盘加载[%s]图片，耗时半秒。。。", image)
	t.x = x
	t.y = y
	return t
}

func (t *Tile) Draw() {
	fmt.Printf("在位置 (%d,%d) 上绘制图片：[%s]\n", t.x, t.y, t.image)
}

type Drawable interface {
	Draw(x, y int) //绘制方法，接收地图坐标。
}

type Water struct {
	image string //河流图片材质
}

func NewWater() *Water {
	w := &Water{image: "河流"}
	fmt.Printf("从磁盘加载[%s]图片，耗时半秒。。。", w.image)
	return w
}

func (w *Water) Draw(x, y int) {
	fmt.Printf("在位置(%d,%d)上绘制图片：[%s]\n", x, y, w.image)
}

type Grass struct {
	image string //草坪图片材质
}

func NewGrass() *Grass {
	g := &Grass{image: "草坪"}
	fmt.Printf("从磁盘加载[%s]图片，耗时半秒。。。", g.image)
	return g
}

func (g *Grass) Draw(x, y int) {
	fmt.Printf("在位置(%d,%d)上绘制图片：[%s]\n", x, y, g.image)
}

type Stone struct {
	image string //石路图片材质
}

func NewStone() *Stone {
	s := &Stone{image: "石路"}
	fmt.Printf("从磁盘加载[%s]图片，耗时半秒。。。", s.image)
	return s
}

func (s *Stone) Draw(x, y int) {
	fmt.Printf("在位置(%d,%d)上绘制图片：[%s]\n", x, y, s.image)
}

type House struct {
	image string //房子图片材质
}

func NewHouse() *House {
	h := &House{image: "房子"}
	fmt.Printf("从磁盘加载[%s]图片，耗时一秒。。。", h.image)
	return h
}

func (h *House) Draw(x, y int) {
	println("将图层切到最上层。。。") //房子盖在地上，所以切换到顶层图层。
	fmt.Printf("在位置(%d,%d)上绘制图片：[%s]\n", x, y, h.image)
}

// Factory 图件工厂
type Factory struct {
	images map[string]Drawable //图库
}

func NewFactory() *Factory {
	return &Factory{images: make(map[string]Drawable, 0)}
}

func (f *Factory) GetDrawable(image string) Drawable {
	//缓存里如果没有图件，则实例化并放入缓存。
	if _, ok := f.images[image]; !ok {
		switch image {
		case "河流":
			f.images[image] = NewWater()
		case "草坪":
			f.images[image] = NewGrass()
		case "石路":
			f.images[image] = NewStone()
		}
	}
	//缓存里必然有图，直接取得并返回。
	return f.images[image]
}
