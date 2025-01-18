package proxy

import (
	"strings"
)

/**
代理模式
*/

// Internet 互联网访问接口
type Internet interface {
	Access(url string)
}

// modem 调制解调器
type modem struct {
}

func (m *modem) Access(url string) {
	println("正在访问：", url)
}

// RouterProxy 路由器代理类
type RouterProxy struct {
	modem     Internet //持有被代理类引用
	blackList []string // 黑名单
}

func NewRouterProxy() Internet {
	r := &RouterProxy{
		modem:     new(modem), //实例化被代理类
		blackList: []string{"电影", "游戏", "音乐", "小说"},
	}
	println("拨号上网...连接成功！")
	return r
}

func (r *RouterProxy) Access(url string) {
	//循环黑名单
	for _, keyword := range r.blackList {
		if strings.Contains(url, keyword) {
			println("禁止访问：", url)
			return
		}
	}
	r.modem.Access(url) //正常访问互联网
}

// Intranet 局域网访问接口
type Intranet interface {
	FileAccess(path string)
}

// Switch 交换机
type Switch struct {
}

func (s *Switch) FileAccess(path string) {
	println("访问内网：", path)
}

/**
  反射实现动态代理
*/

type KeywordFilter struct {
	blackList []string
	// 被代理的真实对象,猫、交换机、或是别的什么都是。
	origin any
}

func NewKeywordFilter(origin any) *KeywordFilter {
	k := &KeywordFilter{
		blackList: []string{"电影", "游戏", "音乐", "小说"},
		origin:    origin, //注入被代理对象
	}
	println("开启关键字过滤模式...")
	return k
}

//func (k *KeywordFilter) Invoke(origin any, args []any) error {
//	//要被切入方法面之前的业务逻辑
//	arg := args[0].(string)
//	for _, keyword := range k.blackList {
//		if strings.Contains(arg, keyword) {
//			println("禁止访问：", arg)
//			return errors.New("禁止访问：" + arg)
//		}
//	}
//	proxys.Proxy.NewProxyInstance(origin, func(obj any, method proxys.InvocationMethod, args []reflect.Value) []reflect.Value {
//		return []reflect.Value{reflect.ValueOf("This is a proxy function")}
//	})
//	// 调用真实的被代理对象方法
//}
