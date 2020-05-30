package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/phoenixking25/kubectl-mtb/benchmarks"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the Multi-Tenancy Benchmarks",

	Run: func(cmd *cobra.Command, args []string) {
		if benchmarks.BenchmarkSuite.GetTotalBenchmarks() == 0 {
			color.HiRed("No Benchmark found.")
			os.Exit(1)
		}
		color.HiWhite("Total Benchmarks: %d", benchmarks.BenchmarkSuite.GetTotalBenchmarks())

		for _, b := range benchmarks.BenchmarkSuite.Benchmarks {
			color.White("[%s] %s", b.ID, b.Title)
		}
	},
}
