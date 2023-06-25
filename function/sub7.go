package function

import "fmt"

func SubFunc7() {
	// 配列
	var a [2]int
	// 配列に代入
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// 配列、宣言と代入を同時に
	var b [2]int = [2]int{10, 20}
	fmt.Println(b)
	fmt.Println(b[0])

	// 配列に入っている値の取り出し
	fmt.Println(a[0])
	fmt.Println(a[1])

	// 配列の値の更新
	a[0] = 500
	fmt.Println(a[0])

	// varなしで配列の宣言
	cc := [3]int{1, 2, 3}
	fmt.Println(cc)

	// 配列の型の確認
	fmt.Printf("ccの型は %T です。 \n", cc)
	// 配列の要素数の確認
	fmt.Println("要素数は、", len(cc), "です。")

	// len関数、データのbyte数を計測できる
	var en string = "Golang"
	var ja string = "Go言語"
	fmt.Println(en, "のbyte数は", len(en), "です。")
	fmt.Println(ja, "のbyte数は", len(ja), "です。")
}
