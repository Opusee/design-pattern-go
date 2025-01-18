package visitor

import (
	"testing"
	"time"
)

func TestVisitor(t *testing.T) {
	strToDate := func(str string) time.Time {
		if date, err := time.ParseInLocation("2006-01-02", str, time.Local); err != nil {
			panic(err)
		} else {
			return date
		}
	}

	products := []Product{
		NewCandy("小黑兔奶糖", strToDate("2018-10-01"), 20.0),
		NewWine("猫泰白酒", strToDate("2017-01-01"), 1000.0),
		NewFruit("草莓", strToDate("2018-12-26"), 10.0, 2.5),
	}

	visitor := NewDiscountVisitor(strToDate("2019-01-01"))
	for _, p := range products {
		p.Accept(visitor)
	}
}
