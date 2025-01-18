package proxy

import (
	"fmt"
	"reflect"
	"testing"
)

/**
  @link: http://dlj.bz/ubpn20
*/

func TestProxy(t *testing.T) {
	p := NewRouterProxy()
	p.Access("http://www.电影.com")
	p.Access("http://www.游戏.com")
	p.Access("ftp://www.学习.com/java")
	p.Access("http://www.工作.com")
}

type hello struct {
	name string
}

func TestReflect(t *testing.T) {
	h := &hello{name: "测试"}
	to := reflect.TypeOf(h)
	vo := reflect.ValueOf(h)
	fmt.Printf("%v\n", to.Kind())
	fmt.Printf("%v\n", to.Elem().Kind())
	fmt.Printf("%v\n", vo.Kind())
	fmt.Printf("%v\n", vo)
	fmt.Printf("%v\n", to)
	fmt.Printf("%v\n", to.Elem())
	voe := vo.Elem()
	fmt.Printf("%v\n", voe)
	fmt.Printf("%v\n", voe.NumField())
	fmt.Printf("%v\n", voe.NumMethod())
	to = to.Elem()
	println("--------------")
	for i := 0; i < voe.NumField(); i++ {
		tof := to.Field(i)
		voef := voe.Field(i)
		fmt.Printf("%v func: %t\n", tof, tof.Type.Kind() == reflect.Func)
		fmt.Printf("%v can set: %t\n", voef, voef.CanSet())
		fmt.Printf("%v\n", voef.Type())
	}
}
