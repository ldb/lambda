# lambda
_lambda_ is an implementation of the untyped Lambda Calculus.

## The Lambda calculus

"[The Lambda calculus](https://en.wikipedia.org/wiki/Lambda_calculus) is a formal system in mathematical logic for 
expressing computation based on function abstraction and application using variable binding and substitution."

### Grammar

The grammar for the untyped Lambda calculus is fairly simple. Here it is in [EBNF](https://en.wikipedia.org/wiki/Extended_Backus–Naur_form):
```
variable    = "a" | "b" | ...  | "z";
application = term, " ", term;
abstraction = "\", letter, ".", term;
term        = variable | "(", application, ")" | "(", abstraction, ")";
``` 

Examples for valid `terms` ("lambda-terms"):

```
v
(x y)
(\x.x)
(\x.(x y)
((\x.((\w.(w y)) x)) z)
```

## The lambda language
### Grammar  
We will deviate a small bit from the lambda calculus' grammar by allowing the outermost parentheses to be omitted by 
adding the following rule to the grammer:
```
lambdaterm = variable | application | abstraction | term;
``` 

As a result, these are also valid terms in _lambda_, on top of the ones mentioned previously:

```
x y
\x.x
\x.(x y)
(x y) (x y)
```

The final grammar now looks like this:

```
variable    = "a" | "b" | ...  | "z";
application = term, " ", term;
abstraction = "\", letter, ".", term;
term        = variable | "(", application, ")" | "(", abstraction, ")";
lambdaterm  = variable | application | abstraction | term;
```