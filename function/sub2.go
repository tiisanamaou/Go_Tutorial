package function

import "fmt"

// 関数：func
func func1(xx int) string {
	return fmt.Sprint("Function ", xx, " Golang")
}

func SubFunc2() {
	// 変数
	var number int = 10
	var (
		num1 int = 100
		num2 int = 200
	)

	//配列
	var array1 = [3]string{}
	array1[0] = "Red"
	array1[1] = "Green"
	array1[2] = "Blue"

	//スライス
	var slice1 = []string{}
	slice1 = append(slice1, "red")
	slice1 = append(slice1, "green")
	slice1 = append(slice1, "blue")

	// IF文
	var x int = 10
	var y int = 5
	if x > y {
		fmt.Println("x is ", x)
	}
	if x < y {
		fmt.Println("000")
	} else if x == y {
		fmt.Println("111")
	} else {
		fmt.Println(x, y)
	}

	//Switch文
	var mode string = "running"
	switch mode {
	case "running":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("不明")
	}

	//For文
	var a1 int = 0
	var a2 int = 3
	for a1 < a2 {
		a1++
		fmt.Println(a1)
	}

	// 表示
	fmt.Println("Hello")
	fmt.Println(number)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(array1[0], array1[1], array1[2])
	fmt.Println(slice1[0], slice1[1], slice1[2])
	fmt.Println(func1(0000))
	fmt.Println(func1(1234))
}
