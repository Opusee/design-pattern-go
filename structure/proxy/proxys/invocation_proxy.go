package proxys

import (
	"fmt"
	"reflect"
)

/**
  参考：
  @link: https://github.com/gohutool/boot4go-proxyW
*/

var Proxy = invocationProxy{}

type invocationProxy struct {
}

type InvocationHandler func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value

func (ip invocationProxy) NewProxyInstance(itf any, handler InvocationHandler) any {

	t := reflect.TypeOf(itf)

	if t.Kind() != reflect.Ptr {
		panic("Need a pointer of interface struct")
	}

	if t.Elem().Kind() != reflect.Struct {
		panic("Need a pointer of interface struct")
	}

	t = t.Elem()
	ot := reflect.ValueOf(itf).Elem()
	n := ot.NumField()

	for idx := 0; idx < n; idx++ {
		f := t.Field(idx)
		of := ot.Field(idx)

		if f.Type.Kind() == reflect.Func {

			if !of.CanSet() {
				fmt.Printf("field %v is a readonly func, ignore proxy\n", f.Name)
				continue
			}

			fmt.Printf("field %v is a func, proxy now\n", f.Name)

			target := reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
				//param := make([]any, 0, len(args))
				//
				//for _, o := range args {
				//	param = append(param, o.Interface())
				//}

				method := InvocationMethod{Name: f.Name, Type: f.Type}

				rtn := handler(itf, method, args)

				//if rtn == nil {
				//	return []reflect.Value{}
				//}
				//
				//rtnV := make([]reflect.Value, 0, len(rtn))
				//
				//for _, o := range rtn {
				//	rtnV = append(rtnV, reflect.ValueOf(o))
				//}

				return rtn
			})
			of.Set(target)
		} else {
			fmt.Printf("field %v is not a func, ignore proxy\n", f.Name)
		}
	}

	return itf
}
