package prototype

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPrototype(t *testing.T) {
	// 敌机制造工厂
	factory := NewEnemyPlaneFactory()
	planes := make([]EnemyPlane, 0)
	for i := 0; i < 10; i++ {
		//此处随机位置产生敌机
		plane := factory.GetInstance(rand.Intn(200))
		planes = append(planes, *plane)
	}
	fmt.Printf("生产敌机: %+v", planes)
}
