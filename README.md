# TSwift
Proyecto compiladores crado con 

| GO  | ANTLR  |  ANGULAR  |
|--------------|--------------|--------------|
| go version go1.21.0 windows/amd64| ANTLR Parser Generator  Version 4.13.1||

**Índice**   
1. [Primer apartado](#i1)
2. [Segundo apartado](#i2)
3. [Segundo apartado](#i3)
4. [Segundo apartado](#i4)


## Requisitos <a name="i1"></a>
Es útil tener conocimientos previos de programación y comprensión de conceptos de lenguajes formales y gramáticas. 

- Tener instalado antlr, en especial la herramienta de `antlr-parse`.
 - Tener instalado go (golang)
 - Angular
 - IDE de preferencia, GoLand de preferencia 

## Instalación <a name="i2"></a>

### ANTLR
[ANTLR4](https://www.github.com](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md#windows-specific-issues)
[Prblemas con Windows](https://www.github.com](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md#windows-specific-issues)

*NOTA* válida que la versión de java sea la correcta, este instalada y concuerde con la última aceptada por antlr4.

## Iniciar Proyecto <a name="i3"></a>


## Estructura Proyecto GO <a name="i4"></a>
1. **ast**
	1. `NTExpression` : todas las estructuras que implementar la interfaz Expressioner, estos son las instrucciones del lenguaje, los NoTerminales.
 	2. `TSStruts` : contiene las estructuras básicas para nuestro analizador, valores primitivos, tabla de simbolos/scope, contexto.
3. `Grammar` : gramática y archivos ejemplo
4. `Services` : servicios para el proyecto
5. `Interpreter` : contiene las estructuras y metódos para procesar la gramática e la implementación del modelo VISITOR.
 
![Sin título](https://github.com/leof300/TSwift/assets/69184023/a34e323e-4a67-463c-9b43-6190ea2955bc)


## Estructura Proyecto ANGULAR <a name="i4"></a>


## Grámatica ANTRL <a name="i5"></a>
ANTLR (ANother Tool for Language Recognition) es una potente herramienta para el análisis y procesamiento de lenguajes. Diseñada para ayudar a los desarrolladores a construir analizadores léxicos y sintácticos eficientes. Utiliza un enfoque descendente (top-down), recursivos predictivos (LL(*)), para generar analizadores sintácticos. Esto significa que ANTLR comienza con una gramática que representa las reglas sintácticas de un lenguaje y utiliza estas reglas para construir un analizador sintáctico descendente.

### Gramatica Lexer
```python
/*
 * Lexer Rules
 */
lexer grammar TSLexer_rules;

NL: '\r'? '\n' -> skip;
WS: [ \t]+ -> skip ;


VAR: 'var';
LET: 'let';

NIL: 'nil';

STRING: 'String';
INT: 'Int';
FLOAT: 'Float';
BOOL: 'Bool';
CHARACTER: 'Character';

IF: 'if';
ELSE: 'else';
SWITCH: 'switch';
DEFAULT:'default';
CASE:'case';
WHILE: 'while';
IN: 'in';
FOR: 'for';
GUARD:'guard';
CONTINUE:'continue';
BREAK:'break';
RETURN:'return';
PRINT: 'print';

STRUCT:'struct';
SELF:'self';
MUTATING:'mutating';
FUNC:'func';
INOUT:'inout';
/*
APPEND:'append';
REMOVELAST:'removeLast';
REMOVE:'remove';
ISEMPTY:'IsEmpty';
COUNT:'count';
*/

fragment DIGIT : [0-9] /*{System.out.println("found an end");}*/ ;

VBOOL:   'true'|'false' ;
VSTRING : '"' (ESC|.)*? '"';
VFLOAT:   [0-9]+ '.' [0-9]+ ;
VINTEGER: [0-9]+ ;

//SINGLE : '\'' .*? '\'' -> type(STRING) ;
ID: [a-zA-Z_] [a-zA-Z0-9_]* ; // match usual identifier spec


//SL_COMMENT: '//' .*?  '\r'? '\n' -> type(NL);
SL_COMMENT: '//' .*?  ~[\r\n]* -> skip;
ML_COMMENT: '/*' .*? '*/' -> skip;

fragment ESC : '\\"' | '\\\\' ; // 2-char sequences \" and \\

//LINE_ESCAPE : '\\' '\r'? '\n' -> skip;
```
### Gramatica Parser

```python
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
        |   arrayValue                #EVArray
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

arrayValue : '[' (expr (','expr)* )? ']';
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
```

## Archivos de Prueba
- [Operaciones](http://www.ejemplo.com)
- [Funciones](http://www.ejemplo.com)
- [Estructuras de Control](http://www.ejemplo.com)






