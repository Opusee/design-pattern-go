package chain

import "testing"

func TestChain(t *testing.T) {
	flightJohn := NewStaff("张飞")
	flightJohn.SetNextApprover(NewManager("关羽")).SetNextApprover(NewCEO("刘备"))

	//高层接触不到也没必要接触，直接找员工张飞审批。
	flightJohn.Approve(1000)
	println("***********************")
	flightJohn.Approve(4000)
	println("***********************")
	flightJohn.Approve(9000)
	println("***********************")
	flightJohn.Approve(88000)
}
