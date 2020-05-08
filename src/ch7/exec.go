package ch7

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"unicode/utf8"
)

func Exec() {
	v := S("hoge")

	if s, err := ToString(v); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s.String())
	}
}

type Stringer interface {
	String() string
}

// errorインターフェースを継承したものを第2引数に返す
func ToString(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}

	// エラーだった場合
	return nil, MyError("CastError")
}

// `Error() string` を実装したカスタムのerror
type MyError string

func (e MyError) Error() string {
	return string(e)
}

type S string

func (s S) String() string {
	return string(s)
}

type RuneScanner struct {
	r   io.Reader
	buf [16]byte
}

func NewRuneScanner(r io.Reader) *RuneScanner {
	return &RuneScanner{r: r}
}

func (s *RuneScanner) Scan() (rune, error) {
	n, err := s.r.Read(s.buf[:])
	if err != nil {
		return 0, err
	}

	r, size := utf8.DecodeLastRune(s.buf[:])
	if r == utf8.RuneError {
		return 0, errors.New("RunError")
	}

	s.r = io.MultiReader(bytes.NewReader(s.buf[size:n]), s.r)

	return r, nil
}

func Exec2() {
	s := NewRuneScanner(strings.NewReader("Hello, 世界"))

	for {
		r, err := s.Scan()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(r)
	}
}
