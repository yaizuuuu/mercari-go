package main

import "fmt"

// IntelliJのプライングイン `File Watcher` で `go fmt` `goimports` を自動で実行する
// https://qiita.com/na_ga/items/288e34360edb3e05ca44
func main() {
	var price int
	fmt.Println("値段>")
	fmt.Scan(&price)

	fmt.Printf("%d円\n", price)
}

// go vet
// 組み込みの静的解析コマンド

// golint
// コーディングスタイルの問題を検出する

// errcheck
// エラーのチェックを行っているかを確認してくれる

// staticcheck
// サードパーティ版の `go vet`
// `go vet` より細かい

// golangci-lint
// `golint` の類似
// `golint` よりスター数が多いためこちらが主流になっていく可能性も

// reviewdog
// GitHubのコメントにLintで指摘された箇所にたいしてコメントをつけてくれる
// https://swet.dena.com/entry/2018/09/18/142413

// delve
// Go専用のデバッガ
// IntelliJやGoLandがあれば必要ないが、使い方を見てみるのはあり
