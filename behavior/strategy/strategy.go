package strategy

/**
策略模式
*/

// Strategy 算法标准
type Strategy interface {
	Calculate(a, b int) int //操作数，被操作数
}

// Addition 实现算法接口
type Addition struct {
}

// Calculate 加数与被加数
func (add *Addition) Calculate(a, b int) int {
	return a + b //这里我们做加法运算
}

type Subtraction struct {
}

func (sub *Subtraction) Calculate(a, b int) int {
	return a - b //这里我们做减法运算
}

// Calculator 计算器类
type Calculator struct {
	strategy Strategy //拥有某种算法策略
}

// SetStrategy 接入算法策略
func (cal *Calculator) SetStrategy(strategy Strategy) {
	cal.strategy = strategy
}

// GetResult 返回具体策略的结果
func (cal *Calculator) GetResult(a, b int) int {
	return cal.strategy.Calculate(a, b)
}
