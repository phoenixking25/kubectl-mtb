package benchmarksuite

import (
	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
)

// Collection of Benchmarks
type BenchmarkSuite struct {
	Version    string
	Title      string
	Benchmarks []*benchmark.Benchmark
}

func (bs *BenchmarkSuite) GetTotalBenchmarks() int {
	return len(bs.Benchmarks)
}

func (bs *BenchmarkSuite) AddBenchmark(benchmark *benchmark.Benchmark) {
	bs.Benchmarks = append(bs.Benchmarks, benchmark)
}

func (bs *BenchmarkSuite) SortBenchmarks() {
	if bs.GetTotalBenchmarks() <= 1 {
		return
	}

	// split := strings.Split(str, "-")
	// profileLevel = split[1][len(val)-1:]
	// orderLevel =
	return
}
