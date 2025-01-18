package facade

/**
门面模式
*/

// VegVendor 菜贩子
type vegVendor struct {
}

func (v *vegVendor) Sell() {
	println("菜贩子卖菜。。。")
}

// 厨师
type chef struct {
}

func (c *chef) Cook() {
	println("厨师烹饪。。。")
}

type cleaner struct {
}

func (c *cleaner) Clean() {
	println("收拾桌子")
}

func (c *cleaner) Wash() {
	println("洗碗...")
}

type waiter struct {
}

func (w *waiter) Order() {
	println("接待、入座、点菜")
}

func (w *waiter) Server() {
	println("上菜...")
}

type Facade struct {
	vv      *vegVendor
	chef    *chef
	waiter  *waiter
	cleaner *cleaner
}

func NewFacade() *Facade {
	f := new(Facade)
	//开门前就找菜贩子准备好蔬菜
	f.vv = new(vegVendor)
	f.vv.Sell()
	//当然还得雇佣好各类饭店服务人员
	f.chef = new(chef)
	f.waiter = new(waiter)
	f.cleaner = new(cleaner)
	return f
}

func (f *Facade) ProvideService() {
	//接待，入座，点菜
	f.waiter.Order()
	//找厨师做饭
	f.chef.Cook()
	//上菜
	f.waiter.Server()
	//收拾桌子，洗碗，以及其他工序……
	f.cleaner.Clean()
	f.cleaner.Wash()
}
