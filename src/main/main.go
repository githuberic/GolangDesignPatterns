package main

import (
	"AdapterPattern"
	"Bridge"
	"BuilderPattern"
	"FactoryPattern"
	"FilterPattern"
	"SingletonPattern"
	"fmt"
)

func main() {
	testFilterPattern()
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
	director := BuilderPattern.Director{Builder: builder}
	computer := director.Create("I7", "32G", "4T")
	fmt.Println(*computer)
}

func testAdapterPattern() {
	audioPlayer := AdapterPattern.AudioPlayer{}

	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "far far away.vlc")
	audioPlayer.Play("avi", "mind me.avi")
}

func testBridge() {
	redCircle := Bridge.Circle{}
	redCircle.Circle(100, 100, 10, &Bridge.RedCircle{})
	greenCircle := Bridge.Circle{}
	greenCircle.Circle(100, 100, 10, &Bridge.GreenCircle{})
	redCircle.Draw()
	greenCircle.Draw()
}

func testFilterPattern() {
	var persons []FilterPattern.Person
	persons = append(persons, FilterPattern.GetPerson("Robert", "Male", "Single"))
	persons = append(persons, FilterPattern.GetPerson("John", "Male", "Married"))
	persons = append(persons, FilterPattern.GetPerson("Laura", "Female", "Married"))
	persons = append(persons, FilterPattern.GetPerson("Diana", "Female", "Single"))
	persons = append(persons, FilterPattern.GetPerson("Mike", "Male", "Single"))
	persons = append(persons, FilterPattern.GetPerson("Bobby", "Male", "Single"))
	male := new(FilterPattern.CriteriaMale)
	fmt.Println(male.MeetCriteria(persons))
	female := new(FilterPattern.CriteriaFemale)
	fmt.Println(female.MeetCriteria(persons))
	single := new(FilterPattern.CriteriaSingle)
	fmt.Println(single.MeetCriteria(persons))
	singleMale := new(FilterPattern.AndCriteria)
	singleMale.AndCriteria(single, male)
	fmt.Println(singleMale.MeetCriteria(persons))
	singleFemale := new(FilterPattern.AndCriteria)
	singleFemale.AndCriteria(single, female)
	fmt.Println(singleFemale.MeetCriteria(persons))
}
