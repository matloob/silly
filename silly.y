%{

package main

%}

%union {
  val int
  ident string
  node node
  list *stmtList
}

%token '{' '}' ';' '=' '+' '/' '-' '*' '(' ')' '#'

%token <val> VAL
%token <ident> IDENT

%type <node> expr stmt

%type <list> stmtList

%%

top: expr { r = $1 }

// ignoring order of operations... there will be s/r conflict & ambiguous output
expr: VAL { $$ = &intNode{$1} }
| expr '+' expr { $$ = &binopNode{"+", $1, $3} }
| expr '-' expr { $$ = &binopNode{"-", $1, $3} }
| expr '/' expr { $$ = &binopNode{"/", $1, $3} }
| expr '*' expr { $$ = &binopNode{"*", $1, $3} }
| IDENT '=' expr { $$ = &assignNode{$1, $3} }
| IDENT { $$ = &identNode{$1} }

| '{' stmtList '}' { $$ = $2; }

stmtList:
{ $$ = nil }
| stmt stmtList { $$ = &stmtList{$1, $2} }

stmt:
expr ';' { $$ = $1 }


%%
