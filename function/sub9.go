package function

import "fmt"

func SubFunc9() {
	// 基本型の変数のコピー
	var a9 int = 10
	a92 := a9
	fmt.Println(a9)
	fmt.Println(a92)
	a92 = 20
	fmt.Println(a9)
	fmt.Println(a92)

	// int型の配列をコピー
	// 値渡し
	var b9 [3]int = [3]int{1, 2, 3}
	b92 := b9
	fmt.Println(b9)
	fmt.Println(b92)
	b92[2] = 1000
	fmt.Println(b9)
	fmt.Println(b92)

	// スライスをコピー
	// 参照渡し（アドレスのコピーをしている）
	sl9 := []int{10, 20, 30}
	sl92 := sl9
	sl92[2] = 100000
	fmt.Println(sl9)
	fmt.Println(sl92)
	// アドレスを確認
	fmt.Printf("sl9  = %p\n", sl9)
	fmt.Printf("sl92 = %p\n", sl92)

	// copy関数
	// スライス宣言
	sl93 := []int{1, 2, 3, 4, 5}
	// スライスのメモリ領域の確保
	sl94 := make([]int, 5)
	fmt.Println(sl94)
	// copy関数、sl93の値をsl94にコピーする
	c := copy(sl94, sl93)
	fmt.Println(c, sl94)
	// アドレスの確認
	fmt.Printf("sl93 = %p\n", sl93)
	fmt.Printf("sl94 = %p\n", sl94)
	// sl94[0]の値を変更
	sl94[0] = 90
	fmt.Println(sl93)
	fmt.Println(sl94)
}
