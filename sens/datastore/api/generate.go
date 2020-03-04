package main

import (
	"fmt"
	"go/importer"
	"log"
	"os"
	"text/template"

	"github.com/senslabs/alpha/sens/types"
)

var oms [][]string

func main() {
	log.SetFlags(log.Lshortfile)
	oms = [][]string{
		[]string{"sleep", "Sleep"},
		[]string{"device", "Device"},
	}
	fmt.Println(oms)
	for _, om := range oms {
		generateDb(om[0], om[1])
		generateRest(om[0], om[1])
	}
	generateMain()
}

func findModels() {
	p, err := importer.Default().Import("models")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p, p.Imports())
}

func generateDb(object string, model string) {
	t, err := template.ParseFiles("api/templates/db.tpl")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(fmt.Sprintf("api/db/%s.go", object))
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(f, types.Map{
		"Object": object,
		"Model":  model,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func generateRest(object string, model string) {
	t, err := template.ParseFiles("api/templates/rest.tpl")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(fmt.Sprintf("api/rest/main/%s.go", object))
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(f, types.Map{
		"Object": object,
		"Model":  model,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func generateMain() {
	t, err := template.ParseFiles("api/templates/main.tpl")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("api/rest/main/main.go")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(f, types.Map{
		"Models": []string{"Sleep", "Device"},
	})

	if err != nil {
		log.Fatal(err)
	}
}
