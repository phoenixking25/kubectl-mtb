package cmd

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type BenchmarkConfig struct {
	Version        string
	Title          string
	Profile_Levels []Level
}

type Level struct {
	level      int
	title      string
	Benchmarks []Benchmark
}

type Benchmark struct {
	Id             string
	Title          string
	Benchmark_type string
	Catefory       string
	Remediation    string
	Description    string
	PKG            string
}

func ReadConfig(path string) (*BenchmarkConfig, error) {
	var Config *BenchmarkConfig

	file, err := LoadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(file, &Config); err != nil {
		return nil, err
	}

	if Config == nil {
		return Config, errors.New("Please fill in a valid/non-empty yaml file")
	}
	return Config, nil
}

func LoadFile(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	return ioutil.ReadFile(path)
}
