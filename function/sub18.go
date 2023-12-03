package function

import "fmt"

type Poniter struct {
	num int
}

func SubFunc18() {
	p := "ポインタ"
	fmt.Println(p)

	// アドレスが出力される
	var pointer *string = &p
	fmt.Println(pointer)

	// ポインタの中身が出力される
	fmt.Println(*pointer)

	//
	var pp Poniter
	pp.PoniterA(0)
	fmt.Println(pp.num)
	pp.PoniterB(0)
	fmt.Println(pp.num)
}

// func (変数名 構造体) 関数名(変数名 型) 型{}
func (p Poniter) PoniterA(n int) {
	p.num = n + 3
}

func (p *Poniter) PoniterB(n int) {
	p.num = n + 3
}
