%{
package main

const (
  INT = iota
  STRING
)

type field struct {
  tokType int
  name string
}

type ent struct {
  name string
  field field
}

%}

%union{
  ent ent
  field field
  val string
  tokType int
}

%token <val>     IDENT EN_TOK INT_T_TOK STRING_T_TOK LS RS
%type  <ent>     ent
%type  <field>   field
%type  <tokType> type

%start main

%%

main: ent
  {
    yylex.(*Lex).result = $1
  }

ent: EN_TOK IDENT LS field RS
  {
    $$ = ent{name: $2, field: $4}
  }

field: IDENT type
  {
    $$ = field{name: $1, tokType: $2}
  }
type:
  INT_T_TOK {$$ = INT}
| STRING_T_TOK {$$ = STRING}