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
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Id            string `yaml:"id"`
	Title         string `yaml:"title"`
	BenchmarkType string `yaml:"benchmarkType"`
	Category      string `yaml:"category"`
	Description   string `yaml:"description"`
	Remediation   string `yaml:"remediation"`
	ProfileLevel  int64  `yaml:"profileLevel"`
}

func main() {
	var t conf
	data, _ := ioutil.ReadFile("test.yaml")
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	output := blackfriday.Run(data)
	fmt.Println(output)
	f, err := os.Create("test.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	n2, err := f.Write(output)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("--- t:\n%v\n\n", t)
	//cmd.Execute()
}
