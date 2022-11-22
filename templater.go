package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type ConfigValues struct {
	Config map[string]interface{}
}

func loadConfigValues(valuesFilename string) (ConfigValues, error) {
	yfile, err := ioutil.ReadFile(valuesFilename)
	if err != nil {
		return ConfigValues{}, fmt.Errorf("opening config file: %w", err)
	}

	config := make(map[string]interface{})
	err = yaml.Unmarshal(yfile, &config)
	if err != nil {
		return ConfigValues{}, fmt.Errorf("parsing config file: %w", err)
	}

	return ConfigValues{config}, nil
}

func loadTemplate(templateFilename string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(templateFilename)
	if err != nil {
		return tmpl, fmt.Errorf("opening template file: %w", err)
	}
	return tmpl, nil
}

func createRenderedFile(renderedFilename string, tmpl *template.Template, values ConfigValues) error {
	file, err := os.Create(renderedFilename)
	if err != nil {
		return fmt.Errorf("creating target file: %w", err)
	}

	err = tmpl.Execute(file, values)
	if err != nil {
		return fmt.Errorf("applying config values to template: %w", err)
	}

	err = file.Close()
	if err != nil {
		return fmt.Errorf("closing target file: %w", err)
	}

	return nil
}

func main() {
	templateFilename := flag.String("template", "", "Path to template file")
	configFilename := flag.String("config", "", "Path to config file")
	targetFilename := flag.String("target", "result.yaml", "Path to target file")
	flag.Parse()

	tmpl, err := loadTemplate(*templateFilename)
	if err != nil {
		log.Fatal("loadTemplate: ", err)
	}
	values, err := loadConfigValues(*configFilename)
	if err != nil {
		log.Fatal("loadConfigValues: ", err)
	}
	err = createRenderedFile(*targetFilename, tmpl, values)
	if err != nil {
		log.Fatal("createRenderedFile: ", err)
	}
}
