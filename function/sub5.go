package function

import "fmt"

func SubFunc5() {
	fmt.Println("Go")

	// for文
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// continue
	// 3のときはスキップ
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// break
	// 3のときに中断
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	//
	arr := [3]string{"Golang", "PHP", "C#"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// if文
	var a int = 2
	if a == 2 {
		fmt.Println("if文")
	}

	// if else
	if a == 2 {
		fmt.Println("aの値は2です")
	} else {
		fmt.Println("aの値は2以外です")
	}

	// if,else if,else
	var b int = 1
	if b == 2 {
		fmt.Println("22")
	} else if b == 1 {
		fmt.Println("11")
	} else {
		fmt.Println("00")
	}

	// switch文
	// 1か2のとき
	// 3か4のとき
	var num int = 3
	switch num {
	case 1, 2:
		fmt.Println("1か2です")
	case 3, 4:
		fmt.Println("3か4です")
	default:
		fmt.Println("どれでもありません")
	}

	// defer
	// 遅延処理
	fmt.Println("プログラムスタート！")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("deferの処理より前に実行される")
}
