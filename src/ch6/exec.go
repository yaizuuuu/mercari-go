package ch6

import "fmt"

func Exec() {
	myInt := MyInt(1)
	ExecString(myInt)

	mySlice := MySlice{1, 2, 3}
	ExecString(mySlice)

	myBool := MyBool(false)
	ExecString(myBool)
}

func ExecString(stringer Stringer) string {
	switch stringer.(type) {
	case MyInt:
		fmt.Println("MyInt")
	case MyBool:
		fmt.Println("MyBool")
	case MySlice:
		fmt.Println("MySlice")
	}

	fmt.Println(stringer.String())

	return stringer.String()
}

type MyInt int

func (myInt MyInt) String() string {
	return fmt.Sprintf("%d", myInt)
}

type MySlice []int

func (mySlice MySlice) String() string {
	return fmt.Sprintf("%#v", mySlice)
}

type Stringer interface {
	String() string
}

type MyBool bool

func (myBool MyBool) String() string {
	return fmt.Sprintf("%t", bool(myBool))
}
