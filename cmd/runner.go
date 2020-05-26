package cmd

import (
	"fmt"
	"os"

	bcr "github.com/phoenixking25/kubectl-mtb/benchmarks/block_cluster_resources"
	bpc "github.com/phoenixking25/kubectl-mtb/benchmarks/block_privileged_containers"
)

var err error
var result bool

var PackageMap = map[string]func(string, string) (bool, error){
	"block_privileged_containers": bpc.Run,
	"block_cluster_resources":     bcr.Run,
}

func runBenchmarks() {
	if tenant == "" {
		fmt.Println("Please provide tenant-admin service account name")
		os.Exit(1)
	}
	if tenantNamespace == "" {
		fmt.Println("Please provide tenant namespace name")
		os.Exit(1)
	}

	Config, err := ReadConfig("/home/phoenix/GO/src/github.com/phoenixking25/kubectl-mtb/benchmarks/list.yaml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(Config.Title)
	for _, benchmark := range Config.Profile_Levels[0].Benchmarks {
		fmt.Println(benchmark.Id)
		result, err = PackageMap[benchmark.PKG](tenant, tenantNamespace)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(benchmark.Title + " configured")
		}
	}

}
