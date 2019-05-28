package main

import (
	"BuilderPattern"
	"FactoryPattern"
	"SingletonPattern"
	"fmt"
)

func main() {
	testBuilderPattern()
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

func testBuilderPattern() {
	builder := new(BuilderPattern.ComputerBuilder)
	director := BuilderPattern.Director {Builder: builder}
	computer := director.Create("I7", "32G", "4T")
	fmt.Println(*computer)
}
