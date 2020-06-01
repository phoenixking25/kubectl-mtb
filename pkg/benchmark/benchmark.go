package benchmark

import (
	"errors"

	"github.com/phoenixking25/kubectl-mtb/util"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/kubernetes"
)

// Benchmark struct
type Benchmark struct {
	ID            string `yaml:"id"`
	Title         string `yaml:"title"`
	BenchmarkType string `yaml:"benchmarkType"`
	Category      string `yaml:"category"`
	Description   string `yaml:"description"`
	Remediation   string `yaml:"remediation"`
	ProfileLevel  string `yaml:"profileLevel"`
	Run           func(string, string, *kubernetes.Clientset, *kubernetes.Clientset) (bool, error)
}

func (b *Benchmark) GetProfileLevel() string {
	return b.ProfileLevel
}

func (b *Benchmark) GetRemediation() string {
	return b.Remediation
}

func (b *Benchmark) ReadConfig(path string) error {
	file, err := util.LoadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(file, b); err != nil {
		return err
	}

	if b == nil {
		return errors.New("Please fill in a valid/non-empty yaml file")
	}
	return nil
}
