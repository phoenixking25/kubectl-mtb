package benchmarksuite

import (
	"sort"
	"strings"

	"github.com/phoenixking25/kubectl-mtb/pkg/benchmark"
)

// BenchmarkSuite - Collection of benchmarks
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

	sort.Slice(bs.Benchmarks, func(j, i int) bool {

		profileLevelI := bs.Benchmarks[i].ProfileLevel
		profileLevelJ := bs.Benchmarks[j].ProfileLevel
		orderNoI := bs.Benchmarks[i].ID[len(bs.Benchmarks[i].ID)-1:]
		orderNoJ := bs.Benchmarks[j].ID[len(bs.Benchmarks[j].ID)-1:]

		switch strings.Compare(profileLevelI, profileLevelJ) {
		case -1:
			return true
		case 1:
			return false
		}
		return orderNoI > orderNoJ
	})

	return
}
