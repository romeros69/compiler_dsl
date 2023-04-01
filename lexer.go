// Code generated by golex. DO NOT EDIT.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

type Lex struct {
	input  []byte
	pos    int
	result AST
	err    error
}

var buf = bytes.NewBuffer(nil)

func NewLex(input []byte) *Lex {
	return &Lex{input: input}
}

func (l *Lex) Next() byte {
	if l.pos >= len(l.input) || l.pos == -1 {
		if l.pos == len(l.input) {
			buf.WriteByte(l.input[l.pos-1])
		}
		l.pos = -1
		return 0
	}
	if l.pos != 0 {
		buf.WriteByte(l.input[l.pos-1])
	}
	l.pos++
	return l.input[l.pos-1]
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

yystate0:

	buf.Reset()

	goto yystart1

yystate1:
	c = l.Next()
yystart1:
	switch {
	default:
		goto yystate3 // c >= '\x01' && c <= '\b' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\x1f' || c >= '!' && c <= '+' || c >= '.' && c <= '@' || c >= '[' && c <= '`' || c == '|' || c >= '~' && c <= 'ÿ'
	case c == ',':
		goto yystate6
	case c == '-':
		goto yystate7
	case c == '\n':
		goto yystate5
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate4
	case c == '\x00':
		goto yystate2
	case c == 'c':
		goto yystate11
	case c == 'd':
		goto yystate17
	case c == 'e':
		goto yystate23
	case c == 'i':
		goto yystate29
	case c == 'l':
		goto yystate32
	case c == 'r':
		goto yystate36
	case c == 's':
		goto yystate40
	case c == 'u':
		goto yystate47
	case c == '{':
		goto yystate53
	case c == '}':
		goto yystate54
	case c >= 'A' && c <= 'Z' || c == 'a' || c == 'b' || c >= 'f' && c <= 'h' || c == 'j' || c == 'k' || c >= 'm' && c <= 'q' || c == 't' || c >= 'v' && c <= 'z':
		goto yystate9
	}

yystate2:
	c = l.Next()
	goto yyrule15

yystate3:
	c = l.Next()
	goto yyrule16

yystate4:
	c = l.Next()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ' || c == 's':
		goto yystate5
	}

yystate5:
	c = l.Next()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ' || c == 's':
		goto yystate5
	}

yystate6:
	c = l.Next()
	goto yyrule4

yystate7:
	c = l.Next()
	switch {
	default:
		goto yyrule16
	case c == '>':
		goto yystate8
	}

yystate8:
	c = l.Next()
	goto yyrule5

yystate9:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate10:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate11:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'r':
		goto yystate12
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate10
	}

yystate12:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'e':
		goto yystate13
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate13:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'a':
		goto yystate14
	case c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate14:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 't':
		goto yystate15
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate15:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'e':
		goto yystate16
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate16:
	c = l.Next()
	switch {
	default:
		goto yyrule6
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate17:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'e':
		goto yystate18
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate18:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'l':
		goto yystate19
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate10
	}

yystate19:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'e':
		goto yystate20
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate20:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 't':
		goto yystate21
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate21:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'e':
		goto yystate22
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate22:
	c = l.Next()
	switch {
	default:
		goto yyrule9
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate23:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'n':
		goto yystate24
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yystate24:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 't':
		goto yystate25
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate25:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'i':
		goto yystate26
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate10
	}

yystate26:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 't':
		goto yystate27
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate27:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'y':
		goto yystate28
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate10
	}

yystate28:
	c = l.Next()
	switch {
	default:
		goto yyrule13
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate29:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'n':
		goto yystate30
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yystate30:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 't':
		goto yystate31
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate31:
	c = l.Next()
	switch {
	default:
		goto yyrule11
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate32:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'i':
		goto yystate33
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate10
	}

yystate33:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 's':
		goto yystate34
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate10
	}

