%{
package main

import (
    "log"
    "bytes"
    "errors"
    "fmt"
)

type Lex struct {
    input []byte
    pos int
    result ent
    err error
}

var buf = bytes.NewBuffer(nil)

func NewLex(input []byte) *Lex {
    return &Lex{input: input}
}

func (l *Lex) Next() byte {
    if l.pos >= len(l.input) || l.pos == -1 {
        if l.pos == len(l.input) {
            buf.WriteByte(l.input[l.pos - 1])
        }
        l.pos = -1
        return 0
    }
    if l.pos != 0 {
        buf.WriteByte(l.input[l.pos - 1])
    }
    l.pos++
    return l.input[l.pos - 1]
}

func (l *Lex) Error(s string) {
    l.err = errors.New(s)
}

func (l *Lex) Backup() {
    if l.pos == -1 {
        return
    }
    l.pos--
}

func (l *Lex) Lex(lval *yySymType) int {
    c := l.Next() // init
%}

%yyc c
%yyn c = l.Next()

S       [a-zA-Z]+
E       entity
INT     int
STRING  string
C       [ \s\t\n\r]+
LS      "{"
RS      "}"

%%
    buf.Reset()

{C}                 // continue

{LS}                {return LS}

{RS}                {return RS}

{INT}               {return INT_T_TOK}

{STRING}            {return STRING_T_TOK}

{E}                 {return EN_TOK}

{S}                 {
                        lval.val = buf.String()
                        return IDENT
                    }

\0                  {return -1} // Exit on EOF or any other error

.                   {fmt.Printf("ERROR 2: %s, buf: %s\n", string(c), buf); return -1}

%%
    fmt.Println("ERROR 3")
    log.Fatal("scanner internal error")
    return 0
}