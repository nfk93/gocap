Run the following command to get the output.
```
 ./gocap ./output
```

The argument must be a path to the directory where all `.cgo` files are.
The output `.go` file will be next to the corresponding `.cgo` file.

Then run output code with
```
 go run ./output/test5.go
```