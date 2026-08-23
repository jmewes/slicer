grammar SpecGrammar;

// Minimal island/tolerant grammar for TypeScript Vitest-style spec files.
// Only the constructs we care about (describe / it / test / beforeEach /
// beforeAll / afterEach / afterAll), including their modifier chains
// (.skip / .only / .todo / .each / .for) and the each/for data-table double
// call, plus enough filler to keep parentheses and braces balanced.
//
// The Go files generated from this grammar are committed to the repo; do
// NOT hand-edit them.

program : element* EOF ;

element
    : suite
    | block
    | parenGroup
    | fillerToken
    ;

// A suite is a call to describe(...), it(...), test(...), beforeEach(...),
// beforeAll(...), afterEach(...) or afterAll(...). It may be preceded by a
// chain of modifiers (e.g. `.skip`, `.only`, `.todo`, `.each`, `.for`) and,
// for the each/for data-table form, a second parenthesised data group
// before the real argument list, e.g. `test.each([...])('...', fn)`.
suite
    : suiteKeyword modifier* dataGroup? LPAREN suiteArg* RPAREN
    ;

modifier
    : DOT IDENT
    ;

// The `([...])` data table of `test.each(...)` / `test.for(...)`.
dataGroup
    : parenGroup
    ;

suiteKeyword
    : DESCRIBE
    | IT
    | TEST
    | BEFOREEACH
    | AFTEREACH
    | BEFOREALL
    | AFTERALL
    ;

// Note: suiteArg intentionally omits raw LPAREN / RPAREN so the outer
// suite's closing `)` is never consumed as an argument. Any parenthesised
// filler (e.g. arrow-function parameter list) is matched via `parenGroup`.
suiteArg
    : STRING
    | block
    | suite
    | parenGroup
    | argFillerToken
    ;

// A `{ ... }` block. Elements inside are walked recursively so nested
// suites are discovered no matter how deeply they are wrapped.
block
    : LBRACE element* RBRACE
    ;

// A balanced `( ... )` group used as filler (for example an arrow function's
// empty or single-parameter list) or as the each/for data table. Contents
// are captured but ignored by the visitor.
parenGroup
    : LPAREN parenContent* RPAREN
    ;

parenContent
    : block
    | parenGroup
    | STRING
    | argFillerToken
    ;

fillerToken
    : STRING
    | IDENT
    | COMMA
    | ARROW
    | DOT
    | DESCRIBE
    | IT
    | TEST
    | BEFOREEACH
    | AFTEREACH
    | BEFOREALL
    | AFTERALL
    | OTHER
    ;

argFillerToken
    : IDENT
    | COMMA
    | ARROW
    | DOT
    | DESCRIBE
    | IT
    | TEST
    | BEFOREEACH
    | AFTEREACH
    | BEFOREALL
    | AFTERALL
    | OTHER
    ;

// -------- Lexer --------

DESCRIBE   : 'describe' ;
IT         : 'it' ;
TEST       : 'test' ;
BEFOREEACH : 'beforeEach' ;
AFTEREACH  : 'afterEach' ;
BEFOREALL  : 'beforeAll' ;
AFTERALL   : 'afterAll' ;

LPAREN : '(' ;
RPAREN : ')' ;
LBRACE : '{' ;
RBRACE : '}' ;
COMMA  : ',' ;
ARROW  : '=>' ;
DOT    : '.' ;

// String literals: single, double, or template quotes. We do not attempt to
// parse escape sequences deeply — just enough to skip over quoted content.
STRING
    : '\'' ( '\\' . | ~['\\\r\n] )* '\''
    | '"'  ( '\\' . | ~["\\\r\n] )* '"'
    | '`'  ( '\\' . | ~[`\\] )* '`'
    ;

// Line and block comments are routed to a hidden channel so the visitor can
// still find them by token index within a suite's block span.
LINE_COMMENT  : '//' ~[\r\n]*        -> channel(HIDDEN) ;
BLOCK_COMMENT : '/*' .*? '*/'        -> channel(HIDDEN) ;

WS : [ \t\r\n]+ -> skip ;

// Identifiers. Placed after the keyword tokens so `describe` / `it` /
// `test` / `beforeEach` / `afterEach` / `beforeAll` / `afterAll` match as
// their dedicated tokens. Combined identifiers such as `fit` / `xit` are
// matched as a single IDENT here, which naturally excludes them from suite
// recognition (removes the fragile isIdentChar hack).
IDENT : [a-zA-Z_$] [a-zA-Z_$0-9]* ;

// Catch-all for any other single character (punctuation, operators, digits
// not starting an identifier, etc.). Kept as a normal token — never skipped
// — so it can appear as filler in the parser rules.
OTHER : . ;
