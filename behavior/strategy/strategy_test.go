package strategy

import "testing"

func TestStrategy(t *testing.T) {
	calculator := new(Calculator)         //实例化计算器
	calculator.SetStrategy(new(Addition)) //接入加法实现
	result := calculator.GetResult(1, 1)
	println(result)

	calculator.SetStrategy(new(Subtraction))
	result = calculator.GetResult(1, 1)
	println(result)
}
