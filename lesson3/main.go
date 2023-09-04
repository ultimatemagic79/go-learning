package main

// 基本的な型について

import (
	"fmt"
	"strconv"
)

// 定数
const Pi = 3.14 // 頭文字が大文字なので外部からアクセス可能

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	c0 = iota
	c1
	c2
)

func main() {
	var i int = 100
	/*
		最大値 最小値
		int8 -128 127
		int16 -32768 32767
		int32 -2147483648 2147483647
		int64 -9223372036854775808 9223372036854775807
	*/

	// intは環境によって変わる
	var i2 int64 = 200
	fmt.Println(i + 50)

	fmt.Println(i2 + 50)

	fmt.Printf("%T\n", i2) // %Tは型を表示するフォーマット
	fmt.Printf("%T\n", int8(i2))

	// 浮動小数点数
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)
	fmt.Printf("%T, %T\n", fl64, float32(fl))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	// uintは符号なし整数
	// complexは複素数

	// 論理値
	var t, f bool = true, false
	fmt.Println(t, f)

	// 文字列
	var s string = "Hello Golang"
	fmt.Println(s)

	fmt.Printf("%T\n", s)

	fmt.Println(`test
	test
	test`)

	fmt.Println(s[0])
	fmt.Println(string(s[0]))

	// 文字列の結合
	fmt.Println(string(s[0]) + string(s[1]))

	// byte
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// 配列, サイズの変更不可
	var arr1 [3]int
	fmt.Println(arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"c", "d"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr4[0])

	arr2[2] = "C"
	fmt.Println(arr2)

	// 配列の長さ
	fmt.Println(len(arr2))

	// interface
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)
	// 計算などは出来ない

	// 型変換
	var i3 int = 300
	fl64 = float64(i3)
	fmt.Printf("fl64 = %T\n", fl64)

	var s2 string = "test"
	fmt.Printf("s2 = %T\n", s2)

	// string convert int
	// strconvで変換
	ii, _ := strconv.Atoi("-42")
	/*
		if err != nil {
			fmt.Println(err)
		}
	*/
	fmt.Println(ii)
	fmt.Printf("%T\n", i)

	byte := []byte(s)
	fmt.Println(byte)
	fmt.Printf("byte = %T\n", byte)

	g2 := string(byte)
	fmt.Println(g2)

	fmt.Println(Pi)
	fmt.Println(c0, c1, c2)

}
