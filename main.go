package main

import "fmt"

func main() {
	l := NewLex([]byte("entity popa"))
	_ = yyParse(l)
	fmt.Printf("ast: %v\n", l.result)
}
