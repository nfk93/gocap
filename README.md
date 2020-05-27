Go to `output/` and run the following command to get the output.
```
 ../gocap github.com/nfk93/gocap/output ../parser/tests/success/test5.cgo
```

The first argument is the package path to `capchan`. 
The second argument must be a path to a code file with extension `.cgo`.
The output will be in current working directory.

Then run output code with
```
 go run test5.go
```