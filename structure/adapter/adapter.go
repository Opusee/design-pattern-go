package adapter

/**
适配器模式
*/

// TriplePin 定义三脚插孔接口
type TriplePin interface {
	// Electrify 参数分别为火线 live，零线 null，地线 earth
	Electrify(l, n, e int)
}

// DualPin 定义两脚插孔接口
type DualPin interface {
	Electrify(l, n int) //这里没有地线
}

type TV struct {
}

func (tv *TV) Electrify(l, n int) {
	println("火线通电：", l)
	println("零线通电：", n)
}

type Adapter struct {
	DualPinDevice DualPin
}

// NewAdapter 创建适配器地时候，需要把双插设备接入进来
func NewAdapter(dualPinDevice DualPin) *Adapter {
	return &Adapter{DualPinDevice: dualPinDevice}
}

// Electrify 适配器实现的是目标接口
func (a *Adapter) Electrify(l, n, _ int) {
	//实际上调用了被适配设备的双插通电，地线e被丢弃了。
	a.DualPinDevice.Electrify(l, n)
}

type ClassAdapter struct {
	tv TV
}

func (c *ClassAdapter) Electrify(l, n, _ int) {
	c.tv.Electrify(l, n)
}
