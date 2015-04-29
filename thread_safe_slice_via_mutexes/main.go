package main

import (
	"fmt"
	"sync"
)

type safeSlice struct {
	mySlice []interface{}
	mutex *sync.RWMutex
}

func (ss *safeSlice) Append(value interface{}) {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()
	ss.mySlice = append(ss.mySlice, value)
}

func NewSafeSlice() *safeSlice {
	return &safeSlice{make([]interface{}, 0), new(sync.RWMutex)}
}

func main() {
	slicer := NewSafeSlice()
	slicer.Append("foo")
	fmt.Println(slicer)
}
