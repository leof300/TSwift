grammar TSParser_rules;
import TSLexer_rules;

start : (   sent
        |   NL
        )* EOF
        ;

sent : expr (NL)+
     ;

expr :  |   expr '%' expr
        |   expr ('*'|'/')expr
        |   expr ('+'|'-') expr
        |   '(' expr ')'
        |   expr '=' expr
        |   VSTRING
        |   VINTEGER
        |   VFLOAT
        |   ID
        ;