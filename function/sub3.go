package function

import "fmt"

func SubFunc3() {
	fmt.Println("Sub3")

	// int
	// 変数の宣言
	var num3 int = 1
	var num4 int64 = 10

	// 値をコマンドラインに出力
	fmt.Println(num3)
	fmt.Println(num4)

	// 型を確認
	fmt.Printf("%T\n", num3)
	fmt.Printf("%T\n", num4)

	// 型変換
	fmt.Println(num3 + int(num4))

	// float
	// 変数の宣言
	var num5 float64 = 1.5
	var num6 float32 = 1.6
	// 型変換
	fmt.Println(num5 + float64(num6))

	// string
	// 変数の宣言
	var string1 string = "Golang_000_日本語"
	fmt.Println(string1)

	// 型の確認
	fmt.Printf("%T\n", string1)

	// stringとbyte型
	fmt.Println(string1[0])
	// アスキーコードを文字に変換
	fmt.Println(string(string1[0]))

	// byte
	// 変数
	b1 := []byte{72, 73}
	b2 := []byte("Go")
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(string(b1))
	fmt.Println(string(b2))

	// bool
	// 変数宣言
	var tt bool = true
	var ff bool = false

	// 値を出力
	fmt.Println(tt, ff)
}
