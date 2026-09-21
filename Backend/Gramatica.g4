grammar Gramatica;


// Reglas Gramaticales


program: block EOF;

block: statement*;

statement
    : mainfunction
    | declaration
    | assignment
    | print
    | if
    | switch
    | for
    | individualBlock
    | transferStatemen
    | call
    ;

// --------------------------
// Funcion Main
mainfunction
    : FUNCTION MAIN '(' ')' '{' block '}'
    ;

// --------------------------
// |Declaracion|
declaration
    : MUT? ID primitiveType '=' expression ';'? # DeclarationExplicitVar
    | MUT? ID primitiveType ';'? # DeclarationImplicitVar
    | MUT? ID '=' '[]' primitiveType '{' listSlice '}' ';'? # DeclarationExplicitSlice
    | MUT? ID '[]' primitiveType ';'? # DeclarationImplicitSlice
    | 'struct' ID '{' struct_field* '}' ';'? # StructDeclaration
    | ID '=' ID '{' lstStruct (',' lstStruct)* '}' ';'? # StructInstanceDeclaration
    | FUNCTION ID '(' parameters? ')' (primitiveType|tiposSlice)? '{' block '}' ';'? # FunctionDeclaration
    ;


 // Parametros de la funcion
parameters
    : parameter (',' parameter)*
    ;
parameter
    : ID primitiveType # NormalTypeParameter
    | ID '[]' primitiveType # SliceTypeParameter
    ;

tiposSlice
    : '[]' primitiveType
    ;

//struct_dcl:
//    'struct' ID '{' struct_field* '}' ';'? # StructDeclaration;

struct_field
    : primitiveType ID ';'? 
    //| ID '[]' primitiveType ';'? # StructFieldSliceDeclaration
    ;    


lstStruct
    : ID ':' expression 
    ;
/*struct_field_init
    : ID ':' expression  # StructFieldInit
    //| ID '=' '[]' primitiveType '{' listSlice '}' ';'? # StructFieldSliceInit
    //| ID '[]' primitiveType ';'? # StructFieldSliceDeclaration
    ;
*/

listSlice
    : expression (',' expression)* 
    ;

// --------------------------
// |Asignacion|
assignment
    : ID '=' expression ';'? # AssignmentExpr
    | ID '+=' expression ';'? # AssignmentPlusExpr
    | ID '-=' expression ';'? # AssignmentMinusExpr
    | ID '++' ';'? # AssignmentIncrementExpr
    | ID '--' ';'? # AssignmentDecrementExpr
    | ID '[' expression ']' '=' expression ';'? # SliceAssignmentExpr
    | expression '=' expression ';'? # AttributeAssignmentExpr
    ;
// --------------------------           

// |Print|
print
    : PRINT '(' printList? ')' ';'? 
    ;

printList
    : expression (',' expression)*
    ;

// --------------------------
// |If|
if
    : IF expression '{' block '}' elseIfBlock* elseBlock?  
    ;

elseIfBlock
    : ELSE IF expression '{' block '}'
    ;

elseBlock
    : ELSE '{' block '}'
    ;

// --------------------------
// |Switch|   
switch
    : SWTCH expression '{' caseBlock+ defaultBlock? '}'
    ;

caseBlock
    : CASE expression ':' block
    ;

defaultBlock
    : 'default' ':' block
    ;

// --------------------------

// |For|
for
    : FOR expression '{' block '}' #ForSimple
    | FOR declaration ';' expression ';' assignment? '{' block '}' #ForDeclaration
    | FOR ID ',' ID 'in' expression '{' block '}' #ForInSlice
    ;

// --------------------------

// Bloque Individual
individualBlock
    : '{' block '}'
    ;

// --------------------------

// Sentencias de Transferencia
transferStatemen
    : TRANSFER_STATEMENT_BREAK ';'? # BreakStatement
    | TRANSFER_STATEMENT_CONTINUE ';'? # ContinueStatement
    | TRANSFER_STATEMENT_RETURN expression? ';'? # ReturnStatement
    ;


