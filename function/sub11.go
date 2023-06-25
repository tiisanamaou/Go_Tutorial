package function

import "fmt"

func SubFunc11() {
	// interface型
	// どんな型の値を入れられる
	// 演算はできない
	var a interface{}
	// 空っぽなので<nil>が帰ってくる
	fmt.Println(a)

	// int
	a = 10
	fmt.Println(a)
	// float
	a = 3.14
	fmt.Println(a)
	// string
	a = "Go"
	fmt.Println(a)
	// 配列
	a = [3]int{3, 2, 1}
	fmt.Println(a)

	// var省略
	b := interface{}("Golang")
	fmt.Println(b)

	// interfaceの型変換
	// 型アサーション
	var d11 interface{} = 13
	var dd int = d11.(int)
	fmt.Printf("a = %d\t%T\n", d11, d11)
	fmt.Printf("b = %d\t%T\n", dd, dd)

	// 型アサーションのエラー判定
	var bb interface{} = 100
	var fl64 float64

	fl64, isFloat := bb.(float64)
	fmt.Println(fl64, isFloat)

	// switch type でinterface型の型アサーション
	var x interface{} = 100
	switch x.(type) {
	case int:
		fmt.Println("変数はint型です。")
	case string:
		fmt.Println("変数はstring型です。")
	case float64:
		fmt.Println("変数はfloat64型です。")
	default:
		fmt.Println("型が不明です。")
	}
}
