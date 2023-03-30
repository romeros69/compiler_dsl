%{
package main

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
    yylex.(*Lex).result = $1
  }

ent: EN_TOK IDENT
  {
    $$ = ent{name: $2}
  }