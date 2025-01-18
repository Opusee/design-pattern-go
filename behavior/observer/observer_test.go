package observer

import "testing"

func TestObserver(t *testing.T) {
	shop := NewShop()
	tanSir := NewPhoneFans("果粉唐僧", shop)
	barJeet := NewHandChopper("剁手族八戒", shop)
	shop.Register(tanSir)
	shop.Register(barJeet)

	// 商店到货
	shop.SetProduct("猪肉炖粉条")
	shop.SetProduct("水果手机【爱疯叉】")
}
