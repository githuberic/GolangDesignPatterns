package TemplatePattern

import "fmt"

type Football struct {
	baseGame BaseGame
}

func (f *Football) Initialize() {
	fmt.Println("Football Game Finished!")
}
func (f *Football) StartPlay() {
	f.baseGame.StartPlay()
}
func (f *Football) EndPlay() {
	fmt.Println("Football Game Started. Enjoy the game!")
}
func (f *Football) Play() {
	f.Initialize()
	f.StartPlay()
	f.EndPlay()
}
