package function

import "fmt"

// 構造体の定義
type SiteData struct {
	Name  string
	URL   string
	Admin string
	Year  int
}

type SiteData2 struct {
	Name  string
	URL   string
	Admin string
	Year  int
}

func SubFunc12() {
	// 構造体
	fmt.Println("構造体")
	var site1 SiteData
	fmt.Println(site1)

	// フィールドを設定
	site1.URL = "https://xxx.com"
	site1.Name = "Golang HP"
	site1.Admin = "User"
	site1.Year = 2023
	fmt.Println(site1)
	// Nameだけ
	fmt.Println(site1.Name)

	// 変数宣言と同時に初期化
	site2 := SiteData2{
		"Yahoo!",
		"https://xxx.com",
		"Yahoo",
		1996,
	}
	fmt.Println(site2)
	fmt.Println(site2.Admin)
	// 値の更新
	site2.Admin = "ヤフー"
	fmt.Println(site2.Admin)
}
