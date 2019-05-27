package main

import (
	"SimplenessFactoryPattern"
)

func main() {
	f := new(SimplenessFactoryPattern.ShapeFactory)
	var s SimplenessFactoryPattern.Shape
	s, ok := f.GetShape("Rectangle")
	if ok{
     
	}

}
