package main

import "fmt"

func detector(v interface{}) string {
	switch v := v.(type) {
	case int:
		return fmt.Sprintf("%d - I'm an int!", v)
	case float64, float32:
		return fmt.Sprintf("%f - I'm a float!", v)
	case string:
		return fmt.Sprintf("%s - I'm a string!", v)
	}
	return "unknown"
}

func main() {
	a := 1
	b := "string"
	c := 1.1

	fmt.Println("a: " + detector(a))
	fmt.Println("b: " + detector(b))
	fmt.Println("c: " + detector(c))
	fmt.Println("d: " + detector(float32(c)))
}
