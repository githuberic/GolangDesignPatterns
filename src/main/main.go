package main

import (
	"FactoryPattern"
	"SingletonPattern"
	"fmt"
)

func main() {
	testSingleton()
}

func testSimplenessFactory() {
	f := new(FactoryPattern.SimplenessFactory)
	var s FactoryPattern.Shape
	s, ok := f.GetShape("Rectangle")
	if ok {
		s.Draw()
	}
}

func testAbstractFactory() {
	f := new(FactoryPattern.AbsFactory)
	color := f.GetColor("Red")
	color.Fill()
	shape := f.GetShape("Rectangle")
	shape.Draw()
}

func testSingleton() {
	s1 := SingletonPattern.GetInstance1()
	s1.Count = 5
	fmt.Println(s1)
	s2 := SingletonPattern.GetInstance1()
	fmt.Println(s2)
}
