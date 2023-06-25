package function

import "fmt"

func pTest(x *int) {
	*x = 0
}

func SubFunc13() {
	var a int = 1000
	fmt.Println(a)
	// メモリアドレス
	fmt.Println(&a)

	// ポインタ型
	// 変数aのアドレスをppに格納する
	var pp *int = &a
	fmt.Println(pp)
	// ポインタ変数のアドレスの中身
	fmt.Println(*pp)
	// ポインタ変数のアドレス
	fmt.Println(&pp)

	// new()関数でポインタを初期化
	var ppn *int = new(int)
	fmt.Printf("ppnのメモリアドレスは、 %pです。\n", ppn)
	fmt.Println(*ppn)

	// ポインタ、値渡し、参照渡し
	var b int = 100
	fmt.Println(b)
	// ポインタ
	var ppp *int = &b
	// 参照渡し（アドレス）なので変数bがx=0で上書きされる
	pTest(ppp)
	fmt.Println(b)
}
