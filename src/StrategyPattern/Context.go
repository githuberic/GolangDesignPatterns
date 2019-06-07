package StrategyPattern

type Context struct {
	strategy Strategy
}

func (c *Context) Context(strategy Strategy) *Context {
	c.strategy = strategy
	return c
}

func (c *Context) ExecuteStrategy(num1 int, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}
