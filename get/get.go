// Command get downloads file from remote gir repository.
// cfg get ssh://git@gitlab.regium.com:33222/xr/configs/nginx/nginx.conf

package get

import (
	"errors"
	"flag"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/arvinkulagin/configurator/git"
)

const (
	_GIT_EXT = ".git"
)

var (
	ErrPath = errors.New("Wrong path to file in repo")
)

func Run() {
	_, err := exec.LookPath("git")
	if err != nil {
		log.Fatal(err)
	}
	fs := flag.NewFlagSet("get", flag.ExitOnError)
	// tag := fs.String("t", "", "Download file from tagged revision")
	fs.Parse(os.Args[2:])
	arg := strings.Join(fs.Args(), " ")

	u, err := url.Parse(arg)
	if err != nil {
		log.Fatal(err)
	}
	if u.Path == "" || u.Path == "/" {
		log.Fatal(ErrPath)
	}
	if len(strings.Split(u.Path, "/")) < 3 {
		log.Fatal(ErrPath)
	}
	repo := strings.Join(strings.Split(u.Path, "/")[:3], "/") + _GIT_EXT
	file := strings.Join(strings.Split(u.Path, "/")[3:], "/")
	u.Path = repo
	data, err := git.GetFile(u.String(), file)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(path.Base(file))
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}
