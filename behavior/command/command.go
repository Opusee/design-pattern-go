package command

/**
命令模式
*/

// Switchable 电器接口
type Switchable interface {
	On()  //通电
	Off() //断电
}

type Device interface {
	Switchable
	ChannelUp()   // 频道+
	ChannelDown() // 频道-
	VolumeUp()    // 音量+
	VolumeDown()  // 音量-
}

type TV struct {
}

func (tv *TV) On() {
	println("电视机启动")
}

func (tv *TV) Off() {
	println("电视机关闭")
}

func (tv *TV) ChannelUp() {
	println("电视机频道+")
}

func (tv *TV) ChannelDown() {
	println("电视机频道-")
}

func (tv *TV) VolumeUp() {
	println("电视机音量+")
}

func (tv *TV) VolumeDown() {
	println("电视机音量-")
}

type Radio struct {
}

func (r *Radio) On() {
	println("收音机启动")
}

func (r *Radio) Off() {
	println("收音机关闭")
}

func (r *Radio) ChannelUp() {
	println("收音机频道+")
}

func (r *Radio) ChannelDown() {
	println("收音机频道-")
}

func (r *Radio) VolumeUp() {
	println("收音机音量+")
}

func (r *Radio) VolumeDown() {
	println("收音机音量-")
}

type Command interface {
	Exe()   //执行命令操作
	UnExe() //反执行命令操作
}

type SwitchCommand struct {
	device Device // 此处持有高级设备接口。
}

func NewSwitchCommand(device Device) *SwitchCommand {
	return &SwitchCommand{device: device}
}

func (s *SwitchCommand) Exe() {
	s.device.On()
}

func (s *SwitchCommand) UnExe() {
	s.device.Off()
}

type ChannelCommand struct {
	device Device
}

func NewChannelCommand(device Device) *ChannelCommand {
	return &ChannelCommand{device: device}
}

func (c *ChannelCommand) Exe() {
	c.device.ChannelUp()
}

func (c *ChannelCommand) UnExe() {
	c.device.ChannelDown()
}

type VolumeCommand struct {
	device Device
}

func NewVolumeCommand(device Device) *VolumeCommand {
	return &VolumeCommand{device: device}
}

func (v *VolumeCommand) Exe() {
	v.device.VolumeUp()
}

func (v *VolumeCommand) UnExe() {
	v.device.VolumeDown()
}

type Controller struct {
	okCommand         Command
	verticalCommand   Command
	horizontalCommand Command
}

// BindOKCommand 绑定OK键命令
func (c *Controller) BindOKCommand(okCommand Command) {
	c.okCommand = okCommand
}

// BindVerticalCommand 绑定上下方向键命令
func (c *Controller) BindVerticalCommand(verticalCommand Command) {
	c.verticalCommand = verticalCommand
}

// BindHorizontalCommand 绑定左右方向键命令
func (c *Controller) BindHorizontalCommand(horizontalCommand Command) {
	c.horizontalCommand = horizontalCommand
}

// ButtonOKHold 开始按键映射命令
func (c *Controller) ButtonOKHold() {
	print("长按OK按键……")
	c.okCommand.Exe()
}

func (c *Controller) ButtonOKClick() {
	print("单击OK按键……")
	c.okCommand.UnExe()
}

func (c *Controller) ButtonUpClick() {
	print("单击↑按键……")
	c.verticalCommand.Exe()
}

func (c *Controller) ButtonDownClick() {
	print("单击↓按键……")
	c.verticalCommand.UnExe()
}

func (c *Controller) ButtonLeftClick() {
	print("单击←按键……")
	c.horizontalCommand.UnExe()
}

func (c *Controller) ButtonRightClick() {
	print("单击→按键……")
	c.horizontalCommand.Exe()
}
