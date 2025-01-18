package command

import "testing"

func TestCommand(t *testing.T) {
	println("===客户端用【可编程式遥控器】操作电器===")
	tv := new(TV)
	radio := new(Radio)
	controller := new(Controller)

	//绑定【电视机】的【命令】到【控制器按键】
	controller.BindOKCommand(NewSwitchCommand(tv))
	//上下调台
	controller.BindVerticalCommand(NewChannelCommand(tv))
	//左右调音
	controller.BindHorizontalCommand(NewVolumeCommand(tv))

	controller.ButtonOKHold()
	controller.ButtonUpClick()
	controller.ButtonUpClick()
	controller.ButtonDownClick()
	controller.ButtonRightClick()

	println("\n")

	//绑定【收音机】的【命令】到【控制器按键】
	controller.BindOKCommand(NewSwitchCommand(radio))
	//上下调台
	controller.BindVerticalCommand(NewChannelCommand(radio))
	//左右调音
	controller.BindHorizontalCommand(NewVolumeCommand(radio))

	controller.ButtonOKHold()
	controller.ButtonUpClick()
	controller.ButtonUpClick()
	controller.ButtonRightClick()
	controller.ButtonDownClick()
}
