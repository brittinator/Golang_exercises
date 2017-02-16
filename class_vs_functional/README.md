# Benchmarking class type methods versus functions


In basic functions, there seems to be no difference between writing class methods, and
just writing the functions and passing in parameters.   

```
go test -bench=Method
1000000	      1795 ns/op
PASS
ok  	1.816s
```

````
go test -bench=Function
1000000	      1652 ns/op
PASS
ok	1.679s
```
