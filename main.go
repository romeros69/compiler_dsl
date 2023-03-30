package main

import "fmt"

func main() {
	l := NewLex([]byte("entity popa"))
	_ = yyParse(l)
	fmt.Println(l.result)
}
