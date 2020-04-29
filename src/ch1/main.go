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

//
