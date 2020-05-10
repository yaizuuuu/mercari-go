package ch9

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

var words = []string{
	"apple",
	"banana",
	"lemon",
}

type Problem struct {
	No     int
	Answer string
	Input  string
}

func (p *Problem) Judge() bool {
	return p.Answer == p.Input
}

type Problems []Problem

func Game() {
	wg := new(sync.WaitGroup)

	isNext := make(chan bool, len(words))
	problems := make(chan Problem, len(words))
	displayedProblems := make(chan Problem, len(words))
	answered := make(Problems, 0, 3)

	bc := context.Background()
	t := 5 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

	wg.Add(1)
	go func() {
		n := 0
		for {
			select {
			case <-isNext:
				if n >= len(words) {
					close(problems)
					close(isNext)
					wg.Done()
					return
				}

				problems <- Problem{
					No:     n + 1,
					Answer: words[n],
				}

				n += 1
			case <-ctx.Done():
				close(problems)
				close(isNext)
				wg.Done()
				fmt.Println("timeUp!")
				return
			}
		}
	}()

	go func() {
		for {
			problem, ok := <-problems

			if !ok {
				return
			}

			fmt.Println(problem.Answer)
			fmt.Print("> ")

			displayedProblems <- problem
		}
	}()

	go func() {
		for {
			problem, ok := <-displayedProblems
			if !ok {
				return
			}

			s := bufio.NewScanner(os.Stdin)
			for s.Scan() {
				problem.Input = s.Text()
				answered = append(answered, problem)
				isNext <- true
				break
			}
		}
	}()

	isNext <- true

	wg.Wait()

	fmt.Println("")
	fmt.Println("=============================================")
	fmt.Println("result")
	fmt.Println("=============================================")

	total := float64(len(words))
	corrected := 0
	for _, v := range answered {
		fmt.Printf("No: %d, IsCorrect: %t, Correct Answer: %s, Your Answer: %s\n", v.No, v.Judge(), v.Answer, v.Input)
		if v.Judge() {
			corrected += 1
		}
	}

	fmt.Printf("Your Score: %.2f%%\n", float64(corrected)/total*100)
}
