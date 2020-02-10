package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"
)

const ChartFile = "Chart.yaml"

func check(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

type Chart struct {
	ApiVersion  string   `yaml:"apiVersion"`
	Name        string   `yaml:"name"`
	Version     string   `yaml:"version"`
	KubeVersion string   `yaml:"kubeVersion"`
	Description string   `yaml:"description"`
	Type        string   `yaml:"type"`
	Keywords    []string `yaml:"keywords"`
	Home        string   `yaml:"home"`
	Sources     []string `yaml:"sources"`
	Repository  string   `yaml:"repository"`
	Icon        string   `yaml:"icon"`
	AppVersion  string   `yaml:"appVersion"`
	Deprecated  bool     `yaml:"deprecated"`
}

func main() {
	chartPath := path.Join(os.Getenv("INPUT_PATH"), ChartFile)
	fmt.Println(fmt.Sprintf(`Reading values from %s`, chartPath))

	dat, readErr := ioutil.ReadFile(chartPath)
	check(readErr)

	chart := Chart{}
	yamlErr := yaml.Unmarshal([]byte(dat), &chart)
	check(yamlErr)

	fmt.Println(fmt.Sprintf(`::set-output name=apiVersion::%s`, chart.ApiVersion))
	fmt.Println(fmt.Sprintf(`::set-output name=name::%s`, chart.Name))
	fmt.Println(fmt.Sprintf(`::set-output name=version::%s`, chart.Version))
	fmt.Println(fmt.Sprintf(`::set-output name=kubeVersion::%s`, chart.KubeVersion))
	fmt.Println(fmt.Sprintf(`::set-output name=description::%s`, chart.Description))
	fmt.Println(fmt.Sprintf(`::set-output name=type::%s`, chart.Type))
	fmt.Println(fmt.Sprintf(`::set-output name=keywords::%s`, chart.Keywords))
	fmt.Println(fmt.Sprintf(`::set-output name=home::%s`, chart.Home))
	fmt.Println(fmt.Sprintf(`::set-output name=sources::%s`, strings.Join(chart.Sources[:], ",")))
	fmt.Println(fmt.Sprintf(`::set-output name=repository::%s`, chart.Repository))
	fmt.Println(fmt.Sprintf(`::set-output name=icon::%s`, chart.Icon))
	fmt.Println(fmt.Sprintf(`::set-output name=appVersion::%s`, chart.AppVersion))
	fmt.Println(fmt.Sprintf(`::set-output name=deprecated::%t`, chart.Deprecated))
}
