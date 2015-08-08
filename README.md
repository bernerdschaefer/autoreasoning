# autoreason

Inspired by [Automated reasoning in F#, Scala, Haskell, C++, and Julia][phdb-auto-reasoning],
this is a small go program
which prints the result of simplifying the expression:

    e = (1 + 0 × x) × 3 + 12.

This is written in go as:

```go
var e Expr = Add{Mult{Add{Const(1), Mult{Const(0), Var("x")}}, Const(3)}, Const(12)}
```

Adjusting the expression reveals
what kind of expressions can be simplified to a constant
(a number is printed to the console).

  [phdb-auto-reasoning]: http://phdp.github.io/posts/2015-04-05-automated-reasoning.html

## License

This software is Copyright (c) 2015 Bernerd Schaefer.
It is free software, and may be redistributed
under the terms specified in the [LICENSE] file.

  [LICENSE]: LICENSE.md
