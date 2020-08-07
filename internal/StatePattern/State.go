package StatePattern

import "fmt"

type State interface {
	DoAction(context *Context)
	ToString() string
}

type StartState struct {
}

func (s *StartState) DoAction(context *Context) {
	fmt.Println("Play is in start state")
	context.SetState(s)
}

func (s *StartState) ToString() string {
	return "Start State"
}

type StopState struct {
}

func (s *StopState) DoAction(context *Context) {
	fmt.Println("Play is in stop state")
	context.SetState(s)
}

func (s *StopState) ToString() string {
	return "Stop State"
}
