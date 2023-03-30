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
  fields []field
}

%}

%union{
  ent ent
  fields []field
  field field
  val string
  tokType int
}

%token <val>	 IDENT EN_TOK INT_T_TOK STRING_T_TOK LS RS
%type  <ent>     ent
%type  <field>   field
%type  <tokType> type
%type  <fields>	 fields

%start main

%%

main: ent
  {
    yylex.(*Lex).result = $1
  }

ent: EN_TOK IDENT LS fields RS
  {
    $$ = ent{name: $2, fields: $4}
  }

fields:
  field
  {
    $$ = []field{$1}
  }
| fields field
  {
    $$ = append($1, $2)
  }

field: IDENT type
  {
    $$ = field{name: $1, tokType: $2}
  }
type:
  INT_T_TOK {$$ = INT}
| STRING_T_TOK {$$ = STRING}