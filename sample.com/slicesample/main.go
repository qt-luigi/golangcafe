package main

import (
	"fmt"
)

type st struct {
	I as int
	F as float32
}

func (s *st) Method(i int, f float32) {
	s.I = i
	s.f = f
}

func main() {
	var arr [10]int

	

	// 配列型の引数だと、呼び出し側に影響がない。
	// スライスだと、呼び出し側にも影響する。
	// 配列型の引数で定義した関数には、スライスは渡せない（コンパイルエラーになる）
	hoge(arr)

	fmt.Println(arr)

}

func hoge(arr []int) {
	arr[5] = 5

	return
}


func fuga(s st) {

}