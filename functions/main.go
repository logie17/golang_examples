package main

import "fmt"

func StandardWithReturn(foo int) int {
	return foo
}

func Standard(foo int) {
	fmt.Println(foo)
}

func WithClosure(foo int) func(int) int {
	return func(bar int) int {
		return foo + bar
	}
}

func HigherOrder(bar int, foo func(int) int) int {
	return foo(bar)
}

func RecursiveFun(foo int) int {
	if foo < 6 {
		return foo + RecursiveFun(foo+1)
	}
	return 0
}

func VaradicFun(list_of_ints ...int) int {
	for _, x := range list_of_ints {
		if x == 5 {
			return x
		}
	}
	return 0
	
}

func main() {
	Standard(1)
	fmt.Println(StandardWithReturn(2))
	thing := WithClosure(3)
	fmt.Println(thing(2))

	fmt.Println(HigherOrder(2, func(a int) int {
		return a * 6
	}))

	fmt.Println(RecursiveFun(1))

	fmt.Println(VaradicFun(1, 3, 4, 5))

	// TODO - optional/default hacky solution
	
}
