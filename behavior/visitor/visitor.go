package visitor

import (
	"fmt"
	"time"
)

/**
访问者模式
*/

type product struct {
	name         string    // 商品名
	producedDate time.Time // 生产日期
	price        float64   // 价格
}

func NewProduct(name string, producedDate time.Time, price float64) *product {
	return &product{
		name:         name,
		producedDate: producedDate,
		price:        price,
	}
}

func (p *product) GetName() string {
	return p.name
}

func (p *product) SetName(name string) {
	p.name = name
}

func (p *product) GetProductDate() time.Time {
	return p.producedDate
}

func (p *product) SetProductDate(producedDate time.Time) {
	p.producedDate = producedDate
}

func (p *product) GetPrice() float64 {
	return p.price
}

func (p *product) SetPrice(price float64) {
	p.price = price
}

type Product interface {
	Acceptable
	GetName() string
	SetName(name string)
	GetProductDate() time.Time
	SetProductDate(producedDate time.Time)
	GetPrice() float64
	SetPrice(price float64)
}

// Candy 糖果类
type Candy struct {
	*product
}

func NewCandy(name string, producedDate time.Time, price float64) *Candy {
	return &Candy{NewProduct(name, producedDate, price)}
}

// Wine 酒类
type Wine struct {
	*product
}

func NewWine(name string, producedDate time.Time, price float64) *Wine {
	return &Wine{NewProduct(name, producedDate, price)}
}

// Fruit 水果
type Fruit struct {
	*product
	weight float64
}

func NewFruit(name string, producedDate time.Time, price float64, weight float64) *Fruit {
	return &Fruit{
		product: NewProduct(name, producedDate, price),
		weight:  weight,
	}
}

func (f *Fruit) GetWeight() float64 {
	return f.weight
}

func (f *Fruit) SetWeight(weight float64) {
	f.weight = weight
}

// Visitor 访问者接口
type Visitor interface {
	CandyVisit(candy *Candy)
	WineVisit(wine *Wine)
	FruitVisit(fruit *Fruit)
}

type DiscountVisitor struct {
	billDate time.Time
}

func NewDiscountVisitor(billDate time.Time) *DiscountVisitor {
	println("结算日期：", billDate.Format("2006-01-02"))
	return &DiscountVisitor{billDate: billDate}
}

func (d *DiscountVisitor) CandyVisit(candy *Candy) {
	fmt.Printf("=====糖果 [%s] 打折后价格=====\n", candy.GetName())
	rate := 0.0
	days := int(d.billDate.Sub(candy.producedDate).Hours() / 24)
	if days > 180 {
		println("超过半年过期糖果，请勿食用！")
	} else {
		rate = 0.9
	}
	discountPrice := candy.GetPrice() * rate
	fmt.Printf("¥%.2f\n", discountPrice)
}

func (d *DiscountVisitor) WineVisit(wine *Wine) {
	fmt.Printf("=====酒品 [%s] 无折扣价格=====\n", wine.GetName())
	fmt.Printf("¥%.2f\n", wine.GetPrice())
}

func (d *DiscountVisitor) FruitVisit(fruit *Fruit) {
	fmt.Printf("=====水果 [%s] 打折后价格=====\n", fruit.GetName())
	rate := 0.0
	days := int(d.billDate.Sub(fruit.producedDate).Hours() / 24)
	if days > 7 {
		println("￥0.00元（超过一周过期水果，请勿食用！）")
	} else if days > 3 {
		rate = 0.5
	} else {
		rate = 1.0
	}
	discountPrice := fruit.GetPrice() * fruit.GetWeight() * rate
	fmt.Printf("¥%.2f\n", discountPrice)
}

type Acceptable interface {
	Accept(visitor Visitor) // 主动接受拜访者
}

func (c *Candy) Accept(visitor Visitor) {
	visitor.CandyVisit(c)
}

func (w *Wine) Accept(visitor Visitor) {
	visitor.WineVisit(w)
}

func (f *Fruit) Accept(visitor Visitor) {
	visitor.FruitVisit(f)
}
