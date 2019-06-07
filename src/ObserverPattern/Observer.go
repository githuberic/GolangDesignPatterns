package ObserverPattern

import (
	"fmt"
)

type Observer interface {
	Update()
}

type BinaryObserver struct {
	subject *Subject
}

func (b *BinaryObserver) BinaryObserver(subject *Subject) {
	b.subject = subject
	b.subject.Attach(b)
}

func (b *BinaryObserver) Update() {
	fmt.Println("Binary String: ", b.subject.GetState())
}

type OctalObserver struct {
	subject *Subject
}

func (o *OctalObserver) OctalObserver(subject *Subject) {
	o.subject = subject
	o.subject.Attach(o)
}

func (o *OctalObserver) Update() {
	fmt.Println("Octal String:", o.subject.GetState())
}

type HexaObserver struct {
	subject *Subject
}

func (h *HexaObserver) HexaObserver(subject *Subject) {
	h.subject = subject
	h.subject.Attach(h)
}

func (h *HexaObserver) Update() {
	fmt.Println("Hex String:", h.subject.GetState())
}
