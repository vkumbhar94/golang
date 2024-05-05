# Benchmark testing

## install benchstat

```zsh
go install golang.org/x/perf/cmd/benchstat
```


## Running benchmarks

```zsh
go test -bench=. -benchmem
# to stop other tests from running
go test -bench=. -benchmem -run=^$

# to run specific benchmark
go test -bench=^BenchmarkAdd2$ -benchmem -run=^$

# to run specific benchmark with specific number of iterations
go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10

# to run specific benchmark with specific number of iterations and save the output to a file
go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10 > output.txt

# to run specific benchmark with specific number of iterations and save the output to a file and append the output

go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10 >> output.txt

# to run specific benchmark with specific number of iterations and save the output to a file and append the output and print the output to the console

go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10 | tee -a output.txt

# to set b.N to a specific value of time

go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10 -benchtime=10s

# to set b.N to a specific value of number of iterations

go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10 -benchtime=10x


# to set memory profiling

go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10 -benchtime=10x -memprofile memprofile.out

# to set cpu profiling

go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10 -benchtime=10x -cpuprofile cpuprofile.out

# to set block profiling

go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10 -benchtime=10x -blockprofile blockprofile.out

# to set mutex profiling

go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10 -benchtime=10x -mutexprofile mutexprofile.out

# to set trace profiling

go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10 -benchtime=10x -trace trace.out

# to set trace profiling with specific block size

go test -bench=^BenchmarkAdd2$ -benchmem -run=^$ -count=10 -benchtime=10x -trace trace.out -blockprofile blockprofile.out

```