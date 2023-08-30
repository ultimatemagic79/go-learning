package main

import "fmt"

// 変数宣言
// var 変数名 型 = 値
var i5 int = 5

func outer() {
	var out string = "outer"
	fmt.Println(out)
}

func main() {
	var i int = 1
	fmt.Println(i)

	var s string = "Hello World"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 2
		s2 string = "Go lang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 3
	s3 = "Go lang"
	fmt.Println(i3, s3)

	i = 4
	fmt.Println(i)

	// 暗黙的な定義
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)
	i4 = 4
	fmt.Println(i4)

	fmt.Println(i5)

	outer()

	// fmt.Println(out)
	// var in string = "in"
}
