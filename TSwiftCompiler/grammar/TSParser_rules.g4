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