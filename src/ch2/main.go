package ch2

import (
	"fmt"
	"math/rand"
	"time"
)

func Exec() {
	useVariable()

	useConst()

	tryOperator()

	evenAndOdd()

	fortune()
}

func useVariable() {
	msg := "Hello, 世界"

	fmt.Println("変数", msg)
}

func useConst() {
	const msg = "Hello, 世界"

	fmt.Println("定数", msg)

	const (
		a = 1 + 2
		b
		c
	)

	fmt.Println("定数で右辺を省略する", a, b, c)

	const (
		d = iota
		e
	)

	fmt.Println("iotaで連続した値を作る", d, e)
}

func tryOperator() {
	n := 100 + 200
	m := n + 100

	msg := "hoge" + "fuga"

	fmt.Println("演算子を試す", n, m, msg)
}

func evenAndOdd() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("[if文]%d - 偶数\n", i)
		} else {
			fmt.Printf("[if文]%d - 奇数\n", i)
		}

		switch i % 2 {
		case 0:
			fmt.Printf("[switch文]%d - 偶数\n", i)
		default:
			fmt.Printf("[switch文]%d - 奇数\n", i)
		}
	}
}

func fortune() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6) + 1

	fmt.Println(s)

	switch s {
	case 1:
		fmt.Println("凶")
	case 2, 3:
		fmt.Println("吉")
	case 4, 5:
		fmt.Println("中吉")
	default:
		fmt.Println("大吉")
	}
}
