# Benchmark test

## Blog
[Builder Pattern with Zero Alloc in Golang](https://blog.devgenius.io/builder-pattern-with-zero-alloc-in-golang-3c04365c62fe)

```zsh
go test -bench=. -benchmem
```

# Benchmark result

```zsh
goos: darwin
goarch: amd64
pkg: github.com/vkumbhar94/golang/cmd/zeroalocbuilder
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkFooBuilder-12          1000000000               0.2810 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/vkumbhar94/golang/cmd/zeroalocbuilder        0.508s
```