package fill

import (
	"errors"
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
	"text/template"

	"github.com/arvinkulagin/configurator/consul"
)

var (
	ErrEmptyFlag = errors.New("You must specify -c and -o flags")
)

func Run() {
	fs := flag.NewFlagSet("fill", flag.ExitOnError)
	consulLink := fs.String("c", "", "Fill template with Consul data")
	fileName := fs.String("o", "", "File to save result config")
	fs.Parse(os.Args[2:])
	if *consulLink == "" {
		log.Fatal(ErrEmptyFlag)
	}
	arg := strings.Join(fs.Args(), " ")

	tmpl, err := template.ParseFiles(arg)
	if err != nil {
		log.Fatal(err)
	}

	if !strings.HasPrefix(*consulLink, "http://") {
		*consulLink = "http://" + *consulLink
	}
	consulURL, err := url.Parse(*consulLink)
	if err != nil {
		log.Fatal(err)
	}
	cns, err := consul.New(consulURL.Host)
	if err != nil {
		log.Fatal(err)
	}
	data, err := cns.Get(consulURL.Path)
	if err != nil {
		log.Fatal(err)
	}
	output, err := os.Create(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	err = tmpl.Execute(output, data)
	if err != nil {
		log.Fatal(err)
	}
}
