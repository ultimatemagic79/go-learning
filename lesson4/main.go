package main

import "fmt"

// 関数
// func 関数名(引数 型) 戻り値の型 {}

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return result
}

func Noreturn() {
	fmt.Println("No return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 4)
	fmt.Println(i2, i3)
	Noreturn()
}