yystate34:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 't':
		goto yystate35
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate35:
	c = l.Next()
	switch {
	default:
		goto yyrule10
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate36:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'e':
		goto yystate37
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate37:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'a':
		goto yystate38
	case c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate38:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'd':
		goto yystate39
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate10
	}

yystate39:
	c = l.Next()
	switch {
	default:
		goto yyrule7
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate40:
	c = l.Next()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	case c == 's':
		goto yystate41
	case c == 't':
		goto yystate42
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'r' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate41:
	c = l.Next()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	case c == 's':
		goto yystate41
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate10
	}

yystate42:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'r':
		goto yystate43
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate10
	}

yystate43:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'i':
		goto yystate44
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate10
	}

yystate44:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'n':
		goto yystate45
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yystate45:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'g':
		goto yystate46
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate10
	}

yystate46:
	c = l.Next()
	switch {
	default:
		goto yyrule12
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate47:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'p':
		goto yystate48
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate10
	}

yystate48:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'd':
		goto yystate49
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate10
	}

yystate49:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'a':
		goto yystate50
	case c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate50:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 't':
		goto yystate51
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate51:
	c = l.Next()
	switch {
	default:
		goto yyrule14
	case c == 'e':
		goto yystate52
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate52:
	c = l.Next()
	switch {
	default:
		goto yyrule8
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate53:
	c = l.Next()
	goto yyrule2

yystate54:
	c = l.Next()
	goto yyrule3

yyrule1: // {C}
	{
		// continue
		goto yystate0
	}
yyrule2: // {LS}
	{
		{
			l.Backup()
			return LS_TOK
		}
		goto yystate0
	}
yyrule3: // {RS}
	{
		{
			l.Backup()
			return RS_TOK
		}
		goto yystate0
	}
yyrule4: // {COMMA}
	{
		{
			l.Backup()
			return COMMA_TOK
		}
		goto yystate0
	}
yyrule5: // {ARROW}
	{
		{
			l.Backup()
			return ARROW_TOK
		}
		goto yystate0
	}
yyrule6: // {CREATE}
	{
		{
			l.Backup()
			return CREATE_TOK
		}
		goto yystate0
	}
yyrule7: // {READ}
	{
		{
			l.Backup()
			return READ_TOK
		}
		goto yystate0
	}
yyrule8: // {UPDATE}
	{
		{
			l.Backup()
			return UPDATE_TOK
		}
		goto yystate0
	}
yyrule9: // {DELETE}
	{
		{
			l.Backup()
			return DELETE_TOK
		}
		goto yystate0
	}
yyrule10: // {LIST}
	{
		{
			l.Backup()
			return LIST_TOK
		}
		goto yystate0
	}
yyrule11: // {INT}
	{
		{
			l.Backup()
			return INT_T_TOK
		}
		goto yystate0
	}
yyrule12: // {STRING}
	{
		{
			l.Backup()
			return STRING_T_TOK
		}
		goto yystate0
	}
yyrule13: // {E}
	{
		{
			l.Backup()
			return EN_TOK
		}
		goto yystate0
	}
yyrule14: // {S}
	{
		{
			lval.val = buf.String()
			return IDENT
		}
		goto yystate0
	}
yyrule15: // \0
	{
		{
			fmt.Printf("ERROR 1 -- {%s}\n", buf.String())
			return -1
		} // Exit on EOF or any other error
		goto yystate0
	}
yyrule16: // .
	if true { // avoid go vet determining the below panic will not be reached
		{
			fmt.Printf("ERROR 2: %s, buf: %s\n", string(c), buf)
			return -1
		}
		goto yystate0
	}
	panic("unreachable")

yyabort: // no lexem recognized
	// silence unused label errors for build and satisfy go vet reachability analysis
	{
		if false {
			goto yyabort
		}
		if false {
			goto yystate0
		}
		if false {
			goto yystate1
		}
	}

	fmt.Println("ERROR 3")
	log.Fatal("scanner internal error")
	return 0
}
