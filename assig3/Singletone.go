package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Data: "Initial data for the singleton",
		}
	})
	return instance
}

func main() {
	instance := GetInstance()
	instance.Data = "New data for the singleton"
	instance2 := GetInstance()
	fmt.Println("Instance Data:", instance2.Data)
}
