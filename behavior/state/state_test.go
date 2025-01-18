package state

import "testing"

func TestState(t *testing.T) {
	s := NewSwitcher()
	s.SwitchOff()
	s.SwitchOn()
	s.SwitchOff()
	s.SwitchOn()
	s.SwitchOn()
}
