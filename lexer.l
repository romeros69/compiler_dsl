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
    result AST
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
ARROW   "->"
COMMA   [\,]
CREATE  "create"
READ    "read"
UPDATE  "update"
DELETE  "delete"
LIST    "list"

%%
    buf.Reset()

{C}                 // continue

{LS}                {l.Backup(); return LS_TOK}

{RS}                {l.Backup(); return RS_TOK}

{COMMA}             {l.Backup(); return COMMA_TOK}

{ARROW}             {l.Backup(); return ARROW_TOK}

{CREATE}            {
                        l.Backup()
                        return CREATE_TOK
                    }

{READ}              {
                        l.Backup()
                        return READ_TOK
                    }

{UPDATE}            {
                        l.Backup()
                        return UPDATE_TOK
                    }

{DELETE}            {
                        l.Backup()
                        return DELETE_TOK
                    }

{LIST}              {
                        l.Backup()
                        return LIST_TOK
                    }

{INT}               {l.Backup(); return INT_T_TOK}

{STRING}            {l.Backup(); return STRING_T_TOK}

{E}                 {l.Backup(); return EN_TOK}

{S}                 {

                        lval.val = buf.String()
                        return IDENT
                    }

\0                  {fmt.Printf("ERROR 1 -- {%s}\n", buf.String()); return -1} // Exit on EOF or any other error

.                   {fmt.Printf("ERROR 2: %s, buf: %s\n", string(c), buf); return -1}

%%
    fmt.Println("ERROR 3")
    log.Fatal("scanner internal error")
    return 0
}