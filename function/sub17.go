package function

import "fmt"

type Player struct {
	Name  string
	Level int
}

// スライスを格納する型を作成（値渡し）
type Players []Player

func SubFunc17() {
	fmt.Println("SUB17")
	player1 := Player{Name: "AAA", Level: 10}
	player2 := Player{Name: "BBB", Level: 5}
	player3 := Player{Name: "CCC", Level: 20}

	// Players型の変数の宣言
	players := Players{}
	// 変数に構造体のデータをスライスに代入する
	players = append(players, player1, player2, player3)
	fmt.Println(players)
}
