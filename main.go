package main

import (
	"flag"
	"log"
	"os"
	"text/template"

	"github.com/arvinkulagin/configurator/consul"
)

var (
	path      = flag.String("tpl", "", "Template file")
	outgoing  = flag.String("result", "", "Result configuration file name")
	addr      = flag.String("consul", "", "Consul node address")
	prefix    = flag.String("key", "", "Consul config key (may be complex: 'user/test/mongodb')")
	asAnsible = flag.Bool("ansible", false, "Use as Ansible module")
)

func main() {
	flag.Parse()

	tmpl, err := template.ParseFiles(*path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	result, err := os.Create(*outgoing)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer result.Close()

	cons, err := consul.New(*addr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	data, err := cons.Get(*prefix)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = tmpl.ExecuteTemplate(result, *path, data)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
