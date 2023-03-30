%{
package main

import "fmt"

type ent struct {
  name string
}

%}

%union{
  ent ent
  val string
}

%token <val> IDENT EN_TOK
%type <ent> ent

%start main

%%

main: ent
  {
    fmt.Println("11111")
    yylex.(*Lex).result = $1
  }

ent: EN_TOK IDENT
  {
    fmt.Println("22222")
    $$ = ent{name: $2}
  }