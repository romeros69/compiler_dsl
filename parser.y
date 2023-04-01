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
  actions []string
}

%}

%union{
  ent ent
  fields []field
  field field
  val string
  tokType int
  actions []string
}

%token <val>	 IDENT EN_TOK INT_T_TOK STRING_T_TOK LS_TOK RS_TOK ARROW_TOK COMMA_TOK CREATE_TOK READ_TOK UPDATE_TOK DELETE_TOK LIST_TOK
%type  <ent>     ent
%type  <field>   field
%type  <tokType> type
%type  <fields>	 fields
%type  <actions> actions
%type  <val>  	 action

%start main

%%

main: ent
  {
    yylex.(*Lex).result = $1
  }

ent: EN_TOK IDENT LS_TOK fields RS_TOK ARROW_TOK actions
  {
    $$ = ent{
    	name: $2,
    	fields: $4,
    	actions: $7,
    }
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

actions:
  action
  {
    $$ = []string{$1}
  }
| actions COMMA_TOK action
  {
    $$ = append($1, $3)
  }

action:
  CREATE_TOK 	{$$ = "create"}
| READ_TOK 	{$$ = "read"}
| UPDATE_TOK 	{$$ = "update"}
| DELETE_TOK 	{$$ = "delete"}
| LIST_TOK 	{$$ = "list"}


type:
  INT_T_TOK 	{$$ = INT}
| STRING_T_TOK 	{$$ = STRING}