package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.fekos")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error in close file, continue...")
		}
	}(file)

	inputBuf := make([]byte, func(f *os.File) int64 {
		stat, _ := f.Stat()
		return stat.Size()
	}(file))

	_, err = file.Read(inputBuf)
	if err == io.EOF {
		fmt.Println("end, ok")
	}

	l := NewLex(inputBuf)
	_ = yyParse(l)
	fmt.Printf("ast: %#+v\n", l.result)
}
