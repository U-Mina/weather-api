package main

import "fmt"

func main() {
	names := [4]string{"hello", "thursday", "just", "pass"}
	a := names[:1]
	b := names[0:]
	fmt.Println(names, a, b)

	a[0] = "happy"
	fmt.Println(names, a, b)
}