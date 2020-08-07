package TemplatePattern

import "fmt"

type Cricket struct {
	baseGame BaseGame
}

func (c *Cricket) Initialize() {
	fmt.Println("Cricket Game Finished!")
}
func (c *Cricket) StartPlay() {
	fmt.Println("Cricket Game Initialized! Start playing.")
}
func (c *Cricket) EndPlay() {
	c.baseGame.EndPlay()
}
func (c *Cricket) Play() {
	c.Initialize()
	c.StartPlay()
	c.EndPlay()
}