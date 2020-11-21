# lambda
An implementation of the untyped Lambda Calculus.

## The Lambda calculus

"[The Lambda calculus](https://en.wikipedia.org/wiki/Lambda_calculus) is a formal system in mathematical logic for 
expressing computation based on function abstraction and application using variable binding and substitution."

### Grammar

The grammar for the untyped Lambda calculus is fairly simple. Here it is in [EBNF](https://en.wikipedia.org/wiki/Extended_Backus–Naur_form):
```
variable    = "a" | "b" | ...  | "z";
application = "(", term, " ", term, ")";
abstraction = "(", "\", letter, ".", term, ")";
term        = variable | application | abstraction;
```

Examples for valid `terms` ("lambda-terms"):

```
v
(x y)
(\x.x)
(\x.(x y)
((\x.((\w.(w y)) x)) z)
```
