# lambda
_lambda_ is an implementation of the untyped Lambda Calculus.

## The Lambda calculus

"[The Lambda calculus](https://en.wikipedia.org/wiki/Lambda_calculus) is a formal system in mathematical logic for 
expressing computation based on function abstraction and application using variable binding and substitution."

### Grammar

The grammar for the untyped Lambda calculus is fairly simple. Here it is in [EBNF](https://en.wikipedia.org/wiki/Extended_Backus–Naur_form):
```
variable    = "a" | "b" | ... | "z";
application = term, " ", term;
abstraction = "λ", letter, ".", term;
term        = variable | "(", application, ")" | "(", abstraction, ")";
``` 

Examples for valid `terms` ("lambda-terms"):

```
v
(x y)
(λx.x)
(λx.(x y)
((\x.((\w.(w y)) x)) z)
```

## The _lambda_ language
### Grammar
We will make some small changes to the lambda calculus' grammar for our convenience.

First, we add a rule to allow omitting the outermost parentheses:
```
lambdaterm  = variable | application | abstraction | term;
```

Secondly, we make a small change to the `abstraction` rule:
```
abstraction = "\", letter, "." , ( abstraction | term );
```
This is allows us to write nested abstractions more easily, without having to use many parentheses, for example `\x.\y.(x y)`.

We also substitute the `λ` symbol for `\ ` to make it easier to type.

As a result, these are also valid terms in _lambda_, on top of the ones mentioned previously:

```
x y
\x.x
\x.(x y)
(x y) (x y)
\x.\y.(x y)
```

The final grammar now looks like this:

```
variable    = "a" | "b" | ... | "z";
application = term, " ", term;
abstraction = "\", letter, "." , ( abstraction | term );
term        = variable | "(", application, ")" | "(", abstraction, ")";
lambdaterm  = variable | application | abstraction | term;
```
