package TemplatePattern

import "fmt"

type Game interface {
	Initialize()
	StartPlay()
	EndPlay()
	Play()
}

type BaseGame struct {
}

func (b *BaseGame) Initialize() {
	fmt.Println("base Initialize")
}
func (b *BaseGame) StartPlay() {
	fmt.Println("base StartPlay")
}
func (b *BaseGame) EndPlay() {
	fmt.Println("base EndPlay")
}
func (b *BaseGame) Play() {
	b.Initialize()
	b.StartPlay()
	b.EndPlay()
}
