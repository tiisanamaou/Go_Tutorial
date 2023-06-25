package function

import "fmt"

func SubFunc4() {
	fmt.Println("Go")

	TESTFunction1()
	TESTFunction2(10, 25)

	// i に return の値（戻り値）を代入している
	i := TESTFunction3(20, 33)
	fmt.Println(i)

	// 複数の引数、戻り値
	dd, rr := TESTFunction4(10, 5, 2)
	fmt.Printf("xとzを足した値は %d です \n", dd)
	fmt.Printf("yとzを足した値は %d　です \n", rr)

	//
	ii := TESTFunction5(90, 10)
	fmt.Printf("xとyをかけた値は %d です \n", ii)
}

// 関数
// 引数なし、戻り値なし
func TESTFunction1() {
	fmt.Println("引数と戻り値がない関数")
}

// 引数あり、戻り値なし
func TESTFunction2(x int, y int) {
	fmt.Println(x + y)
}

// 引数あり、戻り値あり
func TESTFunction3(x int, y int) int {
	return x + y
}

// 複数の、引数、戻り値の場合
func TESTFunction4(x int, y int, z int) (int, int) {
	var d int = x + z
	var r int = y + z
	return d, r
}

// 戻り値の変数宣言
func TESTFunction5(x int, y int) (z int) {
	z = x * y
	return
}
