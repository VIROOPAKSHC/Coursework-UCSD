package main

import "fmt"

func values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}

func variables() {
	var a = "Initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // every int, float variables are by default set to 0
	fmt.Println(e)

	var f float32
	fmt.Println(f)
}

func main() {
	// values()
	variables()
}
