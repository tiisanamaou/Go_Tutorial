package function

import "fmt"

// 構造体の定義
type Character2 struct {
	Name  string
	Job   string
	Level int
	HP    int
	MP    int
	Money int
}

// メソッドの定義
func (a Character2) LevelUpMsg() {
	fmt.Println(a.Name, "はレベルアップした！")
}

// ポインタレシーバー
func (a *Character2) LevelUpCal() {
	fmt.Println(a.Name, "はレベルアップした！")
	a.Level++
	a.HP += 100
	a.MP += 100
	a.Money += 1500
	fmt.Println(a.Name, "はレベルが", a.Level, "HPが", a.HP, "MPが", a.MP, "所持金が", a.Money, "にアップした！")
}

// メソッドを定義（引数を渡す）
func (b *Character2) JobChange(att string) {
	b.Job = att
	fmt.Println(b.Name, "は", b.Job, "に進化した。")
}

func SubFunc15() {
	fmt.Println("SUB15")

	player1 := Character2{
		Name:  "まお",
		Level: 1,
		HP:    100,
		MP:    100,
		Money: 5000,
	}

	// メソッドを呼び出し
	player1.LevelUpMsg()

	// ポインタレシーバー
	player2 := Character2{
		Name:  "ゆうしゃ",
		Job:   "勇者",
		Level: 10,
		HP:    1000,
		MP:    1000,
		Money: 100000,
	}
	// 初期値の確認
	fmt.Println(player2)
	// メソッドの呼び出し
	player2.LevelUpCal()
	//メソッドの呼び出し
	player2.JobChange("闇落ち勇者")
}
