package builder

import (
	"github.com/duke-git/lancet/v2/slice"
	"strings"
)

// Building 建筑物
type Building struct {
	// 用来模拟房子组件的堆叠
	BuildingComponents []string
}

func NewBuilding() *Building {
	return &Building{BuildingComponents: make([]string, 0)}
}

// SetBasement 地基
func (b *Building) SetBasement(basement string) {
	b.BuildingComponents = append(b.BuildingComponents, basement)
}

// setWall 墙体
func (b *Building) setWall(wall string) {
	b.BuildingComponents = append(b.BuildingComponents, wall)
}

// setRoof 屋顶
func (b *Building) setRoof(roof string) {
	b.BuildingComponents = append(b.BuildingComponents, roof)
}

// ToString 自下而上打印最终完成的房子
func (b *Building) ToString() string {
	components := b.BuildingComponents
	slice.Reverse(components)
	return strings.Join(components, "\n")
}

// Builder 施工方接口
type Builder interface {
	BuildBasement()
	BuildWall()
	BuildRoof()
	GetBuilding() *Building
}

// HouseBuilder 别墅施工方
type HouseBuilder struct {
	house *Building
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{house: NewBuilding()}
}

func (h *HouseBuilder) BuildBasement() {
	println("挖地基，部署管道、线缆，水泥加固，搭建围墙、花园。")
	h.house.SetBasement("╬╬╬╬╬╬╬╬")
}

func (h *HouseBuilder) BuildWall() {
	println("搭建木质框架，石膏板封墙并粉饰内外墙。")
	h.house.setWall("｜田｜田 田|")
}

func (h *HouseBuilder) BuildRoof() {
	println("建造木质屋顶、阁楼，安装烟囱，做好防水。")
	h.house.setWall("╱◥███◣")
}

func (h *HouseBuilder) GetBuilding() *Building {
	return h.house
}

// ApartmentBuilder 高层公寓施工方
type ApartmentBuilder struct {
	apartment *Building
}

func NewApartmentBuilder() *ApartmentBuilder {
	return &ApartmentBuilder{apartment: NewBuilding()}
}

func (a *ApartmentBuilder) BuildBasement() {
	println("深挖地基，修建地下车库，部署管道、线缆、风道。")
	a.apartment.SetBasement("╚═════════╝")
}

func (a *ApartmentBuilder) BuildWall() {
	println("搭建多层建筑框架，建造电梯井，钢筋混凝土浇灌。")
	// 此处假设固定8层
	for i := 0; i < 8; i++ {
		a.apartment.setWall("║ □ □ □ □ ║")
	}
}

func (a *ApartmentBuilder) BuildRoof() {
	println("封顶，部署通风井，做防水层，保温层。")
	a.apartment.setRoof("╔═════════╗")
}

func (a *ApartmentBuilder) GetBuilding() *Building {
	return a.apartment
}

// Director 工程总监
type Director struct {
	Builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.Builder = builder
}

func (d *Director) Direct() *Building {
	println("=====工程项目启动=====")
	// 第一步，打好地基；
	d.Builder.BuildBasement()
	// 第二步，建造框架、墙体；
	d.Builder.BuildWall()
	// 第三步，封顶；
	d.Builder.BuildRoof()
	println("=====工程项目竣工=====")
	return d.Builder.GetBuilding()
}
