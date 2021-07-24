# Go Learning Path

## About Go
Fans of Go (called gophers) describe Go as having the expressiveness of dynamic languages like Python or Ruby, with the performance of compiled languages like C or C++.

The language is open source, and was started by engineers at Google. It's written using a C-style syntax, has statically typed variables, manages memory using garbage collection, and is compiled into stand-alone executables.

Go is noted for the concurrent programming features built into the language core, the networking packages in the standard library (such as a web server), fast compilation and execution speed. Its simple, minimalistic and consistent language design make for a delightful experience, while the abundant and thoughtful tooling addresses traditional problems such as consistent formatting and documentation.

The home page for Go is golang.org, and there is an excellent interactive tutorial at tour.golang.org.

## Run Go

#### Run your code
```
go run file_name.go
```

#### Run your tets
```
go test
```

##### BenchmarkXxx()

BenchmarkXxx() is a benchmarking function. These functions follow the form `func BenchmarkXxx(*testing.B)` and can be used to test the performance of your implementation. They may not be present in every exercise, but when they are you can run them by including the `-bench` flag with the `go test` command, like so: 

```
go test -v --bench . --benchmem
```

You will see output similar to the following:
```
BenchmarkXxx   	2000000000	         0.46 ns/op
```
This means that the loop ran 2000000000 times at a speed of 0.46 ns per loop.

While benchmarking can be useful to compare different iterations of the same exercise, keep in mind that others will run the same benchmarks on different machines, with different specs, so the results from these benchmark tests may vary.
