/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var profileLevel int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-mtb",
	Short: "Multi-Tenancy Benchmarking",
	Long: `Command line tool to validate the multi-tenancy
in a Kubernetes cluster by running different benchmarks
to check the configurations and privileges set for tenants.`,

	// Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().IntVar(&profileLevel, "pl", 3, "Profile Level of the Benchmark.")

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(testCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {}
