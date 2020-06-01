module github.com/phoenixking25/kubectl-mtb

go 1.13

require (
	github.com/fatih/color v1.9.0
	github.com/google/uuid v1.1.1
	github.com/heketi/heketi v9.0.0+incompatible
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.3.2
	gopkg.in/yaml.v2 v2.2.8

	k8s.io/api v0.18.2
	k8s.io/apiextensions-apiserver v0.18.2 // indirect
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/kubernetes v1.16.2
	k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20191003000013-35e20aa79eb8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191028232452-c47e10e6d5a3
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191030190112-bb31b70367b7
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20191003001037-3c8b233e046c
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20191016114015-74ad18325ed5
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191028230319-1a481fb1e32d
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191025232453-66dd06a864dd
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20191016115129-c07a134afb42
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20191004115455-8e001e5d1894
	k8s.io/component-base => k8s.io/component-base v0.0.0-20191029070825-5e0e35147053
	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20191025232916-446748cffdda
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20191003003551-0eecdcdcc049
	k8s.io/klog => k8s.io/klog v1.0.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20191016112429-9587704a8ad4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20191016114939-2b2b218dc1df
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190918143330-0270cf2f1c1d
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191016114407-2e83b6f20229
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20191016114748-65049c67a58b
	k8s.io/kubectl => k8s.io/kubectl v0.0.0-20191031072635-2ba9448df4cc
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20191016114556-7841ed97f1b2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20191016115753-cf0698c3a16b
	k8s.io/metrics => k8s.io/metrics v0.0.0-20191016113814-3b1a734dba6e
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20191016112829-06bb3c9d77c9
	k8s.io/utils => k8s.io/utils v0.0.0-20191030222137-2b95a09bc58d
)
