%{
package main

import (
    "log"
    "bytes"
    "errors"
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
        l.pos = -1
        return 0
    }
    buf.WriteByte(l.input[l.pos])
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

S   ^[_a-zA-Z][_a-zA-Z0-9]*
E entity

%%
    buf.Reset()  // Code before the first rule is executed before every scan cycle (state 0 action)

[ \t\n\r]+          // Ignore whitespace

{E}                 {return EN_TOK}

{S}                 {lval.val = buf.String(); return IDENT}

\0                  {return -1} // Exit on EOF or any other error

.                   {return -1}

%%
    log.Fatal("scanner internal error")
    return 0

} // *1 this right brace