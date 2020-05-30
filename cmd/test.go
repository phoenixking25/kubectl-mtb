package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/phoenixking25/kubectl-mtb/benchmarks"
	"github.com/spf13/cobra"
)

func init() {
	testCmd.Flags().StringP("namespace", "n", "", "Tenant namespace")
	testCmd.Flags().StringP("tenant-admin", "t", "", "Tenant service account name")
}

var (
	tenant          string
	tenantNamespace string
	err             error

	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Run the Benchmarks against the K8s cluster",

		Run: func(cmd *cobra.Command, args []string) {

			tenant, err = cmd.Flags().GetString("tenant-admin")
			if err != nil {
				color.Red("Error: Unable to get `tenant-admin` from command-line :%v", err.Error())
				os.Exit(1)
			}

			tenantNamespace, err = cmd.Flags().GetString("namespace")
			if tenantNamespace == "" {
				color.Red("Error: Unable to get `namespace` from command-line :%v", err.Error())
				os.Exit(1)
			}

			color.HiBlue("%s %s", benchmarks.BenchmarkSuite.Title, benchmarks.BenchmarkSuite.Version)

			for _, b := range benchmarks.BenchmarkSuite.Benchmarks {
				result, err := b.Run(tenant, tenantNamespace)

				if result {
					color.HiGreen("[%s] %s Configured", b.ID, b.Title)
				} else {
					color.Red("[%s] %s Not configured", b.ID, b.Title)
					color.HiRed("Error: %s", err.Error())
					color.HiYellow("Remediation: %s", b.GetRemediation())
				}

				fmt.Println("")
			}
		},
	}
)
