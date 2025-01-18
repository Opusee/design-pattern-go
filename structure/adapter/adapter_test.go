package adapter

import "testing"

func TestAdapter(t *testing.T) {
	tv := new(TV)
	// 使用适配器的方式，创建时需要传入对象
	ap := NewAdapter(tv)
	ap.Electrify(0, 1, 2)

	/**
	  适配器方式：
	    缺点：创建时必须传入对象
	    优点：拓展性强，适配器并不在乎传入的对象是什么，风扇、热水壶、电脑等，只要是两脚插孔的设备都可以
	*/

	// 使用类适配器的方式
	ca := new(ClassAdapter)
	ca.Electrify(0, 1, 2)

	/**
	  类适配器方式：
		缺点：只适用于电视机
		优点：创建简单无需传入对象，如果有其他品牌的电视机类型，只需要继承 TV 类型即可
	*/

	// 工作中用哪种，视具体情况而定，熟练掌握，灵活应用
}
