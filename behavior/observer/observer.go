package observer

import "strings"

/**
观察者模式
*/

type Shop struct {
	Product string
	Buyers  []Buyer // 持有买家的引用
}

func NewShop() *Shop {
	return &Shop{
		Product: "无商品",
		Buyers:  make([]Buyer, 0),
	}
}

// Register 为了主动通知买家，买家得来店里注册。
func (s *Shop) Register(buyer Buyer) {
	s.Buyers = append(s.Buyers, buyer)
}

func (s *Shop) GetProduct() string {
	return s.Product
}

func (s *Shop) SetProduct(product string) {
	s.Product = product // 到货了
	s.NotifyBuyers()    // 到货后通知买家
}

// NotifyBuyers 通知所有注册买家
func (s *Shop) NotifyBuyers() {
	for _, b := range s.Buyers {
		b.Inform()
	}
}

type buyer struct {
	Name string
	Shop *Shop
}

func NewBuyer(name string, shop *Shop) *buyer {
	return &buyer{
		Name: name,
		Shop: shop,
	}
}

type Buyer interface {
	Inform()
}

type PhoneFans struct {
	*buyer
}

// NewPhoneFans 调用父类进行构造
func NewPhoneFans(name string, shop *Shop) *PhoneFans {
	return &PhoneFans{NewBuyer(name, shop)}
}

func (p *PhoneFans) Inform() {
	product := p.Shop.GetProduct()
	if strings.Contains(product, "水果手机") { //此买家只买水果牌手机
		println(p.Name, "购买：", product)
	}
}

type HandChopper struct {
	*buyer
}

func NewHandChopper(name string, shop *Shop) *HandChopper {
	return &HandChopper{NewBuyer(name, shop)}
}

func (h *HandChopper) Inform() {
	product := h.Shop.GetProduct()
	println(h.Name, "购买：", product)
}
