package StatePattern

type Context struct {
	state State
}

func (c *Context) Context() {
	c.state = nil
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) GetState() State {
	return c.state
}
