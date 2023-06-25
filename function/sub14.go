package function

import "fmt"

// 構造体の定義
type Character struct {
	Name  string
	HP    int
	MP    int
	Money int
}

// 関数
// 値渡し
func upscore(player Character) {
	fmt.Println(player.Name, "のステータスがアップしました！")
	player.HP += 10
	player.MP += 20
	player.Money += 500
}

// 参照渡し
func upscore2(player *Character) {
	fmt.Println(player.Name, "のステータスがアップしました！")
	player.HP += 100
	player.MP += 200
	player.Money += 5000
}

func SubFunc14() {
	fmt.Println("SUB14")

	// new()関数を使用してポインタ定義
	player1 := new(Character)
	fmt.Println(player1)
	fmt.Println(*player1)

	// &を使用してポインタ定義
	player2 := &Character{}
	fmt.Println(player2)
	fmt.Println(*player2)

	// 使い方
	player3 := &Character{
		Name:  "まお",
		HP:    100,
		MP:    50,
		Money: 1000,
	}
	fmt.Println(player3)

	// 値渡しの関数
	// 値渡しなのでplayer3の値は変わらない
	upscore(*player3)
	fmt.Println(player3)

	// 参照渡しの関数
	// 参照渡しなのでplayer3の値が書き換わる
	upscore2(player3)
	fmt.Println(player3)
}
