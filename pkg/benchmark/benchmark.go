package benchmark

import (
	"errors"

	"github.com/phoenixking25/kubectl-mtb/util"
	"gopkg.in/yaml.v2"
)

// Benchmark struct
type Benchmark struct {
	ID            string
	Title         string
	BenchmarkType string
	Category      string
	Remediation   string
	Description   string
	ProfileLevel  string
	Run           func(string, string) (bool, error)
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
