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