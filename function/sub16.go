package function

import "fmt"

// 構造体の中に構造体を埋め込む
type Monster struct {
	Name      string
	Attribute string
}
type MonsterData struct {
	Monster Monster
	Hp      int
	Mp      int
	attack  int
	defense int
}

func SubFunc16() {
	fmt.Println("SUB16")

	// 構造体の初期化
	slimeData := MonsterData{
		Monster: Monster{
			Name:      "スライム",
			Attribute: "スライム属",
		},
		Hp:      50,
		Mp:      10,
		attack:  10,
		defense: 10,
	}
	fmt.Println(slimeData)

	// 埋め込んだ構造体の値を出力する
	fmt.Println(slimeData.Monster)

	// Nameの値を出力する
	fmt.Println(slimeData.Monster.Name)
}
