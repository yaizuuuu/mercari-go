package ch9

import (
	"bufio"
	"fmt"
	"io"
	"sync"
	"time"
)

func Exec() {
	//defer fmt.Println("main done")
	//
	//go func() {
	//	defer fmt.Println("goroutine 1 done")
	//
	//	time.Sleep(3 * time.Second)
	//}()
	//
	//go func() {
	//	defer fmt.Println("goroutine 2 done")
	//
	//	time.Sleep(1 * time.Second)
	//}()
	//
	//time.Sleep(1 * time.Second)
	//
	//ch := input(os.Stdin)
	//for {
	//	receive, ok := <-ch
	//
	//	if !ok {
	//		break
	//	}
	//
	//	fmt.Print(">")
	//	fmt.Println(receive)
	//}

	var m sync.Mutex
	m.Lock()

	go func() {
		time.Sleep(3 * time.Second)
		m.Unlock()
		fmt.Println("unlock 1")
	}()

	m.Lock()
	m.Unlock()
	fmt.Println("unlock 2")
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)

		for s.Scan() {
			if s.Text() == "" {
				break
			}

			ch <- s.Text()
		}

		close(ch)
	}()

	return ch
}
