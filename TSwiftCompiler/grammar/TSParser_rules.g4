grammar TSParser_rules;
import TSLexer_rules;

start : lsents;

lsents : sents+ EOF
        ;

sents : expr NL         #SentExpr
      |   NL            #SentNL
      | declar          #SDecl
      | declar op='=' expr #SDeclAsig

     ;

expr :      expr op='%' expr        #EModule
        |   expr op=('*'|'/')expr   #EMulDiv
        |   expr op=('+'|'-') expr  #EAddSub
        |   '(' expr ')'            #EParent
        |   expr op='=' expr        #EAssign
        |   expr op='+=' expr       #EAsAdd
        |   expr op='-=' expr       #ESubAdd
        |   <assoc=right> '-' expr  #ENeg
        |   VSTRING                 #EVString
        |   VINTEGER                #EVInteger
        |   VFLOAT                  #EVFloat
        |   VBOOL                   #EVBOOL
        |   ID                      #EID
        ;


declar : 'var' ID ':' STRING    #SDStr
        |'var' ID ':' INT       #SDInt
        |'var' ID ':' FLOAT     #SDFlt
        |'var' ID ':' BOOL      #SDBool
        |'var' ID ':' CHARACTER #SDChr
       ;