package benchmarks

import (
	"github.com/phoenixking25/kubectl-mtb/benchmarks/block_cluster_resources"
	"github.com/phoenixking25/kubectl-mtb/benchmarks/block_privileged_containers"
	"github.com/phoenixking25/kubectl-mtb/benchmarks/require_run_as_non_root"
	benchmarksuite "github.com/phoenixking25/kubectl-mtb/pkg/benchmarks_suite"
)

var BenchmarkSuite = &benchmarksuite.BenchmarkSuite{
	Version: "1.0.0",
	Title:   "Multi-Tenancy Benchmarks",
}

func init() {
	BenchmarkSuite.AddBenchmark(block_privileged_containers.BPCbenchmark)
	BenchmarkSuite.AddBenchmark(block_cluster_resources.BCRbenchmark)
	BenchmarkSuite.AddBenchmark(require_run_as_non_root.RRANRbenchmark)

	// Sort according to ID and ProfileLevel
	BenchmarkSuite.SortBenchmarks()
}
