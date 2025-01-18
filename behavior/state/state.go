package state

/**
状态模式
*/

type State interface {
	SwitchOn(switcher *Switcher)  //开
	SwitchOff(switcher *Switcher) //关
}

type On struct {
}

func (o *On) SwitchOn(*Switcher) {
	println("WARN!!!通电状态无需再开")
}

func (o *On) SwitchOff(switcher *Switcher) {
	switcher.SetState(new(Off))
	println("OK...灯灭")
}

type Off struct {
}

func (o *Off) SwitchOn(switcher *Switcher) {
	switcher.SetState(new(On))
	println("OK...灯亮")
}

func (o *Off) SwitchOff(*Switcher) {
	println("WARN!!!断电状态无需再关")
}

type Switcher struct {
	state State
}

func NewSwitcher() *Switcher {
	//开关的初始状态设置为“关”
	return &Switcher{new(Off)}
}

func (s *Switcher) GetState() State {
	return s.state
}

func (s *Switcher) SetState(state State) {
	s.state = state
}

func (s *Switcher) SwitchOn() {
	s.state.SwitchOn(s) //这里调用的是当前状态的开方法
}

func (s *Switcher) SwitchOff() {
	s.state.SwitchOff(s) //这里调用的是当前状态的关方法
}
