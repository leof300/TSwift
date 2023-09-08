grammar TSParser_rules;
import TSLexer_rules;

start : lins EOF;

lins : ins*;

ins :   functions       #IFunc
    |   declar          #IDecl
    |   decons          #ICons
    |   callFunction    #SCallFunction
    ;

lsents : sents*;

sents : if                      #SIf
        | switch                #SSwitch
        | while                 #SWhile
        | guard                 #SGuard
        | for                   #SFor
        | declar                #SDecl
        | decons                #SCons
        | print                 #SPrint
        | expr                  #SentExpr
        | strans                #SentTrans
     ;

expr :      <assoc=right> op='-' expr  #ENeg
        |   <assoc=right> op='!' expr  #ENot
        |   expr op='%' expr        #EModule
        |   expr op=('*'|'/')expr   #EMulDiv
        |   expr op=('+'|'-') expr  #EAddSub
        |   expr op=('=='|'!='|'>'|'<'|'>='|'<=') expr  #ERel
        |   callFunction            #EFunction
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

        |   NIL                     #ENIL
        ;


declar :    VAR ID ':' tstypes  '?'             #SDType
        |   VAR ID '=' expr                     #SDecAssign
        |   VAR ID ':' tstypes op='='  expr     #SDecTAssign
       ;

decons :    LET ID '=' expr                     #SConsAss
        |   LET ID ':' tstypes op='='  expr     #SConsTypeAss
        ;

tstypes : STRING | INT | FLOAT | BOOL | CHARACTER ;

block : term='{' lsents '}';

if :  IF  expr block            #RIf
    | IF  expr block ELSE block #RIfElse
    | IF  expr block ELSE if    #RIfEIf
;

switch: SWITCH input=expr '{' (CASE expr ':' lsents)+ default? '}';

default: DEFAULT ':' lsents ;

while: WHILE expr block;

guard: GUARD expr ELSE block;

for:    FOR ID IN (expr|rango) block;

rango:  VINTEGER '..' VINTEGER;

strans:   CONTINUE
        | BREAK
        | RETURN expr?
        ;

print: PRINT '(' expr (',' expr)* ')';

functions: FUNC ID '(' ( parameter (',' parameter)* )? ')' ('->' tstypes)? '{' lsents '}' ;

parameter:  alias=(ID|'_')?  ID ':' INOUT? tstypes;

callFunction: ID '(' ( callParameter (','callParameter)* )? ')';

callParameter: (ID':')? op='&'? expr ;
