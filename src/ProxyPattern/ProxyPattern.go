package ProxyPattern

import "fmt"

type Image interface {
	Display()
}

type RealImage struct {
	fileName string
}

func (r *RealImage) Display() {
	fmt.Println("Displaying " + r.fileName)
}

func (r *RealImage) loadFromDisk(fileName string) {
	fmt.Println("Loading " + r.fileName)
}

func NewRealImage(fileName string) *RealImage {
	realImage := new(RealImage)
	realImage.fileName = fileName
	realImage.loadFromDisk(fileName)
	return realImage
}

type ProxyImage struct {
	fileName  string
	realImage *RealImage
}

func (r *ProxyImage) Display() {
	if r.realImage == nil {
		r.realImage = NewRealImage(r.fileName)
	}
	r.realImage.Display()
}

func NewProxyImage(fileName string) *ProxyImage {
	realImage := new(ProxyImage)
	realImage.fileName = fileName
	return realImage
}