// --------------------------
// Llamadas
call
    : ID '(' arguments? ')' ';'? # FunctionCall
    ;

arguments
    : expression (',' expression)*
    ;

// --------------------------

// |Tipos Primitivos|

primitiveType
    : PRIMITIVE_INT 
    | PRIMITIVE_FLOAT 
    | PRIMITIVE_BOOL 
    | PRIMITIVE_STRING 
    ;
    
// --------------------------

// Funciones de los Slice
// Pueden ir en expresiones porque me retornan valores primitivos como int, string, bool, f64
slicePrimitiveValueFunction
    : INDEXOF '(' ID ',' expression ')' # IndexOfFunction
    | JOIN '(' ID ',' expression ')' # JoinFunction
    | LENGTH '(' ID ')' # LengthFunction
    ;

// --------------------------


// |Expresiones|

expression
    : '!' right=expression                        # NotExpr
    | '-' right=expression                        # NegExpr 
    | '&' right=ID                        # RefExpr     
    | '*' right=ID                        # DerefExpr
    | expression '.' ID #StructAccess
    | left=expression op=('*'|'/'|'%') right=expression     # OpExpr
    | left=expression op=('+'|'-') right=expression     # OpExpr
    | left=expression op=('>='|'>') right=expression    # OpExpr
    | left=expression op=('<='|'<') right=expression    # OpExpr
    | left=expression op=('=='|'!=') right=expression   # OpExpr
    | left=expression op='&&' right=expression             # OpExpr
    | left=expression op='||' right=expression             # OpExpr
    | '(' expression ')' # ParExpr
    | slicePrimitiveValueFunction # SlicePrimitiveValue
    | APPEND '(' ID ',' expression ')'  # AppendExpr
    | ID '[' expression ']' # SliceAccess
    | ID '(' arguments? ')' # FunctionCallExpr 
    | ATOI '(' expression  ')' # AtoiExpr
    | PARSEFLOAT '(' expression  ')' # ParseFloatExpr
    | TYPEOF '(' expression  ')' # TypeOfExpr
    | ID # IdExpr
    | INT # IntExpr
    | FLOAT # FloatExpr
    | BOOL # BoolExpr
    | STRING # StrExpr
    ;
// --------------------------

// --------------
// Reservadas
// --------------

MAIN: 'main';
FUNCTION: 'fn';
MUT: 'mut';
PRINT: 'println';
IF: 'if';
ELSE: 'else';
SWTCH: 'switch';
CASE: 'case';
FOR: 'for';
TRANSFER_STATEMENT_BREAK: 'break';
TRANSFER_STATEMENT_CONTINUE: 'continue';
TRANSFER_STATEMENT_RETURN: 'return';
PRIMITIVE_INT: 'int';
PRIMITIVE_FLOAT: 'f64';
PRIMITIVE_BOOL: 'bool';
PRIMITIVE_STRING: 'string';
BOOL    : 'true' | 'false';
// --------------
// SOlo para slices
INDEXOF: 'indexOf';
JOIN: 'join';
LENGTH: 'len';
APPEND: 'append';
// --------------
ATOI: 'Atoi';
PARSEFLOAT: 'parseFloat';
TYPEOF: 'typeOf';

// --------------
// TOKENS
// --------------

ID      : [a-zA-Z_][a-zA-Z0-9_]* ;
FLOAT   : [0-9]+ '.' [0-9]* | '.' [0-9]+ ;
INT     : [0-9]+ ;
STRING  : '"' (ESC_SEQ | ~["\\\r\n])* '"' ;
//STRING  : '"' (~["\r\n] | '""')* '"' ;


fragment ESC_SEQ
    : '\\' [btnr"\\]   // permite: \b \t \n \r \" \\
    ;

// ---------------------
// ESPACIOS Y COMENTARIOS
// ---------------------

LINE_COMMENT    : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT   : '/*' .*? '*/' -> skip ;
WS              : [ \t\r\n]+ -> skip ;