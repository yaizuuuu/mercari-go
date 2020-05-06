package ch5

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var msg = flag.String("msg", "デフォルト値", "説明")
var nn int

func init() {
	flag.IntVar(&nn, "nn", 1, "回数")
}

func Exec() {
	fmt.Println(os.Args)

	for _, v := range os.Args {
		fmt.Println(v)
	}

	// ここで実際に設定される
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, nn))

	// os.Argsはフラグと引数の両方, flag.Argsは引数のみを取得できる
	fmt.Println(flag.Args())
}

var nb bool

func init() {
	flag.BoolVar(&nb, "n", false, "行数を表示する")
}

func Mycat() {
	flag.Parse()

	fmt.Println(nb)

	target := catTarget{
		fileNames:     flag.Args(),
		displayedLine: nb,
	}

	target.readFiles()
}

type catTarget struct {
	fileNames     []string
	displayedLine bool
	currentLine   int
}

func (target *catTarget) readFiles() {
	for _, fileName := range target.fileNames {
		target.readFile(fileName)
	}
}

func (target *catTarget) readFile(fileName string) {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal("ファイルが存在していません")
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		target.currentLine += 1

		if target.displayedLine {
			fmt.Printf("%d: %s\n", target.currentLine, scanner.Text())
		} else {
			fmt.Println(scanner.Text())
		}
	}

	defer f.Close()
}
