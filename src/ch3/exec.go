package ch3

import (
	"fmt"
	"net/http"
)

func Exec() {
	fixProblem()

	useSlice()
}

func fixProblem() {
	sum := float64(5 + 6 + 3)

	avg := sum / 3

	if avg > 4.5 {
		fmt.Println("Good")
	}
}

func useSlice() {
	//n1 := 19
	//n2 := 86
	//n3 := 1
	//n4 := 12
	//
	//sum := n1 + n2 + n3 + n4
	//fmt.Println(sum)

	ns := []int{19, 86, 1, 12}

	var sum int
	for _, v := range ns {
		sum += v
	}

	fmt.Println(sum)

	r := Result{
		point:  100,
		userId: 2,
		gameId: 5,
	}
	fmt.Printf("%#v\n", r)

	// 型のエイリアス
	type Applicant = http.Client
	// %Tで元の型名を表示させる
	fmt.Printf("%T\n", Applicant{})

	for i := 0; i <= 100; i++ {
		evenAndOdd(i)
	}

	x, y := 1, 2
	fmt.Println(x, y)
	x, y = swap(x, y)
	fmt.Println(x, y)

	m, n := 10, 20
	fmt.Println(m, n)
	swap2(&m, &n)
	fmt.Println(m, n)

	var myInt MyInt
	fmt.Println(myInt)
	myInt.Inc()
	fmt.Println(myInt)
}

type Result struct {
	point  int
	userId int
	gameId int
}

func evenAndOdd(i int) {
	fmt.Print(i)

	if i%2 == 0 {
		fmt.Println(" - 偶数")
	} else {
		fmt.Println(" - 奇数")
	}
}

func swap(x, y int) (int, int) {
	return y, x
}

func swap2(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

type MyInt int

func (n *MyInt) Inc() { *n++ }
