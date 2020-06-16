package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/phoenixking25/kubectl-mtb/benchmarks"
	"github.com/phoenixking25/kubectl-mtb/util"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type testOptions struct {
	tenant          string
	tenantNamespace string
	err             error
	pass            int
	fail            int
	kclient         *kubernetes.Clientset
	tclient         *kubernetes.Clientset
}

type groupResource struct {
	APIGroup    string
	APIResource metav1.APIResource
}

func init() {
	testCmd.Flags().StringP("namespace", "n", "", "Tenant namespace")
	testCmd.Flags().StringP("tenant-admin", "t", "", "Tenant service account name")
}

var (
	options = &testOptions{
		pass: 0,
		fail: 0,
	}

	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Run the Benchmarks against the K8s cluster",

		Run: func(cmd *cobra.Command, args []string) {
			options.validate(cmd, args)
			options.run()
		},
	}
)

func (t *testOptions) validate(cmd *cobra.Command, args []string) {
	t.tenant, _ = cmd.Flags().GetString("tenant-admin")
	if t.tenant == "" {
		color.Red("Error: Unable to get `namespace` from command-line")
		os.Exit(1)
	}

	t.tenantNamespace, _ = cmd.Flags().GetString("namespace")
	if t.tenantNamespace == "" {
		color.Red("Error: Unable to get `namespace` from command-line")
		os.Exit(1)
	}

	// Client with given config
	t.kclient, t.err = util.NewKubeClient()
	if t.err != nil {
		color.Red(t.err.Error())
		os.Exit(1)
	}

	// Check Namespace
	_, t.err = t.kclient.CoreV1().Namespaces().Get(context.TODO(), t.tenantNamespace, metav1.GetOptions{})
	if t.err != nil {
		color.Red(t.err.Error())
		os.Exit(1)
	}

	// Check Tenant as a SA
	_, t.err = t.kclient.CoreV1().ServiceAccounts(t.tenantNamespace).Get(context.TODO(), t.tenant, metav1.GetOptions{})
	if t.err != nil {
		color.Red(t.err.Error())
		os.Exit(1)
	}

	// Client representing tenant-admin SA
	t.tclient, t.err = util.ImpersonateWithUserClient(t.tenant, t.tenantNamespace)
	if t.err != nil {
		color.Red(t.err.Error())
		os.Exit(1)
	}

	resources := []util.GroupResource{
		{
			APIGroup: "",
			APIResource: metav1.APIResource{
				Name: "pods",
			},
		},
	}
	verbs := []string{"create"}

	for _, resource := range resources {
		for _, verb := range verbs {
			access, msg, err := util.RunAccessCheck(t.tclient, t.tenantNamespace, resource, verb)
			if err != nil {
				color.Red(t.err.Error())
				os.Exit(1)
			}
			if !access {
				color.Red(msg)
				os.Exit(1)
			}
		}
	}
}

func (t *testOptions) run() {
	color.HiBlue("=====%s %s=====", benchmarks.BenchmarkSuite.Title, benchmarks.BenchmarkSuite.Version)

	fmt.Println("")

	for _, b := range benchmarks.BenchmarkSuite.Benchmarks {
		result, err := b.Run(t.tenant, t.tenantNamespace, t.kclient, t.tclient)

		if result {
			t.pass = t.pass + 1

			color.HiGreen("[%s] %s Configured", b.ID, b.Title)
		} else {
			t.fail = t.fail + 1

			color.Red("[%s] %s Not configured", b.ID, b.Title)
			color.HiRed("Error: %s", err.Error())
			color.HiYellow("Remediation: %s", b.GetRemediation())
		}

		fmt.Println("")
	}
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("Passed: %s  Failed: %s\n", green(t.pass), red(t.fail))
}
