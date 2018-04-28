# tokenizer
Tokenizer for Express written in Go

```program : stmt+ ; 

stmt    : expr <EOS> ;

expr    : term sec_op term |
          term             ;

sec_op  : ADD     |
          VEC_ADD |
          SUB     |
          VEC_SUB ;

ADD     : "+"  ;
VEC_ADD : ".+" ;
SUB     : "-"  ;
VEC_SUB : ".-" ;

term    : factor pri_op term |
          factor             ;

pri_op  : MULT     |
          DIV      |
          VEC_MULT |
          VEC_DIV  |
          MOD      ;

MULT      : "*"  ;
DIV       : "/"  ;
VEC_MULT  : ".*" ;
VEC_DIV   : "./" ;
MOD       : "%"  ;
VEC_MOD   : ".%" ;

factor  : L_PAREN expr R_PAREN |
          <literal>            |
          var                  ;

L_PAREN : "(" ;
R_PAREN : ")" ;

var     : type ident ;

type    : "var"  |
          number |
          char   |
          string |
          array  |
          bool   ;

ident   : <identifier> ;

(* 
  we need to make something for series/formulas
*)

number  : binary   |
          octal    |
          hex      |
          int      |
          float    |
          fraction |
          radical  |
          complex  ;

binary  : [ "0" | "1" ]+ ;
octal   : [ "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" ]+ ;
hex     : [ "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | 
            "8" | "9" |"A" | "B" | "C" | "D" | "E" | "F" ]+ ;
int     : [ "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9" ]+ ;
float   : []```
