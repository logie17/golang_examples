package main

import (
	"fmt"
)

type commandAction int

const (
	append1 commandAction = iota
	at
	close1
	delete
	len1
	update
)

type safeSlice chan commandData

type commandData struct {
	action commandAction
	index int
	value interface{}
	result chan<-interface{}
	data chan<-[]interface{}
	cb UpdateFunc
}

type SafeSlice interface {
	Append(interface{})
	At(int) interface{}
	Close() []interface{}
	Delete (int)
	Len () int
	Update(int, UpdateFunc)
}

type UpdateFunc func(interface{}) interface {}

func (ss safeSlice) Append(value interface{}) {
	ss<-commandData{action: append1, value: value}
}

func (ss safeSlice) Delete(index int) {
	ss<-commandData{action: delete, index: index}
}

func (ss safeSlice) Update(index int, cb UpdateFunc) {
	ss<-commandData{action: update, index: index, cb: cb}
}

func (ss safeSlice) Len() int {
	reply := make(chan interface{})
	ss<-commandData{action: len1, result: reply}
	result := <-reply
	return result.(int)
}

func (ss safeSlice) At(index int) interface{} {
	reply := make(chan interface{})
	ss <- commandData{action: at, index: index, result: reply }
	result := <-reply
	return result
}

func (ss safeSlice) Close() []interface{} {
	reply := make(chan []interface{})
	ss <- commandData{action: close1, data: reply }
	result := <-reply
	return result
}

func (ss safeSlice) Run () {
	store := make([]interface{}, 0)

	for command := range ss {
		switch command.action {
		case append1:
			store = append(store, command.value)
		case delete:
			i := command.index
			store = append(store[:i], store[i+1:]...)
		case at:
			value := store[command.index]
			command.result <- value
		case len1:
			value := len(store)
			command.result <- value
		case update:
			i := command.index
			cb := command.cb
			store[i] = cb(store[i])
		case close1:
			close(ss)
			command.data <- store
		}
	}
}

func NewSS() SafeSlice {
	ss := make(safeSlice)
	go ss.Run()
	return ss
}

func main() {
	test := NewSS()
	test.Append("foo")
	fmt.Println(test.At(0))
	test.Append("bar")
	fmt.Println(test.At(1))
	fmt.Println(test.Len())
	test.Delete(0)
	fmt.Println(test.At(0))
	fmt.Println(test.Len())

	cb := func(v interface{}) interface{} {
		return "woop!"
	}
	test.Update(0, cb)
	fmt.Println(test.At(0))

	fmt.Println(test.Close())
}
