package SingletonPattern

import (
	"fmt"
	"sync"
)

type SingleObject struct {
	Count int
}

var singleObj *SingleObject

func (p *SingleObject) Singleton() {
	fmt.Println("Singleton...")
}

func GetInstance1() *SingleObject {
	if singleObj == nil {
		singleObj = new(SingleObject)
	}
	return singleObj
}

var lock *sync.Mutex = &sync.Mutex{}

func GetInstance2() *SingleObject {
	lock.Lock()
	defer lock.Unlock()
	if singleObj == nil {
		singleObj = new(SingleObject)
	}
	return singleObj
}

func GetInstance3() *SingleObject {
	if singleObj == nil {
		lock.Lock()
		defer lock.Unlock()
		singleObj = new(SingleObject)
	}
	return singleObj
}
