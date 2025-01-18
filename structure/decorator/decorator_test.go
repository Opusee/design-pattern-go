package decorator

import "testing"

func TestDecorator(t *testing.T) {
	//口红包裹粉底，再包裹女友。
	NewLipstick(NewFoundationMakeup(new(Girl))).Show()
	println()
}
