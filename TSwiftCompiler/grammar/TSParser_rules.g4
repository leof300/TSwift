grammar TSParser_rules;
import TSLexer_rules;

start : lsents;

lsents : sents+ EOF
        ;

sents : expr NL #SentExpr
      |   NL    #SentNL
     ;

expr :      expr '%' expr           #EModule
        |   expr ('*'|'/')expr      #EMulDiv
        |   expr ('+'|'-') expr     #EAddSub
        |   '(' expr ')'            #EParent
        |   expr '=' expr           #EAssign
        |   VSTRING                 #EVString
        |   VINTEGER                #EVInteger
        |   VFLOAT                  #EVFloat
        |   VBOOL                   #EVBOOL
        |   ID                      #EID

        ;