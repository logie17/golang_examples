package main

import (
	"fmt"
)

type Thing interface {
	quack()
}

type thing struct {
	item1 string
}

func (this thing) quack(){
	fmt.Println(this.item1)
	fmt.Println("Quack!")
}

func newThing() Thing {
	return thing{"Ugly Duckling!"}
}

// Slice Example
func AggegateDucks(ducks []Thing) {
	fmt.Println(ducks)
}

// Varadic Example
func AggregateVaradicDucks(ducks ...Thing) {
	fmt.Println(ducks)
}

func main () {
	t1 := newThing()
	t1.quack()

	t2 := thing{"Pretty Duck"}
	t2.quack()

	AggegateDucks([]Thing{t1, t2})

	AggregateVaradicDucks(t1,t2)
}
