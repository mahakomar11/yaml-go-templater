package main

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type ConfigValues struct {
	Config map[string]interface{}
}

func loadConfigValues(valuesFilename string) ConfigValues {
	yfile, err := ioutil.ReadFile(valuesFilename)
	if err != nil {
		log.Fatal("Opening config file: ", err)
	}

	config := make(map[string]interface{})
	err2 := yaml.Unmarshal(yfile, &config)
	if err2 != nil {
		log.Fatal("Parsing config file: ", err2)
	}

	return ConfigValues{config}
}

func loadTemplate(templateFilename string) *template.Template {
	tmpl, err := template.ParseFiles(templateFilename)
	if err != nil {
		log.Fatal("Opening template file: ", err)
	}
	return tmpl
}

func createRenderedFile(renderedFilename string, tmpl *template.Template, values ConfigValues) {
	file, err := os.Create(renderedFilename)
	if err != nil {
		log.Fatal("Creating target file: ", err)
	}

	err = tmpl.Execute(file, values)
	if err != nil {
		log.Fatal("Applying config values to template: ", err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal("Closing target file:", err)
	}
}

func main() {
	templateFilename := flag.String("template", "", "Path to template file")
	configFilename := flag.String("config", "", "Path to config file")
	targetFilename := flag.String("target", "result.yaml", "Path to target file")
	flag.Parse()

	tmpl := loadTemplate(*templateFilename)
	values := loadConfigValues(*configFilename)
	createRenderedFile(*targetFilename, tmpl, values)
}
