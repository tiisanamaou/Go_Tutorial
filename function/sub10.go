package function

import "fmt"

func SubFunc10() {
	// map
	var map1 = map[string]int{"Apple": 180, "Banana": 200, "Orange": 300}
	var map2 = map[string]int{
		"Apple":  180,
		"Banana": 200,
		"Orange": 300,
	}
	// varを省略
	map3 := map[string]int{"Go": 30, "C": 70, "PHP": 10}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	// make関数でmapを作成
	// 空のmap作成
	map4 := make(map[string]string)
	fmt.Println(map4)

	// キーを指定して値を取り出す
	// Appleに紐付いている値（180）を取り出す
	fmt.Println(map1["Apple"])

	// mapの値の更新
	map1["Apple"] = 2000
	fmt.Println(map1)

	//mapの値の追加
	map1["melon"] = 1000
	fmt.Println(map1)

	// mapの値の削除
	delete(map1, "Apple")
	fmt.Println(map1)

	// mapのエラー対応
	str, ok := map1["Apple"]
	if !ok {
		fmt.Println("値がありません")
	}
	fmt.Println(str, ok)
	//
	str2, ok2 := map1["melon"]
	if !ok2 {
		fmt.Println("値がありません")
	}
	fmt.Println(str2, ok2)
}
