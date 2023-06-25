package function

import "fmt"

func SubFunc8() {
	// スライス
	var sliceNum []int
	fmt.Println(sliceNum)

	var sliceNum2 []int = []int{30, 20, 10}
	fmt.Println(sliceNum2)
	fmt.Println(sliceNum2[0])

	// 配列の中身を変更する
	sliceNum2[0] = 400
	fmt.Println(sliceNum2[0])

	// スライスを省略形で宣言
	sliceNum3 := []string{"Gooo", "Go言語", "12345"}
	fmt.Println(sliceNum3)

	// 多次元スライス、リスト
	var sliceNum4 = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(sliceNum4)

	// append関数でスライスの要素を追加する
	sliceNum3 = append(sliceNum3, "追加した要素")
	fmt.Println(sliceNum3)

	// スライスの値をfor文で取得する
	for i := 0; i < len(sliceNum3); i++ {
		fmt.Println(sliceNum3[i])
	}

	// make関数でうスライスのメモリを割り当てる
	sliceNum5 := make([]int, 7)
	fmt.Println(sliceNum5)
	sliceNum5[2] = 1000
	fmt.Println(sliceNum5)

	// 要素数とメモリ容量を計測する
	// len関数、cap関数
	fmt.Println("要素数は", len(sliceNum5), "です。")
	fmt.Println("メモリ容量は", cap(sliceNum5), "です。")
}
