package ch4

import (
	"fmt"
	"time"

	"github.com/tenntenn/greeting"
	greeting2 "github.com/tenntenn/greeting/v2"
)

func Exec() {
	fmt.Println(greeting.Do())
	fmt.Println(greeting2.Do(time.Now()))
}
