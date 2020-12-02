# lambda
_lambda_ is an implementation of the untyped Lambda Calculus.

## The Lambda calculus

"[The Lambda calculus](https://en.wikipedia.org/wiki/Lambda_calculus) is a formal system in mathematical logic for 
expressing computation based on function abstraction and application using variable binding and substitution."

### Grammar

The grammar for the untyped Lambda calculus is fairly simple. Here it is in [EBNF](https://en.wikipedia.org/wiki/Extended_Backus–Naur_form):
```
variable    = "v" | variable, "'";
application = term, " ", term;
abstraction = "λ", letter, ".", term;
term        = variable | "(", application, ")" | "(", abstraction, ")";
``` 

Examples for valid `terms` ("lambda-terms"):

```
v
(v v')
(λv.v)
(λv.(v v')
((λv.((λv''.(v'' v')) v)) v''')
```

## The _lambda_ language
### Grammar
We make some small changes to the lambda calculus' grammar for our convenience:

First, we replace the variable rule by the following one:
```
variable = "a" | "b" | ... | "z";
```
While the original rule technically allowed for an infinite number of variables, 
we very likely won't run into issues limiting ourselves to just 26 variable names.

Secondly, we add a rule to allow omitting the outermost parentheses:
```
lambdaterm  = variable | application | abstraction | term;
```

Now, we make a small change to the `abstraction` rule:
```
abstraction = "\", letter, "." , ( abstraction | term );
```
This is allows us to write nested abstractions more easily, without having to use many parentheses, for example `λx.λy.(x y)`.

Finally, also substitute the `λ` symbol for `\ ` to make it easier to type.

As a result, these are also valid terms in _lambda_:

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
