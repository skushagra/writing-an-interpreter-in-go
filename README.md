# Interpreter in GoLang

## About
This repo builds on instructions given in the book `Writing an interpreter in Go`. Building a simple interpreter for a language called `Monkey` which is introduced in the book.

## Notes
I have taken the following notes while reading the book.
Components of an interpreter

1. Lexer
2. Parser
3. Abstract Syntax Tree
4. Internal Object System
5. Evaluater

# Lexing

## Lexical Analysis

We transform the code twice before execution, from `Source Code → Tokens → AST.`

The process of transformation from `Source Code → Tokens` is called lexing or lexical analysis.

Its done by a lexer, also called tokenizer or scanner.

Tokens are small, categorizable data structures that are then fed to a parser which transforms it into AST.

Input given to the lexer

```python
let x = 5 + 5;
```

The lexer output

```python
[
LET,
IDENTIFIER("x"),
EQUAL_SIGN,
INTEGER(5),
PLUS_SIGN,
INTEGER(5),
SEMICOLON
]
```

Some parsers convert the 5 to an integer in later stages like parser or even further, not during lexing, where it is kept as an integer.

## Defining our tokens

We have two extra types `ILLEGAL` and `EOF` . Illegal stands for something we do not know about and EOF is End Of File which tells the parser later to stop parsing the beyond this point.

## The Lexer
