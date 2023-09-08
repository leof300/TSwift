grammar TSParser_rules;
import TSLexer_rules;

start : lins EOF;

lins : ins*;

ins :   functions       #IFunc
    |   declar          #IDecl
    |   decons          #ICons
    |   callFunction    #SCallFunction
    |   tsprint           #IPrint
    ;

lsents : sents*;

sents : if                      #SIf
        | switch                #SSwitch
        | while                 #SWhile
        | guard                 #SGuard
        | for                   #SFor
        | declar                #SDecl
        | decons                #SCons
        | tsprint                 #SPrint
        | expr                  #SentExpr
        | strans                #SentTrans
     ;

expr :      <assoc=right> op='-' expr  #ENeg
        |   <assoc=right> op='!' expr  #ENot
        |   expr op='%' expr        #EModule
        |   expr op=('*'|'/')expr   #EMulDiv
        |   expr op=('+'|'-') expr  #EAddSub
        |   expr op=('=='|'!='|'>'|'<'|'>='|'<=') expr  #ERel
        |   tsfunctions             #ETSFunctions
        |   callFunction            #EFunction
//        |   arrayDef                #EVArray
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


declar :    VAR ID ':' (tstypes | arr='[' tstypes ']')  '?'             #SDType
        |   VAR ID '=' expr                                             #SDecAssign
        |   VAR ID ':' (tstypes | arr='[' tstypes ']')  op='='  expr     #SDecTAssign
       ;

//arrayDef : '[' (ID (','ID)* )? ']';
//
//getArray : ID '[' expr ']';

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

tsprint: PRINT '(' expr (',' expr)* ')';

functions: FUNC ID '(' ( parameter (',' parameter)* )? ')' ('->' tstypes)? '{' lsents '}' ;

parameter:  alias=(ID|'_')?  ID ':' INOUT? (tstypes | arr='[' tstypes ']') ;

callFunction: ID '(' ( callParameter (','callParameter)* )? ')';

callParameter: (ID':')? op='&'? expr ;

tsfunctions:  STRING '(' expr ')'       #ConvertString
            | INT '(' expr ')'          #ConvertInt
            | FLOAT '(' expr ')'        #ConvertFloat
            ;

