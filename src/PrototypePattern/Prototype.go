package PrototypePattern

import "fmt"

type PersonalInfo struct {
	name string
	sex  string
	age  string
}

type WorkExperience struct {
	timeArea string
	company  string
}

type Resume struct {
	PersonalInfo
	WorkExperience
}

func (s *Resume) SetPersonalInfo(name string, sex string, age string) {
	s.name = name
	s.sex = sex
	s.age = age
}

func (s *Resume) SetWorkExperience(timeArea string, company string) {
	s.timeArea = timeArea
	s.company = company
}

func (s *Resume) Display() {
	fmt.Println(s.name, s.sex, s.age)
	fmt.Println("经历：", s.timeArea, s.company)
}

func (s *Resume) Clone() *Resume {
	resume := *s
	return &resume
}
