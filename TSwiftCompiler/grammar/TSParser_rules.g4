grammar TSParser_rules;
import TSLexer_rules;

start : lsents EOF;

lsents : sents*
        ;


sents : if                      #SIf
        | switch                #SSwitch
        | declar                #SDecl
        | declar op='=' expr    #SDeclAsig
        | print                 #SPrint
        //| NL                    #SentNL
        | expr                #SentExpr
     ;

expr :      <assoc=right> op='-' expr  #ENeg
        |   <assoc=right> op='!' expr  #ENot
        |   expr op='%' expr        #EModule
        |   expr op=('*'|'/')expr   #EMulDiv
        |   expr op=('+'|'-') expr  #EAddSub
        |   expr op=('=='|'!='|'>'|'<'|'>='|'<=') expr  #ERel
        |   '(' expr ')'            #EParent
        |   expr op='&&' expr       #ERelAnd
        |   expr op='||' expr       #ERelOr
        |   <assoc=right> expr op='=' expr        #EAssign
        |   <assoc=right> expr op='+=' expr       #EAsAdd
        |   <assoc=right> expr op='-=' expr       #ESubAdd
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

block : term='{' lsents '}';

if :  IF  expr block            #RIf
    | IF  expr block ELSE block #RIfElse
    | IF  expr block ELSE if    #RIfEIf
;

switch: SWITCH input=expr '{' (CASE expr ':' lsents)+ default? '}';

default: DEFAULT ':' lsents ;

print: PRINT '(' expr (',' expr)* ')';
